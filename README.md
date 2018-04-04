### Tempatrature stats
This is a test project

#### Run with docker

    docker build -t hello-world .
    docker run --rm --name hello-world -d -p 3000:3333 hello-world

    # Check logs
    docker logs -f hello-world


#### Run manually

    go test
    go build
    ./tel-hello-world


#### Usage

    curl http://localhost:3000
