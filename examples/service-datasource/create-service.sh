#!/usr/bin/env bash

set -euox pipefail

# upsert team
echo '{
  "team": "{
    "name": "Team 1",
    "slug": "t1"
  },
  "create_if_missing": true
}' | curl -X PUT --data @- -H "API-Token: $CHRONOSPHERE_API_TOKEN" "https://$CHRONOSPHERE_ORG.chronosphere.io/api/v1/config/teams/t1"

echo '{
  "service": {
    "name": "http_server",
    "slug": "http_server",
    "team_slug": "t1",
    "notification_policy_slug": "np-t1",
    "description": "http_server is a monolith which powers non-data APIs"
  },
  "create_if_missing": true
}' | curl -X PUT --data @- -H "API-Token: $CHRONOSPHERE_API_TOKEN" "https://$CHRONOSPHERE_ORG.chronosphere.io/api/v1/config/services/gateway"
