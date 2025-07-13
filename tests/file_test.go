package tests

import (
	"sapphire/internal/file"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileNaming(t *testing.T) {
	t.Run("the name of a file that isn't named is 'unnamed file'", func(t *testing.T) {
		var f = file.MakeFile(nil, 0, "", "")
		assert.True(t, f.Name == "unnamed file")
	})
}

func TestFilePedigree(t *testing.T) {
	t.Run("a file without a parent doesn't have a pedigree", func(t *testing.T) {
		var f = file.MakeFile(nil, 0, "", "")
		assert.True(t, f.Pedigree == "")
	})

	t.Run("the pedigree of a file with a parent is equal to the parent's pedigree and name, separated by a forward slash", func(t *testing.T) {
		var parent = file.MakeRepository(nil, 0, "parent", "")
		var f = file.MakeFile(parent, 0, "file", "")
		assert.True(t, f.Pedigree == "/parent")
	})

	t.Run("the pedigrees of 2 files that have different parents of the same name, but different grandparents with different names, aren't equal", func(t *testing.T) {
		var firstGrandparent = file.MakeRepository(nil, 1, "firstGrandparent", "")
		var secondGrandparent = file.MakeRepository(nil, 2, "secondGrandparent", "")
		var firstParent = file.MakeRepository(firstGrandparent, 1, "parent", "")
		var secondParent = file.MakeRepository(secondGrandparent, 2, "parent", "")
		var firstFile = file.MakeFile(firstParent, 1, "firstFile", "")
		var secondFile = file.MakeFile(secondParent, 2, "secondFile", "")
		assert.True(t, firstFile.Pedigree == "/firstGrandparent/parent")
		assert.True(t, secondFile.Pedigree == "/secondGrandparent/parent")
		assert.True(t, firstFile.Pedigree != secondFile.Pedigree)
	})

	t.Run("the pedigrees of 2 files that have different parents of the same name, and different grandparents of the same name, are equal", func(t *testing.T) {
		var firstGrandparent = file.MakeRepository(nil, 1, "grandparent", "")
		var secondGrandparent = file.MakeRepository(nil, 2, "grandparent", "")
		var firstParent = file.MakeRepository(firstGrandparent, 1, "parent", "")
		var secondParent = file.MakeRepository(secondGrandparent, 2, "parent", "")
		var firstFile = file.MakeFile(firstParent, 1, "firstFile", "")
		var secondFile = file.MakeFile(secondParent, 2, "secondFile", "")
		assert.True(t, firstFile.Pedigree == "/grandparent/parent")
		assert.True(t, secondFile.Pedigree == "/grandparent/parent")
		assert.True(t, firstFile.Pedigree == secondFile.Pedigree)
	})
}

func TestFileIdentity(t *testing.T) {
	t.Run("2 files with the same pedigree, name, and index are identical", func(t *testing.T) {
		var parent = file.MakeRepository(nil, 1, "parent", "")
		var firstFile = file.MakeFile(parent, 1, "file", "")
		var secondFile = file.MakeFile(parent, 1, "file", "")
		assert.True(t, firstFile.Equals(secondFile))
	})

	t.Run("2 files with the same pedigree and name, but different indexes, are identical", func(t *testing.T) {
		var parent = file.MakeRepository(nil, 1, "parent", "")
		var firstFile = file.MakeFile(parent, 1, "file", "")
		var secondFile = file.MakeFile(parent, 2, "file", "")
		assert.True(t, firstFile.Equals(secondFile))
	})

	t.Run("2 files with the same pedigree and index, but different names, are identical", func(t *testing.T) {
		var parent = file.MakeRepository(nil, 1, "parent", "")
		var firstFile = file.MakeFile(parent, 1, "firstFile", "")
		var secondFile = file.MakeFile(parent, 1, "secondFile", "")
		assert.True(t, firstFile.Equals(secondFile))
	})

	t.Run("2 files with different pedigree aren't identical", func(t *testing.T) {
		var firstParent = file.MakeRepository(nil, 1, "firstParent", "")
		var secondParent = file.MakeRepository(nil, 1, "secondParent", "")
		var firstFile = file.MakeFile(firstParent, 1, "firstFile", "")
		var secondFile = file.MakeFile(secondParent, 1, "secondFile", "")
		assert.False(t, firstFile.Equals(secondFile))
	})

	t.Run("2 files with the same pedigree, but different indexes and names, aren't identical", func(t *testing.T) {
		var parent = file.MakeRepository(nil, 1, "parent", "")
		var firstFile = file.MakeFile(parent, 1, "firstFile", "")
		var secondFile = file.MakeFile(parent, 2, "secondFile", "")
		assert.False(t, firstFile.Equals(secondFile))
	})
}

func TestFileRelations(t *testing.T) {
	t.Run("for 2 unique files with the same parent, the file with the lower index is related to the other as the elder sibling, and the file with the greater index is related to the other as the younger sibling", func(t *testing.T) {
		var parent = file.MakeRepository(nil, 1, "parent", "")
		var firstFile = file.MakeFile(parent, 1, "firstFile", "")
		var secondFile = file.MakeFile(parent, 2, "secondFile", "")
		var firstFileRel = firstFile.Relation(secondFile)
		var secondFileRel = secondFile.Relation(firstFile)
		assert.True(t, firstFileRel.FirstParty == file.ELDER_SIBLING)
		assert.True(t, secondFileRel.FirstParty == file.YOUNGER_SIBLING)
	})

	t.Run("for 2 identical orphan files, the relationship between them is indeterminate", func(t *testing.T) {
		var firstFile = file.MakeFile(nil, 1, "file", "")
		var secondFile = file.MakeFile(nil, 1, "file", "")
		var firstFileRel = firstFile.Relation(secondFile)
		var secondFileRel = secondFile.Relation(firstFile)
		assert.True(t, firstFileRel.FirstParty == file.INDETERMINATE)
		assert.True(t, secondFileRel.FirstParty == file.INDETERMINATE)
	})

	t.Run("for 2 unique orphan files, the relationship between them is distant", func(t *testing.T) {
		var firstFile = file.MakeFile(nil, 1, "firstFile", "")
		var secondFile = file.MakeFile(nil, 2, "secondFile", "")
		var firstFileRel = firstFile.Relation(secondFile)
		var secondFileRel = secondFile.Relation(firstFile)
		assert.True(t, firstFileRel.FirstParty == file.DISTANT)
		assert.True(t, secondFileRel.FirstParty == file.DISTANT)
	})

	t.Run("for 2 identical files with the same parent, the relationship between them is indeterminate", func(t *testing.T) {
		var parent = file.MakeRepository(nil, 1, "parent", "")
		var firstFile = file.MakeFile(parent, 1, "firstFile", "")
		var secondFile = file.MakeFile(parent, 1, "secondFile", "")
		var firstFileRel = firstFile.Relation(secondFile)
		var secondFileRel = secondFile.Relation(firstFile)
		assert.True(t, firstFileRel.FirstParty == file.INDETERMINATE)
		assert.True(t, secondFileRel.FirstParty == file.INDETERMINATE)
	})

	t.Run("the relationship between a parent and child", func(t *testing.T) {
		var parent = file.MakeRepository(nil, 1, "parent", "")
		var child = file.MakeFile(parent, 1, "child", "")
		var parentRel = parent.Relation(child)
		var childRel = child.Relation(parent)
		assert.True(t, parentRel.FirstParty == file.PARENT)
		assert.True(t, childRel.FirstParty == file.CHILD)
	})
}

func TestAddFileRelations(t *testing.T) {
	t.Run("when an orphan file with no relatives, adds a relationship with another unique orphan file, the first file has a single relative, and they're related as distant relatives", func(t *testing.T) {
		firstFile := file.MakeFile(nil, 1, "firstFile", "")
		secondFile := file.MakeFile(nil, 2, "secondFile", "")
		err := firstFile.Add(secondFile)
		rel, exists := firstFile.Related(secondFile)
		assert.True(t, err == nil)
		assert.True(t, len(*firstFile.Relatives) == 1)
		assert.True(t, exists)
		assert.True(t, rel.FirstParty == file.DISTANT)
	})

	t.Run("a file with a parent, has at least its parent as a relative", func(t *testing.T) {
		parent := file.MakeRepository(nil, 0, "parent", "")
		f := file.MakeFile(parent, 0, "file", "")
		assert.True(t, len(*f.Relatives) == 1)
		assert.True(t, (*f.Relatives)[0] == parent)
	})

	t.Run("a file that only has a relation with its parent, and adds relations with 2 files, an elder and younger sibling, has a total of 3 relatives", func(t *testing.T) {
		parent := file.MakeRepository(nil, 1, "parent", "")
		f2 := file.MakeFile(parent, 2, "file", "")
		f1 := file.MakeFile(parent, 1, "elder file", "")
		f3 := file.MakeFile(parent, 3, "younger file", "")
		err1 := f2.Add(f1)
		err2 := f2.Add(f3)
		assert.True(t, err1 == nil)
		assert.True(t, err2 == nil)
		assert.True(t, len(*f2.Siblings()) == 2)
		assert.True(t, len(*f2.Relatives) == 3)
		assert.True(t, (*f2.Siblings())[0] == f1)
		assert.True(t, (*f2.Siblings())[1] == f3)
	})

	t.Run("2 identical orphan files, with no parents, can't add relations with each other", func(t *testing.T) {
		f1 := file.MakeFile(nil, 1, "file", "")
		f2 := file.MakeFile(nil, 1, "file", "")
		err := f1.Add(f2)
		assert.True(t, err.Error() == "2 identical files cannot have a relation")
	})

	t.Run("2 identical files, with the same parent, can't add relations with each other", func(t *testing.T) {
		parent := file.MakeRepository(nil, 1, "parent", "")
		f1 := file.MakeFile(parent, 1, "file", "")
		f2 := file.MakeFile(parent, 1, "file", "")
		err := f1.Add(f2)
		assert.True(t, err.Error() == "2 identical files cannot have a relation")
	})

	t.Run("a file with no relatives, upon adding a relation with the same unique file twice, ends up with only 1 relative", func(t *testing.T) {
		f1 := file.MakeFile(nil, 1, "first file", "")
		f2 := file.MakeFile(nil, 2, "second file", "")
		err1 := f1.Add(f2)
		err2 := f1.Add(f2)
		assert.True(t, err1 == nil)
		assert.True(t, err2 == nil)
		assert.True(t, len(*f1.Relatives) == 1)
		assert.True(t, (*f1.Relatives)[0] == f2)
	})
}
