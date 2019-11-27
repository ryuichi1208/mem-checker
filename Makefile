GO_BIN 	:= $(shell which go)
GO_FLG	:= -a -v -x
OBJS 	:= main.go http.go
PROGRAM := mem-checker

.PHONY: info
info:
	@$(GO_BIN) version

.PHONY: build
build:
	@$(GO_BIN) build -o $(PROGRAM) $(GO_FLG) $(OBJS)

.PHONY: clean
clean:
	@rm -f $(PROGRAM)
