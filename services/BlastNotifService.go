package services

import (
	"encoding/json"
	"fmt"
	"log"
	"rose-be-go/auth"
	"rose-be-go/constants"
	"rose-be-go/db"
	"rose-be-go/host/blast_notif"
	"rose-be-go/models"
	"rose-be-go/models/blastnotifmodels"
	"rose-be-go/models/dbmodels"
	"rose-be-go/models/dto"
	"time"
)

// BlastNotifService struct
type BlastNotifService struct {
	BlastNotifHost       *blast_notif.BlastNotifHost
	BlastNotifRepository *db.BlastNotifRepository
}

// InitBlastNotifService ...
func InitBlastNotifService() *BlastNotifService {
	return &BlastNotifService{
		BlastNotifHost:       blast_notif.InitBlastNotifHost(),
		BlastNotifRepository: db.InitBlastNotifRepository(),
	}
}

// SendAll ...
func (svc *BlastNotifService)SendAll(req dto.ReqBlastNotificationSendAllDto, res *models.Response)  {
	fmt.Println("<< BlastNotifService - SendAll >>")

	reqBlast := blastnotifmodels.ReqNotifAll{
		Tilte: req.Title,
		Desc: req.Desc,
		Target: "",
	}

	blastNotif := dbmodels.BlastNotif{
		Title: req.Title,
		Desc: req.Desc,
		CreatedBy: auth.UserLogin.Name,
		CreatedAt: time.Now(),
	}

	byteRes, err := svc.BlastNotifHost.Send(reqBlast, constants.BLAST_NOTIF_SEND_ALL)
	if err != nil {
		res.ErrCode = constants.EC_FAILED_BLAST_NOTIF
		res.ErrDesc = constants.EC_FAILED_BLAST_NOTIF_DESC
		blastNotif.Status = err.Error()
		svc.BlastNotifRepository.Save(&blastNotif)

		return

	}
	log.Println(string(byteRes))
	var resNotif []blastnotifmodels.ResNotifAll
	json.Unmarshal(byteRes, &resNotif)

	if resNotif[0].Rc != constants.ERR_SUCCESS {
		res.ErrCode = constants.EC_FAILED_BLAST_NOTIF
		res.ErrDesc = constants.EC_FAILED_BLAST_NOTIF_DESC
		blastNotif.Status = resNotif[0].Msg
		svc.BlastNotifRepository.Save(&blastNotif)

		return
	}

	blastNotif.Status = "SUCCESS"
	svc.BlastNotifRepository.Save(&blastNotif)


	res.ErrCode = constants.ERR_SUCCESS
	res.ErrDesc = constants.ERR_SUCCESS_MSG



}

