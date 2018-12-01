# Build Stage
FROM lacion/alpine-golang-buildimage:1.11 AS build-stage

LABEL app="build-advent"
LABEL REPO="https://github.com/hfjn/advent"

ENV PROJPATH=/go/src/github.com/hfjn/advent

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

ADD . /go/src/github.com/hfjn/advent
WORKDIR /go/src/github.com/hfjn/advent

RUN make build-alpine

# Final Stage
FROM hfjn/advent

ARG GIT_COMMIT
ARG VERSION
LABEL REPO="https://github.com/hfjn/advent"
LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:/opt/advent/bin

WORKDIR /opt/advent/bin

COPY --from=build-stage /go/src/github.com/hfjn/advent/bin/advent /opt/advent/bin/
RUN chmod +x /opt/advent/bin/advent

# Create appuser
RUN adduser -D -g '' advent
USER advent

ENTRYPOINT ["/usr/bin/dumb-init", "--"]

CMD ["/opt/advent/bin/advent"]
