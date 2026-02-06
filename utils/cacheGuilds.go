package utils

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/nezts08/discordOAuth2WithGo/models"
)

var BotGuildsCache map[string]bool
var LastBotUpdate time.Time

func GetBotGuilds() (map[string]bool, error) {
	if time.Since(LastBotUpdate) < 5*time.Minute && BotGuildsCache != nil {
		return BotGuildsCache, nil
	}

	req, _ := http.NewRequest("GET", "https://discord.com/api/v10/users/@me/guilds", nil)
	req.Header.Set("Authorization", "Bot "+os.Getenv("TOKEN"))

	httpClient := &http.Client{Timeout: 10 * time.Second}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var botGuilds []models.GuildDC
	if err := json.NewDecoder(resp.Body).Decode(&botGuilds); err != nil {
		return nil, err
	}

	newMap := make(map[string]bool)
	for _, g := range botGuilds {
		if g.ID != "" {
			newMap[g.ID] = true
		}
	}

	BotGuildsCache = newMap
	LastBotUpdate = time.Now()

	return BotGuildsCache, nil
}

func ForceBotCacheUpdate() {
	LastBotUpdate = time.Time{}
}
