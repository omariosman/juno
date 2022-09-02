package stubs

import (
	"github.com/NethermindEth/juno/pkg/felt"
	"github.com/NethermindEth/juno/pkg/types"
)

type Block struct {
	Block             *types.Block
	BlockWithTxHashes string
	BlockWithTxs      string
}

var Block0Mainnet = Block{
	Block: &types.Block{
		BlockHash:    new(felt.Felt).SetHex("0x047c3637b57c2b079b93c61539950c17e868a28f46cdef28f88521067f21e943"),
		ParentHash:   new(felt.Felt).SetHex("0x0000000000000000000000000000000000000000000000000000000000000000"),
		BlockNumber:  0,
		Status:       types.BlockStatusAcceptedOnL1,
		Sequencer:    new(felt.Felt).SetHex("0x0"),
		NewRoot:      new(felt.Felt).SetHex("0x021870ba80540e7831fb21c591ee93481f5ae1bb71ff85a86ddd465be4eddee6"),
		AcceptedTime: 0,
		TimeStamp:    1637069048,
		TxCount:      18,
		TxCommitment: nil,
		TxHashes: []*felt.Felt{
			new(felt.Felt).SetHex("0x00e0a2e45a80bb827967e096bcf58874f6c01c191e0a0530624cba66a508ae75"),
			new(felt.Felt).SetHex("0x012c96ae3c050771689eb261c9bf78fac2580708c7f1f3d69a9647d8be59f1e1"),
			new(felt.Felt).SetHex("0x000ce54bbc5647e1c1ea4276c01a708523f740db0ff5474c77734f73beec2624"),
			new(felt.Felt).SetHex("0x01c924916a84ef42a3d25d29c5d1085fe212de04feadc6e88d4c7a6e5b9039bf"),
			new(felt.Felt).SetHex("0x00a66c346e273cc49510ef2e1620a1a7922135cb86ab227b86e0afd12243bd90"),
			new(felt.Felt).SetHex("0x05c71675616b49fb9d16cac8beaaa65f62dc5a532e92785055c15c825166dbbf"),
			new(felt.Felt).SetHex("0x060e05c41a6622592a2e2eff90a9f2e495296a3be9596e7bc4dfbafce00d7a6a"),
			new(felt.Felt).SetHex("0x05634f2847140263ba59480ad4781dacc9991d0365145489b27a198ebed2f969"),
			new(felt.Felt).SetHex("0x00b049c384cf75174150a2540835cc2abdcca1d3a3750298a1741a621983e35a"),
			new(felt.Felt).SetHex("0x0227f3d9d5ce6680bdf2991576c1a90aca8184ca26055bae92d16c58e3e13340"),
			new(felt.Felt).SetHex("0x0376ff82431b52ca1fbc4942de80bc1b01d8e5cd1eeab5a277b601b510f2cab2"),
			new(felt.Felt).SetHex("0x025f20c74821d84f62989a71fceef08c967837b63bae31b279a11343f10d874a"),
			new(felt.Felt).SetHex("0x02d10272a8ba726793fd15aa23a1e3c42447d7483ebb0b49df8b987590fe0055"),
			new(felt.Felt).SetHex("0x00b05ba5cd0b9e0464d2c1790ad93a159c6ef0594513758bca9111e74e4099d4"),
			new(felt.Felt).SetHex("0x04d16393d940fb4a97f20b9034e2a5e954201fee827b2b5c6daa38ec272e7c9c"),
			new(felt.Felt).SetHex("0x009e80672edd4927a79f5384e656416b066f8ef58238227ac0fcea01952b70b5"),
			new(felt.Felt).SetHex("0x0387b5b63e40d4426754895fe52adf668cf8fde2a02aa9b6d761873f31af3462"),
			new(felt.Felt).SetHex("0x04f0cdff0d72fc758413a16db2bc7580dfec7889a8b921f0fe08641fa265e997"),
		},
		EventCount:      0,
		EventCommitment: nil,
	},
	BlockWithTxHashes: `{
		"jsonrpc": "2.0",
		"result": {
			"status": "ACCEPTED_ON_L1",
			"block_hash": "0x047c3637b57c2b079b93c61539950c17e868a28f46cdef28f88521067f21e943",
			"parent_hash": "0x0000000000000000000000000000000000000000000000000000000000000000",
			"block_number": 0,
			"new_root": "0x021870ba80540e7831fb21c591ee93481f5ae1bb71ff85a86ddd465be4eddee6",
			"timestamp": 1637069048,
			"sequencer_address": "0x0000000000000000000000000000000000000000000000000000000000000000",
			"transactions": [
				"0x00e0a2e45a80bb827967e096bcf58874f6c01c191e0a0530624cba66a508ae75",
				"0x012c96ae3c050771689eb261c9bf78fac2580708c7f1f3d69a9647d8be59f1e1",
				"0x000ce54bbc5647e1c1ea4276c01a708523f740db0ff5474c77734f73beec2624",
				"0x01c924916a84ef42a3d25d29c5d1085fe212de04feadc6e88d4c7a6e5b9039bf",
				"0x00a66c346e273cc49510ef2e1620a1a7922135cb86ab227b86e0afd12243bd90",
				"0x05c71675616b49fb9d16cac8beaaa65f62dc5a532e92785055c15c825166dbbf",
				"0x060e05c41a6622592a2e2eff90a9f2e495296a3be9596e7bc4dfbafce00d7a6a",
				"0x05634f2847140263ba59480ad4781dacc9991d0365145489b27a198ebed2f969",
				"0x00b049c384cf75174150a2540835cc2abdcca1d3a3750298a1741a621983e35a",
				"0x0227f3d9d5ce6680bdf2991576c1a90aca8184ca26055bae92d16c58e3e13340",
				"0x0376ff82431b52ca1fbc4942de80bc1b01d8e5cd1eeab5a277b601b510f2cab2",
				"0x025f20c74821d84f62989a71fceef08c967837b63bae31b279a11343f10d874a",
				"0x02d10272a8ba726793fd15aa23a1e3c42447d7483ebb0b49df8b987590fe0055",
				"0x00b05ba5cd0b9e0464d2c1790ad93a159c6ef0594513758bca9111e74e4099d4",
				"0x04d16393d940fb4a97f20b9034e2a5e954201fee827b2b5c6daa38ec272e7c9c",
				"0x009e80672edd4927a79f5384e656416b066f8ef58238227ac0fcea01952b70b5",
				"0x0387b5b63e40d4426754895fe52adf668cf8fde2a02aa9b6d761873f31af3462",
				"0x04f0cdff0d72fc758413a16db2bc7580dfec7889a8b921f0fe08641fa265e997"
			]
		},
		"id": 1
	}`,
	BlockWithTxs: `{
		"jsonrpc": "2.0",
		"result": {
			"status": "ACCEPTED_ON_L1",
			"block_hash": "0x047c3637b57c2b079b93c61539950c17e868a28f46cdef28f88521067f21e943",
			"parent_hash": "0x0000000000000000000000000000000000000000000000000000000000000000",
			"block_number": 0,
			"new_root": "0x021870ba80540e7831fb21c591ee93481f5ae1bb71ff85a86ddd465be4eddee6",
			"timestamp": 1637069048,
			"sequencer_address": "0x0000000000000000000000000000000000000000000000000000000000000000",
			"transactions": [
				{
					"txn_hash": "0x00e0a2e45a80bb827967e096bcf58874f6c01c191e0a0530624cba66a508ae75",
					"class_hash": "0x010455c752b86932ce552f2b0fe81a880746649b9aee7e0d842bf3f52378f9f8",
					"version": "0x0",
					"type": "DEPLOY",
					"contract_address": "0x020cfa74ee3564b4cd5435cdace0f9c4d43b939620e4a0bb5076105df0a626c6",
					"contract_address_salt": "0x0546c86dc6e40a5e5492b782d8964e9a4274ff6ecb16d31eb09cee45a3564015",
					"constructor_calldata": [
						"0x06cf6c2f36d36b08e591e4489e92ca882bb67b9c39a3afccf011972a8de467f0",
						"0x07ab344d88124307c07b56f6c59c12f4543e9c96398727854a322dea82c73240"
					]
				},
				{
					"txn_hash": "0x012c96ae3c050771689eb261c9bf78fac2580708c7f1f3d69a9647d8be59f1e1",
					"class_hash": "0x010455c752b86932ce552f2b0fe81a880746649b9aee7e0d842bf3f52378f9f8",
					"version": "0x0",
					"type": "DEPLOY",
					"contract_address": "0x031c9cdb9b00cb35cf31c05855c0ec3ecf6f7952a1ce6e3c53c3455fcd75a280",
					"contract_address_salt": "0x0012afa0f342ece0468ca9810f0ea59f9c7204af32d1b8b0d318c4f2fe1f384e",
					"constructor_calldata": [
						"0x00cfc2e2866fd08bfb4ac73b70e0c136e326ae18fc797a2c090c8811c695577e",
						"0x05f1dd5a5aef88e0498eeca4e7b2ea0fa7110608c11531278742f0b5499af4b3"
					]
				},
				{
					"txn_hash": "0x000ce54bbc5647e1c1ea4276c01a708523f740db0ff5474c77734f73beec2624",
					"max_fee": "0x0000000000000000000000000000000000000000000000000000000000000000",
					"version": "0x0",
					"signature": [],
					"nonce": "",
					"type": "INVOKE",
					"contract_address": "0x020cfa74ee3564b4cd5435cdace0f9c4d43b939620e4a0bb5076105df0a626c6",
					"entry_point_selector": "0x012ead94ae9d3f9d2bdb6b847cf255f1f398193a1f88884a0ae8e18f24a037b6",
					"calldata": [
						"0x000000000000000000000000c84dd7fd43a7defb5b7a15c4fbbe11cbba6db1ba"
					]
				},
				{
					"txn_hash": "0x01c924916a84ef42a3d25d29c5d1085fe212de04feadc6e88d4c7a6e5b9039bf",
					"max_fee": "0x0000000000000000000000000000000000000000000000000000000000000000",
					"version": "0x0",
					"signature": [],
					"nonce": "",
					"type": "INVOKE",
					"contract_address": "0x031c9cdb9b00cb35cf31c05855c0ec3ecf6f7952a1ce6e3c53c3455fcd75a280",
					"entry_point_selector": "0x0218f305395474a84a39307fa5297be118fe17bf65e27ac5e2de6617baa44c64",
					"calldata": [
						"0x020cfa74ee3564b4cd5435cdace0f9c4d43b939620e4a0bb5076105df0a626c6",
						"0x0000000000000000000000000000000000000000000000000000000000000000"
					]
				},
				{
					"txn_hash": "0x00a66c346e273cc49510ef2e1620a1a7922135cb86ab227b86e0afd12243bd90",
					"max_fee": "0x0000000000000000000000000000000000000000000000000000000000000000",
					"version": "0x0",
					"signature": [],
					"nonce": "",
					"type": "INVOKE",
					"contract_address": "0x020cfa74ee3564b4cd5435cdace0f9c4d43b939620e4a0bb5076105df0a626c6",
					"entry_point_selector": "0x0317eb442b72a9fae758d4fb26830ed0d9f31c8e7da4dbff4e8c59ea6a158e7f",
					"calldata": [
						"0x0007dbfec95c10bbc2fd3f37a89ae6e027826134f955251d11c784a6b34fdf50",
						"0x0000000000000000000000000000000000000000000000000000000000000002",
						"0x04e7e989d58a17cd279eca440c5eaa829efb6f9967aaad89022acbe644c39b36",
						"0x0453ae0c9610197b18b13645c44d3d0a407083d96562e8752aab3fab616cecb0"
					]
				},
				{
					"txn_hash": "0x05c71675616b49fb9d16cac8beaaa65f62dc5a532e92785055c15c825166dbbf",
					"max_fee": "0x0000000000000000000000000000000000000000000000000000000000000000",
					"version": "0x0",
					"signature": [],
					"nonce": "",
					"type": "INVOKE",
					"contract_address": "0x031c9cdb9b00cb35cf31c05855c0ec3ecf6f7952a1ce6e3c53c3455fcd75a280",
					"entry_point_selector": "0x027c3334165536f239cfd400ed956eabff55fc60de4fb56728b6a4f6b87db01c",
					"calldata": [
						"0x031c9cdb9b00cb35cf31c05855c0ec3ecf6f7952a1ce6e3c53c3455fcd75a280",
						"0x0317eb442b72a9fae758d4fb26830ed0d9f31c8e7da4dbff4e8c59ea6a158e7f",
						"0x0000000000000000000000000000000000000000000000000000000000000004",
						"0x04be52041fee36ab5199771acf4b5d260d223297e588654e5c9477df2efa542a",
						"0x0000000000000000000000000000000000000000000000000000000000000002",
						"0x00299e2f4b5a873e95e65eb03d31e532ea2cde43b498b50cd3161145db5542a5",
						"0x03d6897cf23da3bf4fd35cc7a43ccaf7c5eaf8f7c5b9031ac9b09a929204175f"
					]
				},
				{
					"txn_hash": "0x060e05c41a6622592a2e2eff90a9f2e495296a3be9596e7bc4dfbafce00d7a6a",
					"max_fee": "0x0000000000000000000000000000000000000000000000000000000000000000",
					"version": "0x0",
					"signature": [],
					"nonce": "",
					"type": "INVOKE",
					"contract_address": "0x031c9cdb9b00cb35cf31c05855c0ec3ecf6f7952a1ce6e3c53c3455fcd75a280",
					"entry_point_selector": "0x0218f305395474a84a39307fa5297be118fe17bf65e27ac5e2de6617baa44c64",
					"calldata": [
						"0x020cfa74ee3564b4cd5435cdace0f9c4d43b939620e4a0bb5076105df0a626c6",
						"0x0000000000000000000000000000000000000000000000000000000000000001"
					]
				},
				{
					"txn_hash": "0x05634f2847140263ba59480ad4781dacc9991d0365145489b27a198ebed2f969",
					"max_fee": "0x0000000000000000000000000000000000000000000000000000000000000000",
					"version": "0x0",
					"signature": [],
					"nonce": "",
					"type": "INVOKE",
					"contract_address": "0x031c9cdb9b00cb35cf31c05855c0ec3ecf6f7952a1ce6e3c53c3455fcd75a280",
					"entry_point_selector": "0x019a35a6e95cb7a3318dbb244f20975a1cd8587cc6b5259f15f61d7beb7ee43b",
					"calldata": [
						"0x020cfa74ee3564b4cd5435cdace0f9c4d43b939620e4a0bb5076105df0a626c6",
						"0x05aee31408163292105d875070f98cb48275b8c87e80380b78d30647e05854d5"
					]
				},
				{
					"txn_hash": "0x00b049c384cf75174150a2540835cc2abdcca1d3a3750298a1741a621983e35a",
					"class_hash": "0x010455c752b86932ce552f2b0fe81a880746649b9aee7e0d842bf3f52378f9f8",
					"version": "0x0",
					"type": "DEPLOY",
					"contract_address": "0x06ee3440b08a9c805305449ec7f7003f27e9f7e287b83610952ec36bdc5a6bae",
					"contract_address_salt": "0x05098705e4d57a8620e5b387855ef4dc82f0ccd84b7299dc1179b87517249127",
					"constructor_calldata": [
						"0x048cba68d4e86764105adcdcf641ab67b581a55a4f367203647549c8bf1feea2",
						"0x0362d24a3b030998ac75e838955dfee19ec5b6eceb235b9bfbeccf51b6304d0b"
					]
				},
				{
					"txn_hash": "0x0227f3d9d5ce6680bdf2991576c1a90aca8184ca26055bae92d16c58e3e13340",
					"class_hash": "0x010455c752b86932ce552f2b0fe81a880746649b9aee7e0d842bf3f52378f9f8",
					"version": "0x0",
					"type": "DEPLOY",
					"contract_address": "0x0735596016a37ee972c42adef6a3cf628c19bb3794369c65d2c82ba034aecf2c",
					"contract_address_salt": "0x0060bc7461113e4af46fd52e5ecbc5c3f4be92ed7f1329d53721f9bfbc0370cc",
					"constructor_calldata": [
						"0x002f50710449a06a9fa789b3c029a63bd0b1f722f46505828a9f815cf91b31d8",
						"0x02a222e62eabe91abdb6838fa8b267ffe81a6eb575f61e96ec9aa4460c0925a2"
					]
				},
				{
					"txn_hash": "0x0376ff82431b52ca1fbc4942de80bc1b01d8e5cd1eeab5a277b601b510f2cab2",
					"max_fee": "0x0000000000000000000000000000000000000000000000000000000000000000",
					"version": "0x0",
					"signature": [],
					"nonce": "",
					"type": "INVOKE",
					"contract_address": "0x06ee3440b08a9c805305449ec7f7003f27e9f7e287b83610952ec36bdc5a6bae",
					"entry_point_selector": "0x03d7905601c217734671143d457f0db37f7f8883112abd34b92c4abfeafde0c3",
					"calldata": [
						"0x01e2cd4b3588e8f6f9c4e89fb0e293bf92018c96d7a93ee367d29a284223b6ff",
						"0x071d1e9d188c784a0bde95c1d508877a0d93e9102b37213d1e13f3ebc54a7751"
					]
				},
				{
					"txn_hash": "0x025f20c74821d84f62989a71fceef08c967837b63bae31b279a11343f10d874a",
					"class_hash": "0x010455c752b86932ce552f2b0fe81a880746649b9aee7e0d842bf3f52378f9f8",
					"version": "0x0",
					"type": "DEPLOY",
					"contract_address": "0x031c887d82502ceb218c06ebb46198da3f7b92864a8223746bc836dda3e34b52",
					"contract_address_salt": "0x063d1a6f8130958509e2e695c25b147f43f66f56bba49fddb7ee363d8f57a774",
					"constructor_calldata": [
						"0x00df28e613c065616a2e79ca72f9c1908e17b8c913972a9993da77588dc9cae9",
						"0x01432126ac23c7028200e443169c2286f99cdb5a7bf22e607bcd724efa059040"
					]
				},
				{
					"txn_hash": "0x02d10272a8ba726793fd15aa23a1e3c42447d7483ebb0b49df8b987590fe0055",
					"max_fee": "0x0000000000000000000000000000000000000000000000000000000000000000",
					"version": "0x0",
					"signature": [],
					"nonce": "",
					"type": "INVOKE",
					"contract_address": "0x031c887d82502ceb218c06ebb46198da3f7b92864a8223746bc836dda3e34b52",
					"entry_point_selector": "0x0218f305395474a84a39307fa5297be118fe17bf65e27ac5e2de6617baa44c64",
					"calldata": [
						"0x0735596016a37ee972c42adef6a3cf628c19bb3794369c65d2c82ba034aecf2c",
						"0x0000000000000000000000000000000000000000000000000000000000000001"
					]
				},
				{
					"txn_hash": "0x00b05ba5cd0b9e0464d2c1790ad93a159c6ef0594513758bca9111e74e4099d4",
					"max_fee": "0x0000000000000000000000000000000000000000000000000000000000000000",
					"version": "0x0",
					"signature": [],
					"nonce": "",
					"type": "INVOKE",
					"contract_address": "0x0735596016a37ee972c42adef6a3cf628c19bb3794369c65d2c82ba034aecf2c",
					"entry_point_selector": "0x0218f305395474a84a39307fa5297be118fe17bf65e27ac5e2de6617baa44c64",
					"calldata": [
						"0x031c887d82502ceb218c06ebb46198da3f7b92864a8223746bc836dda3e34b52",
						"0x0000000000000000000000000000000000000000000000000000000000000000"
					]
				},
				{
					"txn_hash": "0x04d16393d940fb4a97f20b9034e2a5e954201fee827b2b5c6daa38ec272e7c9c",
					"max_fee": "0x0000000000000000000000000000000000000000000000000000000000000000",
					"version": "0x0",
					"signature": [],
					"nonce": "",
					"type": "INVOKE",
					"contract_address": "0x06ee3440b08a9c805305449ec7f7003f27e9f7e287b83610952ec36bdc5a6bae",
					"entry_point_selector": "0x0317eb442b72a9fae758d4fb26830ed0d9f31c8e7da4dbff4e8c59ea6a158e7f",
					"calldata": [
						"0x01a7cf8b8027ec2d8fd04f1277f3f8ae6379ca957c5fec9ee7f59d56d86a26e4",
						"0x0000000000000000000000000000000000000000000000000000000000000002",
						"0x028dff6722aa73281b2cf84cac09950b71fa90512db294d2042119abdd9f4b87",
						"0x057a8f8a019ccab5bfc6ff86c96b1392257abb8d5d110c01d326b94247af161c"
					]
				},
				{
					"txn_hash": "0x009e80672edd4927a79f5384e656416b066f8ef58238227ac0fcea01952b70b5",
					"max_fee": "0x0000000000000000000000000000000000000000000000000000000000000000",
					"version": "0x0",
					"signature": [],
					"nonce": "",
					"type": "INVOKE",
					"contract_address": "0x031c887d82502ceb218c06ebb46198da3f7b92864a8223746bc836dda3e34b52",
					"entry_point_selector": "0x019a35a6e95cb7a3318dbb244f20975a1cd8587cc6b5259f15f61d7beb7ee43b",
					"calldata": [
						"0x06ee3440b08a9c805305449ec7f7003f27e9f7e287b83610952ec36bdc5a6bae",
						"0x05f750dc13ed239fa6fc43ff6e10ae9125a33bd05ec034fc3bb4dd168df3505f"
					]
				},
				{
					"txn_hash": "0x0387b5b63e40d4426754895fe52adf668cf8fde2a02aa9b6d761873f31af3462",
					"max_fee": "0x0000000000000000000000000000000000000000000000000000000000000000",
					"version": "0x0",
					"signature": [],
					"nonce": "",
					"type": "INVOKE",
					"contract_address": "0x06ee3440b08a9c805305449ec7f7003f27e9f7e287b83610952ec36bdc5a6bae",
					"entry_point_selector": "0x03d7905601c217734671143d457f0db37f7f8883112abd34b92c4abfeafde0c3",
					"calldata": [
						"0x0449908c349e90f81ab13042b1e49dc251eb6e3e51092d9a40f86859f7f415b0",
						"0x02670b3a8266d5046696a4b79f7433d117d3a19166f15bbd8585822c4e9b7491"
					]
				},
				{
					"txn_hash": "0x04f0cdff0d72fc758413a16db2bc7580dfec7889a8b921f0fe08641fa265e997",
					"max_fee": "0x0000000000000000000000000000000000000000000000000000000000000000",
					"version": "0x0",
					"signature": [],
					"nonce": "",
					"type": "INVOKE",
					"contract_address": "0x06ee3440b08a9c805305449ec7f7003f27e9f7e287b83610952ec36bdc5a6bae",
					"entry_point_selector": "0x03d7905601c217734671143d457f0db37f7f8883112abd34b92c4abfeafde0c3",
					"calldata": [
						"0x0449908c349e90f81ab13042b1e49dc251eb6e3e51092d9a40f86859f7f415b0",
						"0x06cb6104279e754967a721b52bcf5be525fdc11fa6db6ef5c3a4db832acf7804"
					]
				}
			]
		},
		"id": 1
	}`,
}
