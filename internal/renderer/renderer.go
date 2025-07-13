package renderer

import (
	"fmt"
	"log"
	"regexp"
	"sapphire/internal/file"
	"sort"
	"strings"
	"time"
)

func renderArtifact(a *file.Artifact) *Artifact {
	name := formatArtifactName(a.Name, a.Metadata)
	path := formatArtifactPath(a.Path())
	renderedArtifact := MakeArtifact(name, a.Location, path, a.Content, a.Metadata)
	return renderedArtifact
}

func renderFile(f *file.File) *File {
	renderedFile := MakeFile(f.Name, f.Location, f.Path())
	return renderedFile
}

func RenderRepository(repo *file.Repository) *Repository {
	children := repo.Children()
	sortChildren(children)
	artifacts := make([]IArtifact, 0)
	files := make([]*File, 0)
	links := make([]*Link, 0)

	for i := 0; i < len(*children); i++ {
		f := (*children)[i]
		switch v := f.(type) {
		case *file.Artifact:
			if f.GetName() != "index" {
				a := renderArtifact(v)
				artifacts = append(artifacts, a)
				isLinkLive := v.Content != "" // add other conditions for which the link isn't live
				link := MakeLink(a.Path, a.Name, isLinkLive, a.Metadata, nil)
				links = append(links, link)
			}
		case *file.Repository:
			r := RenderRepository(v)
			artifacts = append(artifacts, r)
			containsLinks := r.Links != nil && len(*r.Links) > 0
			isLinkLive := r.Content != "" // add other conditions for which the link isn't live
			var link *Link
			if containsLinks {
				link = MakeLink(r.Path, r.Name, isLinkLive, r.Metadata, r.Links)
			} else {
				link = MakeLink(r.Path, r.Name, isLinkLive, r.Metadata, nil)
			}
			links = append(links, link)
		case *file.File:
			renderedFile := renderFile(v)
			files = append(files, renderedFile)
		}
	}

	content := repo.Content()
	metadata := repo.Metadata()
	categories := repo.Categories
	location := repo.Location
	path := repo.Path()
	path = formatArtifactPath(path)
	name := formatArtifactName(repo.Name, metadata)

	renderedRepo := MakeRepository(
		name,
		location,
		path,
		content,
		metadata,
		categories,
		&artifacts,
		&files,
		&links,
	)

	return renderedRepo
}

func getIndex(f file.IFile) *uint {
	switch v := f.(type) {
	case *file.Artifact:
		if v.Metadata != nil {
			return &v.Metadata.Index
		}
	case *file.Repository:
		metadata := v.Metadata()
		if metadata != nil {
			return &metadata.Index
		}
	}
	return nil
}

func getDate(f file.IFile) *time.Time {
	switch v := f.(type) {
	case *file.Artifact:
		if v.Metadata != nil {
			return &v.Metadata.Date
		}
	case *file.Repository:
		metadata := v.Metadata()
		if metadata != nil {
			return &metadata.Date
		}
	}
	return nil
}

func getTitleOrName(f file.IFile) *string {
	switch v := f.(type) {
	case *file.Artifact:
		if v.Metadata != nil && v.Metadata.Title != "" {
			return &v.Metadata.Title
		}
		return &v.Name
	case *file.Repository:
		metadata := v.Metadata()
		if metadata != nil && metadata.Title != "" {
			return &metadata.Title
		}
		return &v.Name
	}
	name := f.GetName()
	return &name
}

func sortChildren(children *[]file.IFile) {
	sort.SliceStable(*children, func(i, j int) bool {
		a, b := (*children)[i], (*children)[j]

		aIndex, bIndex := getIndex(a), getIndex(b)
		indicesExist := aIndex != nil && bIndex != nil
		firstIndexExists, secondIndexExists := aIndex != nil, bIndex != nil

		if indicesExist {
			canSortByIndex := *aIndex != *bIndex
			if canSortByIndex {
				return *aIndex < *bIndex
			}
		} else if firstIndexExists {
			return true
		} else if secondIndexExists {
			return false
		}

		aDate, bDate := getDate(a), getDate(b)
		datesExist := aDate != nil && bDate != nil

		if datesExist && !aDate.Equal(*bDate) {
			return aDate.After(*bDate)
		}

		aName, bName := getTitleOrName(a), getTitleOrName(b)
		return *aName < *bName
	})
}

func formatArtifactPath(path string) string {
	re := regexp.MustCompile(`\s+`)
	path = re.ReplaceAllString(path, "-")
	return path
}

func formatArtifactName(artifactName string, metadata *file.Metadata) string {
	var name string
	if metadata != nil && metadata.Title != "" {
		name = metadata.Title
	} else {
		name = strings.ReplaceAll(artifactName, "-", " ")
	}
	return name
}

// the following functions are used for processing inputs and logging outputs in the CLI

func LogError(errMsg string) {
	log.Fatal(errMsg)
}

func LogMessage(msg string) {
	fmt.Println(msg)
}
