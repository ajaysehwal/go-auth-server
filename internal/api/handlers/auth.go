package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ajaysehwal/go-rest-server/internal/services"
)

type Credentials struct{
	Email string `json:"email"`
	Password string `json:"password"`
}



func RegisterHandler(svc *services.AuthService) http.HandlerFunc {
   return func (w http.ResponseWriter, r *http.Request){
	  w.Header().Set("Content-Type", "application/json")
      var creds Credentials
	 if err:=json.NewDecoder(r.Body).Decode(&creds);err!=nil{
		w.WriteHeader(http.StatusBadRequest);
		json.NewEncoder(w).Encode(APIError{Message:"Invaild request"})
		return;
	 }
	 println(creds.Email,creds.Password)
	 userId,err:= svc.Register(creds.Email,creds.Password);
	 if err !=nil{
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(APIError{Message:err.Error()})
		return;
	 }
	 json.NewEncoder(w).Encode(map[string]int{"id":userId})
   }
}

func LoginHandler(svc *services.AuthService) http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")
        var creds Credentials
		if err:=json.NewDecoder(r.Body).Decode(&creds); err !=nil{
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(APIError{Message: "email or password is incorrect"})
			return;
		}
		token, err:=svc.Login(creds.Email,creds.Password);
		if err != nil{
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(APIError{Message:err.Error()})
			return;
		}
		json.NewEncoder(w).Encode(map[string]string{"jwt":token})
	}
}