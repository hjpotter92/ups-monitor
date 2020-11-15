#!/bin/sh

SCRIPT_PATH=$(realpath $(dirname $0))
cd "$SCRIPT_PATH"

if [ -d ".venv" ]
then
    . .venv/bin/activate
else
    python3 -m venv .venv
    . .venv/bin/activate
    pip install -U requirements.txt
fi

python3 script.py
