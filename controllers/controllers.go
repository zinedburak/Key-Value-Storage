package controllers

import (
	"Key_Value_Storage/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var Db *models.Store

func GetKeyValue(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		bodyBytes, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		ct := r.Header.Get("content-type")
		if ct != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			w.Write([]byte(fmt.Sprintf("Content type must be 'application/json, but you provide '%s'", ct)))
		}
		var cKey map[string]string
		if err := json.Unmarshal(bodyBytes, &cKey); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		jsonBytes, err := json.Marshal(Db.Get(cKey["key"]))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonBytes)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
	}
}

func SetKeyValue(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		bodyBytes, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		ct := r.Header.Get("content-type")
		if ct != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			w.Write([]byte(fmt.Sprintf("Content type must be 'application/json, but you provide '%s''", ct)))
		}
		var cKeyValue models.KeyValue
		if err := json.Unmarshal(bodyBytes, &cKeyValue); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		Db.Set(cKeyValue.Key, cKeyValue.Value)
		return

	}

}

func GetAllKeyValue(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		jsonBytes, err := json.Marshal(Db.GetAll())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonBytes)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
	}
}

func FlushAllData(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		Db.FlushAllData()
		jsonBytes, err := json.Marshal(Db.GetAll())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		w.Header().Add("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonBytes)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
	}
}
