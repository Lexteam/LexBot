package modules

import (
    "fmt"
    "github.com/bwmarrin/discordgo"
)

var (
    BOT *discordgo.Session
)

func InitBot() {
    var err error

    BOT, err = discordgo.New(
        CONFIG.Section("DISCORD").Key("email").String(),
        CONFIG.Section("DISCORD").Key("password").String())
    if err != nil {
        fmt.Println("Error creating Discord session", err)
    }
}
