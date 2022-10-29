FROM golang:1.19.2 AS build

WORKDIR /app

COPY . .

RUN go build ./src/server.go

FROM gcr.io/distroless/base-debian11

WORKDIR /

COPY --from=build /app/config.json .
COPY --from=build /app/test.db .
COPY --from=build /app/server .

EXPOSE 8080

#USER nonroot:nonroot

ENTRYPOINT [ "/server" ]
