# 🚀 Discord Dashboard with Go & Fiber

A powerful, high-performance Discord Dashboard built with **Go (Golang)**, **Fiber**, and **MongoDB**. This project features a custom OAuth2 implementation, a smart caching system for Discord Guilds to avoid rate limits, and a dynamic UI using Pug.

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Fiber](https://img.shields.io/badge/Fiber-000000?style=for-the-badge&logo=gofiber&logoColor=white)
![Discord](https://img.shields.io/badge/Discord-5865F2?style=for-the-badge&logo=discord&logoColor=white)
![MongoDB](https://img.shields.io/badge/MongoDB-47A248?style=for-the-badge&logo=mongodb&logoColor=white)

---

## ✨ Features

- **OAuth2 Authentication**: Secure login flow using Discord's official API.
- **Smart Guild Caching**: Custom caching logic that updates every 5 minutes to prevent Discord API Rate Limits.
- **Real-time Sync**: Automatic database refresh when the bot is invited to a new server using `state` parameter tracking.
- **Dynamic UI**: Styled with **Bulma CSS** and rendered with **Pug** templates.
- **Robust Data Handling**: Custom JSON unmarshalling to handle Discord's "string vs int64" permission fields.

---

## 🛠️ Tech Stack

- **Backend**: [Go (Fiber v3)](https://gofiber.io/)
- **Database**: MongoDB (via official mongo-driver)
- **Templating**: Pug (via [Jade](https://github.com/Joker/jade) engine)
- **Styling**: [Bulma CSS](https://bulma.io/)
- **Auth**: JWT (JSON Web Tokens) & Discord OAuth2

---

## 🚀 Getting Started

### 1. Prerequisites

- Go 1.21 or higher.
- A Discord Application (from [Discord Developer Portal](https://discord.com/developers/applications)).
- A MongoDB instance.

### 2. Environment Variables

Create a `.env` file in the root directory. Use the following structure:

```env
PORT=3000
CLIENT_ID=your_client_id
CLIENT_SECRET=your_client_secret
CLIENT_REDIRECT=http://localhost:3000/auth/redirect
TOKEN=your_bot_token
JWT_SECRET=your_secret_key
DB_USER=admin
DB_PASSWORD=password
DB_NAME=Testing

```
