package bytespool

import (
	"errors"
)

// BytesPool manage all byte pool
type BytesPool struct {
	data           []byte
	length         []uint // record every element length
	maxElementSize int
	maxElementNum  int
}

// NewBytesPool alloc byte buffer once
func NewBytesPool(elementSize, elementNum int) *BytesPool {
	return &BytesPool{
		data:           make([]byte, elementNum*elementSize),
		length:         make([]uint, elementNum),
		maxElementNum:  elementNum,
		maxElementSize: elementSize,
	}
}

// Set data to bytes pool at index position. index begin from 0.
func (bp *BytesPool) Set(index int, data []byte) error {
	if index < 0 || index >= bp.maxElementNum {
		return errors.New("invalid index")
	}

	l := len(data)
	if len(data) > bp.maxElementSize {
		return errors.New("data size is bigger than element size")
	}

	bp.length[index] = uint(l)
	begin := index * bp.maxElementSize
	copy(bp.data[begin:begin+l], data)

	return nil
}

// Get data from pool at index. range of index is [0, maxElementNum]
func (bp *BytesPool) Get(index int) ([]byte, error) {
	if index >= bp.maxElementNum {
		return nil, errors.New("invalid index")
	}

	b := make([]byte, bp.length[index])
	begin := bp.maxElementSize * index
	end := begin + int(bp.length[index])
	copy(b, bp.data[begin:end])

	return b, nil
}
