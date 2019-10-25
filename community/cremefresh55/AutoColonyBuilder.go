// Contact for questions, bug reports, suggestions:
// Email:   umut91c@gmail.com
// Discord: cremefresh55#1208

//Version 2.3
//GOAL: For a new universe or a new colony we want to build in perfect order

//NOTE: At any giving point the new step value will be printed! 
//In case the script stops, you have to put in the latest step 
//value here:
step = 1

//SETTING----------------------------------------------------------------------
//______IMPORTANT!!______ Put in the planet you want this script to run for:
celestial, _ = GetCelestial("P:1:88:10")
include_research = true //true , for new universe; false, for building up a colony
//SETTING DONE-----------------------------------------------------------------
//The number in any list at the bottom is the step number in which everything should be BuildingsArr
//The given Example below follows the QUICK START GUIDE: to fast build small cargo
// https://ogame.fandom.com/wiki/Quick_Start_Guide
//________________________________________________________________________________
// FEEL FREE TO MAKE YOUR OWN ORDERS
//NOTE: Don't worry if you make any mistakes, like forgeting a number or 
//putting a number twice: I have error checker
//Plants
metalmine = [2,3,5,6,10,17,18,35,40]
crystalmine = [8,11,12,15,20,31,39]
deutsynth = [14,21,23,24,26,34]
solarplant = [1,4,7,9,13,16,19,22,25,33,38]

//Facilities
robofactory = [27,28]
shipyard = [30,32]
lab = [29]

//Research 
energy = [36]
combustion = [37,41]
computer = []
laser = []
espionage = []
impuls = []
astro =[]

//Ships
smallcargo = [42]
eprobes =[]
coloship =[]


//__________________DONT WORRY ABOUT THE PART BELOW_____________________________
//Error Checker
all_length = len(solarplant)+len(metalmine)+len(crystalmine)+len(deutsynth)+len(robofactory)+len(shipyard)+len(lab)+len(energy)+len(combustion)+len(computer)+len(laser)+len(smallcargo)+len(eprobes)+len(espionage)+len(impuls)+len(astro)+len(coloship)
step_check = 0
for q = 0; q < all_length;q++{
    step_check = step_check +1
    amount_counter = 0
      //Check Solarplant
    for i = 0; i < len(solarplant); i++ {
        if(solarplant[i] == step_check){
            amount_counter = amount_counter+1
        }
    }
    //Check Metallmine
    for i = 0; i < len(metalmine); i++ {
        if(metalmine[i] == step_check){
          amount_counter = amount_counter+1
        }
    }
      //Check Crystalmine
    for i = 0; i < len(crystalmine); i++ {
        if(crystalmine[i] == step_check){
             amount_counter = amount_counter+1
        }
    }
      //Check Deutsynth
    for i = 0; i < len(deutsynth); i++ {
        if(deutsynth[i] == step_check){
          amount_counter = amount_counter+1
        }
    }
    //Check RoboFactory
    if(robofactory != []){
        for i = 0; i < len(robofactory); i++ {
            if(robofactory[i] == step_check){
               amount_counter = amount_counter+1
            }
        }
    }
    //Check Shipyard
     if(shipyard != []){
        for i = 0; i < len(shipyard); i++ {
            if(shipyard[i] == step_check){
               amount_counter = amount_counter+1
            }
        }
     }
    //Check Lab
     if(lab != []){
        for i = 0; i < len(lab); i++ {
            if(lab[i] == step_check){
                amount_counter = amount_counter+1
            }
        }
     }
    //Energy
     if(energy != []){
        for i = 0; i < len(energy); i++ {
            if(energy[i] == step_check){
               amount_counter = amount_counter+1
            }
        }
     }
    //Combustion
     if(combustion != []){
        for i = 0; i < len(combustion); i++ {
            if(combustion[i] == step_check){
                 amount_counter = amount_counter+1
            }
        }
     }
    //Computer
       if(computer != []){
        for i = 0; i < len(computer); i++ {
            if(computer[i] == step_check){
                amount_counter = amount_counter+1
            }
        }
       }
    //Laser
       if(laser != []){
        for i = 0; i < len(laser); i++ {
            if(laser[i] == step_check){
            amount_counter = amount_counter+1
            }
        }
       }
       
    //Small Cargo
       if(smallcargo != []){
        for i = 0; i < len(smallcargo); i++ {
            if(smallcargo[i] == step_check){
             amount_counter = amount_counter+1
            }
        }
       }
        //Eprobes
       if(eprobes != []){
        for i = 0; i < len(eprobes); i++ {
            if(eprobes[i] == step_check){
             amount_counter = amount_counter+1
            }
        }
       }
        //Eprobes
       if(espionage != []){
        for i = 0; i < len(espionage); i++ {
            if(espionage[i] == step_check){
             amount_counter = amount_counter+1
            }
        }
       }
         //Impuls
       if(impuls  != []){
        for i = 0; i < len(impuls); i++ {
            if(impuls[i] == step_check){
             amount_counter = amount_counter+1
            }
        }
       }
        // Astro
       if(astro  != []){
        for i = 0; i < len(astro); i++ {
            if(astro[i] == step_check){
             amount_counter = amount_counter+1
            }
        }
       }
          // Coloship
       if(coloship  != []){
        for i = 0; i < len(coloship); i++ {
            if(coloship[i] == step_check){
             amount_counter = amount_counter+1
            }
        }
       }
    if(amount_counter == 0){
     print("You forgot to put the step number: "+step_check)
    metalmine[100000] //just for throwing an error
    }
    if(amount_counter > 1){
    print("You put in the step number: "+step_check +" ,times:  "+amount_counter)
     metalmine[100000] //just for throwing an error
    }
       
       
}

next_step = true
//___________________________________________________________________________________________________________________________________
for{
    Sleep(500*step)
    //Do next step?
    buildingID, buildingCountdown, researchID, researchCountdown = celestial.ConstructionsBeingBuilt()
    productionLine = GetProduction(celestial.ID)[0]
    if(buildingCountdown == 0 && len(productionLine) == 0 && next_step == false){
        if(include_research == true &&  researchCountdown == 0){
        step = step+1
        next_step = true
        }
        if(include_research == false){
        step = step+1
        next_step = true
        }
     
    }
    
    if(next_step == true){
     //Check Solarplant
    for i = 0; i < len(solarplant); i++ {
        if(solarplant[i] == step){
            celestial.Build(SOLARPLANT, 0)
            buildingID, buildingCountdown, researchID, researchCountdown = celestial.ConstructionsBeingBuilt()
            if(buildingID == SOLARPLANT){
            next_step = false
            print("Solarplant was built at step: "+step)
            }
        
        }
    }
    //Check Metallmine
    for i = 0; i < len(metalmine); i++ {
        if(metalmine[i] == step){
            celestial.Build(METALMINE, 0)
            buildingID, buildingCountdown, researchID, researchCountdown = celestial.ConstructionsBeingBuilt()
            if(buildingID == METALMINE){
             next_step = false
            print("Metalmine was built at step: "+step)
            }
        }
    }
      //Check Crystalmine
    for i = 0; i < len(crystalmine); i++ {
        if(crystalmine[i] == step){
            celestial.Build(CRYSTALMINE, 0)
             buildingID, buildingCountdown, researchID, researchCountdown = celestial.ConstructionsBeingBuilt()
            if(buildingID == CRYSTALMINE){
             next_step = false
            print("CRYSTALMINE was built at step: "+step)
            }
        }
    }
      //Check Deutsynth
    for i = 0; i < len(deutsynth); i++ {
        if(deutsynth[i] == step){
            celestial.Build(DEUTERIUMSYNTHESIZER, 0)
             buildingID, buildingCountdown, researchID, researchCountdown = celestial.ConstructionsBeingBuilt()
            if(buildingID == DEUTERIUMSYNTHESIZER){
             next_step = false
            print("DEUTERIUMSYNTHESIZER was built at step: "+step)
            }
        }
    }
    //Check RoboFactory
    if(robofactory != []){
        for i = 0; i < len(robofactory); i++ {
            if(robofactory[i] == step){
                celestial.Build(ROBOTICSFACTORY, 0)
                  buildingID, buildingCountdown, researchID, researchCountdown = celestial.ConstructionsBeingBuilt()
                    if(buildingID == ROBOTICSFACTORY){
                     next_step = false
                    print("ROBOTICSFACTORY was built at step: "+step)
                    }
            }
        }
    }
    //Check Shipyard
     if(shipyard != []){
        for i = 0; i < len(shipyard); i++ {
            if(shipyard[i] == step){
                celestial.Build(SHIPYARD, 0)
                  buildingID, buildingCountdown, researchID, researchCountdown = celestial.ConstructionsBeingBuilt()
                if(buildingID == SHIPYARD){
                 next_step = false
                print("SHIPYARD was built at step: "+step)
                }
            }
        }
     }
    //Check Lab
     if(lab != []){
        for i = 0; i < len(lab); i++ {
            if(lab[i] == step){
                celestial.Build(RESEARCHLAB, 0)
                 buildingID, buildingCountdown, researchID, researchCountdown = celestial.ConstructionsBeingBuilt()
                if(buildingID == RESEARCHLAB){
                 next_step = false
                print("RESEARCHLABwas built at step: "+step)
                }
            }
        }
     }
    //Energy
     if(energy != []){
        for i = 0; i < len(energy); i++ {
            if(energy[i] == step){
                celestial.Build(ENERGYTECHNOLOGY, 0)
                  buildingID, buildingCountdown, researchID, researchCountdown = celestial.ConstructionsBeingBuilt()
                if(researchID == ENERGYTECHNOLOGY){
                 next_step = false
                print("ENERGYTECHNOLOGY was built at step: "+step)
                }
            }
        }
     }
    //Combustion
     if(combustion != []){
        for i = 0; i < len(combustion); i++ {
            if(combustion[i] == step){
                celestial.Build(COMBUSTIONDRIVE, 0)
                  buildingID, buildingCountdown, researchID, researchCountdown = celestial.ConstructionsBeingBuilt()
                if(researchID == COMBUSTIONDRIVE){
                 next_step = false
                print("COMBUSTIONDRIVEwas built at step: "+step)
                }
            }
        }
     }
    //Computer
       if(computer != []){
        for i = 0; i < len(computer); i++ {
            if(computer[i] == step){
                celestial.Build(COMPUTERTECHNOLOGY, 0)
                 buildingID, buildingCountdown, researchID, researchCountdown = celestial.ConstructionsBeingBuilt()
                if(researchID == COMPUTERTECHNOLOGY){
                 next_step = false
                print("COMPUTERTECHNOLOGY was built at step: "+step)
                }
            }
        }
       }
    //Laser
       if(laser != []){
        for i = 0; i < len(laser); i++ {
            if(laser[i] == step){
                celestial.Build(LASERTECHNOLOGY, 0)
                 buildingID, buildingCountdown, researchID, researchCountdown = celestial.ConstructionsBeingBuilt()
                if(researchID == LASERTECHNOLOGY){
                 next_step = false
                print("LASERTECHNOLOGYwas built at step: "+step)
                }
            }
        }
       }
       
    //Small Cargo
       if(smallcargo != []){
        for i = 0; i < len(smallcargo); i++ {
            if(smallcargo[i] == step){
                    celestial.Build(SMALLCARGO, 0)
                    productionLine = GetProduction(celestial.ID)[0]
                if(productionLine !=0){
                    next_step = false
                   print("SMALLCARGO was built at step: "+step)
                }
            }
        }
       }
         //Eprobes
       if(eprobes != []){
        for i = 0; i < len(eprobes); i++ {
            if(eprobes[i] == step){
                    celestial.Build(ESPIONAGEPROBE, 0)
                    productionLine = GetProduction(celestial.ID)[0]
                if(productionLine !=0){
                    next_step = false
                   print("ESPIONAGEPROBE was built at step: "+step)
                }
            }
        }
       }
         //Espytech
       if(espionage != []){
        for i = 0; i < len(espionage); i++ {
            if(espionage[i] == step){
                celestial.Build(ESPIONAGETECHNOLOGY, 0)
                buildingID, buildingCountdown, researchID, researchCountdown = celestial.ConstructionsBeingBuilt()
                if(researchID == ESPIONAGETECHNOLOGY){
                 next_step = false
                print("ESPIONAGETECHNOLOGY was built at step: "+step)
                }
            }
        }
       }
           //Impuls
       if(impuls != []){
        for i = 0; i < len(impuls); i++ {
            if(impuls[i] == step){
                celestial.Build(IMPULSEDRIVE, 0)
                buildingID, buildingCountdown, researchID, researchCountdown = celestial.ConstructionsBeingBuilt()
                if(researchID == IMPULSEDRIVE){
                 next_step = false
                print("IMPULSEDRIVE was built at step: "+step)
                }
            }
        }
       }
             //Astro
       if(astro != []){
        for i = 0; i < len(astro); i++ {
            if(astro[i] == step){
                celestial.Build(ASTROPHYSICS, 0)
                buildingID, buildingCountdown, researchID, researchCountdown = celestial.ConstructionsBeingBuilt()
                if(researchID == ASTROPHYSICS){
                 next_step = false
                print("ASTROPHYSICS was built at step: "+step)
                }
            }
        }
       }
                //Coloship
       if(coloship != []){
        for i = 0; i < len(coloship); i++ {
            if(coloship[i] == step){
                celestial.Build(COLONYSHIP, 0)
                productionLine = GetProduction(celestial.ID)[0]
                if(productionLine !=0){
                   next_step = false
                   print("COLONYSHIP was built at step: "+step)
                }
            }
        }
       }
    }
}
