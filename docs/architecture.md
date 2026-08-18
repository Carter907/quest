# Architecture

The universal framework that dictates the layout and design of the program is [cobra](https://github.com/spf13/cobra).

Subcommands live in their own file under the `cmd/` directory.

When commands require more involvement, [charmbracelet/huh](https://github.com/charmbracelet/huh) is used for interactive forms.

## Domain Model
