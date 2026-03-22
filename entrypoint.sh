#!/bin/sh

# Start Backend in the Background
/app/backend_app &

# Start Nginx (Serving Frontend)
nginx -g "daemon off;"
