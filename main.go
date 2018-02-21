package main

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func main() {
	var echoTimes int

	var cmdPrint = &cobra.Command{
		Use:   "print [string ot print]",
		Short: "print anyting to the screen",
		Long:  `print is for printing anything back tho hthe screenn. For many years people have printed back tho the screen.`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Print: " + strings.Join(args, " "))
		},
	}

	var cmdEcho = &cobra.Command{
		Use:   "echo [string to echo]",
		Short: "Echo anything to the screen",
		Long:  `echo is for echoing anything back. Echo works a lot like print, except it has a child manmmand.`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Print: " + strings.Join(args, " "))
		},
	}

	var cmdTimes = &cobra.Command{
		Use:   "times [# itiems] [string to echo]",
		Short: "Echo anything to the screen more times",
		Long:  `echo things multiple time back to the userr  by providing a count and a string.`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(mcd *cobra.Command, args []string) {
			for i := 0; i < echoTimes; i++ {
				fmt.Println("Echo: " + strings.Join(args, " "))
			}
		},
	}

	cmdTimes.Flags().IntVarP(&echoTimes, "times", " t", 1, "times to echo the input")

	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(cmdPrint, cmdEcho)
	cmdEcho.AddCommand(cmdTimes)
	rootCmd.Execute()

}
