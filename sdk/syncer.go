package sdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

type TableColumn struct {
	Name        string   `json:"name"`
	Type        string   `json:"type"`
	CasdoorName string   `json:"casdoorName"`
	IsKey       bool     `json:"isKey"`
	IsHashed    bool     `json:"isHashed"`
	Values      []string `json:"values"`
}

// Syncer has the same definition as https://github.com/casdoor/casdoor/blob/master/object/syncer.go#L24
type Syncer struct {
	Owner       string `xorm:"varchar(100) notnull pk" json:"owner"`
	Name        string `xorm:"varchar(100) notnull pk" json:"name"`
	CreatedTime string `xorm:"varchar(100)" json:"createdTime"`

	Organization string `xorm:"varchar(100)" json:"organization"`
	Type         string `xorm:"varchar(100)" json:"type"`

	Host             string         `xorm:"varchar(100)" json:"host"`
	Port             int            `json:"port"`
	User             string         `xorm:"varchar(100)" json:"user"`
	Password         string         `xorm:"varchar(100)" json:"password"`
	DatabaseType     string         `xorm:"varchar(100)" json:"databaseType"`
	Database         string         `xorm:"varchar(100)" json:"database"`
	Table            string         `xorm:"varchar(100)" json:"table"`
	TablePrimaryKey  string         `xorm:"varchar(100)" json:"tablePrimaryKey"`
	TableColumns     []*TableColumn `xorm:"mediumtext" json:"tableColumns"`
	AffiliationTable string         `xorm:"varchar(100)" json:"affiliationTable"`
	AvatarBaseUrl    string         `xorm:"varchar(100)" json:"avatarBaseUrl"`
	ErrorText        string         `xorm:"mediumtext" json:"errorText"`
	SyncInterval     int            `json:"syncInterval"`
	IsReadOnly       bool           `json:"isReadOnly"`
	IsEnabled        bool           `json:"isEnabled"`

	// Ormer *Ormer `xorm:"-" json:"-"`
}

func (c *Client) GetSyncers() ([]*Syncer, error) {
	queryMap := map[string]string{
		"owner": c.OrganizationName,
	}

	url := c.GetUrl("get-syncers", queryMap)

	bytes, err := c.DoGetBytes(url)
	if err != nil {
		return nil, err
	}

	var syncers []*Syncer
	err = json.Unmarshal(bytes, &syncers)
	if err != nil {
		return nil, err
	}
	return syncers, nil
}

func (c *Client) GetPaginationSyncers(p int, pageSize int, queryMap map[string]string) ([]*Syncer, int, error) {
	queryMap["owner"] = c.OrganizationName
	queryMap["p"] = strconv.Itoa(p)
	queryMap["pageSize"] = strconv.Itoa(pageSize)

	url := c.GetUrl("get-models", queryMap)

	response, err := c.DoGetResponse(url)
	if err != nil {
		return nil, 0, err
	}

	dataBytes, err := json.Marshal(response.Data)
	if err != nil {
		return nil, 0, err
	}

	var syncers []*Syncer
	err = json.Unmarshal(dataBytes, &syncers)
	if err != nil {
		return nil, 0, errors.New("response data format is incorrect")
	}

	return syncers, int(response.Data2.(float64)), nil
}

func (c *Client) GetSyncer(name string) (*Syncer, error) {
	queryMap := map[string]string{
		"id": fmt.Sprintf("%s/%s", c.OrganizationName, name),
	}

	url := c.GetUrl("get-syncer", queryMap)

	bytes, err := c.DoGetBytes(url)
	if err != nil {
		return nil, err
	}

	var syncer *Syncer
	err = json.Unmarshal(bytes, &syncer)
	if err != nil {
		return nil, err
	}
	return syncer, nil
}

func (c *Client) AddSyncer(syncer *Syncer) (bool, error) {
	_, affected, err := c.modifySyncer("add-syncer", syncer, nil)
	return affected, err
}

func (c *Client) UpdateSyncer(syncer *Syncer) (bool, error) {
	_, affected, err := c.modifySyncer("update-syncer", syncer, nil)
	return affected, err
}

func (c *Client) DeleteSyncer(syncer *Syncer) (bool, error) {
	_, affected, err := c.modifySyncer("delete-syncer", syncer, nil)
	return affected, err
}
