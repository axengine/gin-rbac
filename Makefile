SET_OS=""
prefix="gin-rbac"

.PHONY: build
build:
	export GOPROXY="https://goproxy.io,direct"
	mkdir -p ./bin && rm -r ./bin
	mkdir -p ./bin/configs && cp -r configs ./bin
	@if [ ${SET_OS} != "" ]; then\
		GOOS=${SET_OS} go build -o bin/admin/${prefix}_admin cmd/admin/main.go;\
	else\
		go build -o bin/admin/${prefix}_admin cmd/admin/main.go;\
    fi

clean:
	rm -rf ./bin

# swag 1.7.0
.PHONY: docs
docs:
	swag init -d ./cmd/admin --parseDependency --parseDepth 5 -g ./main.go -o ./docs