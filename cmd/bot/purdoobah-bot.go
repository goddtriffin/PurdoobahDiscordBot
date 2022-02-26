package main

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
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
			{"YMSH", "find out the secret definition of YMSH 👀"},
			{"socials", "links social media"},
		},
		thumbnailURL: "https://www.purdoobahs.com/static/image/socials/purdoobahs.webp",
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
		pb.startHealthEndpointServer()
	})

	// filters
	filter, _ := std.NewMsgFilter(context.Background(), pb)
	filter.SetPrefix("/")

	// /help
	pb.On(
		disgord.EvtMessageCreate,

		filter.NotByBot,
		filter.HasPrefix,

		filterNonHelpCommands,
		pb.commandHelp,
	)

	// /ymsh
	pb.On(
		disgord.EvtMessageCreate,

		filter.NotByBot,
		filter.HasPrefix,

		filterNonYMSHCommands,
		pb.commandYMSH,
	)

	// /socials
	pb.On(
		disgord.EvtMessageCreate,

		filter.NotByBot,
		filter.HasPrefix,

		filterNonSocialsCommands,
		pb.commandSocials,
	)

	return pb, nil
}

func (pb *PurdoobahBot) reply(s disgord.Session, evt *disgord.MessageCreate, reply interface{}) {
	_, err := evt.Message.Reply(context.Background(), s, reply)
	if err != nil {
		pb.Logger().Error(fmt.Sprintf("reply error: %+v\n", err))
	}
}

func (pb *PurdoobahBot) startHealthEndpointServer() {
	// create the server
	addr := ":8080"
	srv := &http.Server{
		Addr: addr,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			pb.Logger().Info(fmt.Sprintf("%s %s", req.Method, req.URL.Path))

			if req.URL.Path != "/api/v1/health" {
				w.WriteHeader(http.StatusNotFound)
				return
			}

			w.WriteHeader(http.StatusOK)
			_, err := w.Write([]byte("OK"))
			if err != nil {
				pb.Logger().Error(err)
				return
			}
		}),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// run the server
	pb.Logger().Info(fmt.Sprintf("Health checkpoint is being served at: %s/api/v1/health", addr))
	err := srv.ListenAndServe()
	if err != nil {
		// print error on exit
		pb.Logger().Error(err)
		os.Exit(1)
	}
}
