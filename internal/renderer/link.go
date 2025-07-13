package renderer

import "sapphire/internal/file"

type Link struct {
	Path     string         `json:"path"`
	Label    string         `json:"label"`
	Live     bool           `json:"live"`
	Metadata *file.Metadata `json:"metadata,omitempty"`
	Links    *[](*Link)     `json:"links,omitempty"`
}

func MakeLink(path string, label string, live bool, metadata *file.Metadata, links *[](*Link)) *Link {
	if links == nil {
		defaultLinks := make([](*Link), 0)
		links = &defaultLinks
	}
	link := Link{
		path,
		label,
		live,
		metadata,
		links,
	}
	return &link
}
