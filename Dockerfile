from golang

run mkdir /app

ADD . /app

WORKDIR /app

RUN go build .

EXPOSE 6666

CMD ["/app/notsunnyservice"]