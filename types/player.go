package types

type PlayerErrors struct {
}

type PlayerParams struct {
	Name string `form:"name" binding:"required"`
	Deck string `form:"deck" binding:"required,oneof=tokens dragons goad flying zombies"`
}
