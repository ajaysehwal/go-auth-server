package db

import (
	"database/sql"
	"log"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/file"
)

type DB struct{
	*sql.DB
}

type Config struct{
	ConnString string
	MaxOpenConns int
	MaxIdleConns int
	MaxLifetime time.Duration
}

func ConnectDb(cfg Config) (*DB, error){
	db , err :=sql.Open("postgres",cfg.ConnString)
	if err !=nil{
		return nil, err
	}
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxOpenConns)
	db.SetConnMaxLifetime(cfg.MaxLifetime)

	if err := db.Ping();err!=nil{
		db.Close()
		return nil , err
	}
	log.Println("Connected to PostgreSql")
	return &DB{db}, nil
}

func (d *DB) Migrate(migrateDir string) error{
	source,err :=(&file.File{}).Open("file://"+migrateDir); 
     if err!=nil{
		return err
	 }

    driver, err:=postgres.WithInstance(d.DB,&postgres.Config{})
	if err != nil{
		return err
	}

   m, err:=migrate.NewWithInstance("file",source,"postgres",driver)

   if err !=nil{
	return err
   }
   if err:=m.Up(); err!=nil && err != migrate.ErrNoChange{
	return err
   }
    log.Println("Database migrations applied successfully")

 return nil
}

func (d *DB) Close() error{
	return d.DB.Close()
}