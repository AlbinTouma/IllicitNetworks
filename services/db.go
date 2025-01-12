package db

import (
	//	"database/sql"
	//"fmt"
	"encoding/json"
	"kys/models"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

//http://go-database-sql.org/retrieving.html

func QueryEntity(query string) ([]db_models.Entity, error) {
  db, err := sqlx.Open("sqlite3", "./db/pb_data/data.db")
  if err != nil {
    log.Printf("Error in sqlx %s", err)
  }
  defer db.Close()
  var entities []db_models.Entity

  err = db.Select(&entities, query) 
  if err != nil {
    log.Printf("Error fetching entities from db %s", err)
    return nil, err
  }

  for i := range entities {
    var person db_models.Person
    err := json.Unmarshal([]byte(entities[i].Properties.(string)), &person)
    if err != nil {
      log.Printf("Error unmarshalling json", entities[i].Id, err)
      continue
    }
    entities[i].Properties = person
  }
  return entities, nil
}
