/* this script transport and build defences on every planet/moon
for any suggestion, help or bug report contact me on slack: RudeDude
set those variable as you wish */
homeworld = GetCachedCelestial("M:1:2:3") //this will be the main colony/moon to send transports
standardrocket = 21000 //rocketlaucher you want the script to build
standardllaser = 4000 //light laser you want the script to build
standardhlaser = 2000 //heavy laser you want the script to build
standardgauss = 400 //gauss cannon you want the script to build
standardplasma = 200 //plasma turret you want the script to build
reservedslots = 2 // slot that you want to reserve
minmet = 0 // metal to preserve
mincrys = 0 // crystal to preserve
mindeut = 10000000 // deuterium to preserve
minNanites = 6 // required nanites level to start building the defense, if you want to build defences on moon should be set to zero
minShipyard = 8 //required shipyard level to start building the defense, make sure it matches the requirements for the def you willing to build
whereshouldbuild =  GetCachedPlanets() //planets: GetCachedPlanets(), moons: GetCachedMoons(), planets+moons: GetCachedCelestials()

// the script starts here, you should not edit the code below
func buildefense(celestial, sumrocket, sumllaser, sumhlaser, sumgauss, sumplasma, one, goalrocket, two, goalllaser, three, goalhlaser, four, goalgauss, five, goalplasma) {
    tostay = NewResources(minmet, mincrys, mindeut)
    BID, BCD, _, RCD = celestial.ConstructionsBeingBuilt()
    missingrocket = Max(0, goalrocket - sumrocket)
    missingllaser = Max(0, goalllaser - sumllaser)
    missinghlaser = Max(0, goalhlaser - sumhlaser)
    missingauss = Max(0, goalgauss - sumgauss)
    missingplasma = Max(0, goalplasma - sumplasma)
    rocketPrice = GetPrice(ROCKETLAUNCHER, 1)
    llaserPrice = GetPrice(LIGHTLASER, 1)
    hlaserPrice = GetPrice(HEAVYLASER, 1)
    gaussPrice = GetPrice(GAUSSCANNON, 1)
    plasmaPrice = GetPrice(PLASMATURRET, 1)
    sumres, _ = homeworld.GetResources()
    res = sumres.Sub(tostay)
    price1 = rocketPrice.Mul(missingrocket)
    price2 = llaserPrice.Mul(missingllaser)
    price3 = hlaserPrice.Mul(missinghlaser)
    price4 = gaussPrice.Mul(missingauss)
    price5 = plasmaPrice.Mul(missingplasma)
    restosend = 0
    goalrocket = 0
    goalllaser = 0
    goalhlaser = 0
    goalgauss = 0
    goalplasma = 0
    if res.Gte(price1.Add(price2).Add(price3).Add(price4).Add(price5)) {
        goalrocket = missingrocket
        goalllaser = missingllaser
        goalhlaser = missinghlaser
        goalgauss = missingauss 
        goalplasma = missingplasma 
        restosend = price1.Add(price2).Add(price3).Add(price4).Add(price5)
    } else if res.Gte(price1.Add(price2).Add(price3).Add(price4)) && (missingrocket > 0 || missingllaser > 0 || missinghlaser > 0 || missingauss > 0) {
        goalrocket = missingrocket
        goalllaser = missingllaser
        goalhlaser = missinghlaser
        goalgauss = missingauss 
        restosend = price1.Add(price2).Add(price3).Add(price4)
    } else if res.Sub(tostay).Gte(price1.Add(price2).Add(price3)) && (missingrocket > 0 || missingllaser > 0 || missinghlaser > 0) {
        goalrocket = missingrocket
        goalllaser = missingllaser
        goalhlaser = missinghlaser
        restosend = price1.Add(price2).Add(price3)
    } else if res.Gte(price1.Add(price2)) && (missingrocket > 0 || missingllaser > 0) {
        goalrocket = missingrocket
        goalllaser = missingllaser
        restosend = price1.Add(price2)
    } else if res.Gte(price1) && missingrocket > 0 {
        goalrocket = missingrocket
        restosend = price1
    } else if missingplasma > 0 {
        goalplasma = res.Div(plasmaPrice)
        restosend = plasmaPrice.Mul(goalplasma)
    } else if missingrocket > 0 {
        goalgauss = res.Div(gaussPrice)
        restosend = gaussPrice.Mul(goalgauss)
    } else if missinghlaser > 0 {
        goalhlaser = res.Div(hlaserPrice)
        restosend = hlaserPrice.Mul(goalhlaser)
    } else if missingllaser > 0 {
        goalllaser = res.Div(llaserPrice)
        restosend = llaserPrice.Mul(goalllaser)
    } else if missingrocket > 0 {
        goalrocket = res.Div(rocketPrice)
        restosend = rocketPrice.Mul(goalrocket)
    }
    Print(celestial.GetCoordinate(), "resources required", restosend)
    fleet = NewFleet()
    fleet.SetOrigin(homeworld)
    fleet.SetDestination(celestial)
    fleet.SetSpeed(HUNDRED_PERCENT)
    fleet.SetMission(TRANSPORT)
    fleet.SetResources(restosend)
    lc, sc = CalcCargo(restosend.Total())
    fleet.AddShips(LARGECARGO, lc)
    fleet, err = fleet.SendNow()
    Print(celestial.GetCoordinate(), "Trasport sent, arrive in ", ShortDur(fleet.ArriveIn))
    SleepSec(fleet.ArriveIn + 5)
    celestial.Build(one, goalrocket)
    celestial.Build(two, goalllaser)
    celestial.Build(three, goalhlaser)
    celestial.Build(four, goalgauss)
    celestial.Build(five, goalplasma)
    Print(celestial.GetCoordinate(), "Defense built") 
    SleepSec(30)
}

//loop def
sync = import("sync")
wg = new(sync.WaitGroup)
i = 0
for celestial in whereshouldbuild { //for planets: GetCachedPlanets(), for moons: GetCachedMoons(), for planets+moons use GetCachedCelestials()
    SleepRandSec(30, 40)
    wg.Add(1)  
    i += 1
    Print("starting thread", i, " for the planet ", celestial.GetCoordinate())
    //the go func start a thread for every celestial
    go func (celestial) {
        for {
            res, err = homeworld.GetResources()
            tostay = NewResources(minmet, mincrys, mindeut)
            queuedef , _, _ = celestial.GetProduction()
            defense, err = celestial.GetDefense()
            fleets, _ = GetFleets()
            slots = GetSlots()
            AvailSlots = slots.Total - slots.InUse
            destination = nil
            for fleet in fleets {
                destination = fleet.Destination
            }
            rocketinqueue = 0
            llaserinqueue = 0
            hlaserinqueue = 0
            gaussinqueue = 0
            plasmainqueue = 0
            for def in queuedef {
                queueID = def.ID
                xiteminqueue = def.Nbr
                if queueID == ROCKETLAUNCHER && xiteminqueue {
                    rocketinqueue += xiteminqueue
                } else if queueID == LIGHTLASER && xiteminqueue {
                    llaserinqueue += xiteminqueue
                } else if queueID == HEAVYLASER && xiteminqueue {
                    hlaserinqueue += xiteminqueue
                } else if queueID == GAUSSCANNON && xiteminqueue {
                    gaussinqueue += xiteminqueue
                } else if queueID == PLASMATURRET && xiteminqueue {
                    plasmainqueue += xiteminqueue
                }
            }
            sumrocket = defense.RocketLauncher + rocketinqueue
            sumllaser = defense.LightLaser + llaserinqueue
            sumhlaser = defense.HeavyLaser + hlaserinqueue
            sumgauss = defense.GaussCannon + gaussinqueue
            sumplasma = defense.PlasmaTurret + plasmainqueue
            BID, BCD, _, _ = celestial.ConstructionsBeingBuilt()
            facilities, err = celestial.GetFacilities()
            NF = facilities.NaniteFactory
            SY = facilities.Shipyard
            if !celestial.GetCoordinate().Equal(destination) && res.Sub(tostay).Gte(GetPrice(PLASMATURRET, 1)) && NF >= minNanites && SY >= minShipyard && BID != NANITEFACTORY && BID != SHIPYARD && (sumrocket < standardrocket || sumllaser < standardllaser || sumhlaser < standardhlaser || sumgauss < standardgauss || sumplasma <standardplasma)  && AvailSlots > reservedslots {
                Print(celestial.GetCoordinate(), "Started func build defense")
                buildefense(celestial, sumrocket, sumllaser, sumhlaser, sumgauss, sumplasma, ROCKETLAUNCHER, standardrocket, LIGHTLASER, standardllaser, HEAVYLASER, standardhlaser, GAUSSCANNON, standardgauss, PLASMATURRET, standardplasma)
                SleepRandSec(30, 60)
            } else if sumrocket >= standardrocket && sumllaser >= standardllaser && sumhlaser >= standardhlaser && sumgauss >= standardgauss && sumplasma >= standardplasma { 
                Print(celestial.GetCoordinate(), "Defense completed on the planet, thread " + i, " done")
                wg.Done()
                return
            }else if NF < minNanites || SY < minShipyard {
                Print(celestial.GetCoordinate(), "Need to build facilities first")
                SleepRandSec(30, 60)
            } else if (sumrocket < standardrocket || sumllaser < standardllaser || sumhlaser < standardhlaser || sumgauss < standardgauss || sumplasma < standardplasma) && !res.Sub(tostay).Gte(GetPrice(PLASMATURRET, 1)) {
                Print(celestial.GetCoordinate(), "Not enough resources, can't build defenses atm")
                SleepRandSec(30, 60)
            } else if (sumrocket < standardrocket || sumllaser < standardllaser || sumhlaser < standardhlaser || sumgauss < standardgauss || sumplasma < standardplasma) && (BID == NANITEFACTORY || BID == SHIPYARD) {
                Print(celestial.GetCoordinate(), "Nanites or Shipyard on building, can't build defenses atm")
                SleepRandSec(30, 60)
            } else if AvailSlots <= reservedslots {
                Print(celestial.GetCoordinate(), "There aren't enough available slots atm")
                SleepRandSec(30, 60)
            } else if celestial.GetCoordinate().Equal(destination) {
                Print(celestial.GetCoordinate(), "Transport already flying to this coords")
                SleepRandSec(30, 60)
            } 
            Rmin = Random(20, 40)
            Print(celestial.GetCoordinate(), "sleep", Rmin, " min")
            SleepMin(Rmin)
            continue
        }
    } (celestial)
}
wg.Wait()
Print("Defences built on every planet, congratulations")
