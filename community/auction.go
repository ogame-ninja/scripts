/* Created by Bull and Notriv
 * It is not perfect. However, a quick start for many
 */
 
/*
 * VERSION 1.00
 */
 

/* DESCRIPTION
 * Automatically bid on auctions
 * Maximum resources adjustable
 * Which resource bidding is adjustable
 */

/*---------------------------------------------------------------------------------------------------------------------------*/



//######################################## SETTINGS START ########################################

highestBid = 5000000                //what is the maximum bid
metBid = true                       //should you bid with metal?
crysBid = false                     //should you bid with crystal?
deutBid = false                     //should you bid with deuterium?
bidHome = "M:1:234:5"               //from which planet should be bid?

//######################################## SETTINGS END ########################################


ownPlayerID = GetCachedPlayer().PlayerID
celt = GetCachedCelestial(bidHome)
if celt == nil {
    LogError(bidHome + " is not one of your planet/moon")
    return
}

func AucDo(ress) {
    bid = {}
    if metBid {
        bid = { celt.GetID() : NewResources(ress, 0, 0) }
    } else if crysBid {
        bid = { celt.GetID() : NewResources(0, ress, 0) }
    } else {
        bid = { celt.GetID() : NewResources(0, 0, ress) }
    }
    return DoAuction(bid)
}

func refreshTime(TimeEnd) {
    switch TimeEnd {     
        case TimeEnd <= 300:                    //5 min
        LogDebug("Only 5 min")
        return Random(2, 5)

        case TimeEnd <= 600:                    //10 min
        LogDebug("Only 10 Min")                        
        return Random(60, 120)

        case TimeEnd <= 1800:                   //30 min
        LogDebug("Only 30 Min")                        
        return Random(180, 300)

        case TimeEnd <= 3600:                   //60 min
        LogDebug("Only 60 Min")                        
        return Random(300, 600)

        default:
        LogError("Unknown TimeEnd value", TimeEnd)
        return Random(5, 10)
    }
}

func customSleep(sleepTime) {
    if sleepTime <= 0 {
        sleepTime = Random(5, 10)
    }
    LogInfo("Wait " + ShortDur(sleepTime))
    Sleep(sleepTime * 1000)
}

func didWon(auc) {
    if auc.HighestBidderUserID == ownPlayerID {
        LogInfo("You won the auction with " + Dotify(auc.CurrentBid) + " resources!")
    }
}

func processAuction() {
    auc, err = GetAuction()
    if err != nil {
        LogError(err)
        return Random(5, 10)
    }
    if auc.HasFinished {
        if auc.Endtime > 7200 {
            LogInfo("There is currently no auction")
        } else {
            LogInfo("Auction has finished")
        }
        didWon(auc)
        return auc.Endtime + 10
    }
    if auc.HighestBidderUserID == ownPlayerID {
        LogInfo("You are the highest bidder!")
        return refreshTime(auc.Endtime)
    }
    if auc.MinimumBid > highestBid {
        LogInfo("Resources exceeded! Wait until the next auction!")
        return auc.Endtime + 10
    }

    ress = auc.MinimumBid - auc.AlreadyBid
    LogInfo("You are not the highest bidder! Bid " + Dotify(ress) + " resources!")
    err = AucDo(ress)
    if err != nil {
        LogError(err)
        return Random(5, 10)
    }
    return refreshTime(auc.Endtime)
}

func doWork() {
    for { // forever process auctions
        sleepTime = processAuction()
        customSleep(sleepTime)
    }
}
doWork()
