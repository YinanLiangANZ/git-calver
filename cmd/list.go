package cmd

import (
	"fmt"
	"github.com/ktr0731/go-fuzzyfinder"
	"github.com/socialviolation/git-calver/git"
	"github.com/spf13/cobra"
)

var listTagCommand = &cobra.Command{
	Use:   "list",
	Short: "Will list all CalVer tags matching the provided format",
	Run: func(cmd *cobra.Command, args []string) {
		f := loadFormat()
		tags, err := git.List(f)
		CheckIfError(err)

		fmt.Printf("found tags: %d\n", len(tags))
		idx, err := fuzzyfinder.Find(tags, func(i int) string {
			return tags[i].Short
		}, fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
			if i == -1 {
				return ""
			}
			return fmt.Sprintf(`Ref: %s
IsBranch: %t
Hash: %s
Commit: %s
`,
				tags[i].Ref,
				tags[i].IsBranch,
				tags[i].Hash,
			)
		}))

		if err != nil {
			fmt.Printf("Cancelled.\n")
			return
		}

		fmt.Printf("You selected %d - %s", idx, tags[idx].Short)
	},
}

func init() {
	rootCmd.AddCommand(listTagCommand)
}
