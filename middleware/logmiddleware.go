package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/puvadon-artmit/gofiber-template/log"
	"go.uber.org/zap"
)

func NewECSLoggerMiddleWare(c *fiber.Ctx) error {
	if err := c.Next(); err != nil {
		method := c.Route().Method
		host := c.Hostname()
		path := c.OriginalURL()
		log.Error(err.Error(),
			zap.String("method", method),
			zap.String("host", host),
			zap.Any("path", path),
		)
		return err
	} else {
		method := c.Route().Method
		host := c.Hostname()
		path := c.OriginalURL()
		log.Info("Request",
			zap.String("method", method),
			zap.String("host", host),
			zap.Any("path", path),
		)
		return nil
	}
}
