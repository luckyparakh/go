package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Device struct {
	ID       int    `json:"id"`
	Mac      string `json:"mac"`
	Firmware string `json:"firmware"`
}
type metrics struct {
	devices       prometheus.Gauge
	info          *prometheus.GaugeVec
	upgrades      *prometheus.CounterVec
	duration      *prometheus.HistogramVec
	loginDuration prometheus.Summary
}

func NewMetrics(reg prometheus.Registerer) *metrics {
	m := &metrics{
		devices: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: "myApp",
			Help:      "Number of connected devices",
			Name:      "connected_devices",
		}),
		info: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "myApp",
			Help:      "Info about my app",
			Name:      "info",
		}, []string{"version"}),
		upgrades: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "myApp",
			Help:      "Number of device Upgrade",
			Name:      "device_upgrade_total",
		}, []string{"type"}),
		duration: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: "myApp",
			Name:      "request_duration_seconds",
			Help:      "Duration of the request.",
			// 4 times larger for apdex score
			// Buckets: prometheus.ExponentialBuckets(0.1, 1.5, 5),
			// Buckets: prometheus.LinearBuckets(0.1, 5, 5),
			Buckets: []float64{0.1, 0.15, 0.2, 0.25, 0.3},
		}, []string{"status", "method"}),
		loginDuration: prometheus.NewSummary(prometheus.SummaryOpts{
			Namespace:  "myApp",
			Name:       "login_request_duration_seconds",
			Help:       "Duration of the login request.",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		}),
	}
	reg.MustRegister(m.devices, m.info, m.upgrades, m.duration, m.loginDuration)
	return m
}

var dvs []Device
var version string

func init() {
	version = "2.10.0"
	dvs = []Device{
		{1, "5F-33-CC-1F-43-82", "2.1.6"},
		{2, "EF-2B-C4-F5-D6-34", "2.1.6"},
	}
}
func main() {
	reg := prometheus.NewRegistry()
	reg.MustRegister(collectors.NewGoCollector())

	m := NewMetrics(reg)
	m.devices.Set(float64(len(dvs)))
	m.info.With(prometheus.Labels{"version": version}).Set(1)

	dMux := http.NewServeMux()
	rdh := registerDevicesHandler{metrics: m}
	mdh := manageDevicesHandler{metrics: m}
	lh := loginHandler{}
	mlh := middleware(lh, m)
	
	dMux.Handle("/devices", rdh)
	dMux.Handle("/devices/", mdh)
	dMux.Handle("/login", mlh)


	pMux := http.NewServeMux()
	promHandler := promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg})
	pMux.Handle("/metrics", promHandler)

	go func() {
		log.Fatal(http.ListenAndServe(":8082", dMux))
	}()
	go func() {
		log.Fatal(http.ListenAndServe(":8081", pMux))
	}()
	select {}
	// http.Handle("/defaultmetrics", promhttp.Handler())
	// http.Handle("/metrics", promHandler)
	// http.HandleFunc("/devices", getDevices)
	// http.ListenAndServe(":8081", nil)
}

func getDevices(w http.ResponseWriter, r *http.Request, m *metrics) {
	now := time.Now()
	b, err := json.Marshal(dvs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	m.duration.With(prometheus.Labels{"status": "200", "method": "GET"}).Observe(time.Since(now).Seconds())
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func createDevice(w http.ResponseWriter, r *http.Request, m *metrics) {
	var dv Device

	err := json.NewDecoder(r.Body).Decode(&dv)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	dvs = append(dvs, dv)
	m.devices.Set(float64(len(dvs)))
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Device created!"))
}

type registerDevicesHandler struct {
	metrics *metrics
}

func (rdh registerDevicesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getDevices(w, r, rdh.metrics)
	case "POST":
		createDevice(w, r, rdh.metrics)
	default:
		w.Header().Set("Allow", "GET, POST")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func upgradeDevice(w http.ResponseWriter, r *http.Request, m *metrics) {
	path := strings.TrimPrefix(r.URL.Path, "/devices/")

	id, err := strconv.Atoi(path)
	if err != nil || id < 1 {
		http.NotFound(w, r)
	}

	var dv Device
	err = json.NewDecoder(r.Body).Decode(&dv)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i := range dvs {
		if dvs[i].ID == id {
			dvs[i].Firmware = dv.Firmware
		}
	}
	sleep(1000)
	m.upgrades.With(prometheus.Labels{"type": "router"}).Inc()
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("Upgrading..."))
}

type manageDevicesHandler struct {
	metrics *metrics
}

func (mdh manageDevicesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "PUT":
		upgradeDevice(w, r, mdh.metrics)
	default:
		w.Header().Set("Allow", "PUT")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func sleep(ms int) {
	rand.Seed(time.Now().UnixNano())
	now := time.Now()
	n := rand.Intn(ms + now.Second())
	time.Sleep(time.Duration(n) * time.Millisecond)
}

type loginHandler struct{}

func (l loginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sleep(200)
	w.Write([]byte("Welcome to the app!"))
}

func middleware(next http.Handler, m *metrics) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		next.ServeHTTP(w, r)
		m.loginDuration.Observe(time.Since(now).Seconds())
	})
}
