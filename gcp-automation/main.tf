# 1. terraform configuration
terraform{
     required_providers {
        google = {
         source = "hashicorp/google"
         version = "7.15.0"
    }
  }
}

# 2. provider configuration
provider "google" {
    project = "gcloud-learn-483710"
    region = "us-central1"
    zone  = "us-central1-a"
}
# 3. provider's resource we need configuration
# 3.1 VPC network
resource "google_compute_network" "my-project-vpc-1"{
    name = "my-project-vpc-1"
    auto_create_subnetworks = false
}

# 3.2 Subnet
resource "google_compute_subnetwork" "my-project-vpc-1" {
    name = "my-project-vpc-1"
    ip_cidr_range = "10.128.0.0/20"
    region = "us-central1"
    network = google_compute_network.my-project-vpc-1.id 
}
# 3.3 Compute Enginer Instance
resource "google_compute_instance" "public-instance-1" {
    name = "public-instance-1"
    machine_type = "e2-micro"
    zone = "us-central1-a"

    boot_disk {
        initialize_params{
            image = "debian-cloud/debian-11"
        }
    }
    
    network_interface{
        network = google_compute_network.my-project-vpc-1.id
        subnetwork = google_compute_subnetwork.my-project-vpc-1.id

        access_config {
        #skip this block so that we will get the external ip    
        }
    }
}


