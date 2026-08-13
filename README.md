# quest

Quest is a command-line tool for storing and structuring knowledge. You use guides (markdown files) to explain and describe a certain topic or idea. This tool helps you structure and format them into a single file (a knowledge graph) with the file format .kng

.kng files can then be distributed and unpacked through quest and opened in Obsidian or some other application for viewing and editing.

## Why Quest?

Quest facilitates the distribution and management of knowledge.

You can use the file format on education platforms for curriculum building and learning path creation (walks on the graph).

Being able to structure information using a standard format is important because uncertainty tends to get in the way of comprehension when learning something new.

You can check that that format has been followed, allowing you to be more confident about what you're reading. This standardization of knowledge allows you to focus on perfecting the content itself within defined constraints.

## Terms

- Guide - a single markdown file that explains just one concept and may use previous guides as support.
- Prerequisite - a guide that is required for understanding another guide.
- Knowledge Graph - a [Directed acyclic graph](https://en.wikipedia.org/wiki/Directed_acyclic_graph) that stores a network of guides and their prerequisites. This idea is similar to a [Citation graph](https://en.wikipedia.org/wiki/Citation_graph), but stores general guide documents and edges are directed toward the documents that reference them.

## Guide Metadata

> [!NOTE]
> When creating metadata about language itself, it's important to realize that we are rubbing up against philosophical and linguistic
> barriers. There is no universally accepted way of measuring "scope" or "clarity" in a piece of text. Therefore, defining these terms
> qualitatively gives us some breathing room for the fuzziness in their definitions.

**Properties**:

- `prerequisites` - a list of required guides. This property establishes the edges in the knowledge graph.
- `clarity` - how dense is concepts is the content itself. How much detail appears in the average sentence.
  - vague
  - introductory
  - detailed
  - strict
- `scope` - how much content is covered in a guide; how many concepts or things were explained.
  Scope is qualitative:
  - Definition - smallest scope (singular term)
  - Description - smaller scope using more examples and comparisons than a definition.
  - Explanation - medium guide with multiple descriptions (average case)
  - Lesson - multiple explanations (largest content type)

**Tags**:

Tags are subjects or topic descriptors that summarize the content of the entire guide.

**Title**:

The title of a guide is defined by its file name.
