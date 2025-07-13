package renderer

import "sapphire/internal/file"

type IArtifact interface {
	GetName() string
	GetLocation() string
	GetPath() string
	GetContent() string
	GetMetadata() *file.Metadata
	GetArtifacts() *[]IArtifact
	GetFiles() *[](*File)
	GetLinks() *[](*Link)
	GetNextLink() *Link
	GetPrevLink() *Link
	SetNextLink(link *Link)
	SetPrevLink(link *Link)
}

type Artifact struct {
	Name     string         `json:"name,omitempty"`
	Location string         `json:"location,omitempty"`
	Path     string         `json:"path,omitempty"`
	Content  string         `json:"content,omitempty"`
	Metadata *file.Metadata `json:"metadata,omitempty"`
	Next     *Link          `json:"next,omitempty"`
	Prev     *Link          `json:"prev,omitempty"`
}

func MakeArtifact(name string, location string, path string, content string, metadata *file.Metadata) *Artifact {
	artifact := Artifact{
		name,
		location,
		path,
		content,
		metadata,
		nil,
		nil,
	}
	return &artifact
}

func (a *Artifact) GetName() string {
	return a.Name
}
func (a *Artifact) GetLocation() string {
	return a.Location
}
func (a *Artifact) GetPath() string {
	return a.Path
}
func (a *Artifact) GetContent() string {
	return a.Content
}
func (a *Artifact) GetMetadata() *file.Metadata {
	return a.Metadata
}
func (a *Artifact) GetArtifacts() *[]IArtifact {
	return nil
}
func (a *Artifact) GetFiles() *[](*File) {
	return nil
}
func (a *Artifact) GetLinks() *[](*Link) {
	return nil
}
func (a *Artifact) GetNextLink() *Link {
	return a.Next
}
func (a *Artifact) GetPrevLink() *Link {
	return a.Prev
}
func (a *Artifact) SetNextLink(link *Link) {
	a.Next = link
}
func (a *Artifact) SetPrevLink(link *Link) {
	a.Prev = link
}
