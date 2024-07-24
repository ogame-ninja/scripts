package main

import (
	"nja/pkg/nja"
	"nja/pkg/wrapper"
)

func main() {
	master := findCelestialWithHigherFleetValue()
	nja.Print("master is", master.GetCoordinate())
}

func findCelestialWithHigherFleetValue() wrapper.Celestial {
	lfBonuses, _ := nja.GetCachedLfBonuses()
	fleets, _ := nja.GetFleets()
	var master wrapper.Celestial
	var maxVal int64
	for _, celestial := range nja.GetCachedCelestials() {
		ships, _ := celestial.GetShips()
		value := ships.FleetValue(lfBonuses)
		coord := celestial.GetCoordinate()
		for _, fleet := range fleets {
			if (fleet.Origin.Equal(coord) && fleet.Mission != nja.PARK) || (fleet.Destination.Equal(coord) && fleet.Mission == nja.PARK) {
				value += fleet.Ships.FleetValue(lfBonuses)
			}
		}
		if value > maxVal {
			maxVal = value
			master = celestial
		}
	}
	return master
}
