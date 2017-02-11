FROM golang:1.7
# app specific env
ENV BUILDPATH /go/src/middleware-jwt
ENV TOOLS /var/exec
ENV GOBINARIES /go/bin

# Install reflex
WORKDIR $GOBINARIES
RUN go get github.com/cespare/reflex
RUN go get github.com/pborman/uuid
RUN go get github.com/dgrijalva/jwt-go
RUN go get golang.org/x/crypto/bcrypt
RUN go get github.com/garyburd/redigo/redis
RUN go get github.com/codegangsta/negroni
RUN go get github.com/gorilla/mux

ENV PORT 6060
# dockeer/app port
EXPOSE $PORT

# Make directories and add files as needed
RUN mkdir -p $TOOLS
ADD build.sh $TOOLS
ADD reflex.conf $TOOLS
RUN chmod +x $TOOLS/build.sh

# Execute reflex.
WORKDIR $BUILDPATH
CMD ["reflex","-c","/var/exec/reflex.conf"]
