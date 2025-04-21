
all:
	clear; echo "Please, use some Makefile command!"

mantis:
	cd src && clear; go build -o ../bin/mantis ./langs/Mantis/cmd

uasm:
	cd src && clear; go build -o ../bin/uasm ./langs/UASM/cmd

test:
	clear; go test -v src/langs/Mantis/cmd/main_test.go; go test -v src/langs/UASM/cmd/main_test.go;
