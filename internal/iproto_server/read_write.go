package iprotoserver

import (
	"fmt"

	"github.com/K0STYAa/vk_iproto/pkg/iproto"
	"github.com/sirupsen/logrus"
	"github.com/vmihailenco/msgpack/v5"
)

// swiftlint:disable:next line_length
func STORAGE_READ(iprotoserver *IprotoServer, body []byte) ([]byte, error) { //nolint: lll,golint,revive,nosnakecase,stylecheck
	var req iproto.ReqReadArgs

	err := msgpack.Unmarshal(body, &req)
	if err != nil {
		return nil, fmt.Errorf(errTemplate, err)
	}
	// LOG REQUEST
	logrus.Info(fmt.Sprintf("[REQUEST]: STORAGE_READ(%v)", req))

	resp, err := iprotoserver.usecase.Read(req)

	if err != nil {
		return nil, fmt.Errorf(errTemplate, err)
	}

	responseByte, err := msgpack.Marshal(resp)
	if err != nil {
		return nil, fmt.Errorf(errTemplate, err)
	}

	return responseByte, nil
}

// swiftlint:disable:next line_length
func STORAGE_REPLACE(iprotoserver *IprotoServer, body []byte) ([]byte, error) { //nolint: lll,golint,revive,nosnakecase,stylecheck
	var req iproto.ReqReplaceArgs

	err := msgpack.Unmarshal(body, &req)
	if err != nil {
		return nil, fmt.Errorf(errTemplate, err)
	}
	// LOG REQUEST
	logrus.Info(fmt.Sprintf("[REQUEST]: STORAGE_REPLACE(%v)", req))

	resp, err := iprotoserver.usecase.Replace(req)
	if err != nil {
		return nil, fmt.Errorf(errTemplate, err)
	}

	responseByte, err := msgpack.Marshal(resp)
	if err != nil {
		return nil, fmt.Errorf(errTemplate, err)
	}

	return responseByte, nil
}
