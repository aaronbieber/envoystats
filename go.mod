module github.com/aaronbieber/envoystats

go 1.20

replace github.com/aaronbieber/carbonclient => ../carbonclient

replace github.com/aaronbieber/envoyclient => ../envoyclient

require (
	github.com/aaronbieber/carbonclient v0.0.0-00010101000000-000000000000
	github.com/aaronbieber/envoyclient v0.0.0-00010101000000-000000000000
	github.com/sirupsen/logrus v1.9.3
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/MacIt/pickle v1.0.0 // indirect
	github.com/boltdb/bolt v1.3.1 // indirect
	golang.org/x/sys v0.10.0 // indirect
)
