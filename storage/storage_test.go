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


func TestQueryGet(t *testing.T) {
	s := initStorageWithData()
	cmd := Command{Key: KEY}
	reply := s.Query("get", cmd)
	assert.Equal(t, true, reply.Exists)
	assert.Equal(t, VAL, reply.Value)
	assert.Equal(t, KEY, reply.Key)
}

func TestQuerySetGet(t *testing.T) {
	s := initStorageWithData()
	setCmd := Command{Key: "newKey", Value: "newVal"}
	s.Query("set", setCmd)
	getCmd := Command{Key: "newKey"}
	reply := s.Query("get", getCmd)
	assert.Equal(t, true, reply.Exists)
	assert.Equal(t, "newVal", reply.Value)
	assert.Equal(t, "newKey", reply.Key)
}

func TestQueryDel(t *testing.T) {
	s := initStorageWithData()
	delCmd := Command{Key: KEY}
	reply := s.Query("del", delCmd)
	assert.Equal(t, true, reply.Exists)
}

func TestQueryDelNotExistent(t *testing.T) {
	s:= initStorageWithData()
	delCmd := Command{Key: "NonExistent"}
	reply := s.Query("del", delCmd)
	assert.Equal(t, false, reply.Exists)
}
