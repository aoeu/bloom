package bloom

import (
	"testing"
)


func TestPut(t  *testing.T) {
	bf, err := New(5)
	if err != nil {
		t.Errorf("%+v", err)
	}
	if err := bf.Put([]byte("foo")); err != nil {
		t.Errorf("%+v", err)
	}
}