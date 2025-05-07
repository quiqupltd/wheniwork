#!/bin/bash

esc run wheniwork/prod -- bash -c 'curl -d "{\"email\": \"$WHENIWORK_EMAIL\",\"password\":\"$WHENIWORK_PASSWORD\"}" https://api.wheniwork.com/2/login -H "W-Key: $WHENIWORK_KEY"'