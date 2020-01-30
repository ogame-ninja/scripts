/* Created by Bull
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
func AucDo(ress){
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
    LogInfo("Wait " + ShortDur(sleepTime))
    Sleep(sleepTime * 1000)
}

func doWork() {
    for {
        auc, err = GetAuction()
        if err != nil {
            LogDebug(err)
            customSleep(Random(5, 10))
            continue
        }
        if auc.HasFinished {
            LogInfo("Auction has finished")
            customSleep(auc.Endtime + 10)
            continue
        }
        if auc.HighestBidderUserID == ownPlayerID {
            LogInfo("You are the highest bidder!")
            customSleep(refreshTime(auc.Endtime))
            continue
        }
        if auc.MinimumBid > highestBid {
            LogInfo("Resources exceeded! Wait until the next auction!")
            customSleep(auc.Endtime + 10)
            continue
        }

        ress = auc.MinimumBid - auc.AlreadyBid
        LogInfo("You are not the highest bidder! Bid " + Dotify(ress) + " resources!")
        err = AucDo(ress)
        if err != nil {
            LogDebug(err)
        }
        customSleep(refreshTime(auc.Endtime))
    }
}
doWork()
