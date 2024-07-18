package main

import (
    "nja/pkg/nja"
)

func main() {
    planetsCoord := []string{"1:106:10"} // Your planets
    checkInterval := int64(5)            // Check every 5min

    //----------------------------------------------
    
    for {
        for _, planetCoord := range planetsCoord {
            planet, err := nja.GetPlanet(planetCoord)
            if err != nil {
                nja.LogError(err)
                continue
            }
            facilities, err := planet.GetFacilities()
            if err != nil {
                nja.LogError(err)
                continue
            }
            defenses, err := planet.GetDefense()
            if err != nil {
                nja.LogError(err)
                continue
            }
            abm := defenses.AntiBallisticMissiles
            ipm := defenses.InterplanetaryMissiles
            possibleABM := (facilities.MissileSilo * 10) - (ipm * 2) - abm
            if possibleABM > 0 {
                if err := planet.BuildDefense(nja.ANTIBALLISTICMISSILES, possibleABM); err != nil {
                    nja.LogError(err)
                } else {
                    nja.Printf("Building %d ABM", possibleABM)
                }
            }
        }
        nja.Sleep(checkInterval * 60 * 1000) // Sleep 5min
    }
}
