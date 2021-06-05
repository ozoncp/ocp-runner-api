module github.com/ozoncp/ocp-runner-api

go 1.16

require (
	github.com/envoyproxy/protoc-gen-validate v0.6.1
	github.com/golang/mock v1.5.0
	github.com/google/uuid v1.2.0
	github.com/kr/pretty v0.1.0 // indirect
	github.com/onsi/ginkgo v1.16.3
	github.com/onsi/gomega v1.13.0
	github.com/pkg/math v0.0.0-20141027224758-f2ed9e40e245
	github.com/stretchr/testify v1.7.0
	golang.org/x/net v0.0.0-20210525063256-abc453219eb5 // indirect
	golang.org/x/sys v0.0.0-20210603125802-9665404d3644 // indirect
	google.golang.org/genproto v0.0.0-20210604141403-392c879c8b08
	google.golang.org/grpc v1.38.0
	google.golang.org/protobuf v1.26.0
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
)

replace github.com/ozoncp/ocp-runner-api/pkg/ocp-runner-api => ./pkg/ocp-runner-api
