# buildbot-dashboard [![Join the chat at https://gitter.im/ghophp/buildbot-dashboard](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/ghophp/buildbot-dashboard?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)
If you have a CI/CD setup, it is really important to keep a central display and follow the builds and process going on the CI tools. This project aims to be the dashboard that will allow you to do present better results at a central display with `buildbot`.

## Download [![Build Status](https://semaphoreci.com/api/v1/projects/44130239-880c-468f-9fa7-b976a355676a/611030/badge.svg)](https://semaphoreci.com/ghophp/buildbot-dashboard)
The project has some precompiled binaries, if your enviroment match one of the releases above, this is the most simple way to use this project.

| OS                                                                                                                | Arch   | Size | SHA1 Binary                              |
| ----------------------------------------------------------------------------------------------------------------- | ------ | ---- | ---------------------------------------- |
| [FreeBSD](https://github.com/ghophp/buildbot-dashboard/releases/download/0.1.0/buildbot-dashboard.freebsd.tar.gz) | 64-bit | 6,5MB | 2c45eb46dd63fe038f34031afb472ad2349c3202 |
| [OSX](https://github.com/ghophp/buildbot-dashboard/releases/download/0.1.0/buildbot-dashboard.osx.tar.gz)         | 64-bit | 6,5MB | 61492371b00982802dc13013817c6ff791f7a1df |
| [Linux](https://github.com/ghophp/buildbot-dashboard/releases/download/0.1.0/buildbot-dashboard.linux.tar.gz)     | 64-bit | 6,5MB | a2e3c6fc412a80034b50504e5df0e9ed3cb74126 |

## Running
`buildbot` is the only required flag, you must provide the base url of the running builbot.
```sh
$ ./buildbot_dashboard -h
-buildbot string
	buildbot url eg. http://10.0.0.1/
-filter string
	regex applied over the builder name
-invalidate int
	cache invalidate in seconds (default and min 5 minutes) (default 10)
-refresh int
	refresh rate in seconds (default and min 10 seconds) (default 10)
```

How to use with [runit](https://github.com/ghophp/buildbot-dashboard/wiki/runit).

## Manual Build
As this project is built in `go` you can build it in multiple platforms with:
```sh
$ go get -d github.com/ghophp/buildbot-dashboard
$ cd $GOPATH/src/github.com/ghophp/buildbot-dashboard
$ make
```
This will fetch the code and generate the assets as binary, then you can `go install` or `go build`, and then run:
```sh
$ buildbot-dashboard --buildbot="http://10.0.0.1/"
```
If you have `$GOPATH/bin` at your `$PATH` then you can run from everywhere the command.

You can also run with the `go run` with:
```sh
$ go run *.go -builbot="http://10.0.0.1/"
```

## Features
- Non-reload monitoring through `websockets`
- Enhanced UI with better visualization of the builders
- Easy usage with single command (`binary assets`)
- `Filter` options allow you to just show what matters
- `Base64` communication with websockets (light)
- Save arrangement of the dashboard
- `Zoom compliant` layout

## Preview
![Apache Board](/preview/preview_apache.gif?raw=true "Apache Board")

As we use gridster and a structured style, if your buildbot has a lot of builders, you can take advantage of the `browser zoom` to reduce the size of the blocks and so, use more space to move your grid.

![Apache Board Small](/preview/preview_apache_small.gif?raw=true "Apache Board Small")

## Martini
As this project is build over `martini` please consider setting this `env variables` when deploy:
```sh
export PORT=3000
export MARTINI_ENV=production
```

## Clean Cache
By default the project will cache the builders to avoid the delay of reloading it everytime. If you insert a new builder or remove one, you can force the reload of the cache in the UI, there is [a button on left side](https://github.com/ghophp/buildbot-dashboard/wiki/). This button will force the builders to be reloaded and the localStorage will also be cleaned.

## File limits
If you have a lot of builders on your buildbot and a filesystem with small limit of files, you may run into `too many open files` problem. As we want to keep the system as fast as possible on feedback a building process, we spam the routines to fetch the information from each builder all at the same time. The dashboard will not crash, `but it will be delays on presenting the state`. Please keep a look at logs for this problem, and if start to happen it is recommended to increase the file limit on your system (`ulimit`).
