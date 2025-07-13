package renderer

import "sapphire/internal/file"

type Repository struct {
	*Artifact
	Categories []string     `json:"categories"`
	Artifacts  *[]IArtifact `json:"artifacts,omitempty"`
	Files      *[](*File)   `json:"files,omitempty"`
	Links      *[](*Link)   `json:"links,omitempty"`
}

func MakeRepository(
	name string,
	location string,
	path string,
	content string,
	metadata *file.Metadata,
	categories []string,
	artifacts *[]IArtifact,
	files *[](*File),
	links *[](*Link),
) *Repository {
	a := MakeArtifact(name, location, path, content, metadata)
	if artifacts == nil {
		defaultArtifacts := make([]IArtifact, 0)
		artifacts = &defaultArtifacts
	}
	if files == nil {
		defaultFiles := make([](*File), 0)
		files = &defaultFiles
	}
	if links == nil {
		defaultLinks := make([](*Link), 0)
		links = &defaultLinks
	}
	return &Repository{
		a,
		categories,
		artifacts,
		files,
		links,
	}
}

func (r Repository) GetArtifacts() *[]IArtifact {
	return r.Artifacts
}
func (r Repository) GetFiles() *[](*File) {
	return r.Files
}
func (r Repository) GetLinks() *[](*Link) {
	return r.Links
}
