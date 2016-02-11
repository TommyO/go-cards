package cards

import(
	"math/rand"
	"encoding/json"
)

type CardStack []interface{}

func (p *CardStack) remove(i int) interface{} {
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

func (p *CardStack) Reset() *CardStack {
	p = NewCardStack()
	return p
}

func (p *CardStack) Pop() interface{} {
	return p.remove(-1)
}

func (p *CardStack) Push(card interface{}) {
	*p = append(*p, card)
}

func (p *CardStack) Shift() interface{} {
	return p.remove(0)
}

func (p *CardStack) Unshift(card interface{}) {
	*p = append(CardStack{ card }, (*p)[:]...)
}

func (p *CardStack) Remove(i int) interface{} {
	return p.remove(i)
}

func (p *CardStack) IndexOf(card interface{}) int {
	for i, c := range *p {
		if c == card {
			return i
		}
	}
	return -1
}

func (p *CardStack) All() []interface{} {
	return (*p)[:]
}

func (p *CardStack) Length() int {
	return len(*p)
}

func (p *CardStack) Swap (i, j int) {
	left := (*p)[i]
	right := (*p)[j]
	(*p)[j] = left
	(*p)[i] = right
}

func NewCardStack() *CardStack {
	return &CardStack{}
}

//

type Deck struct {
	Source     *CardStack
	Trash      *CardStack
	Out        *CardStack
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

func (p *Deck) DiscardAll(cards *CardStack) bool {
	length := cards.Length()
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
		Source: NewCardStack(),
		Trash: NewCardStack(),
		Out: NewCardStack(),
	}
}
