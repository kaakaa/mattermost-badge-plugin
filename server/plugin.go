package main

import (
	"github.com/gorilla/mux"
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/plugin"
)

const TriggerWord = "badge"

type Plugin struct {
	plugin.MattermostPlugin
	router *mux.Router
}

func (p *Plugin) OnActivate() error {
	p.router = p.InitAPI()
	return p.API.RegisterCommand(&model.Command{
		Trigger:          TriggerWord,
		DisplayName:      "Matter-Badge",
		Description:      "Generate badge for invitation to Mattermost",
		AutoComplete:     true,
		AutoCompleteDesc: "Generate badge URL",
	})
}

func (p *Plugin) OnDeactivate() error {
	return p.API.UnregisterCommand("", TriggerWord)
}
