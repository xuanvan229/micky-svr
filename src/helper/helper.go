package helper

import "net/http"

func SimpleLog(err error) {
	if err != nil {
		panic(err)
	}
	return
}

func Log(err error, w http.ResponseWriter) http.ResponseWriter {
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusInternalServerError)
	return w
}
