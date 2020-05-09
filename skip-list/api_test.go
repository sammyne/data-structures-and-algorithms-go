package skiplist_test

import (
	"strconv"
	"testing"

	skiplist "github.com/sammyne/data-structures-go/skip-list"
)

func newSkipList(values ...int) *skiplist.SkipList {
	list := skiplist.NewSkipList()
	for _, v := range values {
		list.Insert(v)
	}

	return list
}

func TestSkipListDelete(t *testing.T) {
	testVector := []struct {
		list   *skiplist.SkipList
		value  int
		expect string
	}{
		{
			newSkipList(9, 1, 7, 7, 3),
			3,
			"1 7 9",
		},
		{
			newSkipList(9, 1, 7),
			7,
			"1 9",
		},
		{
			newSkipList(9, 1),
			1,
			"9",
		},
		{
			newSkipList(1),
			1,
			"",
		},
		{
			newSkipList(9, 1, 7, 7, 3),
			4,
			"1 3 7 9",
		},
	}

	for i, v := range testVector {
		v.list.Delete(v.value)
		if got, _ := v.list.LevelString(0); v.expect != got {
			t.Fatalf("#%d fail: expect '%s', got '%s'", i, v.expect, got)
		}
	}
}

func TestSkipListFind(t *testing.T) {
	testVector := []struct {
		list      *skiplist.SkipList
		value     int
		expectErr error
	}{
		{
			newSkipList(9, 1, 7, 3),
			3,
			nil,
		},
		{
			newSkipList(9, 1, 7, 3),
			4,
			skiplist.ErrValueNotFound,
		},
	}

	for i, v := range testVector {
		if n, err := v.list.Find(v.value); err != v.expectErr {
			t.Fatalf("#%d unexpected error: expect '%v', got '%v'", i, v.expectErr, err)
		} else if err == nil && n.Value() != v.value {
			t.Fatalf("#%d invalid value: expect %v, got %d", i, v.value, n.Value())
		}
	}
}

func TestSkipListInsert(t *testing.T) {
	values := []int{9, 1, 7, 7, 3}

	list := skiplist.NewSkipList()

	{
		list.Insert(values[0])

		n, err := list.Find(values[0])
		if err != nil {
			t.Fatalf("missing values[0]: %v", err)
		}

		maxLevel := n.MaxLevel()

		expect := strconv.Itoa(n.Value())
		for i := 0; i <= maxLevel; i++ {
			got, err := list.LevelString(i)
			if err != nil {
				t.Fatalf("[0] invalid level[%d]: unexpected error %v", i, err)
			} else if expect != got {
				t.Fatalf("[0] invalid level[%d]: expect %s, got %s", i, expect, got)
			}
		}

		for i := maxLevel + 1; i < skiplist.MaxLevel; i++ {
			if _, err := list.LevelString(i); err != skiplist.ErrLevelNotFound {
				t.Fatalf("[0] invalid level[%d]: expect error %s, got %v", i,
					skiplist.ErrLevelNotFound, err)
			}
		}
	}

	{
		list.Insert(values[1])

		if _, err := list.Find(values[1]); err != nil {
			t.Fatalf("missing values[1]: %v", err)
		}

		const expect = "1 9"
		if got, err := list.LevelString(0); err != nil {
			t.Fatalf("[1] invalid level[%d]: unexpected error %v", 0, err)
		} else if expect != got {
			t.Fatalf("[1] invalid level[%d]: expect %s, got %s", 0, expect, got)
		}
	}

	{
		v := values[2]
		list.Insert(v)

		if _, err := list.Find(v); err != nil {
			t.Fatalf("missing values[2]: %v", err)
		}

		const expect = "1 7 9"
		if got, err := list.LevelString(0); err != nil {
			t.Fatalf("[2] invalid level[%d]: unexpected error %v", 0, err)
		} else if expect != got {
			t.Fatalf("[2] invalid level[%d]: expect %s, got %s", 0, expect, got)
		}
	}

	{
		v := values[3]
		list.Insert(v)

		if _, err := list.Find(v); err != nil {
			t.Fatalf("missing values[3]: %v", err)
		}

		const expect = "1 7 9"
		if got, err := list.LevelString(0); err != nil {
			t.Fatalf("[3] invalid level[%d]: unexpected error %v", 0, err)
		} else if expect != got {
			t.Fatalf("[3] invalid level[%d]: expect %s, got %s", 0, expect, got)
		}
	}

	{
		v := values[4]
		list.Insert(v)

		if _, err := list.Find(v); err != nil {
			t.Fatalf("missing values[4]: %v", err)
		}

		const expect = "1 3 7 9"
		if got, err := list.LevelString(0); err != nil {
			t.Fatalf("[4] invalid level[%d]: unexpected error %v", 0, err)
		} else if expect != got {
			t.Fatalf("[4] invalid level[%d]: expect %s, got %s", 0, expect, got)
		}
	}
}
