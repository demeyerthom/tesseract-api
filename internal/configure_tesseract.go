// This file is safe to edit. Once it exists it will not be overwritten

package internal

import (
	"crypto/tls"
	"github.com/demeyerthom/tesseract-api/internal/handlers"
	"github.com/demeyerthom/tesseract-api/pkg/tesseract"
	"net/http"

	"github.com/demeyerthom/tesseract-api/internal/operations"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
)

//go:generate swagger generate server --target ../../tesseract-api --name Tesseract --spec ../swagger.json --model-package ./internal/models --server-package ./internal --exclude-main

func configureFlags(api *operations.TesseractAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.TesseractAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	api.JSONConsumer = runtime.JSONConsumer()

	tesClient := tesseract.NewClient()

	api.MultipartformConsumer = runtime.DiscardConsumer

	api.JSONProducer = runtime.JSONProducer()

	api.ConvertHandler = handlers.NewConvertHandler(tesClient)

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
