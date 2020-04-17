package services

import (
	"context"
	"github.com/wzije/covid19-collection/configs"
	"github.com/wzije/covid19-collection/data"
	"github.com/wzije/covid19-collection/domains"
	"github.com/wzije/covid19-collection/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"strings"
	"time"
)

func CaseCollection() *mongo.Collection {
	db, err := configs.DBConnect()
	if err != nil {
		log.Fatal(err.Error())
	}

	return db.Collection("cases")
}

func CaseInfoCollection() *mongo.Collection {
	db, err := configs.DBConnect()
	if err != nil {
		log.Fatal(err.Error())
	}

	return db.Collection("case_infos")
}

// --- case info ---
//get latest cases
func GetLatestCasesProvince() (data.CaseProvincePlain, error) {

	ctx := context.Background()

	//matchStage := bson.D{{"$match", bson.D{{"podcast", id}}}}
	//sortStage := bson.D{{"$sort", bson.D{{"createdat", -1}}}}

	//groupStage := bson.D{{
	//	"$group", bson.D{{"_id", "$createdat"}, {"Confirmed", bson.D{{"$sum", "$confirmed"},}}}}}

	pipeline := make([]bson.M, 0)
	err := bson.UnmarshalExtJSON([]byte(strings.TrimSpace(`
		[
			{ "$group": {
				"_id": "$id",
				"province" : { "$last": "$province" },
				"createdat": { "$last": "$createdat" },
				"updatedat": { "$last": "$updatedat" },
				"confirmed": { "$last": "$confirmed" },
				"deaths": { "$last": "$deaths" },
				"recovered": { "$last": "$recovered" },
				"total": { "$sum": "$confirmed" }
			
			} }
		]
	`)), true, &pipeline)

	if err != nil {
		return data.CaseProvincePlain{}, err
	}

	csr, err := CaseCollection().Aggregate(ctx, pipeline)

	if err != nil {
		return data.CaseProvincePlain{}, err
	}

	defer csr.Close(ctx)

	cases := make([]domains.Case, 0)

	if err = csr.All(ctx, &cases); err != nil {
		panic(err)
	}

	var totalConfirmed int = 0
	var totalDeaths int = 0
	var totalRecovered int = 0

	for i, _ := range cases {
		cases[i].CreatedAt = cases[i].CreatedAtID()
		cases[i].UpdatedAt = cases[i].UpdatedAtID()
		totalConfirmed += cases[i].Confirmed
		totalDeaths += cases[i].Deaths
		totalRecovered += cases[i].Recovered
	}

	result := data.CaseProvincePlain{
		Cases:       cases,
		Confirm:     totalConfirmed,
		Deaths:      totalDeaths,
		Recovered:   totalRecovered,
		LastUpdated: getLatestUpdated(),
	}

	return result, nil
}

//get all cases
func GetAllCasesProvince() (domains.Case, error) {
	var c domains.Case

	err := CaseCollection().
		FindOne(
			context.Background(),
			domains.Case{CreatedAt: time.Now()}).
		Decode(&c)

	if err != nil {
		return domains.Case{}, err
	}

	return c, nil
}

//insert case
func StoreCaseProvince(c domains.Case) {
	_, err := CaseCollection().
		InsertOne(context.Background(), c)

	if err != nil {
		log.Fatal(err.Error())
	}
}

//---- case info ---
func GetCaseInfos() ([]domains.CaseInfo, error) {
	ctx := context.Background()

	csr, err := CaseInfoCollection().Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	defer csr.Close(ctx)

	result := make([]domains.CaseInfo, 0)
	for csr.Next(ctx) {
		var row domains.CaseInfo
		err := csr.Decode(&row)
		if err != nil {
			log.Fatal(err.Error())
		}

		result = append(result, row)
	}

	return result, nil
}

//get latest updatet
func getLatestUpdated() time.Time {
	pipeline := make([]bson.M, 0)
	err := bson.UnmarshalExtJSON([]byte(strings.TrimSpace(`
		[
			{"$sort": {"lastupdate": -1}}
		]
	`)), true, &pipeline)

	if err != nil {
		log.Fatal(err.Error())
	}

	csr, err := CaseInfoCollection().Aggregate(context.Background(), pipeline)

	if err != nil {
		log.Fatal(err.Error())
	}

	defer csr.Close(context.Background())

	caseInfos := make([]domains.CaseInfo, 0)

	if err = csr.All(context.Background(), &caseInfos); err != nil {
		panic(err)
	}

	return caseInfos[0].LastDateID()

}

//validate is today updated
func isTodayUpdated() bool {

	ci, err := GetCaseInfos()
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, el := range ci {
		if utils.IsSameDate(time.Now(), el.LastUpdate) {
			log.Print("Same Day")
			return true
		}
	}

	return false
}

//store case info
func StoreCaseInfo() {

	now := time.Now()

	_, err := CaseInfoCollection().
		InsertOne(
			context.Background(),
			domains.CaseInfo{LastUpdate: now},
		)

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("Update case info successfully")

	//return true, nil
}
