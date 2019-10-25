// Contact for questions, bug reports, suggestions:
// Email:   umut91c@gmail.com
// Discord: cremefresh55#1208

//Version 1.2

//Goal: Auto build def again after getting attacked

// Adjust factor manually along universe progression. you can also use numbers like 1.3 or 1.57
// These are the factors from LEFT TO RIGHT for every planet in the same order as you see in your browser from TOP to BOTTOM.
// NOTE: When you change the order of the planets in your settings, you MUST change the order in the list "factors" also
// WARNING: You MUST put as many numbers seperated with a comma as you have planets! Dont worry it will not work if you forget and spit out an error

// SETTINGS--------------------------------------------------------------------
factors = [1 , 2 ,1 ,0 ,1 ,1 ,1] // Example for 7 Planets , if you dont want to include a planet, just set to 0 

//Explain: With the factors above every planet except planet 2 and planet 4 will 
//automaticly build the defense you setup below. Planet 2 will always 
//build twice the amount you setup below. Planet 4 will build no defence.

// Here you set the ratio of the defences to be build (change number to your liking)
rocket_launcher_desired = 10
light_laser_desired= 50
heavy_laser_desired= 20
gauss_cannon_desired= 5
ion_cannon_desired=  10
plasma_desired= 1
anti_balistic_desired = 2

checkInterval = 5 // Check every 5 min -- also change this to your liking, but 5 min is fine, if not to low
// SETTINGS DONE----------------------------------------------------------------

Planets = GetPlanets()

//ERROR HANDLER
//This is will spit out an error and stop the script, if you have more or less values in factors than actual planets
if(len(factors) != len(Planets)){
    print("You don't have as many factor entries as you have Planets!!")
}
factors[len(Planets)-1]

for{
    i=0;
    for planet in Planets {
        celestial = GetCachedCelestial(planet.Coordinate)
        allDefense, _ = celestial.GetDefense()
        
        productionLine = GetProduction(planet.ID)[0]
        if(len(productionLine)) == 0{
            if(allDefense.RocketLauncher < Round(rocket_launcher_desired * factors[i])){
                celestial.Build(ROCKETLAUNCHER, Round(rocket_launcher_desired * factors[i]) - allDefense.RocketLauncher)
            }
            if(allDefense.LightLaser <  Round(light_laser_desired*factors[i])){
                 celestial.Build(LIGHTLASER,Round(light_laser_desired*factors[i])- allDefense.LightLaser)
            }
            if(allDefense.HeavyLaser < Round(heavy_laser_desired*factors[i])){
                 celestial.Build(HEAVYLASER,Round(heavy_laser_desired*factors[i])- allDefense.HeavyLaser)
            }
            if(allDefense.GaussCannon < Round(gauss_cannon_desired*factors[i])){
                celestial.Build(GAUSSCANNON,Round(gauss_cannon_desired*factors[i])- allDefense.GaussCannon)
            }
            if(allDefense.IonCannon< Round(ion_cannon_desired*factors[i])){
                celestial.Build(IONCANNON,Round(ion_cannon_desired*factors[i]) - allDefense.IonCannon)
            }
            if(allDefense.PlasmaTurret< Round(plasma_desired*factors[i])){
                celestial.Build(PLASMATURRET,Round(plasma_desired*factors[i])- allDefense.PlasmaTurret)
            }
            if(allDefense.SmallShieldDome == 0){
               celestial.Build(SMALLSHIELDDOME,1)
            }
            if(allDefense.LargeShieldDome == 0){
                celestial.Build(LARGESHIELDDOME,1)
            }
            if(allDefense.AntiBallisticMissiles  < Round(anti_balistic_desired*factors[i])){
               celestial.Build(ANTIBALLISTICMISSILES,Round(anti_balistic_desired*factors[i]) - allDefense.AntiBallisticMissiles )
            }
        }
        i = i+1
    }
   Sleep(checkInterval * 60 * 1000) 
}