package main

import (
	"nja/pkg/nja"
	"nja/pkg/ogame"
	"nja/pkg/wrapper"
)

var (
	origin            = "1:106:10"                         // Celestial from where to spy from
	spyFromSystem     = 1                                  // Lower limit system to spy
	spyToSystem       = 299                                // Upper limit system to spy
	spy2send          = 2                                  // Number of spy probes to send
	playersToIgnore   = []string{"Nickname1", "Nickname2"} // Name of players to ignore
	alliancesToIgnore = []string{"Alliance1", "Alliance2"} // Name of alliances to ignore
)

func main() {
	//--------------------------------------------------------------------------------

	originCelestial := nja.GetCachedCelestial(origin)
	if originCelestial == nil {
		nja.LogError("invalid origin coordinate", origin)
		return
	}

	// Foreach systems in the defined range, spy all active moons
	if spyFromSystem < 1 {
		spyFromSystem = 1
	}
	if spyToSystem > int(nja.SYSTEMS) {
		spyToSystem = int(nja.SYSTEMS)
	}
	for system := spyFromSystem; system <= spyToSystem; system++ {
		systemInfos, _ := nja.GalaxyInfos(originCelestial.GetCoordinate().Galaxy, int64(system))
		for i := 1; i <= 15; i++ {
			planetInfo := systemInfos.Position(int64(i))
			if !shouldSkip(planetInfo) && planetInfo.Moon != nil {
				pCoord := planetInfo.Coordinate
				mCoord := pCoord.Moon()
				spyCoord(mCoord, originCelestial)
			}
		}
	}
}

// Sends spy probes to a coordinate
func spyCoord(coord ogame.Coordinate, originCelestial wrapper.Celestial) {
	nja.Print(originCelestial)
	for {
		slots := nja.GetSlots()
		if slots.InUse < slots.Total {
			fleet := nja.NewFleet()
			fleet.SetOrigin(originCelestial)
			fleet.SetDestination(coord)
			fleet.SetMission(nja.SPY)
			fleet.AddShips(nja.ESPIONAGEPROBE, int64(spy2send))
			_, _ = fleet.SendNow()
			nja.Print("Sending probes to", coord)
			break
		} else {
			nja.SleepSec(30) // Wait 30s
		}
	}
}

// Skip if player is inactive or player from my alliance
func shouldSkip(planetInfo *ogame.PlanetInfos) bool {
	if planetInfo == nil {
		return true
	}

	// Check ignored players list
	for _, playerName := range playersToIgnore {
		if planetInfo.Player.Name == playerName {
			return true
		}
	}

	// Check ignored alliances list
	for _, allianceName := range alliancesToIgnore {
		if planetInfo.Alliance != nil && planetInfo.Alliance.Name == allianceName {
			return true
		}
	}

	return planetInfo.Inactive || planetInfo.Vacation
}
