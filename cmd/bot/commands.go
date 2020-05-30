package main

import (
	"fmt"
	"strings"

	"github.com/andersfylling/disgord"
)

func (pb *PurdoobahBot) commandHelp(s disgord.Session, evt *disgord.MessageCreate) {
	pb.reply(s, evt, &disgord.Embed{
		Description: "**PurdoobahBot Help**",
		Color:       15844367,
		Fields: []*disgord.EmbedField{
			{Name: "`!commands`", Value: "displays commands"},
		},
		Thumbnail: &disgord.EmbedThumbnail{URL: pb.thumbnailURL},
	})

	pb.Logger().Info(fmt.Sprintf("%s (%s) called !help\n", evt.Message.Author.Username, evt.Message.Author.ID))
}

func (pb *PurdoobahBot) commandCommands(s disgord.Session, evt *disgord.MessageCreate) {
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
		Description: "**PurdoobahBot Commands**",
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

func (pb *PurdoobahBot) commandPR(s disgord.Session, evt *disgord.MessageCreate) {
	// convert list of PR account links to list of Disgord fields
	fields := []*disgord.EmbedField{}
	for name, URL := range pb.socialMedia {
		fields = append(fields, &disgord.EmbedField{
			Name:  strings.Title(strings.ToLower(name)),
			Value: URL,
		})
	}

	pb.reply(s, evt, &disgord.Embed{
		Description: "**PurdoobahBot Commands**",
		Color:       15844367,
		Fields:      fields,
		Thumbnail: &disgord.EmbedThumbnail{
			URL: "https://www.purdoobahs.com/res/image/logo/purdoobahs-white-768x768.png",
		},
	})

	pb.Logger().Info(fmt.Sprintf("%s (%s) called !pr\n", evt.Message.Author.Username, evt.Message.Author.ID))
}

func (pb *PurdoobahBot) commandWebsite(s disgord.Session, evt *disgord.MessageCreate) {
	pb.reply(s, evt, pb.socialMedia["website"])

	pb.Logger().Info(fmt.Sprintf("%s (%s) called !website\n", evt.Message.Author.Username, evt.Message.Author.ID))
}

func (pb *PurdoobahBot) commandInstagram(s disgord.Session, evt *disgord.MessageCreate) {
	pb.reply(s, evt, pb.socialMedia["instagram"])

	pb.Logger().Info(fmt.Sprintf("%s (%s) called !instagram\n", evt.Message.Author.Username, evt.Message.Author.ID))
}

func (pb *PurdoobahBot) commandFacebook(s disgord.Session, evt *disgord.MessageCreate) {
	pb.reply(s, evt, pb.socialMedia["facebook"])

	pb.Logger().Info(fmt.Sprintf("%s (%s) called !facebook\n", evt.Message.Author.Username, evt.Message.Author.ID))
}

func (pb *PurdoobahBot) commandYoutube(s disgord.Session, evt *disgord.MessageCreate) {
	pb.reply(s, evt, pb.socialMedia["youtube"])

	pb.Logger().Info(fmt.Sprintf("%s (%s) called !youtube\n", evt.Message.Author.Username, evt.Message.Author.ID))
}

func (pb *PurdoobahBot) commandGithub(s disgord.Session, evt *disgord.MessageCreate) {
	pb.reply(s, evt, pb.socialMedia["github"])

	pb.Logger().Info(fmt.Sprintf("%s (%s) called !github\n", evt.Message.Author.Username, evt.Message.Author.ID))
}

func (pb *PurdoobahBot) commandEmail(s disgord.Session, evt *disgord.MessageCreate) {
	pb.reply(s, evt, fmt.Sprintf("` %s `", pb.socialMedia["email"]))

	pb.Logger().Info(fmt.Sprintf("%s (%s) called !email\n", evt.Message.Author.Username, evt.Message.Author.ID))
}
