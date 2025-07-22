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
  castHash?: string;
}

interface AdjacentLink {
  path: string;
  name: string;
}

type Reply = {
  hash: string;
  text: string;
  username: string;
  displayName: string;
  pfp: string;
  formattedDate: string;
  replies: Reply[];
}

type Reactions = {
  likes?: number;
  recasts?: number;
  comments?: number;
}

export {
  type Link,
  type Artifact,
  type AdjacentLink,
  type Reply,
  type Reactions,
}
