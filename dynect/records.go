package dynect

type AllRecordsResponse struct {
	ResponseBlock
	Data []string `json:"data"`
}

/*
The base struct for record data returned from the Dyn REST API.

It should never be directly passed to the *Client.Do() function for marshaling
response data to. Instead, it should aid in the composition of a more-specific
response struct.
*/
type BaseRecord struct {
	FQDN string `json:"fqdn"`
	RecordId string `json:"record_id"`
	RecordType string `json:"record_type"`
	TTL string `json:"ttl"`
	Zone string `json:"zone"`
}

/*
Represents an A record.
*/
type ARecord struct {
	BaseRecord
	RData ARDataBlock `json:"rdata"`
}

/*
Represents an AAAA record.

It uses the ADataBlock struct for is nested record data, as there is zero
difference in the structure between an A record and an AAAA record; the
only difference is in the value of the data, itself.
*/
type AAAARecord struct {
	BaseRecord
	RData ARDataBlock `json:"rdata"`
}

type ARDataBlock struct {
	Address string `json:"address"`
}
