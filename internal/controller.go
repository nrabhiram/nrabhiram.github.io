package internal

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sapphire/internal/file"
	"sapphire/internal/parser"
	"sapphire/internal/renderer"
	"strings"
)

type Controller struct {
	schema         *renderer.Schema
	flattenedLinks *map[*renderer.Repository]*[]*renderer.Link
}

type artifactResult struct {
	route    string
	artifact renderer.IArtifact
	err      error
}

type schemaRepoResult struct {
	repoPath string
	repo     *file.Repository
	err      error
}

func MakeController() *Controller {
	defaultFlattenedLinks := make(map[*renderer.Repository]*[]*renderer.Link)
	defaultSchema := make(renderer.Schema)
	return &Controller{
		schema:         &defaultSchema,
		flattenedLinks: &defaultFlattenedLinks,
	}
}

func (c *Controller) findArtifactInRepo(repo renderer.IArtifact, path string) renderer.IArtifact {
	if repo.GetPath() == path {
		return repo
	} else {
		for _, artifact := range *repo.GetArtifacts() {
			if artifact.GetPath() == path {
				return artifact
			} else {
				artifacts := artifact.GetArtifacts()
				if artifacts != nil && len(*artifacts) != 0 {
					foundArtifact := c.findArtifactInRepo(artifact, path)
					if foundArtifact != nil {
						return foundArtifact
					}
				}
			}
		}
	}
	return nil
}

func (c *Controller) flattenLinks(links *[](*renderer.Link)) *[](*renderer.Link) {
	var result [](*renderer.Link)

	for _, link := range *links {
		result = append(result, link)
		if link.Links != nil {
			result = append(result, (*c.flattenLinks(link.Links))...)
		}
	}

	return &result
}

func (c *Controller) getFlattenedLinks(repo *renderer.Repository) *[](*renderer.Link) {
	flattenedLinks := make([]*renderer.Link, 0)
	if repo != nil {
		if links, exists := (*c.flattenedLinks)[repo]; exists && links != nil {
			flattenedLinks = *links
		}
	}
	return &flattenedLinks
}

func (c *Controller) getAdjacentLiveLink(alias string, currentPath string, next bool) *renderer.Link {
	repo, ok := (*c.schema)[alias]
	if !ok {
		return nil
	}

	links := c.getFlattenedLinks(repo)
	var currentIndex int = -1

	for i := 0; i < len(*links); i++ {
		link := (*links)[i]
		if link.Path == currentPath {
			currentIndex = i
			break
		}
	}

	if currentIndex == -1 {
		return nil // currentPath not found
	}

	var startPoint, endPoint int

	if next {
		startPoint = currentIndex + 1
		endPoint = len(*links) - 1
	} else {
		startPoint = currentIndex - 1
		endPoint = 0
	}

	if startPoint < 0 || startPoint >= len(*links) {
		return nil // startPoint is out of bounds
	}

	for i := startPoint; (next && i <= endPoint) || (!next && i >= endPoint); {
		link := (*links)[i]
		if link.Live {
			return link
		}
		if next {
			i++
		} else {
			i--
		}
	}

	return nil
}

func (c *Controller) GetArtifact(alias string, path string) (renderer.IArtifact, error) {
	var err error
	repo, ok := (*c.schema)[alias]

	if !ok {
		err = fmt.Errorf("repository with an alias of \"%s\" doesn't exist", alias)
		return nil, err
	}

	pathExists := path != ""

	if pathExists {
		artifact := c.findArtifactInRepo(repo, path)
		if artifact == nil {
			err = fmt.Errorf("artifact with a path of \"%s\" in the repository named \"%s\" doesn't exist", path, alias)
			return nil, err
		}
		nextLink := c.getAdjacentLiveLink(alias, artifact.GetPath(), true)
		prevLink := c.getAdjacentLiveLink(alias, artifact.GetPath(), false)
		artifact.SetNextLink(nextLink)
		artifact.SetPrevLink(prevLink)
		return artifact, err
	} else {
		nextLink := c.getAdjacentLiveLink(alias, repo.GetPath(), true)
		prevLink := c.getAdjacentLiveLink(alias, repo.GetPath(), false)
		repo.SetNextLink(nextLink)
		repo.SetPrevLink(prevLink)
		return repo, err
	}
}

func (c *Controller) GetRoutes() *[]string {
	var allRoutes []string

	for _, repo := range *c.schema {
		paths := c.getFlattenedLinks(repo)
		for _, link := range *paths {
			allRoutes = append(allRoutes, link.Path)
		}
	}

	return &allRoutes
}

func (c *Controller) generateSchema(repoPaths []string, clientProjectRoot string, outputPath string) error {
	workers := runtime.GOMAXPROCS(0)

	limitChan := make(chan bool, workers)
	resultChan := make(chan schemaRepoResult)
	done := make(chan bool)

	for i, repoPath := range repoPaths {
		go func(i int, repoPath string) {
			select {
			case <-done:
				return
			default:
				limitChan <- true
				defer func() {
					<-limitChan
				}()
				absPath, err := filepath.Abs(filepath.Join(clientProjectRoot, repoPath))
				if err != nil {
					resultChan <- schemaRepoResult{
						err: fmt.Errorf("failed to resolve path: %w", err),
					}
					return
				}

				repository, err := parser.BuildRepository(absPath, clientProjectRoot, uint(i))
				if err != nil {
					resultChan <- schemaRepoResult{
						err:      fmt.Errorf("failed to add repository: %w", err),
						repoPath: repoPath,
					}
				} else {
					resultChan <- schemaRepoResult{
						repo:     repository,
						repoPath: repoPath,
						err:      nil,
					}
				}
				return
			}
		}(i, repoPath)
	}

	for i := 0; i < len(repoPaths); i++ {
		result := <-resultChan
		if result.err != nil {
			// handle error, maybe write it to an extra channel
			close(done)
		} else {
			repository := result.repo
			repoPath := result.repoPath
			renderedRepo := renderer.RenderRepository(repository)
			key := parser.CreateAlias(repoPath)
			(*c.schema)[key] = renderedRepo
			(*c.flattenedLinks)[renderedRepo] = c.flattenLinks(renderedRepo.GetLinks())
		}
	}

	file, err := os.Create(filepath.Join(outputPath, "schema.json"))
	if err != nil {
		return fmt.Errorf("the following error occured while creating the schema file:\n%w", err)
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(*c.schema); err != nil {
		return fmt.Errorf("the following error occured while generating the schema:\n%w", err)
	}

	return nil
}

func (c *Controller) ProcessFlags() (*string, *string, *string) {
	return parser.ProcessFlags()
}

func (c *Controller) GenerateSchema(repoPaths *string, outputPath *string) {
	if *repoPaths == "" {
		renderer.LogError("comma-separated list of repoPaths is required for generate-schema")
	}

	paths := parser.ProcessRepoPaths(repoPaths)

	cwd, err := os.Getwd()
	if err != nil {
		renderer.LogError(err.Error())
	}

	err = c.generateSchema(paths, cwd, *outputPath)

	if err != nil {
		renderer.LogError("error generating schema:\n" + err.Error())
	}

	renderer.LogMessage("schema generated successfully!")
}

func (c *Controller) Build(repoPaths *string, outputPath *string) {
	workers := runtime.GOMAXPROCS(0)

	limitChan := make(chan bool, workers)
	resultChan := make(chan artifactResult)
	done := make(chan bool)

	defer close(done)

	c.GenerateSchema(repoPaths, outputPath)

	routes := c.GetRoutes()
	generatedArtifacts := make(map[string]renderer.IArtifact)

	for _, route := range *routes {
		go func(route string) {
			limitChan <- true
			defer func() {
				<-limitChan
			}()

			select {
			case _, ok := <-done:
				if !ok {
					return
				}
			default:
				parts := strings.Split(route, "/")
				if len(parts) < 2 {
					resultChan <- artifactResult{
						route:    route,
						err:      fmt.Errorf("could not generate artifact data due to an invalid route format: %s", route),
						artifact: nil,
					}
					return
				}

				alias := strings.Split(route, "/")[1]

				artifact, err := c.GetArtifact(alias, route)
				if err != nil {
					err = fmt.Errorf("unable to get the artifact for the route %s due to the following error\n: %s", route, err.Error())
				}
				resultChan <- artifactResult{
					route:    route,
					err:      err,
					artifact: artifact,
				}
				return
			}
		}(route)
	}

	for i := 0; i < len(*routes); i++ {
		result := <-resultChan
		if result.err != nil {
			renderer.LogMessage(fmt.Sprintf("Error processing route %s: %v", result.route, result.err))
			return
		} else {
			generatedArtifacts[result.route] = result.artifact
		}
	}

	outputFilePath := filepath.Join(*outputPath, "artifacts.json")
	file, err := os.Create(outputFilePath)
	if err != nil {
		renderer.LogError("error creating artifacts file: \n" + err.Error())
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(generatedArtifacts); err != nil {
		renderer.LogError("error writing artifacts to file: \n" + err.Error())
	}

	renderer.LogMessage("artifacts successfully written to: " + outputFilePath)
}
