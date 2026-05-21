package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"runtime"
	"time"
)

func main() {
	crashAfter := 30*time.Second + time.Duration(rand.Intn(60))*time.Second
	log.SetOutput(os.Stderr)
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	log.Printf("starting pid=%d cpus=%d crash_in=%s", os.Getpid(), runtime.NumCPU(), crashAfter)

	for i := 0; i < runtime.NumCPU(); i++ {
		go burnCPU()
	}
	go logErrors()

	time.Sleep(crashAfter)
	log.Fatal("FATAL: unrecoverable error, shutting down")
}

func burnCPU() {
	x := 1.0
	for {
		for range 1_000_000 {
			x = math.Sin(x)*math.Cos(x) + math.Sqrt(math.Abs(x)+1)
		}
		_ = x
	}
}

func logErrors() {
	errors := []string{
		"connection to database timed out",
		"failed to read from cache: connection refused",
		"request handler panic recovered",
		"TLS handshake error: remote error: tls: bad certificate",
		"disk write latency exceeded threshold: 500ms",
		"memory allocation failed: out of memory",
		"upstream service returned HTTP 503",
		"failed to renew lease: context deadline exceeded",
	}
	for seq := 1; ; seq++ {
		delay := time.Duration(500+rand.Intn(2000)) * time.Millisecond
		time.Sleep(delay)
		msg := errors[rand.Intn(len(errors))]
		fmt.Fprintf(os.Stderr, "ERROR [worker-%d] %s seq=%d\n", rand.Intn(4), msg, seq)
	}
}
