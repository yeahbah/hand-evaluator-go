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

func TestParseHand(t *testing.T) {
	hand := "Ks 10s"
	want := uint64(0)

	want |= uint64(1) << (RankKing + (Spades * 13)) 
	want |= uint64(1) << (RankTen + (Spades * 13))	

	var actual, err = ParseHand(hand)
	if err != nil {
		t.Fatalf("Unable to parse %s", hand)
	}

	if actual != want {
		t.Fatalf("Incorrect hand mask value. Want: %d, Got: %d", want, actual)
	}

}

func TestParseHandWithBoard(t *testing.T) {
	hand := "9h Qd"
	board := "5d 8h Js"
	want := uint64(0)

	want |= uint64(1) << (Rank9 + (Hearts * 13)) 
	want |= uint64(1) << (RankQueen + (Diamonds * 13))	
	want |= uint64(1) << (Rank5 + (Diamonds * 13))	
	want |= uint64(1) << (Rank8 + (Hearts * 13))	
	want |= uint64(1) << (RankJack + (Spades * 13))	

	var actual, err = ParseHandWithBoard(hand, board)
	if err != nil {
		t.Fatalf("Unable to parse %s", hand)
	}

	if actual != want {
		t.Fatalf("Incorrect hand mask value. Want: %d, Got: %d", want, actual)
	}
}

func TestParseCard(t *testing.T) {
	card := "6c"
	want := Rank6 + (Clubs * 13)
	got := ParseCard(card)
	if got != want {
		t.Fatalf("Incorrect card value. Want %d, Got: %d", want, got)
	}
}