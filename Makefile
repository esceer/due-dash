build-amd:
	docker build -f Dockerfile.amd -t due-dash .

build-arm:
	docker build -f Dockerfile.arm --platform=linux/arm64 -t due-dash .

build: build-arm

run: build
	docker run --rm -it -p 443:443 -v ~/Documents/due-dash-data:/app/database --name due-dash due-dash
