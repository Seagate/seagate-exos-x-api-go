package exosx

import (
	"crypto/md5"
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/Seagate/seagate-exos-x-api-go/pkg/common"
	"k8s.io/klog/v2"
)

// SessionValid : Determine if a session is valid, if not a login is required
func (client *Client) SessionValid(addr, username string) bool {

	// addr may include protocol and ip address, or only ip address
	// Example: addr == "http://1.2.3.4" OR "https://1.2.3.4" OR "1.2.3.4"
	// client.Addr only contains ip address
	ipaddress := addr
	if strings.HasPrefix(addr, "http://") {
		ipaddress = strings.Replace(addr, "http://", "", 1)
	} else if strings.HasPrefix(addr, "https://") {
		ipaddress = strings.Replace(addr, "https://", "", 1)
	}

	if client.Addr == ipaddress && client.Username == username {
		if client.SessionKey == "" {
			klog.Infof("SessionKey is invalid: %q", client.SessionKey)
			return false
		}
		klog.Infof("client is already configured for API address %q, session is valid", addr)
		return true
	}

	return false
}

// StoreCredentials : Called to store ip address, protocol, username, and password for the client
func (client *Client) StoreCredentials(addr string, protocol string, username string, password string) {

	// addr may include protocol and ip address, or only ip address
	// Example: addr == "http://1.2.3.4" OR "https://1.2.3.4" OR "1.2.3.4"
	// client.Addr only contains ip address
	ipaddress := addr
	if strings.HasPrefix(addr, "http://") {
		ipaddress = strings.Replace(addr, "http://", "", 1)
		if protocol == "" {
			protocol = "http"
		}
	} else if strings.HasPrefix(addr, "https://") {
		ipaddress = strings.Replace(addr, "https://", "", 1)
		if protocol == "" {
			protocol = "https"
		}
	}

	// Store a default protocol if not supplied
	if protocol == "" {
		protocol = "https"
	}

	// Validate that the plain IP Address, with no protocol, is valid
	if net.ParseIP(ipaddress) == nil {
		klog.V(0).InfoS("IP Address is invalid", "addr", addr, "protocol", protocol, "ipaddress", ipaddress)
	}

	// Store the login credentials in the Client object
	client.Username = username
	client.Password = password
	client.Addr = ipaddress
	client.Protocol = protocol
}

// Login : Called automatically, may be called manually if credentials changed
func (client *Client) Login() error {

	userpass := fmt.Sprintf("%s_%s", client.Username, client.Password)
	hash := fmt.Sprintf("%x", md5.Sum([]byte(userpass)))
	res, _, err := client.FormattedRequest("/login/%s", hash)

	if err != nil {
		return err
	}

	client.SessionKey = res.ObjectsMap["status"].PropertiesMap["response"].Data

	return nil
}

// CloseConnections: Called to close any idle connections
func (client *Client) CloseConnections() error {
	client.HTTPClient.CloseIdleConnections()
	return nil
}

// GetPoolType: Returns the pool type for a specified pool
func (client *Client) GetPoolType(pool string) (string, error) {
	return GetPoolType(client.Info, pool)
}

// CreateVolume : creates a volume with the given name, capacity in the given pool
func (client *Client) CreateVolume(name, size, pool, poolType string) (*common.VolumeObject, *common.ResponseStatus, error) {
	request := ""
	if poolType == "Virtual" {
		request = fmt.Sprintf("/create/volume/pool/\"%s\"/size/%s/tier-affinity/no-affinity/\"%s\"", pool, size, name)
	} else {
		request = fmt.Sprintf("/create/volume/pool/\"%s\"/size/%s/\"%s\"", pool, size, name)
	}

	data, status, err := client.FormattedRequest(request)

	// Fill in Volume properties for the volume data object returned
	volume := common.VolumeObject{}
	if err == nil && status.ResponseTypeNumeric == 0 && len(data.Objects) > 0 {
		volume.ObjectName = data.Objects[0].Name

		if val, ok := data.Objects[0].PropertiesMap["blocks"]; ok {
			volume.Blocks, _ = strconv.ParseInt(val.Data, 10, 64)
		}
		if val, ok := data.Objects[0].PropertiesMap["blocksize"]; ok {
			volume.BlockSize, _ = strconv.ParseInt(val.Data, 10, 64)
		}
		if val, ok := data.Objects[0].PropertiesMap["health"]; ok {
			volume.Health = val.Data
		}
		if _, ok := data.Objects[0].PropertiesMap["size-numeric"]; ok {
			volume.SizeNumeric, _ = strconv.ParseInt(data.Objects[0].PropertiesMap["size-numeric"].Data, 10, 64)
		}
		if val, ok := data.Objects[0].PropertiesMap["storage-pool-name"]; ok {
			volume.StoragePoolName = val.Data
		}
		if val, ok := data.Objects[0].PropertiesMap["storage-type"]; ok {
			volume.StorageType = val.Data
		}
		if val, ok := data.Objects[0].PropertiesMap["tier-affinity"]; ok {
			volume.TierAffinity = val.Data
		}
		if val, ok := data.Objects[0].PropertiesMap["total-size"]; ok {
			volume.TotalSize = val.Data
		}
		if val, ok := data.Objects[0].PropertiesMap["volume-name"]; ok {
			volume.VolumeName = val.Data
		}
		if val, ok := data.Objects[0].PropertiesMap["volume-type"]; ok {
			volume.VolumeType = val.Data
		}
		if val, ok := data.Objects[0].PropertiesMap["wwn"]; ok {
			volume.Wwn = val.Data
		}
	}

	// For API versions that do not return a representation of the volume object, use ShowVolumes to fill in the data
	if err == nil && status.ResponseTypeNumeric == 0 && volume.Wwn == "" {
		volumes, status, err := client.ShowVolumes(name)
		if err == nil && status.ResponseTypeNumeric == 0 {
			if len(volumes) > 0 && volumes[0].ObjectName == "volume" {
				volume.Blocks = volumes[0].Blocks
				volume.BlockSize = volumes[0].BlockSize
				volume.Health = volumes[0].Health
				volume.SizeNumeric = volumes[0].SizeNumeric
				volume.StoragePoolName = volumes[0].StoragePoolName
				volume.StorageType = volumes[0].StorageType
				volume.TierAffinity = volumes[0].TierAffinity
				volume.TotalSize = volumes[0].TotalSize
				volume.VolumeName = volumes[0].VolumeName
				volume.VolumeType = volumes[0].VolumeType
				volume.Wwn = strings.ToLower(volumes[0].Wwn)
			}
		}
	}

	return &volume, status, err
}

// GetPortals: Return a list of portals for the storage system
func (client *Client) GetPortals() (string, error) {
	return GetPortals(client.Info)
}

// CreateNickname : Create a nickname for an initiator. The Storage API policy is to prohibit mapping of initiators which are not either
// (a) presently connected to the array or (b) represented by an entry in the initiator nickname table.
func (client *Client) CreateNickname(name, iqn string) (*common.ResponseStatus, error) {
	_, status, err := client.FormattedRequest("/set/initiator/id/\"%s\"/nickname/\"%s\"", iqn, name)
	return status, err
}

// MapVolume : map a volume to an initiator using a specified LUN
func (client *Client) MapVolume(name, initiator, access string, lun int) (*common.ResponseStatus, error) {
	_, status, err := client.FormattedRequest("/map/volume/access/%s/lun/%d/initiator/\"%s\"/\"%s\"", access, lun, initiator, name)
	return status, err
}

// ShowVolumes : get information about volumes
func (client *Client) ShowVolumes(volumes ...string) ([]common.VolumeObject, *common.ResponseStatus, error) {

	request := ""
	returnVolumes := []common.VolumeObject{}

	if len(volumes) == 0 {
		request = "/show/volumes/"
	} else {
		request = fmt.Sprintf("/show/volumes/\"%s\"", strings.Join(volumes, ","))
	}

	data, status, err := client.FormattedRequest(request)

	// Fill in Volume properties for all volume data objects returned
	if err == nil && status.ResponseTypeNumeric == 0 {
		for _, object := range data.Objects {
			if object.Name == "volume" {
				volume := common.VolumeObject{}
				volume.ObjectName = object.Name
				volume.Blocks, _ = strconv.ParseInt(object.PropertiesMap["blocks"].Data, 10, 64)
				volume.BlockSize, _ = strconv.ParseInt(object.PropertiesMap["blocksize"].Data, 10, 64)
				volume.Health = object.PropertiesMap["health"].Data
				volume.SizeNumeric, _ = strconv.ParseInt(object.PropertiesMap["size-numeric"].Data, 10, 64)
				volume.StoragePoolName = object.PropertiesMap["storage-pool-name"].Data
				volume.StorageType = object.PropertiesMap["storage-type"].Data
				volume.TierAffinity = object.PropertiesMap["tier-affinity"].Data
				volume.TotalSize = object.PropertiesMap["total-size"].Data
				volume.VolumeName = object.PropertiesMap["volume-name"].Data
				volume.VolumeType = object.PropertiesMap["volume-type"].Data
				volume.Wwn = object.PropertiesMap["wwn"].Data

				returnVolumes = append(returnVolumes, volume)
			}
		}
	}

	return returnVolumes, status, err
}

// UnmapVolume : unmap a volume from an initiator
func (client *Client) UnmapVolume(name, initiator string) (*common.ResponseStatus, error) {
	if len(initiator) == 0 {
		_, status, err := client.FormattedRequest("/unmap/volume/\"%s\"", name)
		return status, err
	}
	_, status, err := client.FormattedRequest("/unmap/volume/initiator/\"%s\"/\"%s\"", initiator, name)
	return status, err
}

// ExpandVolume : extend a volume if there is enough space on the vdisk
func (client *Client) ExpandVolume(name, size string) (*common.ResponseStatus, error) {
	_, status, err := client.FormattedRequest("/expand/volume/size/\"%s\"/\"%s\"", size, name)
	return status, err
}

// DeleteVolume : deletes a volume
func (client *Client) DeleteVolume(name string) (*common.ResponseStatus, error) {
	_, status, err := client.FormattedRequest("/delete/volumes/\"%s\"", name)
	return status, err
}

// DeleteHost : deletes a host by its ID or nickname
func (client *Client) DeleteHost(name string) (*common.ResponseStatus, error) {
	_, status, err := client.FormattedRequest("/delete/hosts/\"%s\"", name)
	return status, err
}

// ShowHostMaps : list the volume mappings for given host
// If host is an empty string, mapping for all hosts is shown
func (client *Client) ShowHostMaps(host string) ([]Volume, *common.ResponseStatus, error) {
	// We don't use "/show/maps/initiator/<host>" here because
	// the maps for an initiator with a nickname or in a host group will not
	// be returned. Instead we get all initiator mappings and filter by initiator
	klog.Infof("++ ShowHostMaps(%v)", host)
	res, status, err := client.FormattedRequest("/show/maps/initiator/")
	if err != nil {
		return nil, status, err
	}

	mappings := make([]Volume, 0)
	for _, rootObj := range res.Objects {
		if host != "" {
			id, err := rootObj.GetProperties("id")
			if err != nil || id[0].Data != host {
				continue
			}
		}
		if rootObj.Name != "initiator-view" {
			continue
		}

		for i, object := range rootObj.Objects {
			if object.Name == "volume-view" {
				vol := Volume{}
				vol.fillFromObject(&object)
				klog.Infof("++ volume[%v]: %v", i, vol)
				mappings = append(mappings, vol)
			}
		}
	}

	return mappings, status, err
}

// ShowSnapshots : Show one snaphot, or all snapshots, or all snapshots for a volume
func (client *Client) ShowSnapshots(snapshotId string, sourceVolumeId string) ([]common.SnapshotObject, *common.ResponseStatus, error) {

	request := ""
	returnSnapshots := []common.SnapshotObject{}

	// Create the correctly formatted request command
	if sourceVolumeId != "" {
		request = fmt.Sprintf("/show/snapshots/volume/%q", sourceVolumeId)
	} else if snapshotId != "" {
		request = fmt.Sprintf("/show/snapshots/pattern/%q", snapshotId)
	} else {
		request = "/show/snapshots"
	}

	data, status, err := client.FormattedRequest(request)

	// Fill in Snapshot properties for all data objects returned
	if err == nil && status.ResponseTypeNumeric == 0 {
		for _, object := range data.Objects {
			if object.Name == "snapshots" {
				snapshot := common.SnapshotObject{}
				snapshot.ObjectName = object.Name
				snapshot.CreationTime, _ = common.CreationTimeFromString(object.PropertiesMap["creation-date-time-numeric"].Data)
				snapshot.CreationDateTime = object.PropertiesMap["creation-date-time"].Data
				snapshot.CreationDateTimeNumeric, _ = strconv.ParseInt(object.PropertiesMap["creation-date-time-numeric"].Data, 10, 64)
				snapshot.MasterVolumeName = object.PropertiesMap["master-volume-name"].Data
				snapshot.Name = object.PropertiesMap["name"].Data
				snapshot.StoragePoolName = object.PropertiesMap["storage-pool-name"].Data
				snapshot.TotalSize = object.PropertiesMap["total-size"].Data
				snapshot.TotalSizeNumeric, _ = strconv.ParseInt(object.PropertiesMap["total-size-numeric"].Data, 10, 64)
				snapshot.VolumeParent = object.PropertiesMap["volume-parent"].Data

				returnSnapshots = append(returnSnapshots, snapshot)
			}
		}
	}

	return returnSnapshots, status, err
}

// CreateSnapshot : create a snapshot in a snap pool and the snap pool if it doesn't exsits
func (client *Client) CreateSnapshot(name string, snapshotName string) (*common.ResponseStatus, error) {
	_, status, err := client.FormattedRequest("/create/snapshots/volumes/%q/%q", name, snapshotName)
	return status, err
}

// DeleteSnapshot : delete a snapshot
func (client *Client) DeleteSnapshot(names ...string) (*common.ResponseStatus, error) {
	_, status, err := client.FormattedRequest("/delete/snapshot/%q", strings.Join(names, ","))
	return status, err
}

// CopyVolume : create an new volume by copying another one or a snapshot
func (client *Client) CopyVolume(sourceName string, destinationName string, pool string) (*common.ResponseStatus, error) {
	_, status, err := client.FormattedRequest("/copy/volume/destination-pool/%q/name/%q/%q", pool, destinationName, sourceName)
	return status, err
}

func (client *Client) GetVolumeMapsHostNames(name string) ([]string, *common.ResponseStatus, error) {
	if name != "" {
		name = fmt.Sprintf("\"%s\"", name)
	}
	klog.V(2).Infof("++ GetVolumeMapsHostNames(%v)", name)
	res, status, err := client.FormattedRequest("/show/maps/%s", name)
	if err != nil {
		return []string{}, status, err
	}

	hostNames := []string{}
	for _, rootObj := range res.Objects {
		if rootObj.Name != "volume-view" {
			continue
		}

		for i, object := range rootObj.Objects {
			hostName := object.PropertiesMap["identifier"].Data
			if object.Name == "host-view" && hostName != "all other initiators" {
				klog.Infof("++ hostName[%v]: %v", i, hostName)
				hostNames = append(hostNames, hostName)
			}
		}
	}

	return hostNames, status, err
}
