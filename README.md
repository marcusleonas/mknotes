# mknotes

A cli i made to help manage my notes. I tend to use it with
neovim to speed up my notetaking.

## Installation

Download the latest release and put it in your PATH.

## Commands

### `init`

Initialises a new "vault".

**Flags:**

- `-d, --directory` the directory for the new vault. required.
- `-g, --git` create a git repo inside the new folder. default `false`.

### `new`

Creates a new note using a template.

**Flags:**

- `-n, --name` name of the note. required.
- `-t, --template` name of the template. default `default`

## Templates

mknotes has a useful templating feature to fill data when creating notes.
It currently only has a couple of options. If you have any suggestions,
let me know by opening a feature request!

- `Name` name of the note.
- `Timestamp` timestamp of when the note was made in the RFC3339 format.
