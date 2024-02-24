"""Functionality for the to-do app database."""

import json
from configparser import ConfigParser
from pathlib import Path
from typing import Any, NamedTuple

from todo import DB_READ_ERROR, DB_WRITE_ERROR, JSON_ERROR, SUCCESS

DEFAULT_DB_FILE_PATH = Path.home().joinpath("." + Path.home().stem + "_todo.json")


def get_database_path(config_file: Path) -> Path:
    """Return the current path to the to-do database."""
    config_parser = ConfigParser()
    config_parser.read(config_file)
    return Path(config_parser["General"]["database"])


def init_database(db_path: Path) -> int:
    """Create the to-do database."""
    try:
        db_path.write_text("[]")  # Empty to-do list
    except OSError:
        return DB_WRITE_ERROR
    else:
        return SUCCESS


class DBResponse(NamedTuple):
    """
    The database response containing the list of todo's and an error code.

    :param todo_list: the list of todo's
    :param error: a numeric error code, to be mapped to `todo.ERRORS`
    """

    todo_list: list[dict[str, Any]]
    error: int


class DatabaseHandler:
    """Handle interactions with the to-do database."""

    def __init__(self: "DatabaseHandler", db_path: Path) -> None:
        """
        Initialise the database handler for a given database.

        :param db_path: the file path to the todo database.
        :type db_path: pathlib.Path

        """
        self._db_path = db_path

    def read_todos(self: "DatabaseHandler") -> DBResponse:
        """Read the JSON todo database and return a `DBResponse` instance."""
        try:
            with self._db_path.open("r") as db:
                try:
                    return DBResponse(json.load(db), SUCCESS)
                except json.JSONDecodeError:  # Catch wrong JSON format
                    return DBResponse([], JSON_ERROR)
        except OSError:  # Catch file IO errors
            return DBResponse([], DB_READ_ERROR)

    def write_todos(self: "DatabaseHandler", todo_list: list[dict[str, Any]]) -> DBResponse:
        """Write to the JSON todo database and return a `DBResponse` isntance."""
        try:
            with self._db_path.open("w") as db:
                json.dump(todo_list, db, indent=4)
                return DBResponse(todo_list, SUCCESS)
        except OSError:  # Catch file IO errors
            return DBResponse(todo_list, DB_WRITE_ERROR)
