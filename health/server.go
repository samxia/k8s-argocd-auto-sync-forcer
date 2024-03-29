/**
 */
package health

import (
	"net/http"

	"github.com/samxyg/k8s-argocd-sync-forcer/k8s"
	"github.com/samxyg/k8s-argocd-sync-forcer/logger"
)

func Start() {
	// set liveness route
	http.HandleFunc("/live", func(w http.ResponseWriter, r *http.Request) {
		// get the watching event status
		if k8s.GetwatchingEventsStatus() {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusBadGateway)
			logger.Warn("Live checking failed")
		}
		logger.Debug("Liveness check")
	})

	// set readiness route
	http.HandleFunc("/ready", func(w http.ResponseWriter, r *http.Request) {
		// get the watching event status
		if k8s.GetwatchingEventsStatus() {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusBadGateway)
			logger.Warn("Ready checking failed")
		}
		logger.Debug("Readiness check")
	})

	// start HTTP server
	logger.Info("Starting health checking server...")
	logger.Fatal(http.ListenAndServe(":8080", nil))
}
