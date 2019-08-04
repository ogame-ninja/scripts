func callback() {
    err = BuyOfferOfTheDay()
    Print("Buy offer of the day", err)
}
CronExec("@00h00", callback) // Execute callback every day at midnight
<-OnQuitCh