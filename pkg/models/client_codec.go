package models

import (
	"log"
	"net"
	"net/rpc"

	"github.com/vmihailenco/msgpack/v5"
)

func GetClientCodec() *rpc.Client{
	client, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Dial error:", err)
	}

	rpcCodec := &MsgPackRpcClientCodec{
        conn: client,
        dec: msgpack.NewDecoder(client),
        enc: msgpack.NewEncoder(client),
    }
	return rpc.NewClientWithCodec(rpcCodec)
}

type MsgPackRpcClientCodec struct {
	conn net.Conn
	dec  *msgpack.Decoder
	enc  *msgpack.Encoder
}

func (c *MsgPackRpcClientCodec) WriteRequest(r *rpc.Request, body interface{}) error {
	if err := c.enc.Encode(r); err != nil {
		return err
	}
	if err := c.enc.Encode(body); err != nil {
		return err
	}
	return nil
}

func (c *MsgPackRpcClientCodec) ReadResponseHeader(r *rpc.Response) error {
	return c.dec.Decode(&r)
}

func (c *MsgPackRpcClientCodec) ReadResponseBody(body interface{}) error {
	return c.dec.Decode(body)
}

func (c *MsgPackRpcClientCodec) Close() error {
	return c.conn.Close()
}