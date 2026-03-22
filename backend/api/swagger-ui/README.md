## Update Swagger UI
Download the latest Swagger UI dist from the official [GitHub Releases page](https://github.com/swagger-api/swagger-ui/releases/latest). Copy and overwrite the `dist` folder in `swagger-ui` module. Replace the `url` inside the `swagger-initializer.js` file by `/spec/api.yaml`.

## Usage
```go
func main() {
	e := echo.New()
	// This is where the api.yaml is located
	e.Static("/spec", "api/spec")
	// This is where the swagger ui files are located
	e.Static("/swagger-ui", "api/swagger-ui/dist")
}
```

## Reference
https://medium.com/@ribice/serve-swaggerui-within-your-golang-application-5486748a5ed4

https://github.com/flowchartsman/swaggerui