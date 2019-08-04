master = "4:212:8"
slaves = ["4:208:8", "4:212:10"]
when = ["@8h15", "@15h43"] // Will run every day at these times
lc = 1000
sc = 1000

//------------------------------------------------------------------------------

for time in when {
    entryID, err = CronExec(time, func() {
        for slave in slaves {
            fleet = NewFleet()
            fleet.SetOrigin(slave)
            fleet.SetDestination(master)
            fleet.SetMission(TRANSPORT)
            fleet.SetAllResources()
            fleet.AddShips(LARGECARGO, lc)
            fleet.AddShips(SMALLCARGO, sc)
            fleet, err = fleet.SendNow()
            Print(fleet.ID, err)
            Sleep(Random(1000, 4000)) // Sleep between 1 and 4 seconds
        }
    })
    if err == nil {
        Print("Starts cronjob " + entryID + ". Spec: " + time)
    } else {
        LogError("Failed to create cronjob for " + time + ". " + err)
    }
}

<-OnQuitCh