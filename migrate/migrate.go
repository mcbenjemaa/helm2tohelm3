package migrate

import (
	"fmt"
	"github.com/medmedchiheb/helm2tohelm3/exec"
	"github.com/medmedchiheb/helm2tohelm3/utils"
	"github.com/sirupsen/logrus"
	"strings"
)

func MigrateConfiguration() {

	logrus.Infof("------ > Migrate Helm v2 configuration\n\n")
	c := exec.Command{
		Command: "echo y | helm3 2to3 move config",
		Dryrun:  true,
		Execute: true,
	}

	logrus.Infof("Executing '%s' DryRun mode", c.Command)
	err := c.Exec()
	if err != nil {
		logrus.Errorf("Could not execute the script %v:", err)
	}

	c.DryRunMode(false)
	logrus.Infof("Executing '%s'", c.Command)
	if utils.YesNo("Are you sure you want to Migrate Helm v2 configuration") {
		err = c.Exec()
		if err != nil {
			logrus.Errorf("Could not execute the script %v:", err)
		}
	}
}



func MigrateReleases() {
	logrus.Infof("------> Migrate Helm v2 releases\n\n")

	s := getReleases()
	logrus.Infof("Helm releases: %v\n", s)

    logrus.Infof("Migrating Release DryRun mode: %v", s[0])
	logrus.Infof("If something wrong Please stop \n")

	migrateRelease(s[0], true)

	logrus.Infof("Migrating All Releases! \n\n")
	if utils.YesNo("Will migrate all releases to helm 3 without asking for confirmation, Are yu sure!") {
		for _, r := range s {
			err := migrateRelease(r, false)
			if err != nil {
				logrus.Warn("%s Release could not be migrated: %v", r, err)
			}
			logrus.Infof("%s migrated successfully to Helm3", r)
		}
	}

}



func Cleanup() {
	logrus.Infof("------> Clean up of Helm v2 data\n\n")
	c := exec.Command{
		Command: "echo y | helm3 2to3 cleanup ",
		Dryrun:  true,
		Execute: true,
	}

	logrus.Infof("Executing '%s' DryRun mode", c.Command)
	err := c.Exec()
	if err != nil {
		logrus.Errorf("Could not execute the script %v:", err)
	}

	c.DryRunMode(false)
	logrus.Infof("Executing '%s'", c.Command)
	if utils.YesNo("Are you sure you want to Cleanup Helm v2 data") {
		err = c.Exec()
		if err != nil {
			logrus.Errorf("Could not execute the script %v:", err)
		}
	}
}



func getReleases() []string {
	c := exec.Command{
		Command: "helm ls -q",
	}
	out, err := c.ExecAndGetOutput()
	if err != nil {
		logrus.Errorf("could not execute script %v", err)
	}

	s := strings.Fields(out)

   return s
}


func migrateRelease(r string, b bool)  error {
	c := exec.Command{
		Command: "helm3 2to3 convert " + r,
		Execute: true,
	}
	c.DryRunMode(b)
	logrus.Infof("Executing '%s'", c.Command)
	err := c.Exec()
	if err != nil {
		return fmt.Errorf("Could not execute the script %v:", err)
	}

	return err
}


func Reset() {
	logrus.Infof("------> Reset of Helm v3 data\n\n")
	c := exec.Command{
		Command: "kubectl delete secrets --all-namespaces --selector \"owner=helm\"",
		Execute: true,
	}

	c.DryRunMode(false)
	logrus.Infof("Executing '%s'", c.Command)
	if utils.YesNo("Are you sure you want to reset Helm v3 data") {
		err := c.Exec()
		if err != nil {
			logrus.Errorf("Could not execute the script %v:", err)
		}
	}
}