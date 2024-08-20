FROM golang:alpine AS build

# Builds the application
RUN apk add --update git
WORKDIR /go/src/github.com/rubenbupe/go-auth-server
COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/rubenbupe-go-auth-server ./cmd/api/main.go

# Creates a minimal image with the application
FROM scratch
COPY --from=build /go/bin/rubenbupe-go-auth-server /go/bin/rubenbupe-go-auth-server
ENTRYPOINT ["/go/bin/rubenbupe-go-auth-server"]