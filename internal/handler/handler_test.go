package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mespinosago/unitag/internal/handler/mocks"
	"github.com/mespinosago/unitag/internal/model"
	"github.com/mespinosago/unitag/internal/service"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTripsHandler_GetURL(t *testing.T) {
	t.Run("should return bad request if format of code is invalid", func(t *testing.T) {
		res := httptest.NewRecorder()
		gin.SetMode(gin.TestMode)
		c, _ := gin.CreateTestContext(res)
		c.AddParam("code", "wrong_code_format")
		c.Request = httptest.NewRequest(http.MethodGet, "/wrong_code_format", nil)

		h := NewHandler(nil)

		h.GetURL(c)

		_, err := io.ReadAll(res.Body)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("should return not found if code does not exist", func(t *testing.T) {
		res := httptest.NewRecorder()
		gin.SetMode(gin.TestMode)
		c, _ := gin.CreateTestContext(res)
		c.AddParam("code", "XXXXXX")
		c.Request = httptest.NewRequest(http.MethodGet, "/XXXXXX", nil)

		serv := mocks.NewService(t)
		serv.EXPECT().GetURL("XXXXXX", model.Options{
			Language: model.LanguageEnglish,
			OS:       model.OSMac,
			Browser:  model.BrowserChrome,
		}).Return("", service.ErrCodeNotFound)

		h := NewHandler(serv)

		h.GetURL(c)

		_, err := io.ReadAll(res.Body)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, res.Code)
	})

	t.Run("should redirect to url if code is correct", func(t *testing.T) {
		gin.SetMode(gin.TestMode)
		res := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(res)
		c.AddParam("code", "r7TH8k")
		c.Request = httptest.NewRequest(http.MethodGet, "/r7TH8k", nil)
		c.Request.Header.Add("Accept-Language", "en-US")
		c.Request.Header.Add("User-Agent", "Android")

		serv := mocks.NewService(t)
		serv.EXPECT().GetURL("r7TH8k", model.Options{
			Language: model.LanguageEnglish,
			OS:       model.OSAndroid,
			Browser:  model.BrowserChrome,
		}).Return("https://www.google.com", nil)

		h := NewHandler(serv)

		h.GetURL(c)
		_, err := io.ReadAll(res.Body)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusFound, res.Code)
	})
}
