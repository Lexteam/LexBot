package controllers

import (
    "strconv"
    "io/ioutil"
    "encoding/json"
    "github.com/google/go-github/github"
    "github.com/lexteam/lexbot/modules"
    "github.com/lexteam/lexbot/utils"
    "gopkg.in/macaron.v1"
)

func GetWebhook(ctx *macaron.Context) {
    if (ctx.Req.Header.Get("X-GitHub-Event") == "push") {
        body, _ := ioutil.ReadAll(ctx.Req.Body().ReadCloser())

        var res github.PushEvent
        json.Unmarshal(body, &res)

        branch := utils.GetBranchName(*res.Ref)
        compare := utils.GetGitioUrl(*res.Compare)

        modules.BOT.ChannelMessageSend(modules.CONFIG.Section("DISCORD").Key("channel").String(),
            "[" + *res.Repo.Name + "] " + *res.Pusher.Name + " pushed " + strconv.Itoa(len(res.Commits)) + " commits to " + branch + " " + compare)

        for _, commit := range res.Commits {
            message := utils.GetShortCommitMessage(*commit.Message)
            id := utils.GetShortCommitID(*commit.ID)

            modules.BOT.ChannelMessageSend(modules.CONFIG.Section("DISCORD").Key("channel").String(),
                *res.Repo.Name + "/" + branch + " " + id + ": " + message + " (By " + *commit.Author.Name + ")")
        }
    }
}
