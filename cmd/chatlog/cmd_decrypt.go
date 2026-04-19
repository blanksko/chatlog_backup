package chatlog

import (
	"fmt"

	"github.com/sjzar/chatlog/internal/chatlog"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(decryptCmd)
	decryptCmd.Flags().StringVarP(&decryptPlatform, "platform", "p", "wechat", "platform (e.g. wechat, qq)")
	decryptCmd.Flags().IntVarP(&decryptVer, "version", "v", 0, "version")
	decryptCmd.Flags().StringVarP(&decryptDataDir, "data-dir", "d", "", "data dir")
	decryptCmd.Flags().StringVarP(&decryptDatakey, "data-key", "k", "", "data key")
	decryptCmd.Flags().StringVarP(&decryptWorkDir, "work-dir", "w", "", "work dir")
}

var (
	decryptPlatform string
	decryptVer      int
	decryptDataDir  string
	decryptDatakey  string
	decryptWorkDir  string
)

var decryptCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "decrypt chat log database",
	Long:  "Decrypt an encrypted chat log database using the provided data key.",
	Run: func(cmd *cobra.Command, args []string) {

		// Validate that a data key was provided before attempting decryption
		if len(decryptDatakey) == 0 {
			log.Error().Msg("data key is required: use -k to provide the data key")
			return
		}

		cmdConf := getDecryptConfig()

		m := chatlog.New()
		if err := m.CommandDecrypt("", cmdConf); err != nil {
			log.Err(err).Msg("failed to decrypt")
			return
		}
		// Print the platform that was used so it's easier to confirm the right one was selected
		fmt.Printf("decrypt success (platform: %s)\n", decryptPlatform)
	},
}

func getDecryptConfig() map[string]any {
	cmdConf := make(map[string]any)
	if len(decryptDataDir) != 0 {
		cmdConf["data_dir"] = decryptDataDir
	}
	if len(decryptDatakey) != 0 {
		cmdConf["data_key"] = decryptDatakey
	}
	if len(decryptWorkDir) != 0 {
		cmdConf["work_dir"] = decryptWorkDir
	}
	if len(decryptPlatform) != 0 {
		cmdConf["platform"] = decryptPlatform
	}
	if decryptVer != 0 {
		cmdConf["version"] = decryptVer
	}
	return cmdConf
}
