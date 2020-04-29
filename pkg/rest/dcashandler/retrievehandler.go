/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package dcashandler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hyperledger/fabric/common/flogging"
	dcas "github.com/trustbloc/fabric-peer-ext/pkg/collections/offledger/dcas/client"
	"github.com/trustbloc/sidetree-core-go/pkg/restapi/common"
	"github.com/trustbloc/sidetree-fabric/pkg/httpserver"
)

var logger = flogging.MustGetLogger("sidetree_peer")

const (
	hashParam    = "hash"
	maxSizeParam = "max-size"
)

// Retrieve manages file retrievals from the DCAS store
type Retrieve struct {
	Config
	path         string
	params       map[string]string
	channelID    string
	dcasProvider dcasClientProvider
}

type dcasClientProvider interface {
	ForChannel(channelID string) (dcas.DCAS, error)
}

// NewRetrieveHandler returns a new Retrieve handler
func NewRetrieveHandler(channelID string, cfg Config, dcasProvider dcasClientProvider) *Retrieve {
	return &Retrieve{
		Config:       cfg,
		path:         fmt.Sprintf("%s/{%s}", cfg.BasePath, hashParam),
		params:       map[string]string{maxSizeParam: fmt.Sprintf("{%s:[0-9]+}", maxSizeParam)},
		dcasProvider: dcasProvider,
		channelID:    channelID,
	}
}

// Path returns the context path
func (h *Retrieve) Path() string {
	return h.path
}

// Params returns the accepted parameters
func (h *Retrieve) Params() map[string]string {
	return h.params
}

// Method returns the HTTP method
func (h *Retrieve) Method() string {
	return http.MethodGet
}

// Handler returns the request handler
func (h *Retrieve) Handler() common.HTTPRequestHandler {
	return h.retrieve
}

// version retrieves the content from the DCAS store by hash
func (h *Retrieve) retrieve(rw http.ResponseWriter, req *http.Request) {
	hash := getHash(req)

	maxSize := getMaxSize(req)

	logger.Debugf("[%s:%s:%s] Retrieving resp for hash [%s] with max-size %d", h.channelID, h.ChaincodeName, h.Collection, hash, maxSize)

	rrw := newRetrieveWriter(rw)

	content, err := h.doRetrieve(hash, maxSize)
	if err != nil {
		rrw.WriteError(err)
		return
	}

	logger.Debugf("[%s:%s:%s] ... retrieved content for hash [%s]: Content: %s", h.channelID, h.ChaincodeName, h.Collection, hash, content)

	rrw.Write(content)
}

func (h *Retrieve) doRetrieve(hash string, maxSize int) ([]byte, error) {
	if hash == "" {
		return nil, newRetrieveError(http.StatusBadRequest, CodeInvalidHash)
	}

	if maxSize == 0 {
		return nil, newRetrieveError(http.StatusBadRequest, CodeMaxSizeNotSpecified)
	}

	content, err := h.retrieveContent(hash)
	if err != nil {
		return nil, err
	}

	if maxSize > 0 && len(content) > maxSize {
		return nil, newRetrieveError(http.StatusBadRequest, CodeMaxSizeExceeded)
	}

	return content, nil
}

func (h *Retrieve) retrieveContent(hash string) ([]byte, error) {
	dcasClient, err := h.dcasProvider.ForChannel(h.channelID)
	if err != nil {
		logger.Errorf("[%s:%s:%s] Could not get DCAS client: %s", h.channelID, h.ChaincodeName, h.Collection, err)

		return nil, newRetrieveError(http.StatusInternalServerError, CodeCasNotReachable)
	}

	content, err := dcasClient.Get(h.ChaincodeName, h.Collection, hash)
	if err != nil {
		logger.Errorf("[%s:%s:%s] Error retrieving DCAS document for hash [%s]: %s", h.channelID, h.ChaincodeName, h.Collection, hash, err)

		return nil, newRetrieveError(http.StatusInternalServerError, CodeCasNotReachable)
	}

	if len(content) == 0 {
		logger.Debugf("[%s:%s:%s] Content not found in DCAS for hash [%s]", h.channelID, h.ChaincodeName, h.Collection, hash)

		return nil, newRetrieveError(http.StatusNotFound, CodeNotFound)
	}

	return content, nil
}

var getHash = func(req *http.Request) string {
	return mux.Vars(req)[hashParam]
}

var getMaxSize = func(req *http.Request) int {
	return maxSizeFromString(mux.Vars(req)[maxSizeParam])
}

func maxSizeFromString(str string) int {
	if str == "" {
		return 0
	}

	size, err := strconv.Atoi(str)
	if err != nil {
		logger.Debugf("Invalid value for parameter [max-size]: %s", err)

		return 0
	}

	return size
}

type retrieveWriter struct {
	*httpserver.ResponseWriter
}

func newRetrieveWriter(rw http.ResponseWriter) *retrieveWriter {
	return &retrieveWriter{
		ResponseWriter: httpserver.NewResponseWriter(rw),
	}
}

func (rw *retrieveWriter) Write(content []byte) {
	rw.ResponseWriter.Write(http.StatusOK, content, httpserver.ContentTypeBinary)
}

func (rw *retrieveWriter) WriteError(err error) {
	readErr, ok := err.(*retrieveError)
	if ok {
		rw.ResponseWriter.Write(readErr.Status(), []byte(readErr.ResultCode()), httpserver.ContentTypeText)
		return
	}

	rw.ResponseWriter.WriteError(err)
}
