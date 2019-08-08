package puddle

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/jmoiron/sqlx"
)

type Commands struct {
	conn *sqlx.DB
}

func (c *Commands) Help(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	if !strings.HasPrefix(m.Content, "!help") {
		return
	}
	_, err := s.ChannelMessageSend(m.ChannelID, `ping - test connection
`)
	if err != nil {
		fmt.Println(err)
	}

}

func (c *Commands) Ping(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	if !strings.HasPrefix(m.Content, "!ping") {
		return
	}
	_, err := s.ChannelMessageSend(m.ChannelID, "pong!")
	if err != nil {
		fmt.Println(err)
	}
}
