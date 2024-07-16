package main

import (
    "ogame/pkg/ogame"
    "nja"
    "fmt"
)

func main() {
    galaxy := int64(4)
    fromSystem := int64(1)
    toSystem := int64(3)
    interval := nja.Random(5*60*1000, 10*60*1000) // 5-10min
    
    //-------------------------------
    
    data := make(map[string][]int)
    
    for {
        for system := fromSystem; system <= toSystem; system++ {
            systemInfos, err := nja.GalaxyInfos(galaxy, system)
            if err != nil {
                nja.Print(err)
                continue
            }
            arr := make([]int, 0)
            systemInfos.Each(func(planetInfos *ogame.PlanetInfos) {
                val := 0
                if planetInfos != nil {
                    val = 1
                }
                arr = append(arr, val)
            })
            key := fmt.Sprintf("%d:%d", galaxy, system)
            if data[key] != nil && data[key] != arr {
                nja.SendTelegram(nja.TELEGRAM_CHAT_ID, "New/Removed planets in "+key)
            }
            data[key] = arr
        }
        nja.Sleep(interval)
    }
}
