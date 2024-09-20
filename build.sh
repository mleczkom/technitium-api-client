#!/bin/bash
rm -rR technitium
~/bin/openapitools/openapi-generator-cli generate -i api.yaml -g go -c api_config.yaml -o ./technitium --git-user-id mleczkom --git-repo-id technitium-api-client
rm -rR ./technitium/test 
#rm ./technitium/go.*
#go mod tidy