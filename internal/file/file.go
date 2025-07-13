package file

import (
	"errors"
)

type IFile interface {
	GetParent() *Repository
	GetName() string
	GetLocation() string
	GetPedigree() string
	GetIndex() uint
	GetRelatives(relations []RelationStatus) *[]IFile
	GetRelations() *map[IFile]Relation
	Acknowledge(p *Repository) error
	Add(f2 IFile) error
}

type File struct {
	Parent    *Repository
	Name      string
	Location  string
	Pedigree  string
	Index     uint
	Relatives *[]IFile
	Relations *map[IFile]Relation
}

func MakeFile(repo *Repository, index uint, name string, location string) *File {
	var err error
	if name == "" {
		name = "unnamed file"
	}
	defaultRelatives := make([]IFile, 0)
	defaultRelations := make(map[IFile]Relation)
	file := File{
		repo,
		name,
		location,
		"",
		index,
		&defaultRelatives,
		&defaultRelations,
	}
	err = file.Acknowledge(repo)
	if err != nil {
		return nil
	} else {
		return &file
	}
}

func (f *File) GetParent() *Repository {
	return f.Parent
}
func (f *File) GetName() string {
	return f.Name
}
func (f *File) GetLocation() string {
	return f.Location
}
func (f *File) GetPedigree() string {
	return f.Pedigree
}
func (f *File) GetIndex() uint {
	return f.Index
}
func (f *File) GetRelations() *map[IFile]Relation {
	return f.Relations
}

func (f1 *File) Equals(f2 IFile) bool {
	return (f1.Pedigree == f2.GetPedigree()) && ((f1.Index == f2.GetIndex()) || (f1.Name == f2.GetName()))
}

func (f1 *File) Relation(f2 IFile) *Relation {
	var relation Relation
	parentsExistAndEqual := f1.Parent != nil && f2.GetParent() != nil && f1.Parent.Equals(f2.GetParent())
	if parentsExistAndEqual {
		if f1.Equals(f2) {
			relation = MakeRelation(INDETERMINATE, INDETERMINATE)
		} else if f1.Index < f2.GetIndex() {
			relation = MakeRelation(ELDER_SIBLING, YOUNGER_SIBLING)
		} else if f1.Index > f2.GetIndex() {
			relation = MakeRelation(YOUNGER_SIBLING, ELDER_SIBLING)
		} else {
			relation = MakeRelation(INDETERMINATE, INDETERMINATE)
		}
	} else {
		f1ParentExists := f1.Parent != nil
		f2ParentExists := f2.GetParent() != nil
		if f1ParentExists && f1.Parent.Equals(f2) {
			relation = MakeRelation(CHILD, PARENT)
		} else if f2ParentExists && f2.GetParent().Equals(f1) {
			relation = MakeRelation(PARENT, CHILD)
		} else if f1.Equals(f2) {
			relation = MakeRelation(INDETERMINATE, INDETERMINATE)
		} else {
			relation = MakeRelation(DISTANT, DISTANT)
		}
	}
	return &relation
}

func (f1 *File) Add(f2 IFile) error {
	var err error
	relation := f1.Relation(f2)
	_, exists := f1.Related(f2)
	if relation.FirstParty == INDETERMINATE {
		err = errors.New("2 identical files cannot have a relation")
		return err
	}
	(*f1.Relations)[f2] = *relation
	if !exists {
		*f1.Relatives = append(*f1.Relatives, f2)
	}
	return err
}

func (f1 *File) GetRelatives(relations []RelationStatus) *[]IFile {
	statusMatchExists := func(rel RelationStatus) bool {
		matchExists := false
		for i := 0; i < len(relations); i++ {
			if relations[i] == rel {
				matchExists = true
				break
			}
		}
		return matchExists
	}
	relatives := *f1.Relatives
	relativesOfStatus := make([]IFile, 0)
	for i := 0; i < len(relatives); i++ {
		f := relatives[i]
		rel, exists := (*f1.Relations)[f]
		if exists && statusMatchExists(rel.SecondParty) {
			relativesOfStatus = append(relativesOfStatus, f)
		}
	}
	return &relativesOfStatus
}

func (f1 *File) Siblings() *[]IFile {
	return f1.GetRelatives([]RelationStatus{ELDER_SIBLING, YOUNGER_SIBLING})
}

func (f1 *File) Related(f2 IFile) (*Relation, bool) {
	relation, exists := (*f1.Relations)[f2]
	if exists {
		return &relation, true
	}
	return nil, false
}

func (f *File) Acknowledge(p *Repository) error {
	f.Parent = p
	var pedigree, repoPedigree, repoName string
	var err error
	if p != nil {
		repoPedigree = p.Pedigree
		repoName = "/" + p.Name
		err = f.Add(p)
	}
	if err != nil {
		return err
	}
	pedigree = repoPedigree + repoName
	f.Pedigree = pedigree
	return err
}

func (f *File) Path() string {
	path := f.Pedigree + "/" + f.Name
	return path
}
