# iron/go:dev is the alpine image with the go tools added
FROM iron/go:dev

WORKDIR /app

# Set an env var that matches your github repo name, replace treeder/dockergo here with your repo name

RUN go get github.com/gorilla/handlers
RUN go get github.com/gorilla/mux


RUN mkdir /go/src/github.com/Igor-Rast/bit/ -p
RUN git clone https://github.com/Igor-Rast/bit; cp -r bit/* /go/src/github.com/Igor-Rast/bit/

#ENV SRC_DIR=/go/src/github.com/Igor-Rast/bit/

# Add the source code:
ADD . /go/src/github.com/gorilla/handlers/
ADD . /go/src/github.com/gorilla/mux/



# Build it:
#RUN cd $SRC_DIR;
#RUN go build -o myapp; cp myapp /app/

RUN go build -o serve


#ENTRYPOINT ["./myapp"]
