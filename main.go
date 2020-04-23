package main

import (
	"github.com/medmedchiheb/helm2tohelm3/backup"
	"github.com/medmedchiheb/helm2tohelm3/migrate"
	"github.com/medmedchiheb/helm2tohelm3/utils"
	"github.com/sirupsen/logrus"
	flag "github.com/spf13/pflag"
	"strings"

	"os"
	"os/exec"
)

var (
	cluster = flag.String("context", "minikube", "The target cluster to run the migrate on.")
	restore = flag.Bool("restore", false, "Restore Helm2 releases")
	actions = flag.String("actions", "move-convert-cleanup", "The Actions you want to execute: 'move' (configuration), 'convert' (releases), 'cleanup'. separate actions by '-'")
	backupDir = flag.String("backup-dir", "", "The path for Backup files")
	reset = flag.Bool("reset", false, "Reset the created Helm3 releases, this will revoke managed helm3 releases, execute this only when the helm2 still manage the releases.")
)

func main() {
	flag.Parse()
    run()
}



func run()  {
	logrus.Infof("Running Helm2toHelm3 migration on: %s", *cluster)
	validate()
	changeK8sContext(*cluster)

	if *restore {
		if utils.YesNo("Are you sure want to restore the data") {
			backup.Restore(*cluster, *backupDir)
		}
	} else if *reset {
		migrate.Reset()
	} else {
		if strings.Contains(*actions, "convert") {
			backup.ExecuteBackup(*cluster, *backupDir)
		}

		if strings.Contains(*actions, "move") {
			migrate.MigrateConfiguration()
		}

		if strings.Contains(*actions, "convert") {
			migrate.MigrateReleases()
		}

		if strings.Contains(*actions, "cleanup") {
			migrate.Cleanup()
		}

		logrus.Infof("Job Done: please check 'helm3 ls' ")
	}

}



func validate() {
	s := *backupDir
	if len(s) > 0 && s[len(s)-1:] != "/" {
		*backupDir = *backupDir+"/"
	}
}



func changeK8sContext(c string) {
	command := "kubectl config use-context " + c

	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		logrus.Errorf("Error: ", err)
	}
}