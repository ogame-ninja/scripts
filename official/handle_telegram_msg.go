strings = import("strings")

func handleOGameMessage(msg) {
    Print("chat: ", msg.Text)
    err = SendTelegram(TELEGRAM_CHAT_ID, "Bot: " + BotID + ", PlayerID: " + msg.SenderID + ", " + msg.Text)
    if err != nil {
        LogError("err: ", err)
    }
}

func handleTelegramMessage(msg) {
    parts = strings.Split(msg.Text, " ")
    if len(parts) >= 2 {
        targetBotID = parts[0]
        cmd         = parts[1]
        if targetBotID != BotID {
            return // Return early, cmd is not for this bot
        }
        switch cmd {
        case "msg":
            if len(parts) < 4 {
                LogError("Invalid number of arguments for msg command")
                return
            }
            playerID = Atoi(parts[2])
            if playerID == 0 {
                LogError("PlayerID argument must be an integer")
                return
            }
            msgToSend = strings.Join(parts[3:], " ")
            err = SendMessage(playerID, msgToSend)
            if err != nil {
                LogError(err)
                return
            }
            Print("Message was sent")
        default:
            LogError("Unknown cmd: ", cmd)
        }
    } else {
        Print("Receved: ", msg)
    }
}

for {
    select {
    case msg = <-OnTelegramMessageReceivedCh:
        handleTelegramMessage(msg)
    case msg = <-OnChatMessageReceivedCh:
        handleOGameMessage(msg)
    }
}
