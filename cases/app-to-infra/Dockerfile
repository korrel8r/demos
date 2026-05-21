FROM golang:1.25 AS build
WORKDIR /src
COPY go.mod main.go ./
RUN CGO_ENABLED=0 go build -o /bad-app .

FROM scratch
COPY --from=build /bad-app /bad-app
ENTRYPOINT ["/bad-app"]
