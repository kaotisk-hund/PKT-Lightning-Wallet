package restrpc

import (
	"io"
	"net/http"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/mux"
	"github.com/pkt-cash/pktd/btcutil/er"
	"github.com/pkt-cash/pktd/lnd/pkthelp"
	"github.com/pkt-cash/pktd/pktlog/log"
)

const (
	helpURI_prefix = "/help"
)

//	the main help messsage
func mainREST_help() pkthelp.Method {
	return pkthelp.Method{
		Name: "pld - Lightning Network Daemon REST interface (pld)",
		Description: []string{
			"For help on a specific command, use the following URIs:",
			"	getinfo          /api/v1/help/meta/getinfo",
			"	getrecoveryinfo  /api/v1/meta/getrecoveryinfo",
			"	debuglevel       /api/v1/meta/debuglevel",
			"	stop             /api/v1/meta/stop",
			"	version          /api/v1/meta/version",
		},
	}
}

//	rest help response protobuf
type restHelpResponse struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,json=name,proto3" json:"name,omitempty"`
	Service              string   `protobuf:"bytes,2,opt,name=category,json=category,omitempty,proto3" json:"category,omitempty"`
	Description          []string `protobuf:"bytes,3,rep,name=description,json=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *restHelpResponse) Reset()         { *m = restHelpResponse{} }
func (m *restHelpResponse) String() string { return proto.CompactTextString(m) }
func (m *restHelpResponse) ProtoMessage()  {}

//	marshal rest help response
func marshalHelp(httpResponse http.ResponseWriter, helpInfo pkthelp.Method) er.R {

	marshaler := jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
		Indent:       "\t",
	}

	s, err := marshaler.MarshalToString(&restHelpResponse{
		Name:        helpInfo.Name,
		Service:     helpInfo.Service,
		Description: helpInfo.Description,
	})
	if err != nil {
		return er.E(err)
	}

	_, err = io.WriteString(httpResponse, s)
	if err != nil {
		return er.E(err)
	}

	return nil
}

//	add the main help HTTP handler
func RestHandlersHelp(router *mux.Router) {
	router.HandleFunc("/", getMainHelp)
}

//	get main help
func getMainHelp(httpResponse http.ResponseWriter, httpRequest *http.Request) {

	//	fill response payload
	if httpRequest.Method != "GET" {
		httpResponse.Header().Set("Content-Type", "text/plain")
		http.Error(httpResponse, "400 - Request should be a GET because the help endpoint requires no input", http.StatusBadRequest)
		return
	}
	err := marshalHelp(httpResponse, mainREST_help())
	if err != nil {
		httpResponse.Header().Set("Content-Type", "text/plain")
		http.Error(httpResponse, "500 - Internal Error", http.StatusInternalServerError)
		log.Errorf("Error replying to request for [%s] from [%s] - error sending error, giving up: [%s]",
			httpRequest.RequestURI, httpRequest.RemoteAddr, err)
	}
}
