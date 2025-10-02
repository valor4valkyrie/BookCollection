package properties

import "github.com/magiconair/properties"

const propertiesLocation = "./properties/properties.yml"

func GetProperty(key string) string {
	p := properties.MustLoadFile(propertiesLocation, properties.UTF8)
	value := p.MustGetString(key)

	return value
}
