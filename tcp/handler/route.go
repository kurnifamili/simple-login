package handler

import (
	"net"
	"bufio"
	"io"
	"log"
	"strings"

	"../../common"
	"errors"
)

func RouteHandler(c net.Conn) {
	for {
		reqStr, err :=bufio.NewReader(c).ReadString(common.TcpMsgDelimiterByte)
		//d := json.NewDecoder(c)
		//var req common.LoginRequest
		//err := d.Decode(&req)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Println("error in reading request", err)
			continue
		}

		respStr, err := handleRequest(reqStr)
		if err != nil {
			c.Write([]byte(common.TcpErrorResponse.ToString() + common.TcpMsgSeparator + err.Error() + common.TcpMsgDelimiterStr))
		} else {
			c.Write([]byte(common.TcpSuccessResponse.ToString() +  common.TcpMsgSeparator + respStr + common.TcpMsgDelimiterStr))
		}
	}

	c.Close()
}

func handleRequest(reqStr string) (respStr string, respErr error) {
	// Strip delimiter suffix
	reqStr = common.TrimSuffix(reqStr, common.TcpMsgDelimiterStr)
	reqArr := strings.Split(reqStr, common.TcpMsgSeparator)

	switch common.ParseStrToTcpRequestType(reqArr[0]) {
	case common.TcpLoginRequest:
		return loginHandler(reqArr[1:])
	default:
		return "", errors.New("no handler found")
	}
}