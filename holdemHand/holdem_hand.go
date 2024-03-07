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
	return parseHand(pocket+board, &cards)
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

func HandDescriptionFromHandType(handValue uint) string {
	sb := strings.Builder{}
	handType := getHandType(handValue)

	switch handType {
	case HighCard:
		topCard := getTopCard(handValue)
		sb.WriteString("High card: ")
		sb.WriteString(RankTable[topCard])
		return sb.String()
	
	case Pair: 
		topCard := getTopCard(handValue)
		sb.WriteString("One pair, ")
		sb.WriteString(RankTable[topCard])
		return sb.String()
	
	case TwoPair:
		topCard := getTopCard(handValue)
		sb.WriteString("Two pair, ")
		sb.WriteString(RankTable[topCard])

		secondCard := getSecondCard(handValue)
		sb.WriteString("'s and ")
		sb.WriteString(RankTable[secondCard])
		
		kicker := getThirdCard(handValue)
		sb.WriteString("'s with a")
		sb.WriteString(RankTable[kicker])
		sb.WriteString(" for a kicker")
	case Trips:
		topCard := getTopCard(handValue)
		sb.WriteString("Three of a kind, ")
		sb.WriteString(RankTable[topCard])
		sb.WriteString("'s")
		return sb.String()

	case Straight:
		topCard := getTopCard(handValue)
		sb.WriteString("A straight, ")
		sb.WriteString(RankTable[topCard])
		sb.WriteString(" high")
		return sb.String()

	case Flush:
		topCard := getTopCard(handValue)
		sb.WriteString("A flush, ")
		sb.WriteString(RankTable[topCard])
		sb.WriteString(" high")
		return sb.String()

	case FullHouse:
		topCard := getTopCard(handValue)
		sb.WriteString(RankTable[topCard])
		sb.WriteString("'s and ")
		
		secondCard := getSecondCard(handValue)
		sb.WriteString(RankTable[secondCard])
		sb.WriteString("'s")
		return sb.String()

	case FourOfAKind:
		topCard := getTopCard(handValue)
		sb.WriteString(RankTable[topCard])
		sb.WriteString("'s")
		return sb.String()

	case StraightFlush:
		sb.WriteString("A straight flush")
		return sb.String()

	}

	return ""
}

// func GetCardMask(cards uint64, suit int) uint {

// }

func MaskToString(mask uint64) string {
	sb := strings.Builder{}	

	count := 0
	for card := range cardsRange(mask) {
		if (count > 0) {
			sb.WriteString(" ")
		}		
		sb.WriteString(card)
		count++
	}

	return sb.String()
}

// This function is faster than Evaluate but provides less information
func EvaluateType(mask uint64) int {
	numCards := bitCount(mask)
	result := HighCard

	x := uint64(0x1FFF)
	ss := uint((mask >> SPADE_OFFSET) & x)
	sc := uint((mask >> CLUB_OFFSET) & x)
	sd := uint((mask >> DIAMOND_OFFSET) & x)
	sh := uint((mask >> HEART_OFFSET) & x)

	ranks := sc | sd | sh | ss
	rankInfo := uint(BitsAndStrTable[ranks])
	numDups := uint(numCards - (rankInfo >> 2))

	if (rankInfo & 0x01) != 0 {
		if (rankInfo & 0x02) != 0 {
			result = Straight
		}
		t := uint(BitsAndStrTable[ss] | BitsAndStrTable[sc] | BitsAndStrTable[sd] | BitsAndStrTable[sh])
		if t & 0x01 != 0 {
			if t & 0x02 != 0 {
				return StraightFlush
			} else {
				result = Flush
			}
		}
		if result != 0 && numDups < 3 {
			return result
		}
	}

	switch numDups {
	case 0:
		return HighCard
	case 1:
		return Pair
	case 2:
		if (ranks ^ (sc ^ sd ^ sh ^ ss)) != 0 {
			return TwoPair
		} else {
			return Trips
		}
	default:
		if ((sc & sd) & (sh & ss)) != 0 {
			return FourOfAKind
		} else if (((sc & sd) | (sh & ss)) & ((sc & sh) | (sd & ss))) != 0 {
			return FullHouse
		} else if result != 0 {
			return result
		} else {
			return TwoPair
		}
	}	
}

func EvaluateMask(mask uint64) {

}

// Evaluates a mask passed as a string and returns a hand value.
// A hand value can be compare against another hand value to
// determine which has the higher value.
func EvaluateHandText(hand string) {

}

func bitCount(mask uint64) uint {
	x := uint64(0x1FFF)
	ss := uint((mask >> SPADE_OFFSET) & x)
	sc := uint((mask >> CLUB_OFFSET) & x)
	sd := uint((mask >> DIAMOND_OFFSET) & x)
	sh := uint((mask >> HEART_OFFSET) & x)
 	result := BitsTable[sc] + BitsTable[ss] + BitsTable[sd] + BitsTable[sh]
	return uint(result)
}

func cardsRange(mask uint64) <- chan string {
	channel := make(chan string)
	go func() {
		for i := 51; i >= 0; i-- {
			if (uint64(1) << i) & mask != 0 {				
				channel <- CardTable[i]
			}
		}
		close(channel)
	}()

	return channel	
}

func getHandType(handValue uint) uint {
	return handValue >> HANDTYPE_SHIFT
}

func getTopCard(handValue uint) uint {
	return (handValue >> TOP_CARD_SHIFT) & CARD_MASK
}

func getSecondCard(handValue uint) uint {
	return (handValue >> SECOND_CARD_SHIFT) & CARD_MASK
}

func getThirdCard(handValue uint) uint {
	return (handValue >> THIRD_CARD_SHIFT) & CARD_MASK
}

func getFourthCard(handValue uint) uint {
	return (handValue >> FOURTH_CARD_SHIFT) & CARD_MASK
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
		return -1
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
