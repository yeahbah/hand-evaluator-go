package holdemHand

import (
	"errors"
)

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
		if cards[*index] == '0' {
			*index++
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
