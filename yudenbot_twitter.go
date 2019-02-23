package main

import (
	"crypto/md5"
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/ChimeraCoder/anaconda"
)

var twitterAPI *anaconda.TwitterApi
var apiHash []byte

type TwitterAuth struct {
	ConsumerKey    string `yaml:"consumerKey" json:"consumerKey"`
	ConsumerSecret string `yaml:"consumerSecret" json:"consumerSecret"`
	AccessToken    string `yaml:"accessToken" json:"accessToken"`
	AccessSecret   string `yaml:"accessSecret" json:"accessSecret"`
}

func getTwitterAPI(auth *TwitterAuth) *anaconda.TwitterApi {
	// calc hash
	str := fmt.Sprintf("%v", *auth)
	s := md5.New()
	hash := s.Sum([]byte(str))

	if twitterAPI == nil || reflect.DeepEqual(apiHash, hash) == false {
		// (re)Authnication
		log.Println("TwitterAPI Authnication")
		log.Println("new authtoken hash : ", hash)
		anaconda.SetConsumerKey(auth.ConsumerKey)
		anaconda.SetConsumerSecret(auth.ConsumerSecret)
		twitterAPI = anaconda.NewTwitterApi(auth.AccessToken, auth.AccessSecret)
		apiHash = hash
	}
	return twitterAPI
}

func tweet(message string, auth *TwitterAuth) (err error) {

	api := getTwitterAPI(auth)
	if api == nil {
		log.Println("Can't Get TwitterAPI Object")
		return *new(error)
	}
	tweet, err := api.PostTweet(message, nil)
	if err != nil {
		log.Println("Error while post tweet : ", err)
		return err
	}
	log.Println("tweet success")
	log.Println(tweet.Text)
	return nil
}

// Schedule
type tweetSchedule struct {
	Event    EventData
	Time     time.Time
	Message  string
	Executed bool
	Hash     []byte
}
type tweetSchedules []tweetSchedule

var hasher = md5.New()

func (s *tweetSchedules) append(e EventData, t time.Time, msg string) {
	h := hasher.Sum([]byte(fmt.Sprintf("%v%v", e, t)))
	if !s.already(h) {
		*s = append(*s,
			tweetSchedule{
				Event:    e,
				Time:     t,
				Message:  msg,
				Executed: false,
				Hash:     h,
			})
		log.Printf("Schedule append : %v\n%v\n", t.In(jst), msg)
	} else {
		log.Printf("Schedule append skip : %v\n%v\n", t.In(jst), msg)
	}
}

func (s *tweetSchedules) already(hash []byte) bool {
	for _, t := range *s {
		if reflect.DeepEqual(t.Hash, hash) {
			return true
		}
		continue
	}
	return false
}
