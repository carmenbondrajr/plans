build:
	@. cd src && go build -o ../bin/plans
test:
	@. cd src/test && go test -v