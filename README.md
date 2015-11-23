# buildbot-dashboard [![Join the chat at https://gitter.im/ghophp/buildbot-dashboard](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/ghophp/buildbot-dashboard?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)
If you have a CI/CD setup, it is really important to keep a central display and follow the builds and process going on the CI tool. This project aims to be the dashboard that will allow you to do present better results at a central display with `buildbot`.

## Download [![Build Status](https://semaphoreci.com/api/v1/projects/44130239-880c-468f-9fa7-b976a355676a/611030/badge.svg)](https://semaphoreci.com/ghophp/buildbot-dashboard)
The project has some precompiled binaries, if your enviroment match one of the releases above, this is the most simple way to use this project.

- [FreeBSD](https://s3-eu-west-1.amazonaws.com/buildbot-dashboard/freebsd/buildbot-dashboard)
- [Linux x64](https://s3-eu-west-1.amazonaws.com/buildbot-dashboard/linux64/buildbot-dashboard)
- [OSX](https://s3-eu-west-1.amazonaws.com/buildbot-dashboard/mac/buildbot-dashboard)

How to use with [runit](https://github.com/ghophp/buildbot-dashboard/wiki/runit).

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
-size string
	generic ui size (small|large default large) (default "large")
```

## Port and Logs
As this project is build over `martini` please consider setting this `env variables` when deploy:
```sh
export PORT=3000
export MARTINI_ENV=production
```

## Manual Build
As this project is built in `go` you can build it in multiple platforms with:
```sh
$ go get -u github.com/ghophp/buildbot-dashboard
$ cd $GOPATH/src/github.com/ghophp/buildbot-dashboard
$ ./install.sh
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

## Preview
![Apache Board](/preview/preview_apache.gif?raw=true "Apache Board")

![Apache Board Small](/preview/preview_apache_small.gif?raw=true "Apache Board Small")

## runit

When using [runit](http://smarden.org/runit/) this schema can be used:

    /services
    `--buildbot-dashboard
       |--app (The buildbot-dashboard binary)
       |--env
       |  `--PORT
       |--run (run script)
       `--log
        |--main/
        `--run (run script for log)

Contents of env/PORT (the environment var):

    8080


Contents of run:
```sh
#!/bin/sh

# get both standar output/error
exec 2>&1

# run like user buildbot
exec chpst -u buildbot -e ./env ./app --buildbot="http://10.0.0.1:8010"
```


Contents of log/run:
```sh
#!/bin/sh
exec svlogd -tt ./main
```

## Todo
- Use violetear
- Allow a totally compressed size to boards with a lot of builders
- Create a fully adapted mode, that by the size of the display, and the number of the projects, keep "walking" through the multiple screens
- Use violetear
