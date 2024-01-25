origin = "3:92:12"
systemsRange = 10

//---------------------------
strings = import("strings")
c = GetCachedCelestial(origin)
cCoord = c.GetCoordinate()
systems = GetSystemsInRangeAsc(cCoord.System, systemsRange)

// Get how many slots we can use for discovery
func getUpdatedSlots() {
    slots = GetSlots()
    return slots.Total - slots.InUse - GetFleetSlotsReserved()
}

totalSlots = getUpdatedSlots()

for system in systems {
    availCoords, err = CoordinatesAvailableForDiscoveryFleet(c, cCoord.Galaxy, system)
    if err != nil {
        LogError(err)
        SleepSec(5)
        continue
    }
    if len(availCoords) == 0 {
        LogDebug("no available coords in [" + cCoord.Galaxy + ", " + system + "]")
        continue
    }
    i = 0
    for {
        if totalSlots == 0 {
            LogDebug("no slots available, wait 8-10min")
            SleepRandMin(8, 10)
            totalSlots = getUpdatedSlots()
            continue
        }
        destination = availCoords[i]
        err = SendDiscoveryFleet(origin, destination)
        if err != nil {
            LogError(err)
            if strings.Contains(err.Error(), "Maximum number of fleets reached") {
                SleepRandMin(8, 10)
                totalSlots = getUpdatedSlots()
            }
            if strings.Contains(err.Error(), "Not enough resources") {
                SleepRandMin(8, 10)
                totalSlots = getUpdatedSlots()
            }
            SleepSec(5)
            continue
        }
        LogDebug("Sent discovery fleet to " + destination)
        totalSlots--
        SleepRandMs(1000, 2000)
        i++
        if i >= len(availCoords) {
            break
        }
    }
    SleepRandMs(1500, 3000)
}
