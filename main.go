package main

import (
	"github.com/spachava753/go-kit-clean-arch/internal/domain/stream"
	"github.com/spachava753/go-kit-clean-arch/internal/domain/stream/repository"
	"github.com/spachava753/go-kit-clean-arch/internal/pkg"
	"github.com/spachava753/go-kit-clean-arch/internal/pkg/transport"
	"log"
)

var Version = "unknown"
var BuildDate = "unknown"

// main is the entry point of the application.
func main() {
	// print info
	log.Println("VERSION:", Version)
	log.Println("BuildDate:", BuildDate)

	// construct your application configuration
	log.Println("loading configuration")
	c, configErr := pkg.MakeAppConfig()
	if configErr != nil {
		log.Fatalf("misconfiguration: %s", configErr)
	}

	// construct your dependencies here, in order that they are required
	streamRepo := repository.MakeHardCodedStreamRepository()
	streamSvc := stream.MakeStreamService(streamRepo)

	log.Println("constructing server")
	server := transport.MakeServer(c, streamSvc)
	if err := server(); err != nil {
		log.Fatalf("server error: %s", err)
	}
}
