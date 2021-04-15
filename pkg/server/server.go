package server

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
	port int

	enableRouting       bool
	routingHTTPServer   *http.Server
	routingRouterCalled bool
	routingHTTPRouter   *mux.Router
}

// Init the configuration of a new service for the current application with
// the provided name.
func Init(name string, port int) *Service {
	return &Service{
		name: name,
		port: port,
	}
}

// ConfigureRouting enables a HTTP router.
func (s *Service) ConfigureRouting() {
	s.enableRouting = true
}

// RoutingRouter returns the router to register new HTTP routes on it.
func (s *Service) RoutingRouter() *mux.Router {
	if !s.enableRouting {
		panic("routing must be enabled to get a routing router")
	}

	if s.routingRouterCalled || s.routingHTTPRouter != nil {
		panic("routing router already called")
	}

	router := mux.NewRouter()
	if len(s.name) > 0 {
		router = router.PathPrefix(s.name).Subrouter()
	}

	s.routingHTTPRouter = router

	s.routingHTTPRouter.Use(middleware.RecoverHandler)
	s.routingHTTPRouter.Use(middleware.HeaderHandler)
	s.routingHTTPRouter.Use(mux.CORSMethodMiddleware(s.routingHTTPRouter))

	s.routingRouterCalled = true

	return s.routingHTTPRouter
}

// Run starts listening in every configure port needed to provide the configured features.
func (s *Service) Run() {
	rand.Seed(time.Now().UTC().UnixNano())

	if s.enableRouting && !s.routingRouterCalled {
		panic("do not configure routing without routes")
	}

	var wg sync.WaitGroup

	if s.enableRouting {
		wg.Add(1)
		go func() {
			defer wg.Done()

			log.Info("routing server enabled")

			s.routingHTTPServer = &http.Server{
				Addr:    fmt.Sprintf(":%d", s.port),
				Handler: s.routingHTTPRouter,
			}
			if err := s.routingHTTPServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatal(err)
			}
		}()
	}

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "%s is ok\n", s.name) })

	s.stopListener()

	log.WithField("name", s.name).Println("instance initialized successfully!")

	wg.Wait()
	os.Exit(0)
}

func (s *Service) stopListener() {
	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)

	go func() {
		sig := <-gracefulStop
		log.WithField("signal", sig).Info("caught OS signal")

		var wg sync.WaitGroup

		if s.enableRouting {
			wg.Add(1)
			go func() {
				defer wg.Done()

				ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
				defer cancel()

				if err := s.routingHTTPServer.Shutdown(ctx); err != nil {
					log.WithField("error", err).Error("cannot shutdown routing HTTP server")
				}
			}()
		}

		wg.Wait()
		os.Exit(0)
	}()
}
