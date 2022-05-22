grpcurl -plaintext \
-d '{"blog":{"author_id":"Agent007","title":"This is a test!","content":"Some content here... just for the test"}}' \
localhost:50051 blog.BlogService/CreateBlog

grpcurl -plaintext \
-d '{"blog":{"author_id":"James Bond","title":"Golden Eye","content":"This is another movie about James Bond, Agent 007"}}' \
localhost:50051 blog.BlogService/CreateBlog

grpcurl -plaintext \
-d '{"blog_id":"628a2764c28367bace32d7d1"' \
localhost:50051 blog.BlogService/ReadBlog