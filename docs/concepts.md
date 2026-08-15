# Concepts

Here you'll find all the theoretical foundations for quest and the system that it's based on. This foundation is based on largely on the discussions that occurred in the [Bluelearn](https://github.com/bluelearn-org/bluelearn) project.

## Terms

- Guide - a single markdown file that explains just one concept and may use previous guides as support.
- Prerequisite - a required guide for understanding another guide. Prerequisites must have the *same scope* as the guide that requires them.
- Clarity - How detailed the content is overall. Strict clarity represents a 1-to-1 with all sub-guides. This means that you'll find a sub-guides content explicitly used in the larger guide.
- Scope - The kind of content goal for a guide. It represents overarching objective for what is how much is to be communicated to the reader.
- Sub-guide - a smaller-scoped guide that is encompassed by a larger-scoped guide.
- Knowledge Graph - a [Directed acyclic graph](https://en.wikipedia.org/wiki/Directed_acyclic_graph) that stores a network of guides. Edges are divided into horizontal connections (`prerequisites`) between guides of the same scope, and vertical connections (`sub_guides`) from larger-scoped guides to smaller ones. This idea is similar to a [Citation graph](https://en.wikipedia.org/wiki/Citation_graph), but instead of any kind of document, guides are used with structural constraints in place for their content and metadata.
