@echo off
echo.

set bggroupSpec=specs/bggroup.yml

echo "...generating bggroup server"
swagger generate server -P rest_model_bggroup.Principal -f %bggroupSpec% -s rest_server_bggroup -t . -m "rest_model_bggroup" --exclude-main

echo "...generating bggroup client"
swagger generate client -P rest_model_bggroup.Principal -f %bggroupSpec% -c rest_client_bggroup -t . -m "rest_model_bggroup"

echo "...generating js client"
swagger-codegen generate -i %bggroupSpec% -l typescript-fetch -o ui/src/api