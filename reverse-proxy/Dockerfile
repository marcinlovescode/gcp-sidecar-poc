# Step 1: Modules caching
FROM golang:1.20-alpine AS modules
COPY ./go.mod ./go.sum /modules/
WORKDIR /modules
RUN go mod download

# Step 2: Builder
FROM golang:1.20-alpine AS builder
COPY --from=modules /go/pkg /go/pkg
COPY . /app
WORKDIR /app
RUN go mod tidy && go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
	go build -o /bin/app .

# Step 3: Final
FROM scratch
COPY --from=builder /bin/app /app
CMD ["/app"]