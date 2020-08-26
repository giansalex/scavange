package rest

// The packages below are commented out at first to prevent an error if this file isn't initially saved.
import (
	// "bytes"
	"crypto/sha256"
	"encoding/hex"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/giansalex/scavenge/x/scavenge/types"
)

func registerTxRoutes(cliCtx context.CLIContext, r *mux.Router) {
	r.HandleFunc("/scavenge/", getCreateScavenge(cliCtx)).Methods("POST")
	r.HandleFunc("/scavenge/commit", getCommitScavenge(cliCtx)).Methods("POST")
	r.HandleFunc("/scavenge/reveal", getRevealScavenge(cliCtx)).Methods("POST")
	r.HandleFunc("/scavenge/delete", getDeleteScavenge(cliCtx)).Methods("DELETE")
}

type createScavengeReq struct {
	BaseReq     rest.BaseReq `json:"base_req"`
	Solution    string       `json:"solution"`
	Description string       `json:"description"`
	Reward      string       `json:"reward"`
}

func getCreateScavenge(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createScavengeReq

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		reward, err := sdk.ParseCoins(req.Reward)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		addr, _ := sdk.AccAddressFromBech32(baseReq.From)
		var solutionHash = sha256.Sum256([]byte(req.Solution))
		var solutionHashString = hex.EncodeToString(solutionHash[:])

		msg := types.NewMsgCreateScavenge(addr, req.Description, solutionHashString, reward)
		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}

type commitScavengeReq struct {
	BaseReq  rest.BaseReq `json:"base_req"`
	Solution string       `json:"solution"`
}

func getCommitScavenge(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req commitScavengeReq

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		addr, _ := sdk.AccAddressFromBech32(baseReq.From)

		var solution = req.Solution
		var solutionHash = sha256.Sum256([]byte(solution))
		var solutionHashString = hex.EncodeToString(solutionHash[:])

		var scavenger = addr.String()
		var solutionScavengerHash = sha256.Sum256([]byte(solution + scavenger))
		var solutionScavengerHashString = hex.EncodeToString(solutionScavengerHash[:])

		msg := types.NewMsgCommitSolution(addr, solutionHashString, solutionScavengerHashString)
		err := msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}

type revealScavengeReq struct {
	BaseReq  rest.BaseReq `json:"base_req"`
	Solution string       `json:"solution"`
}

func getRevealScavenge(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req revealScavengeReq

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		addr, _ := sdk.AccAddressFromBech32(baseReq.From)
		msg := types.NewMsgRevealSolution(addr, req.Solution)
		err := msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}

type deleteScavengeReq struct {
	BaseReq  rest.BaseReq `json:"base_req"`
	Solution string       `json:"solution"`
}

func getDeleteScavenge(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req deleteScavengeReq

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		addr, _ := sdk.AccAddressFromBech32(baseReq.From)
		msg := types.NewMsgDeleteScavenge(addr, req.Solution)
		err := msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}
