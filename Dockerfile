FROM    golang:1.12

LABEL   org.label-schema.name = "Response API Service"
LABEL   org.label-schema.description = "A backend API for handling database events, auth, datasets, and metadata gathering."
LABEL   org.label-schema.url = "https://github.com/lavrahq/response" 
LABEL   org.label-schema.vcs-url = "https://github.com/lavrahq/response-api"
LABEL   org.label-schema.vendor = "Lavra"
LABEL   org.label-schema.schema-version = "1.0"
LABEL   io.lavra.stack.supported = "true"

RUN     mkdir /app
WORKDIR /app

COPY    go.mod .
COPY    go.sum .

RUN     go mod download

COPY    . .

RUN     go build -o app

ENTRYPOINT [ "./app" ]