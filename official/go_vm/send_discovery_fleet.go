package main

import (
	"nja/pkg/nja"
	"strings"
)

func main() {
	origin := "1:2:3"
	systemsRange := 10

	//---------------------------
	c := nja.GetCachedCelestial(origin)
	cCoord := c.GetCoordinate()
	systems := nja.GetSystemsInRangeAsc(cCoord.System, int64(systemsRange))

	totalSlots := getUpdatedSlots()

	for _, system := range systems {
		availCoords, err := nja.CoordinatesAvailableForDiscoveryFleet(c, cCoord.Galaxy, system)
		if err != nil {
			nja.LogError(err)
			nja.SleepSec(5)
			continue
		}
		if len(availCoords) == 0 {
			nja.LogDebugf("no available coords in [%d, %d]", cCoord.Galaxy, system)
			continue
		}
		i := 0
		for {
			if totalSlots == 0 {
				nja.LogDebug("no slots available, wait 8-10min")
				nja.SleepRandMin(8, 10)
				totalSlots = getUpdatedSlots()
				continue
			}
			destination := availCoords[i]
			err = nja.SendDiscoveryFleet(origin, destination)
			if err != nil {
				nja.LogError(err)
				if strings.Contains(err.Error(), "Maximum number of fleets reached") {
					nja.SleepRandMin(8, 10)
					totalSlots = getUpdatedSlots()
				}
				if strings.Contains(err.Error(), "Not enough resources") {
					nja.SleepRandMin(8, 10)
					totalSlots = getUpdatedSlots()
				}
				nja.SleepSec(5)
				continue
			}
			nja.LogDebugf("Sent discovery fleet to %s", destination)
			totalSlots--
			nja.SleepRandMs(1000, 2000)
			i++
			if i >= len(availCoords) {
				break
			}
		}
		nja.SleepRandMs(1500, 3000)
	}
}

// Get how many slots we can use for discovery
func getUpdatedSlots() int64 {
	slots := nja.GetSlots()
	return slots.Total - slots.InUse - nja.GetFleetSlotsReserved()
}
