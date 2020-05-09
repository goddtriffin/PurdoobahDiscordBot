package main

import (
	"context"
	"fmt"
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

func (pb *PurdoobahBot) mux(s disgord.Session, evt *disgord.MessageCreate) {
	if len(evt.Message.Content) == 0 {
		return
	}

	command := strings.ToLower(strings.Fields(evt.Message.Content)[0])

	if command == "!help" {
		pb.replyHelp(s, evt)
	}

	if command == "!ymsh" {
		pb.replyYMSH(s, evt)
	}
}

func (pb *PurdoobahBot) reply(s disgord.Session, evt *disgord.MessageCreate, reply interface{}) {
	_, err := evt.Message.Reply(context.Background(), s, reply)
	if err != nil {
		log.Printf("reply error: %+v\n", err)
	}
}

func (pb *PurdoobahBot) replyHelp(s disgord.Session, evt *disgord.MessageCreate) {
	fields := []*disgord.EmbedField{}
	for _, command := range pb.getCommands() {
		fields = append(fields, &disgord.EmbedField{
			Name:   command.String(),
			Value:  fmt.Sprintf("`%s`", command.description),
			Inline: true,
		})
	}

	help := &disgord.Embed{
		Description: "**PurdoobahBot Commands**",
		Color:       15844367,
		Fields:      fields,
		Thumbnail: &disgord.EmbedThumbnail{
			URL: "https://www.purdoobahs.com/res/image/logo/purdoobahs-white-768x768.png",
		},
	}

	log.Printf("%s (%s) called !help\n", evt.Message.Author.Username, evt.Message.Author.ID)
	pb.reply(s, evt, help)
}

func (pb *PurdoobahBot) replyYMSH(s disgord.Session, evt *disgord.MessageCreate) {
	ymsh := pb.ymsh.String(pb.rand)
	log.Printf("%s (%s) called !YMSH: %s\n", evt.Message.Author.Username, evt.Message.Author.ID, ymsh)
	pb.reply(s, evt, fmt.Sprintf("YMSH stands for...\n||%s||", ymsh))
}

func (pb *PurdoobahBot) getCommands() []command {
	return []command{
		{"help", "lists PurdoobahBot commands"},
		{"YMSH", "secret YMSH definition"},
	}
}
