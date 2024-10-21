package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/redis/go-redis/v9"
)

func main() {
	m := http.NewServeMux()

	rc := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDR"),
		DB:   0,
	})

	m.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		status := rc.Ping(r.Context())
		if err := status.Err(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "error connecting to redis: %s\n", err.Error())
			return
		}
		stats := rc.PoolStats()

		fmt.Fprintf(w, "I am connected to redis with %d connections\n", stats.TotalConns)
	})

	log.Fatal(http.ListenAndServe(":8000", m))
}
