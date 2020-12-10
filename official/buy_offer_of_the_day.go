func callback() {
    SleepRandMin(3, 5) // Pause execution for a random duration between 3min to 5min
    err = BuyOfferOfTheDay()
    Print("Buy offer of the day", err)
}
CronExec("@00h00", callback) // Execute callback every day at midnight
<-OnQuitCh
