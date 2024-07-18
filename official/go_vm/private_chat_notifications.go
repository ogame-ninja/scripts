package main

import (
	"nja/pkg/nja"
)

func main() {
    for {
        msg := <-nja.OnChatMessageReceivedCh
        nja.SendTelegram(nja.TELEGRAM_CHAT_ID, msg.SenderName+": "+msg.Text)
        nja.LogWarn(msg.SenderName+": "+msg.Text)
    }
}
