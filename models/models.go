package models

type Metadata struct {
	Values []Values `header:"metadata" binding:"required"`
}

type Values struct {
	FileName     string
	FilePath     string
	FileSize     string
	Location     string
	ObjectType   string
	CreationTime int64
}

type SearchMetadata struct {
	SearchAttributes map[string]string `header:"search" binding:"required"`
}
