FROM golang:1.16-alpine
COPY src /home/src
WORKDIR /home/src
RUN go mod init db_service
RUN go mod tidy
RUN go build .
CMD ./db_service