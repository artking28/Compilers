
all:
	clear; echo "Please, use some Makefile command!"

mantis:
	cd src; go build -o ../bin/mantis ./cmd/
