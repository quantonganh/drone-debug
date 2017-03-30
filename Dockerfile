FROM alpine:3.4

ADD drone-debug /bin/
ENTRYPOINT ["/bin/drone-debug"]
