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
		log.Printf("failed to initialize PurdoobahBot\n")
		panic(err)
	}
	defer func() {
		err = pb.StayConnectedUntilInterrupted(context.Background())
		if err != nil {
			log.Println(err)
		}
	}()
}
