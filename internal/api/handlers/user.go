package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ajaysehwal/go-rest-server/internal/services"
	"github.com/gorilla/mux"
)



func GetUserHandler(svc * services.UserService) http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request){
		vars:=mux.Vars(r)
		id:=vars["id"];
		user, err :=svc.GetUserServices(id)
		if err !=nil{
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(APIError{Message: "User not found"})
			return
		}
		json.NewEncoder(w).Encode(user)
	}
}