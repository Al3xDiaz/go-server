FROM python:3.9-slim-bullseye as generatestaticfiles
ARG POSTMAN_DOC_GEN_VERSION=1.1.0
WORKDIR /var/usr/data
ADD  https://github.com/karthiks3000/postman-doc-gen/archive/refs/tags/${POSTMAN_DOC_GEN_VERSION}.tar.gz /var/usr/data/postman-doc-gen.tar.gz
RUN tar -xf /var/usr/data/postman-doc-gen.tar.gz
WORKDIR /var/usr/data/postman-doc-gen-${POSTMAN_DOC_GEN_VERSION}
RUN  pip install --upgrade pip && pip install  -r requirements.txt
COPY ./postmanfiles/ .
RUN python postman_doc_gen/postman_doc_gen.py collection.json -e environment.json -o /var/usr/static

# Choose whatever you want, version >= 1.16
FROM golang:1.21.1-alpine
WORKDIR /go-server
RUN go install github.com/cosmtrek/air@latest
COPY go.mod go.sum ./
RUN go mod download
COPY . .
COPY --from=generatestaticfiles /var/usr/static ./static
CMD ["air", "-c", ".air.toml"]