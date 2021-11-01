OS=""
prefix="gin_rbac"

.PHONY: build
build:
	export GOPROXY="https://goproxy.io,direct"
	mkdir -p ./bin && rm -r ./bin
	mkdir -p ./bin/configs && cp -r configs ./bin
	mkdir -p ./bin/web/admin && cp -r web ./bin/admin
	@if [ ${OS} != "" ]; then\
		GOOS=${OS} go build -o bin/admin/${prefix}_admin cmd/admin/main.go;\
	else\
		go build -o bin/admin/${prefix}_admin cmd/admin/main.go;\
    fi

clean:
	rm -rf ./bin

# swag 1.7.0
.PHONY: docs
docs:
	swag init -d ./cmd/admin --parseDependency --parseDepth 5 -g ./main.go -o ./docs