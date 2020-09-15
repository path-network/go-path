package path

// Acknowledgement is a response type returned from HTTP DELETE operations to indicate whether or not the server
// has acknowledged the deletion request of the specified resource
type Acknowledgement struct {
	Acknowledged bool `json:"acknowledged"`
}
