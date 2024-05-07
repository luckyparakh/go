package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	ctx := context.Background()
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()
	// g, ctx := errgroup.WithContext(ctx)

	r := mux.NewRouter()
	r.HandleFunc("/", echo).Methods(http.MethodGet)
	s := http.Server{
		Handler: r,
		Addr:    ":9098",
	}
	go func() error {
		fmt.Println("Server Started")
		return s.ListenAndServe()
	}()
	<-ctx.Done()
	fmt.Println("Interrupt came")
	// Shutdown will wait for all ongoing requests on this thread/routine but no other go-routine.
	// I.e. it will wait for userAddition func to complete but not for notifyUser func

	// Don't same context as above because whe ctx will be cancelled it will also cancel
	//the shutdown call on same time and will not wait for inflight call to complete
	if err := s.Shutdown(context.TODO()); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Shutdown done")
}

func echo(w http.ResponseWriter, r *http.Request) {
	userAddition()
	go notifyUser()
}

func userAddition() {
	fmt.Println("Adding User")
	time.Sleep(5 * time.Second)
	fmt.Println("User added")
}

func notifyUser() {
	fmt.Println("Notify User")
	time.Sleep(7 * time.Second)
	fmt.Println("User notified")
}
