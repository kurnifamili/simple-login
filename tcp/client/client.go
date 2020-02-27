package client

import (
	"bufio"
	"errors"
	"log"
	"net"
	"strings"

	connPool "../../conn-pool"
	"../../common"
)

var (
	client ITcpClient
)

type ITcpClient interface {
	SendRequest(request string) (response string, err error)
}

type tcpClientImpl struct {
	//Pool pool.Pool
	Pool connPool.Pool
}

func GetClient() ITcpClient {
	if client == nil {
		InitClient()
	}
	return client
}

func InitClient() {
	//connPool, err := pool.NewChannelPool(common.TcpInitialConnections, common.TcpMaxConnections, connectionFactory)
	connPoolConfig := &connPool.PoolConfig{
		InitialConns:        common.TcpInitialConnections,
		MaxOpenConns:        common.TcpMaxConnections,
		AllowNewConnOverMax: true,
	}
	connPool, err := connPool.NewPool(connPoolConfig, connectionFactory)
	if err != nil {
		log.Fatal(err)
	}
	client = &tcpClientImpl{
		Pool: connPool,
	}
}

func connectionFactory() (net.Conn, error) {
	return net.Dial("tcp", ":"+common.TcpPort)
}

func (m *tcpClientImpl) SendRequest(request string) (response string, err error) {
	//fmt.Printf("Current TCP open connections %d\n", m.Pool.Len())
	//fmt.Printf("Current TCP open connections %d\n", m.Pool.GetOpenConnsLength())

	//conn, err := m.Pool.Get()
	conn, err := m.Pool.GetConn()
	if err != nil {
		log.Println("Error in gettting available connection!")
		return "", err
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
