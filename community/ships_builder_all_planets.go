// Build ships periodically on a given list of planets
planetsToBuildOn = ["1:1:1", "2:2:2"]
interval = Random(30*60*1000, 60*60*1000) // 30-60min
toBuild = {
	LIGHTFIGHTER: 6,
	HEAVYFIGHTER: 4,
}

func millisecondsToTime(milliseconds) {
    minutes = (milliseconds / 1000) / 60;
    return minutes;
}

for {
    for planet in planetsToBuildOn {
        // Get planet celestial
        planetCelestial, _ = GetCelestial(planet)
        // Iterate over ships to build
        for unitID, nbrToBuild in toBuild {
            // See how many ships we can build
            unitPrice = GetPrice(unitID, 1)
            resources, _ = planetCelestial.GetResources()
            canBuild = resources.Div(unitPrice)
            if canBuild > 0 {
                if canBuild > nbrToBuild {
                    canBuild = nbrToBuild
                }
                Print("Building: "+canBuild+"/"+nbrToBuild+" "+ID2Str(unitID)+" on "+planet)
                Build(planetCelestial.GetID(), unitID, canBuild)
            } else {
                LogWarn("Not enough resources on "+planet+" to build at least 1 of "+nbrToBuild+" "+ID2Str(unitID))
            }
        }
    }
    Print("All ships built. Resume in "+millisecondsToTime(interval)+" minutes.")
    Sleep(interval)
}