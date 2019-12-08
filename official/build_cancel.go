origin = GetCachedCelestial("4:212:8")
wantedBuildingID = FUSIONREACTOR

//------------------------------------

buildingID, buildingCountdown, _, _ = origin.ConstructionsBeingBuilt()
if buildingID.Int64() != 0 {
    LogError("Building already being built", buildingID)
    return
}

// Build something
Print("Building", wantedBuildingID)
err = origin.BuildBuilding(wantedBuildingID)
if err != nil {
    LogError(err)
    return
}

buildingID, buildingCountdown, researchID, researchCountdown = origin.ConstructionsBeingBuilt()
if buildingID != wantedBuildingID {
    LogError("The building being built is not the one we wanted")
    return
}

sleepTime = buildingCountdown - 60
Print("Wait for " + ShortDur(sleepTime))
Sleep(sleepTime * 1000)

// Cancel it
Print("Cancelling", wantedBuildingID)
err = origin.CancelBuilding()
if err != nil {
    LogError(err)
    return
}
