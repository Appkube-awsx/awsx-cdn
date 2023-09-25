package command

import (
	"github.com/Appkube-awsx/awsx-common/authenticate"
	"log"
	"os"

	"github.com/Appkube-awsx/awsx-cdn/command/cloudfrontcmd"
	"github.com/Appkube-awsx/awsx-common/client"
	"github.com/aws/aws-sdk-go/service/cloudfront"
	"github.com/spf13/cobra"
)

var awsxCloudFunctionCmd = &cobra.Command{
	Use:   "cloudFunctionListDetails",
	Short: "cloudFunctionListDetails command gets resource counts",
	Long:  `cloudFunctionListDetails command gets resource counts details of an AWS account`,

	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Command get cloud Function List Details started")

		authFlag, clientAuth, err := authenticate.CommandAuth(cmd)
		if err != nil {
			cmd.Help()
			return
		}
		if authFlag {
			CloudFunctionList(*clientAuth)
		} else {
			cmd.Help()
			return
		}
	},
}

func CloudFunctionList(auth client.Auth) (*cloudfront.ListFunctionsOutput, error) {
	log.Println("Getting aws cloudfront function list")
	client := client.GetClient(auth, client.CLOUD_FRONT_CLIENT).(*cloudfront.CloudFront)
	input := &cloudfront.ListFunctionsInput{}
	functionResponse, err := client.ListFunctions(input)
	if err != nil {
		log.Fatalln("Error:", err)
	}
	log.Println(functionResponse)
	return functionResponse, err
}

func Execute() {
	err := awsxCloudFunctionCmd.Execute()
	if err != nil {
		log.Fatal("There was some error while executing the CLI: ", err)
		os.Exit(1)
	}
}

func init() {
	awsxCloudFunctionCmd.AddCommand(cloudfrontcmd.GetFunctionCmd)
	awsxCloudFunctionCmd.AddCommand(cloudfrontcmd.GetDistributionListCmd)
	awsxCloudFunctionCmd.AddCommand(cloudfrontcmd.GetDistributionConfigWithTagsCmd)

	awsxCloudFunctionCmd.PersistentFlags().String("vaultUrl", "", "vault end point")
	awsxCloudFunctionCmd.PersistentFlags().String("vaultToken", "", "vault token")
	awsxCloudFunctionCmd.PersistentFlags().String("accountId", "", "aws account number")
	awsxCloudFunctionCmd.PersistentFlags().String("zone", "", "aws region")
	awsxCloudFunctionCmd.PersistentFlags().String("accessKey", "", "aws access key")
	awsxCloudFunctionCmd.PersistentFlags().String("secretKey", "", "aws secret key")
	awsxCloudFunctionCmd.PersistentFlags().String("crossAccountRoleArn", "", "aws crossAccountRoleArn is required")
	awsxCloudFunctionCmd.PersistentFlags().String("externalId", "", "aws external id auth")

}
