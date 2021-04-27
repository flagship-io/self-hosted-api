<p align="center">

<img  src="https://mk0abtastybwtpirqi5t.kinstacdn.com/wp-content/uploads/picture-solutions-persona-product-flagship.jpg"  width="211"  height="182"  alt="Flagship"  />

</p>

<h3 align="center">Bring your features to life</h3>

**Visit [https://developers.flagship.io/](https://developers.flagship.io/) to get started with Flagship.**

# Docs

## Running the app

### Running from build

Download the latest release on github and then simply run:

`ENV_ID={your_environment_id} API_KEY={your_api_key} ./app`

The server will run on the port 8080

### Running with Docker

Run the following command to start the server with Docker

`docker run -p 8080:8080 -e ENV_ID={your_env_id} -e API_KEY={your_api_key} flagshipio/self-hosted-api`

## Configuration

You can configure the self-hosted Decision API using 2 ways:

- YAML configuration file
- Environment Variables

### Using a configuration file

Create a `config.yaml` along your app file, or mount it in docker in location /config.yaml:

`docker run -p 8080:8080 -v ./config.yaml:/config.yaml flagshipio/self-hosted-api`

The configuration file should look like this:

```yaml
env_id: "env_id" # Your Flagship Environment ID
api_key: "api_key" # Your Flagship API Key

# Cache
cache:
  type: local # or 'redis' or 'none' (if you do not want to using visitor cache)
  options:
    dbPath: ./data
    #redisHost: 'localhost:6379' # for redis storage
    #redisUsername: username     # (Optional) for redis storage
    #redisPassword: password     # (Optional) for redis storage
```

### Using environment variables

You can override each configuration variables from the configuration file using environment variables.
Just name your env variables the same as the config file, but with the following rules:

- Env variable name should be UPPERCASE
  Example: ENV_ID
- Sub configuration level are defined using a `_` sign
  Example: CACHE_TYPE
  Example: CACHE_OPTIONS_DBPATH

Here is a Docker example using environment variables to setup local caching:

`docker run -p 8080:8080 -e ENV_ID={your_env_id} -e API_KEY={your_api_key} -e CACHE_TYPE=local -e CACHE_OPTIONS_DBPATH=./data -v ./config.yaml:/config.yaml flagshipio/self-hosted-api`

Here is a Docker Compose example of using Redis as a visitor cache engine:

```yaml
version: "3"
services:
  decision:
    build: .
    ports:
      - 8080:8080
    environment:
      ENV_ID: "env_id"
      API_KEY: "api_key"
      CACHE_TYPE: redis
      CACHE_OPTIONS_REDISHOST: "redis:6379"
    depends_on:
      - redis

  redis:
    image: redis
```

### Configuration parameters

You can use the following parameters to customize the Self Hosted API.
Each parameter is named as in the config.yaml file, and the matching environment variable is parenthesis.

| Parameter                                                 |  Type  | Required |                                                     Description                                                     |
| --------------------------------------------------------- | :----: | -------: | :-----------------------------------------------------------------------------------------------------------------: |
| env_id (ENV_ID)                                           | string |      yes |           The Flagship environment ID. You can get it from the Flagship platform. Default to empty string           |
| api_key (API_KEY)                                         | string |      yes |  The Flagship API Key for this environment ID. You can get it from the Flagship platform. Default to empty string   |
| polling_interval (POLLING_INTERVAL)                       |  int   |       no |          The polling frequency (in seconds) to synchronize with your Flagship configuration. Default to 60          |
| cache.type (CACHE_TYPE)                                   | string |       no |    If you want to enable caching for the visitor assignment. Can be "redis" or "local". Default to empty string.    |
| cache.options.dbPath (CACHE_OPTIONS_DBPATH)               | string |       no | If you chose local cache type, this is the path of the file where the cache will be stored. Default to empty string |
| cache.options.redisHost (CACHE_OPTIONS_REDISHOST)         | string |       no |                        If you chose redis cache type, this is the host for your redis server                        |
| cache.options.redisUsername (CACHE_OPTIONS_REDISUSERNAME) | string |       no |                      If you chose redis cache type, this is the username for your redis server                      |
| cache.options.redisUsername (CACHE_OPTIONS_REDISPASSWORD) | string |       no |                      If you chose redis cache type, this is the password for your redis server                      |
| cache.options.redisDb (CACHE_OPTIONS_REDISDB)             |  int   |       no |        If you chose redis cache type, this is the db number for your redis server. Default to 0 (default DB)        |

## Usage

You can find the Swagger API doc at the `/swagger/index.html` URL when running the application.

# Contribute

## Requirements

You need Go 1.12+ to build the app

## Building

Run the following command to build
`CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app`

## Generate swagger files

Install `swaggo/swag` to build swagger files:
`go get -u github.com/swaggo/swag/cmd/swag`

Then run the init command with parseDependency option:
`swag init --parseDependency`
