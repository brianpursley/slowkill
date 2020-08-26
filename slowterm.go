package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {
	exitCode, err := strconv.Atoi(os.Getenv("EXIT"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to parse exit code: %v", err)
		os.Exit(255)
	}

	sleep, err := strconv.Atoi(os.Getenv("SLEEP"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to parse sleep: %v", err)
		os.Exit(255)
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("awaiting signal")
	sig := <-sigs

	fmt.Printf("got %v signal, sleeping for %d seconds...\n", sig, sleep)
	time.Sleep(time.Duration(sleep) * time.Second)

	fmt.Println("done sleeping")
	os.Exit(exitCode)
}
