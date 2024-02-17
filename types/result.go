package types

// type Result struct {
// 	PlayerID uint
// 	Player   Player `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

// 	PodSize                int  `gorm:"not null;check:value >= 3;check:value <= 6;"`
// 	Place                  int  `gorm:"not null;"`
// 	TheCouncilOfWizards    bool `gorm:"default:false"`
// 	DavidAndTheGoliaths    bool `gorm:"default:false"`
// 	Untouchable            bool `gorm:"default:false"`
// 	Cleave                 bool `gorm:"default:false"`
// 	ItsFreeRealEstate      bool `gorm:"default:false"`
// 	IAmTimmy               bool `gorm:"default:false"`
// 	BigBiggerHuge          bool `gorm:"default:false"`
// 	CloseButNoCigar        bool `gorm:"default:false"`
// 	JustAsGarfieldIntended bool `gorm:"default:false"`

// 	Points      int `gorm:"-"`
// 	BonusPoints int `gorm:"-"`
// 	TotalPoints int `gorm:"-"`
// }

type ResultErrors struct {
}

type ResultParams struct {
	PlayerID               uint `form:"player_id" binding:"required"`
	PodSize                int  `form:"pod_size" binding:"required"`
	Place                  int  `form:"place" binding:"required"`
	TheCouncilOfWizards    bool `form:"the_council_of_wizards"`
	DavidAndTheGoliaths    bool `form:"david_and_the_goliaths"`
	Untouchable            bool `form:"untouchable"`
	Cleave                 bool `form:"cleave"`
	ItsFreeRealEstate      bool `form:"its_free_real_estate"`
	IAmTimmy               bool `form:"i_am_timmy"`
	BigBiggerHuge          bool `form:"big_bigger_huge"`
	CloseButNoCigar        bool `form:"close_but_no_cigar"`
	JustAsGarfieldIntended bool `form:"just_as_garfield_intended"`
}
