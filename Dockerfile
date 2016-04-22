FROM scratch

EXPOSE 8000

ADD VERSION /
ADD main.up /

CMD ["/main.up"]
