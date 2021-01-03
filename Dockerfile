FROM golang:1.15.5-alpine AS build

RUN apk add build-base

WORKDIR /
COPY servirtium.go go.* *.tmpl /
COPY cmd/* /cmd/

RUN CGO_ENABLED=0 go build -v -o /bin/servirtium_todobackend_compatibility cmd/todobackend_compatibility.go
RUN chmod 755 /bin/servirtium_todobackend_compatibility

FROM scratch
COPY --from=build /bin/servirtium_todobackend_compatibility /bin/servirtium_todobackend_compatibility
ENTRYPOINT ["/bin/servirtium_todobackend_compatibility"]