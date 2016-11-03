# GKE - Golang Tutorial

## Prerequisite

### Platforms

- Google Cloud Platform
  - Container Registry

### Tools

- docker
- gcloud
- [kubectl](http://kubernetes.io/docs/user-guide/kubectl-cheatsheet/)


## Steps

### 1. Create Dockerfile and go code

Create a Dockerfile.

```docker
FROM golang:1.7.3-onbuild
```

```go
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Container!")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

### 2. Build and Run

Build a new image from Dockerfile.

```shell
# Build
docker build -t [docker-tag] .
# e.g.) docker build -t hello-golang .

# Run to make sure that the image has no problem.
docker run -p 8080:8080 [docker-tag]
# e.g.) docker run -p 8080:8080 hello-golang
```

### 3. Tag

Assign a new alias to the image.

```shell
# Tag
docker tag [docker-tag] [registry-host]/[project-id]/[docker-tag]
# e.g.) docker tag us.gcr.io/xxxxxxxxx/hello-golang
```

| Registry Host | Region |
| :------------ | :----- |
| us.gcr.io     | US     |
| es.gcr.io     | EU     |
| asia.gcr.io   | ASIA   |
| gcr.io        | US (unfixed) |

### 4. Push to Container Registry

```shell
gcloud docker push [registry-host]/[project-id]/[docker-tag]
# e.g.) gcloud docker push us.gcr.io/xxxxxxxxx/hello-golang
```

### 5. Deployment

Create Pod or Replication Controllers on Cloud Shell.

#### Login Cloud Shell

Create Container Cluster on Cloud Dashboard.

And then login Cloud Shell to use `kubectl`.

```shell
gcloud container clusters get-credentials [cluster-name] \
  --zone [cluster-zone] --project [project-id]
```

#### Create Pod or Replication Controllers

```shell
# Create Pod
kubectl create -f pod.yaml
kubectl get pod

# Create Replication Controllers
kubectl create -f replication-controller.yaml
kubectl get rc
```

#### Create Service

```shell
kubectl create -f service.yaml
kubectl get service
# NAME         CLUSTER-IP     EXTERNAL-IP      PORT(S)   AGE
# golang       x.x.x.x        x.x.x.x          80/TCP    7m
# kubernetes   x.x.x.x        <none>           443/TCP   26m
```

### Deploy to GCE

```shell
gcloud compute instances create [instance-name] \
  --zone us-central1-b \
  --image-family container-vm \
  --image-project google-containers \
  --machine-type n1-standard-1 \
  --tags "http-server" \
  --metadata-from-file google-container-manifest=pod.yaml
```

## License

[The MIT License (MIT)](http://kaneshin.mit-license.org/)

## Author

[Shintaro Kaneko](https://github.com/kaneshin) <kaneshin0120@gmail.com>
