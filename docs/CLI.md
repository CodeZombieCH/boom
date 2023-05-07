# CLI

This document describes how to use the Book Management System (boom) command line interface (CLI).

* TOC
{:toc}


## General

Overall usage

    boom <entity> <operation>

with
- `<entity>`: the entity to operate with
- `<operation>`: the operation to apply on the entity


### Filter

List operations usually support filtering of the result, offering
- a simple substring filter against the main property of the entity
- a filter expression

#### Substring Filter

Allows filtering by the entities main field.

Example:

Filter book where *title" matches the substring "for dummies":

    boom book list "for dummies"

#### Filter Expression
Allows granular filtering by the entities properties. Multiple conditions can be expressed by combining them using a "," (comma).

Example:

Filter books where genre is "fiction" and publication date is before 1900

    boom book list --filter "genre=fiction,publication-date<1900-01-01"


## Books

Manage books

    boom book <operation>


### List

List books

    boom book list [<title-filter>] [--filter <filter>]

with
- `<title-filter>`: substring filter applied to the *title* property
- `<filter>`: filter expression, see [Filter](#filter)

### Create

Create book

    boom book create
        --title <title>
        --author <author>
        --publication-date <publication-date>
        --edition <edition>
        --description <description>
        --genre <genre>

with
- `<title>`: title of book
- `<author>`: author of book
- `<publication-date>`: publication date of book in the format YYY-MM-DD
- `<edition>`: edition of the book
- `<description>`: description of book
- `<genre>`: genre of book

### Show

Show details about a book

    boom book show <book-id>

with
- `<book-id>`: ID of book

### Set

Set a property of a book

    boom book set <book-id> <key> <value>

with
- `<book-id>`: ID of book
- `<key>`: key of the property to set
- `<value>`: value of the property to set

### Delete

Delete a book

    boom book delete <book-id>

with
- `<book-id>`: ID of book


## Collections

Manage collections of books

    boom collection <operation>


### List

List collections

    boom collection list [<name-filter>] [--filter <filter>]

with
- `<name-filter>`: substring filter applied to the *name* property
- `<filter>`: filter expression, see [Filter](#filter)

with
- `<filter>`: filter expression

### Create

Create a new collection

    boom collection create
        --name <name>

with
- `<name>`: name of the collection

### Show

Show details about a collection

    boom collection show <collection-id>

with
- `<collection-id>`: ID of collection


### Set

Set a property of a collection

    boom collection set <collection-id> <key> <value>

with
- `<collection-id>`: ID of collection
- `<key>`: key of the property to set
- `<value>`: value of the property to set


### Delete

Delete a collection

    boom collection delete <collection-id>

with
- `<collection-id>`: ID of collection

### Add Book

Add a book to a collection

    boom collection add <collection-id> <book-id>

with
- `<collection-id>`: ID of collection
- `<book-id>`: ID of book


### Remove Book

Remove a book from a collection

    boom collection remove <collection-id> <book-id>

with
- `<collection-id>`: ID of collection
- `<book-id>`: ID of book

### List Books

List books of a collection

    boom collection list-books <collection-id> [<title-filter>] [--filter <filter>]

with
- `<collection-id>`: ID of collection
- `<title-filter>`: substring filter applied to the *title* property
- `<filter>`: filter expression, see [Filter](#filter)

