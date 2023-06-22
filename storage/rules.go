package storage

import (
	"encoding/json"
	"fmt"
	"github.com/NubeIO/lib-uuid/uuid"
	"github.com/tidwall/buntdb"
)

func (inst *db) AddRule(rc *RQLRule) (*RQLRule, error) {
	rc.UUID = uuid.ShortUUID("rql")
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

func (inst *db) UpdateRule(uuid string, rc *RQLRule) (*RQLRule, error) {
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

func (inst *db) DeleteRule(uuid string) error {
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

func (inst *db) SelectRule(uuid string) (*RQLRule, error) {
	var data *RQLRule
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

func (inst *db) SelectAllRules() ([]RQLRule, error) {
	var resp []RQLRule
	err := inst.DB.View(func(tx *buntdb.Tx) error {
		err := tx.Ascend("", func(key, value string) bool {
			var data RQLRule
			err := json.Unmarshal([]byte(value), &data)
			if err != nil {
				return false
			}
			if matchRuleUUID(data.UUID) {
				resp = append(resp, data)
			}
			return true
		})
		return err
	})
	if err != nil {
		fmt.Printf("Error: %s", err)
		return []RQLRule{}, err
	}
	return resp, nil
}
