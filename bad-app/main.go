package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"time"
)

func main() {
	crashAfter := 30*time.Second + time.Duration(rand.Intn(60))*time.Second
	log.SetOutput(os.Stderr)
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)

	burnCPU := envBool("BURN_CPU")
	hogMemoryMB := envInt("HOG_MEMORY_MB", 0)
	log.Printf("starting pid=%d cpus=%d burn_cpu=%v hog_memory_mb=%d crash_in=%s",
		os.Getpid(), runtime.NumCPU(), burnCPU, hogMemoryMB, crashAfter)

	if burnCPU {
		for i := 0; i < runtime.NumCPU(); i++ {
			go cpuLoop()
		}
	}
	if hogMemoryMB > 0 {
		go memoryHog(hogMemoryMB)
	}
	go emitLogs()

	time.Sleep(crashAfter)
	log.Fatal("FATAL: unrecoverable error, shutting down")
}

func envBool(key string) bool {
	v, _ := strconv.ParseBool(os.Getenv(key))
	return v
}

func envInt(key string, def int) int {
	if s := os.Getenv(key); s != "" {
		if n, err := strconv.Atoi(s); err == nil {
			return n
		}
	}
	return def
}

func cpuLoop() {
	x := 1.0
	for {
		for range 1_000_000 {
			x = math.Sin(x)*math.Cos(x) + math.Sqrt(math.Abs(x)+1)
		}
		_ = x
	}
}

func memoryHog(mb int) {
	chunk := 1024 * 1024 // 1 MB
	var hold [][]byte
	for range mb {
		b := make([]byte, chunk)
		for i := range b {
			b[i] = byte(i)
		}
		hold = append(hold, b)
		time.Sleep(100 * time.Millisecond)
	}
	// keep alive
	select {}
}

func emitLogs() {
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
	infos := []string{
		"request completed successfully",
		"health check passed",
		"cache refreshed",
		"connected to upstream service",
		"configuration reloaded",
		"worker started processing batch",
		"metrics snapshot exported",
		"session token renewed",
	}
	for seq := 1; ; seq++ {
		delay := time.Duration(500+rand.Intn(2000)) * time.Millisecond
		time.Sleep(delay)
		worker := rand.Intn(4)
		if rand.Intn(3) == 0 { // ~1/3 info, ~2/3 error
			msg := infos[rand.Intn(len(infos))]
			fmt.Fprintf(os.Stderr, "INFO  [worker-%d] %s seq=%d\n", worker, msg, seq)
		} else {
			msg := errors[rand.Intn(len(errors))]
			fmt.Fprintf(os.Stderr, "ERROR [worker-%d] %s seq=%d\n", worker, msg, seq)
		}
	}
}