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

	commands     []command
	thumbnailURL string
	socialMedia  map[string]string // name -> URL
	ymsh         *ymsh
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

		commands: []command{
			{"help", "displays help"},
			{"commands", "displays commands"},
			{"YMSH", "secret definition of YMSH"},
			{"pr", "links social media"},
			{"website", "links official website"},
			{"instagram", "links official website"},
			{"facebook", "links official website"},
			{"youtube", "links official website"},
		},
		thumbnailURL: "https://www.purdoobahs.com/res/image/logo/purdoobahs-white-768x768.png",
		socialMedia: map[string]string{
			"website":   "https://www.purdoobahs.com/",
			"instagram": "https://www.instagram.com/purdoobahs/",
			"facebook":  "https://www.facebook.com/purdoobahs/",
			"youtube":   "https://www.youtube.com/channel/UCIH2OACGjUeDPfkISb_lp_Q",
		},
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

	switch command {
	case "!help":
		pb.replyHelp(s, evt)
	case "!commands":
		pb.replyCommands(s, evt)
	case "!ymsh":
		pb.replyYMSH(s, evt)
	case "!pr":
		pb.replyPR(s, evt)
	case "!website":
		pb.replyWebsite(s, evt)
	case "!instagram":
		pb.replyInstagram(s, evt)
	case "!facebook":
		pb.replyFacebook(s, evt)
	case "!youtube":
		pb.replyYoutube(s, evt)
	}
}

func (pb *PurdoobahBot) reply(s disgord.Session, evt *disgord.MessageCreate, reply interface{}) {
	_, err := evt.Message.Reply(context.Background(), s, reply)
	if err != nil {
		log.Printf("reply error: %+v\n", err)
	}
}

func (pb *PurdoobahBot) replyHelp(s disgord.Session, evt *disgord.MessageCreate) {
	embedHelp := &disgord.Embed{
		Description: "**PurdoobahBot Help**",
		Color:       15844367,
		Fields: []*disgord.EmbedField{
			{Name: "`!commands`", Value: "displays commands"},
		},
		Thumbnail: &disgord.EmbedThumbnail{URL: pb.thumbnailURL},
	}

	log.Printf("%s (%s) called !help\n", evt.Message.Author.Username, evt.Message.Author.ID)
	pb.reply(s, evt, embedHelp)
}

func (pb *PurdoobahBot) replyCommands(s disgord.Session, evt *disgord.MessageCreate) {
	fields := []*disgord.EmbedField{}
	for _, command := range pb.commands {
		fields = append(fields, &disgord.EmbedField{
			Name:  fmt.Sprintf("`%s`", command.String()),
			Value: command.description,
		})
	}

	embedCommands := &disgord.Embed{
		Description: "**PurdoobahBot Commands**",
		Color:       15844367,
		Fields:      fields,
		Thumbnail: &disgord.EmbedThumbnail{
			URL: "https://www.purdoobahs.com/res/image/logo/purdoobahs-white-768x768.png",
		},
	}

	log.Printf("%s (%s) called !commands\n", evt.Message.Author.Username, evt.Message.Author.ID)
	pb.reply(s, evt, embedCommands)
}

func (pb *PurdoobahBot) replyYMSH(s disgord.Session, evt *disgord.MessageCreate) {
	ymsh := pb.ymsh.String(pb.rand)
	log.Printf("%s (%s) called !YMSH: %s\n", evt.Message.Author.Username, evt.Message.Author.ID, ymsh)
	pb.reply(s, evt, fmt.Sprintf("YMSH stands for... ||%s||", ymsh))
}

func (pb *PurdoobahBot) replyPR(s disgord.Session, evt *disgord.MessageCreate) {
	fields := []*disgord.EmbedField{}
	for name, URL := range pb.socialMedia {
		fields = append(fields, &disgord.EmbedField{
			Name:  strings.Title(strings.ToLower(name)),
			Value: URL,
		})
	}

	embedPR := &disgord.Embed{
		Description: "**PurdoobahBot Commands**",
		Color:       15844367,
		Fields:      fields,
		Thumbnail: &disgord.EmbedThumbnail{
			URL: "https://www.purdoobahs.com/res/image/logo/purdoobahs-white-768x768.png",
		},
	}

	log.Printf("%s (%s) called !pr\n", evt.Message.Author.Username, evt.Message.Author.ID)
	pb.reply(s, evt, embedPR)
}
func (pb *PurdoobahBot) replyWebsite(s disgord.Session, evt *disgord.MessageCreate) {
	log.Printf("%s (%s) called !website\n", evt.Message.Author.Username, evt.Message.Author.ID)
	pb.reply(s, evt, pb.socialMedia["website"])
}
func (pb *PurdoobahBot) replyInstagram(s disgord.Session, evt *disgord.MessageCreate) {
	log.Printf("%s (%s) called !instagram\n", evt.Message.Author.Username, evt.Message.Author.ID)
	pb.reply(s, evt, pb.socialMedia["instagram"])
}
func (pb *PurdoobahBot) replyFacebook(s disgord.Session, evt *disgord.MessageCreate) {
	log.Printf("%s (%s) called !facebook\n", evt.Message.Author.Username, evt.Message.Author.ID)
	pb.reply(s, evt, pb.socialMedia["facebook"])
}
func (pb *PurdoobahBot) replyYoutube(s disgord.Session, evt *disgord.MessageCreate) {
	log.Printf("%s (%s) called !youtube\n", evt.Message.Author.Username, evt.Message.Author.ID)
	pb.reply(s, evt, pb.socialMedia["youtube"])
}
