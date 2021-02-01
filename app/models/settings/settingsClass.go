package settings

import (
	"../database"
	"../log"
	"strconv"
)

var db = database.GetDB()

func (s *ISettings) GetIntValue() int {
	i, err := strconv.Atoi(s.GetValue())
	if err == nil {
		return i
	}
	panic(err)
}

func (s *ISettings) GetValue() string {
	log.Println(s.ID)
	if s.ID == 0 {

		_, err := db.Model(s).Where("prop_key = ?0", s.PropertyKey).SelectOrInsert()
		if err != nil {
			panic(err)
		}
	}
	return s.PropertyValue
}

func (s *ISettings) UpdateValue(v interface{}) *ISettings {

	newValue := ""

	switch v.(type) {

	case bool:
		if s.PropertyType != EPropertyType.BOOL {
			return s
		}
		newValue = strconv.FormatBool(v.(bool))
	case string:
		if s.PropertyType != EPropertyType.STR {
			return s
		}
		newValue = v.(string)
	case int:
		if s.PropertyType != EPropertyType.INT {
			return s
		}
		newValue = strconv.Itoa(v.(int))
	default:
		return s
	}

	log.Println(newValue)
	log.Println("update")
	s.PropertyValue = newValue

	_, err := db.Model(s).Column("prop_value").Where("uid = ?0", s.ID).Update()
	if err != nil {
		panic(err)
	}

	return s
}
