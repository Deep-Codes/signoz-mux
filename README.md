[Guide](https://signoz.io/docs/instrumentation/golang/)

```sh
SERVICE_NAME=signoxMux INSECURE_MODE=true OTEL_METRICS_EXPORTER=none OTEL_EXPORTER_OTLP_ENDPOINT=127.0.0.1:4317 go run main.go
```

Visit http://localhost:4444/users
