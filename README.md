# seagate-exos-x-api-go

[![Go Report Card](https://goreportcard.com/badge/github.com/Seagate/seagate-exos-x-api-go)](https://goreportcard.com/report/github.com/Seagate/seagate-exos-x-api-go)
[![Go Reference](https://pkg.go.dev/badge/github.com/Seagate/seagate-exos-x-api-go.svg)](https://pkg.go.dev/github.com/Seagate/seagate-exos-x-api-go)


# MC API Client 2.0

This is a Go language library for interfacing with Seagate's Management Controller (MC) CLI/API. This library is used to generate an MC OpenAPI specification and then use that specification to generate a Go CLient API library. The OpenAPI specification is generated using an input yaml file and then gathering `meta` data from the MC API to create the OpenAPI specification. This library leverages an HTTP interface to communicate with the Management Controller of a live Seagate RAID Storage System. In addition, a higher-level API layer is used by clients such as the CSI Driver, and that higher layer relies on the auto-generated Go Client library.

The [Seagate EXOS X API](https://www.seagate.com/files/www-content/support-content/raid-systems/_shared/documentation/83-00007047-13-01_G265_SMG.pdf).

The HTTP API operations are used to execute provisioning and management commands.

## MC API Design

This library was designed to generate an MC OpenAPI specification, as well as provide a library that clients can use to communicate with the MC in a Seagate storage array.

For most users of this library, you will only need to access the generated MC API Go library. The original library (v1) was written by hand and included support for a handful of commands used by the Kubernetes CSI Driver. The current library (v2) uses openapi tools to auto-generated a Go Language Client library. This library uses the **MC OpenAPI Specification** (`api/mc-openapi.yaml`) to auto-generate the Go client library. That same spec can be used to generate a Python, JavaScript, or other libraries.

This design uses a layered approach to the client API - where clients can call the `pkg/client` functions to communicate with an MC, or use the `pkg/v1` or `pkg/v2` client API functions. The `pkg/v1` was created to support the original CSI driver and uses the hand-written code which parses XML responses from the MC. The `pkg/v2` client API was also hand-written and leverages the auto-generated `pkg/client` functions. The v2 client parses JSON data returned from the MC and is 4-5 times faster than v1 based on regression test suite timings. Having both a v1 and v2 versions allowed for an easy transition in any client from importing `github.com/Seagate/seagate-exos-x-api-go/pkg/v1` to `github.com/Seagate/seagate-exos-x-api-go/pkg/v2` without having to change function calls.

This library also includes code written to auto-generate a new MC OpenAPI specification, but that step should not be required in most cases. Below are notes regarding current support, a road map, and regression testing notes.

## OpenAPI Implementation

This section defines the commands, response status codes, and resource objects implemented in the OpenAPI specification. This is a superset of what was developed in API v1.
Commands were added that are being used by other clients needing to interface with the MC Storage Controller API.

The following table provides a list of all possible MC API commands and which ones are supported currently in the MC OpenAPI spec and Go Client library. There are a total of 291 commands, with 29 supported, roughly 10%.

| Command               | In MC OpenAPI Spec  | Not Supported Yet |
| :-------------------- | :------------------ | :---------------- |
| `login` | supported | |
| `abort copy` |  | not supported |
| `abort replication` |  | not supported |
| `abort scrub` |  | not supported |
| `activate certificate` |  | not supported |
| `activate disk-firmware` |  | not supported |
| `activate firmware` |  | not supported |
| `add disk-group` |  | not supported |
| `add event-subscriber` |  | not supported |
| `add host-group-members` |  | not supported |
| `add host-members` |  | not supported |
| `add ipv6-address` |  | not supported |
| `add spares` |  | not supported |
| `add storage` |  | not supported |
| `add volume-group-members` |  | not supported |
| `check firmware-upgrade-health` |  | not supported |
| `clear cache` |  | not supported |
| `clear disk-metadata` |  | not supported |
| `clear dns-parameters` |  | not supported |
| `clear events` |  | not supported |
| `clear expander-status` |  | not supported |
| `clear failed-disk` |  | not supported |
| `clear fde-keys` |  | not supported |
| `clear replication-queue` |  | not supported |
| `clear slow-disk` |  | not supported |
| `create certificate-signing-request` |  | not supported |
| `create certificate` |  | not supported |
| `create custom-profile` |  | not supported |
| `create host-group` |  | not supported |
| `create host` |  | not supported |
| `create peer-connection` |  | not supported |
| `create remote-system` |  | not supported |
| `create replication-set` |  | not supported |
| `create schedule` |  | not supported |
| `create snapshots` | supported |  |
| `create snmpdConf` |  | not supported |
| `create user-group` |  | not supported |
| `create volume-group` |  | not supported |
| `create volume-set` |  | not supported |
| `create volume` | supported |  |
| `delete all-snapshots` |  | not supported |
| `delete chap-record` |  | not supported |
| `delete disk-firmware` |  | not supported |
| `delete event-subscriber` |  | not supported |
| `delete host-groups` |  | not supported |
| `delete hosts` | supported |  |
| `delete initiator-nickname` | supported |  |
| `delete peer-connection` |  | not supported |
| `delete pools` |  | not supported |
| `delete remote-system` |  | not supported |
| `delete replication-set` |  | not supported |
| `delete schedule` |  | not supported |
| `delete snapshot` | supported |  |
| `delete task` |  | not supported |
| `delete user-group` |  | not supported |
| `delete user` |  | not supported |
| `delete volume-groups` |  | not supported |
| `delete volumes` | supported |  |
| `dequarantine` |  | not supported |
| `down disk` |  | not supported |
| `erase disk` |  | not supported |
| `exit` |  | not supported |
| `expand disk-group` |  | not supported |
| `expand volume` | supported |  |
| `fail` |  | not supported |
| `help` |  | not supported |
| `map volume` | supported |  |
| `meta` | supported |  |
| `ping` |  | not supported |
| `query metrics` |  | not supported |
| `query peer-connection` |  | not supported |
| `recover disk-group` |  | not supported |
| `recover replication-set` |  | not supported |
| `recover volume` |  | not supported |
| `release volume` |  | not supported |
| `remove certificate` |  | not supported |
| `remove disk-groups` |  | not supported |
| `remove host-group-members` |  | not supported |
| `remove host-members` |  | not supported |
| `remove ipv6-address` |  | not supported |
| `remove mas-service` |  | not supported |
| `remove spares` |  | not supported |
| `remove volume-group-members` |  | not supported |
| `replicate` |  | not supported |
| `rescan` |  | not supported |
| `reset all-statistics` |  | not supported |
| `reset ciphers` |  | not supported |
| `reset controller-statistics` |  | not supported |
| `reset disk-error-statistics` |  | not supported |
| `reset disk-group-statistics` |  | not supported |
| `reset disk-statistics` |  | not supported |
| `reset dns-management-hostname` |  | not supported |
| `reset host-link` |  | not supported |
| `reset host-port-statistics` |  | not supported |
| `reset pool-statistics` |  | not supported |
| `reset smis-configuration` |  | not supported |
| `reset snapshot` |  | not supported |
| `reset volume-statistics` |  | not supported |
| `restart controller` |  | not supported |
| `restart disk` |  | not supported |
| `restart mc` |  | not supported | 
| `restart sc` |  | not supported |
| `restore defaults` |  | not supported |
| `resume replication-set` |  | not supported |
| `rollback volume` |  | not supported |
| `scrub disk-groups` |  | not supported |
| `scrub volume` |  | not supported |
| `set advanced-settings` |  | not supported |
| `set alert` |  | not supported |
| `set chap-record` |  | not supported |
| `set ciphers` |  | not supported |
| `set cli-parameters` |  | not supported |
| `set controller-date` |  | not supported |
| `set custom-profile` |  | not supported |
| `set debug-log-parameters` |  | not supported |
| `set degraded-disk-threshold` |  | not supported |
| `set disk-firmware-update-settings` |  | not supported |
| `set disk-group` |  | not supported |
| `set disk-parameters` |  | not supported |
| `set disk` |  | not supported |
| `set dns-management-hostname` |  | not supported |
| `set dns-parameters` |  | not supported |
| `set email-parameters` |  | not supported |
| `set enclosure` |  | not supported |
| `set event-subscriber` |  | not supported |
| `set expander-fault-isolation` |  | not supported |
| `set expander-phy` |  | not supported |
| `set fde-import-key` |  | not supported |
| `set fde-state` |  | not supported |
| `set host-group` |  | not supported |
| `set host-parameters` |  | not supported |
| `set host-port-mode` |  | not supported |
| `set host` |  | not supported |
| `set infiniband-parameters` |  | not supported |
| `set initiator` | supported |  |
| `set ipv6-address` |  | not supported |
| `set ipv6-network-parameters` |  | not supported |
| `set ldap-parameters` |  | not supported |
| `set led` |  | not supported |
| `set linear` |  | not supported |
| `set mas-service` |  | not supported |
| `set network-parameters` |  | not supported |
| `set ntp-parameters` |  | not supported |
| `set password` |  | not supported |
| `set peer-connection` |  | not supported |
| `set pool` |  | not supported |
| `set prompt` |  | not supported |
| `set protocols` |  | not supported |
| `set proxy-server` |  | not supported |
| `set redfish-persistent-data` |  | not supported |
| `set remote-system` |  | not supported |
| `set replication-set` |  | not supported |
| `set schedule` |  | not supported |
| `set slow-disk-policy` |  | not supported |
| `set slow-disk-threshold` |  | not supported |
| `set snapshot-space` |  | not supported |
| `set snmp-parameters` |  | not supported |
| `set stratus-hybrid` |  | not supported |
| `set support-assist-ese` |  | not supported |
| `set syslog-parameters` |  | not supported |
| `set system` |  | not supported |
| `set task` |  | not supported |
| `set update-server` |  | not supported |
| `set user-group` |  | not supported |
| `set user` |  | not supported |
| `set volume-group` |  | not supported |
| `set volume` |  | not supported |
| `show advanced-settings` | supported |  |
| `show alerts` |  | not supported |
| `show audit-log` |  | not supported |
| `show cache-parameters` | supported |  |
| `show certificates` | supported |  |
| `show chap-records` |  | not supported |
| `show ciphers` |  | not supported |
| `show cli-parameters` |  | not supported |
| `show configuration` |  | not supported |
| `show controller-date` |  | not supported |
| `show controller-statistics` |  | not supported |
| `show controllers` | supported |  |
| `show debug-log-parameters` |  | not supported |
| `show debug-log` |  | not supported |
| `show degraded-disk-policy` |  | not supported |
| `show degraded-disk-thresholds` |  | not supported |
| `show disk-affinity` |  | not supported |
| `show disk-firmware-update-settings` |  | not supported |
| `show disk-group-statistics` |  | not supported |
| `show disk-groups` | supported |  |
| `show disk-parameters` |  | not supported |
| `show disk-statistics` |  | not supported |
| `show disks` | supported |  |
| `show dns-management-hostname` |  | not supported |
| `show dns-parameters` |  | not supported |
| `show email-parameters` |  | not supported |
| `show enclosures` | supported |  |
| `show event-subscriber` |  | not supported |
| `show events` |  | not supported |
| `show expander-status` |  | not supported |
| `show factory-default-status` |  | not supported |
| `show fan-modules` |  | not supported |
| `show fans` | supported |  |
| `show fde-state` |  | not supported |
| `show features` |  | not supported |
| `show fenced-data` |  | not supported |
| `show firmware-bundles` |  | not supported |
| `show firmware-update-status` |  | not supported |
| `show frus` |  | not supported |
| `show host-groups` | supported |  |
| `show host-phy-statistics` |  | not supported |
| `show host-port-statistics` |  | not supported |
| `show infiniband-parameters` |  | not supported |
| `show initiators` | supported |  |
| `show inquiry` |  | not supported |
| `show ipv6-addresses` |  | not supported |
| `show ipv6-network-parameters` |  | not supported |
| `show iscsi-parameters` |  | not supported |
| `show kms-parameters` |  | not supported |
| `show ldap-parameters` |  | not supported |
| `show led-states` |  | not supported |
| `show license` |  | not supported |
| `show maps` | supported |  |
| `show mas-service` |  | not supported |
| `show metrics-list` |  | not supported |
| `show mui` |  | not supported |
| `show network-parameters` |  | not supported |
| `show ntp-status` |  | not supported |
| `show peer-connections` |  | not supported |
| `show pool-statistics` |  | not supported |
| `show pools` | supported |  |
| `show ports` |  | not supported |
| `show power-supplies` | supported |  |
| `show profile` |  | not supported |
| `show profiles` |  | not supported |
| `show protocols` |  | not supported |
| `show provisioning` |  | not supported |
| `show proxy-server` |  | not supported |
| `show redfish-persistent-data` |  | not supported |
| `show redundancy-mode` |  | not supported |
| `show refresh-counters` |  | not supported |
| `show remote-systems` |  | not supported |
| `show replication-sets` |  | not supported |
| `show replication-snapshot-history` |  | not supported |
| `show sas-link-health` |  | not supported |
| `show scem` |  | not supported |
| `show schedules` |  | not supported |
| `show sensor-status` |  | not supported |
| `show service-tag-info` |  | not supported |
| `show sessions` |  | not supported |
| `show shutdown-status` |  | not supported |
| `show slow-disk-policy` |  | not supported |
| `show slow-disk-thresholds` |  | not supported |
| `show smis-debug` |  | not supported |
| `show snapshot-space` |  | not supported |
| `show snapshots` | supported |  |
| `show snmp-parameters` |  | not supported |
| `show spear-errors` |  | not supported |
| `show storage-configurations` |  | not supported |
| `show support-assist-contact` |  | not supported |
| `show support-assist-ese` |  | not supported |
| `show support-assist-telemetry-status` |  | not supported |
| `show support-assist` |  | not supported |
| `show syslog-parameters` |  | not supported |
| `show system-parameters` |  | not supported |
| `show system` | supported |  |
| `show tasks` |  | not supported |
| `show tier-statistics` |  | not supported |
| `show tiers` |  | not supported |
| `show timezones` |  | not supported |
| `show unwritable-cache` |  | not supported |
| `show update-server` |  | not supported |
| `show user-groups` |  | not supported |
| `show user-interactions` |  | not supported |
| `show users` |  | not supported |
| `show versions` | supported |  |
| `show volume-copies` |  | not supported |
| `show volume-labels` |  | not supported |
| `show volume-names` |  | not supported |
| `show volume-reservations` |  | not supported |
| `show volume-statistics` |  | not supported |
| `show volumes` | supported |  |
| `show workload` |  | not supported |
| `shutdown` |  | not supported |
| `simulate fru-fault` |  | not supported |
| `start metrics` |  | not supported |
| `stop metrics` |  | not supported |
| `suspend replication-set` |  | not supported |
| `trust` |  | not supported |
| `unfail controller` |  | not supported |
| `unfence controller` |  | not supported |
| `uninstall license` |  | not supported |
| `unmap volume` | supported |  |
| `verify disk-performance` |  | not supported |


## Generation Steps

| Command                             | Notes |
| :---------------------------------- | :------------------------------------------------------------------------ |
| Update `generator/mc-commands.yaml` | Modify this YAML file which controls MC command generation |
| `make generator`                    | Builds a generator executable for creating the OpenAPI YAML file |
| `make rung`                         | Executes: `go run generator/cmd/main.go -config generator/generator.conf` |
| `make validate`                     | Use `openapi-generator-cli` to validate the MC OpenAPI spec |
| `make generate`                     | Use `openapi-generator-cli` to generate a Go MC client library |
| `make validator`                    | Builds a validator executable for testing the generated client library |
| `make runv`                         | Executes: `go run validator/cmd/main.go -config validator/validator.conf`

**Note:** `validator/cmd/main.go` is used to validate a handful of MC client library commands, extend as desired.

# MC OpenAPI Roadmap

This section describes proposed next steps for implementing remaining commands.
- Create a new MC API command that returns a list of all supported commands along with their parameters and other information needed to auto-generate `mc-commands.yaml`.
- Expand MC OpenAPI support to include all commands.
- Expand MC OpenAPI support to include all command parameters and variations.
- Work on a system of having multiple versions of the MC OpenAPI spec corresponding to various product releases.

# Steps Needed to Update the MC OpenAPI Specification

In order to add a new command that is supported by the MC OpenAPI specification and also supported by the Go Client Library, follow these steps:
- Update `internal/generator/mc-commands.yaml` to include the new command and save your changes:
  - The main sections to include are: `command URI`, `meta`, `include`, `nested`, and `options`.
  - The `command URI` is essentially how you would run a command in a shell. For example, `show volumes` becomes `/show/volumes`.
  - The `meta` tag value is displayed when you run a command and have it return JSON data, it is one of the properties in the JSON dictionary.
  - The `include` tagn almost always contains `status`.
  - The `nested` tag describes all other JSON dictionaries that are nested, and this can be recursive. Review the JSON data returned for a command.
  - The `options` tag describes all possible command arguments, their types, required or not, keyword required or not, and an optional description.
- In a terminal, run `make generator`
- In an editor, copy `internal/generator/generator-example.conf` to `internal/generator/generator.conf` and modify as needed. This conf file specifies which storage array use for generation.
- In a terminal, run `make rung` to create a new `api/mc-openapi.yaml`
- **Note:** The next two commands require `docker` to be installed and working.
- In a terminal, run `make validate` to validate that `api/mc-openapi.yaml` contains no OpenAPI errors
- In a terminal, run `make generate` to generate a new Go client library in `pkg/client`
- **Note:** You may need to update a file or create a new file in `pkg/v2` that leverages the new command generated. This isn't required, but the v2 layer was created to provide a simplified layer between the CSI Driver and the MC API. 
- The last step is to run the regression test suite for v1 and/or v2.

# Regression Testing Using A Live System

This option runs a API Regression test suite using Ginkgo expressive tests.

This option requires a configuration file, for example: `api-regression.conf`.

Each user should create their own copy of `api-regression.conf` and not check it in. Please see `api-regression-example.conf` as an example.
Use the `./api-regression -config <filename>` option to specify your configuration file.

```
#
# api-regression configuration
#
# Properties:
#   'ip' is required and is the IP Address of the storage controller used for testing
#   'protocol" specifies http or https
#   'username' is required and is the login credentials for the storage controller
#   'password' is required and is the login credentials for the storage controller
#
#   'initiator' is an array of host initiators, such as the iSCSI IQN value, or SAS, or FC
#   'pool' is the storage pool used for creating volumes
#
# Notes:
#    To find the iSCSI initiator:
#    sudo cat /etc/iscsi/initiatorname.iscsi | grep -v "##" | awk -F= '{print $2}'
api-regression: 1.0.0

# Example Controller
ip: "<ipaddress>"
protocol: "https"
username: "<username>"
password: "<password>"
initiator: ["iqn.2004-10.com.ubuntu:01:b6f76364a18"]
pool: "A"
```

The suggestion is to create your own configuration file and do not check it into the repo. For example:
`cp api-regression.conf myconfig.conf` and then run `./api-regression -debug 4 -config myconfig.conf --ginkgo.v`.

There are many options for running Ginkgo test cases, here are a few using an `i1.conf` configuration file:

- `make regession` to build the ***api-regression*** executable.
- `./api-regression -config i1.conf -debug 4 --ginkgo.v --ginkgo.fail-fast --ginkgo.focus "v1"` to run all the ***v1*** tests
- `./api-regression -config i1.conf -debug 4 --ginkgo.v --ginkgo.fail-fast --ginkgo.focus "v2"` to run all the ***v2*** tests
- `./api-regression -config i1.conf -debug 4 --ginkgo.v --ginkgo.fail-fast --ginkgo.focus "v2Login"` to run only the ***v2 Login*** tests
- `./api-regression -config i1.conf -debug 4 --ginkgo.v --ginkgo.fail-fast --ginkgo.focus "v2System"` to run only the ***v2 System*** tests
- `./api-regression -config i1.conf -debug 4 --ginkgo.v --ginkgo.fail-fast --ginkgo.focus "v2Volume"` to run only ***v2 Volume*** tests
- `./api-regression -config i1.conf -debug 4 --ginkgo.v --ginkgo.fail-fast --ginkgo.focus "v2Snapshot"` to run only ***v2 Snapshot*** tests
- `./api-regression -config i1.conf -debug 4 --ginkgo.v --ginkgo.fail-fast` to run only all tests

There are also many command line flags:
- `./api-regression --help` to see all options
- `-debug level` to set the debug level to 0, 1, 2, 3, or 4
- `-config filename` to use an alternate configuration file
- `-version` to display the ***api-regression*** version and exit
- `--ginkgo.fail-fast` to stop testing after the first failure
- `--ginkgo.v` to run testing in verbose mode
