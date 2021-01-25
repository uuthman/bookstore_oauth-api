FROM golang:alpine


ENV REPO_URL = github.com/uuthman/bookstore_oauth-api

ENV GOPATH=/app

ENV APP_PATH = $GOPATH/src/$REPO_URL

ENV WORKPATH=$APP_PATH/src
COPY src $WORKPATH
WORKDIR $WORKPATH

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o oauth-api .

EXPOSE 8080

CMD ["./oauth-api"]