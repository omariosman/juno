package internal

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/NethermindEth/juno/internal/rpc"
	"github.com/NethermindEth/juno/internal/rpc/starknet"
	"github.com/golang/mock/gomock"

	"github.com/stretchr/testify/assert"

	"github.com/NethermindEth/juno/internal/db"
)

func initGetBlockWithTxs(t *testing.T) (string, rpc.HttpRpc, *MockBlockManager, *MockTransactionManager) {
	blockManager := NewMockBlockManager(gomock.NewController(t))
	transactionManager := NewMockTransactionManager(gomock.NewController(t))
	rpcHandler, err := starknet.New(nil, blockManager, transactionManager, nil, nil, nil)
	if err != nil {
		t.Fatal(err)
	}

	var (
		addr      = "localhost:6060"
		pattern   = "/rpc"
		url       = "http://" + addr + pattern
		rpcServer = rpc.NewHttpRpc(addr, pattern, rpcHandler, nil)
	)

	return url, rpcServer, blockManager, transactionManager
}

func TestRpc_GetBlockWithTxs(t *testing.T) {
	t.Run("without params", func(t *testing.T) {
		url, rpcServer, _, _ := initGetBlockWithTxs(t)
		request := newRequest(url, `{"jsonrpc": "2.0", "method": "starknet_getBlockWithTxs", "id": 1}`)
		response := httptest.NewRecorder()
		rpcServer.ServeHTTP(response, request)

		body, _ := ioutil.ReadAll(response.Body)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, ErrInvalidParams, string(body))
	})

	t.Run("with empty param list", func(t *testing.T) {
		url, rpcServer, _, _ := initGetBlockWithTxs(t)
		request := newRequest(url, `{"jsonrpc": "2.0", "id": 1, "method": "starknet_getBlockWithTxs", "params": []}`)
		response := httptest.NewRecorder()
		rpcServer.ServeHTTP(response, request)

		body, _ := ioutil.ReadAll(response.Body)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, ErrInvalidBlockId, string(body))
	})

	t.Run("with empty named params", func(t *testing.T) {
		url, rpcServer, _, _ := initGetBlockWithTxs(t)
		request := newRequest(url, `{"jsonrpc": "2.0", "id": 1, "method": "starknet_getBlockWithTxs", "params": {}}`)
		response := httptest.NewRecorder()
		rpcServer.ServeHTTP(response, request)

		body, _ := ioutil.ReadAll(response.Body)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, ErrInvalidBlockId, string(body))
	})

	t.Run("param list with more than one param", func(t *testing.T) {
		url, rpcServer, _, _ := initGetBlockWithTxs(t)
		request := newRequest(url, `{"jsonrpc": "2.0", "id": 1, "method": "starknet_getBlockWithTxs", "params": [0,0]}`)
		response := httptest.NewRecorder()
		rpcServer.ServeHTTP(response, request)

		body, _ := ioutil.ReadAll(response.Body)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, ErrInvalidParams, string(body))
	})

	t.Run("block not found on the database", func(t *testing.T) {
		url, rpcServer, blockManager, transactionManager := initGetBlockWithTxs(t)
		defer blockManager.ctrl.Finish()
		defer transactionManager.ctrl.Finish()
		request := newRequest(url, `{
			"jsonrpc": "2.0", 
			"id": 1,
			"method": "starknet_getBlockWithTxs",
			"params": {
				"block_id": 0
			}
		}`)

		blockManager.EXPECT().
			GetBlockByNumber(uint64(0)).
			Return(nil, db.ErrNotFound)

		response := httptest.NewRecorder()
		rpcServer.ServeHTTP(response, request)

		body, _ := ioutil.ReadAll(response.Body)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, ErrInvalidBlockId, string(body))
	})

	//t.Run("block found on the database (with id number)", func(t *testing.T) {
	//	url, rpcServer, blockManager, transactionManager := initGetBlockWithTxs(t)
	//	defer blockManager.ctrl.Finish()
	//	defer transactionManager.ctrl.Finish()
	//	request := newRequest(url, `{
	//		"jsonrpc": "2.0",
	//		"id": 1,
	//		"method": "starknet_getBlockWithTxs",
	//		"params": {
	//			"block_id": 0
	//		}
	//	}`)
	//
	//	blockManager.EXPECT().
	//		GetBlockByNumber(uint64(0)).
	//		Return(stubs.Block0Mainnet.Block, nil)
	//	transactionManager.EXPECT().
	//		GetTransaction(new(felt.Felt).SetHex("0x00e0a2e45a80bb827967e096bcf58874f6c01c191e0a0530624cba66a508ae75")).
	//		Return(stubs.TransactionStubs["0x00e0a2e45a80bb827967e096bcf58874f6c01c191e0a0530624cba66a508ae75"].Txn, nil)
	//
	//	response := httptest.NewRecorder()
	//	rpcServer.ServeHTTP(response, request)
	//
	//	body, _ := ioutil.ReadAll(response.Body)
	//
	//	assert.Equal(t, http.StatusOK, response.Code)
	//	assert.JSONEq(t, stubs.Block0Mainnet.BlockWithTxs, string(body))
	//})
	//
	//t.Run("block found on the database (with id hash)", func(t *testing.T) {
	//	url, rpcServer, blockManager := initGetBlockWithTxHashes(t)
	//	defer blockManager.ctrl.Finish()
	//	request := newRequest(url, `{
	//		"jsonrpc": "2.0",
	//		"id": 1,
	//		"method": "starknet_getBlockWithTxHashes",
	//		"params": {
	//			"block_id": "0x047c3637b57c2b079b93c61539950c17e868a28f46cdef28f88521067f21e943"
	//		}
	//	}`)
	//
	//	blockManager.EXPECT().
	//		GetBlockByHash(new(felt.Felt).SetHex("0x047c3637b57c2b079b93c61539950c17e868a28f46cdef28f88521067f21e943")).
	//		Return(stubs.Block0Mainnet.Block, nil)
	//
	//	response := httptest.NewRecorder()
	//	rpcServer.ServeHTTP(response, request)
	//
	//	body, _ := ioutil.ReadAll(response.Body)
	//
	//	assert.Equal(t, http.StatusOK, response.Code)
	//	assert.JSONEq(t, stubs.Block0Mainnet.BlockWithTxHashes, string(body))
	//})
}
