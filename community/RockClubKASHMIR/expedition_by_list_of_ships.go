/***** This script is created by RockClubKASHMIR <discord @RockClubKASHMIR#8058> *****\
discord channel for your personal orders and support - https://discord.gg/zgTuT3

v5.6
 
DESCRIPTION
 1. This script always respects the reserved slots
 2. Possibility to send EXPEDITION fleets from more than 1 planet/moon
 3. You can set up your ship list by 2 methods (or by combination of both of them):
    a. Automatic: All ships with quantity 0 that you set will be calculated automatically (full quantity divided by the free EXPO slots)
       - if sendAtOnce = true, all ships set with quantity 0 will be sent at once.
    b. Set quantity of all ships by yourself:
       - the ships set up with this method will be accepted literally, and if any of your ships is even 1 less, the fleet will not be sent
 4. Possibility to send your EXPEDITION fleets at a range of your solar system or to your solar system only
 5. Evenly distribution of EXPEDITION slots per each moon/planet or use all EXPEDITION slots per every planet/moon
 6. Check for EXPEDITION Debris and recycle them (if you are Discoverer and have Pathfinders)
 7. Possibility to make a scan and recycle debris at a range of your solar system or to your solar system only
 8. You can set a minimum amount of pathfinders for recycling
 9. Sends Pathfinders to same debris more than once only if already sended ships are not enough to get all resources
10. Possibility to repeat the sending of EXPEDITION fleets many times - you can set how many
11. You can start this script at a specific time. Sending of the fleets will stop after the number of repeats that you set
*/

homes = ["M:1:2:3"] // Replace M:1:2:3 with your coordinate - M for the moon, P for planet.
// You can add as many planets/moons you want - the home list must look like this: homes = ["M:1:2:3", "M:2:2:3"]

shipsList = {LARGECARGO: 0, LIGHTFIGHTER: 0, PATHFINDER: 1}// Set your Ships list

SystemsRange = false // Do you want to send your EXPO fleet to Range coordinates? true = YES / false = NO
sendWhenFleetBack = false // Do you want every time to wait until all EXPEDITION fleets back before send them all again, for each planet/moon? true = YES / false = NO
sendAtOnce = false //Do you want to send the ships set with quantity 0 at once? true = YES / false = NO

DurationOfExpedition = 1 // Set duration (in hours) of the EXPEDITION: minimum 1 - maximum 8
RangeRadius = 5  // Set this if SystemsRange = true or/and PathfinderSystemsRange = true 

PathfindersDebris = true // Do you want to get EXPO debrises? true = YES / false = NO
Pnbr = 5  // The script will ignore debris less than for PATHFINDERS that you set - The Maximum PATHFINDERS is limited only of your PATHFINDERS on the current moon/planet! You can set this value from 1, to the number you want
PathfinderSystemsRange = true // Do you want to check/get EXPO debris in range systems? true = YES / false = NO

Repeat = true // Do you want to repeat the full cycle of fleet sending? true = YES / false = NO
HowManyCycles = 5 // Set the limit of repeats of whole cycle of EXPO fleet sending - 0 means forewer

myTime = "09:33:00"// Set your start Time; Hour: 00 - 23, Minute: 00 - 59
useStartTime = false // Do you want to run this script at specific time every day? true = YES / false = NO

//----- Please, don't change the code below -----\\
current = 0
wrong = []
curentco = {}
waves = {}
homeworld = nil
PauseFarmingBot()
StopHunter()
i = 0
ei = 0
er = nil
err = nil
flag = 0
cng = 0
cycle = 0
endFlag = 0
fleetFlag = 0
RepeatTimes = 1
calc = 0
if (Pnbr < 1) {Pnbr = 1}
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
if len(shipsList) > 0 {
    for ShipID, num in shipsList {
        if num == 0 {calc = 1}
    }
} else {
    Print("Your Ship's list is emty!")
    StopScript(__FILE__)
}
if !IsDiscoverer() {
    Print("You are not Discoverer and cannot get the EXPO Debris!")
    PathfindersDebris = false
}
if useStartTime == false {
    hour, minute, sec = Clock()
    startHour = hour
    startMin = minute
    startSec = sec + 3
    if startSec >= 60 {
        startSec = startSec - 60
        startMin = startMin + 1
        if startMin >= 60 {
            startMin = startMin - 60
            startHour = startHour + 1
        }
        if startHour >= 24 {startHour = startHour - 24}
    }
    myTime = ""+startSec+" "+startMin+" "+startHour+" * * *"
}
if HowManyCycles == 0 {HowManyCycles = false}
if homeworld != nil {
    CronExec(myTime, func() {
        slotMarker = 0
        ls = GetSlots()
        Sleep(2000)
        totalUsl = ls.Total - GetFleetSlotsReserved()
        totalExpSlots = ls.ExpTotal
        for home = current; home <= len(homes)-1; home++ {
            pp = 0
            Dtarget = 0
            marker = home
            ls = GetSlots()
            homeworld = GetCachedCelestial(homes[home])
            if homeworld.Coordinate.IsMoon() {
                Print("Your Moon is: "+homeworld.Coordinate)
            } else {Print("Your Planet is: "+homeworld.Coordinate)}
            fromSystem = homeworld.GetCoordinate().System - RangeRadius
            toSystem = homeworld.GetCoordinate().System + RangeRadius
            if fromSystem < 1 {fromSystem = 1}
            if toSystem > 499 {toSystem = 499}
            crdn = fromSystem
            ExpsTemp = 0
            if SystemsRange == true && cycle >= len(homes)-1 {
                for id, num in curentco {
                    if id == homes[home] {crdn = num}
                }
            }
            totalSlots = totalUsl
            slots = ls.InUse
            bk = 0
            currentTime = bk
            times = totalExpSlots
            if slots < totalSlots {
                slots = ls.ExpInUse
                totalSlots = totalExpSlots
                ExpsTemp = 1
                if slots == totalSlots {fleetFlag = 2}
            } else {fleetFlag = 1}
            if err != nil {slots = totalSlots}
            if slots < totalSlots {
                Expos = totalExpSlots - slots
                if sendWhenFleetBack == false {
                    slotMarker = totalExpSlots-marker
                    times = slotMarker/len(homes)
                    if times > Floor(times) {times = Floor(times) + 1}
                    if times < 1 {times = 1}
                }
                Flts, _ = GetFleets()
                bk = 0
                for f in Flts {
                    if f.Mission == EXPEDITION {
                        hh, _ = ParseCoord(homes[home])
                        if hh == f.Origin {bk = bk + 1}
                    }
                }
                currentTime = bk
                if sendAtOnce == true {
                    times = 1
                    Expos = 1
                }
                Expos = times - bk
                if Expos <= 0 {
                    currentTime = times
                    Print("There are no EXPO fleets to send here!")
                } else {Print(Expos+" slots will be used")}
                for time = currentTime; time < times; time++ {
                    myShips, _ = homeworld.GetShips()
                    tt = 0
                    rtt = 0
                    ExpFleet = {}
                    if ExpsTemp == 0 {
                        totalSlots = totalUsl
                        ls = GetSlots()
                        slots = ls.InUse
                        Sleep(800)
                        if slots < totalSlots {
                            slots = ls.ExpInUse
                            Sleep(800)
                            totalSlots = totalExpSlots
                            if slots == totalSlots {fleetFlag = 2}
                        } else {fleetFlag = 1}
                    }
                    if err != nil {slots = totalSlots}
                    if slots < totalSlots {
                        ExpsTemp == 0
                        if SystemsRange == false {
                            Dtarget, _ = ParseCoord(homeworld.GetCoordinate().Galaxy+":"+homeworld.GetCoordinate().System+":"+16)
                        }
                        if SystemsRange == true {
                            if crdn > toSystem {crdn = fromSystem}
                            Dtarget, _ = ParseCoord(homeworld.GetCoordinate().Galaxy+":"+crdn+":"+16)
                        }
                        explist = []
                        Sleep(Random(13000, 18000)) // For avoiding ban
                        Flts, _ = GetFleets()
                        fleet = NewFleet()
                        fleet.SetOrigin(homeworld)
                        fleet.SetDestination(Dtarget)
                        fleet.SetSpeed(HUNDRED_PERCENT)
                        fleet.SetMission(EXPEDITION)
                        if len(shipsList) > 0 {
                            for ShipID, num in shipsList {
                                rtt = rtt + 1
                                fleetInAir = 0
                                if myShips.ByID(ShipID) != 0 {
                                    if num == 0 {
                                        if sendAtOnce == false {
                                            for f in Flts {
                                                ships = f.Ships
                                                if f.Mission == EXPEDITION {
                                                    if homeworld.Coordinate == f.Origin {
                                                        if ships.ByID(ShipID) != 0 {
                                                            fleetInAir = fleetInAir + ships.ByID(ShipID)
                                                        }
                                                    }
                                                }
                                            }
                                            fleetInAir = fleetInAir + myShips.ByID(ShipID)
                                            num = Floor(fleetInAir/times)
                                            temp = (num/100)*40
                                            if myShips.ByID(ShipID) < num && myShips.ByID(ShipID) >= temp {num = myShips.ByID(ShipID)}
                                            if myShips.ByID(ShipID) < num && myShips.ByID(ShipID) < temp {num = 0}
                                        }
                                        if sendAtOnce == true {num = myShips.ByID(ShipID)}
                                        if num < 1 {num = 0}
                                        if num > 0 {
                                            ExpFleet[ShipID] = num
                                            tt = tt + 1
                                        }
                                    } else {
                                        if ShipID != PATHFINDER {
                                            if myShips.ByID(ShipID) >= num {
                                                ExpFleet[ShipID] = num
                                                tt = tt + 1
                                            }
                                        }
                                        if ShipID == PATHFINDER {
                                            if myShips.ByID(ShipID) >= num {
                                                ExpFleet[ShipID] = num
                                                tt = tt + 1
                                            }
                                            if len(shipsList) > 1 && myShips.ByID(ShipID) < num {
                                                num = myShips.ByID(ShipID)
                                                ExpFleet[ShipID] = num
                                                tt = tt + 1
                                            }
                                        }
                                    }
                                }
                            }
                        }
                        fleet.SetDuration(DurationOfExpedition)
                        if rtt == tt {
                            for ShipID, nbr in ExpFleet {
                                fleet.AddShips(ShipID, nbr)
                                explist += ShipID+": "+nbr
                            }
                        }
                        a, err = fleet.SendNow()
                        if err == nil {
                            cng = 1
                            slots = slots + 1
                            waves[homes[home]] = 1
                            Print(explist+" are sended successfully to "+Dtarget)
                            if SystemsRange == true {
                                if crdn <= toSystem {crdn++}
                                curentco[homes[home]] = crdn
                            }
                        } else {
                            time = times
                            Print("The fleet is NOT sended! "+err)
                            er = err
                            err = nil
                        }
                        if marker >= len(homes)-1 {err = er}
                    }
                    if slots == totalSlots && err == nil {
                        time = times
                        fleetFlag = 2
                    }
                    if err != nil {slots = totalSlots}
                    Sleep(Random(1500, 3000))
                }
            } else {home = len(homes)-1}
            if PathfindersDebris == true {
                dflag = 0
                abr = 0
                nbr = 0
                curSystem = fromSystem
                if PathfinderSystemsRange == false {
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
                                    } else {dflag = 1}
                                }
                            }
                        }
                        ls = GetSlots()
                        Sleep(Random(1000, 3000))
                        aaz = ls.InUse
                        if aaz < totalUsl {
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
                                a, b = f.SendNow()
                                if b == nil {
                                    aaz = aaz + 1
                                    if aaz == totalUsl {
                                        fleetFlag = 1
                                        system = toSystem
                                    }
                                    if nbr < systemInfos.ExpeditionDebris.PathfindersNeeded {Print("You don't have enough Ships for this debris field!")}
                                    if nbr > 1 {
                                        Print(nbr+" Pathfinders are sended successfully!")
                                    } else {Print(nbr+" Pathfinder is sended successfully!")}
                                } else {
                                    if nbr > 1 {
                                        Print("The Pathfinders are NOT sended! "+b)
                                    } else {Print("The Pathfinder is NOT sended! "+b)}
                                    break
                                }
                            } else {Print("Needed ships already are sended!")}
                        } else {fleetFlag = 1}
                    }
                }
                if pp == 0 {Print("Not found any debris!")}
            }
            if cycle <= len(homes)-1 {cycle++}
            ls = GetSlots()
            if ls.InUse == totalUsl || ls.ExpInUse == totalExpSlots {
                if sendWhenFleetBack == true && sendAtOnce == false {
                    if ls.ExpInUse < totalExpSlots {fleetFlag = 0}
                }
                home = len(homes)-1
                slots = totalSlots
                current = marker
            }
            if sendAtOnce == true {
                if marker >= len(homes)-1 {
                    if ls.ExpInUse < totalExpSlots && err == nil {
                        err = "no ships to send"
                        slots = totalSlots
                    }
                }
            }
            if home >= len(homes)-1 {
                for slots == totalSlots {
                    delay = Random(7*60, 13*60) // 7 - 13 minutes in seconds
                    if Repeat == true {
                        if err != nil {
                            slots = GetSlots().ExpInUse
                            expslots = slots
                            if slots > 0 {
                                Print("Please wait for the landing of your EXPO ships! Recheck after "+ShortDur(delay))
                                Sleep(delay*1000)
                                expslots = GetSlots().ExpInUse
                                if slots > expslots {
                                    err = nil
                                    er = nil
                                } else {slots = totalSlots}
                            } else {
                                if cng == 0 {
                                    Print("All your EXPO ships are on the ground! Please, check your deuterium and make sure that you set the ships list correctly, then start the script again!")
                                    RepeatTimes = HowManyCycles
                                    useStartTime = false
                                    endFlag = 1
                                }
                            }
                        } else {
                            if fleetFlag == 0 {Print("Please, wait till all your EXPO fleets arrives! Re-check after "+ShortDur(delay))}
                            if fleetFlag == 1 {Print("All slots are busy now! Please, wait "+ShortDur(delay))}
                            if fleetFlag == 2 {Print("All EXPO slots are busy! Please, wait "+ShortDur(delay))}
                            Sleep(delay*1000)
                            ls = GetSlots()
                            if fleetFlag == 1 {slots = ls.InUse}
                            if  fleetFlag == 0 || fleetFlag == 2 {
                                slots = ls.ExpInUse
                                if sendWhenFleetBack == true && slots >= 1 {
                                    if slots < totalSlots {
                                        fleetFlag = 0
                                        slots = totalSlots
                                    }
                                }
                            }
                        }
                    } else {
                        slots = 1
                        totalSlots = 3
                    }
                }
                if RepeatTimes != HowManyCycles {
                    if marker >= len(homes)-1 {
                        if len(waves) == len(homes) {
                            if HowManyCycles != false {
                                if Repeat == true {Print("You make full cycle of fleet sending "+RepeatTimes+"!")}
                                RepeatTimes++
                                waves = {}
                            }
                        }
                        current = -1
                        cng = 0
                        err = nil
                        er = nil
                    }
                    if Repeat == true {home = current}
                } else {
                    if endFlag == 0 {Print("You have reached the limit of repeats that you have set")}
                    Sleep(3000)
                }
            }
            Sleep(Random(1000, 3000))
        }
        if useStartTime == false {StopScript(__FILE__)}
    })
} else {
    Print("You typed wrong coordinates! - "+wrong)
    StopScript(__FILE__)
}
<-OnQuitCh
