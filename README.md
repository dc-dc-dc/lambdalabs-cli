# LambdaLabs CLI

### How to use

Install the binary and include it into path

### Environment Variables Supported

```env
LAMBDA_API_KEY=<your private api key>
DEFAULT_REGION=<default region to use ex: us-west-1>
```

Alternatively you can create a ~/.lambda file and set the variables there

## Instance commands

List your running instances
```bash
lambdacli instance list
```

Create an instance
```bash
lambdacli instance create -region=us-west-1 --type=gpu_1x_a10 -ssh-keys=laptop
```

## SSH Keys
List ssh keys
```bash
lambdacli ssh list
```

Create an ssh key
```bash
lambdacli ssh add --name laptop
```