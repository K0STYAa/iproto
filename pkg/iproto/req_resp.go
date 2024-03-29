package iproto

type Header struct {
	FuncID     uint32
	BodyLength uint32
	RequestID  uint32
}

type Request struct {
	Header Header
	Body   []byte
}

type Response struct {
	Header     Header
	ReturnCode uint32
	Body       []byte
}

// Arguments structure.
type ReqReplaceArgs struct {
	ID int
	S  string
}

type RespReplaceArgs struct{}

type ReqReadArgs struct {
	ID int
}

type RespReadArgs struct {
	S string
}
