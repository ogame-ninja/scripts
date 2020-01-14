// Contact for questions, bug reports, suggestions:
// Email:   umut91c@gmail.com
// Discord: cremefresh55#1208

//Version 1.0

// Goal: We always want to have enough capacity to save all our resources.
// Description: This script automaticly builds Large Cargos when any planet has  
// more resources than total ship capacity. 
// Note: It only builds Large Cargos when no ships/defences already being build

// SETTINGS--------------------------------------------------------------------
checkInterval = 10  // Check every 10 min  
// SETTINGS DONE---------------------------------------------------------------



Planets = GetPlanets()
researches = GetResearch()
hypertech = researches.HyperspaceTechnology
smallCargoCap = 0.05*hypertech*5000 + 5000
for {
    for planet in Planets {
        celestial = GetCachedCelestial(planet.Coordinate)
        
        //Calculate all total Resources on a planet
        resources, err = GetResources(celestial.GetID())
        totalRes = resources.Total()
    
        //Calculate the total capacity of all the ships combined on a planet
        allShips, _ = celestial.GetShips()
        totalShipsCapacity = Cargo(allShips)
        
        //Only build when no ships already being build
        productionLine = GetProduction(planet.ID)[0]
        if(len(productionLine) == 0 && totalRes > totalShipsCapacity){
            SmallCargosToBuild = Round((totalRes-totalShipsCapacity)/smallCargoCap)
            celestial.Build(SMALLCARGO, SmallCargosToBuild)  
            print("Too few capacity to save all the ressources")
            print("Build Small Cargo on planet: "+planet.Coordinate)
        }
    }
    Sleep(checkInterval * 60 * 1000) // Sleep 10min
}

