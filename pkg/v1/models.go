package exosx

import "strconv"

// model : interface to allow generic conversion from raw response to user-object
type model interface {
	fillFromObject(obj *Object) error
}

// Volume : volume-view representation
type Volume struct {
	LUN int
}

func (m *Volume) fillFromObject(obj *Object) error {
	lun, err := strconv.Atoi(obj.PropertiesMap["lun"].Data)
	if err != nil {
		return err
	}

	m.LUN = lun
	return nil
}
