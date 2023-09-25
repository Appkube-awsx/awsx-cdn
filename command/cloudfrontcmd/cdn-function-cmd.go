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

// GetFunctionCmd represents the getCdnFunction command
var GetFunctionCmd = &cobra.Command{
	Use:   "getCdnFunction",
	Short: "Cloudfront function",
	Long:  `Cloudfront function`,
	Run: func(cmd *cobra.Command, args []string) {

		authFlag, clientAuth, err := authenticate.SubCommandAuth(cmd)
		if err != nil {
			cmd.Help()
			return
		}

		if authFlag {
			functionName, _ := cmd.Flags().GetString("functionName")
			if functionName != "" {
				GetCdnFunction(functionName, *clientAuth)
			} else {
				log.Fatalln("function name not provided. program exit")
			}
		}

	},
}

func GetCdnFunction(functionName string, auth client.Auth) *cloudfront.GetFunctionOutput {
	log.Println("Getting aws cloudfront function")
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
	GetFunctionCmd.Flags().StringP("functionName", "f", "", "function name")

	if err := GetFunctionCmd.MarkFlagRequired("functionName"); err != nil {
		fmt.Println(err)
	}
}
