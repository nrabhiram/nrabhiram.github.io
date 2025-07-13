package file

type Artifact struct {
	*File
	*Metadata
	Content string
}

func MakeArtifact(repo *Repository, index uint, name string, location string, content string, metadata *Metadata) *Artifact {
	file := MakeFile(repo, index, name, location)
	if file == nil {
		return nil
	}
	newArtifact := Artifact{
		file,
		metadata,
		content,
	}
	return &newArtifact
}

func (a *Artifact) Path() string {
	pedigree := a.GetPedigree()
	var slug string
	if a.Name != "index" {
		slug = a.Metadata.Slug
		if slug == "" {
			slug = a.Name
		}
		slug = "/" + slug
	}
	return pedigree + slug
}
