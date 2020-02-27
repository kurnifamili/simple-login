package common

import (
	"strings"
	"strconv"
	"log"
)

func TrimSuffix(s string, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		s = s[:len(s)-len(suffix)]
	}
	return s
}

func ParseStrToTcpRequestType(s string) TcpRequestType {
	i, err := strconv.ParseInt(s, 10, 16)
	if err != nil {
		log.Printf(err.Error())
	}
	return TcpRequestType(i)
}
