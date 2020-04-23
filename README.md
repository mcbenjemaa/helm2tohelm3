# Helm2ToHelm3

This small go app to help you migrate from helm2 to helm3,
it uses the nice plugin: `helm-2to3`.

The purpose for this app is only for ensuring automation.
The migration guide: https://helm.sh/blog/migrate-from-helm-v2-to-helm-v3/


### Prerequisite

- Kubernetes cluster (kubectl)

- Helm 2 already has Tiller

- Install Helm 3 (make sure to rename the binary to `helm3` and store it in your path.)

- helm-2to3 plugin: `helm3 plugin install https://github.com/helm/helm-2to3`


### Usage:

The script will run those jobs: 
- backup helm2 data to file.
- Move configuration to Helm3
- Migrate Helm releases to Helm3
- Cleanup.


The script will run the dry-run first then prompt you to select [Yes/No] for any given action, but please be aware that it will run Dryrun for the first Release only when migrating Releases,
and then after the confirmation of the prompt, it will run all migration without dry-run and without confirmation.

If you notice something wrong within the dry-run you can stop the script and choose NO.

The script will create Backup files in the provided location, by default it's the current directoy: e.g. `backup/output-files`. see [output files](backup/output-files/README.md)

### Build and run.

To build the application:

```
go install
```

Run the script:
```
helm2tohelm3 --context minikube --backup-dir somedir/
```


#### Available options:

| Cli Flag       | Description | Example |
| -------------- | -------------- |:-----------:|
| `--context`    | The target cluster to run the migrate on. (default "mcs-eu-dev-dom") | `helm2tohelm3 --context minikube`
| `--restore`    | After creating helm2 Backup you can restore the data the the cluster.  | `helm2tohelm3 --context mcs-eu-dev-prof --restore`
| `--actions`    | To ignore same actions then this flag is good for you: default to `move-convert-cleanup` (please use the separator between actions `-`)  | `helm2tohelm3 --actions convert-cleanup`
| `--backup-dir` | The directory where you want to store the backup files, or to load from | `helm2tohelm3 --context mcs-eu-dev-prof --backup-dir backup/output-files`
| `--reset`      | Reset the created Helm3 releases, this will revoke managed helm3 releases, execute this only when the helm2 still manage the releases.  | `helm2tohelm3 --context mcs-eu-dev-dom --reset`