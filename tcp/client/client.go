package client

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"strings"

	"../../common"
	"github.com/fatih/pool"
)

var (
	client ITcpClient
)

type ITcpClient interface {
	SendLoginRequest(request *common.LoginRequest) (string, error)
}

type TcpClientImpl struct {
	Pool pool.Pool
}

func GetClient() ITcpClient {
	if client == nil {
		InitClient()
	}
	return client
}

func InitClient() {
	connPool, err := pool.NewChannelPool(common.TcpInitialConnections, common.TcpMaxConnections, connectionFactory)
	if err != nil {
		log.Fatal(err)
	}
	client = &TcpClientImpl{
		Pool: connPool,
	}
}

func connectionFactory() (net.Conn, error) {
	return net.Dial("tcp", ":"+common.TcpPort)
}

func (m *TcpClientImpl) SendLoginRequest(request *common.LoginRequest) (string, error) {
	fmt.Printf("Current TCP open connections %d\n", m.Pool.Len())

	conn, err := m.Pool.Get()
	if err != nil {
		log.Println("Error in gettting available connection!")
	}
	defer conn.Close()

	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	enc := json.NewEncoder(rw)
	err = enc.Encode(request)
	//rw.WriteString(request.Username + " " + request.Password + common.TcpMsgDelimiterStr)
	if err != nil {
		log.Println(err)
		return "", err
	}

	err = rw.Flush()
	if err != nil {
		log.Println(err)
		return "", err
	}

	response, err := rw.ReadString(common.TcpMsgDelimiterByte)
	if err != nil {
		log.Println(err)
		return "", err
	}

	respArr := strings.Split(response, " ")
	if respArr[0] != "UserID" || len(respArr) != 2 {
		return "", errors.New(response)
	}

	return respArr[1], nil
}
