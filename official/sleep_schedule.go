// Put the bot to sleep at different time during the week.

oneMinute      = 60
oneHour        = 60 * oneMinute
fifteenMinutes = 15 * oneMinute
eightHours     =  8 * oneHour
tenHours       = 10 * oneHour

func startSleep(period) {
    return func() {
        Sleep(Random(0, fifteenMinutes * 1000)) // Random sleep 0-15min
        Print("Going to sleep for " + ShortDur(period))
        DisableNJA()
        Sleep(period * 1000)
        EnableNJA()
        Print("Waking up")
    }
}

CronExec("0 0 23 * * 0", startSleep(tenHours))   // Sun - Sleep at: 23h for 10h
CronExec("0 0 21 * * 1", startSleep(eightHours)) // Mon - Sleep at: 21h for  8h
CronExec("0 0 21 * * 2", startSleep(eightHours)) // Tue - Sleep at: 21h for  8h
CronExec("0 0 21 * * 3", startSleep(eightHours)) // Wed - Sleep at: 21h for  8h
CronExec("0 0 21 * * 4", startSleep(eightHours)) // Thu - Sleep at: 21h for  8h
CronExec("0 0 21 * * 5", startSleep(eightHours)) // Fri - Sleep at: 21h for  8h
CronExec("0 0 23 * * 6", startSleep(tenHours))   // Sat - Sleep at: 23h for 10h
<-OnQuitCh
