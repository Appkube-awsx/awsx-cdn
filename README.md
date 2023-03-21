# CLoudfront CLi's

## To list all the Cloudfront function,run the following command:

```bash
awsx-cloudfront --zone <zone> --acccessKey <acccessKey> --secretKey <secretKey> --crossAccountRoleArn <crossAccountRoleArn> --externalId <externalId> --env <env>
```

## To retrieve the configuration details of a specific CDN function in cloudfrontcmd, run the following command:

```bash
awsx-cloudfront getConfigData -t <table> --zone <zone> --accessKey <accessKey> --secretKey <secretKey> --crossAccountRoleArn <crossAccountRoleArn> --external <externalId> --env <env> --functionName <functionName>
```

## To retrieve the cost and usage details of a specific CDN function in cloudfrontcmd run the following command:

```bash
awsx-cloudfront getCostData -t <table> --zone <zone> --accessKey <accessKey> --secretKey <secretKey> --crossAccountRoleArn <crossAccountRoleArn> --external <externalId> --env <env>
```
