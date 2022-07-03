1. docker-compose.yml 部署的是Zookeeper和Kafka ```docker-compose up -d```启动

2. asynqUI 是Asynq配套的UI，可以在浏览器中查看任务的执行状态，他的镜像通过```docker pull hibiken/asynqmon```下载

3. 其他的Dockerfile都是service下的各个文件的镜像生成文件，使用时需要将对应的Dockerfile在项目根目录下进行```docker build .```构建镜像