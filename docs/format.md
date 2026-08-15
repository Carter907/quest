# File Format

`.kng` files have a particular format that makes them ideal for housing knowledge.


## Guide Metadata

This archive format contains markdown files denoted as "Guides". These guides have a frontmatter that carries their metadata. When formatting these markdown files you should make sure all fields specified below are included.

> [!NOTE]
> When creating metadata about language itself, it's important to realize that we are rubbing up against philosophical and linguistic
> barriers. There is no universally accepted way of measuring "scope" or "clarity" in a piece of text. Therefore, defining these terms
> qualitatively gives us some breathing room for the fuzziness in their definitions.

**Properties**:

- `prerequisites` - a list of required guides of the *same* scope (horizontal edges in the knowledge graph).
- `sub_guides` - an optional list of encompassed guides of a *smaller* scope (vertical hierarchy). Guides of higher scope may encompass smaller guides.
- `clarity` - how dense in concepts the content itself is, dictating how `sub_guides` are represented in the text. 
  - strict (sub-guides appear 1-to-1 in the parent guide, using exact or near-exact content)
  - detailed (sub-guides are represented closely but with slight summarization)
  - introductory (sub-guides are generalized, pointing to the concept without exact reproduction)
  - vague (sub-guides are referenced loosely, with the least exact reproduction of content)
- `scope` - how much content is covered in a guide; how many concepts or things were explained.
  Scope is qualitative:
  - definition - smallest scope (singular term)
  - description - smaller scope using more examples and comparisons than a definition.
  - explanation - medium guide with multiple descriptions (average case)
  - lesson - multiple explanations (largest content type)

**Tags**:

Tags are subjects or topic descriptors that summarize the content of the entire guide.

**Title**:

The title of a guide is defined by its file name.
