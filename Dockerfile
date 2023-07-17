#FROM ##########################/default-gitlab-ci/golang:1.18-alpine as builder
#
#RUN apk update && apk upgrade --ignore alpine-baselayout && \
#    apk add --no-cache bash git openssh curl tzdata
#
#RUN git config --global url."###################################".insteadOf "https://gitlab.##########"
#
#WORKDIR /build
#
#COPY go.mod go.sum* .env ./
#RUN go env -w GOPRIVATE=gitlab.########
#RUN go mod download
#
#COPY . .
#
#RUN CGO_ENABLED=0 GOOS=linux go build -o /main main.go
#
#FROM scratch
#
#COPY --from=builder /etc/ssl/###################
#COPY --from=builder main .
#COPY --from=builder /build/.env .
#COPY --from=builder /usr/share/############
#ENV TZ=Asia/Almaty
#
#EXPOSE ####
#
#ENTRYPOINT ["./main", "-config-path=./.env"]
#
