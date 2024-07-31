package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"x-app-go/services"
)

func PostTweet(w http.ResponseWriter, r *http.Request) {
	var tw services.Tweet
	userID := r.Header.Get("id")
	err := json.NewDecoder(r.Body).Decode(&tw)
	if err != nil {
		ResponseBuilder(w, Response{
			Msg:  err.Error(),
			Code: http.StatusBadRequest,
		})
		return
	}

	err = validateTweet(tw)
	if err != nil {
		ResponseBuilder(w, Response{
			Msg:  err.Error(),
			Code: http.StatusBadRequest,
		})
		return
	}

	err = tw.InsertTweet(userID, tw)
	if err != nil {
		res := Response{
			Msg:  fmt.Sprintf("cannot create tweet: %s", err.Error()),
			Code: http.StatusInternalServerError,
		}
		ResponseBuilder(w, res)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{
		Msg:  fmt.Sprintf("posted tweet: %s", tw.Content),
		Code: http.StatusCreated,
	})

	return
}

func validateTweet(tw services.Tweet) error {
	switch tw.Content {
	case "":
		return fmt.Errorf("empty tweet")
	case "error":
		return fmt.Errorf("unexpected error")
	}
	if len(tw.Content) > 280 {
		return fmt.Errorf("support a maximum content of 280")
	}

	return nil
}

func GetTweets(w http.ResponseWriter, r *http.Request) {
	var tw services.Tweet
	userID := r.Header.Get("id")
	if userID == "" {
		ResponseBuilder(w, Response{
			Msg:  "empty id",
			Code: http.StatusBadRequest,
		})
		return
	}

	tweets, err := tw.GetTweetsOfUser(userID)
	if err != nil {
		res := Response{
			Msg:  fmt.Sprintf("cannot obtain tweets: %s", err.Error()),
			Code: http.StatusInternalServerError,
		}
		ResponseBuilder(w, res)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tweets)

	return
}
