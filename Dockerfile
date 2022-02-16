FROM golang:1.17

WORKDIR /Key_Value_Storage
COPY ./ /Key_Value_Storage

RUN go get ./...

EXPOSE 8000

CMD [ "go", "run", "main.go" ]