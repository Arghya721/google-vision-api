variable "service_account_key_json" {}
variable "project_id" {}
variable "artifact_registry_url" {}
variable "region" {}

provider "google" {
    credentials = var.service_account_key_json
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
