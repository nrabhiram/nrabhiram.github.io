.PHONY: clean # Remove the executable files that had been created 
clean:
	mkdir -p bin
	find bin -type f -not -name 'golangci-lint' -delete

.PHONY: build # Compile your application code into binary code that can be run as an executable
build: clean
	go build -o bin/main cmd/sapphire/main.go

.PHONY: dev # Run the application in development mode without having to build and run the executable file
dev:
	go run cmd/sapphire/main.go

.PHONY: server # Run the application in development mode without having to build and run the executable file
server:
	go run cmd/server/main.go

.PHONY: run # Build the application and run the compiled executable file 
run: build
	./bin/main --command=$(COMMAND) --repoPaths=$(REPO_PATHS) --output=$(OUTPUT)

.PHONY: deps # Install the project dependencies
deps:
	go mod tidy
	go mod download

.PHONY: install-lint # Manually download the linter's executable as recommended by gloangci-lint maintainers
install-lint:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ./bin v1.58.2

.PHONY: install # Set up the project by installing all of the necessary dependencies and the linter
install: deps install-lint

.PHONY: format # Format all of your Go code by adhering to conventions
format:
	gofmt -w $(shell find . -type f -name '*.go')

.PHONY: lint # Run the linter to check if code isn't complying to style guides 
lint:
	./bin/golangci-lint run

.PHONY: test # Run all of the tests written in the project
test:
	go test ./tests/... -v | sed 's/_/ /g'

.PHONY: help # Generate list of targets with descriptions
help:
	@max=$$(grep '^.PHONY: .* #' Makefile | sed 's/\.PHONY: \(.*\) #.*/\1/' | awk '{ print length }' | sort -nr | head -1); \
	grep '^.PHONY: .* #' Makefile | sed 's/\.PHONY: \(.*\) # \(.*\)/\1 \2/' | awk -v max=$$max '{ printf "%-*s ", max+1, $$1; for (i=2; i<=NF; i++) printf "%s ", $$i; print "" }'