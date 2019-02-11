package main

import (
	"fmt"

	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/plugin"
)

const (
	// Parameters: SiteURL, PluginID, teamID, channelID,
	BadgeMarkdownTemplate = "![Mattermost Badge](%s/plugins/%s/teams/%s/channels/%s/badge.svg)"
	// Parameters: SiteURL, TeamHandle, ChannelHandle
	ChannelURLTemplate = "%s/%s/channels/%s"
)

func (p *Plugin) ExecuteCommand(c *plugin.Context, args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	// TODO: Command arguments. e.g.: /badge ~[channel] :[markdown|url]
	siteURL := args.SiteURL
	teamID := args.TeamId
	channelID := args.ChannelId

	team, err := p.API.GetTeam(teamID)
	if err != nil {
		return nil, err
	}
	ch, err := p.API.GetChannel(channelID)
	if err != nil {
		return nil, err
	}

	badge := fmt.Sprintf(BadgeMarkdownTemplate, siteURL, manifest.Id, teamID, channelID)
	channelURL := fmt.Sprintf(ChannelURLTemplate, siteURL, team.Name, ch.Name)
	badgeLink := fmt.Sprintf("[%s](%s)", badge, channelURL)
	return &model.CommandResponse{
		ResponseType: model.COMMAND_RESPONSE_TYPE_EPHEMERAL,
		Text:         fmt.Sprintf("### You get Mattermost Invitation Badge!\n\n%s\n\nCopy & Paste the following link.\n**`%s`**", badgeLink, badgeLink),
	}, nil
}
