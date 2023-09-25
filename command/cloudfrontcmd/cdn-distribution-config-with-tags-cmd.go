package cloudfrontcmd

import (
	"encoding/json"
	"fmt"
	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-common/client"
	"github.com/aws/aws-sdk-go/service/cloudfront"
	"log"

	"github.com/spf13/cobra"
)

// GetDistributionConfigWithTagsCmd represents the getDistributionConfigWithTags command
var GetDistributionConfigWithTagsCmd = &cobra.Command{
	Use:   "getDistributionConfigWithTags",
	Short: "Distribution Config with Tags of cloudfront",
	Long:  `Distribution Config with Tags of cloudfront`,
	Run: func(cmd *cobra.Command, args []string) {

		authFlag, clientAuth, err := authenticate.SubCommandAuth(cmd)
		if err != nil {
			cmd.Help()
			return
		}

		if authFlag {
			CdnDistributionConfigWithTagList(*clientAuth)
		}

	},
}

type Cdn struct {
	Distribution       interface{} `json:"distribution"`
	DistributionConfig interface{} `json:"distribution_config"`
	Tags               interface{} `json:"tags"`
}

func CdnDistributionConfigWithTagList(auth client.Auth) (string, error) {
	log.Println("Getting aws cloudfront distribution config list with tags")
	distributionList, err := CloudFrontDistributionList(auth)
	if err != nil {
		log.Fatalln("Error:", err)
	}
	client := client.GetClient(auth, client.CLOUD_FRONT_CLIENT).(*cloudfront.CloudFront)
	all := []Cdn{}
	for _, distributionItem := range distributionList.DistributionList.Items {
		configInput := &cloudfront.GetDistributionConfigInput{
			Id: distributionItem.Id,
		}
		distributionConfigOutput, err := client.GetDistributionConfig(configInput)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		tagInput := &cloudfront.ListTagsForResourceInput{
			Resource: distributionItem.ARN,
		}
		tagOutput, err := client.ListTagsForResource(tagInput)

		cdn := Cdn{
			Distribution:       distributionItem,
			DistributionConfig: distributionConfigOutput,
			Tags:               tagOutput,
		}
		all = append(all, cdn)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

	}
	jsonData, err := json.Marshal(all)
	log.Println(string(jsonData))
	return string(jsonData), err
}
