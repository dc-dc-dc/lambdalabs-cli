# LambdaLabs CLI

### How to use

Install the binary and include it into path

### Environment Variables Supported

LAMBDA_API_KEY=<your private api key>
DEFAULT_REGION=<default region to use ex: us-west-1>

Alternatively you can create a ~/.lambda file and set the variables there

### Instance commands

List your running instances
```
lambdacli instance list
```

Create an instance
```
lambdacli instance create -region=us-west-1 --type=gpu_1x_a10 -ssh-keys=laptop
```

### SSH Keys
Create an ssh key
```
lambdacli ssh add --name laptop
```