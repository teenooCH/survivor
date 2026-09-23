package collision_test

import (
	"testing"

	"survivor/internal/domain/collision"
	"survivor/internal/domain/node2D"
)

func TestManager_ProcessCollisions(t *testing.T) {
	manager := collision.NewManager()

	first := collision.NewCollider(
		"first",
		collision.NewMask(collision.LayerPlayer, collision.LayerEnemy|collision.LayerProjectile),
		collision.NewCircle(10),
	)
	second := collision.NewCollider(
		"second",
		collision.NewMask(collision.LayerEnemy, collision.LayerPlayer),
		collision.NewCircle(10),
	)
	third := collision.NewCollider(
		"third",
		collision.NewMask(collision.LayerProjectile, collision.LayerPlayer),
		collision.NewCircle(10),
	)

	parentFirst := node2D.New("parent-first")
	parentFirst.AddChild(first)

	parentSecond := node2D.New("parent-second")
	parentSecond.AddChild(second)
	parentSecond.SetPosition(10, 0)

	parentThird := node2D.New("parent-third")
	parentThird.AddChild(third)
	parentThird.SetPosition(50, 0)

	var firstHit, secondHit, thirdHit *collision.Collider

	first.SetCollisionHandler(func(other *collision.Collider) {
		firstHit = other
	})
	second.SetCollisionHandler(func(other *collision.Collider) {
		secondHit = other
	})
	third.SetCollisionHandler(func(other *collision.Collider) {
		thirdHit = other
	})
	manager.AddCollider(first)
	manager.AddCollider(second)
	manager.AddCollider(third)

	manager.ProcessCollisions()

	if firstHit != second {
		t.Errorf("first collider handler received %v, want second collider", firstHit)
	}

	if secondHit != first {
		t.Errorf("second collider handler received %v, want first collider", secondHit)
	}

	if thirdHit != nil {
		t.Errorf("third collider handler received %v, want first collider", thirdHit)
	}
}

func TestManager_ProcessCollisions_DoesNotDispatchWhenMasksDoNotMatch(t *testing.T) {
	manager := collision.NewManager()

	first := collision.NewCollider(
		"first",
		collision.NewMask(collision.LayerPlayer, collision.LayerEnemy),
		collision.NewCircle(10),
	)
	second := collision.NewCollider(
		"second",
		collision.NewMask(collision.LayerEnemy, 0),
		collision.NewCircle(10),
	)

	parentFirst := node2D.New("parent-first")
	parentFirst.AddChild(first)

	parentSecond := node2D.New("parent-second")
	parentSecond.AddChild(second)

	hits := 0

	first.SetCollisionHandler(func(*collision.Collider) { hits++ })
	second.SetCollisionHandler(func(*collision.Collider) { hits++ })
	manager.AddCollider(first)
	manager.AddCollider(second)

	manager.ProcessCollisions()

	if hits != 0 {
		t.Errorf("collision handlers called %d times, want 0", hits)
	}
}
