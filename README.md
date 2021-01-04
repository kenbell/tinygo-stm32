# TinyGo STM32 Library

Extended TinyGo support for STM32 MCUs.

## Philosophy / Purpose

TinyGo provides native support for a wide range of embedded systems, including a number of STM32 MCUs.  TinyGo has an emphasis on mass-market embedded devices, providing a consistent API to easily run apps on a range of different MCUs from many vendors.

STM32 MCUs support a wide range of configuration options both internally and externally with custom PCBs, where clock speeds, clock configurations, power management, can be adjusted to fit the end application.  This library aims to act as an extension to TinyGo making it easier to use TinyGo in these custom configurations.

The goals of this library are:

1. Integrate deeply with TinyGo (adopt TinyGo abstractions and interfaces, duplication logic only where essential).

2. Be (mostly) compatible with TinyGo so that any code written to use TinyGo, with only minor modification can be used on custom PCBs / specific scenarios.

3. Contribute upstream - upstream functionality where possible (and it makes sense) into TinyGo.

## Status: Experimental

This project is in an experimental phase - APIs and implementation are unstable / evolving rapidly.

## Getting Started

This project is currently split over three Git repos:

| Repo | Purpose |
|------|---------|
|[kenbell/tinygo](https://github.com/kenbell/tinygo)|Fork of TinyGo including hooks into TinyGo that enable this library.  The goal is to upstream these changes into TinyGo in time.|
|[kenbell/tinygo-stm32](https://github.com/kenbell/tinygo-stm32)|The main library.|
|[kenbell/tinygo-stm32-examples](https://github.com/kenbell/tinygo-stm32-examples)|Examples of using the library|

For now, the recommended approach is to clone all three repos adjacent to each other (in the same directory), then follow the instructions to [build TinyGo from source](https://tinygo.org/getting-started/linux/#source-install).
