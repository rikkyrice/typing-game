package handler

type AppHandler interface {
	HealthCheckHandler
	UserHandler
}
