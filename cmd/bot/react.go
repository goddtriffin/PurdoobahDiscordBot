package main

import (
	"context"
	"fmt"

	"github.com/andersfylling/disgord"
)

func (pb *PurdoobahBot) react(s disgord.Session, evt *disgord.MessageCreate) {
	err := evt.Message.React(context.Background(), s, "👍")
	if err != nil {
		pb.Logger().Error(fmt.Sprintf("react error: %+v\n", err))
	}

	err = evt.Message.React(context.Background(), s, "👎")
	if err != nil {
		pb.Logger().Error(fmt.Sprintf("react error: %+v\n", err))
	}
}
