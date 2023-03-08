package storage

import (
	"sync"
)

const (
	storageDataLen = 1000
	firstDataID    = 0
	stringMaxLen   = 256
)

type BaseStorage struct {
	Data         [storageDataLen]string
	StorageState uint8
	Mutex        sync.RWMutex
}
