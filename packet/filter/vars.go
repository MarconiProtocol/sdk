package m_packet_filter

/*
  Each filter package will need to declare a var called 'Filter' that implements the Filter interface
*/
const(
	FILTER_INTERFACE_OBJ_NAME = "Filter"
)

// Either accept or drop a packet that is coming through the filter
type Action int
const(
	ACCEPT Action = iota  // 0
	DROP                  // 1
)

/*
  Response that the filter can use to pass information back to the filter manager
*/
type FilterResponse struct {
	Msg     string
	Status  Action
}