

.PHONY: runclient
runclient:
	go run main.go client --config=config/client.yaml

.PHONY: runserver
runserver:
	go run main.go server --config=config/server.yaml
