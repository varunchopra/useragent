FROM scratch

ADD useragent /useragent
ADD user-agents.json /user-agents.json

EXPOSE 8080

ENTRYPOINT ["/useragent"]
