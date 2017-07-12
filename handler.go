package main

import "encoding/json"

import "github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
import "github.com/gjacquet/condo-notifier/processor"
import "github.com/gjacquet/condo-notifier/condo"

type Response struct {
	Condos []*condo.CondoAd
}

func Handle(evt interface{}, ctx *runtime.Context) (string, error) {
	ads, _ := processor.ParseRss("https://www.kijiji.ca/rss-srp-appartement-condo/ville-de-montreal/c37l1700281?meuble=1")

	bytes, _ := json.MarshalIndent(Response{ads}, "", "  ")

	return string(bytes[:]), nil
}
