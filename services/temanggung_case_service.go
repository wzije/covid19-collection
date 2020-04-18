package services

import (
	"context"
	"github.com/wzije/covid19-collection/configs"
	"github.com/wzije/covid19-collection/data"
	"github.com/wzije/covid19-collection/domains"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"strings"
)

func TemanggungCaseCollection() *mongo.Collection {
	db, err := configs.MongoDB()
	if err != nil {
		log.Fatal(err.Error())
	}

	return db.Collection("temanggung_cases")
}

//get latest cases
func GetLatestTemanggungCases() (data.TemanggungCasePlain, error) {

	ctx := context.Background()

	pipeline := make([]bson.M, 0)
	err := bson.UnmarshalExtJSON([]byte(strings.TrimSpace(`
		[
			{ "$group": {
				"_id": "$id",
				"area" : { "$last": "$area" },
				"createdat": { "$last": "$createdat" },
				"updatedat": { "$last": "$updatedat" },
				"confirmed": { "$last": "$confirmed" },
				"deaths": { "$last": "$deaths" },
				"recovered": { "$last": "$recovered" },
				"odp": { "$last": "$odp" },
				"pdp": { "$last": "$pdp" }
			} }
		]
	`)), true, &pipeline)

	if err != nil {
		return data.TemanggungCasePlain{}, err
	}

	csr, err := TemanggungCaseCollection().Aggregate(ctx, pipeline)

	if err != nil {
		return data.TemanggungCasePlain{}, err
	}

	defer csr.Close(ctx)

	cases := make([]domains.TemanggengCase, 0)

	if err = csr.All(ctx, &cases); err != nil {
		panic(err)
	}

	var totalConfirmed int = 0
	var totalDeaths int = 0
	var totalRecovered int = 0
	var totalODP int = 0
	var totalPDP int = 0

	for i, _ := range cases {
		cases[i].CreatedAt = cases[i].CreatedAtID()
		cases[i].UpdatedAt = cases[i].UpdatedAtID()

		totalConfirmed += cases[i].Confirmed
		totalDeaths += cases[i].Deaths
		totalRecovered += cases[i].Recovered
		totalODP += cases[i].PDP
		totalPDP += cases[i].ODP
	}

	result := data.TemanggungCasePlain{
		Cases:       cases,
		Confirm:     totalConfirmed,
		Deaths:      totalDeaths,
		Recovered:   totalRecovered,
		ODP:         totalODP,
		PDP:         totalPDP,
		LastUpdated: getLatestUpdated(),
	}

	return result, nil
}

//get all cases
func GetAllTemanggungCases() ([]domains.TemanggengCase, error) {
	ctx := context.Background()

	csr, err := ProvinceCaseCollection().
		Find(ctx,
			bson.D{})

	if err != nil {
		return nil, err
	}

	result := make([]domains.TemanggengCase, 0)

	for csr.Next(ctx) {
		var row domains.TemanggengCase
		err := csr.Decode(&row)
		if err != nil {
			log.Fatal(err.Error())
		}

		result = append(result, row)
	}

	return result, nil
}

//insert case
func StoreTemanggungCase(c domains.TemanggengCase) {
	_, err := TemanggungCaseCollection().
		InsertOne(context.Background(), c)

	if err != nil {
		log.Fatal(err.Error())
	}
}
