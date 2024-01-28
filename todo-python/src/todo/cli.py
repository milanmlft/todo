"""CLI for the Python To-Do app."""


import typer

from todo import __app_name__, __version__

app = typer.Typer()


def _version_callback(*, value: bool) -> None:
    if value:
        typer.echo(f"{__app_name__} v{__version__}")
        raise typer.Exit


@app.callback()
def main(
    version: bool | None = typer.Option(  # noqa: ARG001
        None,
        "--version",
        "-v",
        help="Show the application's version and exit.",
        callback=_version_callback,
        is_eager=True,
    ),
) -> None:
    """CLI arguments for the main todo command."""
    return