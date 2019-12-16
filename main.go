package resourcehelper

import (
	"errors"
	"strings"
)

// ResourceItem response data
type ResourceItem struct {
	Patrition      string
	Service        string
	Region         string
	CustomerNumber string
	ResourceType   string
	Resource       string
}

// GetResourceIds get list resources id from list resource string
func GetResourceIds(resourceStr []string) ([]string, error) {
	var resourceIds []string

	for i := 0; i < len(resourceStr); i++ {
		r := GetResourceFormat(resourceStr[i])
		if r == nil {
			continue
		}
		if Contains(resourceIds, r.Resource) {
			continue
		}
		resourceIds = append(resourceIds, r.Resource)
	}
	if resourceIds == nil || len(resourceIds) < 0 {
		return nil, errors.New("Resource Empty")

	}
	return resourceIds, nil
}

// GetResourceFormat parser resource from resource orn
func GetResourceFormat(resourceOrn string) *ResourceItem {
	var result *ResourceItem
	// resourceFormatStr := `orn:%s:%s:%s:%s:%s/%s`
	//check formart of resourceOrn must follow `orn:%s:%s:%s:%s:%s/%s`
	//by pass check format if resourceOrn  "*"
	if resourceOrn == "*" {
		result = &ResourceItem{
			Patrition:      "*",
			Service:        "*",
			Region:         "*",
			CustomerNumber: "*",
			ResourceType:   "*",
			Resource:       "*",
		}
		return result
	}
	//`orn:%s:%s:%s:%s:%s/%s`
	orn := strings.Split(resourceOrn, "/")
	if orn == nil || len(orn) < 2 {
		return nil
	}
	prefixOrn := strings.Split(orn[0], ":")
	if prefixOrn == nil || len(prefixOrn) < 6 || prefixOrn[0] != "orn" {
		return nil
	}
	result = &ResourceItem{
		Patrition:      prefixOrn[1],
		Service:        prefixOrn[2],
		Region:         prefixOrn[3],
		CustomerNumber: prefixOrn[4],
		ResourceType:   prefixOrn[5],
		Resource:       orn[1],
	}
	return result
}
// Contains contain an string item in string slice
func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
