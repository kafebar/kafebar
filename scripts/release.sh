#!/bin/bash

TAG=europe-west4-docker.pkg.dev/kafebar/main/kafebar

docker build . --tag=$TAG

docker push $TAG

gcloud run deploy kafebar \
    --project kafebar \
    --image $TAG \
    --region europe-west4