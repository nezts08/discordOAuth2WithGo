package models

type UserDC struct {
	ID       string    `json:"id"`
	Username string    `json:"username"`
	Avatar   string    `json:"avatar"`
	Guilds   []GuildDC `json:"guilds"`
}
