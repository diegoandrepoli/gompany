# Gompany
Gompany as golang company manager, this application combine different data sources. Importing data on setup with csv files and dispose API to add and update companies.

## Requisites
To use Gompany application your need Golang, get [here](https://golang.org/dl/).

## Setup application
Clone this repository with command git clone or download sorce code.

```console
git clone https://github.com/diegoandrepoli/gompany.git
```
This command get project dependences, up to postgresql database on docker and import start data with csv files. To setup run makefile with parameter setup.

```console
make setup
```

## Start application
This command to start server application, for this execute makefile with start parameter.

```console
make start
```

## Running tests
To running tests application execute makefile with parameter check.

```console
make check
```
