package authentication

import "github.com/gin-gonic/gin"

//authenticate exchanges credentials for a valid token
func (s *Service) authenticate(c *gin.Context) {
	//MOD: Can access s.ctx which contains needed globals
}
