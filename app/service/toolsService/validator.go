package toolsService

import "github.com/gogf/gf/util/gvalid"

type CreateToolsValidator struct {
	Name        string `v:"name@required|max-length:255"`
	Icon        string `v:"icon@max-length:255"`
	Description string `v:"description@required|max-length:255"`
}

func (validator *CreateToolsValidator) Validate() error {
	if err := gvalid.CheckStruct(validator, nil); err != nil {
		return err
	}
	return nil
}
