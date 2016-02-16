package deck

import(
	"math/rand"
	"encoding/json"
)

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

func (p *CardSet) Peek(i int) interface{} {
	return (*p)[i]
}

func (p *CardSet) Reset() *CardSet {
	p = NewCardSet()
	return p
}

func (p *CardSet) Pop() interface{} {
	return p.remove(-1)
}

func (p *CardSet) Push(card interface{}) {
	*p = append(*p, card)
}

func (p *CardSet) Shift() interface{} {
	return p.remove(0)
}

func (p *CardSet) Unshift(card interface{}) {
	*p = append(CardSet{ card }, (*p)[:]...)
}

func (p *CardSet) Remove(i int) interface{} {
	return p.remove(i)
}

func (p *CardSet) IndexOf(card interface{}) int {
	for i, c := range (*p)[:] {
		if c == card {
			return i
		}
	}
	return -1
}

func (p *CardSet) All() []interface{} {
	return (*p)[:]
}

func (p *CardSet) Length() int {
	return len(*p)
}

func (p *CardSet) Swap(i, j int) {
	left := (*p)[i]
	right := (*p)[j]
	(*p)[j] = left
	(*p)[i] = right
}

func NewCardSet() *CardSet {
	return &CardSet{}
}

//

type Deck struct {
	Source     *CardSet
	Trash      *CardSet
	Out        *CardSet
}

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

func (p *Deck) Shuffle() *Deck {
	l := p.Source.Length()
	for i := l - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		p.Source.Swap(i, j)
	}
	return p
}

func (p *Deck) Draw() interface{} {
	out := p.Source.Shift()
	p.Out.Push(out)
	return out
}

func (p *Deck) discard(i int) {
	p.Trash.Unshift(p.Out.Remove(i))
}

func (p *Deck) Discard(card interface{}) bool {
	if i := p.Out.IndexOf(card); i != -1 {
		p.discard(i)
		card = nil
		return true
	}
	return false
}

func (p *Deck) DiscardSet(cards *CardSet) bool {
	// validate them first
	for _, card := range *cards {
		i := p.Out.IndexOf(card)
		if i >= 0 {
			return false
		}
	}
	for _, card := range *cards {
		i := p.Out.IndexOf(card)
		if i >= 0 {
			p.discard(i)
		}
	}
	return true
}

func (p *Deck) Remaining() int {
	return p.Source.Length()
}

func (p *Deck) IndexOf(card interface{}) (int, int, int) {
	return p.Source.IndexOf(card), p.Out.IndexOf(card), p.Trash.IndexOf(card)
}

func (p *Deck) String() string {
	out, _ := json.Marshal(p)
	return string(out)
}

func (p *Deck) AddCard(card interface{}) *Deck {
	p.Source.Push(card)
	return p
}

func NewDeck() *Deck {
	return &Deck{
		Source: NewCardSet(),
		Trash: NewCardSet(),
		Out: NewCardSet(),
	}
}
