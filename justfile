build:
    @go build -o ./bin/lawnchair ./cmd/lawnchair

run:
    ./bin/lawnchair run -c ./bin/config.yml