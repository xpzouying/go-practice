package main

import (
	"encoding/binary"
	"errors"

	"github.com/spaolacci/murmur3"
)

const (
	headerEntrySize  = 4
	defCacheSize     = 128  // default count of cache block
	defMaxBufferSize = 4096 // default buffer size
)

type shard struct {
	items map[uint64]uint32 // map for cache, key is hash key, value is index of data in bytes array
	data  []byte            // data array, like: [ [00 00 01 10] [11 01 11 01 11] ... ]
	tail  int
}

func newShard() *shard {
	return &shard{
		items: make(map[uint64]uint32, defCacheSize),
		data:  make([]byte, defMaxBufferSize),
	}
}

func (s *shard) set(key []byte, entry []byte) {
	hashKey := keyToHashKey(key)

	data := wrapEntry(entry)
	s.items[hashKey] = uint32(s.tail)
	s.tail += copy(s.data[s.tail:], data)
}

func (s *shard) get(key []byte) ([]byte, error) {
	hashKey := keyToHashKey(key)
	idx, ok := s.items[hashKey]
	if !ok {
		return nil, errors.New("key not found")
	}

	// get wrap entry
	blobLen := binary.LittleEndian.Uint32(s.data[idx : idx+headerEntrySize])
	return s.data[idx+headerEntrySize : idx+headerEntrySize+blobLen], nil
}

// wrap entry with (data_length, data_body)
func wrapEntry(entry []byte) []byte {
	blobLen := len(entry)
	blob := make([]byte, headerEntrySize+blobLen) // [(blob_head), (blob_body)]

	binary.LittleEndian.PutUint32(blob[:headerEntrySize], uint32(blobLen))
	copy(blob[headerEntrySize:], entry)
	return blob
}

func keyToHashKey(key []byte) uint64 {
	hasher := murmur3.New64()
	hasher.Write(key)
	return hasher.Sum64()
}
