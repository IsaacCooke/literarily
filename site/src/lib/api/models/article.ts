import Writer from "./writer";

class Article {
  id: number;
  content: string;
  title: string;
  length: number;
  dateUploaded: Date;
  readCount: number;
  thumbnailUrl: string;
  writer: Writer;
}
