package main

import (
	"fmt"
	"strings"

	"github.com/andersfylling/disgord"
)

func (pb *PurdoobahBot) commandHelp(s disgord.Session, evt *disgord.MessageCreate) {
	// convert list of commands to list of Disgord fields
	fields := []*disgord.EmbedField{}
	for _, command := range pb.commands {
		fields = append(fields, &disgord.EmbedField{
			Name:   fmt.Sprintf("`%s`", command.String()),
			Value:  command.description,
			Inline: true,
		})
	}

	pb.reply(s, evt, &disgord.Embed{
		Description: "**PurdoobahBot Help**",
		Color:       15844367,
		Fields:      fields,
		Thumbnail:   &disgord.EmbedThumbnail{URL: pb.thumbnailURL},
	})

	pb.Logger().Info(fmt.Sprintf("%s (%s) called !commands\n", evt.Message.Author.Username, evt.Message.Author.ID))
}

func (pb *PurdoobahBot) commandYMSH(s disgord.Session, evt *disgord.MessageCreate) {
	ymsh := pb.ymsh.String(pb.rand)
	pb.reply(s, evt, fmt.Sprintf("YMSH stands for... ||%s||", ymsh))

	pb.Logger().Info(fmt.Sprintf("%s (%s) called !YMSH: %s\n", evt.Message.Author.Username, evt.Message.Author.ID, ymsh))
}

func (pb *PurdoobahBot) commandSocials(s disgord.Session, evt *disgord.MessageCreate) {
	// convert list of PR account links to list of Disgord fields
	fields := []*disgord.EmbedField{}
	for name, URL := range pb.socialMedia {
		fields = append(fields, &disgord.EmbedField{
			Name:  strings.Title(strings.ToLower(name)),
			Value: URL,
		})
	}

	pb.reply(s, evt, &disgord.Embed{
		Description: "**PurdoobahBot Socials**",
		Color:       15844367,
		Fields:      fields,
		Thumbnail: &disgord.EmbedThumbnail{
			URL: pb.thumbnailURL,
		},
	})

	pb.Logger().Info(fmt.Sprintf("%s (%s) called !pr\n", evt.Message.Author.Username, evt.Message.Author.ID))
}
