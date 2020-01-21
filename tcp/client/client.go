package client

import (
	"bufio"
	"errors"
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
	SendRequest(request string) (response string, err error)
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

func (m *TcpClientImpl) SendRequest(request string) (response string, err error) {
	//fmt.Printf("Current TCP open connections %d\n", m.Pool.Len())

	conn, err := m.Pool.Get()
	if err != nil {
		log.Println("Error in gettting available connection!")
	}
	defer conn.Close()

	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	_, err = rw.WriteString(request + common.TcpMsgDelimiterStr)
	if err != nil {
		log.Println(err)
		return "", err
	}

	err = rw.Flush()
	if err != nil {
		log.Println(err)
		return "", err
	}

	respStr, err := rw.ReadString(common.TcpMsgDelimiterByte)
	if err != nil {
		log.Println(err)
		return "", err
	}

	respStr = common.TrimSuffix(respStr, common.TcpMsgDelimiterStr)
	respArr := strings.Split(respStr, " ")
	if respArr[0] == common.TcpErrorResponse.ToString() {
		return "", errors.New(strings.Join(respArr[1:], " "))
	}

	return strings.Join(respArr[1:], " "), nil
}
