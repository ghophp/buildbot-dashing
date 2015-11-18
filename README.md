# buildbot-dashing

[![Join the chat at https://gitter.im/ghophp/buildbot-dashboard](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/ghophp/buildbot-dashboard?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)
If you have a buildbot setup at your company, it is really ugly to put it into a central display and follow the builds. This project aims to be the interface that will allow you to do that.

## Precompiled Binaries
The project has some precompiled binaries at the release area, so please take a look if any of the releases match your environment, so you can just download the binary and run normally with:
```sh
$ ./buildbot-dashboard --buildbot="http://10.0.0.1/"
```

## Manual Build
As this project is built in `go` you can build it in multiple platforms with:
```sh
$ go get -u -v github.com/ghophp/buildbot-dashboard
$ cd $GOPATH/src/github.com/ghophp/buildbot-dashboard
$ go install
$ buildbot-dashboard --buildbot="http://10.0.0.1/"
```
This will fetch the version directly from github and install it on `$GOPATH/bin`, if you have `$GOPATH/bin` at your `$PATH` then you can run from everywhere the command.

## Interface
When the project run, you should be able to access your localhost at the port 3000 (configurable via `PORT` env variable) and see the following:
