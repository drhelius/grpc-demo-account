FROM registry.access.redhat.com/ubi8-minimal
COPY bin/grpc-demo-account /
EXPOSE 8080 5000
CMD ["/grpc-demo-account"]