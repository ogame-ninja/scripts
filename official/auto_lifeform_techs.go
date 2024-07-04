coord = "1:106:10"
wantedTier1 = [
    INTERGALACTICENVOYS,
    ACOUSTICSCANNING,
    FUSIONDRIVES,
    TELEKINETICTRACTORBEAM,
    ENHANCEDSENSORTECHNOLOGY,
    AUTOMATEDTRANSPORTLINES,
]
useArtefacts = true

errors = import("errors")

// Hack to get the lifeform type for this planet
lfBuildings, _ = GetLfBuildings(coord)
lfType = lfBuildings.LifeformType

// Check the slots (up to slotNumber), and ensure it matches the wanted techs.
// Return true if all techs are the ones we want.
func checkSlots() {
    lfResearch, err = GetLfResearchDetails(coord)
    if err != nil {
        return false, err
    }
    for i = 0; i < 6; i++ {
        slot = lfResearch.Slots[i]
        if !slot.TechID.IsSet() {
            return false, errors.New("tech not set")
        }
        if wantedTier1[i] != slot.TechID {
            return false, nil
        }
    }
    return true, nil
}

// Return either or not the tech can be selected directly for the lfType
func canSelectTech(lfType, techID) {
    return (lfType == HUMANS  && techID in LfTechnologiesHumansArr ) || (lfType == ROCKTAL && techID in LfTechnologiesRocktalArr) || (lfType == MECHAS  && techID in LfTechnologiesMechasArr ) || (lfType == KAELESH && techID in LfTechnologiesKaeleshArr)
}

for {
    for {
        lfResearch, err = GetLfResearchDetails(coord)
        if err != nil {
            LogError("wait 5-6 min", err)
            SleepRandMin(5, 6)
            continue
        }
        slotNumber = lfResearch.AvailableSlot()
        if slotNumber >= 1 && slotNumber <= 6 {
            
            wantedTech = wantedTier1[slotNumber-1]
            
            // Select the research directly if the one we want is the one from our own lf type
            if canSelectTech(lfType, wantedTech) {
                LogInfo("select research for #", slotNumber)
                _ = SelectLfResearchSelect(coord, slotNumber)
            } else { // otherwise we pick randomly
                if useArtefacts {
                    if lfResearch.ArtefactsCollected >= 200 {
                        err = SelectLfResearchArtifacts(coord, slotNumber, wantedTech)
                        if err != nil {
                            LogError(err)
                        } else {
                            LogInfo("bought with artefacts #", slotNumber)
                        }
                    } else {
                        LogInfo("not enough artefacts ", lfResearch.ArtefactsCollected)
                    }
                } else {
                    _ = SelectLfResearchRandom(coord, slotNumber)
                }
            }
        }
        if slotNumber >= 1 && slotNumber <= 6 {
            LogError("wait 5-10 sec")
            SleepRandSec(5, 10)
            continue
        }
        break
    }

    slotsOk, err = checkSlots(6)
    if err != nil {
        LogInfo("wait 1-2 hours", err)
        SleepRandHour(1, 2)
        continue
    }
    if slotsOk {
        if slotNumber == 6 { // if we just set the 6th slot, and all matches, then we can stop here.
            LogInfo("we now have the wanted techs")
            return
        }
    } else {
        LogInfo("reset the tree")
        err = FreeResetTree()
        if err != nil {
            LogError(err)
        }
    }

    LogInfo("wait 1-2 hours")
    SleepRandHour(1, 2)
}
