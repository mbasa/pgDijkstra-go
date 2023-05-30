package main

/*
#include "postgres.h"
#include "utils/builtins.h"
#include "fmgr.h"
#include <string.h>

#cgo darwin LDFLAGS: -ldl -Wl,-undefined,dynamic_lookup
#cgo linux  LDFLAGS: -ldl -Wl,--unresolved-symbols=ignore-in-object-files

static Datum get_arg(PG_FUNCTION_ARGS, uint i) {
	return PG_GETARG_DATUM(i);
}

static char* datum_to_cstring(Datum val) {
    return DatumGetCString(text_to_cstring((struct varlena *)val));
}

static int datum_to_int32(Datum val) {
	return DatumGetInt32(val);
}

static float8 datum_to_float(Datum val) {
	return DatumGetFloat8(val);
}

*/
import "C"
import (
	"encoding/json"
	"fmt"
	"log"
	"unsafe"

	"github.com/RyanCarrier/dijkstra"
)

type GraphType struct {
	ID     int     `json:"id"`
	Source int     `json:"source"`
	Target int     `json:"target"`
	Cost   float64 `json:"cost"`
}

func getParam(fcinfo *C.FunctionCallInfoBaseData, itemNum C.uint) C.Datum {
	return C.get_arg((*C.struct_FunctionCallInfoBaseData)(unsafe.Pointer(fcinfo)),
		itemNum)
}

func dijkstraFindPath(graphArr []GraphType, sNode int, tNode int) string {
	gg := dijkstra.NewGraph()
	edgeMap := make(map[string]int, len(graphArr)*2)

	//
	// Populating the Graph
	//
	for _, m := range graphArr {
		msource := fmt.Sprintf("%d", m.Source)
		mtarget := fmt.Sprintf("%d", m.Target)

		sIndex := gg.AddMappedVertex(msource)
		tIndex := gg.AddMappedVertex(mtarget)

		s := fmt.Sprintf("%d,%d", sIndex, tIndex)
		edgeMap[s] = m.ID

		t := fmt.Sprintf("%d,%d", tIndex, sIndex)
		edgeMap[t] = m.ID

		if m.Source != m.Target {
			gg.AddMappedArc(msource, mtarget, int64(m.Cost*1000000))
			gg.AddMappedArc(mtarget, msource, int64(m.Cost*1000000))
		}
	}

	//
	// Doing the Dijkstra Search
	//
	s, _ := gg.GetMapping(fmt.Sprintf("%d", sNode))
	t, _ := gg.GetMapping(fmt.Sprintf("%d", tNode))

	best, err := gg.Shortest(s, t)

	if err != nil {
		log.Print("Error in Search:", err)
		return ""
	}

	//
	// Formatting the Output that containse the Edge IDs
	//
	pathLen := len(best.Path)
	retVal := "["

	for i := 0; i < pathLen-2; i++ {
		s := fmt.Sprintf("%d,%d", best.Path[i], best.Path[i+1])
		retVal = retVal + fmt.Sprintf("{\"edge_id\":%d},", edgeMap[s])
	}

	st := fmt.Sprintf("%d,%d", best.Path[pathLen-2], best.Path[pathLen-1])
	retVal = retVal + fmt.Sprintf("{\"edge_id\":%d}]", edgeMap[st])

	return retVal
}

//export pgDijkstraGo
func pgDijkstraGo(fcinfo *C.FunctionCallInfoBaseData) *C.text {

	a := C.datum_to_cstring(getParam(fcinfo, 0))
	b := C.datum_to_int32(getParam(fcinfo, 1))
	c := C.datum_to_int32(getParam(fcinfo, 2))

	var GraphArr []GraphType

	err := json.Unmarshal([]byte(C.GoString(a)), &GraphArr)

	if err != nil {
		log.Printf("Error Parsing JSON data", err)
		return C.cstring_to_text(C.CString(""))
	}

	return C.cstring_to_text(C.CString(
		dijkstraFindPath(GraphArr, int(b), int(c))))
}

func main() {
}
