//====== This script is created by RockClubKASHMIR ======

fromSystem = 1 // Your can change this value as you wish
toSystem = 339 // Your can change this value as you wish
shipsList = {PATHFINDER: 5, LARGECARGO: 300, ESPIONAGEPROBE: 11, BOMBER: 1, DESTROYER: 1}// Your can change ENTYRE List, even to left only 1 type of ships!
//-------
curSystem = fromSystem
origin = nil
master = 0
nbr = 0
err = 0
slots = GetSlots().ExpInUse
for celestial in GetCachedCelestials() {
    ships, _ = celestial.GetShips()
    flts = 0
    for ShipID in shipsList {
        if ships.ByID(ShipID) > 0 {
            flts = flts + ships.ByID(ShipID)
        }
    }
    if flts > master {
        master = flts
        origin = celestial // Your Planet(or Moon)
    }
}
if origin != nil {
    Print("Your origin is "+origin.Coordinate)
    for system = curSystem; system <= toSystem; system++ {
        Destination = NewCoordinate(origin.GetCoordinate().Galaxy, system, 16, PLANET_TYPE)
        if slots < GetSlots().ExpTotal {
            ships, _ = origin.GetShips()
            if Destination != 0 {
                Sleep(Random(6000, 10000))
                f = NewFleet()
                f.SetOrigin(origin)
                f.SetDestination(Destination)
                f.SetSpeed(HUNDRED_PERCENT)
                f.SetMission(EXPEDITION)
                for id, nbr in shipsList {
                    if ships.ByID(id) > 0 {
                        if ships.ByID(id) < nbr {nbr = ships.ByID(id)}
                        f.AddShips(id, nbr)
                    }
                }
                a, err = f.SendNow()
                if err == nil {
                    Print("The ships are sended successfully to "+Destination)
                    SendTelegram(TELEGRAM_CHAT_ID, "The ships are sended successfully to "+Destination)
                } else {
                    Print("The fleet is NOT sended! "+err)
                    SendTelegram(TELEGRAM_CHAT_ID, "The fleet is NOT sended! "+err)
                    curSystem = system-1
                    slots = GetSlots().ExpTotal
                }
            }
        } else {
            for slots == GetSlots().ExpTotal {
                if err != 0 {
                    Print("Please wait till ships lands! Recheck after "+ShortDur(2*60))
                    Sleep(120000)
                    ships, _ = origin.GetShips()
                    for ShipID in shipsList {
                        if ships.ByID(ShipID) > 0 {slots = GetSlots().ExpInUse}
                    }
                } else {
                    Print("All Fleet slots are busy now! Please, wait "+ShortDur(2*60))
                    Sleep(2*60*1000)
                    slots = GetSlots().ExpInUse
                }
            }
        }
        if system >= toSystem {
            curSystem = fromSystem-1
            system = curSystem
        }
    }
} else {Print("You don't have ships from the Ships list on your Planets/Moons!")}
