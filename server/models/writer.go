package models

import "github.com/graphql-go/graphql"

type Writer struct {
  ID int `json:"id"`
  FirstName string `json:"firstName"`
  LastName string `json:"lastName"`
  ProfileUrl string `json:"profileUrl"`
}

var WriterType = graphql.NewObject(graphql.ObjectConfig{
  Name: "Writer",
  Fields: graphql.Fields {
    "ID": &graphql.Field{
      Type: graphql.Int,
    },
    "FirstName": &graphql.Field{
      Type: graphql.String,
    },
    "LastName": &graphql.Field{
      Type: graphql.String,
    },
    "ProfileUrl": &graphql.Field{
      Type: graphql.String,
    },
  },
})
