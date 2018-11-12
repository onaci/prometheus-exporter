
IMAGE="$(REGISTRY)/onaci/prometheus-exporter:master"

run:
	docker run \
		-p 8080:8080 \
		-v /var/run/docker.sock:/var/run/docker.sock \
		-it --rm \
			$(IMAGE)

build:
	docker build -t $(IMAGE) .