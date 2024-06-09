// Repatriate all resources from all planets and moons to the master coordinate

master = "1:2:3"

//------------------------------------------------------------------------------
masterCoord, _ = ParseCoord(master)
for celestial in GetCachedCelestials() {
	if celestial.GetCoordinate().Equal(masterCoord) {
		continue
	}
	resources, err = celestial.GetResources()
	ships, err = celestial.GetShips()
	lc, sc, cargo = CalcFastCargo(ships.LargeCargo, ships.SmallCargo, resources.Total())
    fleet = NewFleet()
    fleet.SetOrigin(celestial)
    fleet.SetDestination(master)
    fleet.SetMission(TRANSPORT)
    fleet.SetAllResources()
    fleet.AddShips(LARGECARGO, lc)
    fleet.AddShips(SMALLCARGO, sc)
    fleet, err = fleet.SendNow()
    Print(fleet.ID, err)
    SleepRandSec(1, 4)
}
