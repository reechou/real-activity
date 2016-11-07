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
}

func (al *ActLogic) Run() {
	if al.cfg.Debug {
		EnableDebug()
	}

	logrus.Infof("activity server starting on[%s]..", al.cfg.Host)
	logrus.Infoln(http.ListenAndServe(al.cfg.Host, nil))
}

func WriteJSON(w http.ResponseWriter, code int, v interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(v)
}

func EnableDebug() {
	logrus.SetLevel(logrus.DebugLevel)
}
