
all:
	clear; echo "Please, use some Makefile command!"

mantis:
	clear; go build -o bin/mantis src/langs/Mantis/cmd/main.go

uasm:
	clear; go build -o bin/uasm src/langs/UASM/cmd/main.go

test:
	clear; go test -v src/langs/Mantis/cmd/main_test.go; go test -v src/langs/UASM/cmd/main_test.go;
