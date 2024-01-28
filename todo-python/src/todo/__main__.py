"""To-Do entry point script."""

from todo import __app_name__, cli


def main() -> None:
    """Run the todo CLI app."""
    cli.app(prog_name=__app_name__)


if __name__ == "__main__":
    main()
