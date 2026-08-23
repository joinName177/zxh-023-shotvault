FROM golang:1.26-alpine AS build
WORKDIR /src
COPY . .
RUN go build -trimpath -o /out/shotvault ./cmd/shotvault
FROM alpine:3.22
COPY --from=build /out/shotvault /usr/local/bin/shotvault
ENTRYPOINT ["/usr/local/bin/shotvault"]
