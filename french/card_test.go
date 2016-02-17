package french

import(
	"fmt"
	"testing"
)

func Example() {
	cards := NewDecks(2, 4)

	fmt.Println(cards[13].(Card).Label)
	fmt.Println(cards[107].(Card).Label == WhiteJoker.Label)

	// Output:
	// Ace of Hearts
	// true
}

func TestCard(t *testing.T) {
	cards := NewDecks(2, 4)

	if l := len(cards); l != 108 {
		t.Error(
			"For new deck length",
			"expected", 108,
			"got", l,
		)
	}

	if card := cards[0].(Card); card.Id != AceOfSpades.Id || card.Deck != 0 {
		t.Error(
			"For validate first card",
			"expected", AceOfSpades.Id, 0,
			"got", card.Id, card.Deck,
		)
	}

	if card := cards[1].(Card); card.Id != DeuceOfSpades.Id || card.Deck != 0 {
		t.Error(
			"For validate second card",
			"expected", DeuceOfSpades.Id, 0,
			"got", card.Id, card.Deck,
		)
	}

	if card := cards[13].(Card); card.Id != AceOfHearts.Id || card.Deck != 0 {
		t.Error(
			"For validate first card of second suit",
			"expected", AceOfHearts.Id, 0,
			"got", card.Id, card.Deck,
		)
	}

	if card := cards[14].(Card); card.Id != DeuceOfHearts.Id || card.Deck != 0 {
		t.Error(
			"For validate second card of second suit",
			"expected", DeuceOfHearts.Id, 0,
			"got", card.Id, card.Deck,
		)
	}

	if card := cards[52].(Card); card.Id != AceOfSpades.Id || card.Deck != 1 {
		t.Error(
			"For validate first card of second deck",
			"expected", AceOfSpades.Id, 1,
			"got", card.Id, card.Deck,
		)
	}

	if card := cards[104].(Card); card.Id != BlackJoker.Id {
		t.Error(
			"For validate first joker",
			"expected", BlackJoker.Id, BlackJoker.Label,
			"got", card.Id, card.Label,
		)
	}

	if card := cards[107].(Card); card.Id != WhiteJoker.Id {
		t.Error(
			"For validate last joker",
			"expected", WhiteJoker.Id, WhiteJoker.Label,
			"got", card.Id, card.Label,
		)
	}

}
