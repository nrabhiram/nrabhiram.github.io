package tests

import (
	"fmt"
	"sapphire/internal/file"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMakeFileChildren(t *testing.T) {
	t.Run("2 identical files, that have same indexes, but different names, can't both be children of a repository", func(t *testing.T) {
		r := file.MakeRepository(nil, 0, "repo", "")
		f1 := file.MakeFile(r, 0, "first file", "")
		f2 := file.MakeFile(r, 0, "second file", "")
		err1 := r.Make(f1)
		err2 := r.Make(f2)
		expectedErrorMessage := fmt.Sprintf("\"%s\" and \"%s\" can't be contained in the repository named \"%s\" because they have the same index", f2.GetName(), f1.GetName(), r.Name)
		assert.True(t, err1 == nil)
		assert.True(t, err2 != nil)
		assert.True(t, err2.Error() == expectedErrorMessage)
	})

	t.Run("2 identical files, that have same names, but different indexes, can't both be children of a repository", func(t *testing.T) {
		r := file.MakeRepository(nil, 0, "repo", "")
		f1 := file.MakeFile(r, 0, "file", "")
		f2 := file.MakeFile(r, 1, "file", "")
		err1 := r.Make(f1)
		err2 := r.Make(f2)
		expectedErrorMessage := fmt.Sprintf("files in the repository named \"%s\" have the same name as \"%s\"", r.Name, f1.GetName())
		assert.True(t, err1 == nil)
		assert.True(t, err2 != nil)
		assert.True(t, err2.Error() == expectedErrorMessage)
	})

	t.Run("2 completely identical files, can't both be children of a repository", func(t *testing.T) {
		r := file.MakeRepository(nil, 0, "repo", "")
		f1 := file.MakeFile(r, 0, "file", "")
		f2 := file.MakeFile(r, 0, "file", "")
		err1 := r.Make(f1)
		err2 := r.Make(f2)
		expectedErrorMessage := fmt.Sprintf("files in the repository named \"%s\" have the same name as \"%s\"", r.Name, f1.GetName())
		assert.True(t, err1 == nil)
		assert.True(t, err2 != nil)
		assert.True(t, err2.Error() == expectedErrorMessage)
	})

	t.Run("the children of a repository that contains 2 files, adds the other file as its sibling", func(t *testing.T) {
		r := file.MakeRepository(nil, 0, "repo", "")
		f1 := file.MakeFile(r, 0, "first file", "")
		f2 := file.MakeFile(r, 1, "second file", "")
		err1 := r.Make(f1)
		err2 := r.Make(f2)
		assert.True(t, err1 == nil)
		assert.True(t, err2 == nil)
		assert.True(t, len(*f1.Siblings()) == 1)
		assert.True(t, (*f1.Siblings())[0] == f2)
		assert.True(t, len(*f2.Siblings()) == 1)
		assert.True(t, (*f2.Siblings())[0] == f1)
	})

	t.Run("the third child of a repository that contains 3 files, adds the other 2 files as its siblings", func(t *testing.T) {
		r := file.MakeRepository(nil, 0, "repo", "")
		f1 := file.MakeFile(r, 0, "first file", "")
		f2 := file.MakeFile(r, 1, "second file", "")
		f3 := file.MakeFile(r, 2, "third file", "")
		err1 := r.Make(f1)
		err2 := r.Make(f2)
		err3 := r.Make(f3)
		assert.True(t, err1 == nil)
		assert.True(t, err2 == nil)
		assert.True(t, err3 == nil)
		assert.True(t, len(*f3.Siblings()) == 2)
		assert.True(t, (*f3.Siblings())[0] == f1)
		assert.True(t, (*f3.Siblings())[1] == f2)
	})
}

func TestRepoCategories(t *testing.T) {
	t.Run("a repository that makes 2 artifacts, the first one with the categories \"notes\" and \"essays\", and the second one with a single category \"projects\", has 3 categories - \"notes\", \"essays\", and \"projects\"", func(t *testing.T) {
		r := file.MakeRepository(nil, 0, "repo", "")
		m1 := file.MakeMetadata(
			time.Now(),
			"",
			[]string{"notes", "essays"},
			"",
			"",
			"",
			0,
			map[string]string{},
		)
		a1 := file.MakeArtifact(r, 0, "first file", "", "", m1)
		m2 := file.MakeMetadata(
			time.Now(),
			"",
			[]string{"projects"},
			"",
			"",
			"",
			0,
			map[string]string{},
		)
		a2 := file.MakeArtifact(r, 1, "second file", "", "", m2)
		err1 := r.Make(a1)
		err2 := r.Make(a2)
		assert.True(t, err1 == nil)
		assert.True(t, err2 == nil)
		assert.True(t, len(r.Categories) == 3)
		assert.ElementsMatch(t, r.Categories, []string{"notes", "essays", "projects"})
	})

	t.Run("a repository that makes 2 artifacts, the first one with the categories \"notes\" and \"essays\", and the second one with the categories \"essays\" and \"projects\", has 3 categories - \"notes\", \"essays\", and \"projects\"", func(t *testing.T) {
		r := file.MakeRepository(nil, 0, "repo", "")
		m1 := file.MakeMetadata(
			time.Now(),
			"",
			[]string{"notes", "essays"},
			"",
			"",
			"",
			0,
			map[string]string{},
		)
		a1 := file.MakeArtifact(r, 0, "first file", "", "", m1)
		m2 := file.MakeMetadata(
			time.Now(),
			"",
			[]string{"essays", "projects"},
			"",
			"",
			"",
			0,
			map[string]string{},
		)
		a2 := file.MakeArtifact(r, 1, "second file", "", "", m2)
		err1 := r.Make(a1)
		err2 := r.Make(a2)
		assert.True(t, err1 == nil)
		assert.True(t, err2 == nil)
		assert.True(t, len(r.Categories) == 3)
		assert.ElementsMatch(t, r.Categories, []string{"notes", "essays", "projects"})
	})
}

func TestPrimaryArtifact(t *testing.T) {
	m := file.MakeMetadata(
		time.Now(),
		"",
		[]string{"essays", "projects"},
		"",
		"",
		"",
		0,
		map[string]string{},
	)

	t.Run("when an artifact with a name of \"index\" is made a child of a repository with no children, the repository's content is equal to the artifact's content", func(t *testing.T) {
		r := file.MakeRepository(nil, 0, "repo", "")
		a := file.MakeArtifact(r, 0, "index", "", "hello world", m)
		err := r.Make(a)
		assert.True(t, err == nil)
		assert.True(t, r.Content() == a.Content)
	})

	t.Run("when an artifact with a name of \"index\" is made a child of a repository with no children, the repository's metadata is equal to the artifact's metadata", func(t *testing.T) {
		r := file.MakeRepository(nil, 0, "repo", "")
		a := file.MakeArtifact(r, 0, "index", "", "hello world", m)
		err := r.Make(a)
		assert.True(t, err == nil)
		assert.True(t, r.Metadata() == a.Metadata)
	})

	t.Run("when an artifact with a name that isn't \"index\" is made a child of a repository with no children, it doesn't have any content", func(t *testing.T) {
		r := file.MakeRepository(nil, 0, "repo", "")
		a := file.MakeArtifact(r, 0, "artifact", "", "hello world", m)
		err := r.Make(a)
		assert.True(t, err == nil)
		assert.True(t, a.Content == "hello world")
		assert.True(t, r.Content() == "")
	})

	t.Run("when an artifact with a name that isn't \"index\" is made a child of a repository with no children, it doesn't have any metadata", func(t *testing.T) {
		r := file.MakeRepository(nil, 0, "repo", "")
		a := file.MakeArtifact(r, 0, "artifact", "", "hello world", m)
		err := r.Make(a)
		assert.True(t, err == nil)
		assert.True(t, a.Metadata == m)
		assert.True(t, r.Metadata() == nil)
	})
}
