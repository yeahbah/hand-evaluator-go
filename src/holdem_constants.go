package holdemHand

const (
	HighCard = iota
	Pair
	TwoPair
	Trips
	Straight
	Flush
	FullHouse
	FourOfAKind
	StraightFlush
)

const (
	Clubs = iota
	Diamonds
	Hearts
	Spades
)

const (
	Rank2 = iota
	Rank3
	Rank4
	Rank5
	Rank6
	Rank7
	Rank8
	Rank9
	RankTen
	RankJack
	RankQueen
	RankKing
	RankAce
)

const CardJoker = 52
const NumberOfCards = 52
const NumberOfCardsWithJoker = 53

const (
	HANDTYPE_SHIFT    int    = 24
	TOP_CARD_SHIFT    int    = 16
	TOP_CARD_MASK     uint32 = 0x000F0000
	SECOND_CARD_SHIFT int    = 12
	THIRD_CARD_SHIFT  int    = 8
	FOURTH_CARD_SHIFT int    = 4
	FIFTH_CARD_SHIFT  int    = 0
	FIFTH_CARD_MASK   uint32 = 0x0000000F
	CARD_WIDTH        int    = 4
	CARD_MASK         uint32 = 0x0F
)
