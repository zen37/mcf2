## Build
FROM golang:1.14-alpine AS build

WORKDIR /src/
COPY main.go go.* /src/
RUN CGO_ENABLED=0 go build -o /mcf2

## Deploy
FROM scratch
COPY --from=build /mcf2 /mcf2
ENTRYPOINT ["/mcf2"]