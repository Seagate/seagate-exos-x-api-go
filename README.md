# seagate-exos-x-api-go

[![Go Report Card](https://goreportcard.com/badge/github.com/Seagate/seagate-exos-x-api-go)](https://goreportcard.com/report/github.com/Seagate/seagate-exos-x-api-go)
[![Go Reference](https://pkg.go.dev/badge/github.com/Seagate/seagate-exos-x-api-go.svg)](https://pkg.go.dev/github.com/Seagate/seagate-exos-x-api-go)


# MC API Client 2.0

This is a Go language library for interfacing with Seagate's Management Controller (MC) CLI/API. This library is used to generate an MC OpenAPI specification
and then to use that specification to generate a Go CLient API library. The OpenAPI specification is generated using an input yaml file and then gathering `meta` data
from the MC API to create the OpenAPI specification. This library leverages an HTTP interface to communicate with the Management Controller of a live Seagate RAID Storage System.

The [Seagate EXOS X API](https://www.seagate.com/files/www-content/support-content/raid-systems/_shared/documentation/83-00007047-13-01_G265_SMG.pdf).

The HTTP API operations are used to execute provisioning and management commands.

## OpenAPI Implementation

This section defines the commands, response status codes, and resource objects implemented in the OpenAPI specification.

The `api/mc-openapi.yaml` spec includes support for the following items.

| Command               | Status        | Notes |
| :-------------------- | :------------ | :------------ |
| `login` | Implemented | Required to generate OpenAPI specification |
| `meta` | Implemented | Required to generate OpenAPI specification |
| | | |
| `show system` |  Implemented | includes `status` and nested `redundancy` resource
| `show fans` | Implemented | |
| `show volumes` | Implemented | |
| `show power-supplies` | Implemented | |
| `show certificate` | Implemented | |
| `show pools` | Implemented | |
| `show disks` | Implemented | |
| `show advanced-settings` | Implemented | |
| `show cache-parameters` | Implemented | |
| `show host-groups` | Implemented | |
| `show enclosures` | Implemented | |
| `show controllers` | Implemented | |
| `show versions` | Implemented | |
| `show versions detail` | Implemented | |
| `show maps` | Implemented | |
| `show maps initiator` | Implemented | |
| `show initiators` | Implemented | |
| `show initiator names` | Implemented | |
| `show snapshots` | Partial | Support for volumes and patterns |
| | | |
| `create volume` | Partial | Support for some parameters, but not all |
| `delete volumes` | Implemented | |
| `create snapshots` | Implemented |  |
| `copy volume` | Implemented |  |
| `delete initiator-nickname` | Implemented |  |
| `delete hosts` | Implemented |  |
| `delete snapshot` | Implemented |  |
| `expand volume` | Implemented |  |
| `set initiator` | Partial | Support for id and nickname, but not profile |
| `map volume` | Partial | Support for access, lun, initiator, but not hosts, ports |
| `unmap volume` | Partial | Support for volume, initiator, but not host |
| | | |


| Resource              | Status        |
| :-------------------- | :------------ |
| `status` | Implemented |
| `redundancy` | Implemented |
| `system` | Implemented |
| `fan` | Implemented |
| `volumes` | Implemented |
| `power-supplies` | Implemented |
| `certificate-status` | Implemented |
| `disk-groups` | Implemented |
| `tiers` | Implemented |
| `pools` | Implemented |
| `drives` | Implemented |
| `advanced-settings-table` | Implemented |
| `controller-cache-parameters` | Implemented |
| `cache-settings` | Implemented |
| `initiator` | Implemented |
| `host` | Implemented |
| `host-group` | Implemented |
| `network-parameters` | Implemented |
| `sas-port` | Implemented |
| `port` | Implemented |
| `expander-ports` | Implemented |
| `expanders` | Implemented |
| `controllers` | Implemented |
| `enclosures` | Implemented |
| `versions` | Implemented |
| `volume-group-view` | Implemented |
| `volume-view-mappings` | Implemented |
| `volume-view` | Implemented |
| `initiator-view` | Implemented |
| `snapshots` | Implemented |

## Generation Steps

| Command                             | Notes |
| :---------------------------------- | :------------------------------------------------------------------------ |
| Update `generator/mc-commands.yaml` | Modify this YAML file which controls MC command generation |
| `make rung`                         | Executes: `go run generator/cmd/main.go -config generator/generator.conf` |
| `make validate`                     | Use `openapi-generator-cli` to validate the MC OpenAPI spec |
| `make generate`                     | Use `openapi-generator-cli` to generate a Go MC client library |
| `make runv`                         | Executes: `go run validator/cmd/main.go -config validator/validator.conf`

**Note:** `validator/cmd/main.go` is used to validate a handful of MC client library commands, extend as desired.

# MC OpenAPI Roadmap

This section describes commands still needed for special use cases.

## Top Priority (Customer)

### Implemented
- show fans
- show volumes
- show power-supplies
- show certificate
- show pools
- show disks
- show advanced-settings
- show cache-parameters
- show host-groups
- show enclosures
- show controllers

## Top Priority (cluster)
- TBD

## Top Priority (CSI Driver)

### Implemented
- /login/{session}
- /create/volume/pool/{pool}/size/{size}/tier-affinity/no-affinity/{name}
- /delete/volumes/{name}
- /show/volumes
- /show/volumes/{name1,name2,name2}
- /show/pools
- /show/controllers
- /show/versions/detail
- /show/maps/{name}
- /show/maps/initiator/
- /show/initiators
- /show/initiator/{initiator}
- /show/snapshots
- /show/snapshots/volume/{volume}
- /show/snapshots/pattern/{volume}
- /create/snapshots/volumes/{volume}/{snapshot}
- /copy/volume/destination-pool/{pool}/name/{destination}/{source}
- /delete/initiator-nickname/{nickname}
- /delete/hosts/{name}
- /delete/snapshot/{name1,name2,name3}
- /expand/volume/size/{size}/{name}
- /set/initiator/id/{initiator}/nickname/{nickname}
- /map/volume/access/{access}/lun/{lun}/initiator/{initiator}/{name}
- /unmap/volume/initiator/{initiator}/{name}
- /unmap/volume/{name}


## Test Using A Live System

This option runs the Go language test cases against a live storage system. Two steps are required:
- Update `test/.env` with the correct system IP Address and credentials
  - See `test/env-example` which provides an example of the file contents
- Run `go test ./test/ -v`

Another option is to define environment variables, which take precedence over .env values
- export TEST_STORAGEIP=http://[ipaddress]
- export TEST_USERNAME=[username]
- export TEST_PASSWORD=[password]
- export TEST_INITIATOR=[initiator]
- export TEST_POOL=[pool]
- Run `go test ./test/ -v`
- unset TEST_STORAGEIP TEST_PASSWORD TEST_USERNAME TEST_INITIATOR TEST_POOL
