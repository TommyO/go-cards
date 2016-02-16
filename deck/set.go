package deck

import(

)

// CardSet is a collection of cards, and some convenient ways to manage them. Deck uses it to
// manage its three stacks internally. It may also be helpful in handling a player's hand.
// It includes methods similar to stack and queue methods from other languages.
type CardSet []interface{}

func (p *CardSet) remove(i int) interface{} {
	if i < 0 {
		i += len(*p)
	}
	if i > p.Length() {
		return nil
	}
	out := (*p)[i]
	if i == 0 {
		*p = (*p)[1:]
	} else if i == len(*p) - 1 {
		*p = (*p)[0:i]
	} else {
		*p = append((*p)[:i], (*p)[i+1:]...)
	}
	return out
}

// Peek reveals a card at a specific index without changing its position in any way.
func (p *CardSet) Peek(i int) interface{} {
	return (*p)[i]
}

// Reset throws away all cards. May be deprecated, as it isn't super helpful anymore.
func (p *CardSet) Reset() *CardSet {
	p = NewCardSet()
	return p
}

// Pop removes and returns the last card from the slice
func (p *CardSet) Pop() interface{} {
	return p.remove(-1)
}

// Push adds a new card to the end of the slice
func (p *CardSet) Push(card interface{}) {
	*p = append(*p, card)
}

// Shift removes and returns the first card from the slice
func (p *CardSet) Shift() interface{} {
	return p.remove(0)
}

// Unshift pushes a card to the front of the slice
func (p *CardSet) Unshift(card interface{}) {
	*p = append(CardSet{ card }, (*p)[:]...)
}

// remove pulls a card from an arbitrary point in the slice
func (p *CardSet) Remove(i int) interface{} {
	return p.remove(i)
}

// IndexOf searches for a card in the slice, and returns its position, or -1 if not found.
func (p *CardSet) IndexOf(card interface{}) int {
	for i, c := range (*p)[:] {
		if c == card {
			return i
		}
	}
	return -1
}

// All returns the embedded slice without updating the slice. Useful?
func (p *CardSet) All() []interface{} {
	return (*p)[:]
}

// Length returns the length of the slice.
func (p *CardSet) Length() int {
	return len(*p)
}

// Swap switches the values of two positions in the slice
func (p *CardSet) Swap(i, j int) {
	left := (*p)[i]
	right := (*p)[j]
	(*p)[j] = left
	(*p)[i] = right
}

// returns a new instance of CardSet
func NewCardSet() *CardSet {
	return &CardSet{}
}
