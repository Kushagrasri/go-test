FROM golang
RUN mkdir /pull
ADD . /pull
WORKDIR /pull
RUN go build -o main .
CMD ["/pull/main"]