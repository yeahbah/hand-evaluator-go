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

// Converts a handvalue into descriptive text
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
		sb.WriteString("'s with a ")
		sb.WriteString(RankTable[kicker])
		sb.WriteString(" for a kicker")
		return sb.String()
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
		sb.WriteString("A fullhouse, ")
		sb.WriteString(RankTable[topCard])
		sb.WriteString("'s and ")

		secondCard := getSecondCard(handValue)
		sb.WriteString(RankTable[secondCard])
		sb.WriteString("'s")
		return sb.String()

	case FourOfAKind:
		topCard := getTopCard(handValue)
		sb.WriteString("Four of a kind, ")
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

// Converts a hand mask to a hand text
func MaskToString(mask uint64) string {
	sb := strings.Builder{}

	count := 0
	for card := range cardsRange(mask) {
		if count > 0 {
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
		if t&0x01 != 0 {
			if t&0x02 != 0 {
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

// Evaluates a long type hand mask and returns a hand value.
// A hand value can be compared against another hand value to
// determine which has the higher value.
func EvaluateMask(mask uint64) (uint, error) {
	numCards := bitCount(mask)
	if numCards < 1 || numCards > 7 {
		return 0, errors.New("Invalid number of cards")
	}

	sc := uint((mask >> CLUB_OFFSET) & 0x1FFF)
	sd := uint((mask >> DIAMOND_OFFSET) & 0x1FFF)
	sh := uint((mask >> HEART_OFFSET) & 0x1FFF)
	ss := uint((mask >> SPADE_OFFSET) & 0x1FFF)

	ranks := sc | sd | sh | ss
	nRanks := BitsTable[ranks]
	numDups := numCards - uint(nRanks)
	result := uint(0)

	var straighOrFlushFunc func(uint) uint = func(suit uint) uint {
		if StraightTable[suit] != 0 {
			return HANDTYPE_VALUE_STRAIGHTFLUSH + uint(StraightTable[suit]<<uint16(TOP_CARD_SHIFT))
		}
		return HANDTYPE_VALUE_FLUSH + TopFiveCardsTable[suit]
	}

	// check for straight, flush or straight flush and return if we
	// determine immediately that this is the best possible mask
	if nRanks >= 5 {
		if BitsTable[ss] >= 5 {
			// if StraightTable[ss] != 0 {
			// 	return HANDTYPE_VALUE_STRAIGHTFLUSH + uint(StraightTable[ss]<<TOP_CARD_SHIFT), nil
			// } else {
			// 	result = HANDTYPE_VALUE_FLUSH + uint(TopFiveCardsTable[ss])
			// }
			result = straighOrFlushFunc(ss)

		} else if BitsTable[sc] >= 5 {
			// if StraightTable[sc] != 0 {
			// 	return HANDTYPE_VALUE_STRAIGHTFLUSH + uint(StraightTable[sc]<<TOP_CARD_SHIFT), nil
			// } else {
			// 	result = HANDTYPE_VALUE_FLUSH + uint(TopFiveCardsTable[sc])
			// }
			result = straighOrFlushFunc(sc)
		} else if BitsTable[sd] >= 5 {
			// if StraightTable[sd] != 0 {
			// 	return HANDTYPE_VALUE_STRAIGHTFLUSH + uint(StraightTable[sd]<<TOP_CARD_SHIFT), nil
			// } else {
			// 	result = HANDTYPE_VALUE_FLUSH + uint(TopFiveCardsTable[sd])
			// }
			result = straighOrFlushFunc(sd)
		} else if BitsTable[sh] >= 5 {
			result = straighOrFlushFunc(sh)
		} else {
			st := uint(StraightTable[ranks])
			if st != 0 {
				result = HANDTYPE_VALUE_STRAIGHT + st<<TOP_CARD_SHIFT
			}
		}

		// Another win -- if there can't be a FH/Quads (numDups < 3),
		// which is true most of the time when there is a made mask, the if we've
		// found a five card mask, just return. This skips the whole process of
		// computing two mask/three_mask/etc
		if result != 0 && numDups < 3 {
			return result, nil
		}
	}

	// by the we're here, either:
	// 1. there's no five-card mask possible (flush or straight), or
	// 2. there's a flush or straight, but we know that there are enough
	//	duplicates to make a full house / quads possible
	switch numDups {
	case 0:
		return HANDTYPE_VALUE_HIGHCARD + TopFiveCardsTable[ranks], nil
	case 1:
		twoMask := ranks ^ (sc ^ sd ^ sh ^ ss)
		result = HANDTYPE_VALUE_PAIR + TopCardTable[twoMask]<<TOP_CARD_SHIFT
		t := ranks ^ twoMask // only one bit set in twoMask

		// get the top five cards in what is left, drop all but the top three
		// cards, and shift them by one to get the three desired kickers
		kickers := (TopFiveCardsTable[t] >> CARD_WIDTH) &^ FIFTH_CARD_MASK
		result += kickers
		return result, nil

	case 2:
		// either two pair or trips
		twoMask := ranks ^ (sc ^ sd ^ sh ^ ss)
		if twoMask != 0 {
			t := ranks ^ twoMask // exactly two bits set in twoMask
			result := HANDTYPE_VALUE_TWOPAIR + (TopFiveCardsTable[twoMask] & (TOP_CARD_MASK | SECOND_CARD_MASK)) + (TopCardTable[t] << THIRD_CARD_SHIFT)
			return result, nil
		}

		threeMask := ((sc & sd) | (sh & ss)) & ((sc & sh) | (sd & ss))
		result := HANDTYPE_VALUE_TRIPS + TopCardTable[threeMask]<<TOP_CARD_SHIFT
		t := ranks & threeMask // only one bit set in the threeMask
		second := TopCardTable[t]
		result += second << SECOND_CARD_SHIFT
		t ^= uint(1) << second
		result += TopCardTable[t] << THIRD_CARD_SHIFT
		return result, nil

	default:
		// possible quads, full house or flush or two pair
		fourMask := sh & sd & sc & ss
		if fourMask != 0 {
			tc := TopCardTable[fourMask]
			result := HANDTYPE_VALUE_FOUR_OF_A_KIND + (tc << TOP_CARD_SHIFT) + ((TopCardTable[ranks^uint(1)<<tc]) << SECOND_CARD_SHIFT)
			return result, nil
		}

		// technically, threeMask as defined below is really the set of bits
		// which are set in three or four of the suits, but since
		// we've already eliminated quads, this is OK.
		// Similarly, twoMask is really twoOrFourMask, but since we're
		// already eliminated quads, we can use this shortcut
		twoMask := ranks ^ (sc ^ sd ^ sh ^ ss)
		if BitsTable[twoMask] != numDups {
			// must be some trips then, which really means there is a
			// full house since numDups >= 3
			threeMask := ((sc ^ sd) | (sh & ss)) & ((sc & sh) | (sd & ss))
			result := HANDTYPE_VALUE_FULLHOUSE
			tc := TopCardTable[threeMask]
			result += tc << TOP_CARD_SHIFT
			t := (twoMask | threeMask) ^ (uint(1) << tc)
			result += TopCardTable[t] << SECOND_CARD_SHIFT
			return result, nil
		}

		// must be two pair
		result := HANDTYPE_VALUE_TWOPAIR
		top := TopCardTable[twoMask]
		result += top << TOP_CARD_SHIFT
		second := TopCardTable[twoMask^uint(1)<<top]
		result += second << SECOND_CARD_SHIFT
		result += TopCardTable[ranks^(uint(1)<<top)^(uint(1)<<second)] << THIRD_CARD_SHIFT
		return result, nil
	}

}

// Evaluates a mask passed as a string and returns a hand value.
// A hand value can be compare against another hand value to
// determine which has the higher value.
func EvaluateHandText(hand string) (uint, error) {
	mask, e := ParseHand(hand)
	if e != nil {
		return 0, errors.New(e.Error())
	}
	return EvaluateMask(mask)
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

func cardsRange(mask uint64) <-chan string {
	channel := make(chan string)
	go func() {
		for i := 51; i >= 0; i-- {
			if (uint64(1)<<i)&mask != 0 {
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
