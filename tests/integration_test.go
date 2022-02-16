package tests

import (
	"Key_Value_Storage/controllers"
	"Key_Value_Storage/models"
	"Key_Value_Storage/routes"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func SetupApp() (*fiber.App, *models.Store) {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	routes.Setup(app)
	Db := models.CreateStore("HelloWorld.json", 5*time.Minute)
	controllers.Db = Db
	return app, Db
}

func TestSetKeyValue(t *testing.T) {
	app, Db := SetupApp()

	var pbody models.KeyValue
	pbody.Key = "TestKey"
	pbody.Value = "TestValue"
	body, _ := json.Marshal(pbody)
	req := httptest.NewRequest("POST", "/api/set-key-value", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	_, err := app.Test(req, 100)
	if err != nil {
		fmt.Println(err.Error())
	}
	assert.NotNil(t, Db.Get("TestKey"))
	assert.Equal(t, "TestValue", Db.Get("TestKey"))
}

func TestGetKeyValue(t *testing.T) {
	app, Db := SetupApp()

	Db.Set("TestKey", "TestValue")
	pBody := map[string]string{
		"key": "TestKey",
	}
	body, _ := json.Marshal(pBody)
	req := httptest.NewRequest("GET", "/api/get-key-value", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		fmt.Println(err.Error())
	}
	respBytes, _ := ioutil.ReadAll(resp.Body)
	respString := string(respBytes)
	respString = strings.ReplaceAll(respString, "\"", "")
	dbResp := Db.Get("TestKey")
	assert.NotNil(t, respString, "The response is Nil")
	assert.Equal(t, dbResp, respString, "The response from the DB and API does not match")
}

func TestGetAllKeyValue(t *testing.T) {
	app, Db := SetupApp()
	Db.Set("TestKey", "TestValue")
	Db.Set("TestKey1", "TestValue1")
	Db.Set("TestKey2", "TestValue2")

	req := httptest.NewRequest("GET", "/api/get-all-key-value", nil)

	resp, err := app.Test(req)
	if err != nil {
		fmt.Println(err.Error())
	}
	respBytes, _ := ioutil.ReadAll(resp.Body)
	respString := string(respBytes)
	respString = strings.ReplaceAll(respString, "\"", "")
	assert.NotNil(t, respString, "Response is Empty")

}
