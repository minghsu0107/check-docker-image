package main

import (
	"net/url"
	"os"
	"strings"

	"github.com/minghsu0107/check-docker-image/registry"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true, DisableColors: true})
}

func main() {
	var config Config
	setDefault(&config)
	if err := readEnv(&config); err != nil {
		log.Fatalf("error reading env: %v", err)
	}

	logLevel, _ := log.ParseLevel(config.LogLevel)
	log.SetLevel(logLevel)

	u, err := url.Parse(config.Registry.Url)
	if err != nil {
		log.Fatalf("error parsing registry URL: %v", err)
	}
	var hub *registry.Registry
	switch u.Scheme {
	case "https":
		hub, err = registry.New(config.Registry.Url, config.Registry.Username, config.Registry.Password)
	case "http":
		hub, err = registry.NewInsecure(config.Registry.Url, config.Registry.Username, config.Registry.Password)
	default:
		log.Fatalf("invalid URL scheme: %v", u.Scheme)
	}

	if err != nil {
		log.Fatalf("error connecting to hub: %v", err)
	}

	log.Infof("start searching images %s on %s", strings.Join(config.Images, ","), config.Registry.Url)
	if !checkImages(hub, config.Images) {
		os.Exit(1)
	}
	os.Exit(0)
}
