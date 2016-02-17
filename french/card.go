// Package french provides a common deck used for playing poker and other traditional card games.
package french

import(

)

// Suit defines the standard suits, using ascii hex values for convenience
type Suit int

const (
	Naked Suit = 0
	Spades     = 0x2660
	Hearts     = 0x2661
	Diamonds   = 0x2662
	Clubs      = 0x2663
)

type Card struct {
	Id    int
	Suit  Suit
	Value int
	Label string
	Deck  int
}

// Back is a convenience holder for the ascii card back code
const Back int = 0x1F0A0

var AceOfSpades     = Card{ 0x1F0A1, Spades,    1, "Ace of Spades",     0 }
var	DeuceOfSpades   = Card{ 0x1F0A2, Spades,    2, "Two of Spades",     0 }
var	ThreeOfSpades   = Card{ 0x1F0A3, Spades,    3, "Three of Spades",   0 }
var	FourOfSpades    = Card{ 0x1F0A4, Spades,    4, "Four of Spades",    0 }
var	FiveOfSpades    = Card{ 0x1F0A5, Spades,    5, "Five of Spades",    0 }
var	SixOfSpades     = Card{ 0x1F0A6, Spades,    6, "Six of Spades",     0 }
var	SevenOfSpades   = Card{ 0x1F0A7, Spades,    7, "Seven of Spades",   0 }
var	EightOfSpades   = Card{ 0x1F0A8, Spades,    8, "Eight of Spades",   0 }
var	NineOfSpades    = Card{ 0x1F0A9, Spades,    9, "Nine of Spades",    0 }
var	TenOfSpades     = Card{ 0x1F0AA, Spades,   10, "Ten of Spades",     0 }
var	JackOfSpades    = Card{ 0x1F0AB, Spades,   11, "Jack of Spades",    0 }
var	QueenOfSpades   = Card{ 0x1F0AD, Spades,   12, "Queen of Spades",   0 }
var	KingOfSpades    = Card{ 0x1F0AE, Spades,   13, "King of Spades",    0 }

var AceOfHearts     = Card{ 0x1F0B1, Hearts,    1, "Ace of Hearts",     0 }
var	DeuceOfHearts   = Card{ 0x1F0B2, Hearts,    2, "Two of Hearts",     0 }
var	ThreeOfHearts   = Card{ 0x1F0B3, Hearts,    3, "Three of Hearts",   0 }
var	FourOfHearts    = Card{ 0x1F0B4, Hearts,    4, "Four of Hearts",    0 }
var	FiveOfHearts    = Card{ 0x1F0B5, Hearts,    5, "Five of Hearts",    0 }
var	SixOfHearts     = Card{ 0x1F0B6, Hearts,    6, "Six of Hearts",     0 }
var	SevenOfHearts   = Card{ 0x1F0B7, Hearts,    7, "Seven of Hearts",   0 }
var	EightOfHearts   = Card{ 0x1F0B8, Hearts,    8, "Eight of Hearts",   0 }
var	NineOfHearts    = Card{ 0x1F0B9, Hearts,    9, "Nine of Hearts",    0 }
var	TenOfHearts     = Card{ 0x1F0BA, Hearts,   10, "Ten of Hearts",     0 }
var	JackOfHearts    = Card{ 0x1F0BB, Hearts,   11, "Jack of Hearts",    0 }
var	QueenOfHearts   = Card{ 0x1F0BD, Hearts,   12, "Queen of Hearts",   0 }
var	KingOfHearts    = Card{ 0x1F0BE, Hearts,   13, "King of Hearts",    0 }

var AceOfDiamonds   = Card{ 0x1F0C1, Diamonds,  1, "Ace of Diamonds",   0 }
var	DeuceOfDiamonds = Card{ 0x1F0C2, Diamonds,  2, "Two of Diamonds",   0 }
var	ThreeOfDiamonds = Card{ 0x1F0C3, Diamonds,  3, "Three of Diamonds", 0 }
var	FourOfDiamonds  = Card{ 0x1F0C4, Diamonds,  4, "Four of Diamonds",  0 }
var	FiveOfDiamonds  = Card{ 0x1F0C5, Diamonds,  5, "Five of Diamonds",  0 }
var	SixOfDiamonds   = Card{ 0x1F0C6, Diamonds,  6, "Six of Diamonds",   0 }
var	SevenOfDiamonds = Card{ 0x1F0C7, Diamonds,  7, "Seven of Diamonds", 0 }
var	EightOfDiamonds = Card{ 0x1F0C8, Diamonds,  8, "Eight of Diamonds", 0 }
var	NineOfDiamonds  = Card{ 0x1F0C9, Diamonds,  9, "Nine of Diamonds",  0 }
var	TenOfDiamonds   = Card{ 0x1F0CA, Diamonds, 10, "Ten of Diamonds",   0 }
var	JackOfDiamonds  = Card{ 0x1F0CB, Diamonds, 11, "Jack of Diamonds",  0 }
var	QueenOfDiamonds = Card{ 0x1F0CD, Diamonds, 12, "Queen of Diamonds", 0 }
var	KingOfDiamonds  = Card{ 0x1F0CE, Diamonds, 13, "King of Diamonds",  0 }

var AceOfClubs      = Card{ 0x1F0D1, Clubs,     1, "Ace of Clubs",      0 }
var	DeuceOfClubs    = Card{ 0x1F0D2, Clubs,     2, "Two of Clubs",      0 }
var	ThreeOfClubs    = Card{ 0x1F0D3, Clubs,     3, "Three of Clubs",    0 }
var	FourOfClubs     = Card{ 0x1F0D4, Clubs,     4, "Four of Clubs",     0 }
var	FiveOfClubs     = Card{ 0x1F0D5, Clubs,     5, "Five of Clubs",     0 }
var	SixOfClubs      = Card{ 0x1F0D6, Clubs,     6, "Six of Clubs",      0 }
var	SevenOfClubs    = Card{ 0x1F0D7, Clubs,     7, "Seven of Clubs",    0 }
var	EightOfClubs    = Card{ 0x1F0D8, Clubs,     8, "Eight of Clubs",    0 }
var	NineOfClubs     = Card{ 0x1F0D9, Clubs,     9, "Nine of Clubs",     0 }
var	TenOfClubs      = Card{ 0x1F0DA, Clubs,    10, "Ten of Clubs",      0 }
var	JackOfClubs     = Card{ 0x1F0DB, Clubs,    11, "Jack of Clubs",     0 }
var	QueenOfClubs    = Card{ 0x1F0DD, Clubs,    12, "Queen of Clubs",    0 }
var	KingOfClubs     = Card{ 0x1F0DE, Clubs,    13, "King of Clubs",     0 }

var BlackJoker      = Card{ 0x1F0CF, Naked,    15, "Black Joker",       0 }
var WhiteJoker      = Card{ 0x1F0DF, Naked,    15, "White Joker",       0 }

// NewSpades creates a slice of all spades in a standard deck
func NewSpades() []interface{} {
	return []interface{}{
		AceOfSpades, DeuceOfSpades, ThreeOfSpades, FourOfSpades, FiveOfSpades, 
		SixOfSpades, SevenOfSpades, EightOfSpades, NineOfSpades, TenOfSpades,
		JackOfSpades, QueenOfSpades, KingOfSpades,
	}
}

// NewHearts creates a slice of all hearts in a standard deck
func NewHearts() []interface{} {
	return []interface{}{
		AceOfHearts, DeuceOfHearts, ThreeOfHearts, FourOfHearts, FiveOfHearts,
		SixOfHearts, SevenOfHearts, EightOfHearts, NineOfHearts, TenOfHearts,
		JackOfHearts, QueenOfHearts, KingOfHearts,
	}
}

// NewDiamonds creates a slice of all diamonds in a standard deck
func NewDiamonds() []interface{} {
	return []interface{}{
		AceOfDiamonds, DeuceOfDiamonds, ThreeOfDiamonds, FourOfDiamonds, FiveOfDiamonds,
		SixOfDiamonds, SevenOfDiamonds, EightOfDiamonds, NineOfDiamonds, TenOfDiamonds,
		JackOfDiamonds, QueenOfDiamonds, KingOfDiamonds,
	}
}

// NewClubs creates a slice of all clubs in a standard deck
func NewClubs() []interface{} {
	return []interface{}{
		AceOfClubs, DeuceOfClubs, ThreeOfClubs, FourOfClubs, FiveOfClubs,
		SixOfClubs, SevenOfClubs, EightOfClubs, NineOfClubs, TenOfClubs,
		JackOfClubs, QueenOfClubs, KingOfClubs,
	}
}

func pushToDeck(deck []interface{}, cards []interface{}, id int, offset int) []interface{} {
	for j, card := range cards {
		card.(Card).Deck = id
		deck[offset + j] = card
	}
	return deck
}

// NewDecks builds a new slice with as many decks and jokers requested
func NewDecks(count int, jokers int) []interface{} {
	out := make([]interface{}, (count * 52) + jokers)
	added := 0
	for i := 0; i < count; i++ {
		out = pushToDeck(out, NewSpades(), i, added)
		out = pushToDeck(out, NewHearts(), i, added)
		out = pushToDeck(out, NewDiamonds(), i, added)
		out = pushToDeck(out, NewClubs(), i, added)
		added += 52
	}
	for i := 0; i < jokers; i++ {
		if i % 2 == 0 {
			out[added + i] = BlackJoker
		} else {
			out[added + i] = WhiteJoker
		}
	}
	return out
}
