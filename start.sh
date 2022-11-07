#!/bin/sh

SCRIPT_PATH=$(realpath $(dirname $0))
cd "$SCRIPT_PATH"

echo "Starting UPS metric update script"

if [ -d ".venv" ]; then
  source .venv/bin/activate
  echo "Activated virtual env"
else
  python3 -m venv .venv
  echo "Created new virtual env"
  source .venv/bin/activate
  pip install -U requirements.txt
  echo "Update venv with required packages"
fi

echo "Running script now"
python3 script.py
echo "Script finished"
