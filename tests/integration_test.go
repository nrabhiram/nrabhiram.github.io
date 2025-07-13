package tests

import (
	"os"
	"path/filepath"
	"sapphire/internal"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func formatSchemaOutput(jsonStr string, basePath string) string {
	jsonStr = strings.ReplaceAll(jsonStr, basePath, "/notes")
	var filteredLines []string
	lines := strings.Split(jsonStr, "\n")
	for i := range lines {
		lines[i] = strings.TrimLeft(lines[i], " \t")
		if !strings.Contains(lines[i], `"location":`) {
			filteredLines = append(filteredLines, lines[i])
		}
	}
	return strings.Join(filteredLines, "\n")
}

func TestGenerateSchemaIntegration(t *testing.T) {
	contentStructure := []struct {
		Path    string
		Content string
	}{
		{
			"recipes/index.md",
			`---
			title: My favourite recipes
			description: my notes on recipes for really cool dishes
			categories: cooking, recipes, food
			date: 2023-04-13T12:03:00+02:00
			index: 1
			---

			i've noted down some of my favourite recipes here. hope you like them`,
		},
		{
			"recipes/steak.md",
			`---
			title: Amazing Steak
			summary: this recipe is for making medium-rare steak
			categories: cooking, recipes, food
			date: 2023-04-13
			index: 1
			---

			this is some dummy content and instructions`,
		},
		{
			"recipes/pasta.md",
			`---
			title: Beautiful White Sauce Pasta
			summary: this recipe contains information regarding how one could make white-sauce pasta
			categories: cooking, recipes, food
			date: 2023-04-13
			index: 2
			---

			This is the recipe for white-sauce pasta.`,
		},
		{
			"recipes/ham-burgers.md",
			`---
			summary: this recipe teaches you how to make delicious burgers
			categories: cooking, recipes, food, junk
			date: 2023-04-13
			index: 2
			---

			This is my rendition of the Beef Rendang Burger from McDonalds.`,
		},
		{
			"recipes/idli.md",
			`---
			summary: idlis are a staple in south indian cuisines, and they're very easy to make
			categories: cooking, recipes, food
			date: 2023-04-10T12:09:00+02:00
			index: 1
			---`,
		},
		{
			"recipes/dosa.md",
			`---
			title: Crispy Dosa!
			summary: in this recipe, you'll learn how to make dosa and peanut chutney
			categories: cooking, recipes, food, diets
			date: 2023-04-13
			index: 2
			---`,
		},
		{
			"ts/index.md",
			`---
			title: notes on typescript
			description: these are some of the notes i had written on important ts topics like generics, interfaces, type guards, decorators, etc.
			categories: swe, ts
			date: 2023-04-13T13:03:00+02:00
			index: 1
			---

			ts is a superset`,
		},
		{
			"ts/generics/index.md",
			`---
			description: generics and how they can help you write better abstractions
			categories: swe, ts, generics
			date: 2023-04-13
			index: 1
			---

			hi`,
		},
		{
			"ts/interfaces.md",
			`---
			description: interfaces and how they are used
			categories: swe, ts, lol
			date: 2023-04-13
			index: 2
			---`,
		},
	}

	projectRoot, err := os.Getwd()
	assert.NoError(t, err, "Failed to get current working directory")

	tempDir, err := os.MkdirTemp("", "test-schema")
	assert.NoError(t, err, "Failed to create temporary directory")
	defer os.RemoveAll(tempDir)

	contentDir := filepath.Join(tempDir, "notes")
	err = os.Mkdir(contentDir, 0755)
	assert.NoError(t, err, "Failed to create content directory")

	recipesDir := filepath.Join(contentDir, "recipes")
	err = os.Mkdir(recipesDir, 0755)
	assert.NoError(t, err, "Failed to create recipes directory")

	tsDir := filepath.Join(contentDir, "ts")
	err = os.Mkdir(tsDir, 0755)
	assert.NoError(t, err, "Failed to create ts directory")

	genericsDir := filepath.Join(tsDir, "generics")
	err = os.Mkdir(genericsDir, 0755)
	assert.NoError(t, err, "Failed to create generics directory")

	for i := range contentStructure {
		file := contentStructure[i]
		path := filepath.Join(contentDir, file.Path)
		err = os.WriteFile(path, []byte(file.Content), 0644)
		assert.NoError(t, err, "Failed to write create markdown file")
	}

	controller := internal.MakeController()

	relativeRepoPath, err := filepath.Rel(projectRoot, contentDir)
	assert.NoError(t, err, "Failed to get relative repo path")

	relativeOutputPath, err := filepath.Rel(projectRoot, tempDir)
	assert.NoError(t, err, "Failed to get relative output path")

	controller.GenerateSchema(&relativeRepoPath, &relativeOutputPath)

	schemaFile := filepath.Join(tempDir, "schema.json")
	_, err = os.Stat(schemaFile)
	assert.NoError(t, err, "Schema file was not created")

	expectedSchema := `{
		"notes": {
			"name": "notes",
			"location": "notes",
			"path": "/notes",
			"categories": [],
			"artifacts": [
				{
					"name": "My favourite recipes",
					"location": "notes/recipes",
					"path": "/notes/recipes",
					"content": "i've noted down some of my favourite recipes here. hope you like them",
					"metadata": {
						"date": "2023-04-13T12:03:00+02:00",
						"slug": "",
						"categories": [
							"cooking",
							"recipes",
							"food"
						],
						"title": "My favourite recipes",
						"index": 1,
						"others": {
							"description": "my notes on recipes for really cool dishes"
						}
					},
					"categories": [
						"cooking",
						"recipes",
						"food",
						"diets",
						"junk"
					],
					"artifacts": [
						{
							"name": "idli",
							"location": "notes/recipes/idli.md",
							"path": "/notes/recipes/idli",
							"metadata": {
								"date": "2023-04-10T12:09:00+02:00",
								"slug": "",
								"categories": [
									"cooking",
									"recipes",
									"food"
								],
								"summary": "idlis are a staple in south indian cuisines, and they're very easy to make",
								"index": 1
							}
						},
						{
							"name": "Amazing Steak",
							"location": "notes/recipes/steak.md",
							"path": "/notes/recipes/steak",
							"content": "this is some dummy content and instructions",
							"metadata": {
								"date": "2023-04-13T00:00:00Z",
								"slug": "",
								"categories": [
									"cooking",
									"recipes",
									"food"
								],
								"title": "Amazing Steak",
								"summary": "this recipe is for making medium-rare steak",
								"index": 1
							}
						},
						{
							"name": "Beautiful White Sauce Pasta",
							"location": "notes/recipes/pasta.md",
							"path": "/notes/recipes/pasta",
							"content": "This is the recipe for white-sauce pasta.",
							"metadata": {
								"date": "2023-04-13T00:00:00Z",
								"slug": "",
								"categories": [
									"cooking",
									"recipes",
									"food"
								],
								"title": "Beautiful White Sauce Pasta",
								"summary": "this recipe contains information regarding how one could make white-sauce pasta",
								"index": 2
							}
						},
						{
							"name": "Crispy Dosa!",
							"location": "notes/recipes/dosa.md",
							"path": "/notes/recipes/dosa",
							"metadata": {
								"date": "2023-04-13T00:00:00Z",
								"slug": "",
								"categories": [
									"cooking",
									"recipes",
									"food",
									"diets"
								],
								"title": "Crispy Dosa!",
								"summary": "in this recipe, you'll learn how to make dosa and peanut chutney",
								"index": 2
							}
						},
						{
							"name": "ham burgers",
							"location": "notes/recipes/ham-burgers.md",
							"path": "/notes/recipes/ham-burgers",
							"content": "This is my rendition of the Beef Rendang Burger from McDonalds.",
							"metadata": {
								"date": "2023-04-13T00:00:00Z",
								"slug": "",
								"categories": [
									"cooking",
									"recipes",
									"food",
									"junk"
								],
								"summary": "this recipe teaches you how to make delicious burgers",
								"index": 2
							}
						}
					],
					"files": [],
					"links": [
						{
							"path": "/notes/recipes/idli",
							"label": "idli",
							"live": false,
							"metadata": {
								"date": "2023-04-10T12:09:00+02:00",
								"slug": "",
								"categories": [
									"cooking",
									"recipes",
									"food"
								],
								"summary": "idlis are a staple in south indian cuisines, and they're very easy to make",
								"index": 1
							},
							"links": []
						},
						{
							"path": "/notes/recipes/steak",
							"label": "Amazing Steak",
							"live": true,
							"metadata": {
								"date": "2023-04-13T00:00:00Z",
								"slug": "",
								"categories": [
									"cooking",
									"recipes",
									"food"
								],
								"title": "Amazing Steak",
								"summary": "this recipe is for making medium-rare steak",
								"index": 1
							},
							"links": []
						},
						{
							"path": "/notes/recipes/pasta",
							"label": "Beautiful White Sauce Pasta",
							"live": true,
							"metadata": {
								"date": "2023-04-13T00:00:00Z",
								"slug": "",
								"categories": [
									"cooking",
									"recipes",
									"food"
								],
								"title": "Beautiful White Sauce Pasta",
								"summary": "this recipe contains information regarding how one could make white-sauce pasta",
								"index": 2
							},
							"links": []
						},
						{
							"path": "/notes/recipes/dosa",
							"label": "Crispy Dosa!",
							"live": false,
							"metadata": {
								"date": "2023-04-13T00:00:00Z",
								"slug": "",
								"categories": [
									"cooking",
									"recipes",
									"food",
									"diets"
								],
								"title": "Crispy Dosa!",
								"summary": "in this recipe, you'll learn how to make dosa and peanut chutney",
								"index": 2
							},
							"links": []
						},
						{
							"path": "/notes/recipes/ham-burgers",
							"label": "ham burgers",
							"live": true,
							"metadata": {
								"date": "2023-04-13T00:00:00Z",
								"slug": "",
								"categories": [
									"cooking",
									"recipes",
									"food",
									"junk"
								],
								"summary": "this recipe teaches you how to make delicious burgers",
								"index": 2
							},
							"links": []
						}
					]
				},
				{
					"name": "notes on typescript",
					"location": "notes/ts",
					"path": "/notes/ts",
					"content": "ts is a superset",
					"metadata": {
						"date": "2023-04-13T13:03:00+02:00",
						"slug": "",
						"categories": [
							"swe",
							"ts"
						],
						"title": "notes on typescript",
						"index": 1,
						"others": {
							"description": "these are some of the notes i had written on important ts topics like generics, interfaces, type guards, decorators, etc."
						}
					},
					"categories": [
						"swe",
						"ts",
						"lol"
					],
					"artifacts": [
						{
							"name": "generics",
							"location": "notes/ts/generics",
							"path": "/notes/ts/generics",
							"content": "hi",
							"metadata": {
								"date": "2023-04-13T00:00:00Z",
								"slug": "",
								"categories": [
									"swe",
									"ts",
									"generics"
								],
								"index": 1,
								"others": {
									"description": "generics and how they can help you write better abstractions"
								}
							},
							"categories": [
								"swe",
								"ts",
								"generics"
							],
							"artifacts": [],
							"files": [],
							"links": []
						},
						{
							"name": "interfaces",
							"location": "notes/ts/interfaces.md",
							"path": "/notes/ts/interfaces",
							"metadata": {
								"date": "2023-04-13T00:00:00Z",
								"slug": "",
								"categories": [
									"swe",
									"ts",
									"lol"
								],
								"index": 2,
								"others": {
									"description": "interfaces and how they are used"
								}
							}
						}
					],
					"files": [],
					"links": [
						{
							"path": "/notes/ts/generics",
							"label": "generics",
							"live": true,
							"metadata": {
								"date": "2023-04-13T00:00:00Z",
								"slug": "",
								"categories": [
									"swe",
									"ts",
									"generics"
								],
								"index": 1,
								"others": {
									"description": "generics and how they can help you write better abstractions"
								}
							},
							"links": []
						},
						{
							"path": "/notes/ts/interfaces",
							"label": "interfaces",
							"live": false,
							"metadata": {
								"date": "2023-04-13T00:00:00Z",
								"slug": "",
								"categories": [
									"swe",
									"ts",
									"lol"
								],
								"index": 2,
								"others": {
									"description": "interfaces and how they are used"
								}
							},
							"links": []
						}
					]
				}
			],
			"files": [],
			"links": [
				{
					"path": "/notes/recipes",
					"label": "My favourite recipes",
					"live": true,
					"metadata": {
						"date": "2023-04-13T12:03:00+02:00",
						"slug": "",
						"categories": [
							"cooking",
							"recipes",
							"food"
						],
						"title": "My favourite recipes",
						"index": 1,
						"others": {
							"description": "my notes on recipes for really cool dishes"
						}
					},
					"links": [
						{
							"path": "/notes/recipes/idli",
							"label": "idli",
							"live": false,
							"metadata": {
								"date": "2023-04-10T12:09:00+02:00",
								"slug": "",
								"categories": [
									"cooking",
									"recipes",
									"food"
								],
								"summary": "idlis are a staple in south indian cuisines, and they're very easy to make",
								"index": 1
							},
							"links": []
						},
						{
							"path": "/notes/recipes/steak",
							"label": "Amazing Steak",
							"live": true,
							"metadata": {
								"date": "2023-04-13T00:00:00Z",
								"slug": "",
								"categories": [
									"cooking",
									"recipes",
									"food"
								],
								"title": "Amazing Steak",
								"summary": "this recipe is for making medium-rare steak",
								"index": 1
							},
							"links": []
						},
						{
							"path": "/notes/recipes/pasta",
							"label": "Beautiful White Sauce Pasta",
							"live": true,
							"metadata": {
								"date": "2023-04-13T00:00:00Z",
								"slug": "",
								"categories": [
									"cooking",
									"recipes",
									"food"
								],
								"title": "Beautiful White Sauce Pasta",
								"summary": "this recipe contains information regarding how one could make white-sauce pasta",
								"index": 2
							},
							"links": []
						},
						{
							"path": "/notes/recipes/dosa",
							"label": "Crispy Dosa!",
							"live": false,
							"metadata": {
								"date": "2023-04-13T00:00:00Z",
								"slug": "",
								"categories": [
									"cooking",
									"recipes",
									"food",
									"diets"
								],
								"title": "Crispy Dosa!",
								"summary": "in this recipe, you'll learn how to make dosa and peanut chutney",
								"index": 2
							},
							"links": []
						},
						{
							"path": "/notes/recipes/ham-burgers",
							"label": "ham burgers",
							"live": true,
							"metadata": {
								"date": "2023-04-13T00:00:00Z",
								"slug": "",
								"categories": [
									"cooking",
									"recipes",
									"food",
									"junk"
								],
								"summary": "this recipe teaches you how to make delicious burgers",
								"index": 2
							},
							"links": []
						}
					]
				},
				{
					"path": "/notes/ts",
					"label": "notes on typescript",
					"live": true,
					"metadata": {
						"date": "2023-04-13T13:03:00+02:00",
						"slug": "",
						"categories": [
							"swe",
							"ts"
						],
						"title": "notes on typescript",
						"index": 1,
						"others": {
							"description": "these are some of the notes i had written on important ts topics like generics, interfaces, type guards, decorators, etc."
						}
					},
					"links": [
						{
							"path": "/notes/ts/generics",
							"label": "generics",
							"live": true,
							"metadata": {
								"date": "2023-04-13T00:00:00Z",
								"slug": "",
								"categories": [
									"swe",
									"ts",
									"generics"
								],
								"index": 1,
								"others": {
									"description": "generics and how they can help you write better abstractions"
								}
							},
							"links": []
						},
						{
							"path": "/notes/ts/interfaces",
							"label": "interfaces",
							"live": false,
							"metadata": {
								"date": "2023-04-13T00:00:00Z",
								"slug": "",
								"categories": [
									"swe",
									"ts",
									"lol"
								],
								"index": 2,
								"others": {
									"description": "interfaces and how they are used"
								}
							},
							"links": []
						}
					]
				}
			]
		}
	}`

	basePath := filepath.Join(tempDir, "notes")

	fileContent, err := os.ReadFile(schemaFile)
	assert.NoError(t, err, "Failed to read schema file")

	expectedSchemaNormalized := formatSchemaOutput(expectedSchema, basePath)
	actualContentNormalized := formatSchemaOutput(string(fileContent), basePath)

	assert.JSONEq(t, expectedSchemaNormalized, actualContentNormalized, "Schema content does not match expected content")
}
