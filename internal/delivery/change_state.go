package delivery

import (
	"errors"

	"github.com/K0STYAa/vk_iproto/pkg/models"
	"github.com/vmihailenco/msgpack/v5"
)

func ADM_STORAGE_SWITCH_READONLY(d *Delivery, body []byte) ([]byte, error) {
	var unmarshal_body interface{}
	msgpack.Unmarshal(body, &unmarshal_body)
	if unmarshal_body != nil {
		return nil, errors.New("Incorrect body for function ADM_STORAGE_SWITCH_READONLY")
	}

	d.service.ChangeState(models.ReadOnly)
	return nil, nil
}

func ADM_STORAGE_SWITCH_READWRITE(d *Delivery, body []byte) ([]byte, error) {
	var unmarshal_body interface{}
	msgpack.Unmarshal(body, &unmarshal_body)
	if unmarshal_body != nil {
		return nil, errors.New("Incorrect body for function ADM_STORAGE_SWITCH_READWRITE")
	}

	d.service.ChangeState(models.ReadWrite)
	return nil, nil
}

func ADM_STORAGE_SWITCH_MAINTENANCE(d *Delivery, body []byte) ([]byte, error) {
	var unmarshal_body interface{}
	msgpack.Unmarshal(body, &unmarshal_body)
	if unmarshal_body != nil {
		return nil, errors.New("Incorrect body for function ADM_STORAGE_SWITCH_MAINTENANCE")
	}

	d.service.ChangeState(models.Maintenance)
	return nil, nil
}