package http

import (
	"fmt"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/opencloud-eu/opencloud/pkg/cors"
	opencloudmiddleware "github.com/opencloud-eu/opencloud/pkg/middleware"
	"github.com/opencloud-eu/opencloud/pkg/service/http"
	"github.com/opencloud-eu/opencloud/pkg/version"
	svc "github.com/opencloud-eu/opencloud/services/thumbnails/pkg/service/http/v0"
	"github.com/opencloud-eu/opencloud/services/thumbnails/pkg/thumbnail/storage"
	"go-micro.dev/v4"
)

// Server initializes the http service and server.
func Server(opts ...Option) (http.Service, error) {
	options := newOptions(opts...)

	service, err := http.NewService(
		http.TLSConfig(options.Config.HTTP.TLS),
		http.Logger(options.Logger),
		http.Name(options.Config.Service.Name),
		http.Version(version.GetString()),
		http.Namespace(options.Config.HTTP.Namespace),
		http.Address(options.Config.HTTP.Addr),
		http.Context(options.Context),
		http.TraceProvider(options.TraceProvider),
	)
	if err != nil {
		options.Logger.Error().
			Err(err).
			Msg("Error initializing http service")
		return http.Service{}, fmt.Errorf("could not initialize http service: %w", err)
	}

	handle := svc.NewService(
		svc.Logger(options.Logger),
		svc.Config(options.Config),
		svc.Middleware(
			middleware.RealIP,
			middleware.RequestID,
			opencloudmiddleware.Cors(
				cors.Logger(options.Logger),
				cors.AllowedOrigins(options.Config.HTTP.CORS.AllowedOrigins),
				cors.AllowedMethods(options.Config.HTTP.CORS.AllowedMethods),
				cors.AllowedHeaders(options.Config.HTTP.CORS.AllowedHeaders),
				cors.AllowCredentials(options.Config.HTTP.CORS.AllowCredentials),
			),
			opencloudmiddleware.Version(
				options.Config.Service.Name,
				version.GetString(),
			),
			opencloudmiddleware.Logger(options.Logger),
		),
		svc.ThumbnailStorage(
			storage.NewFileSystemStorage(
				options.Config.Thumbnail.FileSystemStorage,
				options.Logger,
			),
		),
	)

	{
		handle = svc.NewInstrument(handle, options.Metrics)
		handle = svc.NewLogging(handle, options.Logger)
	}

	if err := micro.RegisterHandler(service.Server(), handle); err != nil {
		return http.Service{}, err
	}

	return service, nil
}
