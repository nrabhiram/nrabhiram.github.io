interface Link {
  name: string;
  path: string;
  summary?: string;
  categories?: string[];
  items?: Link[];
}

interface Artifact {
  name: string;
  summary: string;
  content?: string;
  path: string;
  categories?: string[];
  date: string;
  dateEdited?: string;
  files?: any[];
  prev?: Link;
  next?: Link;
}

interface AdjacentLink {
  path: string;
  name: string;
}

export {
  type Link,
  type Artifact,
  type AdjacentLink,
}
