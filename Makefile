# Fetch git latest tag
LATEST_GIT_TAG:=$(shell git describe --tags $(git rev-list --tags --max-count=1))

dep:
	dep ensure -v
	mkdir -p vendor/github.com/tendermint vendor/github.com/ethereum
	git clone -b v0.11.0 --single-branch --depth 1 https://github.com/tendermint/iavl vendor/github.com/tendermint/iavl
	git clone -b v1.8.17 --single-branch --depth 1 https://github.com/ethereum/go-ethereum vendor/github.com/ethereum/go-ethereum

clean:
	rm -rf build

build: clean
	mkdir -p build
	go build -o build/heimdalld cmd/heimdalld/main.go
	go build -o build/heimdallcli cmd/heimdallcli/main.go
	go build -o build/bridge bridge/bridge.go

contracts:
	# mkdir -p contracts/validatorset
	# abigen --abi=contracts/validatorset/validatorset.abi --pkg=validatorset --out=contracts/validatorset/validatorset.go

init-heimdall:
	./build/heimdalld init

show-account-heimdall:
	./build/heimdalld show-account

run-heimdall:
	./build/heimdalld start

reset-heimdalld:
	./build/heimdalld unsafe-reset-all

rest-server:
	./build/heimdallcli rest-server

start:
	mkdir -p ./logs
	./build/heimdalld start > ./logs/heimdalld.log &
	./build/heimdallcli rest-server > ./logs/heimdallcli.log &

start-bridge:
	mkdir -p ./logs
	./build/bridge start > ./logs/bridge.log &

start-all: start-bridge start
	mkdir -p ./logs
	tail -f ./logs/heimdalld.log ./logs/heimdallcli.log ./logs/bridge.log

#
# docker commands
#

build-docker:
	@echo Fetching latest tag: $(LATEST_GIT_TAG)
	git checkout $(LATEST_GIT_TAG)
	docker build -t "maticnetwork/heimdall:$(LATEST_GIT_TAG)" -f docker/Dockerfile .

push-docker:
	@echo Pushing docker tag image: $(LATEST_GIT_TAG)
	docker push "maticnetwork/heimdall:$(LATEST_GIT_TAG)"

build-docker-develop:
	docker build -t "maticnetwork/heimdall:develop" -f docker/Dockerfile.develop .

.PHONY: dep build