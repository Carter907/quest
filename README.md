# quest

Quest is a command-line tool for storing and structuring knowledge. You use guides (markdown files) to explain and describe a certain topic or idea. This tool helps you structure and format them into a single file (a knowledge graph) with the file format .kng

.kng files can then be distributed and unpacked through quest and opened in Obsidian or some other application for viewing and editing.

- read [concepts.md](/docs/concepts.md) for more theoretical foundations.
- read [format.md](/docs/format.md) to understand the .kng file format and guide metadata

# Why Quest?

Quest facilitates the distribution and management of knowledge.

You can use the file format on education platforms for curriculum building and learning path creation (walks on the graph).

Being able to structure information using a standard format is important because uncertainty tends to get in the way of comprehension when learning something new.

You can check that that format has been followed, allowing you to be more confident about what you're reading. This standardization of knowledge allows you to focus on perfecting the content itself within defined constraints.


## Getting Started

**Building the binary**:

```sh
go build -o qst
```

**Running `qst`**:

```sh
./qst
```
