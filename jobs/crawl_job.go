package jobs

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"github.com/wzije/covid19-collection/services"
	"os"
	"time"
)

var loc, _ = time.LoadLocation(os.Getenv("TZ"))

func jobInfo() {
	fmt.Printf("run cron job at %q \n", time.Now())
}

func RunJob() {

	s1 := gocron.NewScheduler(loc)

	_, _ = s1.Every(1).Minute().Do(jobInfo)
	_, _ = s1.Every(1).Day().At("16:00").Do(services.CrawlAll)

	// scheduler starts running jobs and current thread continues to execute
	s1.Start()
}
