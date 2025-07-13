# Sapphire

## Installation

`sapphire` requires [`go`](https://go.dev/doc/install).

## Usage

Here's a list of commands that you can run:

```txt
clean         Remove the executable files that had been created 
build         Compile your application code into binary code that can be run as an executable 
dev           Run the application in development mode without having to build and run the executable file 
run           Build the application and run the compiled executable file 
deps          Install the project dependencies 
install-lint  Manually download the linter's executable as recommended by gloangci-lint maintainers 
install       Set up the project by installing all of the necessary dependencies and the linter 
format        Format all of your Go code by adhering to conventions 
lint          Run the linter to check if code isn't complying to style guides 
test          Run all of the tests written in the project 
help          Generate list of targets with descriptions
```

## Development

### Setup

1. `git clone https://github.com/nrabhiram/go-invade.git`
2. Run `make install` to install of the project dependencies
3. Build the project for production: `make build`
4. Run the project locally: `make dev`

### Dev Loop

- `make format` → run the autoformatter
- `make lint` → run the linter
- `make test` → run the specs
- `make build` → build the application for distribution
- `make dev` → run the application in development mode

## License

The project is available as open source under the terms of the [MIT License](LICENSE).

## Author

[![Twitter](https://img.shields.io/badge/follow-%40nrabhiram-1DA1F2?style=flat&logo=Twitter)](https://twitter.com/nrabhiram)
[![LinkedIn](https://img.shields.io/badge/connect-%40abhiramreddy-%230077B5?style=flat&logo=LinkedIn)](https://www.linkedin.com/in/abhiram-reddy-23285b196/)