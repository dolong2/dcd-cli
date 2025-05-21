package cmd

import (
	"github.com/dolong2/dcd-cli/api/exec"
	"github.com/spf13/cobra"
)

// profileCmd represents the profile command
var profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "현재 프로필 조회",
	Long:  `로그인된 유저 정보의 프로필을 조회합니다.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		profile, err := exec.GetProfile()
		if err != nil {
			return err
		}

		cmd.Println("Email: ", profile.User.Email)
		cmd.Println("Name:", profile.User.Name)
		cmd.Println("Status: ", profile.User.Status)
		cmd.Println()

		for _, ws := range profile.WorkspaceList {
			cmd.Printf("[Workspace] %s\n", ws.Title)
			for i, app := range ws.ApplicationList {
				prefix := "└─"
				if len(ws.ApplicationList) > 1 && i < len(ws.ApplicationList)-1 {
					prefix = "├─"
				}

				cmd.Printf("  %s [Application]\n", prefix)
				cmd.Printf("  │   • ID: %s\n", app.Id)
				cmd.Printf("  │   • Name: %s\n", app.Name)
				cmd.Printf("  │   • Description: %s\n", app.Description)
			}
			cmd.Println()
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(profileCmd)
}
