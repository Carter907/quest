# quest

quest is a command-line tool for storing and structuring knowledge. You use guides (markdown files) to explain and describe a certain topic or idea. This tool helps you structure and format them into a single file (a knowledge graph) with the file format .kng

.kng files can then be distributed and unpacked through quest and opened in Obsidian or some other application for viewing and editing.

## Ontology

- Guide - a single markdown file that explains just one concept and uses previous guides as support.
- Prerequisite - a guide that is required for understanding another guide.
- Knowledge Graph - a directed acyclic graph that stores a network of guides and their prerequisites. This data structure is similar to a [concept map](https://en.wikipedia.org/wiki/Concept_map), but stores different types of guides and their metadata instead of singular concepts, and edges flow from prerequisites to their upstream guides.

## Guide Metadata
**Properties**:

- `clarity` - how detailed the guide is overall.
- `scope` - how much content is covered in a guide; how many concepts or things were explained.
  Scope is qualitative:
  - Definition - smallest scope (singular term)
  - Description - smaller scope using more examples and comparisons than a definition.
  - Explanation - medium guide with multiple descriptions (average case)
  - Lesson - multiple explanations (largest content type)
- `prerequisites` - a list of required guides. This property establishes the edges in the knowledge graph. 
- `subject` - a topic descriptor applied to a guide.

**Tags**:

Tags are used directly but are still compatible with the format.
