//////////////////////////////////////
// Change these settings to your own.
// If you do not want the bot to log you out remove the Logout()
// The bot will turn off any sleep mode and log back in before the fleet lands to be safe.
// To set these to planets, simply remove the M: part.
sendFrom = "M:1:2:3" // Sending fleet from here.
sendTo = "M:1:2:4" // Sending fleet to here.
deutToLeave = 1500000 // Enter how much deut to leave behind when fleetsaving.

// Enter your Telegram Chat ID.
TelegramID = TELEGRAM_CHAT_ID

/* Sends at the desired fleet speed, if 20% fleet speed results in a 8 hour flight time one way then - 
that's how long the total flight time will be. */
fleetSpeed = TEN_PERCENT // TWENTY_PERCENT ... HUNDRED_PERCENT (Set to whichever speed you desire)
//////////////////////////////////////

// Calculate Deut we have from where we are fleetsaving from.
func enoughDeutCheck() {
    celestial = GetCachedCelestial(sendFrom)
    resources, err = GetResourcesDetails(sendFrom)
    return resources.Deuterium.Available
}

// Calculate deut to take.
celestial = GetCachedCelestial(sendFrom)
resources, err = GetResourcesDetails(celestial.GetID())
deutToTake = resources.Deuterium.Available - deutToLeave

// Variables for Telegram usage.
universeName = GetUniverseName()
playerName = GetCachedPlayer().PlayerName
uniPlayerName = universeName + " - " + playerName

// Creates a new fleet object for fleetsaving.
mainFleet = NewFleet()
mainFleet.SetOrigin(sendFrom)
mainFleet.SetDestination(sendTo)
mainFleet.SetSpeed(fleetSpeed)
mainFleet.SetMission(PARK)
mainFleet.SetAllShips()
_, fuel = mainFleet.FlightTime()
mainFleet.SetAllMetal()
mainFleet.SetAllCrystal()
mainFleet.SetDeuterium(deutToTake - fuel)
enoughDeutForFlight = enoughDeutCheck()
fleet, err = mainFleet.SendNow()

// If we do not have enough Deut the fleet cannot be sent, warning message is sent to Telegram.
if enoughDeutForFlight < fuel {
    SendTelegram(TelegramID, uniPlayerName + " WARNING: Not enough Deut to fleetsave, please check and try again!")
    return
}

// Calculates half the arrival time in order to recall the deploy half-way through the flight.
half = fleet.ArriveIn / 2 
Print("Fleetsaving for: ", ShortDur((half * 2)))

// Telegram Message for Fleetsaving.
SendTelegram(TelegramID, uniPlayerName + " Fleetsaving for: " + ShortDur((half*2)))

// Logs out the bot.
Logout()

// Recalls the deploy half-way through the flight with slight randomisation
Sleep(Random(half * 980, half * 1010))

// Stops Sleep mode.
StopSleepMode()

// Logs you in if logged out.
if !IsLoggedIn() {
    Login()
}

// Waits 3-12 seconds, recalls fleet, waits 3-12 seconds then logs out again.
(Sleep(Random(3,12)*1000))
CancelFleet(fleet.ID)
SendTelegram(TelegramID, uniPlayerName + " Recalled fleet")
(Sleep(Random(3,12)*1000))
StartSleepMode()
Logout()

// Logs you back in before your fleet lands, that way if defender is active your fleet is safe when it lands.
Sleep(Random(half * 800, half * 900))
StopSleepMode()
Login()
SendTelegram(TelegramID, uniPlayerName + " Fleet arriving soon")
