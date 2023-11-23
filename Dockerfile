FROM golang:1.21-alpine as build
RUN apk add --no-cache ca-certificates
ADD . /go/src/github.com/skpr/cognito-to-dashboard
WORKDIR /go/src/github.com/skpr/cognito-to-dashboard
RUN CGO_ENABLED=0 GOOS=linux GOARCH=${ARCH} go build -a -o bin/cognito-to-dashboard github.com/skpr/cognito-to-dashboard/cmd/cognito-to-dashboard

FROM scratch
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /go/src/github.com/skpr/cognito-to-dashboard/bin/cognito-to-dashboard /usr/local/bin/cognito-to-dashboard
ENV GIN_MODE=release
ENTRYPOINT ["/usr/local/bin/cognito-to-dashboard"]
