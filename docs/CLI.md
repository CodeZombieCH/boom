# CLI

This document describes how to use the Book Management System (boom) command line interface (CLI).

{:toc}


## General

Overall usage

    boom <entity> <operation>

with
- `<entity>`: the entity to operate with
- `<operation>`: the operation to apply on the entity


## Books

Manage books

    boom book <operation>


### List

List books

    boom book list [--filter <filter>]

with
- `<filter>`: filter expression

### Create

Create book

    boom book create
        --title <title>
        --author <author>
        --published-date <published-date>
        --edition <edition>
        --description <description>
        --genre <genre>

with
- `<filter>`: filter expression

### Show

Show details about a book

    boom book show <book-id>

with
- `<book-id>`: ID of book


### Set

Show details about a book

    boom book set <book-id> <key> <value>

with
- `<book-id>`: ID of book


### Delete

Delete a book

    boom book delete <book-id>

with
- `<book-id>`: ID of book



## Collections


Manage books

    boom book <operation>


### List

List books

    boom book list [--filter <filter>]

with
- `<filter>`: filter expression

### Create

Create book

    boom book create
        --title <title>
        --author <author>
        --published-date <published-date>
        --edition <edition>
        --description <description>
        --genre <genre>

with
- `<filter>`: filter expression

### Show

Show details about a book

    boom book show <book-id>

with
- `<book-id>`: ID of book


### Set

Show details about a book

    boom book set <book-id> <key> <value>

with
- `<book-id>`: ID of book


### Delete

Delete a book

    boom book delete <book-id>

with
- `<book-id>`: ID of book


