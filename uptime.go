package gopingdom

import (
	"encoding/json"
	"fmt"
	"sort"
)

func (r *RestClient) UptimeGetChecks() (cks Checks, error error) {
	uri := "/checks?include_severity=true&include_tags=true"

	data, err := r.Get(uri)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return nil, err
	}

	if r.debug {
		fmt.Printf("Response: %s", data)
	}

	var result ChecksResult
	json.Unmarshal(data, &result)
	sort.Sort(result.Checks)
	return result.Checks, nil
}

func (r *RestClient) UptimeGetChecksMap() (cks map[int]Check, error error) {

	listMap := make(map[int]Check)
	c, e := r.UptimeGetChecks()
	if error != nil {
		fmt.Println("Error: " + e.Error())
		return nil, e
	}

	for _, check := range c {
		listMap[check.ID] = check
	}

	return listMap,error
}

func (r *RestClient) UptimeGetDownChecksMap() (cks map[int]Check, error error) {
	listMap := make(map[int]Check)
	c, e := r.UptimeGetChecks()
	if error != nil {
		fmt.Println("Error: " + e.Error())
		return nil, e
	}

	for _, check := range c {
		if check.Status != "paused" && check.Status != "up"{
			listMap[check.ID] = check
		}

	}

	return listMap,error
}

func (r *RestClient) UptimeGetCheckDetails(id int) (cks Check, error error) {
	uri := "/checks/" + fmt.Sprintf("%d", id)

	data, err := r.Get(uri)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return cks, err
	}

	if r.debug {
		fmt.Printf("Response: %s", data)
	}

	var result CheckResult
	json.Unmarshal(data, &result)
	return result.Checks, nil
}

func (r *RestClient) UptimeGetProbes() (prb []Probe, err error) {
	uri := "/probes"

	data, err := r.Get(uri)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return nil, err
	}

	var result ProbesResult
	json.Unmarshal(data, &result)
	return result.Probes, nil
}
