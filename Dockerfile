FROM golang

RUN mkdir /go/src/github.com/Igor-Rast/bit/ -p
RUN git clone https://github.com/Igor-Rast/bit; cp -r bit/* /go/src/github.com/Igor-Rast/bit/
 
ADD . /go/src/github.com/Igor-Rast/front-server
RUN go get github.com/Igor-Rast/front-server ...

RUN ls -la /go/bin/
RUN ls -la /go/src/github.com/*
ENTRYPOINT /go/bin/front-server
 
EXPOSE 8040


