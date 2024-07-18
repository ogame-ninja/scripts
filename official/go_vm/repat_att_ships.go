package main

import (
    "nja/pkg/ogame"
    "nja/pkg/nja"
)

func main() {
    masterCoord := "1:2:3"

    //------------------------------------------------------------------------------
    
    attShips := []ogame.ID{nja.LIGHTFIGHTER, nja.HEAVYFIGHTER, nja.CRUISER,
        nja.BATTLESHIP, nja.BOMBER, nja.DESTROYER, nja.DEATHSTAR, nja.BATTLECRUISER}
    
    celestials := nja.GetCachedCelestials()
    for _, celestial := range celestials {
        ships, _ := celestial.GetShips()
        fleet := nja.NewFleet()
        fleet.SetOrigin(celestial)
        fleet.SetDestination(masterCoord)
        fleet.SetMission(nja.PARK)
        for _, shipID := range attShips {
            fleet.AddShips(shipID, ships.ByID(shipID))
        }
        f, err := fleet.SendNow()
        if err != nil {
            nja.LogError(err)
        } else {
            nja.Print(f)
        }
    }
}
