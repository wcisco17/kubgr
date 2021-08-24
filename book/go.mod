module github.com/wcisco17/kubgr/book

go 1.16

replace github.com/wcisco17/kubgr/client => ../client

replace github.com/wcisco17/kubgr/db => ../db

require (
	github.com/andybalholm/brotli v1.0.3 // indirect
	github.com/gofiber/fiber/v2 v2.17.0
	github.com/iancoleman/strcase v0.2.0 // indirect
	github.com/prisma/prisma-client-go v0.10.0 // indirect
	github.com/valyala/fasthttp v1.29.0 // indirect
	github.com/wcisco17/kubgr/client v0.0.0-00010101000000-000000000000 // indirect
	github.com/wcisco17/kubgr/db v0.0.0-00010101000000-000000000000
	golang.org/x/sys v0.0.0-20210823070655-63515b42dcdf // indirect
)
