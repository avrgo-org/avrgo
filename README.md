# AVRGo - Go compiler for AVR microcontroller family

AVRGo is a Go compiler to use for specific AVR type microcontrollers.

AVRGo initially using the same concept with TinyGo which is reuses libraries used by the [Go language tools](https://golang.org/pkg/go/) alongside [LLVM](http://llvm.org) to provide an alternative way to compile programs written in the Go programming language.

## Installation
Todo installation instructure.

## Supported boards/targets
Bellow are microcontrollers which will be supported 

* [Arduino Nano33 IoT](https://store.arduino.cc/nano-33-iot)

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
