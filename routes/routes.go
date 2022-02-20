package routes

import (
	"Key_Value_Storage/controllers"
	"net/http"
)

func Setup() {
	http.HandleFunc("/api/get-key-value", controllers.GetKeyValue)
	http.HandleFunc("/api/set-key-value", controllers.SetKeyValue)
	http.HandleFunc("/api/get-all-key-value", controllers.GetAllKeyValue)
	http.HandleFunc("/api/save-all-key-value", controllers.FlushAllData)
}
