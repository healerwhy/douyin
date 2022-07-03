docker run --rm --name asynqmon \
    --network test-net \
    -p 8080:8080 \
    hibiken/asynqmon --redis-addr=localhost:6379 --redis-password=password
