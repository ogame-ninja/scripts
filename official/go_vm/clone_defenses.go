package main

import (
    "ogame/pkg/ogame"
    "nja"
)

func main() {
    master := nja.GetCachedCelestial("1:106:10")
    slave := nja.GetCachedCelestial("1:103:11")
    
    for {
        masterDef, _ := master.GetDefense()
        slaveDef, _ := slave.GetDefense()
        prodQueue, _, _ := slave.GetProduction() // get build queue
        
        // DefencesArr is a built-in array that contains all the defense entity IDs
        for _, defID := range nja.DefencesArr {
            delta := masterDef.ByID(defID) - slaveDef.ByID(defID)
            
            // Remove defenses that are already in build queue
            for _, item := range prodQueue {
                if item.ID == defID {
                    delta -= item.Nbr
                }
            }
            
            if delta > 0 {
                slave.BuildDefense(defID, delta)
                nja.Print("Build", delta, defID) // eg: Build 12 RocketLauncher
            }
        }
        
        nja.Sleep(60 * 60 * 1000) // Check again in 1 hour
    }
}
