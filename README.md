# AVRGo - Go compiler for AVR microcontroller family

![technology Go](https://img.shields.io/badge/technology-go-blue.svg) 
[![GoReportCard](https://goreportcard.com/badge/github.com/avrgo-org/avrgo)](https://goreportcard.com/report/avrgo-org/avrgo)
[![Test Coverage](https://api.codeclimate.com/v1/badges/06627dd8eca5aaebed94/test_coverage)](https://codeclimate.com/github/avrgo-org/avrgo/test_coverage)
[![Build Status](https://travis-ci.org/avrgo-org/avrgo.svg?branch=master)](https://travis-ci.org/avrgo-org/avrgo)
[![Maintainability](https://api.codeclimate.com/v1/badges/06627dd8eca5aaebed94/maintainability)](https://codeclimate.com/github/avrgo-org/avrgo/maintainability)


AVRGo is a Go compiler to use for AVR type microcontrollers.

AVRGo initially combine the concept of :
* AVR-GCC (GNU C compiler for AVR microcontroller)

* TinyGo which is reuses libraries used by the [Go language tools](https://golang.org/pkg/go/) alongside [LLVM](http://llvm.org) to provide an alternative way to compile programs written in the Go programming language.

## Installation
Todo installation instructure.

## Supported boards/targets
Bellow are Atmel's microcontrollers which will be supported 

* [Atmel SAM D21/E/G/J](https://cdn.sparkfun.com/datasheets/Dev/Arduino/Boards/Atmel-42181-SAM-D21_Datasheet.pdf)

## Documentation
Todo documentation.

## Contributing

Your contributions are welcome!

Please take a look at our [CONTRIBUTING.md](./CONTRIBUTING.md) document for details.

## Project Scope

Goals:

* Have very small binary sizes. Don't pay for what you don't use.
* Support for AVR microcontrollers family only.
* Good CGo support, with no more overhead than a regular function call.
* Support most standard library packages and compile most Go code without modification.

## License

This project is licensed under the BSD 3-clause license, just like the [Go project](https://golang.org/LICENSE) itself.
