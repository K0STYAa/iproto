package internal

import "sync"

type Service struct {
    Data      		[1000]string
    StorageState 	int
    Mutex        	sync.RWMutex
}