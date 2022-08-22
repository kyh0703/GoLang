module gitlab.com/ipron-cloud/call-app

go 1.19

replace (
	gitlab.com/ipron-cloud/call-app/api => /api
	gitlab.com/ipron-cloud/call-app/api/rpc => /api/rpc
	gitlab.com/ipron-cloud/call-app/app/controller => /controllers/v1
	gitlab.com/ipron-cloud/call-app/app/models => /models
	gitlab.com/ipron-cloud/call-app/pkg/config => /config
	gitlab.com/ipron-cloud/call-app/pkg/middleware => /middleware
)

require (
	github.com/arsmn/fiber-swagger/v2 v2.31.1
	github.com/caarlos0/env v3.5.0+incompatible
	github.com/gofiber/fiber/v2 v2.36.0
	github.com/gofiber/template v1.6.30
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	gitlab.com/ipron-cloud/grpc-idl v1.1.32
	google.golang.org/grpc v1.48.0
)

require (
	github.com/KyleBanks/depth v1.2.1 // indirect
	github.com/PuerkitoBio/purell v1.1.1 // indirect
	github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578 // indirect
	github.com/andybalholm/brotli v1.0.4 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/jsonreference v0.19.6 // indirect
	github.com/go-openapi/spec v0.20.4 // indirect
	github.com/go-openapi/swag v0.19.15 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/klauspost/compress v1.15.0 // indirect
	github.com/mailru/easyjson v0.7.6 // indirect
	github.com/swaggo/files v0.0.0-20210815190702-a29dd2bc99b2 // indirect
	github.com/swaggo/swag v1.8.1 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.38.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	golang.org/x/net v0.0.0-20220225172249-27dd8689420f // indirect
	golang.org/x/sys v0.0.0-20220227234510-4e6760a101f9 // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/tools v0.1.9 // indirect
	google.golang.org/genproto v0.0.0-20211208223120-3a66f561d7aa // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
