package main

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestPathTransformFunc(t *testing.T) {
	key := "somemycoursework"
	pathKey := CASPathTransformFunc(key)
	expectedFilename := "f979e8bf9ff685078fa7bf471347a1dc8b5de691"
	expectedPathName := "f979e/8bf9f/f6850/78fa7/bf471/347a1/dc8b5/de691"
	if pathKey.PathName != expectedPathName {
		t.Errorf("have %s want %s", pathKey.PathName, expectedPathName)
	}

	if pathKey.Filename != expectedFilename {
		t.Errorf("have %s want %s", pathKey.Filename, expectedFilename)
	}
}

func TestStore(t *testing.T) {
	s := newStore()
	id := generateID()
	defer teardown(t, s)

	for i := 0; i < 50; i++ {
		key := fmt.Sprintf("foo_%d", i)
		data := []byte("some jpg bytes")

		if _, err := s.writeStream(id, key, bytes.NewReader(data)); err != nil {
			t.Error(err)
		}

		if ok := s.Has(id, key); !ok {
			t.Errorf("expected to have key %s", key)
		}

		_, r, err := s.Read(id, key)
		if err != nil {
			t.Error(err)
		}

		b, _ := io.ReadAll(r)
		if string(b) != string(data) {
			t.Errorf("want %s have %s", data, b)
		}

		if err := s.Delete(id, key); err != nil {
			t.Error(err)
		}

		if ok := s.Has(id, key); ok {
			t.Errorf("expected to NOT have key %s", key)
		}
	}
}

func newStore() *Store {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}
	return NewStore(opts)
}

func teardown(t *testing.T, s *Store) {
	if err := s.Clear(); err != nil {
		t.Error(err)
	}
}
