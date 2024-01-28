"""Functionality for the to-do app database."""

from configparser import ConfigParser
from pathlib import Path

from todo import DB_WRITE_ERROR, SUCCESS

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
