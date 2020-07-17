// generated code; DO NOT EDIT; 2020-07-17 16:25:13.778601 +0200 CEST m=+0.032530623"
//
// Client for Choria RPC Agent 'scout'' Version 0.0.1 generated using Choria version 0.14.0

package scoutclient

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/choria-io/go-choria/protocol"
	rpcclient "github.com/choria-io/go-choria/providers/agent/mcorpc/client"
)

// GossValidateRequester performs a RPC request to scout#goss_validate
type GossValidateRequester struct {
	r    *requester
	outc chan *GossValidateOutput
}

// GossValidateOutput is the output from the goss_validate action
type GossValidateOutput struct {
	details *ResultDetails
	reply   map[string]interface{}
}

// GossValidateResult is the result from a goss_validate action
type GossValidateResult struct {
	stats   *rpcclient.Stats
	outputs []*GossValidateOutput
}

// Stats is the rpc request stats
func (d *GossValidateResult) Stats() Stats {
	return d.stats
}

// ResultDetails is the details about the request
func (d *GossValidateOutput) ResultDetails() *ResultDetails {
	return d.details
}

// HashMap is the raw output data
func (d *GossValidateOutput) HashMap() map[string]interface{} {
	return d.reply
}

// JSON is the JSON representation of the output data
func (d *GossValidateOutput) JSON() ([]byte, error) {
	return json.Marshal(d.reply)
}

// ParseOutput parses the result value from the GossValidate action into target
func (d *GossValidateOutput) ParseGossValidateOutput(target interface{}) error {
	j, err := d.JSON()
	if err != nil {
		return fmt.Errorf("could not access payload: %s", err)
	}

	err = json.Unmarshal(j, target)
	if err != nil {
		return fmt.Errorf("could not unmarshal JSON payload: %s", err)
	}

	return nil
}

// Do performs the request
func (d *GossValidateRequester) Do(ctx context.Context) (*GossValidateResult, error) {
	dres := &GossValidateResult{}

	handler := func(pr protocol.Reply, r *rpcclient.RPCReply) {
		output := &GossValidateOutput{
			reply: make(map[string]interface{}),
			details: &ResultDetails{
				sender:  pr.SenderID(),
				code:    int(r.Statuscode),
				message: r.Statusmsg,
				ts:      pr.Time(),
			},
		}

		err := json.Unmarshal(r.Data, &output.reply)
		if err != nil {
			d.r.client.errorf("Could not decode reply from %s: %s", pr.SenderID(), err)
		}

		if d.outc != nil {
			d.outc <- output
			return
		}

		dres.outputs = append(dres.outputs, output)
	}

	res, err := d.r.do(ctx, handler)
	if err != nil {
		return nil, err
	}

	dres.stats = res

	return dres, nil
}

// EachOutput iterates over all results received
func (d *GossValidateResult) EachOutput(h func(r *GossValidateOutput)) {
	for _, resp := range d.outputs {
		h(resp)
	}
}

// Vars is an optional input to the goss_validate action
//
// Description: Path to a file to use as template variables
func (d *GossValidateRequester) Vars(v string) *GossValidateRequester {
	d.r.args["vars"] = v

	return d
}

// Failures is the value of the failures output
//
// Description: The number of tests that failed
func (d *GossValidateOutput) Failures() int64 {
	val := d.reply["failures"]
	return val.(int64)
}

// Results is the value of the results output
//
// Description: The full test results
func (d *GossValidateOutput) Results() []interface{} {
	val := d.reply["results"]
	return val.([]interface{})
}

// Runtime is the value of the runtime output
//
// Description: The time it took to run the tests, in seconds
func (d *GossValidateOutput) Runtime() int64 {
	val := d.reply["runtime"]
	return val.(int64)
}

// Success is the value of the success output
//
// Description: Indicates if the test passed
func (d *GossValidateOutput) Success() string {
	val := d.reply["success"]
	return val.(string)
}

// Summary is the value of the summary output
//
// Description: A human friendly test result
func (d *GossValidateOutput) Summary() string {
	val := d.reply["summary"]
	return val.(string)
}

// Tests is the value of the tests output
//
// Description: The number of tests that were run
func (d *GossValidateOutput) Tests() int64 {
	val := d.reply["tests"]
	return val.(int64)
}