//==== This script is created by RockClubKASHMIR ====

fromSystem = 1 // Your can change this value as you wish
toSystem = 499 // Your can change this value as you wish
Rnbr = 1  // When Rnbr = 1, the script will search only debris for minimum 2 Recyclers. You can change this value as you wish
//----
curSystem = fromSystem
origin = nil
flts = 0
nbr = 0
err = 0
i = 1
for celestial in GetCachedCelestials() {
    ships, _ = celestial.GetShips()
    if ships.Recycler > flts {
        flts = ships.Recycler
        origin = celestial // Your Planet(or Moon), with more Recyclers
    }
}
if origin != nil {
    Print("Your origin is "+origin.Coordinate)
    for system = fromSystem; system <= toSystem; system++ {
        slots = GetSlots().InUse+GetFleetSlotsReserved()
        systemInfos, err = GalaxyInfos(origin.GetCoordinate().Galaxy, system)
        planetInfo = systemInfos.Position(i)
        if slots < GetSlots().Total {
            if planetInfo != nil {
                Sleep(Random(1000, 2000)) // For more human reation...
                Print("Checking "+planetInfo.Coordinate)
                if planetInfo.Debris.RecyclersNeeded > Rnbr { 
                    ships, _ = origin.GetShips()
                    Print("Found Metal:"+planetInfo.Debris.Metal+" and Crystal:"+planetInfo.Debris.Crystal+" at "+planetInfo.Coordinate+", need "+planetInfo.Debris.RecyclersNeeded+" Recyclers!")
                    Sleep(Random(6*1000, 12*1000)) // Pause between 6 and 12 seconds
                    f = NewFleet()
                    f.SetOrigin(origin)
                    f.SetDestination(planetInfo.Coordinate)
                    f.SetSpeed(HUNDRED_PERCENT)
                    f.SetMission(RECYCLEDEBRISFIELD)
                    if planetInfo.Debris.RecyclersNeeded > ships.Recycler {
                        nbr = ships.Recycler
                    } else {nbr = planetInfo.Debris.RecyclersNeeded}
                    f.AddShips(RECYCLER, nbr)
                    a, err = f.SendNow()
                    if err == nil {
                        Print(nbr+" Recyclers are sended successfully!")
                        SendTelegram(TELEGRAM_CHAT_ID, nbr+" Recyclers are sended successfully!")
                    } else {
                        Print("The Recyclers are NOT sended! "+err)
                        SendTelegram(TELEGRAM_CHAT_ID, "The Recyclers are NOT sended! "+err)
                        curSystem = system-1
                        slots = GetSlots().Total
                    }
                }
            }
            if i < 15 {
            i++
            } else {i = 1}
            
        } else {
            for slots == GetSlots().Total {
                if err != 0 {
                    Print("Please wait till ships lands! Recheck after "+ShortDur(2*60))
                    Sleep(2*60*1000) // 2 minutes
                    ships, _ = origin.GetShips()
                    if ships.Recycler > 0 {slots = GetSlots().InUse+GetFleetSlotsReserved()}
                } else {
                    Print("All Fleet slots are busy now! Please, wait "+ShortDur(2*60))
                    Sleep(2*60*1000) // 2 minutes
                    slots = GetSlots().InUse
                }
            }
        }
        if system >= toSystem {
            curSystem = fromSystem-1
            system = curSystem
        }
    }
} else {Print("You don't have Recyclers on your Planets/Moons!")}
