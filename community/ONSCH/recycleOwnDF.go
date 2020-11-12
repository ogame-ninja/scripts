func checkFreeSlots() {
	for {
		slots = GetSlots()
		if slots.InUse >= (slots.Total - GetFleetSlotsReserved()) {
			fleet, _ = GetFleets(); Sleep(Random(500,800))
			waitTime = 0
			for f in fleet {
				if waitTime == 0 || f.BackIn < waitTime {
					waitTime = f.BackIn
				}
			}
			LogInfo("[RECYCLE OWN DF] dont have slot for Recycler, wait", ShortDur(waitTime))
			Sleep((waitTime + Random(5, 10)) * 1000)
		} else {
			break
		}
	}
}

allCelestials = GetCachedCelestials()

for celestial in allCelestials {
	if celestial.GetType() == PLANET_TYPE {
		planetInfo, err = GetPlanetInfo(celestial.Coordinate); Sleep(Random(500,800))

		rn = planetInfo.Debris.RecyclersNeeded
		
		if rn == 0 {
			LogInfo("[RECYCLE OWN DF] No Debris field to recycle on " + celestial.Coordinate.Debris())
			continue
		}

		ships, _ = GetShips(celestial.ID); Sleep(Random(500,800))

		if ships.Recycler == 0 {
			LogWarn("[RECYCLE OWN DF] No Recycler on " + celestial.Coordinate)
			continue
		}

		for {
		    rnToSend = 0
			if rn > 0 {
				if rn > ships.Recycler {
					rnToSend = ships.Recycler
				} else {
					rnToSend = rn
				}
				//wait for free slot
				checkFreeSlots()

				//send recs
				f = NewFleet()
				f.SetOrigin(celestial.Coordinate.Planet())
				f.SetDestination(celestial.Coordinate.Debris())
				f.SetSpeed(HUNDRED_PERCENT)
				f.SetMission(RECYCLEDEBRISFIELD)
				f.AddShips(RECYCLER, rnToSend)
				
				fleet, err = f.SendNow(); Sleep(Random(500,800))

				rn = rn - rnToSend

				if rn > 0 {
					LogInfo("[RECYCLE OWN DF] Send " + fleet.Ships.Recycler + " Recycler to recycle Debris Field at " + fleet.Destination + " | wait " + ShortDur(fleet.BackIn) + " to send Recycler again")
					Sleep((fleet.BackIn*1000) + Random(2000,5000))
				} else {
					LogInfo("[RECYCLE OWN DF] Send " + fleet.Ships.Recycler + " Recycler to recycle Debris Field at " + fleet.Destination)
				}
			} else {
				break
			}
		}
	}
}
LogInfo("[RECYCLE OWN DF] All own Debris Fields are recycled")