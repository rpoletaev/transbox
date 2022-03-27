gen:
	gen-server
	
.PHONY: gen-server

gen-server:
	swagger generate server -f ./ports/http/swagger.yml -t ./ports/http/gen -A entities-service