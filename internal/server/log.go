package server

import (
	"fmt"
	"sync"
)

var (
	ErrOffsetNotFound = fmt.Errorf("offset not found")
)

// Record record struct
type Record struct {
	Value  []byte `json:"value"`
	Offset uint64 `json:"offset"`
}

// Log log struct
type Log struct {
	mu      sync.Mutex
	records []Record
}

// NewLog returns a new instance of Log
func NewLog() *Log {
	return &Log{}
}

// Append append a record to the log
func (c *Log) Append(record Record) (uint64, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	record.Offset = uint64(len(c.records))
	c.records = append(c.records, record)
	return record.Offset, nil
}

// Read read a record at index offset
func (c *Log) Read(offset uint64) (Record, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if offset >= uint64(len(c.records)) {
		return Record{}, ErrOffsetNotFound
	}
	return c.records[offset], nil
}
