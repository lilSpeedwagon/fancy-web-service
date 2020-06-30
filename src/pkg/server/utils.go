package server

const (
	serverAddress = ""
	serverPort    = "80"
)

func getServerAddress() string {
	return serverAddress + ":" + serverPort
}

func validateDbUrl(url string) bool {
	if len(url) == 0 {
		return false
	}

	return true
}
