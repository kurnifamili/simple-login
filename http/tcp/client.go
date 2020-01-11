package tcp

import (
	"bufio"
	"net"
	"log"
	"../common"
	"encoding/json"
	"strings"
	"errors"
)

const (
	TcpPort = "8091"
)
type ITcpClient interface {
	SendLoginRequest(request *common.LoginRequest) (string, error)
}

type TcpClientImpl struct {
}

func NewTcpClient() ITcpClient {
	return &TcpClientImpl{
	}
}

func (m *TcpClientImpl) SendLoginRequest(request *common.LoginRequest) (string, error) {
	conn, err := net.Dial("tcp", ":"+TcpPort)
	if err != nil {
		log.Fatal(err)
	}

	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	enc := json.NewEncoder(rw)
	err = enc.Encode(request)
	if err != nil {
		return "", err
	}

	err = rw.Flush()
	if err != nil {
		return "", err
	}

	response, err := rw.ReadString('\n')
	if err != nil {
		//fmt.Println("Error in reading reply from TCP server, err:", err.Error())
		return "", err
	}

	respArr := strings.Split(response, " ")
	if respArr[0] != "UserID" || len(respArr) != 2 {
		return "", errors.New(response)
	}

	return respArr[1], nil
}

//func Open() (*bufio.ReadWriter, error) {
//	conn, err := net.Dial("tcp", TcpPort)
//	if err != nil {
//		return nil, errors.Wrap(err, "Dialing "+TcpPort+" failed")
//	}
//	return bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn)), nil
//}