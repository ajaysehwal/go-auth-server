package main

import (
	"log"
	"net/http"
	"time"

	"github.com/ajaysehwal/go-rest-server/internal/config"
	"github.com/ajaysehwal/go-rest-server/internal/db"
	"github.com/gorilla/mux"
)


func main(){
   r:=mux.NewRouter()
   cfg:=config.LoadConfig()
   dbCfg:=db.Config{
	ConnString: cfg.DBConn,
	MaxOpenConns: 25,
	MaxIdleConns: 25,
	MaxLifetime: 5*time.Minute,
   }
   dbConn, err:= db.ConnectDb(dbCfg)
   if err !=nil{
	log.Fatalf("DB init Failed: %s" ,err)
   }
   defer dbConn.Close();

   if err:=dbConn.Migrate("./internal/db/migrations");err!=nil{
	log.Fatalf("Migration Failed : %v",err)
   }
   
   r.HandleFunc("/",homeHandler).Methods("GET")
   r.HandleFunc("users/{id}",UserHandler).Methods("GET")
   
   
   log.Println("Server starting on : 8080")
   if err := http.ListenAndServe(":8080",r); err !=nil{
	log.Fatalf("server failed %s",err)
   }
	
}

func homeHandler(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("Hello from go server"))
}

func UserHandler(w http.ResponseWriter, r *http.Request){
	vars:=mux.Vars(r);
	id:=vars["id"]
	w.Write([]byte("User ID: "+id))
   
}