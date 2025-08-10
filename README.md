
# Auth User Service


Authentication  service leverages OpenAPI for self-documenting API.

## OpenApi build commands


```
oapi-codegen -generate models,gorilla-server,client -package openapi -o ./generated/openapi/openapi.gen.go ./api/openapi/openapi.yaml
```

