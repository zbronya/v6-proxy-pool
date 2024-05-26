package main

import (
	"fmt"
	"github.com/zbronya/v6-proxy/config"
	"github.com/zbronya/v6-proxy/proxy"
	"github.com/zbronya/v6-proxy/sysutils"
	"log"
	"net/http"
)

func main() {
	cfg := config.ParseFlags()
	if cfg.CIDR == "" {
		log.Fatal("cidr is required")
	}

	if cfg.AutoForwarding {
		sysutils.SetV6Forwarding()
	}

	if cfg.AutoRoute {
		sysutils.AddV6Route(cfg.CIDR)
	}

	if cfg.AutoIpNoLocalBind {
		sysutils.SetIpNonLocalBind()
	}

	p := proxy.NewProxyServer(cfg)

	fmt.Println("Starting server on  %s:%d", cfg.Bind, cfg.Port)
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", cfg.Bind, cfg.Port), p)

	if err != nil {
		log.Fatal(err)
	}

}
