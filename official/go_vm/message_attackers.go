package main

import (
	"nja/pkg/nja"
	"nja/pkg/ogame"
)

// Define a bunch of messages
var msgs = []string{
	"hey online :)",
	"o/ Online.",
	"Online dude :)",
	"hey, online bro :)",
	"o/ Online buddy!",
	"Online buddy!",
	"im here :(",
	"o/",
}

func main() {
	// Dictionary of [ArrivalTimestamp -> AttackEvent]
	attacks := make(map[int64]ogame.AttackEvent)

	// Infinite loop. Verify if we got new messages in our channels.
	nja.LogInfo("Start send message to attackers script")
	for {
		select {
		case attackEvent := <-nja.OnAttackCh:
			OnAttack(attacks, attackEvent)
		case attackEvent := <-nja.OnAttackDoneCh:
			OnAttackDone(attacks, attackEvent)
		case attackEvent := <-nja.OnAttackCancelledCh:
			OnAttackCancelled(attacks, attackEvent)
		}
	}
}

// OnNewAttack Send a random message to attacker.
func OnAttack(attacks map[int64]ogame.AttackEvent, attackEvent ogame.AttackEvent) {
	if IsNewAttacker(attacks, attackEvent.AttackerID) {
		msg := msgs[nja.Random(0, int64(len(msgs)-1))]            // Pick a random message
		nja.SendMessage(attackEvent.AttackerID, msg)              // Send it to attacker
		nja.LogWarn("`"+msg+"` sent to ", attackEvent.AttackerID) // Log the event
	}
	attacks[attackEvent.ArrivalTime.Unix()] = attackEvent
}

func OnAttackDone(attacks map[int64]ogame.AttackEvent, attackEvent ogame.AttackEvent) {
	delete(attacks, attackEvent.ArrivalTime.Unix())
}

func OnAttackCancelled(attacks map[int64]ogame.AttackEvent, attackEvent ogame.AttackEvent) {
	delete(attacks, attackEvent.ArrivalTime.Unix())
}

// Either or not the attackerID is new
func IsNewAttacker(attacks map[int64]ogame.AttackEvent, attackerID int64) bool {
	for _, evt := range attacks {
		if attackerID == evt.AttackerID {
			return false
		}
	}
	return true
}
