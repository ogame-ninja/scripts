// ########################################### VERSION 1.10 ##################################################

/* 
 * Developed by: Bull, Notriv, and LordMike 
 * This script is not perfect but serves as a quick start for many users.
 * 
 * Description:
 * Automatically bid on auctions
 * Maximum resources customizable per type of auctioned item
 */

 // ########################################### VERSION 1.20 ##################################################

/* 
 * Developed by: CellMaster (Based on the original work of Bull, Notriv, and LordMike)
 * This update enhances the auction bidding logic by improving item detection, bid timing, 
 * and filtering mechanisms while adding better logging and tracking features.
 * 
 * Changelog (Version 1.20 - Changes from Version 1.10):
 * 
 * 1️⃣ **Item Name Translation:**
 *    - Implemented `itemNameDictionary`, which maps item names from **Portuguese to English** 
 *      to ensure correct bid recognition.
 *    - The script follows a **Portuguese → English standard**, meaning that **if a user wants 
 *      to adapt it to another language, they should replace the Portuguese terms, not the English keys**.
 * 
 * 2️⃣ **Forbidden Word Filtering:**
 *    - Introduced `forbiddenWords` list to prevent bidding on items containing specific keywords 
 *      (e.g., `"energy"`).
 * 
 * 3️⃣ **Whitelist for Protected Users:**
 *    - Added `whitelist`, preventing the bot from bidding against specified user IDs.
 * 
 * 4️⃣ **Adaptive Auction Timing:**
 *    - Adjusted bid checking frequency based on the **time remaining in the auction**:
 *      - **Bronze items:** 7-10s
 *      - **Silver items:** 5-7s
 *      - **Gold items:** 3-5s
 *      - **Platinum items:** 2-3s
 *    - Ensures **high-value items are prioritized**, while low-value items receive fewer checks.
 * 
 * 5️⃣ **Enhanced Logging & Notifications:**
 *    - Implemented `LogAndNotify()`, a centralized function to log messages and send Discord alerts.
 *    - Added Discord notifications for:
 *      ✅ Winning an auction.
 *      ✅ Skipping an item due to forbidden words.
 *      ✅ Skipping a bid due to a whitelisted user.
 *      ✅ Items with no mapped bid value.
 * 
 * 6️⃣ **Code Optimization & Maintainability:**
 *    - Refactored functions to improve clarity and efficiency:
 *      - `TranslateItemNameToEnglish()`
 *      - `containsForbiddenWord()`
 *      - `isUserWhitelisted()`
 *      - `DetermineMaxBid()`
 *    - Simplified bid evaluation logic for better performance.
 * 
 * 🔹 **Overall Enhancements:**
 * - Smarter auction participation based on **item value and user rules**.
 * - More **efficient and structured** bid management.
 * - Clearer logging and tracking of auction events.
 * 
 * Author's Note:
 * "This update optimizes automation by refining bid control, tracking, and filtering logic, 
 * improving the overall efficiency of auction participation."
 */

 // ########################################### VERSION 1.21 ##################################################

/*
 * Developed by: CellMaster (Based on the original work of Bull, Notriv, and LordMike)
 * This update includes minor fixes, improved dictionary descriptions, and expanded language support.
 * Now, when a whitelisted user is detected, the log will display **both the player's name and their ID**.
 */

 // ########################################### ITEM NAME DICTIONARY ##########################################

 //
// 📝 **Predefined Language Support (Expandable Library):**
// This dictionary **already includes** translations for multiple languages but can be extended as needed.
// The following languages are currently supported:
//
//    🌍 **Existing Translations:**
//    1️⃣ **English** (Reference language, no translation needed)
//    2️⃣ **Portuguese**
//    3️⃣ **Croatian**
//    4️⃣ **Spanish**
//    5️⃣ **French**
//    6️⃣ **German**
//
// 🔍 **How to Add a New Language:**
// - Add the translated item names **in lowercase** as keys in the dictionary.
// - Map them **to the English equivalent** as the value.
//
// ✅ **Example (Adding Croatian with Multiple Variations for a bronze item):**
//    "bronca" : "bronze",  
//    "brončani" : "bronze",
//
// 🛠 **Common Issues & Fixes:**
// ❌ **My language is missing** → Add it following the format above.
// ❌ **Translation is incorrect** → Verify that the item category matches its correct English term.
// ❌ **A single item has multiple variations** → Add **all** variations separately (E.g., Croatian 'bronca' and 'brončani').
//

// ########################################### Dictionary ####################################################

// Dictionary for translating item names (English, Portuguese, Croatian, Spanish, French, German → English)
itemNameDictionary = {
    // English → English
    "bronze" : "bronze",  
    "silver" : "silver",    
    "gold" : "gold",       
    "platinum" : "platinum",

    // Portuguese → English
    "bronze" : "bronze",  
    "prata" : "silver",    
    "ouro" : "gold",       
    "platina" : "platinum",

    // Croatian → English
    "bronca" : "bronze",  
    "brončani" : "bronze",  
    "srebro" : "silver",    
    "srebrni" : "silver",    
    "zlato" : "gold",       
    "zlatni" : "gold",       
    "platina" : "platinum",
    "platinasti" : "platinum",

    // Spanish → English
    "bronce" : "bronze",  
    "plata" : "silver",    
    "oro" : "gold",       
    "platino" : "platinum",

    // French → English
    "bronze" : "bronze",  
    "argent" : "silver",    
    "or" : "gold",       
    "platine" : "platinum",

    // German → English
    "bronze" : "bronze",  
    "silber" : "silver",    
    "gold" : "gold",       
    "platin" : "platinum"
};

// ########################################### SETTINGS START ################################################

// Maximum bid limits for each auctioned item (in resources)
highestBids = {   
    'bronze': 1000000,     // Maximum bid for Bronze items
    'silver': 2500000,    // Maximum bid for Silver items
    'gold': 5000000,      // Maximum bid for Gold items
    'platinum': 10000000  // Maximum bid for Platinum items
};

// Resource selection for bidding (only one should be set to true)
metBid = true    // If true, the bot will place bids using Metal
crysBid = false  // If true, the bot will place bids using Crystal
deutBid = false  // If true, the bot will place bids using Deuterium

// Define the planet or moon from which the bids will be placed
bidHome = "P:1:1:1"  // Coordinates of the celestial object used for bidding

// List of forbidden words (items containing these words will be ignored)
forbiddenWords = ["energy", "energia", "energije", "energía", "énergie", "energie"]  // If an auctioned item contains this word, it will be skipped

// Whitelist of user IDs (if the highest bidder is on this list, the bot will not bid)
// Cehck this link to find out the players ID's https://s{server-code}-{community}.ogame.gameforge.com/api/players.xml (E.g., https://s1-en.ogame.gameforge.com/api/players.xml )

// 🔍 Check this link to find player IDs: https://s{server-code}-{community}.ogame.gameforge.com/api/players.xml 
// 🌍 Example: If you're playing on **server 1 in the English community**, use: https://s1-en.ogame.gameforge.com/api/players.xml 

whitelist = [999999999, 88888888, 77777777]  // List of user IDs that should not be outbid

// Global setting to enable or disable Discord notifications
DISCORD_NOTIFY = true  // If true, auction events will be sent to Discord

// Granular control for Discord notifications (individually enable/disable specific alerts)
NOTIFY_AUCTION_WON_DISCORD = true       // Notify when the bot wins an auction
NOTIFY_FORBIDDEN_WORD_DISCORD = true    // Notify when an item is skipped due to a forbidden word
NOTIFY_WHITELIST_DISCORD = true         // Notify when an item is skipped due to a whitelisted user
NOTIFY_UNMAPPED_ITEM_DISCORD = true     // Notify when an item has no mapped bid value in the dictionary

// ########################################### BID TIME CONTROL ##############################################

/* 
 * AUCTION_TIME_RANGES
 * 
 * This configuration defines the time intervals (in seconds) that the script will use 
 * to check and process bids in auctions based on the time remaining before the auction ends.
 * 
 * The most critical adjustments happen in the **last 5 minutes**, where the time interval 
 * varies depending on the item's category. This ensures that bidding is **more aggressive** 
 * for high-priority items and **less frequent** for low-priority ones:
 * 
 * - **Bronze:** 7 to 10 seconds (Lowest priority, less frequent checks)
 * - **Silver:** 5 to 7 seconds (Moderate priority)
 * - **Gold:** 3 to 5 seconds (High priority)
 * - **Platinum:** 2 to 3 seconds (Highest priority, fastest reaction time)
 * - **Default:** 5 to 10 seconds (Fallback if item category is unknown)
 * 
 * For timeframes greater than 5 minutes, the script follows **fixed** and **less aggressive** 
 * intervals, reducing the number of checks to **optimize resources** and **avoid predictable patterns**.
 * 
 * This approach ensures that the bot **reacts quickly** in critical moments while conserving 
 * resources and avoiding unnecessary activity when there is still plenty of time left.
 */

// Time range settings for auction processing (in seconds)
AUCTION_TIME_RANGES = {
    "5_MINUTES": {
        "bronze": { "min": 7, "max": 10 },       // Custom range for Bronze items
        "silver": { "min": 5, "max": 7 },       // Custom range for Silver items
        "gold": { "min": 3, "max": 5 },         // Custom range for Gold items
        "platinum": { "min": 2, "max": 3 },     // Custom range for Platinum items
        "default": { "min": 5, "max": 10 }       // Default range if category is unknown
    },
    "10_MINUTES": { "min": 60, "max": 120 },  // 301 - 600s (Last 10 min)
    "30_MINUTES": { "min": 180, "max": 300 }, // 601 - 1800s (Last 30 min)
    "60_MINUTES": { "min": 300, "max": 600 }, // 1801 - 3600s (Last 60 min)
    "UNKNOWN": { "min": 5, "max": 10 }        // Fallback
};

// ########################################### SETTINGS END ##################################################

var strings = import("strings")  // Imports the "strings" module, used for string manipulation throughout the script

// ########################################### INITIALIZATION VALIDATION #####################################

ownPlayerID = GetCachedPlayer().PlayerID  // Retrieves the player's unique ID from the game cache

celt = GetCachedCelestial(bidHome)  // Fetches the celestial object (planet or moon) used for bidding

if celt == nil {  
    LogError(bidHome + " is not one of your planets/moons")  // Logs an error if the specified planet/moon does not exist
    return  // Terminates execution to prevent errors in bidding operations
}

// ########################################### LOGGING & NOTIFICATIONS #######################################

// Logs messages and sends Discord notifications if enabled
func LogAndNotify(message, level, notifyFlag) {
    if level == "INFO" {
        LogInfo(message)
    } else if level == "DEBUG" {
        LogDebug(message)
    } else if level == "WARN" {
        LogWarn(message)
    } else if level == "ERROR" {
        LogError(message, "")
    }

    if DISCORD_NOTIFY && notifyFlag {
        err = SendDiscord(DISCORD_WEBHOOK, "[" + level + "] " + message)
        if err != nil {
            LogError("Failed to send Discord notification: " + err, "")
        }
    }
}

//######################################## VALIDATION FUNCTIONS ########################################

// Checks if the item contains a forbidden word
func containsForbiddenWord(name, itemCategory) {
    for word in forbiddenWords {
        if strings.Contains(strings.ToLower(name), strings.ToLower(word)) {
            message = "Item '" + name + "' (" + itemCategory + ") will be ignored due to forbidden word: '" + word + "'";
            LogAndNotify(message, "WARN", NOTIFY_FORBIDDEN_WORD_DISCORD);
            return true;
        }
    }
    return false;
}

// Checks if the highest bidder is on the whitelist
func isUserWhitelisted(HighestBidderUserID, HighestBidder, itemName, itemCategory) {
    for id in whitelist {
        if id == HighestBidderUserID {
            message = "Item '" + itemName + "' (" + itemCategory + ") was ignored because the highest bid is from a whitelisted user: " + HighestBidder + " (ID: " + HighestBidderUserID + ")";
            LogAndNotify(message, "INFO", NOTIFY_WHITELIST_DISCORD);
            return true;
        }
    }
    return false;
}

// If the language is missing or the translation is incorrect, it provides assistance for correction.
func TranslateItemNameToEnglish(name) {
    nameLower = strings.ToLower(name)

    // Otherwise, attempt to translate using the dictionary
    for key, value in itemNameDictionary {
        if strings.Contains(nameLower, key) {
            return value
        }
    }

    return "" // No match found, handled by `DetermineMaxBid()`
}

// ########################################### MAIN FUNCTIONS ################################################

// Determines the maximum bid allowed for an auction item
func DetermineMaxBid(name, auction) {
    itemEnglishName = TranslateItemNameToEnglish(name);
    itemCategory = itemEnglishName != "" ? itemEnglishName + " item" : "Unknown item";

    if itemEnglishName == "" {
        message = "Unable to map '" + name + "' (" + itemCategory + ") to an English name, skipping";
        LogAndNotify(message, "WARN", NOTIFY_UNMAPPED_ITEM_DISCORD);
        return 0, itemCategory;
    }

    if containsForbiddenWord(name, itemCategory) {
        return 0, itemCategory;
    }

    if isUserWhitelisted(auction.HighestBidderUserID, auction.HighestBidder, name, itemCategory) {
        return 0, itemCategory;
    }    

    highestBid = highestBids[itemEnglishName];

    if highestBid != nil {
        LogDebug("Detected '" + name + "' as '" + itemEnglishName + "', highest bid: " + Dotify(highestBid));
        return highestBid, itemCategory;
    }

    message = "Unable to map '" + name + "' (" + itemCategory + ") to a bid value, skipping";
    LogAndNotify(message, "WARN", NOTIFY_UNMAPPED_ITEM_DISCORD);
    return 0, itemCategory;
}

// Places a bid in the auction
func AucDo(ress) {
    LogInfo("Attempting to place bid of " + Dotify(ress) + " resources.")
    bid = {}
    if metBid {
        bid = { celt.GetID() : NewResources(ress, 0, 0) }
    } else if crysBid {
        bid = { celt.GetID() : NewResources(0, ress, 0) }
    } else {
        bid = { celt.GetID() : NewResources(0, 0, ress) }
    }
    result = DoAuction(bid)
    if result != nil {
        LogError("Bid placement failed. Error: " + result)
    } else {
        LogInfo("Bid placed successfully.")
    }
    return result
}

// Processes the current auction and determines the bidding strategy
func processAuction() {
    // Fetch current auction details
    auc, err = GetAuction()
    if err != nil {
        LogError(err)
        return Random(5, 10)
    }
    
    // Check if the auction has finished
    if auc.HasFinished {
        if auc.Endtime > 7200 {
            LogInfo("There is currently no active auction")
        } else {
            LogInfo("Auction has finished")
        }
        // Check if we won the previous auction
        didWon(auc)
        return auc.Endtime + 10
    }
    
    // Determine the highest bid allowed for this item
    highestBid, itemCategory = DetermineMaxBid(auc.CurrentItem, auc)

    // Skip auction if no valid bid amount was determined
    if highestBid <= 0 {
        LogInfo("Skipping auction for '" + auc.CurrentItem + "' (" + itemCategory + ")")
        return auc.Endtime + 10
    }
    
    // Log information if no bid has been placed yet
    if auc.AlreadyBid == 0 {
        LogDebug("Willing to bid " + Dotify(highestBid) + " for '" + auc.CurrentItem + "' (" + itemCategory + ")")
    }
    
    // If we are already the highest bidder, wait for the auction to progress
    if auc.HighestBidderUserID == ownPlayerID {
        LogDebug("Already highest bidder for '" + auc.CurrentItem + "' (" + itemCategory + ") at " + Dotify(auc.CurrentBid) + " / " + Dotify(highestBid) + ", waiting..")
        return refreshTime(auc.Endtime, itemCategory)
    }
    
    // If the required minimum bid exceeds our allowed highest bid, skip the auction
    if auc.MinimumBid > highestBid {
        LogWarn("Resources exceeded for '" + auc.CurrentItem + "' (" + itemCategory + "), currently at " + Dotify(auc.CurrentBid) + " / " + Dotify(highestBid))
        return auc.Endtime + 10
    }
    
    // Calculate the amount to bid
    shouldBid = auc.MinimumBid - auc.AlreadyBid
    LogInfo("Bidding " + Dotify(auc.AlreadyBid) + " + " + Dotify(shouldBid) + " / " + Dotify(highestBid) + " resources for '" + auc.CurrentItem + "' (" + itemCategory + ")")	

    // Attempt to place a bid
    err = AucDo(shouldBid)
    if err != nil {
        LogError(err)
        return Random(5, 10)
    }
    
    // Determine the next wait time based on auction duration
    return refreshTime(auc.Endtime, itemCategory)
}

// Function to determine refresh time based on auction time left
func refreshTime(TimeEnd, itemCategory) {
    if TimeEnd <= 300 {
        LogDebug("Only 5 min left");

        // Determine the range for 5-minute items based on category
        rangeForCategory = AUCTION_TIME_RANGES["5_MINUTES"][itemCategory];
        if rangeForCategory == nil {
            rangeForCategory = AUCTION_TIME_RANGES["5_MINUTES"]["default"];  // Fallback to default
        }
        return Random(rangeForCategory.min, rangeForCategory.max);
    } else if TimeEnd <= 600 {
        LogDebug("Only 10 min left");
        return Random(AUCTION_TIME_RANGES["10_MINUTES"].min, AUCTION_TIME_RANGES["10_MINUTES"].max);
    } else if TimeEnd <= 1800 {
        LogDebug("Only 30 min left");
        return Random(AUCTION_TIME_RANGES["30_MINUTES"].min, AUCTION_TIME_RANGES["30_MINUTES"].max);
    } else if TimeEnd <= 3600 {
        LogDebug("Only 60 min left");
        return Random(AUCTION_TIME_RANGES["60_MINUTES"].min, AUCTION_TIME_RANGES["60_MINUTES"].max);
    } else {
        LogError("Unknown TimeEnd value", TimeEnd);
        return Random(AUCTION_TIME_RANGES["UNKNOWN"].min, AUCTION_TIME_RANGES["UNKNOWN"].max);
    }
}

// Sleeps for a calculated duration before checking the next auction
func customSleep(sleepTime) {
    if sleepTime <= 0 {
        sleepTime = Random(5, 10)
    }
    LogDebug("Sleeping for " + ShortDur(sleepTime) + " before next auction check.")
    Sleep(sleepTime * 1000)
}

// Checks if the player won the auction
func didWon(auc) {
    if auc.HighestBidderUserID == ownPlayerID {
        // Mensagem com nome do item
        message = "You won the auction for '" + auc.CurrentItem + "' with " + Dotify(auc.CurrentBid) + " resources!"
        LogAndNotify(message, "INFO", NOTIFY_AUCTION_WON_DISCORD)
    }
}

// Infinite loop to continuously process auctions
func doWork() {
    for {
        sleepTime = processAuction()
        customSleep(sleepTime)
    }
}
doWork()
