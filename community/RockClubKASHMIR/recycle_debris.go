//==== This script is created by RockClubKASHMIR ====

fromSystem = 1 // Your can change this as you wish
toSystem = 499 // Your can change this as you wish
Rnbr = 1 // if Rnbr = 1, the script will search debris for minimum 2 Recyclers. You can change it to any value you want
//----
curSystem = fromSystem
origin = nil
flts = 0
nbr = 0
err = 0
i = 1
slots = GetSlots().InUse+GetFleetSlotsReserved()
for celestial in GetCachedCelestials() {
    ships, _ = celestial.GetShips()
    if ships.Recycler > flts {
        flts = ships.Recycler
        origin = celestial // Your Planet(or Moon), with biggest amount of Recyclers
    }
}
if origin != nil {
    Print("Your origin is "+origin.Coordinate)
    for system = fromSystem; system <= toSystem; system++ {
        systemInfos, err = GalaxyInfos(origin.GetCoordinate().Galaxy, system)
        planetInfo = systemInfos.Position(i)
        if slots < GetSlots().Total {
            if planetInfo != nil {
                Print("Checking "+planetInfo.Coordinate)
                if planetInfo.Debris.RecyclersNeeded > Rnbr { 
                    ships, _ = origin.GetShips()
                    Print("Found Metal:"+planetInfo.Debris.Metal+" and Crystal:"+planetInfo.Debris.Crystal+" at "+planetInfo.Coordinate+", need "+planetInfo.Debris.RecyclersNeeded+" Recyclers!")
                    Sleep(Random(6*1000, 12*1000))
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
                        Print(nbr+" Pathfinders are sended successfully!")
                        SendTelegram(TELEGRAM_CHAT_ID, nbr+" Pathfinders are sended successfully!")
                    } else {
                        Print("The fleet is NOT sended! "+err)
                        SendTelegram(TELEGRAM_CHAT_ID, "The fleet is NOT sended! "+err)
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
                    Sleep(2*60*1000)
                    ships, _ = origin.GetShips()
                    if ships.Recycler > 0 {slots = GetSlots().InUse+GetFleetSlotsReserved()}
                } else {
                    Print("All Fleet slots are busy now! Please, wait "+ShortDur(2*60))
                    Sleep(2*60*1000)
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
