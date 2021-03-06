package controller

import (
	"encoding/json"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/reechou/real-activity/config"
	"github.com/reechou/real-activity/models"
)

type ActLogic struct {
	cfg *config.Config
}

func NewActLogic(cfg *config.Config) *ActLogic {
	al := &ActLogic{
		cfg: cfg,
	}
	al.init()

	return al
}

func (al *ActLogic) init() {
	models.InitDB(al.cfg)

	http.HandleFunc("/act/add_activity", al.addActivityHandler)
	http.HandleFunc("/act/add_activity_header", al.addActivityHeaderHandler)
	http.HandleFunc("/act/add_activity_items", al.addActivityItemsHandler)
	http.HandleFunc("/act/get_activity_header", al.getActivityHeaderHandler)
	http.HandleFunc("/act/get_activity_items", al.getActivityItemsHandler)
	http.HandleFunc("/act/get_activitys", al.getActivityListHandler)
	http.HandleFunc("/act/del_activity", al.delActivityHandler)
	http.HandleFunc("/act/del_item", al.delActivityItemHandler)
	http.HandleFunc("/act/del_navigation", al.delActivityNavigationHandler)
	http.HandleFunc("/act/del_banner", al.delActivityBannerHandler)
	http.HandleFunc("/act/update_navigation", al.updateActivityNavigationHandler)
	http.HandleFunc("/act/update_banner", al.updateActivityBannerHandler)
	http.HandleFunc("/act/update_activity", al.updateActivityHandler)
}

func (al *ActLogic) Run() {
	if al.cfg.Debug {
		EnableDebug()
	}

	logrus.Infof("activity server starting on[%s]..", al.cfg.Host)
	logrus.Infoln(http.ListenAndServe(al.cfg.Host, nil))
}

func WriteJSON(w http.ResponseWriter, code int, v interface{}) error {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "x-requested-with,content-type")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(v)
}

func EnableDebug() {
	logrus.SetLevel(logrus.DebugLevel)
}
