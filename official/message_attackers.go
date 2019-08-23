// Define a bunch of messages
msgs = [
    "hey online :)",
    "o/ Online.",
    "Online dude :)",
    "hey, online bro :)",
    "o/ Online buddy!",
    "Online buddy!",
    "im here :(",
    "o/",
]

// Dictionary of [ArrivalTimestamp -> AttackEvent]
attacks = {}

// Either or not the attackerID is new
func IsNewAttacker(attackerID) {
    for _, evt in attacks {
        if attackerID == evt.AttackerID {
            return false
        }
    }
    return true
}

// OnNewAttack Send a random message to attacker.
func OnAttack(attackEvent) {
    if IsNewAttacker(attackEvent.AttackerID) {
        msg = msgs[Random(0, len(msgs)-1)]                    // Pick a random message
        SendMessage(attackEvent.AttackerID, msg)              // Send it to attacker
        LogWarn("`"+msg+"` sent to ", attackEvent.AttackerID) // Log the event
    }
    attacks[attackEvent.ArrivalTime.Unix()] = attackEvent
}

func OnAttackDone(attackEvent) {
    delete(attacks, attackEvent.ArrivalTime.Unix())
}

func OnAttackCancelled(attackEvent) {
    delete(attacks, attackEvent.ArrivalTime.Unix())
}

// Infinite loop. Verify if we got new messages in our channels.
LogInfo("Start send message to attackers script")
for {
    select {
    case attackEvent = <-OnAttackCh:
        OnAttack(attackEvent)
    case attackEvent = <-OnAttackDoneCh:
        OnAttackDone(attackEvent)
    case attackEvent = <-OnAttackCancelledCh:
        OnAttackCancelled(attackEvent)
    }
}
