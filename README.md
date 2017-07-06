This project supplements the scratch container with the ca certificates
of /etc/ssl/certs from the Ubuntu trusty Linux image.

An easy way to grab tge CA certs files to run an interactive shell with 
trusy and to copy the file to your local drive.

To run an interactive terminal session in trusty:

<pre>

</pre>

Then in your trusty session install the certs and copy them to the
directory you mapped to you local drive. The commands are:

<pre>
apt-get update -y
apt-get install ca-certificates
cp /etc/ssl/certs/ca-certificates.crt /foo
</pre>

This was inspired by the a [codeship blog article](https://blog.codeship.com/building-minimal-docker-containers-for-go-applications/)
