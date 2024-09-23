package main

import (
	"bytes"
	"io"
	"testing"
)

func TestPathTransformFunc(t *testing.T) {
	key := "somemycoursework"
	pathKey := CASPathTransformFunc(key)
	// Test the pathname
	expectesOriginalKey := "f979e8bf9ff685078fa7bf471347a1dc8b5de691"
	expectedPathName := "f979e/8bf9f/f6850/78fa7/bf471/347a1/dc8b5/de691"
	if pathKey.PathName != expectedPathName {
		t.Errorf("Have %s, want %s", pathKey.PathName, expectedPathName)
	}

	if pathKey.Filename != expectesOriginalKey {
		t.Errorf("Have %s, want %s", pathKey.Filename, expectesOriginalKey)
	}
}

func TestStoreDeleteKey(t *testing.T) {
	opts := StoreOpts {
		PathTransformFunc: CASPathTransformFunc,
	}
	s := NewStore(opts)
	key := "somefile"
	data := []byte("some data")

	if err := s.writeStream(key, bytes.NewReader(data)); err != nil {
		t.Error(err)
	}

	if err := s.Delete(key); err != nil {
		t.Error(err)
	}
}

func TestStore(t *testing.T) {
	opts := StoreOpts {
		PathTransformFunc: CASPathTransformFunc,
	}
	s := NewStore(opts)
	key := "somefile"
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
	s.Delete(key)
}