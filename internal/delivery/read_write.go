package delivery

import (
	"fmt"

	"github.com/K0STYAa/vk_iproto/pkg/models"
	"github.com/sirupsen/logrus"
	"github.com/vmihailenco/msgpack/v5"
)

func STORAGE_READ(delivery *Delivery, body []byte) ([]byte, error) { //nolint: golint,revive,nosnakecase,stylecheck
	var req models.ReqReadArgs

	err := msgpack.Unmarshal(body, &req)
	if err != nil {
		return nil, fmt.Errorf(errTemplate, err)
	}
	// LOG REQUEST
	logrus.Info(fmt.Sprintf("[REQUEST]: STORAGE_READ(%v)", req))

	resp, err := delivery.service.Read(req)

	if err != nil {
		return nil, fmt.Errorf(errTemplate, err)
	}

	responseByte, err := msgpack.Marshal(resp)
	if err != nil {
		return nil, fmt.Errorf(errTemplate, err)
	}

	return responseByte, nil
}

func STORAGE_REPLACE(delivery *Delivery, body []byte) ([]byte, error) { //nolint: golint,revive,nosnakecase,stylecheck
	var req models.ReqReplaceArgs

	err := msgpack.Unmarshal(body, &req)
	if err != nil {
		return nil, fmt.Errorf(errTemplate, err)
	}
	// LOG REQUEST
	logrus.Info(fmt.Sprintf("[REQUEST]: STORAGE_REPLACE(%v)", req))

	resp, err := delivery.service.Replace(req)
	if err != nil {
		return nil, fmt.Errorf(errTemplate, err)
	}

	responseByte, err := msgpack.Marshal(resp)
	if err != nil {
		return nil, fmt.Errorf(errTemplate, err)
	}

	return responseByte, nil
}
