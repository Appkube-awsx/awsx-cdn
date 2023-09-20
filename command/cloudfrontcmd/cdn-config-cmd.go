package cloudfrontcmd

import (
	"fmt"
	"github.com/Appkube-awsx/awsx-common/authenticate"
	"log"

	"github.com/Appkube-awsx/awsx-common/client"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudfront"
	"github.com/spf13/cobra"
)

// GetConfigDataCmd represents the getConfigData command
var GetConfigDataCmd = &cobra.Command{
	Use:   "getConfigData",
	Short: "Config data for cloudfront function",
	Long:  `Config data for cloudfront function`,
	Run: func(cmd *cobra.Command, args []string) {

		authFlag, clientAuth, err := authenticate.SubCommandAuth(cmd)
		if err != nil {
			cmd.Help()
			return
		}

		if authFlag {
			functionName, _ := cmd.Flags().GetString("functionName")
			if functionName != "" {
				GetFunctionList(functionName, *clientAuth)
			} else {
				log.Fatalln("function name not provided. program exit")
			}
		}

	},
}

func GetFunctionList(functionName string, auth client.Auth) *cloudfront.GetFunctionOutput {
	log.Println("Getting aws cloud function Count summary")
	client := client.GetClient(auth, client.CLOUD_FRONT_CLIENT).(*cloudfront.CloudFront)
	input := &cloudfront.GetFunctionInput{
		Name: aws.String(functionName),
	}
	functionResponse, err := client.GetFunction(input)
	if err != nil {
		log.Fatalln("Error:", err)
	}
	log.Println(functionResponse)
	return functionResponse
}

func init() {
	GetConfigDataCmd.Flags().StringP("functionName", "f", "", "function name")

	if err := GetConfigDataCmd.MarkFlagRequired("functionName"); err != nil {
		fmt.Println(err)
	}
}
