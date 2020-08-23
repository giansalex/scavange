package rest

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/giansalex/scavenge/x/scavenge/types"
)

const storeName = "scavenge"

func registerQueryRoutes(cliCtx context.CLIContext, r *mux.Router) {
	r.HandleFunc("/scavenge/list", getRestListScavenges(cliCtx)).Methods("GET")
	r.HandleFunc("/scavenge/get/{solutionHash}", getRestGetScavenges(cliCtx)).Methods("GET")
	r.HandleFunc("/scavenge/commit/{scavenger}/{solution}", getRestCommitScavenges(cliCtx)).Methods("GET")
}

func getRestListScavenges(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryListScavenges)
		res, height, err := cliCtx.QueryWithData(route, nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		cliCtx = cliCtx.WithHeight(height)
		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func getRestGetScavenges(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}
		vars := mux.Vars(r)
		solutionHash := vars["solutionHash"]

		route := fmt.Sprintf("custom/%s/%s/%s", types.QuerierRoute, types.QueryGetScavenge, solutionHash)
		res, height, err := cliCtx.QueryWithData(route, nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		cliCtx = cliCtx.WithHeight(height)
		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func getRestCommitScavenges(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}
		vars := mux.Vars(r)
		solution := vars["solutionHash"]
		var solutionHash = sha256.Sum256([]byte(solution))
		var solutionHashString = hex.EncodeToString(solutionHash[:])

		var scavenger = vars["scavenger"]

		var solutionScavengerHash = sha256.Sum256([]byte(solution + scavenger))
		var solutionScavengerHashString = hex.EncodeToString(solutionScavengerHash[:])

		route := fmt.Sprintf("custom/%s/%s/%s", types.QuerierRoute, types.QueryCommit, solutionScavengerHashString)
		res, height, err := cliCtx.QueryWithData(route, nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		cliCtx = cliCtx.WithHeight(height)
		rest.PostProcessResponse(w, cliCtx, res)
	}
}
