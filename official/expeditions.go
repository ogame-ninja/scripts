// WARNING: This script doesn't care about the "reserved slots"
origin = GetCachedCelestial("4:208:8")
minSystem = 1
maxSystem = 10
ships = {LIGHTFIGHTER: 2, LARGECARGO: 3}
expeditionDuration = 1

//----------------------------------------

func sendExpedition() {
    randomSystem = Random(minSystem, maxSystem)
    destination = NewCoordinate(origin.GetCoordinate().Galaxy, randomSystem, 16, PLANET_TYPE)
    fleet = NewFleet()
    fleet.SetOrigin(origin.GetID())
    fleet.SetDestination(destination)
    fleet.SetSpeed(HUNDRED_PERCENT)
    fleet.SetMission(EXPEDITION)
    for shipID, nbr in ships {
        fleet.AddShips(shipID, nbr)
    }
    fleet.SetDuration(expeditionDuration)
    return fleet.SendNow()
}

for {
    fleets, slots = GetFleets()
    
    // Find next expedition fleet that will come back
    bigNum = 999999999
    minSecs = bigNum
    for fleet in fleets {
        if fleet.Mission == EXPEDITION {
            minSecs = Min(fleet.BackIn, minSecs)
        }
    }
    
    // Sends new expeditions
    expeditionsPossible = slots.ExpTotal - slots.ExpInUse
    for expeditionsPossible > 0 {
        newFleet, err = sendExpedition()
        if err != nil {
            LogError(err)
            break
        } else {
            Print(newFleet)
            minSecs = Min(newFleet.BackIn, minSecs)
            expeditionsPossible--
        }
        Sleep(Random(10000, 20000))
    }
    
    // If we didn't found any expedition fleet and didn't create any, let's wait 5min
    if minSecs == bigNum {
        minSecs = 5 * 60
    }
    
    Sleep((minSecs + 10) * 1000) // Sleep until one of the expedition fleet come back
}
