interface Metadata {
  date: string;
  slug: string;
  categories: string[];
  title: string;
  summary: string;
  index?: number;
  [key: string]: any;
}

interface Link {
  name: string;
  path: string;
  summary: string;
  categories?: string[];
  items?: Link[];
}

interface Artifact {
  name: string;
  summary: string;
  path: string;
  content: string;
  categories?: string[];
  date: string;
  dateEdited?: string;
  files?: any[];
  links?: Link[];
  prev?: Link;
  next?: Link;
}

// For the root structure (artifacts.json)
interface ArtifactsSchema {
  [path: string]: Artifact;
}

interface AdjacentLink {
  path: string;
  name: string;
}

export {
  type Metadata,
  type Link,
  type Artifact,
  type ArtifactsSchema,
  type AdjacentLink,
}
