# Huffman coding

A CLI tool for encoding or decoding .txt files using the huffman algorithm.

(The output is in .huff, but it can easily be opened as a .txt file)

## Usage
```bash
    <encode|decode> <-f|-d> <file/directory> [-o <output>]
```
## Prerequisites

- [Go](https://go.dev).

## Installation

To get started with this project, clone the repository and run (or compile, then run) using the `go` command.

```bash
git clone https://github.com/ikugox/huffman-coding.git
cd huffman-coding
```
You can run the code from `huffman-coding/` using
```bash
go run src/*.go
```
or
```bash
go build -o huffman src/*.go
```
and then run with
``` ./huffman ```
