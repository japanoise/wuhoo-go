# wuhoo-go - easy drawing canvas

This is a golang port of [wuhoo by Nick Vitsas][upstream]; a simple library
which uses native APIs to give you a drawing context & the primitives to create
an event loop.

## Status

At the moment, all that has been done is a fairly thin CGO wrapper around the
upstream library's software renderer, enough to get it to compile on Linux & to
port one of the example programs to Golang.  I have not tested building on MacOS
or Windows yet, but the compiler flags are present, so it may Just Work™.  There
has not currently been any work done on an OpenGL version.

## License

wuhoo-go, like the upstream library, is licensed MIT.

[upstream]: https://github.com/ViNeek/wuhoo
