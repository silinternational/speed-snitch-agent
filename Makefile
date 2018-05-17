dep:
	docker-compose run go dep ensure

builddist:
	docker-compose run go ./build-sign-deploy.sh

mac:
	docker-compose run go bash -c "cd cmd/speedsnitch && GOOS=darwin GOARCH=amd64 go build"

linux:
	docker-compose run go bash -c "cd cmd/speedsnitch && GOOS=linux GOARCH=amd64 go build"

windows:
	docker-compose run go bash -c "cd cmd/speedsnitch && GOOS=windows GOARCH=amd64 go build"