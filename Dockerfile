FROM aankittcoolest/mecab-mozc:1.0

COPY --from=golang:1.21-alpine /usr/local/go/ /usr/local/go/

ENV PATH="/usr/local/go/bin:${PATH}"

RUN mkdir -p /usr/src/app && cd /usr/src/app
WORKDIR /usr/src/app
COPY ./src/go.mod .
COPY ./src/go.sum .
RUN go mod download
COPY ./src/ .

RUN go build -o ../builds/jp-search-combinations main.go

CMD ["../builds/jp-search-combinations"]