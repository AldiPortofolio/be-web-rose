#!/usr/bin/env bash

export GIT_BRANCH_NAME="feature/merchantGroupVa"
export GIT_COMMIT=$(git rev-list -1 HEAD)
export RELEASE_VERSION=V.0.1
export DATE_TIME="25 Oct 2021"
export TEAM_CREATED="Wayang Squad - Srikandi Team"
export VERSION_TYPE="SIT" # SIT UAT PROD
export NAME_SERVICES="rose-be-go"

# ***************************** 
# Deploy
env GOOS=linux GOARCH=amd64 go build -v -ldflags="-X 'rose-be-go/controllers.GitCommit=$GIT_COMMIT' -X 'rose-be-go/controllers.ReleaseVersion=$RELEASE_VERSION' -X 'rose-be-go/controllers.DateTime=$DATE_TIME' -X 'rose-be-go/controllers.TeamCreated=$TEAM_CREATED'
            -X 'rose-be-go/controllers.VersionType=$VERSION_TYPE' -X 'rose-be-go/controllers.GitBranchName=$GIT_BRANCH_NAME' -X 'rose-be-go/controllers.NameServices=$NAME_SERVICES'"
sudo scp -i /users/abdulah/Documents/devsfa.priv rose-be-go devsfa@34.101.126.12:/home/devsfa/abdullah/rose-be-go
# scp rose-be-go nc_ketut@10.10.43.49:/home/nc_ketut/abdul

# ***************************** 


# ***************************** 
# Local run
# env GOARCH=amd64 go build -v -ldflags="-X 'rose-be-go/controllers.GitCommit=$GIT_COMMIT' -X 'rose-be-go/controllers.ReleaseVersion=$RELEASE_VERSION' -X 'rose-be-go/controllers.DateTime=$DATE_TIME' -X 'rose-be-go/controllers.TeamCreated=$TEAM_CREATED'
#             -X 'rose-be-go/controllers.VersionType=$VERSION_TYPE' -X 'rose-be-go/controllers.GitBranchName=$GIT_BRANCH_NAME' -X 'rose-be-go/controllers.NameServices=$NAME_SERVICES'"

# ./rose-be-go

# ***************************** 

# env GOOS=linux GOARCH=amd64 go build
# sudo scp -i -i ~/.ssh/LightsailDefaultKey-ap-southeast-1-new.pem rose-be-go ubuntu@13.228.25.85:/home/ubuntu/rose-be-go
# sudo scp -i /users/abdulah/OttopayAwsLite.pem rose-be-go ubuntu@13.228.25.85:/home/ubuntu/rose-be-go

