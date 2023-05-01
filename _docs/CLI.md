# CLI

This document describes how to use the Book Management System (bms) command line interface (CLI).

{:toc}


## General

Overall usage

    bms <entity> <operation>

with
- `<entity>`: the entity to operate with
- `<operation>`: the operation to apply on the entity


## Books

Manage books

    bms book <operation>


### List

List books

    bms book list [--filter <filter>]

with
- `<filter>`: filter expression

### Create

Create book

    bms book create
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

    bms book show <book-id>

with
- `<book-id>`: ID of book


### Set

Show details about a book

    bms book set <book-id> <key> <value>

with
- `<book-id>`: ID of book


### Delete

Delete a book

    bms book delete <book-id>

with
- `<book-id>`: ID of book

