# File Merger

[![version](https://img.shields.io/badge/version-1.0.0-green.svg)][semver]
[![Build Status](https://travis-ci.org/geektimus/file-merger.svg?branch=master)][travis_url]

This simple application allow the user to concatenate files specified on a json file

## Getting Started

These instructions will get you a copy of the app to be used on your local machine. See deployment for notes on how to use this simple application.

### Prerequisites

No prerequisites are needed to download and use this application, just follow the installation instructions and you will be ready to use the app
in no time. 😄

### Installing

- Clone the repository

```
git clone git@github.com:geektimus/file-merger.git
cd ~/file-merger
```

- Build the executable

```
go build ./...
```

This command will generate a file named "file-merger" (as the folder name), in case you may want to use a different name, you can run the following
command

```
go build -o <other-file-name> ./...
```

After these simple steps you should have a copy of the executable file.

## Running the tests

In case you want to run the tests it's as simple as running this command

```
go test -v ./...
```

In case you want to run some race condition test (Not applicable to this application, not for now), you can run this variation of the previous command

```
go test -v -race ./...
```

## Usage

This application expects three parameters, only one is mandatory

### Parameters

- Input filename: This defines the file name to be used as the source of the paths, it defines the structure that needs to parse before the app can read and merge the files. **required**
- Output filename: This defines the output file name that is going to be used as the result of the merge. **optional**
- Base path: Each path on the input file may be on a folder different from the folder where you are using the app so you need to specify the root path. In case you run the app from the cloned folder, this param should be 'testdata' since the files are under a structure under that folder.

### Examples

```
-- Use the merger to merge 😅 some files in my current folder
./file-merger -i workflow.json (the output file should be named output.txt which is the default)

-- Use the merger to ... some files under the test data folder and the output file should be queries.sql
./file-merger -i testdata/workflow.json -b testdata -o queries.sql

-- Doing the same using the long version of the flags.
./file-merger --inputFile testdata/workflow.json --basePath testdata --outputFile queries.sql
```

## Built With

- [Golang 1.11.x][golang] - Base language

## Pending Task

- Improve testing, adding more test to the other methods in main.
- Add more examples

## Contributing

Please read [CONTRIBUTING.md][contributing] for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer][semver] for versioning. For the versions available, see the [tags on this repository][project_tags].

## Authors

- **Alex Cano** - _Initial work_ - [Geektimus][profile]

See also the list of [contributors][project_contributors] who participated in this project.

## License

[![license](https://img.shields.io/badge/license-MIT-blue.svg)][license]

This project is licensed under the MIT License - see the [LICENSE.md][license] file for details

[travis_url]: https://travis-ci.org/geektimus/file-merger
[golang]: https://golang.org/doc/
[contributing]: CONTRIBUTING.md
[semver]: http://semver.org/
[project_tags]: https://github.com/geektimus/file-merger/tags
[profile]: https://github.com/Geektimus
[project_contributors]: https://github.com/geektimus/file-merger/graphs/contributors
[license]: LICENSE.md
