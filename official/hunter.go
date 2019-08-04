minActivity   = 30   // This number has to be higher than 15
checkInterval = 5    // Check every 5min
targets = {          // Targets to watch
    "Target1": [
        "5:116:1",
        "5:116:8",
        "5:116:12",
        "5:117:9",
        "5:119:8",
        "5:126:8",
        "5:208:10",
    ],
}

//------------------------------------------------------------------------------

if minActivity <= 15 {
    LogError("minActivity must be greater than 15")
    return
}

for {
    for targetName, planetsCoord in targets {
        hasActivity = false
        visitedCoords = {}
        for planetCoord in planetsCoord {
            coord, err = ParseCoord(planetCoord)
            if err != nil {
                LogError("Failed to parse coord " + planetCoord)
                continue
            }
            cached = visitedCoords[coord.Galaxy + ":" + coord.System]
            systemInfos = nil
            if cached == nil {
                systemInfos, err = GalaxyInfos(coord.Galaxy, coord.System)
                if err != nil {
                    LogError("Failed to get system infos for " + coord)
                    continue
                }
                visitedCoords[coord.Galaxy + ":" + coord.System] = systemInfos
            } else {
                systemInfos = cached
            }
            planetInfos, err = systemInfos.Position(coord.Position)
            if err != nil {
                LogError("Failed to get planet infos for " + coord)
                continue
            }
            if planetInfos == nil {
                LogError("Failed to get planet infos for " + coord)
                continue
            }
            if planetInfos.Activity > 0 && planetInfos.Activity < minActivity {
                hasActivity = true
                break
            }
            if planetInfos.Moon != nil && planetInfos.Moon.Activity > 0 && planetInfos.Moon.Activity < minActivity {
                hasActivity = true
                break
            }
        }
        if hasActivity {
            Print("Some activities detected for " + targetName)
        } else {
            Print("No activty in the last " + minActivity + "min for " + targetName)
        }
    }
    Sleep(checkInterval * 60 * 1000) // Sleep 5min
}