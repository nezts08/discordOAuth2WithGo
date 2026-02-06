package controller

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/nezts08/discordOAuth2WithGo/auth"
	"github.com/nezts08/discordOAuth2WithGo/models"
	"github.com/nezts08/discordOAuth2WithGo/repository"
	"github.com/nezts08/discordOAuth2WithGo/utils"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func AuthLoginController(c fiber.Ctx) error {
	url := auth.DiscordOAuthConfig().AuthCodeURL("state")
	return c.Redirect().To(url)
}

func AuthRedirectController(c fiber.Ctx) error {
	code := c.Query("code")
	state := c.Query("state")

	if code == "" {
		return c.Redirect().To("/forbidden")
	}

	if state == "inviting_bot" {
		utils.ForceBotCacheUpdate()
		fmt.Println("🔄 force refresh: User invited the bot.")
	} else {
		fmt.Println("✅ Using cache: Normal login.")
	}

	token, err := auth.DiscordOAuthConfig().Exchange(c.Context(), code)
	if err != nil {
		return c.Redirect().To("/forbidden")
	}

	client := auth.DiscordOAuthConfig().Client(c.Context(), token)

	resp, err := client.Get("https://discord.com/api/users/@me")
	if err != nil {
		return c.Redirect().To("/forbidden")
	}

	defer resp.Body.Close()

	var user models.UserDC

	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return c.Redirect().To("/forbidden")
	}

	guildResp, err := client.Get("https://discord.com/api/users/@me/guilds")
	if err != nil {
		return c.Redirect().To("/forbidden")
	}
	defer guildResp.Body.Close()

	var allGuilds []models.GuildDC
	if err := json.NewDecoder(guildResp.Body).Decode(&allGuilds); err != nil {
		return c.Redirect().To("/forbidden")
	}

	utils.ForceBotCacheUpdate()

	botGuildMap, err := utils.GetBotGuilds()
	if err != nil {
		fmt.Println("Error de caché del bot:", err)
		botGuildMap = make(map[string]bool)
	}

	var adminGuilds []models.GuildDC
	const adminPermission = 0x8

	for i := 0; i < len(allGuilds); i++ {
		guild := &allGuilds[i]

		if guild.Owner || (int64(guild.Permissions)&adminPermission) == adminPermission {

			if botGuildMap[guild.ID] {
				guild.InServer = true
			} else {
				guild.InServer = false
			}
			adminGuilds = append(adminGuilds, *guild)
		}
	}

	_, err = repository.FindUserByDiscordID(user.ID)

	if err == nil {
		repository.UpdateUser(user.ID, bson.M{
			"username": user.Username,
			"avatar":   user.Avatar,
			"guilds":   adminGuilds,
		})
	} else {
		newUser := models.UserDC{
			ID:       user.ID,
			Username: user.Username,
			Avatar:   user.Avatar,
			Guilds:   adminGuilds,
		}
		repository.CreateUser(&newUser)
	}

	jwtToken, err := auth.GenerateJWT(user.ID)
	if err != nil {
		fmt.Println("JWT ERROR:", err)
		return c.Redirect().To("/forbidden")
	}

	c.Cookie(&fiber.Cookie{
		Name:     "auth",
		Value:    jwtToken,
		HTTPOnly: true,
		Secure:   true,
		SameSite: fiber.CookieSameSiteNoneMode,
		Path:     "/",
	})

	return c.Redirect().To("/dashboard")
}

func AuthLogoutController(c fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "auth",
		Value:    "",
		Path:     "/",
		HTTPOnly: true,
		Secure:   true,
		SameSite: fiber.CookieSameSiteNoneMode,
		Expires:  time.Now().Add(-time.Hour),
		MaxAge:   -1,
	})

	return c.Redirect().To("/")
}
