# cliphub
A cross-device web app to allow user sync messages and files.
1. Use gin framework to implement Go web app
2. Message is stored in memory, and file is stored in local
3. Use go embed to package static files
4. Use docker scratch to run web app in docker container

# How to build and deploy
Recommend to use docker to build and deploy:
1. Run `make` to generate docker image in local
2. Run `docker push forzart/cliphub-scratch:latest` to push you local image to dockerhub
3. Run `docker run -d --name cliphub-scratch -p 1122:9000/tcp --restart=always -w /home/golab/cliphub forzart/cliphub-scratch` to run container
