.PHONY: all
all: nova

GOPROXY ?= "https://proxy.golang.org,direct"
.PHONY: nova
nova:
	GIT_COMMIT=`git rev-list -1 HEAD 2>/dev/null || echo ""` && \
	GIT_DATE=`git log -1 --date=short --pretty=format:%ct 2>/dev/null || echo ""` && \
	GOPROXY=$(GOPROXY) \
	go build \
	    -ldflags "-s -w -X github.com/Nova-foundation/go-nova/cmd/nova/launcher.gitCommit=$${GIT_COMMIT} -X github.com/Nova-foundation/go-nova/cmd/nova/launcher.gitDate=$${GIT_DATE}" \
	    -o build/nova \
	    ./cmd/nova


TAG ?= "latest"
.PHONY: nova-image
nova-image:
	docker build \
    	    --network=host \
    	    -f ./docker/Dockerfile.nova -t "nova:$(TAG)" .

.PHONY: test
test:
	go test ./...

.PHONY: coverage
coverage:
	go test -coverprofile=cover.prof $$(go list ./... | grep -v '/gossip/contract/' | grep -v '/gossip/emitter/mock' | xargs)
	go tool cover -func cover.prof | grep -e "^total:"

.PHONY: fuzz
fuzz:
	CGO_ENABLED=1 \
	mkdir -p ./fuzzing && \
	go run github.com/dvyukov/go-fuzz/go-fuzz-build -o=./fuzzing/gossip-fuzz.zip ./gossip && \
	go run github.com/dvyukov/go-fuzz/go-fuzz -workdir=./fuzzing -bin=./fuzzing/gossip-fuzz.zip


.PHONY: clean
clean:
	rm -fr ./build/*
