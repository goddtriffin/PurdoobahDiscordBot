package main

import (
	"context"
	"flag"
	"log"
	"os"
)

func main() {
	token := flag.String("token", "", "Discord Bot Token")
	flag.Parse()

	if *token == "" {
		flag.Usage()
		os.Exit(1)
	}

	pb, err := NewPurdoobahBot(*token)
	if err != nil {
		log.Printf("New PurdoobahBot error")
		panic(err)
	}
	defer pb.StayConnectedUntilInterrupted(context.Background())
}
