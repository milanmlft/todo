""""The todo model-controller."""

from pathlib import Path

from typinr import Any, NamedTuple

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
