# tugbot-result-service-es

[![CircleCI](https://circleci.com/gh/gaia-docker/tugbot-result-service-es.svg?style=svg)](https://circleci.com/gh/gaia-docker/tugbot-result-service-es)
[![codecov](https://codecov.io/gh/gaia-docker/tugbot-result-service-es/branch/master/graph/badge.svg)](https://codecov.io/gh/gaia-docker/tugbot-result-service-es)
[![Go Report Card](https://goreportcard.com/badge/github.com/gaia-docker/tugbot-result-service-es)](https://goreportcard.com/report/github.com/gaia-docker/tugbot-result-service-es)
[![](https://imagelayers.io/badge/gaiadocker/tugbot-result-service-es:latest.svg)](https://imagelayers.io/?images=gaiadocker/tugbot-result-service-es:latest 'Get your own badge on imagelayers.io')

Tugbot Result Service ES

Implements [Result Service API](https://github.com/gaia-docker/tugbot/blob/master/doc/proposal/Result%20Service%20API.md#api-design) - publish test results & events into elasticsearch.

## Usage
```
$ tugbot-result-service-es help

NAME:
   tugbot-result-service-es - Publish test results to elasticsearch.

USAGE:
   tugbot-result-service-es [global options] command [command options] [arguments...]

VERSION:
   1.0.0

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --port value, -p value        http service port (default: "8081")
   --debug                       enable debug mode with verbose logging
   --elasticurl value, -e value  elastic search URL (default: "http://127.0.0.1:9200")
   --help, -h                    show help
   --version, -v                 print the version

```

## Run as docker container
`docker run -it --name result-service-es -p 8081:8081 --link es:esUrl gaiadocker/tugbot-result-service-es ./tugbot-result-service-es -e http://esUrl:9200`

