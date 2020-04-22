package backup

import (
	"fmt"
	"github.com/medmedchiheb/helm2tohelm3/exec"
	"github.com/medmedchiheb/helm2tohelm3/utils"
	"github.com/sirupsen/logrus"
	"os"
)

func ExecuteBackup(c string, path string) {
	if checkBackupFileExists(c, path) {
		logrus.Warnf("Backup file already exists '%s%s-backup-cm.yaml'. Please make sure to execute backup again or not!", path, c)
	}
	if utils.YesNo("Are you sure to create Backup!"){
		command := fmt.Sprintf("kubectl get configmaps --namespace \"kube-system\" --selector \"OWNER=TILLER\" --output \"yaml\" > %s%s-backup-cm.yaml", path, c)

		cmd := exec.Command{
			Command: command,
			Dryrun:  false,
			Execute: true,
		}

		err := cmd.Exec()
		if err != nil {
			logrus.Fatalf("Could not execute the script %v:", err)
		}

		logrus.Infof("Backup %s-backup-cm.yaml saved!", c)
	} else {
		logrus.Warnf("Skipping: Backup")
	}


}

func Restore(c string,  path string)  {

	command := fmt.Sprintf("kubectl apply -f  %s%s-backup-cm.yaml", path, c)

	cmd := exec.Command{
		Command: command,
		Dryrun:  false,
		Execute: true,
	}

	err := cmd.Exec()
	if err != nil {
		logrus.Fatalf("Could not execute the script %v:", err)
	}

	logrus.Infof("The Helm2 Releases are restored to the cluster %s", c)
}


func checkBackupFileExists(c string, path string) bool {
	file := fmt.Sprintf("%s%s-backup-cm.yaml", c)

	if _, err := os.Stat(file); err == nil {
		return true
	} else {
		return false
	}
}