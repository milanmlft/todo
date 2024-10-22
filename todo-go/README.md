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
- [ ] Support setting priority in commands:
    - [ ] Add optional flag to `add` to set priority of new todo
    - [ ] Add `set-priority` to change priority of an existing todo
- [ ] Add `remove` command to remove a todo
