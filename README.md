# docker-compose-promlab

> This repo forked from [dockprom](https://github.com/stefanprodan/dockprom)

This project consists of the commonly used components.

Quick start all components in one step.

## Start

```shell
ADMIN_USER='admin' ADMIN_PASSWORD='admin' ADMIN_PASSWORD_HASH='$2a$14$1l.IozJx7xQRVmlkEQ32OeEEfP5mRxTpbDTCTcXRqn19gXD8YK1pO' docker-compose up -d
```

## Consul

> Add consul for service discovery.

- Register service `go-gin-demo1`

```shell
curl -v  --request PUT \
 --header 'Authorization: Basic YWRtaW46YWRtaW4=' \
 --data '{
  "ID": "go-gin-demo1",
  "Name": "go-gin-demo1",
  "Tags": ["gin","golang"],
  "Address": "http_service1",
  "Port": 5001
}' http://localhost:8500/v1/agent/service/register

```

- Deregister service  `go-gin-demo1`

``` shell
curl \
    --header 'Authorization: Basic YWRtaW46YWRtaW4=' \
    --request PUT \
    http://127.0.0.1:8500/v1/agent/service/deregister/go-gin-demo1

```

### Consul Service Discovery configuration in prometheus.yml

```yaml
  scrape_configs:
  - job_name: 'consul'
    consul_sd_configs:
    - server: 'consul:8500'
      services: []
```
