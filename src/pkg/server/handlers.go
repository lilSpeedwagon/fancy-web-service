package server

import (
	"fmt"
	"net/http"
	"pkg/database"
)

const (
	indexPage = "./page/"

	requestBase   = "/"
	requestPut    = "/put"
	requestRemove = "/remove"
	requestGet    = "/get"

	dataBaseUrl = ""
)

func setHandlers() {
	http.HandleFunc(requestPut, handlePut)
	http.HandleFunc(requestRemove, handleRemove)
	http.HandleFunc(requestGet, handleGet)
	http.Handle(requestBase, http.FileServer(http.Dir(indexPage)))
}

func getDataBase() database.IDataBase {
	db, err := database.OpenDataBase(dataBaseUrl)
	if err != nil {
		logf(err.Error())
		return nil
	}
	return db
}

func handleError(w http.ResponseWriter, err string, code int) {
	logf(err)
	http.Error(w, err, code)
}

func printToRequestBody(writer http.ResponseWriter, format string, args ...interface{}) {
	logf(format, args...)
	_, err := fmt.Fprintf(writer, format, args...)
	if err != nil {
		handleError(writer, err.Error(), http.StatusInternalServerError)
	}
}

func closeRequestBody(request *http.Request) {
	if err := request.Body.Close(); err != nil {
		logf(err.Error())
	}
}

func handlePut(writer http.ResponseWriter, request *http.Request) {
	defer closeRequestBody(request)

	logf("Put request: " + request.RequestURI)

	entriesInserted := 0
	db := getDataBase()

	switch request.Method {
	case "GET":
		logf("Processing GET request")
		for k, v := range request.URL.Query() {
			logf("Key: %s. Value = %s.", k, v[0])

			isInserted, err := db.Put(k, v[0])
			if err != nil {
				handleError(writer, err.Error(), http.StatusInternalServerError)
			}
			if isInserted {
				entriesInserted++
			}
		}
	case "POST":
		handleError(writer, "Not implemented", http.StatusNotImplemented)
		return
	default:
		handleError(writer, "Unsupported request type", http.StatusInternalServerError)
		return
	}

	printToRequestBody(writer, "Put done. Entries inserted %d.", entriesInserted)
}

func handleRemove(writer http.ResponseWriter, request *http.Request) {
	defer closeRequestBody(request)

	logf("Remove request: " + request.RequestURI)

	entriesRemoved := 0
	db := getDataBase()

	switch request.Method {
	case "GET":
		logf("Processing GET request")
		key := request.URL.Query().Get("key")
		logf("Key: %s.", key)

		isRemoved, err := db.Remove(key)
		if err != nil {
			handleError(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		if isRemoved {
			entriesRemoved++
		}
	case "POST":
		handleError(writer, "Not implemented", http.StatusNotImplemented)
		return
	default:
		handleError(writer, "Unsupported request type", http.StatusInternalServerError)
		return
	}

	printToRequestBody(writer, "Remove done. Entries removed: %d.", entriesRemoved)
}

func handleGet(writer http.ResponseWriter, request *http.Request) {
	defer closeRequestBody(request)

	logf("Get request: " + request.RequestURI)

	db := getDataBase()
	var value string

	switch request.Method {
	case "GET":
		logf("Processing GET request")
		key := request.URL.Query().Get("key")
		logf("Key: %s.", key)

		v, err := db.Read(key)
		if err != nil {
			handleError(writer, err.Error(), http.StatusInternalServerError)
			return
		} else {
			value = v
		}
	case "POST":
		handleError(writer, "Not implemented", http.StatusNotImplemented)
		return
	default:
		handleError(writer, "Unsupported request type", http.StatusInternalServerError)
		return
	}

	if len(value) != 0 {
		printToRequestBody(writer, "Get done. Value: %s.", value)
	} else {
		printToRequestBody(writer, "Get done. Value not found.")
	}
}
