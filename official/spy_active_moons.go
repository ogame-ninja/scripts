origin = GetCachedCelestial("1:2:3")           // Celestial from where to spy from
myAlliance = "TNT"                             // Should skip this alliance
spyFromSystem = 1                              // Lower limit system to spy
spyToSystem = 299                              // Upper limit system to spy
spy2send = 2                                   // Number of spy probes to send
playersToIgnore = ["Nickname1", "Nickname2"]   // Name of players to ignore
alliancesToIgnore = ["Alliance1", "Alliance2"] // Name of alliances to ignore

// Skip if player is inactive or player from my alliance
func shouldSkip(planetInfo) {
    if planetInfo == nil { return true }

    // Check ignored players list
    for playerName in playersToIgnore {
        if planetInfo.Player.Name == playerName {
            return true
        }
    }
    
    // Check ignored alliances list
    for allianceName in alliancesToIgnore {
        if (planetInfo.Alliance != nil && planetInfo.Alliance.Name == allianceName) {
            return true
        }
    }
    
    return planetInfo.Inactive || planetInfo.Vacation || (planetInfo.Alliance != nil && planetInfo.Alliance.Name == myAlliance)
}

// Sends spy probes to a coordinate
func spyCoord(coord) {
    for {
        slots = GetSlots()
        if slots.InUse < slots.Total {
                fleet = NewFleet()
                fleet.SetOrigin(origin)
                fleet.SetDestination(coord)
                fleet.SetMission(SPY)
                fleet.AddShips(ESPIONAGEPROBE, spy2send)
                fleet, err = fleet.SendNow()
                Print("Sending probes to", coord)
            break
        } else {
            Sleep(30 * 1000) // Wait 30s
        }
    }
}

// Foreach systems in the defined range, spy all active moons
if spyFromSystem < 1 { spyFromSystem = 1 }
if spyToSystem > 499 { spyToSystem = 499 }
for system = spyFromSystem; system <= spyToSystem; system++ {
    systemInfos, _ = GalaxyInfos(origin.GetCoordinate().Galaxy, system)
    for i = 1; i <= 15; i++ {
        planetInfo = systemInfos.Position(i)
        if !shouldSkip(planetInfo) && planetInfo.Moon != nil {
            pCoord = planetInfo.Coordinate
            mCoord = NewCoordinate(pCoord.Galaxy, pCoord.System, pCoord.Position, MOON_TYPE)
            spyCoord(mCoord)
        }
    }
}
