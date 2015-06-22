package bloom

import (
	"testing"
)

func TestPut(t *testing.T) {
	bf, _ := New(5)
	if err := bf.Put([]byte("foo")); err != nil {
		t.Errorf("%+v", err)
	}
}

func TestMightContain(t *testing.T) {
	bf, _ := New(1)
	key := []byte("foo")
	if actual := bf.MightContain(key); actual != false {
		t.Errorf("Expected false but received %v", actual)
	}
	bf.Put(key)
	if actual := bf.MightContain(key); actual != true {
		t.Errorf("Expected true but received %v", actual)
	}
}
