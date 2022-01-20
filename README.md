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
- `REGISTRY_USERNAME`: username for registry authentication
- `REGISTRY_PASSWORD`: password for registry authentication
- `CHECKED_IMAGES`: comma-separated list of images to be checked
- `LOGLEVEL`: log level, can be `debug`, `info`, `warn`, or `error` (default is `info`)

Or run the docker image directly:
```bash
docker run \
-e REGISTRY_URL=https://harbor.mycompany.com \
-e REGISTRY_USERNAME=harboruser \
-e REGISTRY_PASSWORD=harborpass \
-e CHECKED_IMAGES=myrepo/mysvc1:v1,myrepo/mysvc2:v1,myrepo/mysvc3:v1  \
minghsu0107/check-docker-image
```