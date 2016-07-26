# tugbot-result-service-es

Tugbot Result Service ES

Implements [Result Service API](https://github.com/gaia-docker/tugbot/blob/master/doc/proposal/Result%20Service%20API.md#api-design) 
and publish test results to elastic search

## Usage
```
$ tugbot-result-service-es help

NAME:
   tugbot-result-service-es - Implements Result Service API and publish test results to elastic search

USAGE:
   tugbot-result-service-es [global options] command [command options] [arguments...]

VERSION:
   1.0.0

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --port value, -p value   http service port (default: "8081")
   --loglevel value, -l value  log level (default: "debug")
   --elasticurl value, -e value elastic search URL (default: "http://127.0.0.1:9200")
   --help, -h                  show help
   --version, -v               print the version
```

## Run as docker container
`docker run -it --name result-service-es -p 8081:8081 gaiadocker/tugbot-result-service-es ./tugbot-result-service-es -e http://esUrl:9200`

