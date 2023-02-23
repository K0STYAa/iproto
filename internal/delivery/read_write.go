package delivery

import (
	"github.com/vmihailenco/msgpack/v5"
	"github.com/K0STYAa/vk_iproto/pkg/models"
)

func STORAGE_READ(d *Delivery, body []byte) ([]byte, error) {
	var req models.ReqReadArgs

	err := msgpack.Unmarshal(body, &req)
    if err != nil {
        return nil, err
    }

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