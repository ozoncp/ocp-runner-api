.PHONY: setup
build: vendor-proto .generate .build

PHONY: .generate
.generate:
		IF NOT EXIST "pkg\ocp-runner-api" \
			mkdir "pkg\ocp-runner-api" &&\
			protoc -I vendor.protogen \
					--go_out=pkg/ocp-runner-api --go_opt=paths=import \
					--go-grpc_out=pkg/ocp-runner-api --go-grpc_opt=paths=import \
					--grpc-gateway_out=pkg/ocp-runner-api \
					--grpc-gateway_opt=logtostderr=true \
					--grpc-gateway_opt=paths=import \
					--validate_out lang=go:pkg/ocp-runner-api \
					--swagger_out=allow_merge=true,merge_file_name=api:. \
					api/ocp-runner-api/ocp-runner-api.proto &&\
			move "pkg\ocp-runner-api\github.com\ozoncp\ocp-runner-api\pkg\ocp-runner-api\*" "pkg\ocp-runner-api" &&\
			rmdir /s "pkg\ocp-runner-api\github.com" &&\
			IF NOT EXIST "cmd/ocp-runner-api" \
				mkdir "cmd/ocp-runner-api"

PHONY: .build
.build:
		go build -o bin/ocp-runner-api.exe cmd/ocp-runner-api/main.go

PHONY: install
install: build .install

PHONY: .install
install:
		go install cmd/grpc-server/main.go

PHONY: vendor-proto
vendor-proto: .vendor-proto

PHONY: .vendor-proto
.vendor-proto:
		IF NOT EXIST "vendor.protogen" \
			mkdir "vendor.protogen\api\ocp-runner-api" &&\
			copy "%cd%\api\ocp-runner-api\ocp-runner-api.proto" "%cd%\vendor.protogen\api\ocp-runner-api"\

		IF NOT EXIST "vendor.protogen\google" \
			git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
			mkdir "vendor.protogen\google" &&\
			move "vendor.protogen\googleapis\google\api" "vendor.protogen\google" &&\
			rmdir /s "vendor.protogen\googleapis" ;\

		IF NOT EXIST "vendor.protogen\github.com\envoyproxy" \
			mkdir "vendor.protogen\github.com\envoyproxy" &&\
			git clone https://github.com/envoyproxy/protoc-gen-validate vendor.protogen/github.com/envoyproxy/protoc-gen-validate