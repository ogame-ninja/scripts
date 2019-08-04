fourHours = 4 * 60 * 60

fleet = NewFleet()
fleet.SetOrigin("M:1:2:3")
fleet.SetDestination("M:1:2:4")
fleet.SetSpeed(TEN_PERCENT)
fleet.SetMission(PARK)
fleet.SetAllResources()
fleet.SetAllShips()
fleet.SetRecallIn(fourHours)
fleet, err = fleet.SendNow()

Print(fleet.ID, err)