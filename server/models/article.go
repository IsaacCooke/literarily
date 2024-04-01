package models

import (
	"time"

	"github.com/graphql-go/graphql"
)

type Article struct {
  ID int `json:"id"`
  Content string `json:"content"`
  Title string `json:"title"`
  Length int `json:"length"`
  DateUploaded time.Time `json:"date"`
  ReadCount int `json:"readCount"`
  ThumbnailUrl string `json:"thumbnailUrl"`
  Writer Writer `json:"writer"`
}

var ArticleType = graphql.NewObject(graphql.ObjectConfig{
  Name: "Article",
  Fields: graphql.Fields {
    "ID": &graphql.Field{
      Type: graphql.Int,
    },
    "Content": &graphql.Field{
      Type: graphql.String,
    },
    "Title": &graphql.Field{
      Type: graphql.String,
    },
    "Length": &graphql.Field{
      Type: graphql.Int,
    },
    "DateUploaded": &graphql.Field{
      Type: graphql.DateTime,
    },
    "ReadCount": &graphql.Field{
      Type: graphql.Int,
    },
    "ThumbnailUrl": &graphql.Field{
      Type: graphql.String,
    },
    "Writer": &graphql.Field{
      Type: WriterType,
    },
  },
})
