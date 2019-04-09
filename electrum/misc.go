package electrum

/* TODO:
 * - masternode.announce.broadcast
 * - masternode.list
 * - protx.diff
 * - protx.info
 */

type basicResp struct {
	Result string `json:"result"`
}

// GetFee returns the estimated transaction fee per kilobytes for a transaction
// to be confirmed within a target number of blocks.
// https://electrumx.readthedocs.io/en/latest/protocol-methods.html#blockchain-estimatefee
func (s *Server) GetFee(target uint32) (float32, error) {
	resp := &struct {
		Result float32 `json:"result"`
	}{}

	err := s.request("blockchain.estimatefee", []interface{}{target}, resp)
	if err != nil {
		return -1, err
	}

	return resp.Result, err
}

// GetRelayFee returns the minimum fee a transaction must pay to be accepted into the
// remote server memory pool.
// https://electrumx.readthedocs.io/en/latest/protocol-methods.html#blockchain-relayfee
func (s *Server) GetRelayFee() (float32, error) {
	resp := &struct {
		Result float32 `json:"result"`
	}{}

	err := s.request("blockchain.relayfee", []interface{}{}, resp)
	if err != nil {
		return -1, err
	}

	return resp.Result, err
}

// GetFeeHistogram returns a histogram of the fee rates paid by transactions in the
// memory pool, weighted by transacation size.
// https://electrumx.readthedocs.io/en/latest/protocol-methods.html#mempool-get-fee-histogram
func (s *Server) GetFeeHistogram() (map[uint32]uint32, error) {
	resp := &struct {
		Result map[uint32]uint32 `json:"result"`
	}{}

	err := s.request("mempool.get_fee_histogram", []interface{}{}, resp)
	if err != nil {
		return nil, err
	}

	return resp.Result, err
}
