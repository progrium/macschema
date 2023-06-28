package cmd

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/progrium/macschema/schema"
	"github.com/spf13/cobra"
)

var crawlCmd = &cobra.Command{
	Use:   "crawl",
	Short: "Downloads topics linked from a topic to doc dir",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		l := schema.NewLookup(args[0], flagLang)
		opts := schema.FetchOptions{}
		opts.Debug, _ = cmd.Flags().GetBool("debug")
		ctx := context.Background()
		if !l.DocExists() {
			schema.FetchTopic(ctx, l, opts)
		}
		t, err := schema.ReadTopic(l)
		fatal(err)

		for _, link := range t.Topics {
			fmt.Fprintln(os.Stderr, "=>", link.Name)
			ll := schema.LookupFromPath(link.Path)
			if ll.DocExists() {
				// TODO: check last fetch, version
				continue
			}
			tt := schema.FetchTopic(ctx, ll, fetchOptions(cmd))
			fatal(writeTopic(ll, tt))
			fmt.Fprintf(os.Stderr, "   %s [%s]\n", ll.DocPath, time.Since(tt.LastFetch))
		}
	},
}
