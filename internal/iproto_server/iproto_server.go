package iprotoserver

import (
	"github.com/K0STYAa/iproto/internal/storage"
	"github.com/K0STYAa/iproto/pkg/iproto"
	"github.com/sirupsen/logrus"
	"github.com/vmihailenco/msgpack/v5"
)

type IprotoServer struct {
	storage *storage.Storage
}

func NewIprotoServer(storage storage.Storage) *IprotoServer {
	return &IprotoServer{storage: &storage}
}

func (d *IprotoServer) MainHandler(req iproto.Request) iproto.Response {
	// Find Handler function by FuncID
	handlerFuncMap := map[uint32]func(*IprotoServer, []byte) ([]byte, error){
		0x00010001: ADM_STORAGE_SWITCH_READONLY,    //nolint: nosnakecase
		0x00010002: ADM_STORAGE_SWITCH_READWRITE,   //nolint: nosnakecase
		0x00010003: ADM_STORAGE_SWITCH_MAINTENANCE, //nolint: nosnakecase
		0x00020001: STORAGE_REPLACE,                //nolint: nosnakecase
		0x00020002: STORAGE_READ,                   //nolint: nosnakecase
	}
	iprotoserverFunc := handlerFuncMap[req.Header.FuncID]

	// Call handler for function
	resp, err := iprotoserverFunc(d, req.Body)

	var (
		returnCode uint32
		returnBody []byte
		bodyResp   iproto.RespReadArgs
	)

	if err != nil {
		returnCode = 1
		{ // LOG ERROR
			logrus.Warn("[ERROR]: ", err.Error())
		}

		returnBody, _ = msgpack.Marshal(err.Error())
	} else if err2 := msgpack.Unmarshal(resp, &bodyResp); err2 != nil {
		returnCode = 1
		{ // LOG ERROR
			logrus.Warn("[ERROR]: ", err2.Error())
		}
		returnBody, _ = msgpack.Marshal(err2.Error())
	} else {
		if bodyResp.S != "" { // LOG RESPONSE
			logrus.Info("[RESPONSE]: ", bodyResp.S)
		} else {
			logrus.Debug("[RESPONSE]: ", bodyResp.S)
		}
		returnBody = resp
	}

	result := iproto.Response{
		Header: iproto.Header{
			FuncID:     req.Header.FuncID,
			BodyLength: uint32(len(returnBody)),
			RequestID:  req.Header.RequestID,
		},
		ReturnCode: returnCode,
		Body:       returnBody,
	}

	return result
}
