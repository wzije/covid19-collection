package services

import (
	"errors"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/wzije/covid19-collection/domains"
	"github.com/wzije/covid19-collection/utils"
	"net/http"
	"strings"
	"time"
)

const urlKompas string = "https://www.kompas.com/covid-19"
const urlTemanggung string = "https://corona.temanggungkab.go.id"
const userAgent string = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.113 Safari/537.36"

func CollectAll() error {
	fmt.Printf("start crawl \n")

	if isTodayUpdated() {
		fmt.Printf("sudah update hari ini \n")
		return errors.New("hari ini sudah update")
	} else {
		collectProvince()
		collectTemanggung()
		StoreCaseInfo()
		return nil
	}
}

func CollectProvince() {
	collectProvince()
}

func CollectTemanggung() {
	collectTemanggung()
}

func collectProvince() {

	cl := colly.NewCollector(
		colly.UserAgent(userAgent),
		colly.AllowURLRevisit(),
		// Allow crawling to be done in parallel / async
		colly.Async(true),
		colly.MaxDepth(2),
	)

	cl.Limit(&colly.LimitRule{
		// Filter domains affected by this rule
		DomainGlob: "*",
		// Set a delay between requests to these domains
		Delay:       1 * time.Second,
		Parallelism: 2,
		// Add an additional random delay
		RandomDelay: 1 * time.Second,
	})

	cl.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting => ", r.URL.String())
	})

	cl.OnHTML(".covid__table",
		func(e *colly.HTMLElement) {
			e.ForEach(".covid__row",
				func(rowIdx int, e *colly.HTMLElement) {
					province := e.ChildText(".covid__prov")
					confirmed := strings.Trim(e.ChildText(".-odp"), "Terkonfirmasi: ")
					deaths := strings.Trim(e.ChildText(".-gone"), "Meninggal: ")
					recovered := strings.Trim(e.ChildText(".-health"), "Sembuh: ")

					fmt.Printf(
						"Provinsi: %a \n Positif: %b \n Meninggal: %cl \n Sehat: %d \n ",
						province, confirmed, deaths, recovered)

					//Insert data to db
					StoreProvinceCase(domains.ProvinceCase{
						ID:        rowIdx + 1,
						Confirmed: utils.StringToInt(confirmed),
						Deaths:    utils.StringToInt(deaths),
						Recovered: utils.StringToInt(recovered),
						Province:  province,
						CreatedAt: utils.Time().Now(),
						UpdatedAt: utils.Time().Now(),
					})

					//set latest updated

				})
		})

	cl.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	//start scrap
	cl.Visit(urlKompas)

	// Wait until threads are finished
	cl.Wait()
}

func collectTemanggung() {
	cl := colly.NewCollector(
		colly.UserAgent(userAgent),
		colly.AllowURLRevisit(),
		colly.Async(true),
		colly.MaxDepth(2),
	)

	cl.WithTransport(&http.Transport{
		DisableKeepAlives: true,
	})

	cl.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Delay:       2 * time.Second,
		Parallelism: 2,
		RandomDelay: 2 * time.Second,
	})

	cl.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting => ", r.URL.String())
	})

	cl.OnHTML("#sebaran",
		func(e *colly.HTMLElement) {
			fmt.Print("------- start roll ---------")
			e.ForEach("table tbody tr",
				func(trIdx int, e *colly.HTMLElement) {

					var area string
					var odp int
					var pdp int
					var confirmed int
					var recovered int
					var deaths int

					e.ForEach("td",
						func(tdIdx int, td *colly.HTMLElement) {
							switch tdIdx {
							case 1:
								area = strings.TrimSpace(td.Text)
								break
							case 2:
								odp = utils.StringToInt(td.Text)
								break
							case 3:
								pdp = utils.StringToInt(td.Text)
								break
							case 4:
								confirmed = utils.StringToInt(td.Text)
								break
							case 5:
								recovered = utils.StringToInt(td.Text)
								break
							case 6:
								deaths = utils.StringToInt(td.Text)
								break
							}
						})

					StoreTemanggungCase(
						domains.TemanggengCase{
							ID:        trIdx + 1,
							Area:      area,
							ODP:       odp,
							PDP:       pdp,
							Confirmed: confirmed,
							Recovered: recovered,
							Deaths:    deaths,
							CreatedAt: utils.Time().Now(),
							UpdatedAt: utils.Time().Now(),
						})

				})
		})

	cl.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	//start crawl
	cl.Visit(urlTemanggung)

	cl.Wait()
}
