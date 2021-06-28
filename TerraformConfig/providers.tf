provider "google" {
  credentials = "${file("creds.json")}"
  project = "${var.gcloud-project}"
  region = "${var.gcloud-region}"
}