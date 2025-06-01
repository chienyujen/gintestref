package handle

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/valkey-io/valkey-go"
)

// Define the structure for the data to be stored in Redis
type AssuranceData struct {
	RPCode         string  `json:"RPCode"`
	RuleGroupID    int     `json:"RuleGroupID"`
	AssuranceType  int     `json:"AssuranceType"`
	RuleMode       int     `json:"RuleMode"`
	Status         int     `json:"Status"`
	RuleWeight     int     `json:"RuleWeight"`
	LevelNumber    int     `json:"LevelNumber"`
	WeightScore    float64 `json:"WeightScore"`
	ExpectedResult bool    `json:"ExpectedResult"`
	Description    string  `json:"Description"`
}

// Helper struct for ConditionValue to handle varying structures
type ConditionValueHolder struct {
	Value interface{} `json:"value"`
}

// Structure for "behavior" data items
type BehaviorRule struct {
	RuleGroupID      int                  `json:"RuleGroupID"`
	RPCode           string               `json:"RPCode"`
	RuleCategoryType int                  `json:"RuleCategoryType"`
	RuleSubtype      int                  `json:"RuleSubtype"`
	RuleMode         int                  `json:"RuleMode"`
	Status           int                  `json:"Status"`
	RuleWeight       int                  `json:"RuleWeight"`
	RuleID           int                  `json:"RuleID"`
	ConditionType    int                  `json:"ConditionType"`
	ConditionValue   ConditionValueHolder `json:"ConditionValue"`
	WeightScore      float64              `json:"WeightScore"`
	ExpectedResult   bool                 `json:"ExpectedResult"`
	Priority         int                  `json:"Priority"`
	Description      string               `json:"Description"`
}

// Structure for "custom" data items
type CustomRule struct {
	RuleGroupID      int                  `json:"RuleGroupID"`
	RPCode           string               `json:"RPCode"`
	RuleCategoryType int                  `json:"RuleCategoryType"`
	RuleSubtype      int                  `json:"RuleSubtype"`
	RuleMode         int                  `json:"RuleMode"`
	Status           int                  `json:"Status"`
	RuleWeight       int                  `json:"RuleWeight"`
	RuleID           int                  `json:"RuleID"`
	ConditionType    int                  `json:"ConditionType"`
	ConditionValue   ConditionValueHolder `json:"ConditionValue"`
	WeightScore      float64              `json:"WeightScore"`
	ExpectedResult   bool                 `json:"ExpectedResult"`
	Priority         int                  `json:"Priority"`
	Description      string               `json:"Description"`
}

// Data to be set for "ass"
var assData = map[string][]AssuranceData{
	"8": {
		{
			RPCode:         "abcd99",
			RuleGroupID:    8,
			AssuranceType:  2,
			RuleMode:       1,
			Status:         1,
			RuleWeight:     1,
			LevelNumber:    3,
			WeightScore:    0.7,
			ExpectedResult: false,
			Description:    "",
		},
		{
			RPCode:         "abcd99",
			RuleGroupID:    8,
			AssuranceType:  2,
			RuleMode:       1,
			Status:         1,
			RuleWeight:     1,
			LevelNumber:    2,
			WeightScore:    0.6,
			ExpectedResult: false,
			Description:    "",
		},
		{
			RPCode:         "abcd99",
			RuleGroupID:    8,
			AssuranceType:  2,
			RuleMode:       1,
			Status:         1,
			RuleWeight:     1,
			LevelNumber:    1,
			WeightScore:    0.4,
			ExpectedResult: false,
			Description:    "",
		},
	},
	"11": {
		{
			RPCode:         "abcd99",
			RuleGroupID:    11,
			AssuranceType:  3,
			RuleMode:       1,
			Status:         1,
			RuleWeight:     1,
			LevelNumber:    3,
			WeightScore:    0.57,
			ExpectedResult: false,
			Description:    "test199",
		},
		{
			RPCode:         "abcd99",
			RuleGroupID:    11,
			AssuranceType:  3,
			RuleMode:       1,
			Status:         1,
			RuleWeight:     1,
			LevelNumber:    2,
			WeightScore:    0.44,
			ExpectedResult: false,
			Description:    "test199",
		},
		{
			RPCode:         "abcd99",
			RuleGroupID:    11,
			AssuranceType:  3,
			RuleMode:       1,
			Status:         1,
			RuleWeight:     1,
			LevelNumber:    1,
			WeightScore:    0.33,
			ExpectedResult: false,
			Description:    "test199",
		},
	},
}

// Data to be set for "behavior"
var behaviorData = map[string][]BehaviorRule{
	"7": {
		{
			RuleGroupID:      7,
			RPCode:           "abcd99",
			RuleCategoryType: 3,
			RuleSubtype:      3,
			RuleMode:         1,
			Status:           1,
			RuleWeight:       5,
			RuleID:           10,
			ConditionType:    3,
			ConditionValue:   ConditionValueHolder{Value: []string{"10.0.0.0/8", "192.168.0.0/16"}},
			WeightScore:      0.01,
			ExpectedResult:   false,
			Priority:         1,
			Description:      "範例：檢查IP範圍",
		},
	},
	"8": {
		{
			RuleGroupID:      8,
			RPCode:           "abcd99",
			RuleCategoryType: 4,
			RuleSubtype:      2,
			RuleMode:         2,
			Status:           1,
			RuleWeight:       3,
			RuleID:           11,
			ConditionType:    2,
			ConditionValue:   ConditionValueHolder{Value: map[string]string{"etime": "18:00:00", "stime": "08:00:00"}},
			WeightScore:      0,
			ExpectedResult:   true,
			Priority:         1,
			Description:      "範例：時間範圍",
		},
		{
			RuleGroupID:      8,
			RPCode:           "abcd99",
			RuleCategoryType: 4,
			RuleSubtype:      2,
			RuleMode:         2,
			Status:           1,
			RuleWeight:       3,
			RuleID:           12,
			ConditionType:    2,
			ConditionValue:   ConditionValueHolder{Value: map[string]string{"etime": "23:59:59", "stime": "18:00:00"}},
			WeightScore:      0,
			ExpectedResult:   false,
			Priority:         1,
			Description:      "範例：時間範圍",
		},
	},
}

// Data to be set for "custom"
var customData = map[string][]CustomRule{
	"33": {
		{
			RuleGroupID:      33,
			RPCode:           "abcd99",
			RuleCategoryType: 1,
			RuleSubtype:      1,
			RuleMode:         1,
			Status:           1,
			RuleWeight:       5,
			RuleID:           60,
			ConditionType:    1,
			ConditionValue:   ConditionValueHolder{Value: map[string]int{"max": 91, "min": 78}},
			WeightScore:      0.02,
			ExpectedResult:   false,
			Priority:         1,
			Description:      "範例：檢查設定總數",
		},
		{
			RuleGroupID:      33,
			RPCode:           "abcd99",
			RuleCategoryType: 1,
			RuleSubtype:      1,
			RuleMode:         1,
			Status:           1,
			RuleWeight:       5,
			RuleID:           61,
			ConditionType:    1,
			ConditionValue:   ConditionValueHolder{Value: map[string]int{"max": 100, "min": 90}},
			WeightScore:      0.9,
			ExpectedResult:   false,
			Priority:         2,
			Description:      "範例：檢查設定總數",
		},
	},
	"34": {
		{
			RuleGroupID:      34,
			RPCode:           "abcd99",
			RuleCategoryType: 2,
			RuleSubtype:      5,
			RuleMode:         2,
			Status:           1,
			RuleWeight:       3,
			RuleID:           62,
			ConditionType:    5,
			ConditionValue:   ConditionValueHolder{Value: true},
			WeightScore:      0,
			ExpectedResult:   true,
			Priority:         1,
			Description:      "範例：是否已安裝防毒",
		},
		{
			RuleGroupID:      34,
			RPCode:           "abcd99",
			RuleCategoryType: 2,
			RuleSubtype:      5,
			RuleMode:         2,
			Status:           1,
			RuleWeight:       3,
			RuleID:           63,
			ConditionType:    5,
			ConditionValue:   ConditionValueHolder{Value: false},
			WeightScore:      0,
			ExpectedResult:   false,
			Priority:         2,
			Description:      "範例：是否已安裝防毒",
		},
	},
}

// SetValHandler handles the /setval route
func SetValHandler(client valkey.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Handle "ass" data
		redisKeyAss := "myside:abcd1234:ass"
		for field, dataList := range assData {
			jsonData, err := json.Marshal(dataList)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal data for ass/" + field + ": " + err.Error()})
				return
			}
			err = client.Do(c, client.B().Hset().Key(redisKeyAss).FieldValue().FieldValue(field, string(jsonData)).Build()).Error()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set value in Redis for ass/" + field + ": " + err.Error()})
				return
			}
		}

		// Handle "behavior" data
		redisKeyBehavior := "myside:abcd1234:behavior"
		for field, dataList := range behaviorData {
			jsonData, err := json.Marshal(dataList)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal data for behavior/" + field + ": " + err.Error()})
				return
			}
			err = client.Do(c, client.B().Hset().Key(redisKeyBehavior).FieldValue().FieldValue(field, string(jsonData)).Build()).Error()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set value in Redis for behavior/" + field + ": " + err.Error()})
				return
			}
		}

		// Handle "custom" data
		redisKeyCustom := "myside:abcd1234:custom"
		for field, dataList := range customData {
			jsonData, err := json.Marshal(dataList)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal data for custom/" + field + ": " + err.Error()})
				return
			}
			err = client.Do(c, client.B().Hset().Key(redisKeyCustom).FieldValue().FieldValue(field, string(jsonData)).Build()).Error()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set value in Redis for custom/" + field + ": " + err.Error()})
				return
			}
		}

		c.JSON(http.StatusOK, gin.H{"message": "Data set successfully for all categories"})
	}
}

// GetValHandler handles the /getval route
func GetValHandler(client valkey.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		redisKey := "myside:abcd1234:ass"
		// Using HGetAll to retrieve all fields and values
		result, err := client.Do(c, client.B().Hgetall().Key(redisKey).Build()).AsStrMap()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get value from Redis: " + err.Error()})
			return
		}

		if len(result) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"message": "No data found for the key"})
			return
		}

		// Convert the map of JSON strings back to the desired structure
		responseMap := make(map[string][]AssuranceData)
		for field, jsonStr := range result {
			var dataList []AssuranceData
			if err := json.Unmarshal([]byte(jsonStr), &dataList); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal data for field " + field + ": " + err.Error()})
				return
			}
			responseMap[field] = dataList
		}

		c.JSON(http.StatusOK, responseMap)
	}
}

// UpdatePayload defines the structure for the /updateval request body
type UpdatePayload struct {
	Side   string          `json:"side"`
	Type   string          `json:"type"`
	ID     int             `json:"id"`
	RpCode string          `json:"rpCode"`
	Data   json.RawMessage `json:"data"`
}

// UpdateValHandler handles the /updateval POST route
func UpdateValHandler(client valkey.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload UpdatePayload
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload: " + err.Error()})
			return
		}

		redisKey := fmt.Sprintf("%s:%s:%s", payload.Side, payload.RpCode, payload.Type)
		field := strconv.Itoa(payload.ID)

		var jsonDataToStore []byte
		var err error

		switch payload.Type {
		case "ass":
			var singleData AssuranceData
			if err = json.Unmarshal(payload.Data, &singleData); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to unmarshal 'data' for type 'ass': " + err.Error()})
				return
			}
			// Storing as an array to match the structure set by SetValHandler
			jsonDataToStore, err = json.Marshal([]AssuranceData{singleData})
		case "behavior":
			var singleData BehaviorRule
			if err = json.Unmarshal(payload.Data, &singleData); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to unmarshal 'data' for type 'behavior': " + err.Error()})
				return
			}
			jsonDataToStore, err = json.Marshal([]BehaviorRule{singleData})
		case "custom":
			var singleData CustomRule
			if err = json.Unmarshal(payload.Data, &singleData); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to unmarshal 'data' for type 'custom': " + err.Error()})
				return
			}
			jsonDataToStore, err = json.Marshal([]CustomRule{singleData})
		default:
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'type' in payload: " + payload.Type})
			return
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal data for Redis: " + err.Error()})
			return
		}

		err = client.Do(c, client.B().Hset().Key(redisKey).FieldValue().FieldValue(field, string(jsonDataToStore)).Build()).Error()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set value in Redis: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Data updated successfully for %s/%s, ID: %d", payload.Type, field, payload.ID)})
	}
}
