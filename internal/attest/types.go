package shared

type AttestationResponse struct {
	Quote    string       `json:"quote"`
	UserData UserData     `json:"user_data"`
	TPM      *TPMEvidence `json:"tpm,omitempty"`
}

type TPMEvidence struct {
	AKPubHex          string            `json:"ak_pub_hex"`
	QuoteHex          string            `json:"quote_hex"`
	QuoteSignatureHex string            `json:"quote_signature_hex"`
	PCRsHex           map[string]string `json:"pcrs_hex"`
}

type UserData struct {
	GPUCards     *Cards        `json:"gpu_cards,omitempty"`
	CPUCards     *Cards        `json:"cpu_cards,omitempty"`
	NodeType     string        `json:"node_type"`
	NVCCResponse *NVCCResponse `json:"attestation,omitempty"`
	AuctionName  string        `json:"auction_name"`
	Nonce        string        `json:"nonce"`
	CVMID        string        `json:"cvm_id"`
	QuoteType    string        `json:"quote_type"`
	RAMBytes     int64         `json:"ram_bytes,omitempty"`
	StorageBytes int64         `json:"storage_bytes,omitempty"`
	Version      int           `json:"version,omitempty"`
}

type NVCCResponse struct {
	GPURemote struct {
		AttestationResult bool   `json:"attestation_result"`
		Token             string `json:"token"`
		Valid             bool   `json:"valid"`
	} `json:"gpu_remote"`
	SwitchRemote struct {
		AttestationResult bool   `json:"attestation_result"`
		Token             string `json:"token"`
		Valid             bool   `json:"valid"`
	} `json:"switch_remote"`
}

type Cards []string
