module fiber

go 1.19

replace (
	fiber/api => /api
	fiber/api/rpc => /api/rpc
	fiber/controller => /controllers/v1
	fiber/models => /models
	fiber/pkg/config => /config
	fiber/pkg/middleware => /middleware
)

require (
	github.com/arsmn/fiber-swagger/v2 v2.31.1
	github.com/caarlos0/env v3.5.0+incompatible
	github.com/gofiber/fiber/v2 v2.37.0
	github.com/gofiber/template v1.6.30
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
	github.com/josharian/intern v1.0.0 // indirect
	github.com/klauspost/compress v1.15.0 // indirect
	github.com/mailru/easyjson v0.7.6 // indirect
	github.com/stretchr/testify v1.7.1 // indirect
	github.com/swaggo/files v0.0.0-20210815190702-a29dd2bc99b2 // indirect
	github.com/swaggo/swag v1.8.1 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.39.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	golang.org/x/net v0.0.0-20220225172249-27dd8689420f // indirect
	golang.org/x/sys v0.0.0-20220227234510-4e6760a101f9 // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/tools v0.1.9 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
