
## Package

- Module has packages (folders)

- From the root of the project:
```
go mod init entitlement
go get github.com/gin-gonic/gin
```
- Create the package "entitlementsvc"
- Import the package "entitlementssvc"
```
import (
	"entitlement/entitlementsvc"
    )
```

## References
- [Go Dev GIN Example](https://go.dev/doc/tutorial/web-service-gin)
- [Official REST using GIN Example](https://github.com/gin-gonic/examples)
- [Another GIN Example by Logrocket](https://blog.logrocket.com/rest-api-golang-gin-gorm/)
- [Module & Package - New Approach after Go 1.13](https://www.digitalocean.com/community/tutorials/how-to-use-go-modules)
- [OpenAPI Client for Go](https://github.com/OpenAPITools/openapi-generator)
- [Go Standard Project Layout](https://github.com/golang-standards/project-layout)