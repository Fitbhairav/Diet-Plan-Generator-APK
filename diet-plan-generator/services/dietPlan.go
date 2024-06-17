// package services

// import (
//     "bytes"
//     "encoding/json"
//     "fmt"
//     "io/ioutil"
//     "net/http"
//     "os"

//     "diet-plan-generator/models"
// )

// func GenerateDietPlan(prompt string) (string, error) {
//     apiKey := os.Getenv("OPENAI_API_KEY")
//     chatGPTUrl := "https://api.openai.com/v1/chat/completions"

//     requestData := models.ChatGPTRequest{
//         Messages: []models.ChatGPTMessage{
//             {
//                 Role:    "user",
//                 Content: prompt,
//             },
//         },
//         MaxTokens: 100,
//         Model:     "gpt-3.5-turbo", // Update the model to a supported one
//     }

//     jsonData, err := json.Marshal(requestData)
//     if err != nil {
//         return "", fmt.Errorf("error marshaling request data: %v", err)
//     }

//     client := &http.Client{}
//     req, err := http.NewRequest("POST", chatGPTUrl, bytes.NewReader(jsonData))
//     if err != nil {
//         return "", fmt.Errorf("error creating new request: %v", err)
//     }

//     req.Header.Set("Content-Type", "application/json")
//     req.Header.Set("Authorization", "Bearer "+apiKey)

//     resp, err := client.Do(req)
//     if err != nil {
//         return "", fmt.Errorf("error making request to OpenAI API: %v", err)
//     }
//     defer resp.Body.Close()

//     body, err := ioutil.ReadAll(resp.Body)
//     if err != nil {
//         return "", fmt.Errorf("error reading response body: %v", err)
//     }

//     if resp.StatusCode != http.StatusOK {
//         if resp.StatusCode == http.StatusTooManyRequests {
//             return "", fmt.Errorf("You have exceeded your API quota. Please check your plan and billing details.")
//         }
//         return "", fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
//     }

//     var chatGPTResponse models.ChatGPTResponse
//     err = json.Unmarshal(body, &chatGPTResponse)
//     if err != nil {
//         return "", fmt.Errorf("error unmarshaling response data: %v", err)
//     }

//     if len(chatGPTResponse.Choices) == 0 {
//         return "", fmt.Errorf("no choices returned from the API: %s", string(body))
//     }

//     return chatGPTResponse.Choices[0].Message.Content, nil
// }
package services

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"

    "diet-plan-generator/models"
)

func GenerateDietPlan(request models.RequestBody) (map[string]interface{}, error) {
    rapidApiKey := os.Getenv("RAPIDAPI_KEY")
    rapidApiUrl := "https://ai-diet-planner.p.rapidapi.com/api/generate"

    jsonData, err := json.Marshal(request)
    if err != nil {
        return nil, fmt.Errorf("error marshaling request data: %v", err)
    }

    client := &http.Client{}
    req, err := http.NewRequest("POST", rapidApiUrl, bytes.NewReader(jsonData))
    if err != nil {
        return nil, fmt.Errorf("error creating new request: %v", err)
    }

    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("x-rapidapi-host", "ai-diet-planner.p.rapidapi.com")
    req.Header.Set("x-rapidapi-key", rapidApiKey)

    resp, err := client.Do(req)
    if err != nil {
        return nil, fmt.Errorf("error making request to RapidAPI: %v", err)
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("error reading response body: %v", err)
    }

    if resp.StatusCode != http.StatusOK {
        if resp.StatusCode == http.StatusTooManyRequests {
            return nil, fmt.Errorf("You have exceeded your API quota. Please check your plan and billing details.")
        }
        return nil, fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(body))
    }

    var response map[string]interface{}
    err = json.Unmarshal(body, &response)
    if err != nil {
        return nil, fmt.Errorf("error unmarshaling response data: %v", err)
    }

    // Ensure the response is structured correctly
    if _, ok := response["response"].(map[string]interface{}); !ok {
        return nil, fmt.Errorf("invalid response format: %s", string(body))
    }

    return response, nil
}
