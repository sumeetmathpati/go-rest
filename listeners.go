package main

import (
	"fmt"
	"log"
	"net/url"
	"sort"

	"fyne.io/fyne/v2/data/binding"
)

func onParamsChanged(_ string) {
	log.Println("INFO: Updated params")
	paramsObjs, err := params.Get()
	if err != nil {
		log.Println("ERR: Could not read params")
	}

	requestUrl, err := inputUrl.Get()
	if err != nil {
		msgDie("Could not get input URL")
	}

	// paramsObjs = removeDuplicateParams(paramsObjs)

	parsedUrl, err := url.Parse(requestUrl)
	if err == nil {
		newQueryParams := url.Values{}
		for _, v := range paramsObjs {
			if nameValue, ok := v.(NameValue); ok {
				name, err := nameValue.Name.Get()
				if err != nil {
					msgDie("Could not get param name")
				}
				value, err := nameValue.Value.Get()
				if err != nil {
					msgDie("Could not get param value")
				}
				newQueryParams.Add(name, value)

			}
		}
		parsedUrl.RawQuery = newQueryParams.Encode()
	}

	inputUrl.Set(parsedUrl.String())
}

func onUrlChanged(input string) {
	log.Println("INFO: Url changed")
	parsedUrl, _ := url.Parse(input)
	queryParams := parsedUrl.Query()

	err := params.Set(make([]any, 0))

	if err != nil {
		log.Println("ERR:", err)
	}
	for paramName, v := range queryParams {
		for _, paramValue := range v {
			newNameBindString := binding.NewString()
			newNameBindString.Set(paramName)

			newValueBindString := binding.NewString()
			newValueBindString.Set(paramValue)

			fmt.Printf("%s:%s\n", paramName, paramValue)
			params.Append(NameValue{Name: newNameBindString, Value: newValueBindString})
		}
	}

	paramsCopy, err := params.Get()
	if err != nil {
		msgDie("Can't read params")
	}

	sort.Slice(paramsCopy, func(i, j int) bool {
		if nameValueI, ok := paramsCopy[i].(NameValue); ok {
			if nameValueJ, ok := paramsCopy[i].(NameValue); ok {
				nameI, err := nameValueI.Name.Get()
				if err != nil {
					msgDie("Can't read the param name")
				}
				nameJ, err := nameValueJ.Name.Get()
				if err != nil {
					msgDie("Can't read the param name")
				}
				return nameI < nameJ
			}
		}
		return false
	})

}

func removeDuplicateParams(params []any) []any {
	allKeys := make(map[string]bool)
	list := []any{}
	for _, item := range params {

		NameValue := getNameValueFromInterface(item)
		name, value := getValuesFromNameValue(NameValue)

		if _, v := allKeys[fmt.Sprintf("%s:%s", name, value)]; !v {
			key := fmt.Sprintf("%s:%s", name, value)
			allKeys[key] = true
			list = append(list, item)
		}

	}
	fmt.Println("LEN", len(list))
	return list
}

func getValuesFromNameValue(nameValue NameValue) (string, string) {
	name, err := nameValue.Name.Get()
	if err != nil {
		msgDie("Can't get Name from NameValue")
	}

	value, err := nameValue.Value.Get()
	if err != nil {
		msgDie("Can't get Value from NameValue")
	}

	return name, value
}

func getNameValueFromInterface(nameValue any) (nameValueObj NameValue) {
	if nameValueDetected, ok := nameValue.(NameValue); !ok {
		msgDie("Can't get NameValue")
	} else {
		nameValueObj.Name = nameValueDetected.Name
		nameValueObj.Value = nameValueDetected.Value
	}
	return
}
