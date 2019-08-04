origin = GetCachedCelestial("4:212:6")
toBuild = {LIGHTFIGHTER: 6, HEAVYFIGHTER: 4}

for {
    Print("Need to build:", toBuild)
    for unitID, nbr in toBuild {
        unitPrice = GetPrice(unitID, 1)
        res, _ = origin.GetResources()
        canBuild = res.Div(unitPrice)
        if canBuild > 0 {
            if canBuild > nbr {
                canBuild = nbr
            }
            Build(origin.GetID(), unitID, canBuild)
            Print("Building: "+canBuild+" "+ID2Str(unitID))
            toBuild[unitID] = toBuild[unitID] - canBuild
            if toBuild[unitID] <= 0 {
                toBuild[unitID] = nil
            }
        }
    }
    if len(toBuild) == 0 {
        break
    }
    Sleep(5 * 60 * 1000) // 5min
}
Print("All ships built")