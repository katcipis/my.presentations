all: deps present

present:
	present .

deps:
	go get golang.org/x/tools/cmd/present
