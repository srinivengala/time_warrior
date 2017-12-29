package cmd

import (
	"fmt"

	"github.com/mrcook/time_warrior/manager"
	"github.com/mrcook/time_warrior/timeslip"
	"github.com/spf13/cobra"
)

var resumeCmd = &cobra.Command{
	Use:     "resume",
	Short:   "Resume a paused timeslip",
	Aliases: []string{"r"},
	Args:    cobra.NoArgs,
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		slip, err := resumeTimeSlip()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(slip)
		}
	},
}

func init() {
	rootCmd.AddCommand(resumeCmd)
}

func resumeTimeSlip() (*timeslip.Slip, error) {
	m := manager.NewFromConfig(initializeConfig())

	slipJSON, err := m.PendingTimeSlip()
	if err != nil {
		return nil, err
	}

	slip, err := timeslip.NewFromJSON(slipJSON)
	if err != nil {
		return nil, err
	}

	if err := slip.Resume(); err != nil {
		return nil, err
	}

	if err := m.SavePending(slip.ToJson()); err != nil {
		return nil, err
	}

	return slip, nil
}
