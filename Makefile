all: echo

echo:
	@echo "Makefile"

image:
	./docker_trivy.sh