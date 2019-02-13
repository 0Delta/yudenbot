package main

import (
	"crypto/md5"
	"fmt"
	"log"
	"reflect"

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
