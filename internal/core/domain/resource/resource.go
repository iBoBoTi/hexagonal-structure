package domain

type Resource struct {
	Reference    string `validate:"required" json:"reference" bson:"reference"`
	CreatedOn    string `validate:"required" json:"created_on" bson:"created_on"`
	LastModified string `validate:"required" json:"last_modified" bson:"last_modified"`
	Name         string `validate:"required" json:"name" bson:"name"`
	Value        string `validate:"required" json:"value" bson:"value"`
}
