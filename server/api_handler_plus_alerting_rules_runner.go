package chserver

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"

	rportplus "github.com/IOTech17/neo-rport/plus"
	alertingcap "github.com/IOTech17/neo-rport/plus/capabilities/alerting"
	"github.com/IOTech17/neo-rport/plus/capabilities/alerting/entities/rundata"
	"github.com/IOTech17/neo-rport/server/api"
	"github.com/IOTech17/neo-rport/server/routes"
)

func (al *APIListener) getAlertingCapability() (capEx alertingcap.CapabilityEx, statusCode int, err error) {
	plusManager := al.Server.plusManager
	if plusManager == nil {
		return nil, http.StatusUnauthorized, rportplus.ErrPlusNotAvailable
	}

	capEx = plusManager.GetAlertingCapabilityEx()
	if capEx == nil {
		return nil, http.StatusForbidden, rportplus.ErrCapabilityNotAvailable(rportplus.PlusAlertingCapability)
	}

	return capEx, 0, nil
}

func (al *APIListener) handleTestRules(w http.ResponseWriter, r *http.Request) {
	asCap, status, err := al.getAlertingCapability()
	if err != nil {
		al.jsonErrorResponse(w, status, err)
		return
	}

	runData := rundata.RunData{}

	err = parseRequestBody(r.Body, &runData)
	if err != nil {
		al.jsonError(w, err)
		return
	}

	ctx := context.Background()

	results, errs, err := asCap.RunRulesTest(ctx, &runData, al.Logger)

	if err != nil {
		if errs != nil {
			errPayload := makeValidationErrorPayload(errs)
			al.writeJSONResponse(w, http.StatusBadRequest, errPayload)
			return
		}
		al.jsonErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	response := api.NewSuccessPayload(results)

	al.writeJSONResponse(w, http.StatusOK, response)
}

func (al *APIListener) handleGetSampleData(w http.ResponseWriter, r *http.Request) {
	as, status, err := al.getAlertingService()
	if err != nil {
		al.jsonErrorResponse(w, status, err)
		return
	}

	vars := mux.Vars(r)
	choice := vars[routes.ParamSampleDataChoice]

	sampleData, err := as.GetSampleData(choice)
	if err != nil {
		al.jsonErrorResponse(w, status, err)
		return
	}

	response := api.NewSuccessPayload(sampleData)

	al.writeJSONResponse(w, http.StatusOK, response)
}
