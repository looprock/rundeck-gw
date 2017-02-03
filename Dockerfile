FROM alpine
MAINTAINER Doug Land <dland@ojolabs.com>

COPY rundeck-gw /usr/local/bin/rundeck-gw
RUN chmod +x /usr/local/bin/rundeck-gw
expose 8080
CMD GIN_MODE=release RUNDECK_TOKEN=xxx RUNDECK_URL=http://rundeck.host.com:4440 /usr/local/bin/rundeck-gw
