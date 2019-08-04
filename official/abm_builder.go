planetsCoord = ["1:2:3", "1:2:4", "1:2:5"] // Your planets
checkInterval = 5                          // Check every 5min

//----------------------------------------------

for {
    for planetCoord in planetsCoord {
        planet, err = GetPlanet(planetCoord)
        if err != nil {
            LogError(err)
            continue
        }
        facilities, err = planet.GetFacilities()
        if err != nil {
            LogError(err)
            continue
        }
        defenses, err = planet.GetDefense()
        if err != nil {
            LogError(err)
            continue
        }
        abm = defenses.AntiBallisticMissiles
        ipm = defenses.InterplanetaryMissiles
        possibleABM = (facilities.MissileSilo * 10) - (ipm * 2) - abm
        err = planet.BuildDefense(ANTIBALLISTICMISSILES, possibleABM)
        if err != nil {
            LogError(err)
        } else {
            Print("Building " + possibleABM + " ABM")
        }
    }
    Sleep(checkInterval * 60 * 1000) // Sleep 5min
}