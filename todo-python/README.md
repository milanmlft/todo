# Todo list in Python

[![pre-commit](https://img.shields.io/badge/pre--commit-enabled-brightgreen?logo=pre-commit&logoColor=white)](https://github.com/pre-commit/pre-commit)
[![Tests status][tests-badge]][tests-link]
[![Linting status][linting-badge]][linting-link]
[![License][license-badge]](./LICENSE.md)

<!-- prettier-ignore-start -->
[tests-badge]:              https://github.com/milanmlft/todo/actions/workflows/python-tests.yml/badge.svg
[tests-link]:               https://github.com/milanmlft/todo/actions/workflows/python-tests.yml
[linting-badge]:            https://github.com/milanmlft/todo/actions/workflows/python-linting.yml/badge.svg
[linting-link]:             https://github.com/milanmlft/todo/actions/workflows/python_linting.yml
<!-- prettier-ignore-end -->

A Todo list manager written in Python

## About

### Project Team

Milan Malfait ([milan.malfait94@gmail.com](mailto:milan.malfait94@gmail.com))

### Research Software Engineering Contact

## Getting Started

### Prerequisites

`todo-python` requires Python 3.12.

### Installation

We recommend installing in a project specific virtual environment created using a environment management tool such as [Mamba](https://mamba.readthedocs.io/en/latest/user_guide/mamba.html) or [Conda](https://conda.io/projects/conda/en/latest/). To install the latest development version of `todo-python` using `pip` in the currently active environment run

```sh
git clone https://github.com/milanmlft/todo-python.git
```

and then install in editable mode by running

```sh
pip install -e .
```

### Running Locally

How to run the application on your local system.

### Running Tests

<!-- How to run tests on your local system. -->

Tests can be run across all compatible Python versions in isolated environments using
[`tox`](https://tox.wiki/en/latest/) by running

```sh
tox
```

To run tests manually in a Python environment with `pytest` installed run

```sh
pytest tests
```

again from the root of the repository.

## Roadmap

- [x] Set up initial project structure
- [x] Set up app initialisation
- [ ] Set up todo back-end
- [ ] Implement adding and listing todo's
- [ ] Implement todo completion functionality
- [ ] Implement removing todo's
- [ ] Improve todo displaying

## Acknowledgements

This work was funded by a grant from the JBFC: The Joe Bloggs Funding Council.
