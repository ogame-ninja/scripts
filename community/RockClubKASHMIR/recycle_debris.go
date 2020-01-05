//==== This script is created by RockClubKASHMIR ====

/* This script automatically find your Planet or Moon with highes amount of Recyclers
 * No need to setup your planet!
*/
fromSystem = 1 // Your can change this value as you wish
toSystem = 200 // Your can change this value as you wish
Rnbr = 0  // If Rnbr = 1, the script will search only debris for minimum 2 Recyclers. You can change this value as you wish
times = 1 // if times = 1, the script will full scan 2 times the galaxy, from system, to system you want. You can set this value from 0, to the number you want
//----
cycle = 0
curSystem = fromSystem
origin = nil
flts = 0
nbr = 0
err = nil
i = 1
if (Rnbr < 0) {Rnbr = 0}
if (times < 0) {times = 0}
totalSlots = GetSlots().Total - GetFleetSlotsReserved()
// Start to Search highest amount of Recyclers on all your Planets and Moons(if you have some)
for celestial in GetCachedCelestials() {
    ships, _ = celestial.GetShips()
    if ships.Recycler > flts {
        flts = ships.Recycler
        origin = celestial // Your Planet(or Moon) with highest amount of Recyclers
    }
}
if origin != nil {
    Print("Your origin is "+origin.Coordinate)
    if toSystem > 499 || toSystem == 0 {toSystem = -1}
    if fromSystem > toSystem {Print("Please, type correctly fromSystem and/or toSystem!")}
    for system = curSystem; system <= toSystem; system++ {
        systemInfos, b = GalaxyInfos(origin.GetCoordinate().Galaxy, system)
        planetInfo = systemInfos.Position(i)
        Sleep(Random(500, 1000)) // for avoid ban
        slots = GetSlots().InUse
        if err != nil {slots = totalSlots}
        if slots < totalSlots {
            if b == nil {
                if planetInfo != nil {
                    Print("Checking "+planetInfo.Coordinate)
                    if planetInfo.Debris.RecyclersNeeded > Rnbr { 
                        ships, _ = origin.GetShips()
                        if planetInfo.Debris.Metal == 0 && planetInfo.Debris.Crystal > 0 {Print("Found Crystal: "+planetInfo.Debris.Crystal)}
                        if planetInfo.Debris.Metal > 0 && planetInfo.Debris.Crystal == 0 {Print("Found Metal: "+planetInfo.Debris.Metal)}
                        if planetInfo.Debris.Metal > 0 && planetInfo.Debris.Crystal > 0 {Print("Found Metal: "+planetInfo.Debris.Metal+" and Crystal: "+planetInfo.Debris.Crystal)}
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
                            if planetInfo.Debris.RecyclersNeeded > ships.Recycler {Print("You don't have enough Recyclers for these Debris field!")}
                            if nbr > 1 {
                                Print(nbr+" Recyclers are sended successfully!")
                            } else {Print(nbr+" Recycler is sended successfully!")}
                        } else {
                            if nbr > 1 {
                                Print("The Recyclers are NOT sended! "+err)
                                SendTelegram(TELEGRAM_CHAT_ID, "The Recyclers are NOT sended! "+err)
                            } else {
                                Print("The Recycler is NOT sended! "+err)
                                SendTelegram(TELEGRAM_CHAT_ID, "The Recycler is NOT sended! "+err)
                            }
                        }
                    }
                }
                if i < 15 {
                    i++
                } else {i = 1}
            }
        } else {
            for slots == totalSlots {
                if err != 0 {
                    Print("Please wait till ships lands! Recheck after "+ShortDur(2*60))
                    Sleep(2*60*1000)
                    ships, _ = origin.GetShips()
                    if ships.Recycler > 0 {slots = GetSlots().InUse}
                    err = nil
                } else {
                    Print("All Fleet slots are busy now! Please, wait "+ShortDur(2*60))
                    Sleep(2*60*1000)
                    slots = GetSlots().InUse
                }
                curSystem = system-1
            }
        }
        if b == nil {
            if system >= toSystem {
                if times > 0 {
                    if cycle < times {
                        cycle++
                        if nbr == 0 {Print("Not found any debris! Start searching again...")}
                        curSystem = fromSystem-1
                        system = curSystem
                        Sleep(4000)
                    } else {
                        Print("You made "+(times+1)+" times full scan all systems chosen by you! The script turns off")
                        SendTelegram(TELEGRAM_CHAT_ID, "You made "+(times+1)+" times full scan all systems chosen by you! The script turns off")
                        break
                    }
                } else {
                    Print("You made full scan all systems chosen by you! The script turns off")
                    SendTelegram(TELEGRAM_CHAT_ID, "You made full scan all systems chosen by you! The script turns off")
                    break
                }
            }
        } else {
            Print("Please, type correctly fromSystem and/or toSystem!")
            break
        }
    }
} else {Print("You don't have Recyclers on your Planets/Moons!")}
