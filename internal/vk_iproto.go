package internal

import (
	"sync"
)

const (
	ReadWrite uint8 = iota
	ReadOnly
	Maintenance

	StorageDataLen = 1000
	FirstDataID    = 0
	StringMaxLen   = 256
)

type BaseStorage struct {
	Data         [StorageDataLen]string
	StorageState uint8
	Mutex        sync.RWMutex
}
