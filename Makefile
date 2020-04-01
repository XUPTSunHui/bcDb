.PHONY: envup envdr build runsv runcli

####UP
envup:
	@echo "=======env up...======="
	@docker-compose  -f docker-compose-mysql.yaml up -d
	@echo "env up done"

####DC
envdr:
	@echo "=======env down & remove...========="
	@docker-compose -f docker-compose-mysql.yaml down --volumes --remove-orphans
	@docker rm -f `docker ps -a -q --no-trunc`
	@echo "env down & remove done"

####BUILD
build:
	@echo "========build server & client...======"
	@cd server && go build dbServer.go
	@cd client && go build
	@echo "build server & client done"

####RUNSv
runsv:
	@echo "======run server...========"
	@gnome-terminal
	@cd server && ./dbServer
####RUNCli
runcli:
	@echo "======run client...========"
	@cd client && ./client
