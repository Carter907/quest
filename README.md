# quest-cli

Quest is a command-line tool for storing and structuring knowledge. You use guides (markdown files) to explain and describe a certain topic or idea. This tool helps you structure and format them into a single file (a knowledge graph) with the file format .kng

.kng files can then be distributed and unpacked through quest and opened in Obsidian or some other application for viewing and editing.


## Why Quest?

Quest facilitates the distribution and management of knowledge.

You can use the file format on education platforms for curriculum building and learning path creation (walks on the graph).

Being able to structure information using a standard format is important because otherwise uncertainty is more likely to hamper your ability to learn. Quest enforces this standard format using `qst form`, which validates and packs guides into knowledge graph files.

By offloading validation of content to a concrete system, you can focus on comprehending and building upon the content.


## Installation & Getting Started

**Install via `go install`**:

```sh
go install github.com/Carter907/quest-cli@latest
```

**Or build from source**:

```sh
go build -o qst
```

**Running `qst`**:

```sh
./qst
```

## Concepts

Check out [concepts.md](/docs/concepts.md) for more theoretical foundations.

## Format

Find formatting constraints and metadata specifics in [format.md](/docs/format.md)

## CLI Docs

Visit [the cli docs](/docs/cli/) for comprehensive breakdown of the current `qst` functionality
