CREATE TABLE IF NOT EXISTS results (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	player_id INTEGER NOT NULL,
	pod_size INTEGER NOT NULL,
	place INTEGER NOT NULL,
	the_council_of_wizards BOOLEAN DEFAULT(FALSE),
	david_and_the_goliaths BOOLEAN DEFAULT(FALSE),
	untouchable BOOLEAN DEFAULT(FALSE),
	cleave BOOLEAN DEFAULT(FALSE),
	its_free_real_estate BOOLEAN DEFAULT(FALSE),
	i_am_timmy BOOLEAN DEFAULT(FALSE),
	big_bigger_huge BOOLEAN DEFAULT(FALSE),
	close_but_no_cigar BOOLEAN DEFAULT(FALSE),
	just_as_garfield_intended BOOLEAN DEFAULT(FALSE),
	created_at DATETIME default CURRENT_TIMESTAMP,
    updated_at DATETIME default CURRENT_TIMESTAMP,
	FOREIGN KEY(player_id) REFERENCES players(id)
);

