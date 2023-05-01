package cloudfrontcmd

import (
	"fmt"
	"log"

	"github.com/Appkube-awsx/awsx-cloudfront/authenticator"
	"github.com/Appkube-awsx/awsx-cloudfront/client"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudfront"
	"github.com/spf13/cobra"
)

// getConfigDataCmd represents the getConfigData command
var GetConfigDataCmd = &cobra.Command{
	Use:   "getConfigData",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		vaultUrl := cmd.Parent().PersistentFlags().Lookup("vaultUrl").Value.String()
		accountNo := cmd.Parent().PersistentFlags().Lookup("accountId").Value.String()
		region := cmd.Parent().PersistentFlags().Lookup("zone").Value.String()
		acKey := cmd.Parent().PersistentFlags().Lookup("accessKey").Value.String()
		secKey := cmd.Parent().PersistentFlags().Lookup("secretKey").Value.String()
		crossAccountRoleArn := cmd.Parent().PersistentFlags().Lookup("crossAccountRoleArn").Value.String()
		externalId := cmd.Parent().PersistentFlags().Lookup("externalId").Value.String()

		authFlag := authenticator.AuthenticateData(vaultUrl, accountNo, region, acKey, secKey, crossAccountRoleArn,  externalId)
		print(authFlag)
		// authFlag := true
		if authFlag {
			functionName, _ := cmd.Flags().GetString("functionName")
			if functionName != "" {
				getFunctionList(region, crossAccountRoleArn, acKey, secKey, functionName, externalId)
			} else {
				log.Fatalln("functionName not provided. Program exit")
			}
		}
	},
}

func getFunctionList(region string, crossAccountRoleArn string, accessKey string, secretKey string, functionName string, externalId string) *cloudfront.GetFunctionOutput {
	log.Println("Getting aws cloud function Count summary")
	getClient := client.GetClient(region, crossAccountRoleArn, accessKey, secretKey, externalId)
	input := &cloudfront.GetFunctionInput{
		Name: aws.String(functionName),
	}
	functionResponse, err := getClient.GetFunction(input)
	if err != nil {
		log.Fatalln("Error:", err)
	}
	log.Println(functionResponse)
	return functionResponse
}

func init() {
	GetConfigDataCmd.Flags().StringP("functionName", "t", "", "function Name")

	if err := GetConfigDataCmd.MarkFlagRequired("functionName"); err != nil {
		fmt.Println(err)
	}
}
