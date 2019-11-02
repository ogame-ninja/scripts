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
        hour, min, sec = Clock()
        Print("Debris field is gone at " + hour + "h" + min + "m" + sec)
        break
    }
    Sleep(delaySecs * 1000)
}
Print("Debris field script exit")
