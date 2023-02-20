package recisub

import (
	"sync"
	"sync/atomic"

	"github.com/vela-ssoc/backend-common/model"
	"github.com/vela-ssoc/vela-manager/inward/evtrsk"
	"gorm.io/gorm"
)

func Subscribe(db *gorm.DB) evtrsk.Subscriber {
	return &recSub{db: db}
}

type recSub struct {
	db     *gorm.DB
	mutex  sync.Mutex
	done   atomic.Bool
	events model.RecipientMap
	risks  model.RecipientMap
}

func (rs *recSub) Unset() {
	rs.mutex.Lock()
	rs.done.Store(false)
	rs.mutex.Unlock()
}

func (rs *recSub) Event(evt *model.Event) model.Recipients {
	events := rs.loadEvents()
	if events == nil {
		return nil
	}
	users := events[evt.Typeof]

	return users
	//size := len(users)
	//if size == 0 {
	//	return nil
	//}

	//ret := make(model.Recipients, size)
	//for _, user := range users {
	//	send := user.EvalEvent(evt)
	//	if send {
	//		ret = append(ret, user.rec)
	//	}
	//}

	// return ret
}

func (rs *recSub) Risk(rsk *model.Risk) model.Recipients {
	risks := rs.loadRisks()
	if risks == nil {
		return nil
	}
	users := risks[rsk.RiskType]

	return users
	//size := len(users)
	//if size == 0 {
	//	return nil
	//}

	//ret := make(model.Recipients, size)
	//for _, user := range users {
	//	send := user.EvalRisk(rsk)
	//	if send {
	//		ret = append(ret, user.rec)
	//	}
	//}

	// return ret
}

func (rs *recSub) loadEvents() model.RecipientMap {
	if rs.done.Load() {
		return rs.events
	}

	events, _ := rs.find()

	return events
}

func (rs *recSub) loadRisks() model.RecipientMap {
	if rs.done.Load() {
		return rs.risks
	}

	_, risks := rs.find()

	return risks
}

func (rs *recSub) find() (model.RecipientMap, model.RecipientMap) {
	rs.mutex.Lock()
	defer rs.mutex.Unlock()

	if rs.done.Load() {
		return rs.events, rs.risks
	}
	defer rs.done.Store(true)

	var recs model.Recipients
	rs.db.Find(&recs)
	events, risks := recs.Classify()
	rs.events = events
	rs.risks = risks
	//events := make(map[string][]*subUser, len(evts))
	//risks := make(map[string][]*subUser, len(rsks))
	//
	//for key, val := range evts {
	//	users := events[key]
	//	events[key] = append(users, &subUser{rec: val})
	//}

	return events, risks
}
