package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStore(t *testing.T) {
	opts := StoreOpts{
		PathTransformer: CASPathTransform,
	}
	s := NewStore(opts)
	data := bytes.NewReader([]byte("Abc"))
	if err := s.writeStream("myfile", data); err != nil {
		t.Error(err)
	}

}

func TestCASPathTransform(t *testing.T) {
	pathName := CASPathTransform("myfile")
	assert.Equal(t, pathName.PathName, "b3580/ab45c/b088b/a47ff/070aa/81c2d/ae1be/56ca2")
	assert.Equal(t, pathName.Original, "b3580ab45cb088ba47ff070aa81c2dae1be56ca2")
}
