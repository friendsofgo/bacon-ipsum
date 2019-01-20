# Bacon Ipsum Cli
[![Build Status](https://travis-ci.org/friendsofgo/bacon-ipsum.svg?branch=master)](https://travis-ci.org/friendsofgo/bacon)

This project is purely formative, It has been created in order to illustrate the examples of your respective blog article.

This cli use the Banc Ipsum API for generate random texts and print in the output.

## Getting started

To install the library, run

```sh
go get -u github.com/friendsofgo/bacon-ipsum/cmd/bacon-ipsum
```

## Use the cli

```sh
$ bacon-ipsum --help
usage: bacon-ipsum [OPTIONS]
        bacon-ipsum is a simple tool to generate random text based on a bacon ipsum API

  -paras int
        number of paragraphs (default 5)
  -sentences int
        number of sentences (this overrides paragraphs)
  -type string
        Type of the text to generate (Required) [Valid options: all-meat, meat-and-filler]
  -withLorem
        if it is true the first paragraph start with 'Bacon dolor sit amet'
```
## Contribute
[Contributions](https://github.com/friendsofgo/bacon-ipsum/issues?q=is%3Aissue+is%3Aopen) are more than welcome, if you are interested please fork this repo and send your Pull Request.

## License
MIT License, see [LICENSE](https://github.com/friendsofgo/bacon-ipsum/blob/master/LICENSE)