package pck

import (
	"bitbucket.org/tekion/tbaas/apiContext"
	log "bitbucket.org/tekion/tbaas/log/v1"
	api "bitbucket.org/tekion/tbaas/tapi"
	"net/http"
)

const servicePort = "8082"

// Start - start the service console app
func Start() {

	api.AddNoAuthServiceRoute(
		"get dealer by _id",
		http.MethodGet,
		"/dealer",
		getReport,
	)

	log.GenericInfo(apiContext.TContext{}, "Started TAP backend Service", log.FieldsMap{"port": servicePort})
	api.Start(servicePort, "/test")
}
