package main

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func newStroe() *Store {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}
	return NewStore(opts)
}

func tearDown(t *testing.T, s *Store) {
	if err := s.Clear(); err != nil {
		t.Error(err)
	}
}

func TestPathTransformFunc(t *testing.T) {
	key := "somemycoursework"
	pathKey := CASPathTransformFunc(key)
	// Test the pathname
	expectesFileName := "f979e8bf9ff685078fa7bf471347a1dc8b5de691"
	expectedPathName := "f979e/8bf9f/f6850/78fa7/bf471/347a1/dc8b5/de691"
	if pathKey.PathName != expectedPathName {
		t.Errorf("Have %s, want %s", pathKey.PathName, expectedPathName)
	}

	if pathKey.Filename != expectesFileName {
		t.Errorf("Have %s, want %s", pathKey.Filename, expectesFileName)
	}
}

func TestStore(t *testing.T) {
	s := newStroe()
	defer tearDown(t, s)

	for i := 0; i < 50; i++ {
		key := fmt.Sprintf("somefile_%d", i)
		data := []byte("some data")

		if err := s.writeStream(key, bytes.NewReader(data)); err != nil {
			t.Error(err)
		}

		if ok := s.Has(key); !ok {
			t.Errorf("Expected to have key %s", key)
		}

		r, err := s.Read(key)
		if err != nil {
			t.Error(err)
		}

		b, _ := io.ReadAll(r)
		if string(b) != string(data) {
			t.Errorf("want %s have %s", data, b)
		}

		if err := s.Delete(key); err != nil {
			t.Error(err)
		}

		if ok := s.Has(key); ok {
			t.Errorf("Expected NOT to have key %s", key)
		}
	}
}
