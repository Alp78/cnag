# this Dockerfile creates an image with a unix executable of the app already compiled
# --> env GOOS=linux GOARCH=amd64 go build
# a lot more leightwieght than to embed Golang and compile the script in the container

FROM alpine:3.7
WORKDIR $GOPATH/src/cnag
COPY . $GOPATH/src/cnag
RUN chmod +x $GOPATH/src/cnag

EXPOSE 8080
ENTRYPOINT ./cnag