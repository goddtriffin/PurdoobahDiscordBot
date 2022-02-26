package main

import "fmt"

const commandSignal = "/"

type command struct {
	name        string
	description string
}

func (c command) String() string {
	return fmt.Sprintf("%s%s", commandSignal, c.name)
}

func (c command) Help() string {
	return fmt.Sprintf("%s: %s", c.String(), c.description)
}
