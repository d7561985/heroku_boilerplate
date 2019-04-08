# Golang heroku boilerplate

[![Build Status](https://travis-ci.org/d7561985/heroku_boilerplate.svg?branch=master)](https://travis-ci.org/d7561985/heroku_boilerplate)

## Introduction

> Template for golang project for future deploy on heroku dyno `container stack` from repo deployed on Github

## Code Samples

> You've gotten their attention in the introduction, now show a few code examples. So they get a visualization and as a bonus, make them copy/paste friendly.

## Installation

I'm assume what you has heroku cli installed properly on development machine.

#### create app
As we use github existent repo for handling our app we can create new app on UI and link it here too.
1. create app
1. remember your app name. I use {{.APP_NAME}} template further with this name. 
1. here we should switch app to `container stack` mode through CLI. Look at `container stack` chapter.
1. add buildpack
1. link github
1. do deploy

About how to do it you can read [here](https://devcenter.heroku.com/articles/github-integration)

Note:
 
You can get list of your started apps through cli 
```bash
$ heroku apps
```

#### container stack
For heroku deployment required container stack: 
```
$ heroku stack:set container --app={{.APP_NAME}}
``` 
Option in UI mode i didn't find for doing this.

#### environments
Heroku app should handle some env. variables:
* PORT - http port from app witch will be exposed in https through heroku proxy
* HEROKU_APP_NAME - actual slug of our app(optional). For enable this from console we should enable plugin: 
```bash
$ heroku labs:enable runtime-dyno-metadata -a {{.APP_NAME}}
``` 
where `{{.APP_NAME}}` is applications name in heroku. I use it for webhook live registration.