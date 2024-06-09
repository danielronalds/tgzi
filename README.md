# tgzi

An interactive file picker for compressing files into a 
tar.gz archive

![demo](./resources/tgzi-demo.gif)

## Installation

Can be installed with `go install`

```console
go install github.com/danielronalds/tgzi
```

## Usage

Simply call the program in the directory of the files you'd
like to compress, select the files you'd like and hit enter
to compress them. You'll be prompted for a name, but leaving
this blank results in the default `archive.tar.gz` being
created

```console
tgzi
```

### Keybindings

| Key    | Action                  |
| ------ | ----------------------- |
| up/k   |  Navigate up the list   |
| down/j |  Navigate down the list |
| space  |  Select a file          |
| A      |  Select all files       |
| enter  |  Compress files         |
| ?      |  Toggle help menu       |
