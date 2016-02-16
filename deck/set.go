package deck

import(

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
