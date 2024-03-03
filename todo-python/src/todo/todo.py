""""The todo model-controller."""

from pathlib import Path
from typing import Any, NamedTuple

from todo import DB_READ_ERROR, ID_ERROR
from todo.database import DatabaseHandler


class CurrentTodo(NamedTuple):
    """
    The data model for a given todo.

    Contains the todo's information in a dictionary and an error code, signaling the status of any
    operation performed on the todo.
    """

    todo: dict[str, Any]
    error: int


class Todoer:
    """The main Controller class, handling database connections."""

    def __init__(self: "Todoer", db_path: Path) -> None:
        """Initialise the Todoer class, composing it with an instance of DatabaseHandler."""
        self._db_handler = DatabaseHandler(db_path)  # <-- Composition!

    def add(self: "Todoer", description: list[str], priority: int = 2) -> CurrentTodo:
        """Add a new todo to the database."""
        description_text = " ".join(description)
        if not description_text.endswith("."):
            description_text += "."

        todo = {
            "Description": description_text,
            "Priority": priority,
            "Done": False,
        }

        read = self._db_handler.read_todos()
        if read.error == DB_READ_ERROR:
            return CurrentTodo(todo, read.error)

        read.todo_list.append(todo)
        write = self._db_handler.write_todos(read.todo_list)
        return CurrentTodo(todo, write.error)

    def get_todo_list(self: "Todoer") -> list[dict[str, Any]]:
        """Get the current list of todo's."""
        read = self._db_handler.read_todos()
        return read.todo_list

    def set_done(self: "Todoer", todo_id: int) -> CurrentTodo:
        """Set a todo as done."""
        read = self._db_handler.read_todos()

        if read.error == DB_READ_ERROR:
            return CurrentTodo({}, read.error)

        try:
            todo = read.todo_list[todo_id - 1]
        except IndexError:
            return CurrentTodo({}, ID_ERROR)

        todo["Done"] = True
        write = self._db_handler.write_todos(read.todo_list)
        return CurrentTodo(todo, write.error)

    def remove(self: "Todoer", todo_id: int) -> CurrentTodo:
        """Remove a todo from the database."""
        read = self._db_handler.read_todos()

        if read.error == DB_READ_ERROR:
            return CurrentTodo({}, read.error)

        try:
            todo = read.todo_list.pop(todo_id - 1)
        except IndexError:
            return CurrentTodo({}, ID_ERROR)

        write = self._db_handler.write_todos(read.todo_list)
        return CurrentTodo(todo, write.error)
