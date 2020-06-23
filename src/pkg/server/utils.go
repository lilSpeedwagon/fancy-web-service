package server

const (
	serverAddress = ""
	serverPort    = "80"
)

func getServerAddress() string {
	return serverAddress + ":" + serverPort
}
