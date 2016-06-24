FROM scratch
MAINTAINER vishnuk@google.com

ADD memstress /

ENTRYPOINT ["/memstress", "-logtostderr"]
