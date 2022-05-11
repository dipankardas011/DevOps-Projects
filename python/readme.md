# command to start

## For the Devlopment
```bash
cd python/
docker build --target dev -t python-dev .
docker run --privileged=true -it --rm -v $(pwd):/work python sh
```

## For the prod
```bash
cd python/
docker build --target runtime -t python-prod .
docker run -it --rm python-prod
```
