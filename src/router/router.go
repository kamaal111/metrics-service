package router

import (
	"log"
	"net/http"

	"github.com/kamaal111/metrics-service/src/models"
)

func HandleRequests(port string) {
	mux := http.NewServeMux()

	mux.Handle("/", loggerMiddleware(http.HandlerFunc(rootHandler)))
	// TODO: Deprecate this
	mux.Handle("/collect/", loggerMiddleware(withDeprecateEndpoint(models.VERSION_1_0_0, restrictToHttpMethod(http.MethodPost, http.HandlerFunc(metricsCollectHandler)))))
	mux.Handle("/metrics/data/", loggerMiddleware(restrictToHttpMethod(http.MethodGet, http.HandlerFunc(metricsDataHandler))))
	mux.Handle("/metrics/collect/", loggerMiddleware(restrictToHttpMethod(http.MethodPost, http.HandlerFunc(metricsCollectHandler))))
	mux.Handle("/metrics/register/", loggerMiddleware(restrictToHttpMethod(http.MethodPost, apiKeyRequired(http.HandlerFunc(metricsRegisterHandler)))))

	log.Printf("Listening on %s\n", port)
	err := http.ListenAndServe(port, mux)
	log.Fatal(err)
}
