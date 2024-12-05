package handler

import (
	"encoding/json"
	"net/http"
)

//Устанавливает заголовок Content-Type для ответа, указывая, что тело ответа будет в формате JSON. 
//Это нужно для того, чтобы клиент (например, браузер или Insomnia/Postman) понимал, что пришел JSON, и мог его корректно обработать.
func jsonResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
