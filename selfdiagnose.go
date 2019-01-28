package main

import (
	"net/http"
	"os"

	selfdiagnose "github.com/emicklei/go-selfdiagnose"
	"github.com/emicklei/go-selfdiagnose/task"
)

func addSelfdiagnose() {
	// add http handlers for /internal/selfdiagnose.(html|json|xml)
	selfdiagnose.AddInternalHandlers()
	selfdiagnose.Register(task.ReportHttpRequest{})
	selfdiagnose.Register(task.ReportHostname{})
	selfdiagnose.Register(task.ReportCPU())

	if len(os.Getenv("GOOGLE_CLOUD_PROJECT")) > 0 {
		// https://cloud.google.com/appengine/docs/standard/nodejs/runtime
		m := map[string]interface{}{}
		for _, each := range []string{"GAE_APPLICATION", "GAE_DEPLOYMENT_ID", "GAE_ENV", "GAE_INSTANCE", "GAE_MEMORY_MB", "GAE_RUNTIME",
			"GAE_SERVICE", "GAE_VERSION", "GOOGLE_CLOUD_PROJECT", "PORT"} {
			m[each] = os.Getenv(each)
		}
		selfdiagnose.Register(task.ReportVariables{VariableMap: m, Description: "Google AppEngine Environment"})
	}
	// start a HTTP server
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	webServer := &http.Server{Addr: ":" + port}
	logInfo("HTTP api is listening on %s", port)
	logInfo("open http://localhost:%s/internal/selfdiagnose.html", port)
	webServer.ListenAndServe()
}
