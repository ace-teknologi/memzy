package memory

import (
	"testing"
)

type testObject struct {
	MegaID     string
	MegaAge    int
	MegaColour string
}

func TestStorePutThenGetItem(t *testing.T) {
	s := NewStore("MegaID")
	obj := &testObject{
		MegaID:     "123",
		MegaAge:    27,
		MegaColour: "brown",
	}

	err := s.PutItem(obj)
	if err != nil {
		t.Errorf("Unable to put item: %v", err)
	}

	newObj := &testObject{}

	err = s.GetItem(newObj, map[string]interface{}{"MegaID": "123"})
	if err != nil {
		t.Errorf("Unable to get item: %v", err)
	}

	err = s.GetItem(newObj, map[string]interface{}{"MegaID": "456"})
	if err == nil {
		t.Errorf("Expected an error")
	}
}

func TestStoreGetItemBadKeys(t *testing.T) {
	s := NewStore("derp")
	o := &testObject{}

	badKeys := []map[string]interface{}{
		{"wrongKeyDerp": "nar"},
		{
			"derp":  "soFarSogood",
			"derp2": "tooManyKeys",
		},
		{"derp": 22},
	}

	for _, k := range badKeys {
		err := s.GetItem(o, k)
		if err == nil {
			t.Errorf("Expected an error, got none")
		}
	}

}

func TestStorePutItemBad(t *testing.T) {
	s := NewStore("MegaID")

	badObjects := []interface{}{
		&testObject{},
		1,
		nil,
		&testObject{
			MegaAge: 11,
		},
		struct {
			WrongKey string
		}{
			"derp",
		},
	}

	for _, o := range badObjects {
		err := s.PutItem(o)
		if err == nil {
			t.Errorf("Expected an error, got none")
			return
		}
	}
}
