/***** This script is created by RockClubKASHMIR *****\

--- WARNING!!! This script can work ONLY if you are Discoverer! ---

   DESCRIPTION
   This script find automatically your planet/moon with highgly amount of Pathfinders!
   Newer send more than 1 fleet to the detected debris field at same time, accept only in cases that first fleet is not enough to get all debris
  
  ONLY if the automatic method of finding your moon/planet not satisfied you;
  - replace all rows between //START and //END, with origin = GetCachedCelestial("M:1:2:3") where on "M:1:2:3" you must type your coordinate - M for the moon, P for planet
*/
fromSystem = 1 // Set from what system you want start to scan
toSystem = 499 // Set to what system you want to end to scan
Range = true // Do you want to use check/fly at range coordinates? true = YES / false = NO 

Telegram = false // Do you want to have TELEGRAM messages?  YES = true / NO = false
Pnbr = 5  // Will ignore debris less than for PATHFINDER with quantity as this value. The maximum is not limited even if you left this value as it is! Change it if/as you want.
times = 8 // if times = 5, the script will full scan 6 times the entire galaxy, from system, to system you set. You can set this value from 0, to the number you want
useCycles = false // Do you want to use the limited repeats?  YES = true / NO = false


//----
cycle = 0
origin = nil
flts = 0
//START
for celestial in GetCachedCelestials() {
    ships, _ = celestial.GetShips()
    if ships.Pathfinder > flts {
        flts = ships.Pathfinder
        origin = celestial 
    }
}
//END
nbr = 0
err = nil
if (Pnbr < 1) {Pnbr = 1}
if (times < 0) {times = 0}
if Range != false && Range != true {Range = true}
if useCycles != false && useCycles != true {useCycles = false}
totalSlots = GetSlots().Total - GetFleetSlotsReserved()
curSystem = fromSystem
if origin != nil {
    if IsDiscoverer() {
        Print("Your origin is "+origin.Coordinate)
        if toSystem > 499 || toSystem == 0 {toSystem = -1}
        if fromSystem > toSystem {Print("Please, type correctly fromSystem and/or toSystem!")}
        if Range == false {
            fromSystem = origin.GetCoordinate().System
            toSystem = origin.GetCoordinate().System
            curSystem = fromSystem
        }
        for system = curSystem; system <= toSystem; system++ {
            pp = 0
            dflag = 0
            abr = 0
            nbr = 0
            systemInfos, b = GalaxyInfos(origin.GetCoordinate().Galaxy, system)
            Dtarget, _ = ParseCoord(origin.GetCoordinate().Galaxy+":"+system+":"+16)
            Debris, _ = ParseCoord("D:"+origin.GetCoordinate().Galaxy+":"+system+":"+16)
            Sleep(Random(500, 1500)) // for avoid ban
            slots = GetSlots().InUse
            if err != nil {slots = totalSlots}
            if slots < totalSlots {
                if b == nil {
                    Print("Checking "+Dtarget)
                    if systemInfos.ExpeditionDebris.PathfindersNeeded >= Pnbr { 
                        ships, _ = origin.GetShips()
                        pp = systemInfos.ExpeditionDebris.PathfindersNeeded
                        if systemInfos.ExpeditionDebris.Metal == 0 && systemInfos.ExpeditionDebris.Crystal > 0 {Print("Found Crystal: "+systemInfos.ExpeditionDebris.Crystal)}
                        if systemInfos.ExpeditionDebris.Metal > 0 && systemInfos.ExpeditionDebris.Crystal == 0 {Print("Found Metal: "+systemInfos.ExpeditionDebris.Metal)}
                        if systemInfos.ExpeditionDebris.Metal > 0 && systemInfos.ExpeditionDebris.Crystal > 0 {Print("Found Metal: "+systemInfos.ExpeditionDebris.Metal+" and Crystal: "+systemInfos.ExpeditionDebris.Crystal)}
                        fleet, _ = GetFleets()
                        for f in fleet {
                            if f.Mission == RECYCLEDEBRISFIELD && f.ReturnFlight == false {
                                if Debris == f.Destination {
                                    if f.Ships.Pathfinder < pp {
                                        abr = pp - f.Ships.Pathfinder
                                    } else {dflag = 1}
                                }
                            }
                        }
                        if dflag == 0 {
                            f = NewFleet()
                            f.SetOrigin(origin)
                            f.SetDestination(Dtarget)
                            f.SetSpeed(HUNDRED_PERCENT)
                            f.SetMission(RECYCLEDEBRISFIELD)
                            if abr == 0 {
                                nbr = systemInfos.ExpeditionDebris.PathfindersNeeded
                            } else {nbr = abr}
                            if nbr > ships.Pathfinder {nbr = ships.Pathfinder}
                            f.AddShips(PATHFINDER, nbr)
                            a, err = f.SendNow()
                            if err == nil {
                                if nbr < systemInfos.ExpeditionDebris.PathfindersNeeded {Print("You don't have enough Ships for this debris field!")}
                                if nbr > 1 {
                                    Print(nbr+" Pathfinders are sended successfully!")
                                } else {Print(nbr+" Pathfinder is sended successfully!")}
                            } else {
                                if nbr > 1 {
                                    Print("The Pathfinders are NOT sended! "+err)
                                } else {
                                    Print("The Pathfinder is NOT sended! "+err)
                                }
                            }
                        } else {Print("Needed ships already are sended!")}
                    }
                }
            } else {
                for slots == totalSlots {
                    delay = Random(4*60, 8*60)
                    if err != nil {
                        Print("Please wait till ships lands! Recheck after "+ShortDur(delay))
                        Sleep(delay*1000)
                        ships, _ = origin.GetShips()
                        if ships.Pathfinder > 0 {slots = GetSlots().InUse}
                        err = nil
                    } else {
                        Print("All Fleet slots are busy now! Please, wait "+ShortDur(delay))
                        Sleep(delay*1000)
                        slots = GetSlots().InUse
                    }
                    curSystem = system-1
                }
            }
            if b == nil {
                if system >= toSystem {
                    if useCycles == true {
                        if times > 0 {
                            if cycle < times {
                                cycle++
                                if nbr == 0 {Print("Not found any debris! Start searching again...")}
                                curSystem = fromSystem-1
                                system = curSystem
                                delay = Random(50*60, 90*60)
                                Print("Will Start searching again after "+ShortDur(delay))
                                Sleep(delay*1000)
                            } else {
                                Print("You made "+(times+1)+" times full scan all systems chosen by you! The script turns off")
                                if Telegram == true {SendTelegram(TELEGRAM_CHAT_ID, "You made "+(times+1)+" times full scan all systems chosen by you! The script turns off")}
                                break
                            }
                        } else {
                            Print("You made full scan all systems chosen by you! The script turns off")
                            if Telegram == true {SendTelegram(TELEGRAM_CHAT_ID, "You made full scan all systems chosen by you! The script turns off")}
                            break
                        }
                    } else {
                        if nbr == 0 {Print("Not found any debris! Start searching again...")}
                        curSystem = fromSystem-1
                        system = curSystem
                        delay = Random(50*60, 90*60)
                        Print("Will Start searching again after "+ShortDur(delay))
                        Sleep(delay*1000)
                    }
                }
            } else {
                Print("Please, type correctly fromSystem and/or toSystem!")
                break
            }
        }
    } else {Print("You are not DISCOVERER! The sript will stop automatically")}
} else {Print("You don't have Pathfinders on your Planets/Moons!")}
