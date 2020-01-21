package common

import "strconv"

type TcpRequestType int16
type TcpResponseType int16

const (
	TcpLoginRequest TcpRequestType = 0
)

func (t TcpRequestType) ToInt() int {
	return int(t)
}

func (t TcpRequestType) ToByte() byte {
	return byte(t)
}

func (t TcpRequestType) ToString() string {
	return strconv.Itoa(t.ToInt())
}

const (
	TcpSuccessResponse TcpResponseType = 0
	TcpErrorResponse TcpResponseType = 1
)

func (t TcpResponseType) ToInt() int {
	return int(t)
}

func (t TcpResponseType) ToByte() byte {
	return byte(t)
}

func (t TcpResponseType) ToString() string {
	return strconv.Itoa(t.ToInt())
}