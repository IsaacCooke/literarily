import { onMount } from "svelte";

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

  const response = await fetch("https://localhost:8080/graphql", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ query }),
  });
  articles = response.data;
};
