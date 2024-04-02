package services

import (
	"fmt"
	"log"

	"github.com/IsaacCooke/literarily/data"
	"github.com/IsaacCooke/literarily/models"
	"github.com/graphql-go/graphql"
)

var getAllWriters = &graphql.Field {
  Type: graphql.NewList(models.WriterType),
  Resolve: func(p graphql.ResolveParams) (interface{}, error) {
    db := data.Connect()
    defer db.Close()

    rows, err := db.Query("SELECT * FROM writer ORDER BY first_name ASC")
    if err != nil {
      panic(err)
    }

    var writers []models.Writer

    for rows.Next(){
      var id int
      var firstName string
      var lastName string
      var profileUrl string

      err := rows.Scan(&id, &firstName, &lastName, &profileUrl)

      if err != nil {
        log.Panicln(err)
      }

      writers = append(writers, models.Writer {
        ID: id,
        FirstName: firstName,
        LastName: lastName,
        ProfileUrl: profileUrl,
      })
    }

    if writers == nil {
      return nil, fmt.Errorf("No writers found")
    }

    return writers, nil
  },
}
