package collision

import (
	"iter"
	"slices"
)

type Manager struct {
	colliders []*Collider
}

func NewManager() *Manager {
	return &Manager{
		colliders: make([]*Collider, 0),
	}
}

func (m *Manager) AddCollider(c *Collider) {
	m.colliders = append(m.colliders, c)
}

func (m *Manager) RemoveCollider(c *Collider) {
	i := slices.Index(m.colliders, c)
	if i != -1 {
		m.colliders = slices.Delete(m.colliders, i, i+1)
	}
}

func (m *Manager) Colliders() iter.Seq[*Collider] {
	return slices.Values(m.colliders)
}

func (m *Manager) ProcessCollisions() {
	for a, b := range collisionPairs(m.colliders) {
		if !a.CanCollideWith(b) || !b.CanCollideWith(a) {
			continue
		}

		if !a.OverlapsWith(b) {
			continue
		}

		if a.CollisionHandler() != nil {
			a.CollisionHandler()(b)
		}

		if b.CollisionHandler() != nil {
			b.CollisionHandler()(a)
		}
	}
}

// collisionPairs returns all pairs of s.
func collisionPairs[S ~[]T, T any](s S) iter.Seq2[T, T] {
	return func(yield func(T, T) bool) {
		for i := 0; i < len(s)-1; i++ {
			for j := i + 1; j < len(s); j++ {
				if !yield(s[i], s[j]) {
					return
				}
			}
		}
	}
}
