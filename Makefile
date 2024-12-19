.PHONY: test
test:
	go test -v ./...

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: setup-asdf # install binaries from .tool-versions
setup-asdf:
	cut -d\  -f1 .tool-versions|grep -E '^[^#]'|xargs -L1 asdf plugin add
	asdf install

.PHONY: setup-pre-commit # install pre-commit hooks
setup-pre-commit:
	pre-commit install --install-hook
	pre-commit install --hook-type commit-msg

.PHONY: setup # setup dev environment
setup: setup-asdf setup-pre-commit

.PHONY: container-run # run tests in a container
container-run:
	docker build -t code-challenges-go .
	docker run --rm code-challenges-go
