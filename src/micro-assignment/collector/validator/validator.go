package validator

type DefaultSchemaValidator struct{
}

func (schemaValidator *DefaultSchemaValidator) ValidateSchema() (bool){
	return true
}