package holdemHand

import (
	"fmt"
	"strings"
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

func TestCardRank(t *testing.T) {
	card := ParseCard("Qs")
	want := RankQueen
	got, _ := CardRank(card)

	if want != got {
		t.Fatalf("Incorrect card value. Want %d, got %d", want, got)
	}
}

func TestCardSuit(t *testing.T) {
	card := ParseCard("4d")
	want := Diamonds
	got, _ := CardSuit(card)
	if want != got {
		t.Fatalf("Incorrect suit. Want %d, got %d", want, got)
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

func TestMaskToString(t *testing.T) {
	hand := "Js Ts"
	mask, _ := ParseHand(hand)
	got := MaskToString(mask)
	if got != hand {
		t.Fatalf("MaskToString output does not match hand. Want %s, Got %s", hand, got)
	}
}

func TestEvaluateTypeHighCard(t *testing.T) {

}

func TestEvaluateType(t *testing.T) {
	pocket := "Ad Kh"
	board := "8c 5s 6c Js 10h"
	mask, _ := ParseHandWithBoard(pocket, board)
	handType := EvaluateType(mask)
	want := HighCard
	if handType != want {
		t.Fatalf("EvaluateType() failed. Want %d but got %d", want, handType)
	}

	pocket = "Ad Kh"
	board = "Ac 5s 6c Js 10h"
	mask, _ = ParseHandWithBoard(pocket, board)
	handType = EvaluateType(mask)
	want = Pair
	if handType != want {
		t.Fatalf("EvaluateType() failed. Want %d but got %d", want, handType)
	}

	pocket = "Ad Kh"
	board = "Ac Ks 6c Js 10h"
	mask, _ = ParseHandWithBoard(pocket, board)
	handType = EvaluateType(mask)
	want = TwoPair
	if handType != want {
		t.Fatalf("EvaluateType() failed. Want %d but got %d", want, handType)
	}

	pocket = "Ad Ah"
	board = "Ac Ks 6c Js 10h"
	mask, _ = ParseHandWithBoard(pocket, board)
	handType = EvaluateType(mask)
	want = Trips
	if handType != want {
		t.Fatalf("EvaluateType() failed. Want %d but got %d", want, handType)
	}

	pocket = "2d 3d"
	board = "4c 5s 6c Ad Ah"
	mask, _ = ParseHandWithBoard(pocket, board)
	handType = EvaluateType(mask)
	want = Straight
	if handType != want {
		t.Fatalf("EvaluateType() failed. Want %d but got %d", want, handType)
	}

	pocket = "Ad Kh"
	board = "2d Kd 6d Jd Th"
	mask, _ = ParseHandWithBoard(pocket, board)
	handType = EvaluateType(mask)
	want = Flush
	if handType != want {
		t.Fatalf("EvaluateType() failed. Want %d but got %d", want, handType)
	}

	pocket = "Ad Ah"
	board = "As Kd 6d 6c Th"
	mask, _ = ParseHandWithBoard(pocket, board)
	handType = EvaluateType(mask)
	want = FullHouse
	if handType != want {
		t.Fatalf("EvaluateType() failed. Want %d but got %d", want, handType)
	}

	pocket = "8d 9d"
	board = "As Kd Jd 7d Td"
	mask, _ = ParseHandWithBoard(pocket, board)
	handType = EvaluateType(mask)
	want = StraightFlush
	if handType != want {
		t.Fatalf("EvaluateType() failed. Want %d but got %d", want, handType)
	}

	pocket = "Kh Ah"
	board = "Jh Qh 8d 6c Th"
	mask, _ = ParseHandWithBoard(pocket, board)
	handType = EvaluateType(mask)
	want = StraightFlush
	if handType != want {
		t.Fatalf("EvaluateType() failed. Want %d but got %d", want, handType)
	}
}

func TestEvaluateMask(t *testing.T) {
	pocket := "Ad Kh"
	board := "8c 5s 6c Js 10h"
	handValue, _ := EvaluateHandText(pocket + board)
	want := "High card"
	got := HandDescriptionFromHandType(handValue)
	if !strings.Contains(got, want) {
		t.Fatalf("EvaluateHandText() failed. Want %s but got %s", want, got)
	}

	pocket = "Ad Kh"
	board = "Ac 5s 6c Js 10h"
	handValue, _ = EvaluateHandText(pocket + board)
	want = "One pair"
	got = HandDescriptionFromHandType(handValue)
	if !strings.Contains(got, want) {
		t.Fatalf("EvaluateHandText() failed. Want %s but got %s", want, got)
	}

	pocket = "Ad Kh"
	board = "Ac Ks 6c Js 10h"
	handValue, _ = EvaluateHandText(pocket + board)
	want = "Two pair"
	got = HandDescriptionFromHandType(handValue)
	if !strings.Contains(got, want) {
		t.Fatalf("EvaluateHandText() failed. Want %s but got %s", want, got)
	}

	pocket = "Ad Ah"
	board = "Ac Ks 6c Js 10h"
	handValue, _ = EvaluateHandText(pocket + board)
	got = HandDescriptionFromHandType(handValue)
	want = "Three of a kind"
	if !strings.Contains(got, want) {
		t.Fatalf("EvaluateHandText() failed. Want %s but got %s", want, got)
	}

	pocket = "2d 3d"
	board = "4c 5s 6c Ad Ah"
	handValue, _ = EvaluateHandText(pocket + board)
	got = HandDescriptionFromHandType(handValue)
	want = "A straight"
	if !strings.Contains(got, want) {
		t.Fatalf("EvaluateHandText() failed. Want %s but got %s", want, got)
	}

	pocket = "Ad Kh"
	board = "2d Kd 6d Jd Th"
	handValue, _ = EvaluateHandText(pocket + board)
	got = HandDescriptionFromHandType(handValue)
	want = "A flush"
	if !strings.Contains(got, want) {
		t.Fatalf("EvaluateHandText() failed. Want %s but got %s", want, got)
	}

	pocket = "Ad Ah"
	board = "As Kd 6d 6c Th"
	handValue, _ = EvaluateHandText(pocket + board)
	got = HandDescriptionFromHandType(handValue)
	want = "A fullhouse"
	if !strings.Contains(got, want) {
		t.Fatalf("EvaluateHandText() failed. Want %s but got %s", want, got)
	}

	pocket = "Ad Ah"
	board = "As Kd Ac 6c Th"
	handValue, _ = EvaluateHandText(pocket + board)
	got = HandDescriptionFromHandType(handValue)
	want = "Four of a kind"
	if !strings.Contains(got, want) {
		t.Fatalf("EvaluateHandText() failed. Want %s but got %s", want, got)
	}

	pocket = "8d 9d"
	board = "As Kd Jd 7d Td"
	handValue, _ = EvaluateHandText(pocket + board)
	got = HandDescriptionFromHandType(handValue)
	want = "A straight flush"
	if !strings.Contains(got, want) {
		t.Fatalf("EvaluateHandText() failed. Want %s but got %s", want, got)
	}

	pocket = "Kh Ah"
	board = "Jh Qh 8d 6c Th"
	handValue, _ = EvaluateHandText(pocket + board)
	got = HandDescriptionFromHandType(handValue)
	want = "A straight flush"
	if !strings.Contains(got, want) {
		t.Fatalf("EvaluateHandText() failed. Want %s but got %s", want, got)
	}
}
