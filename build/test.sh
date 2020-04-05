sudo docker run \
  -v $(pwd):/home/autosync/data/ \
  -v $(pwd)/configs:/home/autosync/configs \
  -e ENV="local" \
  --rm \
  binayakd/autosync:v1