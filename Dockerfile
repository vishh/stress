FROM scratch
MAINTAINER vishnuk@google.com

ADD stress /

ENTRYPOINT ["/stress", "-logtostderr"]
