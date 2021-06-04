package co

import (
	"math/rand"
	"sync"
)

// NewSelector constructs a Selector instance
func NewSelector() (*Selector, error) {
	sel := &Selector{}
	sel.prior.addrs = make([]string, 0, 64)
	sel.all.addrs = make([]string, 0, 64)
	sel.all.nodes = map[string]*Node{}

	return sel, nil
}

// Selector is used to select a best chain node to route the requests to
type Selector struct {
	prior struct {
		sync.RWMutex
		addrs []string
	}

	all struct {
		sync.RWMutex
		addrs []string
		nodes map[string]*Node
	}
}

// ReplaceNodes adds and removes nodes
func (s *Selector) ReplaceNodes(add []*Node, removes map[string]bool, removesAll bool) {
	s.all.Lock()

	// reset
	if removesAll || len(removes) > 0 {
		s.all.addrs = s.all.addrs[:0]
		for host := range s.all.nodes {
			if removesAll || removes[host] {
				s.all.nodes[host].Stop()
				delete(s.all.nodes, host)
				continue
			}

			s.all.addrs = append(s.all.addrs, host)
		}
	}

	for i := range add {
		current := add[i]
		if prev, has := s.all.nodes[current.info.Addr]; has {
			prev.Stop()
		} else {
			s.all.addrs = append(s.all.addrs, current.info.Addr)
		}

		s.all.nodes[current.info.Addr] = current
		go current.Start()
	}

	s.all.Unlock()
}

func (s *Selector) setPriors(addrs ...string) {
	s.prior.Lock()
	s.prior.addrs = append(s.prior.addrs[:0], addrs...)
	s.prior.Unlock()
}

// Select tries to choose a node from the candidates
func (s *Selector) Select() (*Node, error) {
	var addr string

	s.prior.RLock()
	psize := len(s.prior.addrs)
	switch psize {
	case 0:

	case 1:
		addr = s.prior.addrs[0]

	default:
		addr = s.prior.addrs[rand.Intn(psize)]
	}
	s.prior.RUnlock()

	s.all.RLock()
	defer s.all.RUnlock()

	if addr != "" {
		if node, ok := s.all.nodes[addr]; ok {
			return node, nil
		}
	}

	allSize := len(s.all.addrs)
	if allSize == 0 {
		return nil, ErrNoNodeAvailable
	}

	return s.all.nodes[s.all.addrs[rand.Intn(allSize)]], nil
}
