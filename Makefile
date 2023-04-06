

REPO_FSPATH=$(PWD)
BIN_FSPATH=$(REPO_FSPATH)/cmd/task
# adjust for cross manually
BIN_NAME=$(BIN_FSPATH)/task-dash
BIN=$(BIN_FSPATH)/$(BIN_NAME)

mod-up:
	cd $(REPO_FSPATH) && go-mod-upgrade

mod-tidy:
	cd $(REPO_FSPATH) && go mod tidy

build:
	cd $(BIN_FSPATH) && go build -o $(BIN_NAME) .
build-cross:
	# todo LDFLags and version...
	cd $(BIN_FSPATH) && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(BIN_NAME)-linux-amd64 .
	cd $(BIN_FSPATH) && CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o $(BIN_NAME)-darwin-amd64 .
	cd $(BIN_FSPATH) && CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o $(BIN_NAME)-windows-amd64.exe .

TEST_FSPATH=$(REPO_FSPATH)/testdata/short_task_notation/Taskfile.yml

TEST_FSPATH=$(REPO_FSPATH)/Taskfile.yml

run-print:
	@echo $(TEST_FSPATH)
run-h:
	$(BIN) -h 
run:
	$(BIN) --dry --menu --taskfile $(TEST_FSPATH)