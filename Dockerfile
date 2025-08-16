FROM golang:1.23

COPY main /main
RUN chmod +x /main

CMD ["/main"]
