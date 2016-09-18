package controllers

import (
    "io/ioutil"
    "encoding/json"
    "github.com/google/go-github/github"
    "github.com/lexteam/lexbot/modules"
    "gopkg.in/macaron.v1"
    "strconv"
    "strings"
    "net/http"
    "log"
    "net/url"
)

func GetWebhook(ctx *macaron.Context) {
    if (ctx.Req.Header.Get("X-GitHub-Event") == "push") {
        body, _ := ioutil.ReadAll(ctx.Req.Body().ReadCloser())

        var res github.PushEvent
        json.Unmarshal(body, &res)

        branch := strings.Split(*res.Ref, "/")[2]
        compare := shortenUrl(*res.Compare)

        modules.BOT.ChannelMessageSend(modules.CONFIG.Section("DISCORD").Key("channel").String(),
            "[" + *res.Repo.Name + "] " + *res.Pusher.Name + " pushed " + strconv.Itoa(len(res.Commits)) + " commits to " + branch + " " + compare)

        for _, commit := range res.Commits {
            message := strings.Split(*commit.Message, "\n")[0]
            id := (*commit.ID)[0:8]

            modules.BOT.ChannelMessageSend(modules.CONFIG.Section("DISCORD").Key("channel").String(),
                *res.Repo.Name + "/" + branch + " " + id + ": " + message + " (By " + *commit.Author.Name + ")")
        }
    }
}

func shortenUrl(longUrl string) string {
    values := url.Values{
        "url": { longUrl },
    }
    response, err := http.PostForm("https://git.io", values)
    if err != nil {
        log.Println("failed to shorten url")
        return longUrl
    }
    shortUrl := response.Header.Get("Location")
    defer response.Body.Close()
    return shortUrl
}
