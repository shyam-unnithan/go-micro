########################################
# Build image for the api service
########################################
FROM builder

# Import the user and groups
# This will be used to run the service
# with a lower privileged user

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

ADD ./bin/service/event/cmd /go/bin/event
ADD ./bin/service/event/config /go/config/
USER appuser:appuser

CMD ["/go/bin/event"]
