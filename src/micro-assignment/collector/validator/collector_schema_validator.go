package validator


// CollectionValidator is an interface to validate collected resources as per provide schema
type ResourceSchemaValidator interface {
   ReadSchemas(path string) (bool)
   ValidateSchema() (bool)
}

