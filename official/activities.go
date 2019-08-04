for {
    celestials, _ = GetCelestials()
    if len(celestials) > 0 {
        celestials[Random(0, len(celestials)-1)].GetFacilities()
    }
    Sleep(Random(5*60*1000, 7*60*1000))
}