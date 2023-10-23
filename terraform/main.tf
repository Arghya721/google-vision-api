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

resource "google_cloud_run_v2_service" "google-vision" {
    name     = "vision"
    location = var.region

    template {
        containers {
            image = "${var.artifact_registry_url}/google-vision:latest"
            ports {
                container_port = 1323
            }
        }
    }

    traffic {
        type = "TRAFFIC_TARGET_ALLOCATION_TYPE_LATEST"
        percent = 100
  }
}

data "google_iam_policy" "noauth" {
    binding {
        role = "roles/run.invoker"
        members = [
            "allUsers",
        ]
    }
}

resource "google_cloud_run_service_iam_policy" "noauth" {
    location = google_cloud_run_v2_service.google-vision.location
    project  = google_cloud_run_v2_service.google-vision.project
    service  = google_cloud_run_v2_service.google-vision.name

    policy_data = data.google_iam_policy.noauth.policy_data
}
