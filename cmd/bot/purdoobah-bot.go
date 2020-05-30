package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/andersfylling/disgord"
	"github.com/andersfylling/disgord/std"
	"github.com/sirupsen/logrus"
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
	logger := &logrus.Logger{
		Out:       os.Stderr,
		Formatter: new(logrus.JSONFormatter),
		Hooks:     make(logrus.LevelHooks),
		Level:     logrus.DebugLevel,
	}

	ymsh, err := newYMSH()
	if err != nil {
		panic(err)
	}

	pb := &PurdoobahBot{
		Client: disgord.New(disgord.Config{
			ProjectName: "PurdoobahBot",
			BotToken:    botToken,
			Logger:      logger,
		}),
		rand: rand.New(rand.NewSource(time.Now().UnixNano())),

		commands: []command{
			{"help", "displays help"},
			{"commands", "displays commands"},
			{"YMSH", "secret definition of YMSH"},
			{"pr", "links social media"},
			{"website", "links official website"},
			{"instagram", "links official Instagram"},
			{"facebook", "links official Facebook"},
			{"youtube", "links official Youtube"},
			{"github", "links official Github"},
			{"email", "links official E-mail"},
		},
		thumbnailURL: "https://www.purdoobahs.com/res/image/logo/purdoobahs-white-768x768.png",
		socialMedia: map[string]string{
			"website":   "https://www.purdoobahs.com/",
			"instagram": "https://www.instagram.com/purdoobahs/",
			"facebook":  "https://www.facebook.com/purdoobahs/",
			"youtube":   "https://www.youtube.com/channel/UCIH2OACGjUeDPfkISb_lp_Q",
			"github":    "https://github.com/purdoobahs",
			"email":     "purdoobahs@gmail.com",
		},
		ymsh: ymsh,
	}

	pb.Ready(func() {
		pb.Logger().Info("PurdoobahBot is online!")
	})

	// filters
	filter, _ := std.NewMsgFilter(context.Background(), pb)
	filter.SetPrefix("!")

	// !help
	pb.On(
		disgord.EvtMessageCreate,

		filter.NotByBot,
		filter.HasPrefix,

		filterNonHelpCommands,
		pb.commandHelp,
	)

	// !command
	pb.On(
		disgord.EvtMessageCreate,

		filter.NotByBot,
		filter.HasPrefix,

		filterNonCommandsCommands,
		pb.commandCommands,
	)

	// !ymsh
	pb.On(
		disgord.EvtMessageCreate,

		filter.NotByBot,
		filter.HasPrefix,

		filterNonYMSHCommands,
		pb.commandYMSH,
	)

	// !pr
	pb.On(
		disgord.EvtMessageCreate,

		filter.NotByBot,
		filter.HasPrefix,

		filterNonPRCommands,
		pb.commandPR,
	)

	// !website
	pb.On(
		disgord.EvtMessageCreate,

		filter.NotByBot,
		filter.HasPrefix,

		filterNonWebsiteCommands,
		pb.commandWebsite,
	)

	// !instagram
	pb.On(
		disgord.EvtMessageCreate,

		filter.NotByBot,
		filter.HasPrefix,

		filterNonInstagramCommands,
		pb.commandInstagram,
	)

	// !facebook
	pb.On(
		disgord.EvtMessageCreate,

		filter.NotByBot,
		filter.HasPrefix,

		filterNonFacebookCommands,
		pb.commandFacebook,
	)

	// !youtube
	pb.On(
		disgord.EvtMessageCreate,

		filter.NotByBot,
		filter.HasPrefix,

		filterNonYoutubeCommands,
		pb.commandYoutube,
	)

	// !github
	pb.On(
		disgord.EvtMessageCreate,

		filter.NotByBot,
		filter.HasPrefix,

		filterNonGithubCommands,
		pb.commandGithub,
	)

	// !email
	pb.On(
		disgord.EvtMessageCreate,

		filter.NotByBot,
		filter.HasPrefix,

		filterNonEmailCommands,
		pb.commandEmail,
	)

	// add up/down vote reactions to every message
	pb.On(
		disgord.EvtMessageCreate,

		filterNonDM,

		pb.react,
	)

	return pb, nil
}

func (pb *PurdoobahBot) reply(s disgord.Session, evt *disgord.MessageCreate, reply interface{}) {
	_, err := evt.Message.Reply(context.Background(), s, reply)
	if err != nil {
		pb.Logger().Error(fmt.Sprintf("reply error: %+v\n", err))
	}
}
