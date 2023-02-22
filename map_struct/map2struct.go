package mapstruct

import "encoding/json"

// MapTransferStruct: map[string]interface{} transfer to struct,
// structure must be a pointer
func MapTransferStruct(dict map[string]interface{}, structure interface{}) error {
	bts, err := json.Marshal(dict)
	if err != nil {
		return err
	}
	return json.Unmarshal(bts, structure)
}

// StructTransferMap: struct transfer to map[string]interface{} ,
func StructTransferMap(structure interface{}, dict map[string]interface{}) error {
	bts, err := json.Marshal(structure)
	if err != nil {
		return err
	}
	return json.Unmarshal(bts, &dict)
}
