# LambdaLabs CLI

### How to use

Create a Lambda Labs API Key and set the environment variable of LAMBDA_API_KEY=<api key here>

### Instance commands

Create an instance
```
go run main.go instance create -region=us-west-1 --type=gpu_1x_a10 -ssh-keys=laptop
```

### SSH Keys
Create an ssh key
```
go run main.go ssh add --name laptop
```