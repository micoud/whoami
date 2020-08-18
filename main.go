package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"time"
)

type portHandler struct {
	port string
}

var port1 string
var port2 string

func init() {
	flag.StringVar(&port1, "port1", "8080", "Port number 1")
	flag.StringVar(&port2, "port2", "8081", "Port number 2")
}

func main() {
	flag.Parse()

	handler1 := &portHandler{port: port1}
	handler2 := &portHandler{port: port2}

	whoami1 := http.NewServeMux()
	whoami1.HandleFunc("/", handler1.whoamiHandler)
	whoami2 := http.NewServeMux()
	whoami2.HandleFunc("/", handler2.whoamiHandler)

	go func() {
		log.Fatal(http.ListenAndServe(":"+port1, whoami1))
	}()
	log.Fatal(http.ListenAndServe(":"+port2, whoami2))
}

func (ph *portHandler) whoamiHandler(w http.ResponseWriter, req *http.Request) {
	u, _ := url.Parse(req.URL.String())
	wait := u.Query().Get("wait")
	if len(wait) > 0 {
		duration, err := time.ParseDuration(wait)
		if err == nil {
			time.Sleep(duration)
		}
	}

	_, _ = fmt.Fprintln(w, "Port called:", ph.port)

	hostname, _ := os.Hostname()
	_, _ = fmt.Fprintln(w, "Hostname:", hostname)

	ifaces, _ := net.Interfaces()
	for _, i := range ifaces {
		addrs, _ := i.Addrs()
		// handle err
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			_, _ = fmt.Fprintln(w, "IP:", ip)
		}
	}

	_, _ = fmt.Fprintln(w, "RemoteAddr:", req.RemoteAddr)
	if err := req.Write(w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
