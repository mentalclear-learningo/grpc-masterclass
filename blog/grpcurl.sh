# Create Blog grpcurls
grpcurl -plaintext \
-d '{"blog":{"author_id":"Agent007","title":"This is a test!","content":"Some content here... just for the test"}}' \
localhost:50051 blog.BlogService/CreateBlog

grpcurl -plaintext \
-d '{"blog":{"author_id":"James Bond","title":"Golden Eye","content":"This is another movie about James Bond, Agent 007"}}' \
localhost:50051 blog.BlogService/CreateBlog

# Read(Find) Blog grpcurls 
grpcurl -plaintext \
-d '{"blog_id":"628a2764c28367bace32d7d1"}' \
localhost:50051 blog.BlogService/ReadBlog

# Update Blog grpcurls
grpcurl -plaintext \
-d '{"blog":{"id":"628a44fc929f20f60a571f77", "author_id":"New Author","title":"Fire Eye","content":"This is an updated existing blog"}}' \
localhost:50051 blog.BlogService/UpdateBlog

# Delete Blog grpcurls 
grpcurl -plaintext \
-d '{"blog_id":"628a86d595a248097fc68a5e"}' \
localhost:50051 blog.BlogService/DeleteBlog

# List Blogs grpcurls
grpcurl -plaintext localhost:50051 blog.BlogService/ListBlog