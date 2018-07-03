# README #

This is my attempt at solving [Programming Praxis](https://programmingpraxis.com/) problems using [golang](https://golang.org).

## Requisites:

golang; I've used 1.10.3.

## To Build:

Put the project in a folder named programmingpraxis inside $GOPATH/src and then run:

    go build programmingpraxis/cmd/rpncalculator

## To run tests:

    go build programmingpraxis/cmd/rpncalculator

## To run binaries for a problem:

    go install programmingpraxis/cmd/rpncalculator
    $GOPATH/bin/rpncalculator