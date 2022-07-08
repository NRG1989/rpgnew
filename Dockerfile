FROM golang:latest

RUN go version
ENV GOPATH=/
ARG GITLAB_USER
ARG GITLAB_TOKEN
ENV GOPRIVATE=git.andersenlab.com

COPY ./ ./

# install psql
RUN apt-get update
RUN apt-get -y install postgresql-client

# set credential to privat repo
RUN echo "machine git.###.com login ${GITLAB_USER} password ${GITLAB_TOKEN}" > ~/.netrc

# build go app
RUN go mod download
RUN go build -o rpg-api-otp ./cmd/main.go

EXPOSE 5014

CMD ["./rpg-api-otp"]