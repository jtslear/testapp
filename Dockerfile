FROM scratch

EXPOSE 8000

ADD testapp.up /

CMD ["/testapp.up"]
