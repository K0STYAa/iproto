package delivery

import (
	"errors"
	"fmt"

	"github.com/K0STYAa/vk_iproto/pkg/models"
	"github.com/sirupsen/logrus"
	"github.com/vmihailenco/msgpack/v5"
)

var (
	ErrIncorrectBodyReadOnly    = errors.New("incorrect body for function ADM_STORAGE_SWITCH_READONLY")
	ErrIncorrectBodyReadWrite   = errors.New("incorrect body for function ADM_STORAGE_SWITCH_READWRITE")
	ErrIncorrectBodyMaintenance = errors.New("incorrect body for function ADM_STORAGE_SWITCH_MAINTENANCE")
)

const errTemplate = "%w"

func ADM_STORAGE_SWITCH_READONLY(delivery *Delivery, body []byte) ([]byte, error) { //nolint: golint,revive,nosnakecase,stylecheck,lll
	var unmarshalBody interface{}

	err := msgpack.Unmarshal(body, &unmarshalBody)
	if err != nil {
		return nil, fmt.Errorf(errTemplate, err)
	}
	// LOG REQUEST
	logrus.Warn(fmt.Sprintf("[REQUEST]: ADM_STORAGE_SWITCH_READONLY(%v)", unmarshalBody))

	if unmarshalBody != nil {
		return nil, ErrIncorrectBodyReadOnly
	}

	delivery.service.ChangeState(models.ReadOnly)

	returnBody, err := msgpack.Marshal(nil)

	if err != nil {
		return returnBody, fmt.Errorf(errTemplate, err)
	}

	return returnBody, nil
}

func ADM_STORAGE_SWITCH_READWRITE(delivery *Delivery, body []byte) ([]byte, error) { //nolint: golint,revive,nosnakecase,stylecheck,lll
	var unmarshalBody interface{}

	err := msgpack.Unmarshal(body, &unmarshalBody)
	if err != nil {
		return nil, fmt.Errorf(errTemplate, err)
	}
	// LOG REQUEST
	logrus.Warn(fmt.Sprintf("[REQUEST]: ADM_STORAGE_SWITCH_READWRITE(%v)", unmarshalBody))

	if unmarshalBody != nil {
		return nil, ErrIncorrectBodyReadWrite
	}

	delivery.service.ChangeState(models.ReadWrite)

	returnBody, err := msgpack.Marshal(nil)

	if err != nil {
		return returnBody, fmt.Errorf(errTemplate, err)
	}

	return returnBody, nil
}

func ADM_STORAGE_SWITCH_MAINTENANCE(delivery *Delivery, body []byte) ([]byte, error) { //nolint: golint,revive,nosnakecase,stylecheck,lll
	var unmarshalBody interface{}

	err := msgpack.Unmarshal(body, &unmarshalBody)
	if err != nil {
		return nil, fmt.Errorf(errTemplate, err)
	}
	// LOG REQUEST
	logrus.Warn(fmt.Sprintf("[REQUEST]: ADM_STORAGE_SWITCH_MAINTENANCE(%v)", unmarshalBody))

	if unmarshalBody != nil {
		return nil, ErrIncorrectBodyMaintenance
	}

	delivery.service.ChangeState(models.Maintenance)

	returnBody, err := msgpack.Marshal(nil)

	if err != nil {
		return returnBody, fmt.Errorf(errTemplate, err)
	}

	return returnBody, nil
}
