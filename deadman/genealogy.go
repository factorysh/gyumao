package deadman

// Genealogy stores deads over time
type Genealogy struct {
	genealogies []*DeadRegistry
	rank        int // Age
}

// New Genealogy with a size and a number of generation
func New(generation int, size uint) *Genealogy {
	genealogies := make([]*DeadRegistry, generation)
	for i := 0; i < generation; i++ {
		genealogies[i] = NewDeadRegistry(size)
	}
	return &Genealogy{
		genealogies: genealogies,
	}
}

// Tick add one more generation
func (g *Genealogy) Tick() {
	// FIXME it's not threadsafe
	g.rank++
	if g.rank >= len(g.genealogies) {
		g.rank = 0
	}
}

// Current generation
func (g *Genealogy) Current() *DeadRegistry {
	return g.genealogies[g.rank]
}

func (g *Genealogy) previous(n int) int {
	l := len(g.genealogies)
	return (g.rank + l - n) % l
}

// Previous generation
func (g *Genealogy) Previous(n int) *DeadRegistry {
	// return nil if n > size
	if n >= (len(g.genealogies)-1) || n < 0 {
		return nil
	}
	return g.genealogies[g.previous(n)]
}

// Crunch n generations with a logical OR
func (g *Genealogy) Crunch(n int) *DeadRegistry {
	if n >= (len(g.genealogies)-1) || n <= 0 {
		return nil
	}
	current := g.Current()
	for i := 1; i < n; i++ {
		current = g.genealogies[g.previous(i)].Or(current)
	}
	return current
}
