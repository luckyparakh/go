package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	ctx := context.Background()
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()
	// g, ctx := errgroup.WithContext(ctx)

	wg := sync.WaitGroup{}
	u := user{&wg}
	r := mux.NewRouter()
	r.HandleFunc("/", u.echo).Methods(http.MethodGet)
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
	wg.Wait()
	// Don't same context as above because whe ctx will be cancelled it will also cancel
	//the shutdown call on same time and will not wait for inflight call to complete
	if err := s.Shutdown(context.TODO()); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Shutdown done")

}

type user struct {
	wg *sync.WaitGroup
}

func (u user) echo(w http.ResponseWriter, r *http.Request) {
	u.wg.Add(2)
	userAddition(u.wg)
	go notifyUser(u.wg)
}

func userAddition(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Adding User")
	time.Sleep(5 * time.Second)
	fmt.Println("User added")
}

func notifyUser(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Notify User")
	time.Sleep(7 * time.Second)
	fmt.Println("User notified")
}
