// Package deck provides the standard tools necessary to manage a deck of cards for most games.
package deck

import(
	"math/rand"
	"encoding/json"
)

// Deck is the main struct. It keeps track of all cards, including cards that are in play
// and have not yet been discarded.
type Deck struct {
	Source     *CardSet
	Trash      *CardSet
	Out        *CardSet
}

// Reset returns all cards in the Trash ,back to the Source. Optionally, ita can also
// return the cards from Out to reset the deck to be used again. It returns itself as a matter of
// convenience, as you'll likely want to call Shuffle immediately afterwards.
func (p *Deck) Reset(all bool) *Deck {
	for p.Trash.Length() > 0 {
		p.Source.Push(p.Trash.Pop())
	}
	if all {
		for p.Out.Length() > 0 {
			p.Source.Push(p.Out.Pop())
		}
	}
	return p
}

// Shuffle shuffles all cards in the Source pile. You likely will want to call Reset first.
func (p *Deck) Shuffle() *Deck {
	l := p.Source.Length()
	for i := l - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		p.Source.Swap(i, j)
	}
	return p
}

// Draw pulls a single card off of the top of the Source pile, and keeps a record of its existence
// in the Out pile.
func (p *Deck) Draw() interface{} {
	out := p.Source.Shift()
	p.Out.Push(out)
	return out
}

func (p *Deck) discard(i int) {
	p.Trash.Unshift(p.Out.Remove(i))
}

// Discard sends a card to the top of the Trash pile. It first validates that the card is
// in the out state, and return false if it could not be moved.
func (p *Deck) Discard(card interface{}) bool {
	if i := p.Out.IndexOf(card); i != -1 {
		p.discard(i)
		card = nil
		return true
	}
	return false
}

// DiscardAll is a convenience method that discards a slice of cards after first verifying
// that they are all valid.
func (p *Deck) DiscardAll(cards []interface{}) bool {
	// validate them first
	for _, card := range cards {
		i := p.Out.IndexOf(card)
		if i >= 0 {
			return false
		}
	}
	for _, card := range cards {
		i := p.Out.IndexOf(card)
		if i >= 0 {
			p.discard(i)
		}
	}
	return true
}

// Remaining returns the number of cards remaining in Source to be played.
func (p *Deck) Remaining() int {
	return p.Source.Length()
}

// IndexOf searches all three card stacks and returns their index place, or -1 if not found.
// Ideally this would always return 2 values of -1 and 1 positive place index.
func (p *Deck) IndexOf(card interface{}) (int, int, int) {
	return p.Source.IndexOf(card), p.Out.IndexOf(card), p.Trash.IndexOf(card)
}

func (p *Deck) String() string {
	out, _ := json.Marshal(p)
	return string(out)
}

// AddCard introduces a card to the deck by adding it to the bottom of the Source pile.
// Usually only called while defining a deck before game play.
func (p *Deck) AddCard(card interface{}) *Deck {
	p.Source.Push(card)
	return p
}

// NewDeck initializes a deck and returns an instance.
func NewDeck() *Deck {
	return &Deck{
		Source: NewCardSet(),
		Trash: NewCardSet(),
		Out: NewCardSet(),
	}
}
