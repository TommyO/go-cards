package deck

import(
	"fmt"
	"testing"
)

type Card struct { Name string }

var AceOfSpades   Card = Card{ "Ace Of Spades" }
var AceOfHearts   Card = Card{ "Ace Of Hearts" }
var AceOfDiamonds Card = Card{ "Ace Of Diamonds" }
var AceOfClubs    Card = Card{ "Ace Of Clubs" }

var Aces []Card = []Card{ AceOfSpades, AceOfHearts, AceOfDiamonds, AceOfClubs }

// Initialize
// AddCard    - done
// IndexOf    - done
// String
// Draw       - done
// Discard    - done
// DiscardAll
// Remaining  - done
// Reset      - done
// Shuffle

func prepDeck(cards []Card) (*Deck, int) {
	deck := NewDeck()

	for _, card := range cards {
		deck.AddCard(card)
	}
	return deck, len(*deck.Source)
}

func TestDeck_Initialize(t *testing.T) {
	deck := &Deck{}
	deck.Initialize()

	if deck.Source == nil || deck.Trash == nil || deck.Out == nil {
		t.Error(
			"For initializing the needed stacks",
			"expected 3 non nil values",
			"got", deck.Source, deck.Trash, deck.Out,
		)
	}
}

func TestDeck_AddCard(t *testing.T) {
	deck := NewDeck()

	for i, card := range Aces {
		deck.AddCard(card)
		if l := len(*deck.Source); l != i + 1 {
			t.Error(
				"For length of deck after deck.AddCard", i, card.Name,
				"expected", i + 1,
				"got", l,
			)
		}
	}
}

func TestDeck_IndexOf(t *testing.T) {

	deck, _ := prepDeck(Aces)

	for i, card := range Aces {
		x, y, z := deck.IndexOf(card)
		if x != i || y != -1 || z != -1 {
			t.Error(
				"For index of card", i, card.Name,
				"expected", i, -1, -1,
				"got", x, y, z,
			)
		}
	}
}

func TestDeck_Draw(t *testing.T) {

	deck, total := prepDeck(Aces)

	for i, expected := range Aces {
		drawn := deck.Draw()

		at := i + 1

		s_len := len(*deck.Source)
		o_len := len(*deck.Out)

		// are we drawing the right card?
		if drawn != expected {
			t.Error(
				"For card drawn",
				"expected", expected.Name,
				"got", drawn.(Card).Name,
			)
		}

		// is the card in the right place in the Out stack?
		out := (*deck.Out)[i]

		if out != drawn {
			t.Error(
				"For card out", i,
				"expected", drawn,
				"got",  out,
			)
		}

		// are the Source and Out piles scaling properly?
		if s_len != total - at || o_len != at || s_len + o_len != total {
			t.Error(
				"For pile lengths",
				"expected", total - at, at, total,
				"got", s_len, o_len, s_len + o_len,
			)
		}
	}
}

func TestDeck_Discard(t *testing.T) {

	deck, total := prepDeck(Aces)

	hand := make([]Card, total)

	for i := 0; i < total; i++ {
		hand[i] = deck.Draw().(Card)
	}

	for i, card := range hand {
		ok := deck.Discard(card)

		at := i + 1

		// did we get a proper discard result?
		if !ok {
			t.Error(
				"For discard",
				"expected", true,
				"got", ok,
			)
			continue
		}

		// are the Out and Trash piles scaling properly?
		t_len := len(*deck.Trash)
		o_len := len(*deck.Out)

		if o_len != total - at || t_len != at || o_len + t_len != total {
			t.Error(
				"For pile lengths",
				"expected", total - at, at, total,
				"got", o_len, t_len, o_len + t_len,
			)
		}

		trashed := (*deck.Trash)[0]

		// is the discarded card in the right place in the Trash stack?
		if trashed != card {
			t.Error(
				"For last discarded", i,
				"expected", card,
				"got",  trashed,
			)
		}
	}
}

func TestDeck_Remaining(t *testing.T) {

	deck, total := prepDeck(Aces)

	for i, _ := range Aces {
		deck.Draw()

		at := i + 1

		r := deck.Remaining()

		if r != total - at {
			t.Error(
				"For remaining",
				"expected", total - at,
				"got", r,
			)
		}
	}
}

func TestDeck_Reset(t *testing.T) {
	deck, total := prepDeck(Aces)

	half := total / 2

	cards := make([]Card, half)

	// draw half, then reset
	for i := 0; i < half; i++ {
		cards[i] = deck.Draw().(Card)
	}

	deck.Reset(false)

	s_len := len(*deck.Source)
	o_len := len(*deck.Out)
	t_len := len(*deck.Trash)

	if s_len != total - half || o_len != half || t_len != 0 {
		t.Error(
			"For reset after drawing",
			"expected", total - half, half, 0,
			"got", s_len, o_len, t_len,
		)
	}

	// discard drawn, then reset
	for _, card := range cards {
		deck.Discard(card)
	}

	deck.Reset(false)

	s_len = len(*deck.Source)
	o_len = len(*deck.Out)
	t_len = len(*deck.Trash)

	if s_len != total || o_len != 0 || t_len != 0 {
		t.Error(
			"For reset after drawing",
			"expected", total, 0, 0,
			"got", s_len, o_len, t_len,
		)
	}

	// draw half, then reset all
	for i := 0; i < half; i++ {
		deck.Draw()
	}

	deck.Reset(true)

	s_len = len(*deck.Source)
	o_len = len(*deck.Out)
	t_len = len(*deck.Trash)

	if s_len != total || o_len != 0 || t_len != 0 {
		t.Error(
			"For reset after drawing",
			"expected", total, 0, 0,
			"got", s_len, o_len, t_len,
		)
	}

}


func Example() {
	deck := NewDeck()

	deck.AddCard(AceOfSpades)
	deck.AddCard(AceOfHearts)
	deck.AddCard(AceOfDiamonds)
	deck.AddCard(AceOfClubs)

//	deck.Shuffle()

//	hand := NewCardSet()

//	hand.Push(deck.Draw())
//	hand.Push(deck.Draw())

	s, o, d := deck.IndexOf(AceOfDiamonds)
	fmt.Printf("%d, %d, %d\n", s, o, d)

//	deck.Discard(hand.Shift())

	// Output:
	// 2, -1, -1
}

func ExampleCardSet() {
	set := NewCardSet()

	set.Push(AceOfSpades)
	set.Push(AceOfHearts)

	set.Unshift(AceOfDiamonds)
	set.Unshift(AceOfClubs)

	fmt.Println(set.Pop().(Card).Name)
	fmt.Println(set.Shift().(Card).Name)

	// Output:
	// Ace Of Hearts
	// Ace Of Clubs
}

func ExampleDeck() {
	deck := NewDeck()

	for _, card := range Aces {
		deck.AddCard(card)
	}

	for deck.Remaining() > 0 {
		drawn := deck.Draw().(Card)

		fmt.Println(drawn.Name)

		deck.Discard(drawn)
	}

	deck.Reset(true)

	// Output:
	// Ace Of Spades
	// Ace Of Hearts
	// Ace Of Diamonds
	// Ace Of Clubs
}
