FROM golang:1.23-alpine AS build
WORKDIR /src
COPY go.mod ./
COPY main.go ./
RUN CGO_ENABLED=0 go build -trimpath -ldflags "-s -w" -o /api .

FROM gcr.io/distroless/static:nonroot
COPY --from=build /api /api
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/api"]
