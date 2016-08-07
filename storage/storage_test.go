package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const KEY = "myKey"
const VAL = "myValue"

func initStorageWithData() *Storage {
	s := Storage{}
	s.Init()
	s.set(KEY, VAL)
	return &s
}

func TestSetGet(t *testing.T) {
	s := initStorageWithData()
	v, ok := s.get(KEY)
	assert.Equal(t, VAL, v)
	assert.Equal(t, true, ok)
}

func TestGetNotExist(t *testing.T) {
	s := initStorageWithData()
	v, ok := s.get("nonExistentKey")
	assert.Equal(t, "", v)
	assert.Equal(t, false, ok)
}

func TestDel(t *testing.T) {
	s := initStorageWithData()
	ok := s.del(KEY)
	assert.Equal(t, true, ok)
}

func TestDelNotExistent(t *testing.T) {
	s := initStorageWithData()
	ok := s.del("noKey")
	assert.Equal(t, false, ok)
}
