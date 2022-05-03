package encodedecode

import (
	"testing"
)

func TestEncodeLink(t *testing.T){
	got := EncodeLink(1913818838563217286)
	want := "kwC5c3ptxrc"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestDecodeLink(t *testing.T){
	got, _ := DecodeLink("eFfdcOt2U3r")
	want := uint64(15022863907731997794)

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func BenchmarkEncodeLink(b *testing.B){
    for i :=0; i < b.N ; i++{
        EncodeLink(1913818838563217286)
    }
}