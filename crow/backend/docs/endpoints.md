# Endpoints

The following is a list of all the supported endpoints.

* `GET /healthz`

Doesn't expect a request body and will respond with a 200 status code when the service is available or a 503 when not.

Note: A 503 will be returned when we support graceful shutdown.
