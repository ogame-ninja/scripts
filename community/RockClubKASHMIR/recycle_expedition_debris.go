/***** This script is created by RockClubKASHMIR <discord @RockClubKASHMIR#8058> *****\
discord channel for your personal orders and support - https://discord.gg/zgTuT3

 v1.1
 
    DESCRIPTION
 1. Always respect reserved slots
 2. Checks for EXPEDITION debris and recycle them (if you are Discoverer and have Pathfinders)
 3. Can check and recycle EXPEDITION debris from more than 1 planet/moon 
 4. Possible to make a scan and recycle debris at range solar systems
 5. You can set the minimum amount of pathfinders for recycle
 6. Sends Pathfinders to same debris only if already sended ships are not enough to get all resources
 7. Possibility to repeat scanning many times - you can set how many
 8. Repeats scanning after a minutes that you are set 
 */

homes = ["M:1:2:3"] // Replace M:1:2:3 with your coordinate - M for the moon, P for planet.
// You can add as many planets/moons you want - the home list must look like this: homes = ["M:1:2:3", "M:2:2:3"]

SystemsRange = true // Do you want to check for debris in range solar systems? true = YES / false = NO
RangeRadius = 30  // Radius around your solar system.Set this if SystemsRange = true
Pnbr = 5  // The script will ignore debris less than for PATHFINDERS that you set - The Maximum PATHFINDERS is limited only of your PATHFINDERS on the current moon/planet! You can set this value from 1, to the number you want

Repeat = true // Do you want to repeat the full scanning for debris? true = YES / false = NO
HowManyRepeats = 5 // Set the limit of repeats of full scanning for EXPO debris - 0 means forewer
PauseBetweenRepeats = 30 // Set the pause between repeats in minutes

//----- Please, don't change the code below -----\\
current = 0
wrong = []
homeworld = nil
PauseFarmingBot()
i = 0
ei = 0
er = nil
err = nil
nbr = 0
endFlag = 0
fleetFlag = 0
RepeatTimes = 1
if Pnbr < 1 {Pnbr = 1}
for home in homes {
    flag = 1
    hh, _ = ParseCoord(home)
    for celestial in GetCachedCelestials() {
        if celestial.Coordinate == hh {
            ei++
            flag = 0
        }
    }
    if flag == 1 {wrong += home}
    i++
}
if ei == len(homes) {homeworld = GetCachedCelestial(homes[0])}
PathfindersDebris = true
if !IsDiscoverer() {
    Print("You are not Discoverer and cannot get the EXPO Debris!")
    PathfindersDebris = false
}
if HowManyRepeats == 0 {HowManyRepeats = false}
if homeworld != nil {
    ls = GetSlots()
    Sleep(2000)
    totalUsl = ls.Total - GetFleetSlotsReserved()
    for home = current; home <= len(homes)-1; home++ {
        pp = 0
        Dtarget = 0
        marker = home
        delay = 0
        ls = GetSlots()
        homeworld = GetCachedCelestial(homes[home])
        if homeworld.Coordinate.IsMoon() {
            Print("Your Moon is: "+homeworld.Coordinate)
        } else {Print("Your Planet is: "+homeworld.Coordinate)}
        fromSystem = homeworld.GetCoordinate().System - RangeRadius
        toSystem = homeworld.GetCoordinate().System + RangeRadius
        if fromSystem < 1 {fromSystem = 1}
        if toSystem > 499 {toSystem = 499}
        totalSlots = totalUsl
        slots = ls.InUse
        if slots < totalSlots {
            
            if PathfindersDebris == true {
                dflag = 0
                abr = 0
                curSystem = fromSystem
                if SystemsRange == false {
                    curSystem = homeworld.GetCoordinate().System
                    toSystem = homeworld.GetCoordinate().System
                }
                for system = curSystem; system <= toSystem; system++ {
                    Sleep(Random(1000, 3000))
                    systemInfos, _ = GalaxyInfos(homeworld.GetCoordinate().Galaxy, system)
                    Dtarget, _ = ParseCoord(homeworld.GetCoordinate().Galaxy+":"+system+":"+16)
                    Debris, _ = ParseCoord("D:"+homeworld.GetCoordinate().Galaxy+":"+system+":"+16)
                    Print("Checking "+Dtarget)
                    if systemInfos.ExpeditionDebris.PathfindersNeeded >= Pnbr {
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
                                    } else {
                                        dflag = 1
                                        nbr = 1
                                    }
                                }
                            }
                        }
                        ls = GetSlots()
                        Sleep(Random(1000, 3000))
                        slots = ls.InUse
                        if slots < totalSlots {
                            if dflag == 0 {
                                myShips, _ = homeworld.GetShips()
                                f = NewFleet()
                                f.SetOrigin(homeworld)
                                f.SetDestination(Dtarget)
                                f.SetSpeed(HUNDRED_PERCENT)
                                f.SetMission(RECYCLEDEBRISFIELD)
                                if abr == 0 {
                                    nbr = systemInfos.ExpeditionDebris.PathfindersNeeded
                                } else {nbr = abr}
                                if nbr > myShips.Pathfinder {nbr = myShips.Pathfinder}
                                f.AddShips(PATHFINDER, nbr)
                                a, err = f.SendNow()
                                if err == nil {
                                    slots = slots + 1
                                    if nbr < systemInfos.ExpeditionDebris.PathfindersNeeded {Print("You don't have enough Ships for this debris field!")}
                                    if nbr > 1 {
                                        Print(nbr+" Pathfinders are sended successfully!")
                                    } else {Print(nbr+" Pathfinder is sended successfully!")}
                                } else {
                                    if nbr > 1 {
                                        Print("The Pathfinders are NOT sended! "+err)
                                    } else {Print("The Pathfinder is NOT sended! "+err)}
                                    system = toSystem
                                    er = err
                                    err = nil
                                }
                            } else {Print("Needed ships already are sended!")}
                        }
                        if slots == totalSlots {
                            fleetFlag = 1
                            if system < toSystem {curSystem = system-1}
                            system = toSystem
                            current = marker-1
                            home = len(homes)-1
                        }
                    }
                    if marker >= len(homes)-1 {err = er}
                }
                if pp == 0 {Print("Not found any debris!")}
            }
        } else {fleetFlag = 1}
        if err != nil {slots = totalSlots}
        if home >= len(homes)-1 {
            for slots == totalSlots {
                delay = Random(7*60, 13*60) // 7 - 13 minutes in seconds
                if Repeat == true {
                    if err != nil {
                        slots = GetSlots().InUse
                        expslots = slots
                        if slots > 0 {
                            Print("Please, wait until Pathfinders returns! Re-check after "+ShortDur(delay))
                            Sleep(delay*1000)
                            expslots = GetSlots().InUse
                            if slots == expslots {slots = totalSlots}
                        }
                    } else {
                        if fleetFlag == 1 {Print("All slots are busy now! Please, wait "+ShortDur(delay))}
                        Sleep(delay*1000)
                        slots = GetSlots().InUse
                    }
                } else {
                    slots = 1
                    totalSlots = 3
                }
            }
            if RepeatTimes != HowManyRepeats {
                if err != nil || fleetFlag == 1 {
                    delay = 3
                    fleetFlag = 0
                } else {delay = Random((PauseBetweenRepeats-5)*60, PauseBetweenRepeats*60)}
                if marker >= len(homes)-1 {
                    if nbr == 0 {
                        Print("Not found any EXPEDITION debris!")
                        Sleep(3000)
                    }
                    if Repeat == true {
                        if RepeatTimes == 1 {Print("You have full scan for debris all coordinates "+RepeatTimes+" time")}
                        if RepeatTimes > 1 {Print("You have full scan for debris all coordinates "+RepeatTimes+" times")}
                    }
                    RepeatTimes++
                    Sleep(2000)
                    Print("Start searching for debris again after "+ShortDur(delay))
                    Sleep(delay*1000)
                    current = -1
                    err = nil
                    er = nil
                    nbr = 0
                }
                if Repeat == true {home = current}
            } else {
                Print("You have reached the limit of repeats that you have set")
                Sleep(3000)
            }
        }
        Sleep(Random(1000, 3000))
    }
} else {Print("You typed wrong coordinates! - "+wrong)}
