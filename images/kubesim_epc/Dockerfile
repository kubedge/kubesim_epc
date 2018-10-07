FROM scratch
ADD kubesim_epc/ca-certificates.crt /etc/ssl/certs/
ADD kubesim_epc/goserv /goserv
EXPOSE 8080
CMD ["/goserv"]
