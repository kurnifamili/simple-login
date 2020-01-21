package common

import (
	"strconv"
	"encoding/binary"
)

type TcpRequestType uint16
type TcpResponseType uint16

const (
	TcpLoginRequest TcpRequestType = 0
)

func (t TcpRequestType) ToInt() int {
	return int(t)
}

func (t TcpRequestType) ToByte() byte {
	return byte(t)
}

func (t TcpRequestType) ToBytes() []byte {
	a := make([]byte, 4)
	binary.LittleEndian.PutUint16(a, uint16(t))
	return a
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