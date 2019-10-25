// Put the bot to sleep at different time during the week.

oneMinute      = 60
oneHour        = 60 * oneMinute
fifteenMinutes = 15 * oneMinute
eightHours     =  8 * oneHour
tenHours       = 10 * oneHour

func newSleepCallback(duration) {
    return func() {
        Sleep(Random(0, fifteenMinutes * 1000)) // Random sleep 0-15min
        Print("Going to sleep for " + ShortDur(period))
        DisableNJA()
        Sleep(duration * 1000)
        EnableNJA()
        Print("Waking up")
    }
}

CronExec("0 0 23 * * 0", newSleepCallback(tenHours))   // Sun - Sleep at: 23h for 10h
CronExec("0 0 21 * * 1", newSleepCallback(eightHours)) // Mon - Sleep at: 21h for  8h
CronExec("0 0 21 * * 2", newSleepCallback(eightHours)) // Tue - Sleep at: 21h for  8h
CronExec("0 0 21 * * 3", newSleepCallback(eightHours)) // Wed - Sleep at: 21h for  8h
CronExec("0 0 21 * * 4", newSleepCallback(eightHours)) // Thu - Sleep at: 21h for  8h
CronExec("0 0 21 * * 5", newSleepCallback(eightHours)) // Fri - Sleep at: 21h for  8h
CronExec("0 0 23 * * 6", newSleepCallback(tenHours))   // Sat - Sleep at: 23h for 10h
<-OnQuitCh
