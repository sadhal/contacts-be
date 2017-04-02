FROM golang:1.8

EXPOSE 6767
USER nobody

RUN mkdir -p /go/src/github.com/sadhal/contacts-be/userservice
WORKDIR /go/src/github.com/sadhal/contacts-be/userservice

COPY . /go/src/github.com/sadhal/contacts-be/userservice
RUN go-wrapper download && go-wrapper install

CMD ["go-wrapper", "run"]