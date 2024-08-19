package main

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestTransformFunc(t *testing.T) {
	key := "MomsFirstPicture"
	pathKey := CASPathTransformFunc(key)
	expectedOriginalKey := "9cd5e6c965171282e1429de38df7c1669576263e"
	expectedPathname := "9cd5e/6c965/17128/2e142/9de38/df7c1/66957/6263e"
	if pathKey.PathName != expectedPathname {
		t.Errorf("have %s want %s", pathKey.PathName, expectedPathname)
	}
	if pathKey.Filename != expectedPathname {
		t.Errorf("have %s want %s", pathKey.Filename, expectedOriginalKey)
	}
}

func TestStoreDeleteKey(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}
	s := NewStore(opts)
	key := "momSpecials"
	data := []byte("some jpg bytes")
	if err := s.writeStream(key, bytes.NewReader(data)); err != nil {
		t.Error(err)
	}
	if err := s.Delete(key); err != nil {
		t.Error(err)
	}
}

func TestStore(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}
	s := NewStore(opts)
	key := "momSpecials"
	data := []byte("some jpeg bytes")
	if err := s.writeStream(key, bytes.NewReader(data)); err != nil {
		t.Error(err)
	}
	if ok := s.Has(key); !ok {
		t.Errorf("expected to have key %s", key)
	}
	r, err := s.Read(key)
	if err != nil {
		t.Error(err)
	}
	b, _ := io.ReadAll(r)
	if string(b) != string(data) {
		t.Errorf("want %s have %s", data, b)
	}
	fmt.Println(string(b))
	s.Delete(key)
}
