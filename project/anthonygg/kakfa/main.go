package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

const (
	APIKeyAPIVersion = 18
)

type ApiVersion struct {
	APIKey     int16
	MinVersion int32
	MaxVersion int32
}
type ApiVersionResponse struct {
	ErrorCode      int16
	ApiVersions    []ApiVersion
	ThrottleTimeMs int32
}

func (avr *ApiVersionResponse) Encode(w io.Writer) error {
	buf := &bytes.Buffer{}
	binary.Write(buf, binary.BigEndian, int32(0))
	binary.Write(buf, binary.BigEndian, avr.ErrorCode)
	binary.Write(buf, binary.BigEndian, int32(len(avr.ApiVersions)))
	for _, api := range avr.ApiVersions {
		binary.Write(buf, binary.BigEndian, api.APIKey)
		binary.Write(buf, binary.BigEndian, api.MinVersion)
		binary.Write(buf, binary.BigEndian, api.MaxVersion)
	}
	binary.Write(buf, binary.BigEndian, avr.ThrottleTimeMs)
	binary.BigEndian.PutUint32(buf.Bytes(), uint32(buf.Len()-4))
	_, err := io.Copy(w, buf)
	return err
}

// https://kafka.apache.org/protocol.html#The_Messages_ApiVersions
type Header struct {
	Size          int32
	APIKey        int16
	APIVersion    int16
	CorrelationID int32
}
type ClientSoftware struct {
	ClientIDLength int16
	Name           []byte // Compact_string
	Version        []byte // Compact_string
}

func readAPIVersion(r *bytes.Reader) (*ClientSoftware, error) {
	cs := ClientSoftware{}

	err := binary.Read(r, binary.BigEndian, &cs.ClientIDLength)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}
	cs.Name = make([]byte, cs.ClientIDLength)
	err = binary.Read(r, binary.BigEndian, &cs.Name)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}
	cs.Version = make([]byte, r.Len())
	err = binary.Read(r, binary.BigEndian, &cs.Version)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}
	// fmt.Printf("%s, %s\n", cs.Name, cs.Version)
	return &cs, nil
}

type Message struct {
	data []byte
}
type Server struct {
	cOffsets map[string]int
	buffer   []Message
	ln       net.Listener
}

func NewServer() *Server {
	return &Server{
		cOffsets: make(map[string]int),
		buffer:   make([]Message, 0),
	}
}
func (s *Server) handleConn(conn net.Conn) error {
	fmt.Println("New Conn", conn.RemoteAddr())
	buf := make([]byte, 1024)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				slog.Info("End of Stream or file or no more data to read from buffer")
				return nil
			}
			slog.Error("Conn error", "err", err.Error())
			return err
		}
		// fmt.Println("From Buffer:", buf[:n])
		rawMsg := buf[:n]
		var h Header
		r := bytes.NewReader(rawMsg)
		err = binary.Read(r, binary.BigEndian, &h)
		if err != nil {
			slog.Error(err.Error())
			return err
		}
		fmt.Printf("%+v\n", h)
		switch h.APIKey {

		case APIKeyAPIVersion:
			cs, err := readAPIVersion(r)
			if err != nil {
				return err
			}
			fmt.Printf("%+v\n", cs)
			resp := ApiVersionResponse{
				ErrorCode:      0,
				ThrottleTimeMs: 10,
				ApiVersions: []ApiVersion{
					{
						APIKey:     0,
						MinVersion: 0,
						MaxVersion: 0,
					},
				},
			}
			resp.Encode(conn)
		default:
			fmt.Println("API Not supported", h.APIKey)
		}

	}
}
func (s *Server) Listen() error {
	ln, err := net.Listen("tcp", ":9092")
	if err != nil {
		return err
	}
	s.ln = ln
	for {
		conn, err := ln.Accept()
		if err != nil {
			if err == io.EOF {
				return err
			}
			slog.Error("", err)
		}
		go s.handleConn(conn)
	}
}

func main() {
	server := NewServer()
	// log.Fatal(server.Listen())
	go func() {

		log.Fatal(server.Listen())
	}()
	time.Sleep(time.Second)
	// fmt.Println("consuming...")
	// consumer()

	fmt.Println("producer")
	producer()
	// select {}
}

func consumer() error {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":   "localhost:9092",
		"group.id":            "myGroup",
		"auto.offset.reset":   "earliest",
		"api.version.request": false,
	})

	if err != nil {
		return err
	}

	c.SubscribeTopics([]string{"myTopic"}, nil)

	// A signal handler or similar could be used to set this to false to break the loop.
	run := true

	for run {
		msg, err := c.ReadMessage(time.Second)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else if !err.(kafka.Error).IsTimeout() {
			// The client will automatically try to recover from all errors.
			// Timeout is not considered an error because it is raised by
			// ReadMessage in absence of messages.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}

	return c.Close()
}

func producer() error {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":   "localhost:9092",
		"api.version.request": false,
	})
	if err != nil {
		slog.Error("New Producer", "err", err.Error())
		panic(err)
	}

	defer p.Close()

	// Delivery report handler for produced messages
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	// Produce messages to topic (asynchronously)
	topic := "myTopic"
	for _, word := range []string{"Welcome", "to", "the", "Confluent", "Kafka", "Golang", "client"} {
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(word),
		}, nil)
		time.Sleep(time.Second)
	}

	// Wait for message deliveries before shutting down
	p.Flush(15 * 1000)
	return nil
}
