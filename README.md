# Envoystats

A simple program to collect solar production and consumption stats from my
Enphase Envoy gateway and send them to Graphite. This program depends on two
other libraries I've written, [envoyclient][ec] and [carbonclient][cc], to
collect the stats from the Envoy and send them to Carbon, respectively.

[ec]: https://github.com/aaronbieber/envoyclient
[cc]: https://github.com/aaronbieber/carbonclient

## Building

Provided you have Go installed and that it's somewhere near version 1.20.6,
building should be as simple as `go build`. To build for other, potentially
older systems (as I do), you may want to disable Cgo so that core C libraries
(like glibc) are statically linked. I've included a make target for that
scenario, so `make build` should do the trick.

It is said that disabling Cgo results in larger binaries, which stands to
reason, but in my own experience with this program, it results in a *smaller
binary*, YMMV.

## Usage

You'll need to set up your own configuration file. Make a copy of the defaults
and edit to your taste:

```sh
$ cp config.default.yml config.yml
```

Then run the program. That's honestly all there is to it.

## License

```text
        DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE 
                    Version 2, December 2004 

 Copyright (C) 2004 Sam Hocevar <sam@hocevar.net> 

 Everyone is permitted to copy and distribute verbatim or modified 
 copies of this license document, and changing it is allowed as long 
 as the name is changed. 

            DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE 
   TERMS AND CONDITIONS FOR COPYING, DISTRIBUTION AND MODIFICATION 

  0. You just DO WHAT THE FUCK YOU WANT TO.
```
