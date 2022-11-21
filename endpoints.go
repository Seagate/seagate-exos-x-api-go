package exosx

import (
	"crypto/md5"
	"fmt"
	"strings"

	"k8s.io/klog"
)

// SessionValid : Determine if a session is valid, if not a login is required
func (client *Client) SessionValid(addr, username string) bool {

	if client.Addr == addr && client.Username == username {
		if client.SessionKey == "" {
			klog.Infof("SessionKey is invalid: %q", client.SessionKey)
			return false
		}
		klog.Infof("client is already configured for API address %q, session is valid", addr)
		return true
	}

	return false
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

// CreateVolume : creates a volume with the given name, capacity in the given pool
func (client *Client) CreateVolume(name, size, pool, poolType string) (*Response, *ResponseStatus, error) {
	if poolType == "Virtual" {
		return client.FormattedRequest("/create/volume/pool/\"%s\"/size/%s/tier-affinity/no-affinity/\"%s\"", pool, size, name)
	} else {
		return client.FormattedRequest("/create/volume/pool/\"%s\"/size/%s/\"%s\"", pool, size, name)
	}
}

// CreateNickname : Create a nickname for an initiator. The Storage API policy is to prohibit mapping of initiators which are not either
// (a) presently connected to the array or (b) represented by an entry in the initiator nickname table.
func (client *Client) CreateNickname(name, iqn string) (*Response, *ResponseStatus, error) {
	return client.FormattedRequest("/set/initiator/id/\"%s\"/nickname/\"%s\"", iqn, name)
}

// MapVolume : map a volume to an initiator using a specified LUN
func (client *Client) MapVolume(name, initiator, access string, lun int) (*Response, *ResponseStatus, error) {
	return client.FormattedRequest("/map/volume/access/%s/lun/%d/initiator/\"%s\"/\"%s\"", access, lun, initiator, name)
}

// ShowVolumes : get informations about volumes
func (client *Client) ShowVolumes(volumes ...string) (*Response, *ResponseStatus, error) {
	if len(volumes) == 0 {
		return client.FormattedRequest("/show/volumes/")
	}
	return client.FormattedRequest("/show/volumes/\"%s\"", strings.Join(volumes, ","))
}

// UnmapVolume : unmap a volume from an initiator
func (client *Client) UnmapVolume(name, initiator string) (*Response, *ResponseStatus, error) {
	if len(initiator) == 0 {
		return client.FormattedRequest("/unmap/volume/\"%s\"", name)
	}

	return client.FormattedRequest("/unmap/volume/initiator/\"%s\"/\"%s\"", initiator, name)
}

// ExpandVolume : extend a volume if there is enough space on the vdisk
func (client *Client) ExpandVolume(name, size string) (*Response, *ResponseStatus, error) {
	return client.FormattedRequest("/expand/volume/size/\"%s\"/\"%s\"", size, name)
}

// DeleteVolume : deletes a volume
func (client *Client) DeleteVolume(name string) (*Response, *ResponseStatus, error) {
	return client.FormattedRequest("/delete/volumes/\"%s\"", name)
}

// DeleteHost : deletes a host by its ID or nickname
func (client *Client) DeleteHost(name string) (*Response, *ResponseStatus, error) {
	return client.FormattedRequest("/delete/hosts/\"%s\"", name)
}

// ShowHostMaps : list the volume mappings for given host
// If host is an empty string, mapping for all hosts is shown
func (client *Client) ShowHostMaps(host string) ([]Volume, *ResponseStatus, error) {
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
func (client *Client) ShowSnapshots(snapshotId string, sourceVolumeId string) (*Response, *ResponseStatus, error) {
	if sourceVolumeId != "" {
		return client.FormattedRequest("/show/snapshots/volume/%q", sourceVolumeId)
	} else if snapshotId != "" {
		return client.FormattedRequest("/show/snapshots/pattern/%q", snapshotId)
	}
	return client.FormattedRequest("/show/snapshots")
}

// CreateSnapshot : create a snapshot in a snap pool and the snap pool if it doesn't exsits
func (client *Client) CreateSnapshot(name string, snapshotName string) (*Response, *ResponseStatus, error) {
	return client.FormattedRequest("/create/snapshots/volumes/%q/%q", name, snapshotName)
}

// DeleteSnapshot : delete a snapshot
func (client *Client) DeleteSnapshot(names ...string) (*Response, *ResponseStatus, error) {
	return client.FormattedRequest("/delete/snapshot/%q", strings.Join(names, ","))
}

// CopyVolume : create an new volume by copying another one or a snapshot
func (client *Client) CopyVolume(sourceName string, destinationName string, pool string) (*Response, *ResponseStatus, error) {
	return client.FormattedRequest("/copy/volume/destination-pool/%q/name/%q/%q", pool, destinationName, sourceName)
}

func (client *Client) GetVolumeMapsHostNames(name string) ([]string, *ResponseStatus, error) {
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
