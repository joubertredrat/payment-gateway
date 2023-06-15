FROM golang:1.18-alpine

RUN apk add --no-cache gcc g++ make git

WORKDIR /go/src/app

RUN wget https://raw.githubusercontent.com/eficode/wait-for/v2.2.4/wait-for -O ./wait-for && chmod +x ./wait-for

ADD . .

RUN go get

RUN go build -o app .

RUN apk del git

ENTRYPOINT [ "/go/src/app/entrypoint.sh" ]
