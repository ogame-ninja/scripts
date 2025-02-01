// By @CellMaster

highscore, err = GetHighscore(1, 1, 1)
if err != nil {
    Printf("Error retrieving the initial ranking page: %v", err)
    return
}
Printf("Page %d: %d players captured", 1, len(highscore.Players))
for page = 2; page <= highscore.NbPage; page++ {
    SleepRandMs(100, 200)
    highscore, err = GetHighscore(1, 1, page)
    if err != nil {
        Printf("Error retrieving ranking from page %d: %v", page, err)
        continue
    }
    Printf("Page %d: %d players captured", page, len(highscore.Players))
}
