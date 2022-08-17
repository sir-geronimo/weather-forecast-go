package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var apiKey string

type Handler interface {
	Get(context *gin.Context)
}

type Controller struct {
	Url string
}

// New adds the handlers for the Weather route group.
func New(routerGroup *gin.RouterGroup) *Controller {
	apiKey = os.Getenv("OPEN_WEATHER_API_KEY")

	ctl := &Controller{
		Url: os.Getenv("OPEN_WEATHER_API_URL") + "/weather",
	}

	routerGroup.GET("/weather/:city", ctl.Get)

	return ctl
}

// Get returns the current weather for a given city.
func (controller *Controller) Get(context *gin.Context) {
	city := context.Param("city")
	if city == "" {
		context.String(http.StatusBadRequest, "Weather is unavailable for this City.")
	}

	url := fmt.Sprintf("%s?appid=%s&units=%s&q=%s", controller.Url, apiKey, "metric", city)

	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)
	body, _ := io.ReadAll(res.Body)

	var response Response
	err := json.Unmarshal(body, &response)
	if err != nil {
		return
	}

	context.IndentedJSON(http.StatusOK, response)
}
