package main

import (
    "fmt"
    "nja/pkg/nja"
)

func main() {
    debrisCoord, _ := nja.ParseCoord("D:1:111:11")
    delaySecs := int64(2)
    for {
        systemInfo, _ := nja.GalaxyInfos(debrisCoord.Galaxy, debrisCoord.System)
        planetInfo := systemInfo.Position(debrisCoord.Position)
        if planetInfo == nil {
            panic("planet not found")
        }
        if planetInfo.Debris.RecyclersNeeded == 0 {
            hour, min, sec := nja.Clock()
            msg := fmt.Sprintf("Debris field is gone at %dh%dm%d", hour, min, sec)
            nja.SendTelegram(nja.TELEGRAM_CHAT_ID, msg)
            break
        }
        nja.Sleep(delaySecs * 1000)
    }
    println("Debris field script exit")
}
