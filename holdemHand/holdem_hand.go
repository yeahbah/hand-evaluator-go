package holdemHand

import (
	"errors"
	_ "fmt"
	"strings"
)

// This function takes a string representing a full or partial holdem mask
// and validates that the text represents valid cards and that no card is duplicated.
// Valid hand: 7c 2d. This is ok too: 2c 3d 4s 5c
// This is invalid: 72
func ValidateHand(hand string) bool {
	if hand == "" {
		return false
	}

	var index int = 0
	var handMask uint64 = 0
	var cards = 0
	var card = 0

	for card = nextCard(hand, &index); card >= 0; card = nextCard(hand, &index) {
		if (handMask & (uint64(1) << card)) != 0 {
			return false
		}
		handMask |= uint64(1) << card
		cards++
	}

	return card == -1 && cards > 0 && index >= len(hand)
}

// This function is provided for convenience. It does the same as ValidateHand() but 
// this one takes in a community board parameter.
// Ok: As 2s, board: Ac 2d 9s
func ValidateHandWithBoard(pocket string, board string) bool {
	return ValidateHand(pocket + board)
}

// Parse the hand string to get a uint64 representation
func ParseHand(hand string) (uint64, error) {
	cards := 0
	return parseHand(hand, &cards)
}

// Provided for convenience. It does the same thing as ParseHand() except it accepts a board parameter.
func ParseHandWithBoard(pocket string, board string) (uint64, error) {
	cards := 0
	return parseHand(pocket + board, &cards)
}

// Get card value of given card string
func ParseCard(card string) int {
	cards := 0
	return nextCard(card, &cards) 
}

// given a card value, return the card rank
func CardRank(card int) (int, error) {
	if card < 0 && card > 52 {
		return -1, errors.New("Invalid card. There are only 52 cards in a deck.")
	}

	return card % 13, nil
}

// given a card value, return the card suit
func CardSuit(card int) (int, error) {
	if card < 0 || card > 52 {
		return -1, errors.New("Invalid card")
	}

	return card / 13, nil
}

func parseHand(hand string, cards *int) (uint64, error) {	
	if strings.Trim(hand, " ") == "" {
		*cards = 0
		return 0, nil
	}

	if !ValidateHand(hand) {
		return 0, errors.New("Bad hand definition") 
	}

	*cards = 0
	index := 0
	handMask := uint64(0)
	for card := nextCard(hand, &index); card >= 0; card = nextCard(hand, &index) {
		handMask |= uint64(1) << card
		*cards++
	}

	return handMask, nil
}

func nextCard(cards string, index *int) int {
	if cards == "" {
		errors.New("cards must be defined.")
	}

	for *index < len(cards) && cards[*index] == ' ' {
		*index++
	}

	if *index >= len(cards) {
		return -1
	}

	rank := 0
	card := cards[*index]

	switch card {
	case '1':
		*index += 1
		if cards[*index] == '0' {			
			rank = RankTen
		} else {
			return -1
		}

	case '2':
		rank = Rank2
	case '3':
		rank = Rank3
	case '4':
		rank = Rank4
	case '5':
		rank = Rank5
	case '6':
		rank = Rank6
	case '7':
		rank = Rank7
	case '8':
		rank = Rank8
	case '9':
		rank = Rank9
	case 'T', 't':
		rank = RankTen
	case 'J', 'j':
		rank = RankJack
	case 'Q', 'q':
		rank = RankQueen
	case 'K', 'k':
		rank = RankKing
	case 'A', 'a':
		rank = RankAce
	default:
		return -2
	}

	*index++

	if *index >= len(cards) {
		return -2
	}

	suit := 0
	card = cards[*index]

	switch card {
	case 'H', 'h':
		suit = Hearts
	case 'D', 'd':
		suit = Diamonds
	case 'S', 's':
		suit = Spades
	case 'C', 'c':
		suit = Clubs
	default:
		return -2
	}
	*index++
	return rank + (suit * 13)
}
