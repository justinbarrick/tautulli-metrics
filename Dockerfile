FROM scratch
COPY ./tautulli-metrics /tautulli-metrics
ENTRYPOINT ["/tautulli-metrics"]
