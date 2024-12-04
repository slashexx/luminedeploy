package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	// "lumine/integrations/cicd/githubactions"
	// "lumine/integrations/cicd/jenkins"
	// "lumine/integrations/docker"
	// "lumine/integrations/monitoring"
	// "lumine/integrations/providers"
	"lumine/backend/controllers"
	// "lumine/templates"
)

func RegisterRoutes(r *mux.Router) {

	r.HandleFunc("/upload", controllers.UploadHandler).Methods("POST")
	r.HandleFunc("/files", controllers.ListFilesHandler).Methods("GET")

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/generate-github-actions", controllers.GitHubActionHandler).Methods("POST")
	api.HandleFunc("/generate-jenkinsfile", controllers.GenerateJenkinsFile).Methods("POST")
	api.HandleFunc("/generate-go-dockerfile", controllers.GenerateGoDockerfile).Methods("POST")
	api.HandleFunc("/setup-prometheus-monitoring", controllers.PrometheusSetupHandler).Methods("POST")
	api.HandleFunc("/setup-ecr-config", controllers.GenerateECRConfig).Methods("POST")
	api.HandleFunc("/setup-eks-config", controllers.GenerateEKSConfig).Methods("POST")
	api.HandleFunc("/setup-sss-config", controllers.GenerateS3Config).Methods("POST")
	api.HandleFunc("/download-zip", controllers.HandleDownloadZip).Methods("GET")
	api.HandleFunc("/docker-deploy", controllers.HandleDockerDeploy).Methods("POST")


	// Health check route
	api.HandleFunc("/health", HealthCheck).Methods("GET")
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"status": "healthy"}
	json.NewEncoder(w).Encode(response)
}
