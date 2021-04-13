package service

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"golang-seed/pkg/middleware"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// Service stores the configuration of the service we are configuring.
type Service struct {
	name string

	enableRouting       bool
	routingHTTPServer   *http.Server
	routingRouterCalled bool
	routingHTTPRouter   *mux.Router

	enableProfiler bool
}

// Init the configuration of a new service for the current application with
// the provided name.
func Init(name string) *Service {
	return &Service{
		name: name,
	}
}

// ConfigureRouting enables a HTTP router.
func (service *Service) ConfigureRouting() {
	service.enableRouting = true
}

// RoutingRouter returns the router to register new HTTP routes on it.
func (service *Service) RoutingRouter() *mux.Router {
	if !service.enableRouting {
		panic("routing must be enabled to get a routing router")
	}

	if service.routingHTTPRouter == nil {
		service.routingHTTPRouter = mux.NewRouter()
	}

	service.routingHTTPRouter.Use(middleware.RecoverHandler)
	service.routingHTTPRouter.Use(middleware.HeaderHandler)

	service.routingRouterCalled = true

	return service.routingHTTPRouter
}

// Run starts listening in every configure port needed to provide the configured features.
func (service *Service) Run() {
	rand.Seed(time.Now().UTC().UnixNano())

	if service.enableRouting && !service.routingRouterCalled {
		panic("do not configure routing without routes")
	}

	var wg sync.WaitGroup

	if service.enableRouting {
		wg.Add(1)
		go func() {
			defer wg.Done()

			log.Info("routing server enabled")

			service.routingHTTPServer = &http.Server{
				Addr:    ":8080",
				Handler: service.routingHTTPRouter,
			}
			if err := service.routingHTTPServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatal(err)
			}
		}()
	}

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "%s is ok\n", service.name) })

	service.stopListener()

	log.WithField("name", service.name).Println("instance initialized successfully!")

	wg.Wait()
	os.Exit(0)
}

func (service *Service) stopListener() {
	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)

	go func() {
		sig := <-gracefulStop
		log.WithField("signal", sig).Info("caught OS signal")

		var wg sync.WaitGroup

		if service.enableRouting {
			wg.Add(1)
			go func() {
				defer wg.Done()

				ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
				defer cancel()

				if err := service.routingHTTPServer.Shutdown(ctx); err != nil {
					log.WithField("error", err).Error("cannot shutdown routing HTTP server")
				}
			}()
		}

		wg.Wait()
		os.Exit(0)
	}()
}
