"""CLI for the Python To-Do app."""


from pathlib import Path
from typing import Optional

import typer

from todo import ERRORS, __app_name__, __version__, config, database, todo

app = typer.Typer()


@app.command()
def init(
    db_path: str = typer.Option(
        str(database.DEFAULT_DB_FILE_PATH),
        "--db-path",
        "-db",
        prompt="to-do databae location?",
    ),
) -> None:
    """Initialise the to-do database."""
    app_init_error = config.init_app(db_path)
    if app_init_error:
        typer.secho(
            f'Creating config file failed with "{ERRORS[app_init_error]}"',
            fg=typer.colors.RED,
        )
        raise typer.Exit(1)

    db_init_error = database.init_database(Path(db_path))
    if db_init_error:
        typer.secho(f'Creating the to-do database failed with "{ERRORS[db_init_error]}"')
        raise typer.Exit(1)

    typer.secho(f"The to-do databse is {db_path}", fg=typer.colors.GREEN)


def get_todoer() -> todo.Todoer:
    """Get the  todo controller."""
    if not config.CONFIG_FILE_PATH.exists():
        typer.secho("Config file not found, please run `todo init`.", fg=typer.colors.RED)
        raise typer.Exit(1)

    db_path = database.get_database_path(config.CONFIG_FILE_PATH)

    if not db_path.exists():
        typer.secho("Database file not found, please run `todo init`.", fg=typer.colors.RED)
        raise typer.Exit(1)

    return todo.Todoer(db_path)


@app.command()
def add(
    description: list[str] = typer.Argument(...),  # noqa: B008, functiona calls in argument
    priority: int = typer.Option(2, "--priority", "-p", min=1, max=3),
) -> None:
    """Add a new todo."""
    todoer = get_todoer()
    todo, error = todoer.add(description, priority)

    if error:
        typer.secho(f'Adding to-do failed with "{ERRORS[error]}"', fg=typer.colors.RED)
        raise typer.Exit(1)
    typer.secho(
        f"""to-do: "{todo['Description']}" was added """ f"""with priority: {priority}""",
        fg=typer.colors.GREEN,
    )


def _version_callback(*, value: bool) -> None:
    if value:
        typer.echo(f"{__app_name__} v{__version__}")
        raise typer.Exit


@app.command("list")
def list_all() -> None:
    """List all to-do's."""
    todoer = get_todoer()
    todo_list = todoer.get_todo_list()
    if len(todo_list) == 0:
        typer.secho("No to-do's found.", fg=typer.colors.RED)
        raise typer.Exit

    typer.secho("\nto-do list:\n", fg=typer.colors.BLUE)
    columns = (
        "ID  ",
        "| Priority  ",
        "| Done  ",
        "| Description  ",
    )
    headers = "".join(columns)
    typer.secho(headers, fg=typer.colors.BLUE, bold=True)
    typer.secho("-" * len(headers), fg=typer.colors.BLUE)

    # Loop over todo list and assign ID to each one
    for idx, item in enumerate(todo_list, start=1):
        desc, priority, done = item.values()
        typer.secho(
            # print the id, priority, done and description with proper padding
            f"{idx}{(len(columns[0]) - len(str(idx))) * ' '}"
            f"| ({priority}){(len(columns[1]) - len(str(priority)) - 4) * ' '}"
            f"| {done}{(len(columns[2]) - len(str(done)) - 2) * ' '}"
            f"| {desc}",
            fg=typer.colors.CYAN if done else typer.colors.YELLOW,
        )
    typer.secho("-" * len(headers) + "\n", fg=typer.colors.BLUE)


@app.callback()
def main(
    version: Optional[bool] = typer.Option(  # noqa: ARG001, UP007
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
