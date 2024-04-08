
.PHONY: run-grpc
# run grpc service
run-grpc:
	@bash run-grpc.sh


.PHONY: run-http
# run grpc service
run-http:
	@bash run-http.sh


.PHONY: run-grpc-nohup
# run service with nohup in local, if you want to stop the server, pass the parameter stop, e.g. make run-grpc-nohup CMD=stop
run-grpc-nohup:
	@bash run-grpc-nohup.sh $(CMD)


.PHONY: run-http-nohup
# run service with nohup in local, if you want to stop the server, pass the parameter stop, e.g. make run-http-nohup CMD=stop
run-http-nohup:
	@bash run-http-nohup.sh $(CMD)


.PHONY: clean
# remove binary and log file
clean:
	@rm -vrf grpc/helloworld/grpc_helloworld
	@rm -vrf grpc_helloworld.log
	@rm -vrf http/helloworld/http_helloworld
	@rm -vrf http_helloworld.log


# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
