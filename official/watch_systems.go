galaxy = 4
fromSystem = 1
toSystem = 3
interval = Random(5*60*1000, 10*60*1000) // 5-10min

//-------------------------------

data = {}

for {
    for system = fromSystem; system <= toSystem; system++ {
        systemInfos, err = GalaxyInfos(galaxy, system)
        if err != nil {
            Print(err)
            continue
        }
        arr = []
        systemInfos.Each(func(planetInfos) {
            arr += planetInfos == nil ? 0 : 1
        })
        key = galaxy+":"+system
        if data[key] != nil && data[key] != arr {
            SendTelegram(TELEGRAM_CHAT_ID, "New/Removed planets in "+key)
        }
        data[key] = arr
    }
    Sleep(interval)
}
