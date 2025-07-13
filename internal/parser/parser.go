package parser

import (
	"flag"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"sapphire/internal/file"
	"strconv"
	"strings"
	"time"
)

func CreateAlias(repoPath string) string {
	absolutePath, err := filepath.Abs(repoPath)
	if err != nil {
		return ""
	}
	return path.Base(absolutePath)
}

func extractMetadata(markdown string) file.Metadata {
	parts := strings.Split(markdown, "---")
	metadata := file.Metadata{
		Date:       time.Time{},
		Categories: []string{},
		Others:     make(map[string]string),
	}

	if len(parts) > 1 {
		frontMatter := parts[1]
		lines := strings.Split(frontMatter, "\n")

		// Regex for full date-time with timezone
		dateTimeRegex := regexp.MustCompile(`^(\d{4})-(\d{2})-(\d{2})[T\s](\d{2}):(\d{2}):(\d{2})(Z|([+-]\d{2}:\d{2}))$`)
		// Regex for date-only
		dateOnlyRegex := regexp.MustCompile(`^(\d{4})-(\d{2})-(\d{2})$`)
		numRegex := regexp.MustCompile(`^\d+$`)

		for _, line := range lines {
			line = strings.TrimSpace(line)
			if line == "" {
				continue
			}

			parts := strings.SplitN(line, ":", 2)
			if len(parts) != 2 {
				continue
			}

			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])

			switch key {
			case "date":
				if dateTimeRegex.MatchString(value) {
					parsedTime, err := time.Parse(time.RFC3339, value)
					if err == nil {
						metadata.Date = parsedTime
					}
				} else if dateOnlyRegex.MatchString(value) {
					parsedTime, err := time.Parse("2006-01-02", value)
					if err == nil {
						metadata.Date = parsedTime
					}
				}
			case "categories":
				categories := strings.Split(value, ",")
				for i := range categories {
					categories[i] = strings.TrimSpace(categories[i])
				}
				metadata.Categories = categories
			case "slug":
				metadata.Slug = value
			case "thumbnail":
				metadata.Thumbnail = value
			case "title":
				metadata.Title = value
			case "summary":
				metadata.Summary = value
			case "index":
				if numRegex.MatchString(value) {
					index, _ := strconv.Atoi(value)
					metadata.Index = uint(index)
				}
			default:
				metadata.Others[key] = value
			}
		}
	}

	return metadata
}

func extractContent(markdown string) string {
	parts := strings.Split(markdown, "---")
	if len(parts) > 2 {
		return strings.TrimSpace(strings.Join(parts[2:], "---"))
	}
	return strings.TrimSpace(markdown)
}

func readDirectory(directoryPath string, clientProjectRoot string, repo *file.Repository) error {
	entries, err := os.ReadDir(directoryPath)
	if err != nil {
		return err
	}

	for index, entry := range entries {
		itemPath := filepath.Join(directoryPath, entry.Name())
		pathFromProjectRoot, err := filepath.Rel(clientProjectRoot, itemPath)
		if err != nil {
			return err
		}

		if entry.IsDir() {
			subRepository := file.MakeRepository(nil, uint(index), entry.Name(), pathFromProjectRoot)
			err := repo.Make(subRepository)
			if err != nil {
				return err
			}
			if err = readDirectory(itemPath, clientProjectRoot, subRepository); err != nil {
				return err
			}
		} else {
			info, err := entry.Info()
			if err != nil {
				return err
			}

			if info.Mode().IsRegular() {
				if strings.HasSuffix(entry.Name(), ".md") {
					content, err := os.ReadFile(itemPath)
					if err != nil {
						return err
					}

					metadata := extractMetadata(string(content))
					markdownContent := extractContent(string(content))
					artifactName := strings.TrimSuffix(entry.Name(), filepath.Ext(entry.Name()))
					artifact := file.MakeArtifact(nil, uint(index), artifactName, pathFromProjectRoot, markdownContent, &metadata)
					err = repo.Make(artifact)
					if err != nil {
						return err
					}
				} else {
					sapphireFile := file.MakeFile(nil, uint(index), entry.Name(), pathFromProjectRoot)
					err := repo.Make(sapphireFile)
					if err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}

func BuildRepository(directoryPath string, clientProjectRoot string, index uint) (*file.Repository, error) {
	repoName := CreateAlias(directoryPath)

	location, err := filepath.Rel(clientProjectRoot, directoryPath)
	if err != nil {
		return nil, fmt.Errorf("could not determine relative path of repository named %s due to the following error:\n%w", repoName, err)
	}

	repo := file.MakeRepository(nil, index, repoName, location)

	err = readDirectory(directoryPath, clientProjectRoot, repo)
	if err != nil {
		return nil, fmt.Errorf("could not process the repository named %s due to the following error:\n%w", repoName, err)
	}

	return repo, nil
}

// the following functions are used for processing inputs from the command-line application

func ProcessFlags() (command *string, repoPaths *string, outputPath *string) {
	command = flag.String("command", "", "Command to execute (GenerateSchema, Build)")
	repoPaths = flag.String("repoPaths", "", "Comma-separated list of repository paths")
	outputPath = flag.String("output", "utils", "Output path for the schema")
	flag.Parse()
	return command, repoPaths, outputPath
}

func ProcessRepoPaths(repoPaths *string) []string {
	paths := strings.Split(*repoPaths, ",")
	return paths
}

func ProcessSchemaPath(outputPath *string) string {
	return filepath.Join(*outputPath, "schema.json")
}
