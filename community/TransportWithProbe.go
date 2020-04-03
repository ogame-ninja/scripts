/* 
    Creaded by Anderson Palma
    Version 2020.04.03.A
    Use this script to quickly move resources using Spy Probes (when universe rules allow it)
    
    For updated version or new scripts go to:
    https://github.com/ogame-ninja/scripts/tree/master/community/
    
    if you use this code, you must:
    1 keep it free
    2 keep this notice
*/


//------------------------------------------------------------------------------
// All you have to edit is here
//------------------------------------------------------------------------------


from = "1:111:1"                // planet with resources
to   = "2:222:2"                // planet to receive resources
repeat = 10                       // Number of time to send
maxprobes = true                // "true" to use All of the available probes
fixedprobes = 999              // Use a especific number of probes when maxprobes = false)
waittime = 240000               // Round trip time  in ms, at this moment, you have to insert manually
wantedresource = 2              // Select the resource to transport
//                                 1 - metal
//                                 2 - cristal  
//                                 3 - deuterium 
//                                 4 - all (wont work with negative deuterium production)


//------------------------------------------------------------------------------
// Do not edit below this line if you are not a coder
//------------------------------------------------------------------------------



func TransportWithProbe() {
    probes=0 
    if maxprobes == true {
        celestial = GetCachedCelestial(from)
        ships, _ = GetShips(celestial.GetID())
        Print("Found " + ships.EspionageProbe + " Spy probes ")
        probes = ships.EspionageProbe
    } else {
        probes = fixedprobes 
    }

    for i = 1; i <= repeat; i++ {
        LogDebug("Sending " + probes + " spy Probes " + i + " of " + repeat + " times " )
            
        fleet = NewFleet()
        fleet.SetOrigin(from)
        fleet.SetDestination(to)
        fleet.SetMission(TRANSPORT)
            
        if wantedresource == 1 { fleet.SetAllMetal()        }
        if wantedresource == 2 { fleet.SetAllCrystal()      }
        if wantedresource == 3 { fleet.SetAllDeuterium()    }
        if wantedresource == 4 { fleet.SetAllResources()    } 
            
        fleet.AddShips(ESPIONAGEPROBE, probes)
        fleet, err = fleet.SendNow()
        LogDebug("Probes Fleet.ID/ERROR: "+ fleet.ID + "/" + err +  "\n waiting " + waittime/1000 + " seconds ")
        Sleep(waittime)
            
            
        RND = Random(5000, 10000)
        LogDebug ("Returned, wait random delay of " + RND/1000 + " Seconds ")
        Sleep(RND) // Sleep between 5 and 10 seconds
   
}
LogDebug ("Sent probes "+ repeat + " times" )
LogDebug ("End of Script" )
Exit()
}

TransportWithProbe()
// if you get
// ERROR [TransportWithProbe.ank [43:29] type invalid does not support member operation]
// it means you coordinates are wrong
