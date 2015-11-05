FROM scratch

EXPOSE 8000

ADD VERSION /
ADD main /

CMD ["/main"]
