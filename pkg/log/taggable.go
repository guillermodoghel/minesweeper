package log

type Taggable interface {
	Tags() Tags
}

type errorWithTags struct {
	error
	tags Tags
}

func (err errorWithTags) Tags() Tags {
	return err.tags
}

func NewErrorWithTags(err error, tags Tags) error {
	return errorWithTags{
		error: err,
		tags:  tags,
	}
}
