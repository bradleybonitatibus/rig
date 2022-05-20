#!/bin/sh

semgrep --error --metrics=on --strict --config=p/golang -o sast.json --json
