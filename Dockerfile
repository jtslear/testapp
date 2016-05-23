FROM scratch

EXPOSE 8000

ADD VERSION /
ADD testapp.up /

CMD ["/testapp.up"]
