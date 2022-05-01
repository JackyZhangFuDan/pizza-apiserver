FROM fedora
ADD pizza-apiserver /
ENTRYPOINT ["/pizza-apiserver"]