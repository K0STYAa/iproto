package internal

import "sync"

type BaseStorage struct {
    Data      		[1000]string
    StorageState 	uint8
    Mutex        	sync.RWMutex
}