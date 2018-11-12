
IMAGE="onaci/prometheus-exporter:latest"

run:
	docker run \
		-p 8080:8080 \
		-v /var/run/docker.sock:/var/run/docker.sock \
		-it --rm \
			$(IMAGE)

build:
	docker build -t $(IMAGE) .

modules:
	go mod tidy
	go mod download