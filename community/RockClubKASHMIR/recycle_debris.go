//==== This script is created by RockClubKASHMIR ====

fromSystem = 1 // Your can change this as you wish
toSystem = 499 // Your can change this as you wish
Rnbr = 1 // if Rnbr = 1, the script will search debris for minimum 2 Recyclers. You can change it to any value you want
//----
curSystem = fromSystem
origin = nil
flts = 0
nbr = 0
err = 0
i = 1
slots = GetSlots().InUse+GetFleetSlotsReserved()
for celestial in GetCachedCelestials() {
    ships, _ = celestial.GetShips()
    if ships.Recycler > flts {
        flts = ships.Recycler
        origin = celestial // Your Planet(or Moon), with biggest amount of Recyclers
    }
}
if origin != nil {
    
} else {Print("You don't have any ships from the desired list of ships!")}
