//---This script is created by RockClubKASHMIR---

fromSystem = 1 // Your can change the value as you wish!
toSystem = 499 // Your can change the value as you wish!
curSystem = fromSystem
origin = nil
flts = 0
nbr = 0
err = 0
slots = GetSlots().InUse
//----
for celestial in GetCachedCelestials() {
    ships, _ = celestial.GetShips()
    if ships.Pathfinder > 0 {
        if ships.Pathfinder > flts {
            flts = ships.Pathfinder
            origin = celestial
        }
    }
}
if origin != nil {
    Print("Your origin is "+origin.Coordinate)
    for system = curSystem; system <= toSystem; system++ {
        systemInfo, _ = GalaxyInfos(origin.GetCoordinate().Galaxy, system)
        Dtarget, _ = ParseCoord(origin.GetCoordinate().Galaxy+":"+system+":"+16)
        if slots < GetSlots().Total {
            ships, _ = origin.GetShips()
            if Dtarget != 0 {
                Print("Checking "+Dtarget)
                if systemInfo.ExpeditionDebris.PathfindersNeeded > 1 { 
                    Print("Found Metal:"+systemInfo.ExpeditionDebris.Metal+" and Crystal:"+systemInfo.ExpeditionDebris.Crystal+" at "+Dtarget+", need "+systemInfo.ExpeditionDebris.PathfindersNeeded+" Pathfinders!")
                    Sleep(6000)
                    f = NewFleet()
                    f.SetOrigin(origin)
                    f.SetDestination(Dtarget)
                    f.SetSpeed(HUNDRED_PERCENT)
                    f.SetMission(RECYCLEDEBRISFIELD)
                    if systemInfo.ExpeditionDebris.PathfindersNeeded > ships.Pathfinder {
                        nbr = ships.Pathfinder
                    } else {nbr = systemInfo.ExpeditionDebris.PathfindersNeeded}
                    f.AddShips(PATHFINDER, nbr)
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
        } else {
            for slots == GetSlots().Total {
                if err != 0 {
                    Print("Please wait till ships lands! Recheck after "+ShortDur(120))
                    Sleep(120000)
                    ships, _ = origin.GetShips()
                    if ships.Pathfinder > 0 {slots = GetSlots().InUse}
                } else {
                    Print("All Fleet slots are busy now! Please, wait "+ShortDur(120))
                    Sleep(120000)
                    slots = GetSlots().InUse
                }
            }
        }
        if system >= toSystem {
            curSystem = fromSystem-1
            system = curSystem
        }
    }
} else {Print("You don't have Pathfinders on your Planets/Moons!")}
