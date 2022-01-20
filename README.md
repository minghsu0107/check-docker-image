# Check Docker Image
A tool to check whether docker images exist in the remote registry.

Build project:
```bash
go build -o check-image .
```
Example usage:
```
export REGISTRY_URL=https://harbor.mycompany.com
export REGISTRY_USERNAME=harboruser
export REGISTRY_PASSWORD=harborpass
export CHECKED_IMAGES=myrepo/mysvc1:v1,myrepo/mysvc2:v1,myrepo/mysvc3:v1
export LOGLEVEL=info
./check-image
```
- `REGISTRY_URL`: registry URL, can be either secure or insecure registry
  - Should be set to `https://registry-1.docker.io` if you are using Dockerhub
- `REGISTRY_USERNAME`: username for registry authentication
  - Can be set empty if all checked images are public
- `REGISTRY_PASSWORD`: password for registry authentication
  - Can be set empty if all checked images are public
- `CHECKED_IMAGES`: comma-separated list of images to be checked
- `LOGLEVEL`: log level, can be `debug`, `info`, `warn`, or `error` (default is `info`)

To obtain the result, you can refer to the following command:
```
./check-image; if [ `echo $?` = '0' ]; then touch SUCCESS; fi
```
If the file `SUCCESS` is created, then the check has passed.

You can also run the docker image directly:
```bash
docker run \
-e REGISTRY_URL=https://harbor.mycompany.com \
-e REGISTRY_USERNAME=harboruser \
-e REGISTRY_PASSWORD=harborpass \
-e CHECKED_IMAGES=myrepo/mysvc1:v7.3.2,myrepo/mysvc2:v8,myrepo/mysvc3:v8.2.0  \
minghsu0107/check-docker-image:v1
```