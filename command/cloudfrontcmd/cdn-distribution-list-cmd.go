package cloudfrontcmd

import (
	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-common/client"
	"github.com/aws/aws-sdk-go/service/cloudfront"
	"log"

	"github.com/spf13/cobra"
)

// GetDistributionListCmd represents the getDistributionList command
var GetDistributionListCmd = &cobra.Command{
	Use:   "getDistributionList",
	Short: "Distribution list of cloudfront",
	Long:  `Distribution list of cloudfront`,
	Run: func(cmd *cobra.Command, args []string) {

		authFlag, clientAuth, err := authenticate.SubCommandAuth(cmd)
		if err != nil {
			cmd.Help()
			return
		}

		if authFlag {
			CloudFrontDistributionList(*clientAuth)
		}

	},
}

func CloudFrontDistributionList(auth client.Auth) (*cloudfront.ListDistributionsOutput, error) {
	log.Println("Getting aws cloudfront distribution list")
	client := client.GetClient(auth, client.CLOUD_FRONT_CLIENT).(*cloudfront.CloudFront)
	input := &cloudfront.ListDistributionsInput{}
	response, err := client.ListDistributions(input)
	if err != nil {
		log.Fatalln("Error:", err)
	}
	log.Println(response)
	return response, err
}
