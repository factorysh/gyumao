package deadman

import (
	"sync"
	"time"
)

// Genealogy stores deads over time
type Genealogy struct {
	genealogies []*DeadRegistry
	rank        int // Age
	lock        sync.Mutex
	duration    time.Duration
}

// New Genealogy with a size and a number of generation
func New(generation int, keys []string, duration time.Duration) *Genealogy {
	genealogies := make([]*DeadRegistry, generation)
	genealogies[0] = NewDeadRegistry(keys)
	for i := 1; i < generation; i++ {
		genealogies[i] = genealogies[0].Ghost()
	}
	return &Genealogy{
		genealogies: genealogies,
		duration:    duration,
	}
}

func (g *Genealogy) AutoTick() {
	if g.duration != 0 {
		go func() {
			t := time.Tick(g.duration)
			for {
				<-t
				g.Tick()
			}
		}()
	}
}

// Tick add one more generation
func (g *Genealogy) Tick() {
	g.lock.Lock()
	defer g.lock.Unlock()
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
