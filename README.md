# Auto Move

A tool to move files based on a set of rules, defined by you. The rules are 
encoded in a JSON file, and many rules can be written (each one in a 
respective) .json file.

Some examples of rules are given in the rules directory and each rule is of a 
respective _type_ with an empty predicate.

## Rules

As said before, each rule is defined in a json file and the basic structure of 
a given rule is defined below:

```
{
    "dirs": {
        "watch": [
            ""
        ],
        "to": ""
    },

    "rule": {
        "type": "",
        "predicate": []
    }
}
```

Auto move can watch an array of directories as show inside the `dirs` object and 
all of the files inside of it, that follows the given predicate, will be moved to
the `to` directory.

A rule is made up of two things, a **type** and a **predicate**. A type can be 
one of the four defined below:

- filetype
- filename
- prefix
- suffix

More types can (and might) be implemented in the code. 

The types itself - well, defines the type of the rule - instructs how a predicate 
will be writtten. A filetype rule will accept predicates of file extensions to be
used, that means the auto-move will look for all defined file extensions in 
predicate in order to know what has to be moved or not.

A filename rule roughly interpreted as a _regex_, although is not a _regex_ per say
it will follow a predicate pattern in order to know what file has to be moved. There
are only two types of predicates that can be used to represent a filename: an unicode 
letter that is represented in the predicate with the letter `L` and a digit that is 
represented in the predicate with the letter `D`. Combined, the two patterns can be
used to match the following filenames:

- `filename-2018-05-23.txt` can be matched with `LLLLLLLL-DDDD-DD-DD.txt`
- `coração-1.txt` can be matched with `LLLLLLL-D.txt`

All alphanumeric digits which aren't specified with the `L` or `D` have to
be math exactly as they are in the filename. One side note on this type is that I 
haven't tried this with non latin letters so be aware that problems might occur. 

The prefix and suffix types are quite easy to understand as they just filter the
prefix or suffix alphanumeric digits of a filename. Note that these two types 
doesn't implement the matching rules of the previously defined type.

## Build

To build the code you can simply call `go` directly with a build instruction or
use the defined Makefile in Unix like systems which support make. An example of
how the project can be build is shown below for reference.

`go build -o auto-move.exe main.go config.go move.go watch.go`