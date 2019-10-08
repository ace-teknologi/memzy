package memory

import (
	"testing"

	"github.com/ace-teknologi/memzy"
)

func TestIterImplementsMemzy(t *testing.T) {
	i := &Iter{}
	if _, ok := interface{}(i).(memzy.Iter); !ok {
		t.Errorf("Iter does not implement memzy.Iter")
	}
}

func TestIter(t *testing.T) {
	s := NewStore("MegaID")
	s.PutItem(&testObject{
		MegaID:     "1",
		MegaAge:    27,
		MegaColour: "Hot Pink",
	})
	s.PutItem(&testObject{
		MegaID:     "2",
		MegaAge:    27,
		MegaColour: "Hot Pink",
	})
	s.PutItem(&testObject{
		MegaID:     "3",
		MegaAge:    27,
		MegaColour: "Hot Pink",
	})
	s.PutItem(&testObject{
		MegaID:     "4",
		MegaAge:    27,
		MegaColour: "Hot Pink",
	})

	i := s.NewIter()
	var count int
	for i.Next() {
		count++
		var obj testObject
		err := i.Current(&obj)
		if err != nil {
			t.Error(err)
		}
	}

	if 4 != count {
		t.Errorf("Expected 4 items in the Iter, got %d", count)
	}

	if i.Err() != nil {
		t.Error(i.Err())
	}

}
