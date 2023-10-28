FROM golang

COPY . /

WORKDIR /

RUN go mod download

RUN go build 

CMD ["./todo"]