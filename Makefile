docker-build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o cliphub .
	docker build -t cliphub-scratch .