Build a small image that can talk to AWS. The secret sauce is the scratchy
base image, which is scratch plus the ca-certificates.crt from
Ubuntu

You can run on premise like this:

<pre>
docker run -e BUCKET=$BUCKET -e AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID -e AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY -e AWS_REGION=$AWS_REGION -e http_proxy=$http_proxy xtracdev/bucketcat
</pre>

This can be mapped to an ECS task def to try out in Amazon.
