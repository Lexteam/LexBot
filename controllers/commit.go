package controllers

import (
    "io/ioutil"
    "encoding/json"
    "github.com/google/go-github/github"
    "github.com/lexteam/lexbot/modules"
    "gopkg.in/macaron.v1"
)

func GetWebhook(ctx *macaron.Context) {
    if (ctx.Req.Header.Get("X-GitHub-Event") == "push") {
        body, _ := ioutil.ReadAll(ctx.Req.Body().ReadCloser())

        var res github.PushEvent
        json.Unmarshal(body, &res)

        modules.BOT.ChannelMessageSend(modules.CONFIG.Section("DISCORD").Key("channel").String(),
            "[" + res.Repo.Name + "] " + res.Pusher.Name + " pushed " + len(res.Commits) + " commits to " + res.Ref + "<" + res.Compare + ">")
        for _, commit := range res.Commits {
            modules.BOT.ChannelMessageSend(modules.CONFIG.Section("DISCORD").Key("channel").String(),
                res.Repo.Name + "/" + res.Ref + " " + commit.SHA + ": " + commit.Message + " (By " + commit.Author.Name + ")")
        }
    }
}
