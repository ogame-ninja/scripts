package main

import "nja/pkg/nja"

func main() {
	for {
		celestials, _ := nja.GetCelestials()
		if len(celestials) > 0 {
			celestials[nja.Random(0, int64(len(celestials)-1))].GetFacilities()
		}
		nja.Sleep(nja.Random(5*60*1000, 7*60*1000))
	}
}
