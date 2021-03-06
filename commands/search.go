package commands

import (
	"errors"

	"github.com/emersion/go-imap"
)

// A SEARCH command.
// See RFC 3501 section 6.4.4
type Search struct {
	Charset  string
	Criteria *imap.SearchCriteria
}

func (cmd *Search) Command() *imap.Command {
	var args []interface{}
	if cmd.Charset != "" {
		args = append(args, "CHARSET", cmd.Charset)
	}
	args = append(args, cmd.Criteria.Format()...)

	return &imap.Command{
		Name:      imap.Search,
		Arguments: args,
	}
}

func (cmd *Search) Parse(fields []interface{}) error {
	if len(fields) == 0 {
		return errors.New("Missing search criteria")
	}

	// Parse charset
	if f, ok := fields[0].(string); ok && f == "CHARSET" {
		if len(fields) < 2 {
			return errors.New("Missing CHARSET value")
		}
		if cmd.Charset, ok = fields[1].(string); !ok {
			return errors.New("Charset must be a string")
		}

		fields = fields[2:]
	}

	cmd.Criteria = &imap.SearchCriteria{}
	return cmd.Criteria.Parse(fields)
}
