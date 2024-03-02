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
	hand := "As Ac"

	// when simple if statement is too easy too read
	Assert(t, func() bool { return !ValidateHand(hand) }, fmt.Sprintf("%s is a valid hand", hand))

}

func TestSadValidateHand(t *testing.T) {
	hand := "Ts Ts"
	if ValidateHand(hand) {
		t.Fatalf("Expecting ValidateHand to return false")
	}
}
