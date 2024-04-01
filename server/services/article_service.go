package services

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/IsaacCooke/literarily/data"
	"github.com/IsaacCooke/literarily/models"
	"github.com/graphql-go/graphql"
)

var getAllArticles = &graphql.Field {
  Type: graphql.NewList(models.ArticleType),
  Resolve: func(p graphql.ResolveParams) (interface{}, error) {
    db := data.Connect()
    defer db.Close()

    rows, err := db.Query("SELECT * FROM article JOIN writer ON article.writer_id = writer.id;")
    if err != nil {
      panic(err)
    }

    var articles []models.Article

    for rows.Next(){
      var id int
      var content string
      var title string
      var length int
      var dateUploaded time.Time
      var readCount int
      var thumbnailUrl string
      var writerID int
      var firstName string
      var lastName string
      var profileUrl string

      // TODO: Check if the columns match this data arrangement (very important!)
      err = rows.Scan(&id, &content, &title, &length, &dateUploaded, &readCount, &thumbnailUrl, &writerID, &writerID, &firstName, &lastName, &profileUrl)

      if err != nil {
        panic(err)
      }

      articles = append(articles, models.Article{
        ID: id,
        Content: content,
        Title: title,
        Length: length,
        DateUploaded: dateUploaded,
        ReadCount: readCount,
        ThumbnailUrl: thumbnailUrl,
        Writer: models.Writer{
          ID: writerID,
          FirstName: firstName,
          LastName: lastName,
          ProfileUrl: profileUrl,
        },
      })
    }

    if articles == nil {
      return nil, fmt.Errorf("No articles found")
    }

    return articles, nil
  },
}

var getArticleByTitle = &graphql.Field{
  Type: models.ArticleType,
  Args: graphql.FieldConfigArgument{
    "title": &graphql.ArgumentConfig{
      Type: graphql.String,
    },
  },
  Resolve: func(params graphql.ResolveParams) (interface{}, error){
    paramTitle := params.Args["title"].(string)

    db := data.Connect()
    defer db.Close()

    var article models.Article
    sqlStatement := `SELECT * FROM article JOIN writer ON article.writer_id = writer.id WHERE title = $1;`

    row := db.QueryRow(sqlStatement, paramTitle)

    err := row.Scan(
      &article.ID,
      &article.Content,
      &article.Title,
      &article.Length,
      &article.DateUploaded,
      &article.ReadCount,
      &article.ThumbnailUrl,
      &article.Writer.ID,
      &article.Writer.ID,
      &article.Writer.FirstName,
      &article.Writer.LastName,
      &article.Writer.ProfileUrl,
    )

    switch err {
    case sql.ErrNoRows:
      log.Printf("No articles with title %s found", paramTitle)
      return nil, err
    case nil:
      return article, nil
    default:
      e := fmt.Sprintf("error: %v occured while reading the database", err)
      log.Println(e)
      return nil, err
    }
  },
}
