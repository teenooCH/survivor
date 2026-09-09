package stack_test

import (
	"slices"
	"testing"

	"github.com/teenooCH/survivor/internal/pkg/stack"
)

func TestStack(t *testing.T) {
	tests := []struct {
		elements []int
	}{
		{[]int{}},
		{[]int{5}},
		{[]int{5, 6}},
		{[]int{5, 6, 7}},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			s := stack.New[int]()

			for _, element := range tt.elements {
				s.Push(element)
			}

			if s.Size() != len(tt.elements) {
				t.Errorf("Expected stack size to be %d, but got %d", len(tt.elements), s.Size())
			}

			for _, element := range slices.Backward(tt.elements) {
				peek, ok := s.Peek()
				if !ok || peek != element {
					t.Errorf("Expected top element to be %v, but got %v", element, peek)
				}

				popped, ok := s.Pop()
				if !ok || popped != element {
					t.Errorf("Expected popped element to be %v, but got %v", element, popped)
				}
			}

			if !s.IsEmpty() {
				t.Errorf("Expected stack to be empty after pop, but it is not")
			}
		})
	}
}
