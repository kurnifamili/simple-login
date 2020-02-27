package common

const (
	HttpPort = "8090"

	TcpPort               = "8091"
	TcpMsgDelimiterStr    = "\n"
	TcpMsgDelimiterByte   = '\n'
	TcpMsgSeparator       = " "
	TcpInitialConnections = 10
	TcpMaxConnections     = 50

	DBMaxOpenConnections    = 150
	DBMaxIdleConnections    = 100
	DBConnMaxLifetimeInSecs = 60
)
