[![tests](https://github.com/khlopko/assure/actions/workflows/ci.yml/badge.svg?branch=main&event=push)](https://github.com/khlopko/assure/actions/workflows/ci.yml)

# assure

Go assertions library.

## Features

- Disable assertions (e.g. when in production).
- Configure default fail handler with your own.
- Create scoped contexts with their own handlers.

## Implementation details

`Assert` function uses a closure to run evaluation. This is made on purpose to be able to eliminate unnecessary
evaluations when assertions are disabled, since Go (at least with my current understanding) cannot express what is
possible with, say, build flags – as build tags doesn't propagate.

