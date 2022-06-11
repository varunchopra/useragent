# useragent

Source code for [usera.gent](https://usera.gent).

To run locally:

    docker run -d -p 8080:8080 varunchopra/useragent

Or download the image from Docker Hub:

    docker pull docker.io/varunchopra/useragent:latest

## Health checks

The `/healthz` endpoint can be used for conducting health checks:

    curl localhost:8080/healthz

## License

MIT (c) 2022 Varun Chopra
