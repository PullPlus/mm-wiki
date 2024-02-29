package work

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/chaiyd/mm-wiki/app/models"
	"github.com/chaiyd/mm-wiki/app/services"
	"github.com/chaiyd/mm-wiki/app/utils"
	"sync"
	"time"
)

var (
	DocSearchWorker = NewDocSearchWork()
)

const (
	// work <LABEL_562>
	RunStatusStop = 0
	// work <LABEL_1501>
	RunStatusRunning = 1
)

type DocSearch struct {
	// <LABEL_1502>，<LABEL_130>，<LABEL_1166>
	lock sync.RWMutex
	// work <LABEL_1167>
	runStatus int
	// work <LABEL_212>
	isTaskRunning bool
	// work <LABEL_1168>
	quit chan bool
}

func NewDocSearchWork() *DocSearch {
	return &DocSearch{
		runStatus:     RunStatusStop,
		isTaskRunning: false,
		quit:          make(chan bool, 1),
	}
}

// Start <LABEL_1625> work
func (d *DocSearch) Start() {
	// <LABEL_973>
	if d.runStatus == RunStatusRunning {
		return
	}
	timer, ok := d.getFullTextSearchConf()
	if !ok {
		return
	}
	d.updateAllDocIndex()
	go func(d *DocSearch, t time.Duration) {
		defer func() {
			e := recover()
			if e != nil {
				logs.Info("[DocSearchWork] load all doc index panic: %v", e)
			}
			d.lock.Lock()
			d.runStatus = RunStatusStop
			d.isTaskRunning = false
			d.lock.Unlock()
		}()
		d.lock.Lock()
		d.runStatus = RunStatusRunning
		d.lock.Unlock()
		for {
			select {
			case <-time.Tick(t):
				if !d.isTaskRunning {
					d.updateAllDocIndex()
				}
			case <-d.quit:
				logs.Info("[DocSearchWork] stop doc index")
				return
			}
		}
	}(d, time.Duration(timer)*time.Second)
}

// Restart <LABEL_1169> work
func (d *DocSearch) Restart() {
	d.Stop()
	time.Sleep(time.Millisecond)
	d.Start()
}

// Stop <LABEL_1626> work
func (d *DocSearch) Stop() {
	d.quit <- true
}

// <LABEL_22>
func (d *DocSearch) getFullTextSearchConf() (timer int64, isOpen bool) {
	fulltextSearchOpen := models.ConfigModel.GetConfigValueByKey(models.ConfigKeyFulltextSearch, "0")
	docSearchTimer := models.ConfigModel.GetConfigValueByKey(models.ConfigKeyDocSearchTimer, "3600")
	timer = utils.Convert.StringToInt64(docSearchTimer)
	// <LABEL_1613> 3600 s
	if timer <= 0 {
		timer = int64(3600)
	}
	if fulltextSearchOpen == "1" {
		return timer, true
	}
	return timer, false
}

func (d *DocSearch) updateAllDocIndex() {

	logs.Info("[DocSearchWork] start load all doc index")

	d.lock.Lock()
	d.isTaskRunning = true
	d.lock.Unlock()

	// <LABEL_974>，<LABEL_1503> 100
	batchUpdateDocNum, _ := beego.AppConfig.Int("search::batch_update_doc_num")
	if batchUpdateDocNum <= 0 {
		batchUpdateDocNum = 100
	}
	services.DocIndexService.UpdateAllDocIndex(batchUpdateDocNum)
	services.DocIndexService.Flush()

	d.lock.Lock()
	d.isTaskRunning = false
	d.lock.Unlock()

	logs.Info("[DocSearchWork] finish all doc index flush")

}
