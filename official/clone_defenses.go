master = GetCachedCelestial("4:119:8")
slave = GetCachedCelestial("4:126:8")

for {
    masterDef, _ = master.GetDefense()
    slaveDef, _ = slave.GetDefense()
    prodQueue, _ = slave.GetProduction() // get build queue
    
    // DefencesArr is a built-in array that contains all the defense entity IDs
    for defID in DefencesArr {
        delta = masterDef.ByID(defID) - slaveDef.ByID(defID)
        
        // Remove defenses that are already in build queue
        for item in prodQueue {
            if item.ID == defID {
                delta -= item.Nbr
            }
        }
        
        if delta > 0 {
            slave.BuildDefense(defID, delta)
            Print("Build", delta, defID) // eg: Build 12 RocketLauncher
        }
    }
    
    Sleep(60 * 60 * 1000) // Check again in 1 hour
}