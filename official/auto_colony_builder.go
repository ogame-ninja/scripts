// Largely inspired by https://github.com/ogame-ninja/scripts/blob/master/community/cremefresh55/AutoColonyBuilder.go

coord = "P:1:2:3"
a = [
    SOLARPLANT,  // lvl 1
    METALMINE,   // lvl 1
    METALMINE,   // lvl 2
    SOLARPLANT,  // lvl 2
    METALMINE,   // lvl 3
    METALMINE,   // lvl 4
    SOLARPLANT,  // lvl 3
    CRYSTALMINE, // lvl 1
    SOLARPLANT,  // ...
    METALMINE,
    CRYSTALMINE,
    CRYSTALMINE,
    SOLARPLANT,
    DEUTERIUMSYNTHESIZER,
    CRYSTALMINE,
    SOLARPLANT,
    METALMINE,
    METALMINE,
    SOLARPLANT,
    CRYSTALMINE,
    DEUTERIUMSYNTHESIZER,
    SOLARPLANT,
    DEUTERIUMSYNTHESIZER,
    DEUTERIUMSYNTHESIZER,
    SOLARPLANT,
    DEUTERIUMSYNTHESIZER,
    ROBOTICSFACTORY,
    ROBOTICSFACTORY,
    RESEARCHLAB,
    SHIPYARD,
    CRYSTALMINE,
    SHIPYARD,
    SOLARPLANT,
    DEUTERIUMSYNTHESIZER,
    METALMINE,
    ENERGYTECHNOLOGY,
    COMBUSTIONDRIVE,
    SOLARPLANT,
    CRYSTALMINE,
    METALMINE,
    COMBUSTIONDRIVE,
    SMALLCARGO,
]

//-----------------------------------------------------------------------

// Return either or not a ogameID is in the current production line
func itemInProductionLine(oid) {
    productionLine, _, _ = celestial.GetProduction()
    for item in productionLine {
        prodID = item.ID
        if prodID == oid {
            return true
        }
    }
    return false
}

// If we build a colony, skip researches and lab
planets = GetCachedPlanets()
skipResearches = len(planets) > 1

// Get celestial object and buildings/researches information
celestial, _ = GetCelestial(coord)
supplies, facilities, _, _, researches, err = GetTechs(celestial)

// Map that keep track of what level to build
m = {}

for i = 0; i < len(a); i++ {
    for { // Loop to retry the same step in the list if something went wrong
        oid = a[i]
        m[oid]++
        if oid == RESEARCHLAB && skipResearches {
            break
        }
        wantedLvl = m[oid]
        lvl = 0
        nbr = 0
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
            ships, err = celestial.GetShips()
            if err != nil {
                LogError(err)
                SleepMin(1)
                continue
            }
            nbrOnCelestial = ships.ByID(oid)
            if nbrOnCelestial >= wantedLvl {
                break
            }
            
            SleepMs(1000)
            
            if itemInProductionLine(oid) {
                SleepMs(1000)
                break
            }
        }
        if lvl >= wantedLvl {
            break
        }
        
        // Actually build the OGameID
        Print("Wants to build " + oid + " lvl " + wantedLvl)
        err = celestial.Build(oid, nbr)
        if err != nil {
            LogError(err)
            SleepMs(1000)
            continue
        }
        SleepMs(100)
        
        // Then verify that the item is actually being built
        wait = 1
        found = true
        if oid.IsShip() || oid.IsDefense() {
            if !itemInProductionLine(oid) {
                found = false
            }
        } else {
            buildingID, buildingCountdown, researchID, researchCountdown = celestial.ConstructionsBeingBuilt()
            if (oid.IsResourceBuilding() || oid.IsFacility()) && buildingID == oid {
                wait = buildingCountdown + 10
            } else if oid.IsTech() && researchID == oid {
                wait = researchCountdown + 10
            } else {
                found = false
            }
        }
        if !found {
            LogError(oid + " is not being built, wait 1min")
            SleepMin(1)
            continue
        }
        
        print(oid + " was built at step: " + i + ", wait " + wait)
        SleepSec(wait)
        supplies, facilities, _, _, researches, err = GetTechs(celestial) // Update values
        break
    }
}
Print("Done building list")
