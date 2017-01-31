# buildbot-dashboard [![Join the chat at https://gitter.im/ghophp/buildbot-dashboard](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/ghophp/buildbot-dashboard?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)
If you have a CI/CD setup, it is really important to keep a central display and follow the builds and process going on the CI tools. This project aims to be the dashboard that will allow you to do present better results at a central display with `buildbot`.

## Download [![Build Status](https://travis-ci.org/ghophp/buildbot-dashboard.svg?branch=master)](https://travis-ci.org/ghophp/buildbot-dashboard) [![Coverage Status](https://coveralls.io/repos/github/ghophp/buildbot-dashboard/badge.svg?branch=master)](https://coveralls.io/github/ghophp/buildbot-dashboard?branch=master)
The project has some precompiled binaries, if your enviroment match one of the releases above, this is the most simple way to use this project.

| OS                                                                                                                | Arch   | Size | SHA1 Binary                              |
| ----------------------------------------------------------------------------------------------------------------- | ------ | ---- | ---------------------------------------- |
| [OSX](https://s3-eu-west-1.amazonaws.com/buildbot-dashboard/master/osx/buildbot-dashboard)         				| 64-bit | 6,5MB | 2df4875acc7127bca5f3a74ad556b641d668eeb5 |
| [Linux](https://s3-eu-west-1.amazonaws.com/buildbot-dashboard/master/linux/buildbot-dashboard)     				| 64-bit | 6,5MB | b8990466151cddc6097e5af84f15f420b366aa6e |

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
	refresh rate in seconds (default and min 20 seconds) (default 20)
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
- Non-reload monitoring
- Enhanced UI with better visualization of the builders
- Easy usage with single command
- `Filter` options allow you to just show what matters
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

---

Thanks to [nbari](https://github.com/nbari) and [paw](https://luckymarmot.com/paw) for helping and supporting open source.

![](/preview/paw.png)
