gcloud functions deploy visitorCounter \
  --project=cis437-hw4-476803 \
  --runtime=go125 \
  --trigger-http \
  --entry-point=VisitorCounter \
  --allow-unauthenticated \
  --region=us-central1