package french

import(
	//"../"
)

type Suit int
const(
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
}

const Back int = 0x1F0A0

var AceOfSpades     = Card{ 0x1F0A1, Spades,    1, "Ace of Spades" }
var	DeuceOfSpades   = Card{ 0x1F0A2, Spades,    2, "Two of Spades" }
var	ThreeOfSpades   = Card{ 0x1F0A3, Spades,    3, "Three of Spades" }
var	FourOfSpades    = Card{ 0x1F0A4, Spades,    4, "Four of Spades" }
var	FiveOfSpades    = Card{ 0x1F0A5, Spades,    5, "Five of Spades" }
var	SixOfSpades     = Card{ 0x1F0A6, Spades,    6, "Six of Spades" }
var	SevenOfSpades   = Card{ 0x1F0A7, Spades,    7, "Seven of Spades" }
var	EightOfSpades   = Card{ 0x1F0A8, Spades,    8, "Eight of Spades" }
var	NineOfSpades    = Card{ 0x1F0A9, Spades,    9, "Nine of Spades" }
var	TenOfSpades     = Card{ 0x1F0AA, Spades,   10, "Ten of Spades" }
var	JackOfSpades    = Card{ 0x1F0AB, Spades,   11, "Jack of Spades" }
var	QueenOfSpades   = Card{ 0x1F0AD, Spades,   12, "Queen of Spades" }
var	KingOfSpades    = Card{ 0x1F0AE, Spades,   13, "King of Spades" }

var AceOfHearts     = Card{ 0x1F0B1, Hearts,    1, "Ace of Hearts" }
var	DeuceOfHearts   = Card{ 0x1F0B2, Hearts,    2, "Two of Hearts" }
var	ThreeOfHearts   = Card{ 0x1F0B3, Hearts,    3, "Three of Hearts" }
var	FourOfHearts    = Card{ 0x1F0B4, Hearts,    4, "Four of Hearts" }
var	FiveOfHearts    = Card{ 0x1F0B5, Hearts,    5, "Five of Hearts" }
var	SixOfHearts     = Card{ 0x1F0B6, Hearts,    6, "Six of Hearts" }
var	SevenOfHearts   = Card{ 0x1F0B7, Hearts,    7, "Seven of Hearts" }
var	EightOfHearts   = Card{ 0x1F0B8, Hearts,    8, "Eight of Hearts" }
var	NineOfHearts    = Card{ 0x1F0B9, Hearts,    9, "Nine of Hearts" }
var	TenOfHearts     = Card{ 0x1F0BA, Hearts,   10, "Ten of Hearts" }
var	JackOfHearts    = Card{ 0x1F0BB, Hearts,   11, "Jack of Hearts" }
var	QueenOfHearts   = Card{ 0x1F0BD, Hearts,   12, "Queen of Hearts" }
var	KingOfHearts    = Card{ 0x1F0BE, Hearts,   13, "King of Hearts" }

var AceOfDiamonds   = Card{ 0x1F0C1, Diamonds,  1, "Ace of Diamonds" }
var	DeuceOfDiamonds = Card{ 0x1F0C2, Diamonds,  2, "Two of Diamonds" }
var	ThreeOfDiamonds = Card{ 0x1F0C3, Diamonds,  3, "Three of Diamonds" }
var	FourOfDiamonds  = Card{ 0x1F0C4, Diamonds,  4, "Four of Diamonds" }
var	FiveOfDiamonds  = Card{ 0x1F0C5, Diamonds,  5, "Five of Diamonds" }
var	SixOfDiamonds   = Card{ 0x1F0C6, Diamonds,  6, "Six of Diamonds" }
var	SevenOfDiamonds = Card{ 0x1F0C7, Diamonds,  7, "Seven of Diamonds" }
var	EightOfDiamonds = Card{ 0x1F0C8, Diamonds,  8, "Eight of Diamonds" }
var	NineOfDiamonds  = Card{ 0x1F0C9, Diamonds,  9, "Nine of Diamonds" }
var	TenOfDiamonds   = Card{ 0x1F0CA, Diamonds, 10, "Ten of Diamonds" }
var	JackOfDiamonds  = Card{ 0x1F0CB, Diamonds, 11, "Jack of Diamonds" }
var	QueenOfDiamonds = Card{ 0x1F0CD, Diamonds, 12, "Queen of Diamonds" }
var	KingOfDiamonds  = Card{ 0x1F0CE, Diamonds, 13, "King of Diamonds" }

var AceOfClubs      = Card{ 0x1F0D1, Clubs,     1, "Ace of Clubs" }
var	DeuceOfClubs    = Card{ 0x1F0D2, Clubs,     2, "Two of Clubs" }
var	ThreeOfClubs    = Card{ 0x1F0D3, Clubs,     3, "Three of Clubs" }
var	FourOfClubs     = Card{ 0x1F0D4, Clubs,     4, "Four of Clubs" }
var	FiveOfClubs     = Card{ 0x1F0D5, Clubs,     5, "Five of Clubs" }
var	SixOfClubs      = Card{ 0x1F0D6, Clubs,     6, "Six of Clubs" }
var	SevenOfClubs    = Card{ 0x1F0D7, Clubs,     7, "Seven of Clubs" }
var	EightOfClubs    = Card{ 0x1F0D8, Clubs,     8, "Eight of Clubs" }
var	NineOfClubs     = Card{ 0x1F0D9, Clubs,     9, "Nine of Clubs" }
var	TenOfClubs      = Card{ 0x1F0DA, Clubs,    10, "Ten of Clubs" }
var	JackOfClubs     = Card{ 0x1F0DB, Clubs,    11, "Jack of Clubs" }
var	QueenOfClubs    = Card{ 0x1F0DD, Clubs,    12, "Queen of Clubs" }
var	KingOfClubs     = Card{ 0x1F0DE, Clubs,    13, "King of Clubs" }

var BlackJoker      = Card{ 0x1F0CF, Naked,    15, "Black Joker" }
var WhiteJoker      = Card{ 0x1F0DF, Naked,    15, "White Joker" }

func Spades() []interface{} {
	return []interface{}{
		AceOfSpades, DeuceOfSpades, ThreeOfSpades, FourOfSpades, FiveOfSpades, 
		SixOfSpades, SevenOfSpades, EightOfSpades, NineOfSpades, TenOfSpades,
		JackOfSpades, QueenOfSpades, KingOfSpades,
	}
}

func Hearts() []interface{} {
	return []interface{}{
		AceOfHearts, DeuceOfHearts, ThreeOfHearts, FourOfHearts, FiveOfHearts,
		SixOfHearts, SevenOfHearts, EightOfHearts, NineOfHearts, TenOfHearts,
		JackOfHearts, QueenOfHearts, KingOfHearts,
	}
}

func Diamonds() []interface{} {
	return []interface{}{
		AceOfDiamonds, DeuceOfDiamonds, ThreeOfDiamonds, FourOfDiamonds, FiveOfDiamonds,
		SixOfDiamonds, SevenOfDiamonds, EightOfDiamonds, NineOfDiamonds, TenOfDiamonds,
		JackOfDiamonds, QueenOfDiamonds, KingOfDiamonds,
	}
}

func Clubs() []interface{} {
	return []interface{}{
		AceOfClubs, DeuceOfClubs, ThreeOfClubs, FourOfClubs, FiveOfClubs,
		SixOfClubs, SevenOfClubs, EightOfClubs, NineOfClubs, TenOfClubs,
		JackOfClubs, QueenOfClubs, KingOfClubs,
	}
}

func Decks(count int, jokers int) []interface{} {
	out := make([]interface{}, 0)
	for i := 0; i < count; i++ {
		out = append(out, Spades()[:], Hearts()[:], Diamonds()[:], Clubs()[:])
	}
	for i := 0; i < jokers; i++ {
		if i % 2 == 0 {
			out = append(out, &BlackJoker)
		} else {
			out = append(out, &WhiteJoker)
		}
	}
	return out
}
