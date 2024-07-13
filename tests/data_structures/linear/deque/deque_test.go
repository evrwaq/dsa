package data_structures_linear_test

import (
	ds_errors "dsa/src/data_structures/errors"
	ds "dsa/src/data_structures/linear/deque"
	"testing"
)

func TestNewDeque(t *testing.T) {
	deque := ds.NewDeque()

	t.Run("Initial Size", func(t *testing.T) {
		if deque.Size() != 0 {
			t.Errorf("expected size 0, got %d", deque.Size())
		}
	})

	t.Run("IsEmpty on new deque", func(t *testing.T) {
		if !deque.IsEmpty() {
			t.Error("expected deque to be empty")
		}
	})
}

func TestDequeAddFront(t *testing.T) {
	deque := ds.NewDeque()

	t.Run("AddFront elements", func(t *testing.T) {
		deque.AddFront(1)
		if deque.Size() != 1 {
			t.Errorf("expected size 1, got %d", deque.Size())
		}

		deque.AddFront(2)
		if deque.Size() != 2 {
			t.Errorf("expected size 2, got %d", deque.Size())
		}

		deque.AddFront(3)
		if deque.Size() != 3 {
			t.Errorf("expected size 3, got %d", deque.Size())
		}

		if deque.IsEmpty() {
			t.Error("expected deque to not be empty")
		}
	})
}

func TestDequeAddBack(t *testing.T) {
	deque := ds.NewDeque()

	t.Run("AddBack elements", func(t *testing.T) {
		deque.AddBack(1)
		if deque.Size() != 1 {
			t.Errorf("expected size 1, got %d", deque.Size())
		}

		deque.AddBack(2)
		if deque.Size() != 2 {
			t.Errorf("expected size 2, got %d", deque.Size())
		}

		deque.AddBack(3)
		if deque.Size() != 3 {
			t.Errorf("expected size 3, got %d", deque.Size())
		}

		if deque.IsEmpty() {
			t.Error("expected deque to not be empty")
		}
	})
}

func TestDequeRemoveFront(t *testing.T) {
	deque := ds.NewDeque()
	deque.AddBack(1)
	deque.AddBack(2)
	deque.AddBack(3)

	t.Run("RemoveFront elements", func(t *testing.T) {
		value, err := deque.RemoveFront()
		if err != nil {
			t.Errorf(ds_errors.UnexpectedError, err)
		}
		if value != 1 {
			t.Errorf("expected value 1, got %v", value)
		}
		if deque.Size() != 2 {
			t.Errorf("expected size 2, got %d", deque.Size())
		}

		value, err = deque.RemoveFront()
		if err != nil {
			t.Errorf(ds_errors.UnexpectedError, err)
		}
		if value != 2 {
			t.Errorf("expected value 2, got %v", value)
		}
		if deque.Size() != 1 {
			t.Errorf("expected size 1, got %d", deque.Size())
		}

		value, err = deque.RemoveFront()
		if err != nil {
			t.Errorf(ds_errors.UnexpectedError, err)
		}
		if value != 3 {
			t.Errorf("expected value 3, got %v", value)
		}
		if deque.Size() != 0 {
			t.Errorf("expected size 0, got %d", deque.Size())
		}
	})

	t.Run("RemoveFront from empty deque", func(t *testing.T) {
		_, err := deque.RemoveFront()
		if err == nil {
			t.Error(ds_errors.ExpectedError)
		}
	})
}

func TestDequeRemoveBack(t *testing.T) {
	deque := ds.NewDeque()
	deque.AddBack(1)
	deque.AddBack(2)
	deque.AddBack(3)

	t.Run("RemoveBack elements", func(t *testing.T) {
		value, err := deque.RemoveBack()
		if err != nil {
			t.Errorf(ds_errors.UnexpectedError, err)
		}
		if value != 3 {
			t.Errorf("expected value 3, got %v", value)
		}
		if deque.Size() != 2 {
			t.Errorf("expected size 2, got %d", deque.Size())
		}

		value, err = deque.RemoveBack()
		if err != nil {
			t.Errorf(ds_errors.UnexpectedError, err)
		}
		if value != 2 {
			t.Errorf("expected value 2, got %v", value)
		}
		if deque.Size() != 1 {
			t.Errorf("expected size 1, got %d", deque.Size())
		}

		value, err = deque.RemoveBack()
		if err != nil {
			t.Errorf(ds_errors.UnexpectedError, err)
		}
		if value != 1 {
			t.Errorf("expected value 1, got %v", value)
		}
		if deque.Size() != 0 {
			t.Errorf("expected size 0, got %d", deque.Size())
		}
	})

	t.Run("RemoveBack from empty deque", func(t *testing.T) {
		_, err := deque.RemoveBack()
		if err == nil {
			t.Error(ds_errors.ExpectedError)
		}
	})
}

func TestDequePeekFront(t *testing.T) {
	deque := ds.NewDeque()
	deque.AddBack(1)
	deque.AddBack(2)
	deque.AddBack(3)

	t.Run("PeekFront element", func(t *testing.T) {
		value, err := deque.PeekFront()
		if err != nil {
			t.Errorf(ds_errors.UnexpectedError, err)
		}
		if value != 1 {
			t.Errorf("expected value 1, got %v", value)
		}
		if deque.Size() != 3 {
			t.Errorf("expected size 3, got %d", deque.Size())
		}
	})

	t.Run("PeekFront from empty deque", func(t *testing.T) {
		deque.RemoveFront()
		deque.RemoveFront()
		deque.RemoveFront()
		_, err := deque.PeekFront()
		if err == nil {
			t.Error(ds_errors.ExpectedError)
		}
	})
}

func TestDequePeekBack(t *testing.T) {
	deque := ds.NewDeque()
	deque.AddBack(1)
	deque.AddBack(2)
	deque.AddBack(3)

	t.Run("PeekBack element", func(t *testing.T) {
		value, err := deque.PeekBack()
		if err != nil {
			t.Errorf(ds_errors.UnexpectedError, err)
		}
		if value != 3 {
			t.Errorf("expected value 3, got %v", value)
		}
		if deque.Size() != 3 {
			t.Errorf("expected size 3, got %d", deque.Size())
		}
	})

	t.Run("PeekBack from empty deque", func(t *testing.T) {
		deque.RemoveBack()
		deque.RemoveBack()
		deque.RemoveBack()
		_, err := deque.PeekBack()
		if err == nil {
			t.Error(ds_errors.ExpectedError)
		}
	})
}
