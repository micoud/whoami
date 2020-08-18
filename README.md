# whoami

For test purposes only!

Starts two webservers on two ports (8080,8081).

Derived from https://github.com/containous/whoami.

Two http endpoints are opened:

* `localhost:<port1>`
* `localhost:<port2>`

## Build, run

Build (yields `./whoami`)

```bash
go build
```

Run

```bash
./whoami [-port1 <port1>] [-port2 <port2>]
```

## Docker

A docker image can be built with `docker build [-t <registry>/whoami:latest] .`

Assuming that you run traefik as a load balancer for docker (e.g. a docker swarm) the container can be deployed via the docker compose file `docker-compose.yml`.
