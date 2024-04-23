import { onMount } from "svelte";

const GRAPQL_ENDPOINT: string = "https://localhost:8080/graphql";

const getAllArticles = async () => {
  let articles;
  let query = `
  {
    getAllArticles {
      ID
      Content
      Title
      Length
      DateUploaded
      ReadCount
      ThumbnailUrl
      Writer {
        ID
        FirstName
        LastName
        ProfileUrl
      }
    }
  }`;

  const response = await fetch(GRAPQL_ENDPOINT, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ query }),
  });
  articles = response.data;
};

const getArticleByWriter = async (writerId: number) => {
  let articles;
  let query = `
  {
    getArticlesByWriter(writerId: ${writerId}) {
      ID
      Title
      ThumbnailUrl
      Content
      Length
      DateUploaded
      ReadCount
      Writer {
        ID
        FirstName
        LastName
        ProfileUrl
      }
    }
  };
`;

  const response = await fetch(GRAPQL_ENDPOINT, {
    method: "POST",
    headers: { 'Content-Type': "application/json" },
    body: JSON.stringify({query}),
  });
  articles = response.data;
};
