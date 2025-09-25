package initialize

import (
	"go-telegram-bot/internal/infrastructure/repository"
)

func (c *Container) InitRepositories() {
	c.UserRepo = repository.NewUserRepository(c.DB)
	c.UserProfileRepo = repository.NewUserProfileRepo(c.DB, c.UserRepo)
	c.ChatRepo = repository.NewChatRepository(c.DB)
	c.MessageRepo = repository.NewMessageRepository(c.DB, c.UserRepo, c.ChatRepo)
}
