#!/usr/bin/env bash
docker run -u $(id -u):$(id -g) -ti -v mcs2lm:/app-bin -v ${PWD}:/app --rm mcs2lm  $@
