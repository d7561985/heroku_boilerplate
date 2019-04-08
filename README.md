# Golang heroku boilerplate

## Introduction

> Template for golang project for future deploy on heroku dyno `container stack` from repo deployed on Github

## Code Samples

> You've gotten their attention in the introduction, now show a few code examples. So they get a visualization and as a bonus, make them copy/paste friendly.

## Installation

#### create app
As we use github existent repo for handling our app we can create new app on UI and link it here too. About how to do it you can read [here](https://devcenter.heroku.com/articles/github-integration)

#### container stack
For heroku deployment required container stack: `heroku stack:set container`. 
Option in UI mode i didn't find for doing this.

#### environments

* PORT - http port from app witch will be exposed in https through heroku proxy
* HEROKU_APP_NAME - for enable this from console we should enable plugin: 
`heroku labs:enable runtime-dyno-metadata -a {{.APP_NAME}}` where `{{.APP_NAME}}` is applications name in heroku (optional). I use it for webhook live registration.