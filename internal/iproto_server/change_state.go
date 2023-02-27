package iprotoserver

import (
	"errors"
	"fmt"

	"github.com/K0STYAa/vk_iproto/internal"
	"github.com/sirupsen/logrus"
	"github.com/vmihailenco/msgpack/v5"
)

var (
	ErrIncorrectBodyReadOnly    = errors.New("incorrect body for function ADM_STORAGE_SWITCH_READONLY")
	ErrIncorrectBodyReadWrite   = errors.New("incorrect body for function ADM_STORAGE_SWITCH_READWRITE")
	ErrIncorrectBodyMaintenance = errors.New("incorrect body for function ADM_STORAGE_SWITCH_MAINTENANCE")
)

const errTemplate = "%w"

func ADM_STORAGE_SWITCH_READONLY(iprotoserver *IprotoServer, body []byte) ([]byte, error) { //nolint: golint,revive,nosnakecase,stylecheck,lll
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

	iprotoserver.usecase.ChangeState(internal.ReadOnly)

	returnBody, err := msgpack.Marshal(nil)

	if err != nil {
		return returnBody, fmt.Errorf(errTemplate, err)
	}

	return returnBody, nil
}

func ADM_STORAGE_SWITCH_READWRITE(iprotoserver *IprotoServer, body []byte) ([]byte, error) { //nolint: golint,revive,nosnakecase,stylecheck,lll
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

	iprotoserver.usecase.ChangeState(internal.ReadWrite)

	returnBody, err := msgpack.Marshal(nil)

	if err != nil {
		return returnBody, fmt.Errorf(errTemplate, err)
	}

	return returnBody, nil
}

func ADM_STORAGE_SWITCH_MAINTENANCE(iprotoserver *IprotoServer, body []byte) ([]byte, error) { //nolint: golint,revive,nosnakecase,stylecheck,lll
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

	iprotoserver.usecase.ChangeState(internal.Maintenance)

	returnBody, err := msgpack.Marshal(nil)

	if err != nil {
		return returnBody, fmt.Errorf(errTemplate, err)
	}

	return returnBody, nil
}
