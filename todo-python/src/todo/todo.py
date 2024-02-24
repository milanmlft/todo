""""The todo model-controller."""

from typinr import Any, NamedTuple


class CurrentTodo(NamedTuple):
    """
    The data model for a given todo.

    Contains the todo's information in a dictionary and an error code, signaling the status of any
    operation performed on the todo.
    """

    todo: dict[str, Any]
    error: int
