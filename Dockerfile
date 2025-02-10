FROM golang:1.23-alpine AS builder
ARG PRIVATE_KEY
RUN apk update && apk add --no-cache git openssh-client make vim build-base
ENV GO111MODULE="on"
RUN go env -w GOPRIVATE=gitlab.imperva.local,gitlab,gitlab.com
RUN git config --global url."ssh://git@gitlab:".insteadOf "https://gitlab.com"
RUN mkdir /root/.ssh/ &&\
    echo "${PRIVATE_KEY}" > /root/.ssh/id_rsa &&\
    chmod 600 /root/.ssh/id_rsa &&\
    echo "StrictHostKeyChecking no " > /root/.ssh/config &&\
    echo "Host gitlab" >> /root/.ssh/config &&\
    echo "  IdentityFile ~/.ssh/id_rsa" >> /root/.ssh/config &&\
    echo "  IdentitiesOnly yes" >> /root/.ssh/config
WORKDIR /myapp
COPY . .
RUN go mod tidy


FROM alpine:latest AS runner
RUN apk --no-cache add tzdata bash
WORKDIR /service


FROM builder AS apiserverbuilder
WORKDIR /myapp/apiserver
RUN go build -tags musl -o /app

FROM builder AS apiserver2builder
WORKDIR /myapp/apiserver2
RUN go build -tags musl -o /app

FROM runner AS apiserver
COPY --from=apiserverbuilder /app /service/app
EXPOSE 8080
ENTRYPOINT ["/service/app"]

FROM runner AS apiserver2
COPY --from=apiserver2builder /app /service/app
EXPOSE 8080
ENTRYPOINT ["/service/app"]