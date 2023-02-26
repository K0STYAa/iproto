package delivery

import (
	"fmt"

	"github.com/K0STYAa/vk_iproto/pkg/models"
	"github.com/sirupsen/logrus"
	"github.com/vmihailenco/msgpack/v5"
)

func STORAGE_READ(d *Delivery, body []byte) ([]byte, error) {
	var req models.ReqReadArgs

	err := msgpack.Unmarshal(body, &req)
	if err != nil {
		return nil, err
	}
	// LOG REQUEST
	logrus.Info(fmt.Sprintf("[REQUEST]: STORAGE_READ(%v)", req))

	resp, err := d.service.Read(req)
	if err != nil {
		return nil, err
	}

	b, err := msgpack.Marshal(resp)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func STORAGE_REPLACE(d *Delivery, body []byte) ([]byte, error) {
	var req models.ReqReplaceArgs

	err := msgpack.Unmarshal(body, &req)
	if err != nil {
		return nil, err
	}
	// LOG REQUEST
	logrus.Info(fmt.Sprintf("[REQUEST]: STORAGE_REPLACE(%v)", req))

	resp, err := d.service.Replace(req)
	if err != nil {
		return nil, err
	}

	b, err := msgpack.Marshal(resp)
	if err != nil {
		return nil, err
	}

	return b, nil
}
