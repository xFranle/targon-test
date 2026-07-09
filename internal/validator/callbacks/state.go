package callbacks

import (
	"github.com/manifold-inc/targon/internal/attest"
	"github.com/manifold-inc/targon/internal/tower"
	"github.com/manifold-inc/targon/internal/validator"

	"github.com/subtrahend-labs/gobt/runtime"
)

func resetState(c *validator.Core) {
	c.Neurons = make(map[string]runtime.NeuronInfo)
	c.MinerNodes = make(map[string][]*attest.MinerNode)
	c.NodeIds = make(map[string]bool)
	// Dont really need to wipe tao price
	c.EmissionPool = nil
	c.MinerErrors = make(map[string]map[string]string)
	c.VerifiedNodes = make(map[string]map[string]*attest.UserData)
	c.Auctions = make(map[string]tower.Auction)
	c.AuctionResults = make(map[string][]*validator.MinerBid)
	c.TaoPrice = nil
}
