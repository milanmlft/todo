"""Tests for the todo CLI app."""

import json
from pathlib import Path
from typing import Any

import pytest
from todo import SUCCESS, __app_name__, __version__, cli
from todo.todo import Todoer
from typer.testing import CliRunner

runner = CliRunner()


def test_version() -> None:
    """Test if CLI app returns the correct version."""
    result = runner.invoke(cli.app, ["--version"])
    assert result.exit_code == 0
    assert f"{__app_name__} v{__version__}\n" in result.stdout


@pytest.fixture()
def mock_json_file(tmp_path: Path) -> Path:
    """Mock todo database JSON file."""
    todo = [{"Description": "Get some milk.", "Priority": 2, "Done": False}]
    db_file = tmp_path / "todo.json"

    with db_file.open("w") as db:
        json.dump(todo, db, indent=4)

    return db_file


test_data1 = {
    "description": ["Clean", "the", "house"],
    "priority": 1,
    "todo": {
        "Description": "Clean the house.",
        "Priority": 1,
        "Done": False,
    },
}
test_data2 = {
    "description": ["Wash the car"],
    "priority": 2,
    "todo": {
        "Description": "Wash the car.",
        "Priority": 2,
        "Done": False,
    },
}


@pytest.mark.parametrize(
    ("description", "priority", "expected"),
    [
        pytest.param(
            test_data1["description"],
            test_data1["priority"],
            (test_data1["todo"], SUCCESS),
        ),
        pytest.param(
            test_data2["description"],
            test_data2["priority"],
            (test_data2["todo"], SUCCESS),
        ),
    ],
)
def test_add(
    mock_json_file: Path, description: list[str], priority: int, expected: list[dict[str, Any]]
) -> None:
    """Test adding todo's to an existing database."""
    todoer = Todoer(mock_json_file)
    assert todoer.add(description, priority) == expected
    new_todo_list = todoer._db_handler.read_todos()
    assert len(new_todo_list) == 2
