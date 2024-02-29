/* Copyright: Making these scripts is very hard work. Please appriciate that and 
              don't share if you want me to keep making these cool scripts */ 
// Contact for questions, bug reports, suggestions:
// Email:   umut91c@gmail.com
// Discord: cremefresh55#1208

//Version 1.7

//Changelogs
/*
Considering last activity of spied inactives

V1.6 and lower
Now calculates time of the longest attack and switches to next planet after the longest attack is back
Spying now consideres if you have enough spy probes 
Spying now consideres reseverve Slots
Added Start / End Timer
Added Sleep Timers to be adjustable
Added option to save after farming
Added option to gather to home planet after farming
*/
 
 
//NOTE: THIS SCRIPT HAS A LOT OF INTENTIONAL DELAYS, DONT WORRY IF IT WORKS OR NOT
//########################-------Description------############################
/*
1) Shuffles the order of the farm planets to make it more human
2) Scans all the inactives without defences in set range with sleep inbetween 
   them to make it more human
3) After a huge delay (3min) it starts attack the top best planets until 
   your slots full full (ofc reserved slot is considered)
4) If another farm planet is on the list -> Repeat 1-3) for this planet
5) All farm planets have farmed now going to optionals A and B
*A) If you have send_to_home = true you will send all your ressources to your
    set home at the end of all farming planets
*B) If you set time and use_timer = true , it will start the whole script over 
   at that time. If use_timer = false it will playe the script ONES! (so you don't
   have to wait for the exact time)
NOTE: You can start the script with already used slots. From the second planet 
on, it will wait until all the slots are free and then start the scan process 
after a certain delay. This makes the script more human
*/
//########################-------Description------############################

//#########################--------SETTINGS--------########################
attack_if_last_active = 0 //this will only attack if last activity was above x min
                           //set to 0 if you just want to attack regardless of activity
                           //set 
esp_probes_for_scans = 3 // Number of Spy Probes for Spying
min_player_rank = 3000 // I recommend you set this to a reasonable number 
                       // if you don't you will slow the script down by huge amount
                       // because the sleep functions accumulate for each spy
additional_cargos = 10 // Additional cargos in %,
delay_between_spy_missions = [250,2000] //0.25 - 2 sek 
delay_after_spy_to_attack_missions = [60*1,60*2] // 1-2min - at least 1 min is recommended
delay_between_attack_missions = [3000,6000] //3-6 sek
delay_after_attack_to_next_planet = [60*2,60*3] // 2-3min

// ### TIMER ###
starttime = "@19h58" //"@06h00" 
end_after_x_hours = 16.5 // Example : 16.5 this will end the script after 16,5 hours after starttime
                      
                      
//### SHUFFLE ### farming order
use_shuffle = false; 
//### TELEGRAM ###
useTelegram = true; //Message for these steps:
                     //Scanning,Spying,Attacking , Gathering (if true)

// ### Saving or Sending to Home planet (Gathering)
home_planet = "P:2:247:7"

send_to_home = true // At the end of the whole farming and with this set true
                    // all your farm planets will send all ressources to home_planet


//### FARMER ###                   
farm_planets = ["P:4:361:9", "P:4:141:6"]//, "P:3:411:9", "P:3:88:8", "P:2:247:7", "P:2:88:9", "P:1:248:9"
sys_radius = 20 // Spy radius of systems around each planet.. always same radius
/*If you want to use the lower_upper_range system down below, you 
!!!!! ---> MUST set this to 0 <<<---!!!!! , because this setting has priority */

//FOR THIS TO USE sys_radius must be 0! :
//WARNING: IF USED: IT HAS TO HAVE AS MANY ENTRIES IN BOTH ARRAYS AS YOU HAVE FARM PLANETS
system_lower_range =[10,30] // first entry corresponds with first planet
system_upper_range = [20,40]  // first entry corresponds with first planet
//Example for: farm_planets = ["P:2:200:10", "P:3:300:9"]
/* For first planet the attack range will be  190<-[200]->220
   For second planet the attack range will be 270<-[300]->340 */
//#########################--------SETTINGS--------########################



sum_all_res = 0
all_inactives_coords=[]
all_inactive_res = []
func calc_highest_flight_time(system_distance,largecargo){
    server = GetServer()
    research = GetResearch()
    impulslvl = research.ImpulseDrive
    combustionlvl = research.CombustionDrive
    Geschwindigkeitsfaktor = 1
    speeduni = server.Settings.FleetSpeed
    Entfernung = system_distance* 95 + 2700
    speedship = 0
    if(largecargo == true){
        speedship = 7500+750*combustionlvl
        
    }else{
        if(impulslvl < 5){
        speedship = 10000+1000*combustionlvl
        }else{
            speedship = 10000+2000*impulslvl
        }
    }
    GeschwindigkeitDesLangsamstenSchiffs = speedship
    temp = (Entfernung * 10 / GeschwindigkeitDesLangsamstenSchiffs)
    time = (3500 / Geschwindigkeitsfaktor) * Sqrt(temp)+ 10
    time = time/speeduni
    return time
}

func shuffle(array) {
    for old_index = len(array) - 1; old_index > 0; old_index-- {
        new_index = Random(0, len(array) - 1)
        temp_array = array[old_index]
        array[old_index] = array[new_index]
        array[new_index] = temp_array
    }
}
func shuffle2(array,lower,upper) {
    for old_index = len(array) - 1; old_index > 0; old_index-- {
        new_index = Random(0, len(array) - 1)
        
        temp_array = array[old_index]
        temp_lower = lower[old_index]
        temp_upper = upper[old_index]
        
        array[old_index] = array[new_index]
        lower[old_index] = lower[new_index]
        upper[old_index] = upper[new_index]
        
        array[new_index] = temp_array
        lower[new_index] = temp_lower
        upper[new_index] = temp_upper
    }
}
func sort(array_res,array_koord){
    n = len(array_res)
    for i = 0; i < n-1; i++ {
        for j = 0; j < n-i-1; j++ {
            if (array_res[j] < array_res[j+1]){
                temp_res = array_res[j]
                 array_res[j] = array_res[j+1]
                 array_res[j+1] = temp_res
                temp_koord = array_koord[j]
                 array_koord[j] = array_koord[j+1]
                 array_koord[j+1] = temp_koord
                 
            }
        }
    }     
}
func send_to_home_planet() {
    for planet in farm_planets {
        if planet != home_planet {
            celestial, _ = GetCelestial(planet)
            res, _ =  celestial.GetResources()
            all_ships, _ =  celestial.GetShips()
            met =  res.Metal
            crys =  res.Crystal
            deut =  res.Deuterium - 75000
            res_to_send = NewResources(met, crys, deut)
            lc, sc, cargo = CalcFastCargo( all_ships.LargeCargo, all_ships.SmallCargo,  res_to_send.Total())
            si = NewShipsInfos()
            si.Set(LARGECARGO, lc)
            si.Set(SMALLCARGO, sc)
            coord_home, _ = ParseCoord(home_planet)
            _, fuel = FlightTime(celestial.GetCoordinate(), coord_home, HUNDRED_PERCENT, *si)
            deut -= fuel
            res_to_send = NewResources(met, crys, deut)
            fleet = NewFleet()
            fleet.SetOrigin(planet)
            fleet.SetDestination(home_planet)
            fleet.SetMission(TRANSPORT)
            fleet.SetSpeed(HUNDRED_PERCENT)
            fleet.SetResources( res_to_send)
            fleet.AddShips(LARGECARGO, lc)
            fleet.AddShips(SMALLCARGO, sc)
            f, err = fleet.SendNow()
            print("test")
            Sleep(Random(3000, 6000))
        }
    }
}
save_fleet = false
sys_a = 0
sys_b = 0
quitter = false
highest_flight_time = 0


func farmer(){
    ts = GetTimestamp() + 3600 * end_after_x_hours
    if use_shuffle == true {
        if(sys_radius > 0){
            shuffle(farm_planets)
        }else{
            shuffle2(farm_planets,system_lower_range,system_upper_range)
        }
    }
             
    for quitter == false{
        slots = GetSlots()
        hard_init_slot = slots.InUse
        planet = []
        planet_counter = 0;
        all_inactives_coords = []
        all_inactive_res = []
    
        //Spying
        for planet_counter = 0 ; planet_counter < len(farm_planets) ; planet_counter++ {
            current_planet, _ = ParseCoord(farm_planets[planet_counter])
            galaxy = current_planet.Galaxy
            system = current_planet.System
            planet, _ = GetCelestial("P:"+Itoa(current_planet.Galaxy)+":"+Itoa(current_planet.System)+":"+Itoa(current_planet.Position))
            if sys_radius > 0 {
            sys_a = system - sys_radius 
            sys_b = system + sys_radius 
            } else {
                sys_a = system - system_lower_range[planet_counter]
                sys_b = system + system_upper_range[planet_counter]
            }
            if(sys_a < 1){
                sys_a = 1
            }
            if(sys_a > 499){
                sys_a = 499
            }
            if(sys_b < 1){
                sys_b = 1
            }
            if(sys_b > 499){
                sys_b = 499
            }
            if useTelegram {
                SendTelegram(TELEGRAM_CHAT_ID, "Scanning planets for farmplanet: "+farm_planets[planet_counter])    
            }
            //Save all spy reports
            counter = 0
            for i = sys_a ; i < sys_b+1 ; i++ {
                systemInfo, _ = GalaxyInfos(galaxy, i)
                for j = 1; j < 16; j++ {
                    if(systemInfo.Position(j) != nil){
                        planetInfo = systemInfo.Position(j)
                        inactive = planetInfo.Inactive
                        rank = planetInfo.Player.Rank
                        vacation =  planetInfo.Vacation
                        banned =  planetInfo.Banned
                        if inactive == true && rank < min_player_rank && vacation == false && banned == false {
                            all_inactives_coords[counter] = planetInfo.Coordinate
                            counter = counter +1  
                        }
                    }
                }
            }   
                
            if useTelegram {
                SendTelegram(TELEGRAM_CHAT_ID, "Scanning complete! Now spying inactives")    
            }
            print("Inactive: "+len(all_inactives_coords))
            slots = GetSlots()
            slots_in_use = slots.InUse
            slots_total = slots.Total
            slots_reserved = GetFleetSlotsReserved()
            if len(all_inactives_coords) > 0 {
                for i = 0 ; i < len(all_inactives_coords) ; i++ {
                    slots = GetSlots()
                    slots_in_use = slots.InUse
                    all_ships, _ = planet.GetShips()
                    spyprobes = all_ships.EspionageProbe 
                    for {
                        if slots_in_use < slots_total-slots_reserved && spyprobes >= esp_probes_for_scans {
                            break
                        }
                        slots = GetSlots()
                        slots_in_use = slots.InUse
                        all_ships, _ =  planet.GetShips()
                        spyprobes = all_ships.EspionageProbe 
                        print("slots "+ slots_in_use)
                        print("spyprobes "+spyprobes)
                        Sleep(5000) // Waiting for free slot
                    }
                    fleet = NewFleet()
                    fleet.SetOrigin(planet)
                    fleet.SetDestination(all_inactives_coords[i])
                    fleet.SetMission(SPY)
                    fleet.AddShips(ESPIONAGEPROBE, esp_probes_for_scans)
                    fleet, err = fleet.SendNow()
                    Sleep(Random(delay_between_spy_missions[0],delay_between_spy_missions[1]))
                }    
            }
            Sleep(1000*10)
            Sleep(Random(1000*delay_after_spy_to_attack_missions[0],1000*delay_after_spy_to_attack_missions[1])) // Give 2min for deploying
                
            //Get total Res of all spyreports
            for i = 0 ; i < len(all_inactives_coords) ; i++{
                report,error1 = GetEspionageReportFor(all_inactives_coords[i])
                if(attack_if_last_active != 0){
                    if((report.LastActivity > attack_if_last_active || report.LastActivity == nil )&& report.HasFleet == true && report.HasDefenses == true && report.RocketLauncher == nil && report.LightLaser  == nil && report.HeavyLaser == nil && report.GaussCannon == nil && report.IonCannon == nil && report.PlasmaTurret == nil && report.SmallShieldDome == nil && report.LargeShieldDome == nil && report.LightFighter == nil && report.HeavyFighter == nil && report.Cruiser == nil && report.Battleship == nil && report.Battlecruiser == nil && report.Bomber == nil && report.Destroyer == nil && report.Deathstar == nil && report.SmallCargo == nil && report.LargeCargo == nil && report.ColonyShip == nil && report.Recycler == nil){
                        all_inactive_res[i] = report.Total()  
                        print("TOTAL RES: "+   all_inactive_res[i]+" COORDS: "+all_inactives_coords[i]  )
                    } else{
                        all_inactive_res[i] = 0
                        print("TOTAL RES: "+   all_inactive_res[i]+" COORDS: "+all_inactives_coords[i]  )
                    }
                }else{
                    if(report.HasFleet == true && report.HasDefenses == true && report.RocketLauncher == nil && report.LightLaser  == nil && report.HeavyLaser == nil && report.GaussCannon == nil && report.IonCannon == nil && report.PlasmaTurret == nil && report.SmallShieldDome == nil && report.LargeShieldDome == nil && report.LightFighter == nil && report.HeavyFighter == nil && report.Cruiser == nil && report.Battleship == nil && report.Battlecruiser == nil && report.Bomber == nil && report.Destroyer == nil && report.Deathstar == nil && report.SmallCargo == nil && report.LargeCargo == nil && report.ColonyShip == nil && report.Recycler == nil){
                        all_inactive_res[i] = report.Total()  
                        print("TOTAL RES: "+   all_inactive_res[i]+" COORDS: "+all_inactives_coords[i]  )
                    } else{
                        all_inactive_res[i] = 0
                        print("TOTAL RES: "+   all_inactive_res[i]+" COORDS: "+all_inactives_coords[i]  )
                    }
                }
            }
            //Check Slots
            slots = GetSlots()
            slots_in_use = slots.InUse
            slots_total = slots.Total
            slots_reserved = GetFleetSlotsReserved()
            attacks_to_make = slots_total - slots.InUse - slots_reserved 
            if useTelegram {
                SendTelegram(TELEGRAM_CHAT_ID, "Found "+attacks_to_make+" targets for planet: "+farm_planets[planet_counter])    
            }
            //Sort Res  & Coords Array by value
            if(len(all_inactives_coords) > 0){
                sort(all_inactive_res,all_inactives_coords)
            }
            // Sorting Out bad spy reports
            attacks_to_make_temp = 0
            for i = 0 ; i < len(all_inactive_res) ; i++{
                if(all_inactive_res[i] != 0){
                    attacks_to_make_temp = attacks_to_make_temp +1
                }
            }
            if(attacks_to_make_temp < attacks_to_make){
                attacks_to_make = attacks_to_make_temp
            }
            if(len(all_inactives_coords) < attacks_to_make ){
                attacks_to_make = len(all_inactives_coords)
            }
            print("Atk Temp: "+attacks_to_make_temp)
            print("Atk To Make: "+attacks_to_make)
            highest_system = 0
                have_large_cargo = false;
                if(len(all_inactives_coords) > 0 &&  len(all_inactive_res) > 0 ){
                for i = 0; i < attacks_to_make ; i++{
                    fleet = NewFleet()
                    fleet.SetOrigin(planet)
                    system = all_inactives_coords[i].System
                    system_dif = Abs(planet.GetCoordinate().System - system)
                    if(system_dif > highest_system){
                        highest_system = system_dif
                    }
                    fleet.SetDestination(all_inactives_coords[i])
                    fleet.SetMission(ATTACK)
                    //Res 
                    res = all_inactive_res[i]
                    //Calc SmallCargos
                    all_ships, _ =  planet.GetShips()
                    lc, sc, cargo = CalcFastCargo( all_ships.LargeCargo, all_ships.SmallCargo,  res/2)
                    if(lc > 0){
                        have_large_cargo  = true
                    }
                    fleet.AddShips(LARGECARGO, lc)
                    fleet.AddShips(SMALLCARGO, sc+Round(sc*additional_cargos/100))
                    fleet, err = fleet.SendNow()
                    Sleep(Random(delay_between_attack_missions[0],delay_between_attack_missions[1])) 
                }  
                highest_flight_time =  calc_highest_flight_time(highest_system,have_large_cargo )
                print("Sleep for: "+ShortDur(highest_flight_time))
            }
            ll_inactives_coords = []
            all_inactive_res = []
            planet = []
            //Pause program until slots are 0 => Ready to Scan Again
            Sleep(highest_flight_time*2*1000)
            if useTelegram {
                SendTelegram(TELEGRAM_CHAT_ID, "Attack complete session for planet: "+farm_planets[planet_counter]+" is complete" )    
            }
            Sleep(Random(delay_after_attack_to_next_planet[0],delay_after_attack_to_next_planet[1])) 
        }
        new_ts = GetTimestamp()
        if(new_ts >= ts){
            print("Break!")
            break
        }
    
    }
    if(send_to_home == true){
        send_to_home_planet() 
        if useTelegram {
            SendTelegram(TELEGRAM_CHAT_ID, "Sending all ressources to homeplanet!" )    
        }
    }
}
CronExec(starttime,farmer) // Execute callback every day at midnight
<-OnQuitCh
