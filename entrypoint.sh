#!/bin/sh

# Starts varnish
varnishd -a :8080 -f /etc/varnish/default.vcl -s malloc,512m &

echo "-------------------- NOTE --------------------"
echo ""
echo "SERVING CACHED PROJECT ON http://localhost:8080."
echo ""
echo "-------------------- NOTE --------------------"

# Starts API
./main