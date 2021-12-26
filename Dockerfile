FROM golang:alpine AS build

RUN apk add gcc sqlite sqlite-dev musl-dev

RUN mkdir /globby
WORKDIR /globby

COPY . .

RUN CGO_ENABLED=0 go build

### Deployment
FROM scratch

COPY --from=build /globby/globby /globby

ENTRYPOINT ["/globby"]
