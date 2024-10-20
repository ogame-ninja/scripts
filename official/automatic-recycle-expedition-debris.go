origin = "1:2:3"
nbPathfindersMin = 10

for {
    s = <-OnSystemInfos
    m = s.ExpeditionDebris.Metal
    c = s.ExpeditionDebris.Crystal
    d = s.ExpeditionDebris.Deuterium
    n = s.ExpeditionDebris.PathfindersNeeded
    if n >= nbPathfindersMin {
        dest = fmt.Sprintf("%d:%d:16", s.Galaxy(), s.System())
        LogInfof("debris field at %s (M: %d, C: %d, D: %d)", dest, m, c, d)
        f = NewFleet()
        f.SetOrigin(origin)
        f.SetDestination(dest)
        f.SetMission(RECYCLEDEBRISFIELD)
        f.AddShips(PATHFINDER, n)
        fleet, err = f.SendNow()
        if err != nil {
            LogErrorf("failed to send %d pathfinders to %s: %v", n, dest, err)
            continue
        }
        LogInfof("sending %d pathfinders to %s : #%d", n, dest, fleet.ID)
    }
}
