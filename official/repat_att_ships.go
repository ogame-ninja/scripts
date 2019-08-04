masterCoord = "1:2:3"

//------------------------------------------------------------------------------

attShips = [LIGHTFIGHTER, HEAVYFIGHTER, CRUISER, BATTLESHIP, BOMBER, DESTROYER, DEATHSTAR, BATTLECRUISER]

for celestial in GetCachedCelestials() {
    ships, _ = celestial.GetShips()
    fleet = NewFleet()
    fleet.SetOrigin(celestial)
    fleet.SetDestination(masterCoord)
    fleet.SetMission(PARK)
    for shipID in attShips {
        fleet.AddShips(shipID, ships.ByID(shipID))
    }
    fleet, err = fleet.SendNow()
    if err != nil {
        LogError(err)
    } else {
        Print(fleet)
    }
}