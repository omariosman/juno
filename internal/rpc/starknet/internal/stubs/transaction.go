package stubs

import (
	"github.com/NethermindEth/juno/pkg/felt"
	"github.com/NethermindEth/juno/pkg/types"
)

type Transaction struct {
	Txn types.IsTransaction
}

var TransactionStubs = map[string]Transaction{
	"0x00e0a2e45a80bb827967e096bcf58874f6c01c191e0a0530624cba66a508ae75": {
		Txn: &types.TransactionDeploy{
			Hash:                new(felt.Felt).SetHex("0x00e0a2e45a80bb827967e096bcf58874f6c01c191e0a0530624cba66a508ae75"),
			ClassHash:           new(felt.Felt).SetHex("0x010455c752b86932ce552f2b0fe81a880746649b9aee7e0d842bf3f52378f9f8"),
			ContractAddress:     new(felt.Felt).SetHex("0x020cfa74ee3564b4cd5435cdace0f9c4d43b939620e4a0bb5076105df0a626c6"),
			ContractAddressSalt: new(felt.Felt).SetHex("0x0546c86dc6e40a5e5492b782d8964e9a4274ff6ecb16d31eb09cee45a3564015"),
			ConstructorCallData: []*felt.Felt{
				new(felt.Felt).SetHex("0x06cf6c2f36d36b08e591e4489e92ca882bb67b9c39a3afccf011972a8de467f0"),
				new(felt.Felt).SetHex("0x07ab344d88124307c07b56f6c59c12f4543e9c96398727854a322dea82c73240"),
			},
		},
	},
}
