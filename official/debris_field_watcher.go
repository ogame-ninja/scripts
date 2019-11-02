debrisCoord = NewCoordinate(1, 2, 3, DEBRIS_TYPE)
delaySecs = 2
for {
    systemInfo, _ = GalaxyInfos(debrisCoord.Galaxy, debrisCoord.System)
    planetInfo = systemInfo.Position(debrisCoord.Position)
    if planetInfo == nil {
        LogError("planet not found")
        break
    }
    if planetInfo.Debris.RecyclersNeeded == 0 {
        SendTelegram(TELEGRAM_CHAT_ID, "Debris field is gone")
        break
    }
    Sleep(delaySecs * 1000)
}
Print("Debris field script exit")
