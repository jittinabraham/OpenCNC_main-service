package nni

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	store "main-service/pkg/store-wrapper"
	devicemodelregistry "main-service/pkg/structures/device-model-registry"
	"main-service/pkg/structures/topology"

	//"git.cs.kau.se/hamzchah/opencnc_kafka-exporter/logger/pkg/logger"

	"github.com/go-openapi/runtime/middleware/header"
	"github.com/gogo/protobuf/jsonpb"
)

//var log = logger.GetLogger()

const PORT uint16 = 8000

func StartServer() {
	fmt.Println("Starting NNI server")

	//log.Infof("Starting NNI server")
	http.HandleFunc("/get_talkers", getTalkers)
	http.HandleFunc("/get_listeners", getListeners)
	http.HandleFunc("/get_streams", getStreams)
	http.HandleFunc("/get_switches", getSwitches)
	http.HandleFunc("/add_switch", addSwitch)

	http.HandleFunc("/add_topology", addTopology)
	http.HandleFunc("/add_model", addDeviceModel)

	//log.Infof("NNI endpoint -> http://localhost:%d/get_talkers", PORT)
	//log.Infof("NNI endpoint -> http://localhost:%d/get_listeners", PORT)
	//log.Infof("NNI endpoint -> http://localhost:%d/get_switches", PORT)
	//log.Infof("NNI endpoint -> http://localhost:%d/get_streams", PORT)
	//log.Infof("NNI endpoint -> http://localhost:%d/add_topology", PORT)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil); err != nil {
		//log.Errorf("Failed to listen and server on %d, with error: %v", PORT, err)
		fmt.Printf("Failed to listen and server on %d, with error: %v", PORT, err)

	}
}

func getListeners(writer http.ResponseWriter, req *http.Request) {
	//TODO
	/*nodeType := "listener"
	if err := checkHeader(req); err != nil {
		http.Error(writer, err.Error(), http.StatusUnsupportedMediaType)
		return
	}

	nodes, err := store.GetAllNodesOfType(nodeType)
	if err != nil {
		//log.Errorf("Failed getting nodes of type %v from store: %v", nodeType, err)
		http.Error(writer, fmt.Sprintf("Failed getting %vs", nodeType), http.StatusInternalServerError)
	}

	// Convert nodes to json
	jsonResp, err := convertNodesToJson(nodes)
	if err != nil {
		//log.Errorf("Failed getting %vs: %v", nodeType, err)
		http.Error(writer, fmt.Sprintf("Failed converting %vs to json (string)", nodeType), http.StatusInternalServerError)
		return
	}

	writer.Header().Add("Content-Type", "application/json; charset=utf-8")
	writer.Write([]byte(jsonResp)) */
}

func addTopology(writer http.ResponseWriter, req *http.Request) {
	if err := checkHeader(req); err != nil {
		http.Error(writer, err.Error(), http.StatusUnsupportedMediaType)
		return
	}
	//log.Info("Received add_topology request")
	fmt.Println("Received add_topology request")
	var topo topology.Topology

	err := jsonpb.Unmarshal(req.Body, &topo)

	topo.Print()

	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnsupportedTypeError

		switch {
		case errors.As(err, &syntaxError):
			msg := fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
			http.Error(writer, msg, http.StatusBadRequest)
		case errors.Is(err, io.ErrUnexpectedEOF):
			msg := "Request body contains badly-formed JSON"
			http.Error(writer, msg, http.StatusBadRequest)
		case errors.As(err, &unmarshalTypeError):
			msg := "Request body contains invalid structure"
			http.Error(writer, msg, http.StatusBadRequest)
		default:
			//log.Error(err.Error())
			fmt.Println(err.Error())

			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	/*if err = nil {
		//log.Errorf("Failed to read req.Body: %v", err)
		http.Error(writer, "Failed to read body", http.StatusBadRequest)
		return
	}*/

	// Store the topology
	err = store.StoreTopology(&topo)
	if err != nil {
		//log.Errorf("Failed storing the topology: %v", err)
		fmt.Printf("Failed storing the topology: %v", err)

		return
	}
	//log.Info("Topology stored successfully!")
	fmt.Println("Topology stored successfully!")

	writer.Header().Add("Content-Type", "application/json; charset=utf-8")
	writer.WriteHeader(http.StatusOK)
}

func addDeviceModel(writer http.ResponseWriter, req *http.Request) {
	if err := checkHeader(req); err != nil {
		http.Error(writer, err.Error(), http.StatusUnsupportedMediaType)
		return
	}
	//log.Info("Received add_model request")
	fmt.Println("Received add_model request")

	var model devicemodelregistry.DeviceModel

	err := jsonpb.Unmarshal(req.Body, &model)

	model.Print()

	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnsupportedTypeError

		switch {
		case errors.As(err, &syntaxError):
			msg := fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
			http.Error(writer, msg, http.StatusBadRequest)
		case errors.Is(err, io.ErrUnexpectedEOF):
			msg := "Request body contains badly-formed JSON"
			http.Error(writer, msg, http.StatusBadRequest)
		case errors.As(err, &unmarshalTypeError):
			msg := "Request body contains invalid structure"
			http.Error(writer, msg, http.StatusBadRequest)
		default:
			//log.Error(err.Error())
			fmt.Println(err.Error())

			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	/*if err != nil {
		//log.Errorf("Failed to read req.Body: %v", err)
		http.Error(writer, "Failed to read body", http.StatusBadRequest)
		return
	}*/

	// Store the topology
	err = store.StoreDeviceModel(&model)
	if err != nil {
		//log.Errorf("Failed storing the device model: %v", err)
		fmt.Printf("Failed storing the device model: %v", err)

		return
	}
	//log.Info("Device model stored successfully!")
	fmt.Println("Device model stored successfully!")

	writer.Header().Add("Content-Type", "application/json; charset=utf-8")
	writer.WriteHeader(http.StatusOK)
}

func getTalkers(writer http.ResponseWriter, req *http.Request) {
	//TODO
	/*nodeType := "talker"
	if err := checkHeader(req); err != nil {
		http.Error(writer, err.Error(), http.StatusUnsupportedMediaType)
		return
	}

	nodes, err := store.GetAllNodesOfType(nodeType)
	if err != nil {
		//log.Errorf("Failed getting nodes of type %v from store: %v", nodeType, err)
		http.Error(writer, fmt.Sprintf("Failed getting %vs", nodeType), http.StatusInternalServerError)
	}

	// Convert nodes to json
	jsonResp, err := convertNodesToJson(nodes)
	if err != nil {
		//log.Errorf("Failed getting %vs: %v", nodeType, err)
		http.Error(writer, fmt.Sprintf("Failed converting %vs to json (string)", nodeType), http.StatusInternalServerError)
		return
	}

	writer.Header().Add("Content-Type", "application/json; charset=utf-8")
	writer.Write([]byte(jsonResp))
	*/
}

func getStreams(writer http.ResponseWriter, req *http.Request) {
	if err := checkHeader(req); err != nil {
		http.Error(writer, err.Error(), http.StatusUnsupportedMediaType)
		return
	}

	streams, err := store.GetAllStreams()
	if err != nil {
		//log.Errorf("Failed getting streams from store: %v", err)
		http.Error(writer, "Failed getting streams from store", http.StatusInternalServerError)
	}

	jsonResp, err := convertStreamsToJson(streams)
	if err != nil {
		//log.Errorf("Failed getting streams: %v", err)
		http.Error(writer, "Failed converting streams to json (string)", http.StatusInternalServerError)
		return
	}

	writer.Header().Add("Content-Type", "application/json; charset=utf-8")
	writer.Write([]byte(jsonResp))
}

func getSwitches(writer http.ResponseWriter, req *http.Request) {
	nodeType := "switch"
	if err := checkHeader(req); err != nil {
		http.Error(writer, err.Error(), http.StatusUnsupportedMediaType)
		return
	}

	nodes, err := store.GetNodes("bridges")
	if err != nil {
		//log.Errorf("Failed getting nodes of type %v from store: %v", nodeType, err)
		fmt.Printf("Failed getting nodes of type %v from store: %v", nodeType, err)

		http.Error(writer, fmt.Sprintf("Failed getting %vs", nodeType), http.StatusInternalServerError)
	}

	// Convert nodes to json
	jsonResp, err := convertNodesToJson(nodes)
	if err != nil {
		//log.Errorf("Failed getting %vs: %v", nodeType, err)
		fmt.Printf("Failed getting %vs: %v", nodeType, err)

		http.Error(writer, fmt.Sprintf("Failed converting %vs to json (string)", nodeType), http.StatusInternalServerError)
		return
	}

	writer.Header().Add("Content-Type", "application/json; charset=utf-8")
	writer.Write([]byte(jsonResp))
}

func addSwitch(writer http.ResponseWriter, req *http.Request) {
	if err := checkHeader(req); err != nil {
		http.Error(writer, err.Error(), http.StatusUnsupportedMediaType)
		return
	}

	//TODO: Implement logic using bridge.proto
}

/////

func checkHeader(req *http.Request) error {
	if req.Header.Get("Content-Type") != "" {
		value, _ := header.ParseValueAndParams(req.Header, "Content-Type")
		if value != "application/json" {
			msg := "Content-Type header is not application/json"
			return errors.New(msg)
		}
	}

	return nil
}
