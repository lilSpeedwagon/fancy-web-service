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

	codeSuccessOk     = 200
	codeErrorInternal = 500
	codeErrorNotImpl  = 501
)

func setHandlers() {
	logMsg("configuring handlers...")
	http.HandleFunc(requestPut, handlePut)
	http.HandleFunc(requestRemove, handleRemove)
	http.HandleFunc(requestGet, handleGet)
	http.Handle(requestBase, http.FileServer(http.Dir(indexPage)))
}

func getDataBase() database.IDataBase {
	db, err := database.OpenDataBase(dataBaseUrl)
	if err != nil {
		logMsg(err.Error())
		return nil
	}
	return db
}

func handleError(w http.ResponseWriter, err string, code int) {
	logMsg(err)
	http.Error(w, err, code)
}

func handlePut(writer http.ResponseWriter, request *http.Request) {
	logMsg("Put request: " + request.RequestURI)

	entriesInserted := 0
	db := getDataBase()

	switch request.Method {
	case "GET":
		logMsg("Processing GET request")
		for k, v := range request.URL.Query() {
			logMsg("Key: " + k + ". Value: " + v[0])

			isInserted, err := db.Put(k, v[0])
			if err != nil {
				handleError(writer, err.Error(), codeErrorInternal)
			}
			if isInserted {
				entriesInserted++
			}
		}
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		/*if err := request.ParseForm(); err != nil {
			fmt.Fprintf(writer, "ParseForm() err: %v", err)
			return
		}
		fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
		name := r.FormValue("name")
		address := r.FormValue("address")
		fmt.Fprintf(w, "Name = %s\n", name)
		fmt.Fprintf(w, "Address = %s\n", address)*/
	default:
		handleError(writer, "Unsupported request type", codeErrorInternal)
		return
	}

	_, err := fmt.Fprintf(writer, "Put done. Entries inserted: %d.", entriesInserted)
	if err != nil {
		handleError(writer, err.Error(), codeErrorInternal)
	}
}

func handleRemove(writer http.ResponseWriter, request *http.Request) {
	logMsg("Remove request: " + request.RequestURI)

	entriesRemoved := 0
	db := getDataBase()

	switch request.Method {
	case "GET":
		logMsg("Processing GET request")
		key := request.URL.Query().Get("key")
		logMsg("Key: " + key)

		isRemoved, err := db.Remove(key)
		if err != nil {
			handleError(writer, err.Error(), codeErrorInternal)
			return
		}
		if isRemoved {
			entriesRemoved++
		}
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		/*if err := request.ParseForm(); err != nil {
			fmt.Fprintf(writer, "ParseForm() err: %v", err)
			return
		}
		fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
		name := r.FormValue("name")
		address := r.FormValue("address")
		fmt.Fprintf(w, "Name = %s\n", name)
		fmt.Fprintf(w, "Address = %s\n", address)*/
	default:
		handleError(writer, "Unsupported request type", codeErrorInternal)
		return
	}

	_, err := fmt.Fprintf(writer, "Remove done. Entries removed: %d.", entriesRemoved)
	if err != nil {
		handleError(writer, err.Error(), codeErrorInternal)
	}
}

func handleGet(writer http.ResponseWriter, request *http.Request) {
	logMsg("Get request: " + request.RequestURI)

	db := getDataBase()
	var value string

	switch request.Method {
	case "GET":
		logMsg("Processing GET request")
		key := request.URL.Query().Get("key")
		logMsg("Key: " + key)

		v, err := db.Read(key)
		if err != nil {
			handleError(writer, err.Error(), codeErrorInternal)
			return
		} else {
			value = v
		}
	case "POST":
	default:
		handleError(writer, "Unsupported request type", codeErrorInternal)
		return
	}

	logMsg("Value: " + value)
	_, err := fmt.Fprintf(writer, "Get done. Value: %s.", value)
	if err != nil {
		handleError(writer, err.Error(), codeErrorInternal)
	}
}
