# Todo list in Go

A Todo list manager written in Go.

## Roadmap

- [x] Implement core functionality for managing todos
- [x] Set up CLI framework with `cobra`
- [x] Add database handler to make todos persist
- [x] Add basic CLI commands to initialise and add todos
- [x] Add `list` command to print current todos
- [x] Add `done` command to mark todos as done
- [ ] Add flag to `list` to filter on "done/not-done" todos
- [x] Add optional flag to `add` to set priority of new todo
- [x] Add `remove` command to remove a todo
- [ ] Use an `enum` for priorities with labels `"low", "medium", "high"` and use these in printing
- [ ] Add `edit` command to modify an existing todo
