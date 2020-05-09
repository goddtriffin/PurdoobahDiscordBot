package main

import (
	"context"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/andersfylling/disgord"
)

// PurdoobahBot is the Discord PurdoobahBot.
type PurdoobahBot struct {
	*disgord.Client

	rand *rand.Rand

	ymsh *ymsh
}

// NewPurdoobahBot creates a new PurdoobahBot.
func NewPurdoobahBot(botToken string) (*PurdoobahBot, error) {
	ymsh, err := newYMSH()
	if err != nil {
		panic(err)
	}

	pb := &PurdoobahBot{
		Client: disgord.New(disgord.Config{
			ProjectName: "PurdoobahBot",
			BotToken:    botToken,
			Logger:      disgord.DefaultLogger(false),
		}),
		rand: rand.New(rand.NewSource(time.Now().UnixNano())),
		ymsh: ymsh,
	}

	pb.On(disgord.EvtMessageCreate, pb.mux)

	return pb, nil
}

func (pb *PurdoobahBot) mux(session disgord.Session, m *disgord.MessageCreate) {
	if strings.ToLower(m.Message.Content) == "!ymsh" {
		ymsh := pb.ymsh.String(pb.rand)
		log.Println("YMSH:", ymsh)
		_, err := m.Message.Reply(context.Background(), session, ymsh)
		if err != nil {
			log.Printf("YMSH Reply error: %+v\n", err)
		}
	}
}
