package internal

import (
	"sync"

	"github.com/K0STYAa/vk_iproto/pkg/models"
)

type BaseStorage struct {
	Data         [models.StorageDataLen]string
	StorageState uint8
	Mutex        sync.RWMutex
}
