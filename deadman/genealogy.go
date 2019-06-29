package deadman

type Genealogy struct {
	genealogies []*DeadRegistry
	rank        int
}

func NewGenealogy(generation int, size uint) *Genealogy {
	genealogies := make([]*DeadRegistry, generation)
	for i := 0; i < generation; i++ {
		genealogies[i] = New(size)
	}
	return &Genealogy{
		genealogies: genealogies,
	}
}

func (g *Genealogy) Tick() {
	// FIXME it's not threadsafe
	g.rank++
	if g.rank >= len(g.genealogies) {
		g.rank = 0
	}
}

func (g *Genealogy) Current() *DeadRegistry {
	return g.genealogies[g.rank]
}

func (g *Genealogy) previous(n int) int {
	l := len(g.genealogies)
	return (g.rank + l - n) % l
}

func (g *Genealogy) Previous(n int) *DeadRegistry {
	// return nil if n > size
	if n >= (len(g.genealogies)-1) || n < 0 {
		return nil
	}
	return g.genealogies[g.previous(n)]
}
