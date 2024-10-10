package routes

import (
	"os"

	"github.com/joho/godotenv"
)

var _ = godotenv.Load("config/.env")
var baseApi = os.Getenv("API_URL")

var (
	RegionApiUrl          = baseApi + "/region"
	PathApiUrl            = baseApi + "/path"
	PlaceApiUrl           = baseApi + "/place"
	TeamApiUrl            = baseApi + "/team"
	TeamPlaceAnswerApiUrl = baseApi + "/answer"
)
