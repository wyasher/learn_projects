package console

import (
	"web-framework/app/console/comman/demo"
	"web-framework/framework"
	"web-framework/framework/cobra"
	"web-framework/framework/command"
)

func RunCommand(container framework.Container) error {
	var rootCmd = &cobra.Command{
		Use:   "web",
		Short: "web 命令",
		Long:  `web 框架提供的命令行工具，使用这个命令行工具能很方便执行框架自带命令，也能很方便编写业务命令`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.InitDefaultHelpFlag()
			return cmd.Help()
		},
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
	}
	rootCmd.SetContainer(container)
	command.AddKernelCommand(rootCmd)
	AddAppCommand(rootCmd)
	return rootCmd.Execute()
}
func AddAppCommand(rootCmd *cobra.Command) {
	rootCmd.AddCommand(demo.InitFoo())
}
