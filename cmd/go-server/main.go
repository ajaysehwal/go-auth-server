package main

import (
	"log"
	"net/http"
	"time"

	"github.com/ajaysehwal/go-rest-server/internal/api"
	"github.com/ajaysehwal/go-rest-server/internal/config"
	"github.com/ajaysehwal/go-rest-server/internal/db"
)


func main(){
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
   router:=api.SetupRouter(dbConn.DB)
  
   log.Println("Server starting on : 8080")
   if err := http.ListenAndServe(":8080",router); err !=nil{
	log.Fatalf("server failed %s",err)
   }
	
}