#!/bin/sh

docker run --rm -v `pwd`:/work --entrypoint=/work/docker_solve.sh python:3.9
