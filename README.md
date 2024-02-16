# Envoystats

A simple program to collect solar production and consumption stats from my
Enphase Envoy gateway and send them to Graphite. This program depends on two
other libraries I've written, [envoyclient][ec] and [carbonclient][cc], to
collect the stats from the Envoy and send them to Carbon, respectively.

[ec]: https://github.com/aaronbieber/envoyclient
[cc]: https://github.com/aaronbieber/carbonclient

## Overview

The Envoy device serves a web portal, and within it you can find a link to a
realtime energy stats feed, which is spit out as a JSON document. Unfortunately,
newer Envoy devices (like mine), now require authentication to your Enphase
account on the internet, which provides you a token that you use to access the
local portal. [There is more information here][guytec].

[guytec]: https://guytec.com/Envoy-S/

I lament the move in this direction, but this was the software version my Envoy
was installed with, so I have no alternatives available. For that reason, this
software is built to do the cloud authentication via the Enphase site (the same
one you log into to view your live energy stats), acquire the token, and then
talk to the Envoy directly.

It's a shame that an internet round-trip is necessary just to get local data off
your local device, but the tokens live for a while, so at least the energy
readings can be very fast, local requests.

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
