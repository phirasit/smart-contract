/**
 * Copyright (c) 2018, 2019 National Digital ID COMPANY LIMITED
 *
 * This file is part of NDID software.
 *
 * NDID is the free software: you can redistribute it and/or modify it under
 * the terms of the Affero GNU General Public License as published by the
 * Free Software Foundation, either version 3 of the License, or any later
 * version.
 *
 * NDID is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
 * See the Affero GNU General Public License for more details.
 *
 * You should have received a copy of the Affero GNU General Public License
 * along with the NDID source code. If not, see https://www.gnu.org/licenses/agpl.txt.
 *
 * Please contact info@ndid.co.th for any further questions
 *
 */

package did

// "encoding/json"

// "github.com/golang/protobuf/proto"
// "github.com/ndidplatform/smart-contract/v4/abci/utils"
// pbData "github.com/ndidplatform/smart-contract/v4/protos/data"
// "github.com/tendermint/tendermint/abci/types"

// FIXME: Need a more efficient way to store Tx report
// Currently, it load and unmarshal the whole report, add more item, then marshal it all back and store
// func writeBurnTokenReport(nodeID string, method string, price float64, data string, app *DIDApplication) error {
// 	key := "SpendGas" + "|" + nodeID
// 	_, chkExists := app.GetStateDB([]byte(key)))
// 	var newReport pbData.Report
// 	newReport.Method = method
// 	newReport.Price = price
// 	newReport.Data = data
// 	if chkExists != nil {
// 		var reports pbData.ReportList
// 		err := proto.Unmarshal([]byte(chkExists), &reports)
// 		if err != nil {
// 			return err
// 		}
// 		reports.Reports = append(reports.Reports, &newReport)
// 		value, err := utils.ProtoDeterministicMarshal(&reports)
// 		if err != nil {
// 			return err
// 		}
// 		app.SetStateDB([]byte(key), []byte(value))
// 	} else {
// 		var reports pbData.ReportList
// 		reports.Reports = append(reports.Reports, &newReport)
// 		value, err := utils.ProtoDeterministicMarshal(&reports)
// 		if err != nil {
// 			return err
// 		}
// 		app.SetStateDB([]byte(key), []byte(value))
// 	}
// 	return nil
// }

// func (app *DIDApplication) getUsedTokenReport(param string, height int64) types.ResponseQuery {
// 	app.logger.Infof("GetUsedTokenReport, Parameter: %s", param)
// 	var funcParam GetUsedTokenReportParam
// 	err := json.Unmarshal([]byte(param), &funcParam)
// 	if err != nil {
// 		return app.ReturnQuery(nil, err.Error(), app.state.Height)
// 	}
// 	key := "SpendGas" + "|" + funcParam.NodeID
// 	_, value := app.GetStateDBVersioned([]byte(key)), height)
// 	if value == nil {
// 		value = []byte("[]")
// 		return app.ReturnQuery(value, "not found", app.state.Height)
// 	}
// 	var result GetUsedTokenReportResult
// 	var reports pbData.ReportList
// 	err = proto.Unmarshal([]byte(value), &reports)
// 	if err != nil {
// 		return app.ReturnQuery(nil, err.Error(), app.state.Height)
// 	}
// 	for _, report := range reports.Reports {
// 		var newRow Report
// 		newRow.Method = report.Method
// 		newRow.Price = float64(report.Price)
// 		newRow.Data = report.Data
// 		result = append(result, newRow)
// 	}
// 	resultJSON, err := json.Marshal(result)
// 	if err != nil {
// 		return app.ReturnQuery(nil, err.Error(), app.state.Height)
// 	}
// 	return app.ReturnQuery(resultJSON, "success", app.state.Height)
// }
