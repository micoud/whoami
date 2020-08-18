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
