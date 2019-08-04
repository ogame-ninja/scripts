for {
    msg = <-OnChatMessageReceivedCh
    SendTelegram(TELEGRAM_CHAT_ID, msg.SenderName+": "+msg.Text)
    LogWarn(msg.SenderName+": "+msg.Text)
}