generate:
	@echo " --- Generate Schema --- "
	go get github.com/99designs/gqlgen@v0.17.24 && go run github.com/99designs/gqlgen
	@echo " --- Generate Done --- "

run:
	go run server.go