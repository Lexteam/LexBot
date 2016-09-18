package main

import (
    "github.com/lexteam/lexbot/controllers"
    "github.com/lexteam/lexbot/modules"
    "gopkg.in/macaron.v1"
    "fmt"
)

func main() {
    // macaron
    s := macaron.Classic()

    // init all the stuff
    modules.InitConfig()
    modules.InitBot()

    // Routes
    s.Post("/commit", controllers.GetWebhook)

    // Run webserver
    s.Run()

    // Run Discord bot
    err := modules.BOT.Open()
    if err != nil {
        fmt.Println("error opening connection", err)
        return
    }
}
