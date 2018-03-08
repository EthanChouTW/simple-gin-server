FROM golang:1.8.1-alpine
RUN apk add -f --no-cache git
RUN go get github.com/codegangsta/gin
RUN go get github.com/gin-gonic/gin
RUN go get github.com/gin-contrib/cors
RUN go get github.com/parnurzeal/gorequest
WORKDIR /go/src/github.com/EthanChouTW/studygroup/src/
ENV PORT=5001
RUN export PORT=$PORT
CMD ["gin", "-a", "5000", "-p", "5001"]
# CMD ["go", "run", "main.go"]