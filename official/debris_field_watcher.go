debrisCoord, _ = ParseCoord("D:1:2:3")
delaySecs = 2
for {
    systemInfo, _ = GalaxyInfos(debrisCoord.Galaxy, debrisCoord.System)
    planetInfo = systemInfo.Position(debrisCoord.Position)
    if planetInfo == nil {
        LogError("planet not found")
        break
    }
    if planetInfo.Debris.RecyclersNeeded == 0 {
        hour, min, sec = Clock()
        SendTelegram(TELEGRAM_CHAT_ID, "Debris field is gone at " + hour + "h" + min + "m" + sec)
        break
    }
    Sleep(delaySecs * 1000)
}
Print("Debris field script exit")
