package delivery

import (
	"errors"
	"fmt"

	"github.com/K0STYAa/vk_iproto/pkg/models"
	"github.com/sirupsen/logrus"
	"github.com/vmihailenco/msgpack/v5"
)

func ADM_STORAGE_SWITCH_READONLY(d *Delivery, body []byte) ([]byte, error) {
	var unmarshal_body interface{}
	err := msgpack.Unmarshal(body, &unmarshal_body)
    if err != nil {
        return nil, err
    }
	// LOG REQUEST
	logrus.Warn(fmt.Sprintf("[REQUEST]: ADM_STORAGE_SWITCH_READONLY(%v)", unmarshal_body))
	if unmarshal_body != nil {
		return nil, errors.New("Incorrect body for function ADM_STORAGE_SWITCH_READONLY")
	}

	d.service.ChangeState(models.ReadOnly)
	return_body, err := msgpack.Marshal(nil)
	return return_body, err
}

func ADM_STORAGE_SWITCH_READWRITE(d *Delivery, body []byte) ([]byte, error) {
	var unmarshal_body interface{}
	err := msgpack.Unmarshal(body, &unmarshal_body)
    if err != nil {
        return nil, err
    }
	// LOG REQUEST
	logrus.Warn(fmt.Sprintf("[REQUEST]: ADM_STORAGE_SWITCH_READWRITE(%v)", unmarshal_body))
	if unmarshal_body != nil {
		return nil, errors.New("Incorrect body for function ADM_STORAGE_SWITCH_READWRITE")
	}

	d.service.ChangeState(models.ReadWrite)
	return_body, err := msgpack.Marshal(nil)
	return return_body, err
}

func ADM_STORAGE_SWITCH_MAINTENANCE(d *Delivery, body []byte) ([]byte, error) {
	var unmarshal_body interface{}
	err := msgpack.Unmarshal(body, &unmarshal_body)
    if err != nil {
        return nil, err
    }
	// LOG REQUEST
	logrus.Warn(fmt.Sprintf("[REQUEST]: ADM_STORAGE_SWITCH_MAINTENANCE(%v)", unmarshal_body))
	if unmarshal_body != nil {
		return nil, errors.New("Incorrect body for function ADM_STORAGE_SWITCH_MAINTENANCE")
	}

	d.service.ChangeState(models.Maintenance)
	return_body, err := msgpack.Marshal(nil)
	return return_body, err
}