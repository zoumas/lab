# Backend Configuration - Environment Variables

The following is a list of environment variables the server expects to be set.

The environment variable names are not final.

## Note

During local development, a `-local` flag can be used to load a .env file from the module root path.
Otherwise, the server considers the environment variables to be set from the DevOps team and will error on absence.

## Variables

* `ADDR` - `hostname:port` the address the server will bind to and listen for incoming requests.
* `CORS_ORIGIN` - the single URL that will be allowed to make requests to the server. This is to be the frontend.
* `DSN` - the Data Source Name for the database
