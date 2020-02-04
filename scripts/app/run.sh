#!/usr/bin/env bash
docker run -u $(id -u):$(id -g) -ti -v mcs2lm:/app-bin --rm mcs2lm /app-bin/mcs2lm $@
