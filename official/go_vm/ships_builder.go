package main

import (
	"nja/pkg/nja"
	"nja/pkg/ogame"
)

func main() {
	origin := nja.GetCachedCelestial("1:106:10")
	toBuild := map[ogame.ID]int64{nja.LIGHTFIGHTER: 6, nja.HEAVYFIGHTER: 4}
	for {
		nja.Print("Need to build:", toBuild)
		for unitID, nbr := range toBuild {
			unitPrice := nja.GetPrice(unitID, 1)
			res, _ := origin.GetResources()
			canBuild := res.Div(unitPrice)
			if canBuild > 0 {
				if canBuild > nbr {
					canBuild = nbr
				}
				nja.Build(origin.GetID(), unitID, canBuild)
				nja.Printf("Building: %d %s", canBuild, nja.ID2Str(unitID))
				toBuild[unitID] = toBuild[unitID] - canBuild
				if toBuild[unitID] <= 0 {
					delete(toBuild, unitID)
				}
			}
		}
		if len(toBuild) == 0 {
			break
		}
		nja.Sleep(5 * 60 * 1000) // 5min
	}
	nja.Print("All ships built")
}
