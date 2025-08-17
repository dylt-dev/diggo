Go over to daylight and push the latest code (check out Howto-release-nightly.md)

Grab the latest release tag. There's a function in `sunbeam.sh` for that ... `git-get-latest-release-tag`

Then an explicit `go get` on the correct version of dylt

eg
```
$ go get github.com/dylt-dev/dylt@v0.0.7-nightly.20240605184642
```

Then `diggo` will be using the latest `dylt`

#### Caveat
I could not get the above to work consistently. Sometimes I had to edit `go.mod` by hand
```
require (
	github.com/dylt-dev/dylt v0.0.7-nightly.20240605201237
    ...
}
```

Hopefully there's a way to get rid of this.