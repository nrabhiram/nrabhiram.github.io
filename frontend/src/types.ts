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
  path: string;
  categories?: string[];
  date: string;
  dateEdited?: string;
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
