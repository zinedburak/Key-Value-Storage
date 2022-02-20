package tests

import (
	"Key_Value_Storage/controllers"
	"Key_Value_Storage/models"
	"Key_Value_Storage/routes"
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestSetKeyValue(t *testing.T) {
	db := setupApp()
	var pBody models.KeyValue
	pBody.Key = "TestKey"
	pBody.Value = "TestValue"
	body, _ := json.Marshal(pBody)
	req := httptest.NewRequest("POST", "/api/set-key-value", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	controllers.SetKeyValue(w, req)
	assert.NotNil(t, db.Get("TestKey"))
	assert.Equal(t, "TestValue", db.Get("TestKey"))
}

func TestGetKeyValue(t *testing.T) {
	db := setupApp()
	db.Set("TestKey", "TestValue")
	pBody := map[string]string{
		"key": "TestKey",
	}
	body, _ := json.Marshal(pBody)
	req := httptest.NewRequest("GET", "/api/get-key-value", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	controllers.GetKeyValue(w, req)
	resp := w.Result()
	respBytes, _ := ioutil.ReadAll(resp.Body)
	respString := string(respBytes)
	respString = strings.ReplaceAll(respString, "\"", "")
	dbResp := db.Get("TestKey")
	assert.NotNil(t, respString, "The response is Nil")
	assert.Equal(t, dbResp, respString, "The response from the DB and API does not match")
}

func TestGetAllKeyValue(t *testing.T) {
	db := setupApp()
	db.Set("TestKey", "TestValue")
	db.Set("TestKey1", "TestValue1")
	db.Set("TestKey2", "TestValue2")

	req := httptest.NewRequest("GET", "/api/get-all-key-value", nil)
	w := httptest.NewRecorder()
	controllers.GetAllKeyValue(w, req)
	resp := w.Result()
	respBytes, _ := ioutil.ReadAll(resp.Body)
	respString := string(respBytes)
	respString = strings.ReplaceAll(respString, "\"", "")
	assert.NotNil(t, respString, "Response is Empty")

}

var routesCheck bool

func setupApp() *models.Store {
	if !routesCheck {
		routes.Setup()
		routesCheck = true
	}
	db := models.CreateStore("Testing.json", 5*time.Minute)
	controllers.Db = db
	return db
}
