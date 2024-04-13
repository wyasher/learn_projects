package command

import "web-framework/framework/cobra"

func AddKernelCommand(root *cobra.Command) {
	root.AddCommand(DemoCommand)
	root.AddCommand(initAppCommand())
}
