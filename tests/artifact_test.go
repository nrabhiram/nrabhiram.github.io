package tests

import (
	"sapphire/internal/file"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestArtifactPath(t *testing.T) {
	metadata := file.MakeMetadata(
		time.Now(),
		"",
		[]string{"notes", "programming"},
		"",
		"",
		"",
		0,
		map[string]string{},
	)

	t.Run("an artifact with a name of 'index', that has no slug, and is stored in an orphan repository with a name of 'notes', has a path of '/notes'", func(t *testing.T) {
		r := file.MakeRepository(nil, 0, "notes", "")
		a := file.MakeArtifact(r, 0, "index", "", "", metadata)
		assert.True(t, a.Path() == "/notes")
	})

	t.Run("an artifact with a name of 'index', that has a slug, and is stored in an orphan repository with a name of 'notes', has a path of '/notes'", func(t *testing.T) {
		metadata := file.MakeMetadata(
			time.Now(),
			"home",
			[]string{"notes", "programming"},
			"",
			"",
			"",
			0,
			map[string]string{},
		)
		r := file.MakeRepository(nil, 0, "notes", "")
		a := file.MakeArtifact(r, 0, "index", "", "", metadata)
		assert.True(t, a.Path() == "/notes")
	})

	t.Run("an artifact with a name of 'note', that has no slug, and is stored in an orphan repository with a name of 'notes', has a path of '/notes/note'", func(t *testing.T) {
		r := file.MakeRepository(nil, 0, "notes", "")
		a := file.MakeArtifact(r, 0, "note", "", "", metadata)
		assert.True(t, a.Path() == "/notes/note")
	})

	t.Run("an artifact with a name of 'note', that has a slug of 'first-note', and is stored in an orphan repository with a name of 'notes', has a path of '/notes/first-note'", func(t *testing.T) {
		metadata := file.MakeMetadata(
			time.Now(),
			"first-note",
			[]string{"notes", "programming"},
			"",
			"",
			"",
			0,
			map[string]string{},
		)
		r := file.MakeRepository(nil, 0, "notes", "")
		a := file.MakeArtifact(r, 0, "note", "", "", metadata)
		assert.True(t, a.Path() == "/notes/first-note")
	})
}
