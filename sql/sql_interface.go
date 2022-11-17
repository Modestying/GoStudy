package sql

type SQL interface {
	ConnectDB() error
	DisconnectDB() error
}
