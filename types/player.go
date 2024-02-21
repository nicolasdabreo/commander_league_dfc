package types

type PlayerErrors struct {
	Name string
	Deck string
}

type PlayerParams struct {
	Name string
	Deck string
}

func (params PlayerParams) Validate() (PlayerErrors, bool) {
	// return validation.ValidateStruct(
	// 	&params,
	// 	validation.Field(params.Name, validation.Required, validation.Length(3, 50)),
	// 	validation.Field(params.Deck, validation.In("tokens", "dragons", "goad", "flying", "zombies")),
	// )

	var errs PlayerErrors
	hasErrs := false

	if len(params.Name) == 0 {
		errs.Name = "is required"
		hasErrs = true
	}

	if len(params.Name) > 50 {
		errs.Name = "too long"
		hasErrs = true
	}

	if len(params.Deck) == 0 {
		errs.Deck = "is required"
		hasErrs = true
	}

	return errs, hasErrs
}
