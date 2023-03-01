package Backend

import (
	"fmt"
	"github.com/cockroachdb/errors"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"strconv"
)

type Backend struct {
	router  http.Handler
	apiPort int
}

func (ca *Backend) catchAll(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusForbidden)
}

func (ca *Backend) Start() error {
	return http.ListenAndServe(fmt.Sprintf(":%d", ca.apiPort), ca.router)
}

func NewBackend() (*Backend, error) {
	// check port (mandatory)
	portNumEnv, eB := os.LookupEnv("BACKEND_PORT")
	if !eB {
		return nil, errors.Errorf("API port is missing")
	}

	portNum, e := strconv.Atoi(portNumEnv)
	if e != nil {
		return nil, errors.Errorf("API port not valid")
	}

	// initialize router
	router := mux.NewRouter()

	// initialize object
	backend := &Backend{router: router, apiPort: portNum}

	// bind router with functions
	router.HandleFunc("/clock", backend.clock).Methods(http.MethodGet)

	router.PathPrefix("/").HandlerFunc(backend.catchAll).Methods(http.MethodGet)

	return backend, nil
}

func StartBackend() error {
	api, err := NewBackend()
	if err != nil {
		return err
	}

	return api.Start()
}
