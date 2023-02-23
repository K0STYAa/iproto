package models

import (
	"net"
	"net/rpc"

	"github.com/vmihailenco/msgpack/v5"
)

type MsgPackRpcServerCodec struct {
	rwc net.Conn
	dec *msgpack.Decoder
	enc *msgpack.Encoder
}

func (c *MsgPackRpcServerCodec) ReadRequestHeader(r *rpc.Request) error {
	return c.dec.Decode(&r)
}

func (c *MsgPackRpcServerCodec) ReadRequestBody(body interface{}) error {
	return c.dec.Decode(body)
}

func (c *MsgPackRpcServerCodec) WriteResponse(r *rpc.Response, body interface{}) error {
	if err := c.enc.Encode(r); err != nil {
		return err
	}
	if err := c.enc.Encode(body); err != nil {
		return err
	}
	return nil
}

func (c *MsgPackRpcServerCodec) Close() error {
	return c.rwc.Close()
}

func GetServerCodec(conn net.Conn) *MsgPackRpcServerCodec{
	return &MsgPackRpcServerCodec{
		rwc: conn,
		dec: msgpack.NewDecoder(conn),
		enc: msgpack.NewEncoder(conn),
	}
}