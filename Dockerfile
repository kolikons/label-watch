FROM scratch
COPY label-watch /usr/bin/label-watch
ENTRYPOINT ["/usr/bin/label-watch"]
