package utils

import (
	"fmt"
	"reflect"
	"strings"
)

func ValidationPaginationDefault(page, limit int) (int, int) {

	// default
	limitDefault := 10
	limitMax := 100
	firstPage := 1

	// transform
	if limit <= 0 {
		limit = limitDefault
	} else if limit > limitMax {
		limit = limitMax
	}
	if page <= 0 {
		page = firstPage
	}

	return page, limit
}

func ValidateStruct(s interface{}) (errs []error) {

	// init
	value := reflect.ValueOf(s)
	typ := reflect.TypeOf(s)

	// validate struct variables
	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		fieldType := typ.Field(i)
		tag := fieldType.Tag.Get("validate")

		// skip if there's no validation tag
		if tag == "" {
			continue
		}

		// check rules
		rules := strings.Split(tag, ",")
		fieldValue := field.Interface()
		for _, rule := range rules {
			parts := strings.SplitN(rule, "=", 2)
			switch parts[0] {
			case "required":
				if fieldValue == reflect.Zero(field.Type()).Interface() {
					errs = append(errs, fmt.Errorf("%s is required", fieldType.Name))
				}
			case "min":
				minValue := parts[1]
				if fieldValue.(int) < atoi(minValue) {
					errs = append(errs, fmt.Errorf("%s must be at least %s", fieldType.Name, minValue))
				}
			case "max":
				maxValue := parts[1]
				if fieldValue.(int) > atoi(maxValue) {
					errs = append(errs, fmt.Errorf("%s must be at most %s", fieldType.Name, maxValue))
				}
			case "email":
				email := fieldValue.(string)
				if !strings.Contains(email, "@") || !strings.Contains(email, ".") {
					errs = append(errs, fmt.Errorf("%s is not a valid email", fieldType.Name))
				}
			}
		}

	}

	return errs
}

func atoi(s string) int {
	i := 0
	for _, r := range s {
		i = i*10 + int(r-'0')
	}
	return i
}