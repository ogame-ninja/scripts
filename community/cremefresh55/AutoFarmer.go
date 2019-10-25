// Contact for questions, bug reports, suggestions:
// Email:   umut91c@gmail.com
// Discord: cremefresh55#1208

//Version 2.2
//This script will start a farming sessions for all the planets and moons you
//have setup at exactly the time you have set. 

// AUTO FARMING________________________________________________________________

// SETTINGS--------------------------------------------------------------------
time = "@01h45" // Use NJT Time (see on the left below the + ScriptButton)
func autofarm() {
min_ressources_to_attack = 100000 //Minimal ressources to attack
min_rank_to_attack = 1500 //Minimal rank to atatck
esp_probes_for_scans = 5 // Amount of Spy Probes to use
//Settings to make it more human:
use_random_probnumber = true // Set false/true :Randomizeses around your esp_probes_for_scan value by -1 , +1, 
randomize_attack_sys = 10 // Example: 10 randomizes all of your set ranges below by -10 , +10

//This is most important part (PLEASE READ CAREFULLY): 
//Step 1: Put exactly [] inside planets_range as you have planets. 
//Example: You have 4 planets:  planets_range = [ [],[],[],[] ]
//Note: The planets from left to right ARE the planets from top to bottom on your screen. 
//Warning: When you change the order of the planets in your screen, you have to change the order in the list also!
//Step 2: Put the attack range for every planet. Empty [] slots will be excluded from session creation
//Example:planets_range= [ [100,300],[1,499], [] ,[200,400] ]  -> there will not be any sessions for the 3th PlanetBuildingsArr
//Step 3: You must do the same for moons_range

planets_range = [ [],    []    ,[]   ,[]   ,[]   ,[] ,[]  ] // This is my setup for 7 Planets
moons_range = [200-300]  // I have no moons so this is empty
// SETTINGS DONE----------------------------------------------------------------
// Common Mistakes:
// 1: Not correct numbers of [] as you have planets/moons 
// 2: Forgetting a "," between two entries []
// 3: Not updating lists after colonization and/or getting a moon
// _____________________________________________________________________________

// FINAL NOTE:
// Choose settings so that you are finished by the time this script is called again
// Thus, your farming sessions will not accumulate if you are gone for a few days


//__________________ DONT WORRY ABOUT THIS PART AND BELOW_______________________

   
    //ERROR Handler
    Planets = GetPlanets()
    Moons = GetMoons()
    //This is will spit out an error and stop the script, if you have more or less entries in planets_to_attack than actual planets
    if(len(planets_range) != len(Planets)){
        print("You don't have as many planet entries as you have Planets!!")
        print("Did you colonized a new planet recently??")
    }
    if(planets_range != []){
       planets_range[len(Planets)-1] 
    }
    if(planets_range == []){
        print("You forgot to put in any slots for planets_range.. You must atleast have 1 planet right?")
        planets_range[len(Planets)-1] 
    }
    if(len(moons_range) != len(Moons)){
        print("You don't have as many moon entries as you have Moons!!")
        print("Did you get a moon recently?")
    }
    if(moons_range != []){
       moons_range[len(Moons)-1]
    }
    
    i=0
    for moon in Moons {
        if(use_random_probnumber == true){
            esp_probes_for_scans = Random(esp_probes_for_scans-1,esp_probes_for_scans+1)
        }
        sys_a = 0
        sys_b = 0
        if(moons_range[i] != []){
            celestial = GetCachedCelestial(moon.Coordinate)
            coordinate = celestial.GetCoordinate()
            galaxy = coordinate.Galaxy
            sys_a=moons_range[i][0] + Random(-randomize_attack_sys,randomize_attack_sys)    
            if(sys_a < 1){
                sys_a = 1
            }
            if(sys_a > 499){
                sys_a = 499
            }
            sys_b=moons_range[i][1]  + Random(-randomize_attack_sys,randomize_attack_sys)   
            if(sys_b < 1){
                sys_b = 1
            }
            if(sys_b > 499){
                sys_b = 499
            }
            NewFarmingSession(moon.GetID(), galaxy,  sys_a,  sys_b, esp_probes_for_scans, 0,min_ressources_to_attack, 0, 0, 0, 0, min_rank_to_attack, false, true, false, false, HUNDRED_PERCENT, 1, 2, 3)
            }
        i = i+1
    }
    i=0;
    for planet in Planets {
        if(use_random_probnumber == true){
            esp_probes_for_scans = Random(esp_probes_for_scans-1,esp_probes_for_scans+1)
        }
        sys_a = 0
        sys_b = 0
        if(planets_range[i] != []){
            celestial = GetCachedCelestial(planet.Coordinate)
            coordinate = celestial.GetCoordinate()
            galaxy = coordinate.Galaxy
            sys_a=planets_range[i][0] + Random(-randomize_attack_sys,randomize_attack_sys)   
              if(sys_a < 1){
                sys_a = 1
            }
            if(sys_a > 499){
                sys_a = 499
            }
            sys_b=planets_range[i][1] + Random(-randomize_attack_sys,randomize_attack_sys)   
              if(sys_b < 1){
                sys_b = 1
            }
            if(sys_b > 499){
                sys_b = 499
            }
           
            NewFarmingSession(planet.GetID(), galaxy,  sys_a,  sys_b, esp_probes_for_scans, 0,min_ressources_to_attack, 0, 0, 0, 0, min_rank_to_attack, false, true, false, false, HUNDRED_PERCENT, 1, 2, 3)
         
            }
        i = i+1
    }
}
CronExec(time,autofarm) // Execute callback every day at midnight
<-OnQuitCh