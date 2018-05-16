dep:
	docker-compose run go dep ensure

builddist:
	docker-compose run go bash -c "cd cmd/speedsnitch && gox -osarch="linux/arm linux/amd64 windows/amd64 darwin/amd64" -output=\"../../dist/{{.OS}}/{{.Arch}}/speedsnitch\""

##
#  Before this can work, you need to have a gpg key pair and create an armored version of it.
#    (Keep track of its password, since you will need to input it during the `make` run.)
# $ gpg --gen-key
# $ gpg --list-secret-keys --keyid-format LONG   # Use the id from this line in the following line
# $ gpg --armor --export 1234567890123456
##
signdist:
	gpg -a -o ./dist/linux/arm/speedsnitch.sig --detach-sig ./dist/linux/arm/speedsnitch
	gpg -a -o ./dist/linux/amd64/speedsnitch.sig --detach-sig ./dist/linux/amd64/speedsnitch
	gpg -a -o ./dist/darwin/amd64/speedsnitch.sig --detach-sig ./dist/darwin/amd64/speedsnitch
	gpg -a -o ./dist/windows/amd64/speedsnitch.exe.sig --detach-sig ./dist/windows/amd64/speedsnitch.exe

mac:
	docker-compose run go bash -c "cd cmd/speedsnitch && GOOS=darwin GOARCH=amd64 go build"

linux:
	docker-compose run go bash -c "cd cmd/speedsnitch && GOOS=linux GOARCH=amd64 go build"

windows:
	docker-compose run go bash -c "cd cmd/speedsnitch && GOOS=windows GOARCH=amd64 go build"