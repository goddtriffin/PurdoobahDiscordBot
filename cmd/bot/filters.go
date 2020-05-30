package main

import (
	"strings"

	"github.com/andersfylling/disgord"
)

func filterNonHelpCommands(evt interface{}) interface{} {
	m := evt.(*disgord.MessageCreate)
	if strings.ToLower(m.Message.Content) != "!help" {
		return nil
	}
	return evt
}

func filterNonCommandsCommands(evt interface{}) interface{} {
	m := evt.(*disgord.MessageCreate)
	if strings.ToLower(m.Message.Content) != "!commands" {
		return nil
	}
	return evt
}

func filterNonYMSHCommands(evt interface{}) interface{} {
	m := evt.(*disgord.MessageCreate)
	if strings.ToLower(m.Message.Content) != "!ymsh" {
		return nil
	}
	return evt
}

func filterNonPRCommands(evt interface{}) interface{} {
	m := evt.(*disgord.MessageCreate)
	if strings.ToLower(m.Message.Content) != "!pr" {
		return nil
	}
	return evt
}

func filterNonWebsiteCommands(evt interface{}) interface{} {
	m := evt.(*disgord.MessageCreate)
	if strings.ToLower(m.Message.Content) != "!website" {
		return nil
	}
	return evt
}

func filterNonInstagramCommands(evt interface{}) interface{} {
	m := evt.(*disgord.MessageCreate)
	if strings.ToLower(m.Message.Content) != "!instagram" {
		return nil
	}
	return evt
}

func filterNonFacebookCommands(evt interface{}) interface{} {
	m := evt.(*disgord.MessageCreate)
	if strings.ToLower(m.Message.Content) != "!facebook" {
		return nil
	}
	return evt
}

func filterNonYoutubeCommands(evt interface{}) interface{} {
	m := evt.(*disgord.MessageCreate)
	if strings.ToLower(m.Message.Content) != "!youtube" {
		return nil
	}
	return evt
}

func filterNonGithubCommands(evt interface{}) interface{} {
	m := evt.(*disgord.MessageCreate)
	if strings.ToLower(m.Message.Content) != "!github" {
		return nil
	}
	return evt
}

func filterNonEmailCommands(evt interface{}) interface{} {
	m := evt.(*disgord.MessageCreate)
	if strings.ToLower(m.Message.Content) != "!email" {
		return nil
	}
	return evt
}
