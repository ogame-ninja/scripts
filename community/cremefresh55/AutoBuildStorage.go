// Contact for questions, bug reports, suggestions:
// Email:   umut91c@gmail.com
// Discord: cremefresh55#1208

//Version 1.1

// Goal: We always want to have enough storage capacity, so we can produce ressources
// Description: This script automatically build the next storage level, when
// any storage is full (on any planet)

//SETTINGS----------------------------------------------------------------------
checkInterval = 10  // Check every 10 min  
//SETTINGS DONE-----------------------------------------------------------------

Planets = GetPlanets()

for {
    for planet in Planets {
        celestial = GetCachedCelestial(planet.Coordinate)
        buildingID, buildingCountdown, researchID, researchCountdown = celestial.ConstructionsBeingBuilt()
        
        if(buildingCountdown == 0){
            resources, err = GetResourcesDetails(celestial.GetID())
            if(resources.Metal.Available >= resources.Metal.StorageCapacity){
                Build(celestial.GetID(), METALSTORAGE, 0) 
                print("Build Metall storage on: " + planet)
            }
            if(resources.Crystal.Available >= resources.Crystal.StorageCapacity){
                Build(celestial.GetID(), CRYSTALSTORAGE, 0) 
                print("Build Crystal storage on: " + planet)
            }
            if(resources.Deuterium.Available >= resources.Deuterium.StorageCapacity){
                Build(celestial.GetID(), DEUTERIUMTANK, 0) 
                print("Build Deuterium storage on: " + planet)
            }
        }
    }
    Sleep(checkInterval * 60 * 1000) 
}