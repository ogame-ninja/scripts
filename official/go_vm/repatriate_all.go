package main

import (
	"nja/pkg/nja"
)

func main() {
	// Repatriate all resources from all planets and moons to the master coordinate

	master := "1:2:3"

	// ------------------------------------------------------------------------------
	masterCoord, _ := nja.ParseCoord(master)
	for _, celestial := range nja.GetCachedCelestials() {
		if celestial.GetCoordinate().Equal(masterCoord) {
			continue
		}
		resources, err := celestial.GetResources()
		ships, err := celestial.GetShips()
		lc, sc, _ := nja.CalcFastCargo(ships.LargeCargo, ships.SmallCargo, resources.Total())
		fleet := nja.NewFleet()
		fleet.SetOrigin(celestial)
		fleet.SetDestination(master)
		fleet.SetMission(nja.TRANSPORT)
		fleet.SetAllResources()
		fleet.AddShips(nja.LARGECARGO, lc)
		fleet.AddShips(nja.SMALLCARGO, sc)
		f, err := fleet.SendNow()
		nja.Print(f.ID, err)
		nja.SleepRandSec(1, 4)
	}
}
