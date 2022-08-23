FROM mongo

WORKDIR /home/

RUN apt-get update -y
RUN apt-get install -y curl

RUN curl -O -L "https://golang.org/dl/go1.19.linux-amd64.tar.gz"

RUN tar -xf "go1.19.linux-amd64.tar.gz"

RUN mv -v go /usr/local

COPY *go* /home/

COPY ./script.sh .

EXPOSE 8080

CMD ["sh", "script.sh"]

