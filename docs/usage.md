# Usage

## Working with `qst`

### 1. Create a new knowledge graph with `qst new`

```sh
qst new Exponents
```

This will create a new directory called "Exponents" and populate it with a `quest.yaml` file and a **starter markdown file**.

### 2. Add your first guides

```sh
qst add "Understanding Exponents" --scope lesson --clarity introductory --tags math,algebra
```

This command creates a markdown file with populated front-matter. You can also specify prerequisites and sub-guides if you already know the structure you're going for.

### 3. Validate the structure

As you create more and more guides and start to draft out a presentable network of guides, you want to make sure the graph is valid under the constraints:

```sh
qst validate
```

`qst validate` will check for the following:

- Cycles
- Non-existent prerequisite
- Non-existent sub-guides
- Missing properties


### 4. Package your knowledge graph

Once you are ready to distribute your knowledge, you use the `qst form` command:

```sh
qst form
```

The current directory is chosen if you don't specify one. When you run this command, a `.kng` file with the same name as the chosen directory will be created. This file is a zip archive housing all markdown files and the `quest.yaml` metafile.

## Summary

This should give you an idea of what the workflow of curating knowledge graphs looks like. Getting the pedagogical structure of a knowledge right takes much more than `qst validate` and a single person. The reason why this CLI is focused on packaging knowledge is because we believe it should be open, free, and easily comprehensible. Sharing `.kng` files is how you can get more insight on how best to build a something presentable.
