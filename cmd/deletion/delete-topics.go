package deletion

import (
	"github.com/deviceinsight/kafkactl/operations"
	"github.com/spf13/cobra"
)

var cmdDeleteTopic = &cobra.Command{
	Use:   "topic",
	Short: "delete a topic",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		(&operations.TopicOperation{}).DeleteTopics(args)
	},
}

func init() {
}
