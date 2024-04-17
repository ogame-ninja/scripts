/***** This script is created by RockClubKASHMIR <discord @RockClubKASHMIR#8058> *****\
  discord channel for your personal orders and support - https://discord.gg/kbsdRCB
  For Donation: https://www.paypal.com/paypalme/RockClubKASHMIR
  
 DESCRIPTION

A fully automated fleet saver that lets you choose between three types of missions for fleet saving.
The fleet saver starts at a randomizing time (between 0 and 35 minutes of the time you already set).
You can choose how many hours your fleets must fly in total.
It saves all your fleets, even from planets (if you have a moon on this planet, the fleet will be sent to the moon first, and then the fleet saver will save your fleet).
It goes to sleep mode only if all your fleets are saved.
The fleet saver recalls your fleets at half the flight time of your fastest fleet and goes into sleep mode for the rest of the flight time.

\******************************************************/
FleetSaverStartTime = "01:07:00" // Must be in format HH:MM:SS only!
FleetTrip = 7 // Set how long you want your fleets to be in the air (in hours)

SaveByColonize = true // It will save your fleets to colonize mission (you should have ColonyShip)
SaveToDebris = false // It will save your fleets to the debris (you should have Recycler)
SaveByDeployment = false // It will save your fleets to deployment mission

//====== DO NOT CHANGE THE CODE BELOW ======\\
if IsLoggedIn() {Login()}
if !IsNJAEnabled() {EnableNJA()}
if !IsRunningDefenderBot() {StartDefenderBot()}
minimumTravelTime = FleetTrip*1800 //calculate a half time of total fleet trip
Missions = [COLONIZE, RECYCLEDEBRISFIELD, PARK] //List with the mission types
if SaveByColonize == true {
    SaveToDebris = false
    SaveByDeployment = false
}

if SaveToDebris == true {
    SaveByColonize = false
    SaveByDeployment = false
}

if SaveByDeployment == true {
    SaveByColonize = false
    SaveToDebris = false
}
if SaveByColonize == false && SaveToDebris == false && SaveByDeployment == false {SaveByColonize = true}
//The fleet saver
func SmartFleetSaver() {
    planets = GetPlanets()
    Sleep(1000)
    moons = GetMoons()
    saved = Random(0, 35) //Randomise start from 0 to 35 minutes
    LogInfo("The Fleet saver will start after "+ShortDur(saved*60))
    SleepMin(saved)
    FleetMarker = 0
    NotSavedShips = 0
    shorterTime = 0
    for planet in planets {
        origin = planet.Coordinate
        m, _ = ParseCoord("M:"+Itoa(origin.Galaxy)+":"+Itoa(origin.System)+":"+Itoa(origin.Position))
        myShips, _ = planet.GetShips()
        for myMoon in moons {
            Sleep(Random(100, 250))
            if m == myMoon.Coordinate {
                if Cargo(myShips) > 0 {
                    fTime, fuel = FlightTime(origin, myMoon.Coordinate, HUNDRED_PERCENT, myShips, PARK)
                    fleet = NewFleet()
                    fleet.SetOrigin(origin)
                    fleet.SetDestination(myMoon.Coordinate)
                    fleet.SetMission(PARK)
                    fleet.SetSpeed(HUNDRED_PERCENT)
                    fleet.SetAllShips()
                    fleet.SetAllResources()
                    _, err = fleet.SendNow()
                    if err == nil {
                        LogInfo("The fleet is sent from planet "+origin+" to your moon "+myMoon.Coordinate+". Please wait "+ShortDur(fTime))
                        SleepSec(fTime+2)
                    } else {
                        LogWarn("The fleet from "+origin+" cannot be sent to the moon - "+err)
                        NotSavedShips = 1
                    }
                }
                myShips, _ = myMoon.GetShips()
                origin = myMoon.Coordinate
            }
        }
        // Saving your current fleet
        if Cargo(myShips) > 0 {
            Mission = Missions[0]
            DestinationCoordinate = nil
            currentTimeOfTrip = 0
            if SaveByColonize == true {
                emptyPlanetCoordinate, flightTime, fuel, err = FindEmptyPlanetWithMinimumTravelTime(origin, myShips, minimumTravelTime)
                if err == nil {
                    DestinationCoordinate = emptyPlanetCoordinate
                    currentTimeOfTrip = flightTime
                    if shorterTime == 0 || shorterTime > currentTimeOfTrip {shorterTime = currentTimeOfTrip}
                } else {LogWarn(err)}
            }
            if SaveToDebris == true {
                debrisCoordinate, flightTime, fuel, err = FindDebrisFieldWithMinimumTravelTime(origin, myShips, minimumTravelTime)
                if err == nil {
                    Mission = Missions[1]
                    DestinationCoordinate = debrisCoordinate
                    currentTimeOfTrip = flightTime
                    if shorterTime == 0 || shorterTime > currentTimeOfTrip {shorterTime = currentTimeOfTrip}
                } else {LogWarn(err)}
            }
            if SaveByDeployment == true {
                TotalFuel = 0
                Mission = Missions[2]
                for moon in moons {
                    if moon.Coordinate != m {
                        flightTime, fuel = FlightTime(origin, moon.Coordinate, TEN_PERCENT, myShips, Mission)
                        if TotalFuel == 0 || TotalFuel > fuel {
                            TotalFuel = fuel
                            DestinationCoordinate = moon.Coordinate
                            currentTimeOfTrip = flightTime
                            if shorterTime == 0 || shorterTime > currentTimeOfTrip {shorterTime = currentTimeOfTrip}
                        }
                    }
                }
            }
            if shorterTime > 0 {
                fTime, fuel = FlightTime(origin, DestinationCoordinate, TEN_PERCENT, myShips, Mission)
                fleet = NewFleet()
                fleet.SetOrigin(origin)
                fleet.SetDestination(DestinationCoordinate)
                fleet.SetMission(Mission)
                fleet.SetSpeed(TEN_PERCENT)
                fleet.SetAllShips()
                fleet.SetAllResources()
                _, err = fleet.SendNow()
                if err == nil {
                    LogInfo("The fleet from "+origin+" is successfully saved to "+DestinationCoordinate)
                } else {
                    LogWarn("The fleet from "+origin+" cannot be saved - "+err)
                    NotSavedShips = 1
                }
            } else {
                LogWarn("The fleet from "+origin+" cannot be saved - "+err)
                NotSavedShips = 1
            }
        }
    }
    
    //Start sleep mode
    if shorterTime > 0 && NotSavedShips == 0 {
        if shorterTime > minimumTravelTime {shorterTime = minimumTravelTime}
        LogInfo("All fleets are successfully saved!")
        LogInfo("Starting the sleep mode!\nThe bot will be enabled after "+ShortDur(shorterTime-10))
        DisableNJA()
        SleepSec(shorterTime-10)
        if IsLoggedIn() {Login()}
        if !IsNJAEnabled() {EnableNJA()}
        if !IsRunningDefenderBot() {StartDefenderBot()}
        LogInfo("The bot is enabled at "+NowTimeString())
        fleets, _ = GetFleets()
        for fleet in fleets {
            if fleet.ReturnFlight == false {
                CancelFleet(fleet.ID)
                LogInfo("The fleet sent from "+fleet.Origin+" is successfully canceled!")
                Sleep(500)
            }
        }
        LogInfo("Starting the sleep mode!\nThe bot will be enabled after "+ShortDur(shorterTime-10))
        DisableNJA()
        SleepSec(shorterTime-10)
        if IsLoggedIn() {Login()}
        if !IsNJAEnabled() {EnableNJA()}
        if !IsRunningDefenderBot() {StartDefenderBot()}
        LogInfo("The bot is enabled at "+NowTimeString())
    } else {LogWarn("The bot cannot be going into sleep mode!")}
}
CronExec(FleetSaverStartTime, SmartFleetSaver)

<-OnQuitCh
