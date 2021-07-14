package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/Joker666/goworkerpool/audit/dynamo"

	"github.com/Joker666/goworkerpool/parser/fastly"
	"github.com/Joker666/goworkerpool/storage/s3"
	"github.com/Joker666/goworkerpool/workerpool"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func main() {
	logger, err := createLogger("DEBUG")
	if err != nil {
		log.Fatal("creating logger: %w", err)
	}

	s3Service := s3.New()
	fastlyParser := fastly.New()
	dynamo := dynamo.New()

	var allTask []*workerpool.Task
	// for i := 1; i <= 20; i++ {
	// 	task := workerpool.NewTask(uuid.NewString())
	// 	allTask = append(allTask, task)
	// }

	pool := workerpool.NewPool(allTask, 5, 100, 10*time.Second)

	// ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	// ctx, cancel := context.WithDeadline(context.Background(), 20*time.Second)
	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}

	pool.Run(ctx, wg)

	fmt.Println("running server")

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		task := workerpool.NewTask(uuid.New().String(), s3Service, s3Service, fastlyParser, dynamo)
		pool.AddTask(task)
		fmt.Fprintf(w, "Hi there, task ID %s!", task.ID)
		fmt.Fprintf(w, "CreatedAt %s!", task.CreatedAt)
	})

	server := http.Server{
		Addr:              ":8080",
		Handler:           r,
		WriteTimeout:      75 * time.Second,
		ReadTimeout:       75 * time.Second,
		ReadHeaderTimeout: 75 * time.Second,
	}
	// Launch server start in a separate goroutine
	go func() {
		logger.
			WithField("bind_address", server.Addr).
			Info("starting server")
		err := server.ListenAndServe()
		logger.
			WithError(err).
			Error("starting server")
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(
		stop,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGHUP,
	)

	<-stop
	// Handle shutdown
	fmt.Println("*********************************\nShutdown signal received\n*********************************")
	cancel()              // Signal cancellation to context.Context
	close(pool.Collector) // closing main task channel

	wg.Wait() // Block here until are workers are done
	logger.Error("All workers done, shutting down!")
	err = server.Shutdown(ctx)
	logger.
		WithError(err).
		Error("server gracefully stopped")
}
