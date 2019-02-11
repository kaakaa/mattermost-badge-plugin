package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mattermost/mattermost-server/plugin"
	"github.com/narqo/go-badge"
)

func (p *Plugin) ServeHTTP(c *plugin.Context, w http.ResponseWriter, r *http.Request) {
	p.router.ServeHTTP(w, r)
}

func (p *Plugin) InitAPI() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", p.handleInfo).Methods("GET")
	r.HandleFunc("/teams/{team_id:[a-z0-9]+}/channels/{channel_id:[a-z0-9]+}/badge.svg", p.handleBadge).Methods("GET")
	return r
}

func (p *Plugin) handleInfo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("MetterBadge"))
}

func (p *Plugin) handleBadge(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	channelID := vars["channel_id"]
	ch, appErr := p.API.GetChannel(channelID)
	if appErr != nil {
		return
	}
	w.Header().Set("Content-Type", "image/svg+xml")
	err := badge.Render("Mattermost", ch.DisplayName, "#5272B4", w)
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusOK)
}
