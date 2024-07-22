// Largely inspired by https://github.com/ogame-ninja/scripts/blob/master/community/cremefresh55/AutoColonyBuilder.go

package main

import (
	"nja/pkg/nja"
	"nja/pkg/ogame"
	"nja/pkg/wrapper"
)

func main() {

	coord := "P:1:106:10"
	var a = []ogame.ID{
		nja.SOLARPLANT,  // lvl 1
		nja.METALMINE,   // lvl 1
		nja.METALMINE,   // lvl 2
		nja.SOLARPLANT,  // lvl 2
		nja.METALMINE,   // lvl 3
		nja.METALMINE,   // lvl 4
		nja.SOLARPLANT,  // lvl 3
		nja.CRYSTALMINE, // lvl 1
		nja.SOLARPLANT,  // ...
		nja.METALMINE,
		nja.CRYSTALMINE,
		nja.CRYSTALMINE,
		nja.SOLARPLANT,
		nja.DEUTERIUMSYNTHESIZER,
		nja.CRYSTALMINE,
		nja.SOLARPLANT,
		nja.METALMINE,
		nja.METALMINE,
		nja.SOLARPLANT,
		nja.CRYSTALMINE,
		nja.DEUTERIUMSYNTHESIZER,
		nja.SOLARPLANT,
		nja.DEUTERIUMSYNTHESIZER,
		nja.DEUTERIUMSYNTHESIZER,
		nja.SOLARPLANT,
		nja.DEUTERIUMSYNTHESIZER,
		nja.ROBOTICSFACTORY,
		nja.ROBOTICSFACTORY,
		nja.RESEARCHLAB,
		nja.SHIPYARD,
		nja.CRYSTALMINE,
		nja.SHIPYARD,
		nja.SOLARPLANT,
		nja.DEUTERIUMSYNTHESIZER,
		nja.METALMINE,
		nja.ENERGYTECHNOLOGY,
		nja.COMBUSTIONDRIVE,
		nja.SOLARPLANT,
		nja.CRYSTALMINE,
		nja.METALMINE,
		nja.COMBUSTIONDRIVE,
		nja.SMALLCARGO,
	}

	//-----------------------------------------------------------------------

	// If we build a colony, skip researches and lab
	planets := nja.GetCachedPlanets()
	skipResearches := len(planets) > 1

	// Get celestial object and buildings/researches information
	celestial, _ := nja.GetCelestial(coord)
	supplies, facilities, _, _, researches, err := nja.GetTechs(celestial)

	// Map that keep track of what level to build
	m := make(map[ogame.ID]int64)

	for i := 0; i < len(a); i++ {
		for { // Loop to retry the same step in the list if something went wrong
			oid := a[i]
			m[oid]++
			if oid == nja.RESEARCHLAB && skipResearches {
				break
			}
			wantedLvl := m[oid]
			lvl := int64(0)
			nbr := int64(0)
			if oid.IsResourceBuilding() {
				lvl = supplies.ByID(oid)
			} else if oid.IsFacility() {
				lvl = facilities.ByID(oid)
			} else if oid.IsTech() {
				if skipResearches {
					break
				}
				lvl = researches.ByID(oid)
			} else if oid.IsShip() || oid.IsDefense() {
				nbr = 1

				// Check the existing ships in shipyard, if we already have it, skip
				ships, err := celestial.GetShips()
				if err != nil {
					nja.LogError(err)
					nja.SleepMin(1)
					continue
				}
				nbrOnCelestial := ships.ByID(oid)
				if nbrOnCelestial >= wantedLvl {
					break
				}

				nja.SleepMs(1000)

				if itemInProductionLine(celestial, oid) {
					nja.SleepMs(1000)
					break
				}
			}
			if lvl >= wantedLvl {
				break
			}

			// Actually build the OGameID
			nja.Printf("Wants to build %d lvl %d", oid, wantedLvl)
			err = celestial.Build(oid, nbr)
			if err != nil {
				nja.LogError(err)
				nja.SleepMs(1000)
				continue
			}
			nja.SleepMs(100)

			// Then verify that the item is actually being built
			wait := int64(1)
			found := true
			if oid.IsShip() || oid.IsDefense() {
				if !itemInProductionLine(oid) {
					found = false
				}
			} else {
				buildingID, buildingCountdown, researchID, researchCountdown, _, _, _, _ := celestial.ConstructionsBeingBuilt()
				if (oid.IsResourceBuilding() || oid.IsFacility()) && buildingID == oid {
					wait = buildingCountdown + 10
				} else if oid.IsTech() && researchID == oid {
					wait = researchCountdown + 10
				} else {
					found = false
				}
			}
			if !found {
				nja.LogErrorf("%d is not being built, wait 1min", oid)
				nja.SleepMin(1)
				continue
			}

			nja.Printf("%d was built at step: %d, wait %d", oid, i, wait)
			nja.SleepSec(wait)
			supplies, facilities, _, _, researches, err = nja.GetTechs(celestial) // Update values
			break
		}
	}
	nja.Print("Done building list")
}

// Return either or not a ogameID is in the current production line
func itemInProductionLine(celestial wrapper.Celestial, oid ogame.ID) bool {
	productionLine, _, _ := celestial.GetProduction()
	for _, item := range productionLine {
		prodID := item.ID
		if prodID == oid {
			return true
		}
	}
	return false
}
