package holdemHand

import (
	"fmt"
	"testing"
)

type AssertFunc func() bool

func Assert(t *testing.T, assertFunc AssertFunc, msg string) {
	if assertFunc() {
		t.Fatalf(msg)
	}
}

func TestHappyValidateHand(t *testing.T) {
	hand := "Ts 10c"

	// when simple if statement is too easy too read
	Assert(t, func() bool { return !ValidateHand(hand) }, fmt.Sprintf("%s is a valid hand", hand))

	hand = "2s 3c 4d 5d"
	Assert(t, func() bool { return !ValidateHand(hand) }, fmt.Sprintf("%s is a valid hand", hand))
}

func TestSadValidateHand(t *testing.T) {
	hand := "Ts Ts"
	if ValidateHand(hand) {
		t.Fatalf("Expecting ValidateHand to return false")
	}
}

func TestValidateHandWithBoard(t *testing.T) {
	hand := "Td 10s"
	board := "10h Jd 3c"
	Assert(t, func() bool {
		return !ValidateHandWithBoard(hand, board)
	}, fmt.Sprintf("%s with board %s is a valid hand and board combination", hand, board))
}
