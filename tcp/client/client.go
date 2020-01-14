package client

import (
	"bufio"
	"net"
	"log"
	"../../common"
	"encoding/json"
	"strings"
	"errors"
)

const (
	TcpPort = "8091"
)

var (
	client ITcpClient
)
type ITcpClient interface {
	SendLoginRequest(request *common.LoginRequest) (string, error)
}

type TcpClientImpl struct {
	Client *bufio.ReadWriter
}

func GetClient() ITcpClient {
	if client == nil {
		InitClient()
	}
	return client
}

func InitClient() {
	conn, err := net.Dial("tcp", ":"+TcpPort)
	if err != nil {
		log.Fatal(err)
	}
	client =  &TcpClientImpl{
		Client: bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn)),
	}
}

func (m *TcpClientImpl) SendLoginRequest(request *common.LoginRequest) (string, error) {
	conn, err := net.Dial("tcp", ":"+TcpPort)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

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