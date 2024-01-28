"""Tests for the todo CLI app."""

from todo import __app_name__, __version__, cli
from typer.testing import CliRunner

runner = CliRunner()


def test_version() -> None:
    """Test if CLI app returns the correct version."""
    result = runner.invoke(cli.app, ["--version"])
    assert result.exit_code == 0
    assert f"{__app_name__} v{__version__}\n" in result.stdout
