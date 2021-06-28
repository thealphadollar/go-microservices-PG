# create network layer
resource "google_compute_network" "micro-pg-network" {
  name = "${var.platform-name}"
}

resource "google_compute_firewall" "ssh" {
  name = "${var.platform-name}-ssh"
  network = "${google_compute_network.micro-pg-network.name}"
  allow {
      protocol = "icmp"
  }
  allow {
      protocol = "tcp"
      ports = ["22", "80", "443"]
  }

  source_ranges = ["0.0.0.0/0"]
}

resource "google_dns_managed_zone" "micro-pg-dns" {
  name = "micro-pg-com"
  dns_name = "micro-pg.com"
  description = "micro PG dns zone"
}

resource "google_compute_subnetwork" "micro-pg" {
  name = "dev-${var.platform-name}-${var.gcloud-region}"
  ip_cidr_range = "10.1.2.0/24"
  network = "${google_compute_network.micro-pg-network.self_link}"
  region = "${var.gcloud-region}"
}

resource "google_container_cluster" "micro-pg-cluster" {
  name = "micro-pg-cluster"
  network = "${google_compute_network.micro-pg-network.name}"
  subnetwork = "${google_compute_subnetwork.micro-pg.name}"
  # zone = "${var.gcloud-zone}"

  initial_node_count = 1

  master_auth {
    username = "thealphadollar"
    password = "thealphadollar"
  }

  node_config {
    machine_type = "n1-standard-1"
    oauth_scopes = [
      "https://www.googleapis.com/auth/projecthosting",
      "https://www.googleapis.com/auth/devstorage.full_control",
      "https://www.googleapis.com/auth/monitoring",
      "https://www.googleapis.com/auth/logging.write",
      "https://www.googleapis.com/auth/compute",
      "https://www.googleapis.com/auth/cloud-platform",
    ]
  }
}

resource "google_dns_record_set" "dev-k8s-endpoint-micro-pg" {
  name = "k8s.dev.${google_dns_managed_zone.micro-pg-dns.dns_name}"
  type = "A"
  ttl = 300
  managed_zone = "${google_dns_managed_zone.micro-pg-dns.name}"
  rrdatas = ["${google_container_cluster.micro-pg-cluster.endpoint}"]
}