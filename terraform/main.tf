variable "service_account_key" {
    default = jsondecode(${secrets.GCP_SERVICE_ACCOUNT_KEY_JSON})
}

variable "project_id" {
    default = "${secrets.PROJECT_ID}"
}

variable "artifact_registry_url" {
    default = "${secrets.GCP_ARTIFACT_REGISTORY_URL}"
}

variable "region" {
    default = "${secrets.REGION}"
}

provider "google" {
    credentials = var.service_account_key
    project     = var.project_id
    region      = var.region
}

data "google_project" "project" {}

resource "google_cloud_run_service" "default" {
    name = "google-vision" 
    location = var.region 

    template {
        spec {
            containers {
                image = "${var.artifact_registry_url}/google-vision:latest"
                ports {
                    container_port = 1323
                }
            }
        }
        traffic {
            percent         = 100
            latest_revision = true
        }
    }
}
