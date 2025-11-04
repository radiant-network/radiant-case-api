# Radiant Case API

Mock API for case creation and updates with Gin + swaggo.

## Requirements
- Go 1.20+
- Docker

## Run Locally

```bash
make install
make run
```

## Build the documentation
```bash
make doc
```

## Display in Swagger
```bash
docker run -p 9090:8080 -e SWAGGER_JSON=/foo/swagger.yaml -v $(pwd)/docs:/foo swaggerapi/swagger-ui
```

Then open http://localhost:9090 in your browser.

Click on authorize and use `test-admin` as bearer token.

## Run reactflow-graph
```bash
cd reactflow-graph
npm install
npm run dev
```

Then open http://localhost:5173 in your browser.