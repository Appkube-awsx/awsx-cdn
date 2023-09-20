package controller

import (
	"github.com/Appkube-awsx/awsx-cdn/command"
	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-common/client"
	"github.com/aws/aws-sdk-go/service/cloudfront"
	"log"
)

func GetCdnByAccountNo(vaultUrl string, vaultToken string, accountNo string, region string) (*cloudfront.ListFunctionsOutput, error) {
	authFlag, clientAuth, err := authenticate.AuthenticateData(vaultUrl, vaultToken, accountNo, region, "", "", "", "")
	return GetCdnByFlagAndClientAuth(authFlag, clientAuth, err)
}

func GetCdnByUserCreds(region string, accessKey string, secretKey string, crossAccountRoleArn string, externalId string) (*cloudfront.ListFunctionsOutput, error) {
	authFlag, clientAuth, err := authenticate.AuthenticateData("", "", "", region, accessKey, secretKey, crossAccountRoleArn, externalId)
	return GetCdnByFlagAndClientAuth(authFlag, clientAuth, err)
}

func GetCdnByFlagAndClientAuth(authFlag bool, clientAuth *client.Auth, err error) (*cloudfront.ListFunctionsOutput, error) {
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	if !authFlag {
		log.Println(err.Error())
		return nil, err
	}
	response, err := command.CloudFunctionList(*clientAuth)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return response, nil
}

func GetCdn(clientAuth *client.Auth) (*cloudfront.ListFunctionsOutput, error) {
	response, err := command.CloudFunctionList(*clientAuth)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return response, nil
}
