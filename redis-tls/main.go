package main

import (
	"crypto/tls"
	"github.com/go-redis/redis/v8"
)

func main() {
	tlsConfig := tls.Config{}

	_ = redis.NewClient(&redis.Options{
		Addr:      "localhost:6379",
		Password:  "", // no password set
		DB:        0,  // use default DB
		TLSConfig: &tlsConfig,
	})
}
