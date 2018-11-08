FROM scratch

COPY bin/* ./

COPY certs/* ./cert

EXPOSE 5000

ENTRYPOINT ["./daemon"]



