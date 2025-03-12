package handlers

type Credentials struct{
	Email string `json: "email"`
	Password string `json:"password"`
}



func RegisterHandler()