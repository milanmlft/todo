"""Tests the config functionality for the to-do app."""


from configparser import ConfigParser
from pathlib import Path

from todo import SUCCESS
from todo.config import CONFIG_FILE_PATH, init_app


def test_init_app(tmp_path: Path) -> None:
    """Tests whether the app config can be initialised."""
    # Arrange
    config = ConfigParser()

    # Act
    result = init_app(str(tmp_path))
    config.read(CONFIG_FILE_PATH)
    expected_db_path = Path(config["General"]["database"])

    # Assert
    assert result == SUCCESS
    assert CONFIG_FILE_PATH.is_file()
    assert CONFIG_FILE_PATH.exists()
    assert expected_db_path.exists()
    assert expected_db_path == tmp_path
