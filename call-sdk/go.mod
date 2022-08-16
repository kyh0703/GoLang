module gitlab.com/ipron-cloud/call-sdk

go 1.19

replace (
	gitlab.com/ipron-cloud/call-sdk/config => /config
	gitlab.com/ipron-cloud/call-sdk/controller => /controller/v1
	gitlab.com/ipron-cloud/call-sdk/rpc => /rpc
)

require (
	github.com/caarlos0/env v3.5.0+incompatible
	github.com/gofiber/fiber/v2 v2.36.0
	github.com/gofiber/template v1.6.30
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	gitlab.com/ipron-cloud/grpc-idl v1.1.27
	google.golang.org/grpc v1.48.0
)

require (
	github.com/andybalholm/brotli v1.0.4 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/klauspost/compress v1.15.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.38.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	golang.org/x/net v0.0.0-20220225172249-27dd8689420f // indirect
	golang.org/x/sys v0.0.0-20220227234510-4e6760a101f9 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20211208223120-3a66f561d7aa // indirect
	google.golang.org/protobuf v1.27.1 // indirect
)
