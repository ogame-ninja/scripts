// Contact for questions, bug reports, suggestions:
// Email:   umut91c@gmail.com
// Discord: cremefresh55#1208

//Version 1.0
/*
This script will be executed onces and built exactly the amount of ships you set
for every planet (except the planets you put in exception).
Setting a time is also possible if you want to build on a regular basis.

*/


// SETTINGS--------------------------------------------------------------------
use_timer = false // if this is false, the script will exectue only once 
                  //if true, it will only exececute on the time you set
time = "@15h10" // Use NJT Time (see on the left below the + ScriptButton)

exception_planets = [] // Example for 2 exceptions: ["P:1:361:9", "P:4:141:6"]

wait_seconds = 10 // 10 seconds until it gets executed to prevent building by accident
                  // Press Stop Script if you startet by accident.
                  
build_lightfighter= 0
build_heavyfighter= 0
build_cruiser = 0
build_battleship = 20
build_battlecruiser = 5
build_bomber =0
build_destroyer =0
build_deathstar = 0

build_small_cargo = 10
build_large_cargo= 50
build_eprobes = 0
build_solar_sats = 0
build_recycler = 0
build_coloships = 0

// SETTINGS DONE----------------------------------------------------------------

Planets = GetPlanets()
 Sleep(wait_seconds * 1000) 



func build_fleet() {
    build = true
    for planet in Planets {
        celestial = GetCachedCelestial(planet.Coordinate)
        build = true
        for i = 0 ; i < len(exception_planets) ; i++{
            if(planet.Coordinate == exception_planets[i]){
                build = false
            }
            
            if(build == true){
                if(build_lightfighter > 0){
                     celestial.Build(LIGHTFIGHTER,build_lightfighter)
                     Sleep(Random(400,1700))
                }
               if(build_heavyfighter > 0){
                     celestial.Build(HEAVYFIGHTER,build_heavyfighter)
                        Sleep(Random(400,1700))
                     
               }
               if(build_cruiser > 0){
                      celestial.Build(CRUISER,build_cruiser)
                         Sleep(Random(400,1700))
               }
               if(build_battleship > 0){
                       celestial.Build(BATTLESHIP,build_battleship)
                          Sleep(Random(400,1700))
               }
               if(build_battlecruiser > 0){
                         celestial.Build(BATTLECRUISER,build_battlecruiser)
                            Sleep(Random(400,1700))
               }
               if(build_bomber > 0){
                     celestial.Build(BOMBER,build_bomber)
                        Sleep(Random(400,1700))
               }
                if(build_destroyer > 0){
                    
                       celestial.Build(DESTROYER,build_destroyer)
                          Sleep(Random(400,1700))
                }
                if(build_deathstar > 0){
                            celestial.Build(DEATHSTAR,build_deathstar)
                               Sleep(Random(400,1700))
                }
                if(build_small_cargo > 0){
                        celestial.Build(SMALLCARGO,build_small_cargo)
                           Sleep(Random(400,1700))
                } 
               if(build_large_cargo > 0){
                             celestial.Build(LARGECARGO,build_large_cargo)
                                Sleep(Random(400,1700))
               }
               if(build_eprobes > 0){
                         celestial.Build(ESPIONAGEPROBE,build_eprobes) 
                            Sleep(Random(400,1700))
               }
                if(build_solar_sats > 0){
                     celestial.Build(SOLARSATELLITE,build_solar_sats)
                        Sleep(Random(400,1700))
                }
                if(build_recycler > 0){
                          celestial.Build(RECYCLER,build_recycler) 
                             Sleep(Random(400,1700))
                }
                 if(build_coloships > 0){
                          celestial.Build(COLONYSHIP,build_coloships)
                             Sleep(Random(400,1700))
                 }
                 Sleep(Random(2000,4000))
                  
            }
        }
       
    
    }
}
if(use_timer){
   CronExec(time,build_fleet) // Execute callback every day at midnight 
   <-OnQuitCh
}else{
    build_fleet()
}




