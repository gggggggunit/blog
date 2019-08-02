package documents

type PostDocuments struct {
	ID      string `bson:"_id,omitempty"`
	Title   string
	Content string
}
