// Contact for questions, bug reports, suggestions:
// Email:   umut91c@gmail.com
// Discord: cremefresh55#1208

//Version 1.0

//Goal: Due to spy probes getting wrecked from time to time, we always want
//to have enough spy probes so that our farming session keeps working properly
//Description: Set the variable "spy_probes_desired" to your desired amount
//of spy probes you want to have at least for each planet


//SETTINGS----------------------------------------------------------------------
spy_probes_desired = 50
checkInterval = 10 // Check every 10 min  
//SETTINGS DONE-----------------------------------------------------------------


Planets = GetPlanets()
for{
    for planet in Planets {
        celestial = GetCachedCelestial(planet.Coordinate)
        allShips, _ = celestial.GetShips()
        
        //Only build when no ships already being build
        productionLine = GetProduction(planet.ID)[0]
        if(len(productionLine) == 0 && allShips.EspionageProbe < spy_probes_desired){
            celestial.Build(ESPIONAGEPROBE,spy_probes_desired - allShips.EspionageProbe)
            print("Build Spy Probes on: "+planet.Coordinate)
        }
    }
   Sleep(checkInterval * 60 * 1000) 
}