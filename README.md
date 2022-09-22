# diolibgen
Diolibgen generate drawio library from svg files. This tool also insert style tag to svg if not exists.
## Quick start
+ Install golan >= 1.17
+ Run command below to install external library
```
~$ go mod tidy
```
+ Deploy svg files to input directory
+ Execute this tool
```
~$ go run main.go
# or
~$ go buid main.go && ./main
```
## File structure
```
.
├── README.md
├── go.mod
├── .gitignore
├── main.go # entry point
├── diolibgen
│   └── diolibgen.go # main process
├── svg
│   ├── svg_test.go
│   ├── svg.go # manage svg file
│   ├── svgReader_test.go
│   └── svgReader.go # svg file reader
├── mxlibrary
│   ├── mxGraphModelTemplate.go
│   ├── mxlibrary_test.go
│   └── mxlibrary.go # manage mxlibrary
├── output # output directory
│   └── .gitkeep
└── input # target svg files
    └── .gitkeep
```