# quest-cli

Quest is a command-line tool for storing and structuring knowledge. You use guides (markdown files) to explain and describe a certain topic or idea. This tool helps you structure and format them into a single file format.

These files can be distributed and unpacked in markdown editors like Obsidian. Integration with this tool in your workflow is made easy through is simple yet comprehensive command-line interface.


## Why Quest?

Quest facilitates the distribution and management of knowledge.

You can use the file format on education platforms for curriculum building and learning path creation (walks on the graph).

Being able to structure information using a standard format is important. Quest helps decrease uncertainty in the learning process by proposing and helping to enforce a structural standard for markdown content. Quest enforces this standard format using `qst validate`, which checks knowledge graph files for any formatting or structural violations (like cycles).

By offloading validation of content to a concrete system, you can focus on comprehending and building upon learning material.


## Installation & Getting Started

**Build from source and install locally**:

```sh
go build -o $(go env GOPATH)/bin/qst .
```

**Running `qst`**:

```sh
qst --help
```

## Concepts

Check out [concepts.md](/docs/concepts.md) for more theoretical foundations.

## Format

Find formatting constraints and metadata specifics in [format.md](/docs/format.md)

## CLI Docs

Visit [the cli docs](/docs/cli/) for comprehensive breakdown of the current `qst` functionality
