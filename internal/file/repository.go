package file

import "fmt"

type Repository struct {
	*File
	Categories    []string
	CategoriesMap map[string]bool
	Primary       *Artifact
}

func MakeRepository(repo *Repository, index uint, name string, location string) *Repository {
	file := MakeFile(repo, index, name, location)
	if file == nil {
		return nil
	}
	newRepo := Repository{
		File:          file,
		Categories:    []string{},
		CategoriesMap: make(map[string]bool),
		Primary:       nil,
	}
	return &newRepo
}

func (r *Repository) Make(f IFile) error {
	children := r.Children()
	err := f.Acknowledge(r)
	if err != nil {
		return err
	}
	err = r.Add(f)
	for i := 0; i < len(*children); i++ {
		child := (*children)[i]
		err1 := child.Add(f)
		err2 := f.Add(child)
		if err1 != nil || err2 != nil || err != nil {
			if child.GetName() == f.GetName() {
				return fmt.Errorf("files in the repository named \"%s\" have the same name as \"%s\"", r.Name, child.GetName())
			} else {
				return fmt.Errorf("\"%s\" and \"%s\" can't be contained in the repository named \"%s\" because they have the same index", f.GetName(), child.GetName(), r.Name)
			}
		}
	}
	a, ok := f.(*Artifact)
	if ok {
		if a.GetName() == "index" {
			r.Primary = a
		}
		for i := 0; i < len(a.Metadata.Categories); i++ {
			category := a.Metadata.Categories[i]
			exists, ok := r.CategoriesMap[category]
			newCategory := !(ok && exists)
			if newCategory {
				r.CategoriesMap[category] = true
				r.Categories = append(r.Categories, category)
			}
		}
	}
	return nil
}

func (r *Repository) Children() *[]IFile {
	return r.GetRelatives([]RelationStatus{CHILD})
}

func (r *Repository) Content() string {
	var content string
	if r.Primary != nil {
		content = r.Primary.Content
	}
	return content
}

func (r *Repository) Metadata() *Metadata {
	var metadata *Metadata
	if r.Primary != nil {
		metadata = r.Primary.Metadata
	}
	return metadata
}
