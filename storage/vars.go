package storage

import (
	"encoding/json"
	"fmt"
	"github.com/NubeIO/lib-uuid/uuid"
	"github.com/tidwall/buntdb"
)

func (inst *db) AddVariable(rc *RQLVariables) (*RQLVariables, error) {
	rc.UUID = uuid.ShortUUID("var")
	data, err := json.Marshal(rc)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return nil, err
	}
	err = inst.DB.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set(rc.UUID, string(data), nil)
		return err
	})
	if err != nil {
		fmt.Printf("Error: %s", err)
		return nil, err
	}
	return rc, nil
}

func (inst *db) UpdateVariable(uuid string, rc *RQLVariables) (*RQLVariables, error) {
	j, err := json.Marshal(rc)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return nil, err
	}
	err = inst.DB.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set(uuid, string(j), nil)
		return err
	})
	if err != nil {
		fmt.Printf("Error: %s", err)
		return nil, err
	}
	return rc, nil
}

func (inst *db) DeleteVariable(uuid string) error {
	err := inst.DB.Update(func(tx *buntdb.Tx) error {
		_, err := tx.Delete(uuid)
		return err
	})
	if err != nil {
		fmt.Printf("Error delete: %s", err)
		return err
	}
	return nil
}

func (inst *db) SelectVariable(uuid string) (*RQLVariables, error) {
	var data *RQLVariables
	err := inst.DB.View(func(tx *buntdb.Tx) error {
		val, err := tx.Get(uuid)
		if err != nil {
			return err
		}
		err = json.Unmarshal([]byte(val), &data)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		fmt.Printf("Error: %s", err)
		return nil, err
	}
	return data, nil

}

func (inst *db) SelectAllVariables() ([]RQLVariables, error) {
	var resp []RQLVariables
	err := inst.DB.View(func(tx *buntdb.Tx) error {
		err := tx.Ascend("", func(key, value string) bool {
			var data RQLVariables
			err := json.Unmarshal([]byte(value), &data)
			if err != nil {
				return false
			}
			if matchVarUID(data.UUID) {
				resp = append(resp, data)
			}
			return true
		})
		return err
	})
	if err != nil {
		fmt.Printf("Error: %s", err)
		return []RQLVariables{}, err
	}
	return resp, nil
}
