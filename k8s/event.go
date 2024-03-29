package k8s

import (
	"context"
	"strings"

	"github.com/samxyg/k8s-argocd-sync-forcer/argo"
	"github.com/samxyg/k8s-argocd-sync-forcer/logger"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// to monitor if running or not
var watchingEventsStatus bool = false

func Watch() {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	watcher, err := clientset.CoreV1().Events("").Watch(context.TODO(), metav1.ListOptions{})
	if err != nil {
		watchingEventsStatus = false
		logger.Fatal(err)
	}
	defer watcher.Stop()

	logger.Info("Watching the event...")

	for event := range watcher.ResultChan() {

		watchingEventsStatus = true
		// change the event type to *v1.Event
		eventObj, ok := event.Object.(*v1.Event)
		if !ok {
			logger.Warn("Failed to convert event to *v1.Event")
			continue
		}

		// get the fields value from event object
		involvedObject := eventObj.InvolvedObject
		source := eventObj.Source
		message := eventObj.Message

		logger.Debug("involvedObject: ", involvedObject)
		logger.Debug("source: ", source)
		logger.Debug("source.componet: ", source.Component)
		logger.Debug("message: ", message)

		// if the event related with argocd, output
		if source.Component == "argocd-application-controller" {
			logger.Info("source:", source.Component)
			logger.Info("message:", message)
		}

		// check if should trigger the force sync
		if source.Component == "argocd-application-controller" &&
			strings.Contains(message, "Sync operation") &&
			strings.Contains(message, "failed") &&
			//strings.Contains(message, `"sidecar.istio.io/inject":"true"`) &&
			strings.Contains(message, "field is immutable") {
			// output target argcd application name
			logger.Info("binggo +++++++ app: ", involvedObject.Name)

			watchingEventsStatus = false

			// go to force sync
			argo.Forcesync(involvedObject.Name)

			watchingEventsStatus = true
		}
	}
}

// get the watching status for livenewss/readiness
func GetwatchingEventsStatus() bool {
	return watchingEventsStatus
}
