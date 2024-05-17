
ci-test:
	go test  ./...
go:
	docker run --rm \
      -v ${PWD}:/local openapitools/openapi-generator-cli:v7.5.0 generate \
      -i https://raw.githubusercontent.com/digitalfemsa/openapi/main/_build/api.yaml \
      -g go \
      -o /local \
      -c /local/config-go.json \
      --global-property modelTests=false \
      --additional-properties=hideGenerationTimestamp=true
