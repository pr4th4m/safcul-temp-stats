### Tempatrature stats
**This is a test project**
Display temperature stats for fridges


#### Tested On
- Golang version: go1.10 darwin/amd64


#### Run with docker

    docker build -t "safcul-temp-stats:dev" .
    docker run --rm --name safcul-temp-stats -it safcul-temp-stats:dev

    # with file flag
    docker run --rm --name safcul-temp-stats -it safcul-temp-stats:dev --file data/fridge_temperature.json


#### Run manually

    cd safcul-temp-stats

    go test
    go build -o stats *.go
    ./stats

    # with file flag
    ./stats --file data/fridge_temperature.json
