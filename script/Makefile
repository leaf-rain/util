GOPATH:=$(shell go env GOPATH)
MODIFY=Mproto/imports/api.proto=github.com/micro/go-micro/v2/api/proto

Port:=8888
Name:=app

.PHONY: build
build:
	if [ -d build ]; then echo "build exists"; else mkdir "build" ; fi
	@echo 正在生成可执行文件...
ifeq ($(findstring WIN,$(OS_NAME)),WIN)
	@echo 编译win软件中...
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -a -ldflags '-extldflags "-static"' -o $(BUILD_NAME) ./cmd/main/
else ifeq ($(findstring Darw,$(OS_NAME)),Darw)
#else ifeq ($(OS_NAME), Darwin)
	@echo 编译mac软件中...
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -a -ldflags '-extldflags "-static"' -o $(BUILD_NAME) ./cmd/main/ && chmod +x ./$(BUILD_NAME)
else ifeq ($(findstring Linux,$(OS_NAME)),Linux)
#else ifeq (,$(OS_NAME), Linux)
	@echo 编译linux软件中...
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-extldflags "-static"' -o $(BUILD_NAME) ./cmd/main/ && chmod +x ./$(BUILD_NAME)

.PHONY: docker
docker:
	make build
	@echo 正在打包镜像...
	docker build . -t $(Name):latest

.PHONY: run
run:
	@echo 正在执行...
localhost:
	make build
	@echo 正在删除运行中文件...
	bash ./build/kill.sh $(Port)
	@echo 正在运行二进制文件...
	nohup ./build/$(Name) &
	@echo 程序启动成功
nohup-localhost:
	make build
	@echo 正在删除运行中文件...
	bash ./build/kill.sh $(Port)
	@echo 正在运行二进制文件...
	nohup ./$(BUILD_NAME) -conf ./infrastructure/config/resource/app/config.yaml > output 2>&1 &
	@echo 程序启动成功
nohup-dev:
	make build
	@echo 正在删除运行中文件...
	bash ./build/kill.sh $(Port)
	@echo 正在运行二进制文件...
	nohup ./$(BUILD_NAME) -conf ./infrastructure/config/resource/app/config_dev.yaml > output 2>&1 &
	@echo 程序启动成功
docker:
	make docker
	@echo 正在运行docker文件
	docker run -itd -p $(Port):$(Port) --name $(Name) $(Name)

.PHONY: kill
kill:
	@echo 正在执行...
port:
	bash ./linux/killAppByPort.sh $(Port)
name:
	bash ./linux/killAppByName.sh $(Name)

