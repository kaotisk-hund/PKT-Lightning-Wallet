// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package walletrpc

import (
	context "context"
	signrpc "github.com/pkt-cash/pktd/lnd/lnrpc/signrpc"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// WalletKitClient is the client API for WalletKit service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WalletKitClient interface {
	//
	//$pld.category: `Unspent`
	//$pld.short_description: `List utxos available for spending`
	//
	//ListUnspent returns a list of all utxos spendable by the wallet with a
	//number of confirmations between the specified minimum and maximum.
	ListUnspent(ctx context.Context, in *ListUnspentRequest, opts ...grpc.CallOption) (*ListUnspentResponse, error)
	//
	//LeaseOutput locks an output to the given ID, preventing it from being
	//available for any future coin selection attempts. The absolute time of the
	//lock's expiration is returned. The expiration of the lock can be extended by
	//successive invocations of this RPC. Outputs can be unlocked before their
	//expiration through `ReleaseOutput`.
	LeaseOutput(ctx context.Context, in *LeaseOutputRequest, opts ...grpc.CallOption) (*LeaseOutputResponse, error)
	//
	//ReleaseOutput unlocks an output, allowing it to be available for coin
	//selection if it remains unspent. The ID should match the one used to
	//originally lock the output.
	ReleaseOutput(ctx context.Context, in *ReleaseOutputRequest, opts ...grpc.CallOption) (*ReleaseOutputResponse, error)
	//
	//DeriveNextKey attempts to derive the *next* key within the key family
	//(account in BIP43) specified. This method should return the next external
	//child within this branch.
	DeriveNextKey(ctx context.Context, in *KeyReq, opts ...grpc.CallOption) (*signrpc.KeyDescriptor, error)
	//
	//DeriveKey attempts to derive an arbitrary key specified by the passed
	//KeyLocator.
	DeriveKey(ctx context.Context, in *signrpc.KeyLocator, opts ...grpc.CallOption) (*signrpc.KeyDescriptor, error)
	//
	//NextAddr returns the next unused address within the wallet.
	NextAddr(ctx context.Context, in *AddrRequest, opts ...grpc.CallOption) (*AddrResponse, error)
	//
	//PublishTransaction attempts to publish the passed transaction to the
	//network. Once this returns without an error, the wallet will continually
	//attempt to re-broadcast the transaction on start up, until it enters the
	//chain.
	PublishTransaction(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*PublishResponse, error)
	//
	//SendOutputs is similar to the existing sendmany call in Bitcoind, and
	//allows the caller to create a transaction that sends to several outputs at
	//once. This is ideal when wanting to batch create a set of transactions.
	SendOutputs(ctx context.Context, in *SendOutputsRequest, opts ...grpc.CallOption) (*SendOutputsResponse, error)
	//
	//EstimateFee attempts to query the internal fee estimator of the wallet to
	//determine the fee (in sat/kw) to attach to a transaction in order to
	//achieve the confirmation target.
	EstimateFee(ctx context.Context, in *EstimateFeeRequest, opts ...grpc.CallOption) (*EstimateFeeResponse, error)
	//
	//PendingSweeps returns lists of on-chain outputs that lnd is currently
	//attempting to sweep within its central batching engine. Outputs with similar
	//fee rates are batched together in order to sweep them within a single
	//transaction.
	//
	//NOTE: Some of the fields within PendingSweepsRequest are not guaranteed to
	//remain supported. This is an advanced API that depends on the internals of
	//the UtxoSweeper, so things may change.
	PendingSweeps(ctx context.Context, in *PendingSweepsRequest, opts ...grpc.CallOption) (*PendingSweepsResponse, error)
	//
	//BumpFee bumps the fee of an arbitrary input within a transaction. This RPC
	//takes a different approach than bitcoind's bumpfee command. lnd has a
	//central batching engine in which inputs with similar fee rates are batched
	//together to save on transaction fees. Due to this, we cannot rely on
	//bumping the fee on a specific transaction, since transactions can change at
	//any point with the addition of new inputs. The list of inputs that
	//currently exist within lnd's central batching engine can be retrieved
	//through the PendingSweeps RPC.
	//
	//When bumping the fee of an input that currently exists within lnd's central
	//batching engine, a higher fee transaction will be created that replaces the
	//lower fee transaction through the Replace-By-Fee (RBF) policy. If it
	//
	//This RPC also serves useful when wanting to perform a Child-Pays-For-Parent
	//(CPFP), where the child transaction pays for its parent's fee. This can be
	//done by specifying an outpoint within the low fee transaction that is under
	//the control of the wallet.
	//
	//The fee preference can be expressed either as a specific fee rate or a delta
	//of blocks in which the output should be swept on-chain within. If a fee
	//preference is not explicitly specified, then an error is returned.
	//
	//Note that this RPC currently doesn't perform any validation checks on the
	//fee preference being provided. For now, the responsibility of ensuring that
	//the new fee preference is sufficient is delegated to the user.
	BumpFee(ctx context.Context, in *BumpFeeRequest, opts ...grpc.CallOption) (*BumpFeeResponse, error)
	//
	//ListSweeps returns a list of the sweep transactions our node has produced.
	//Note that these sweeps may not be confirmed yet, as we record sweeps on
	//broadcast, not confirmation.
	ListSweeps(ctx context.Context, in *ListSweepsRequest, opts ...grpc.CallOption) (*ListSweepsResponse, error)
	//
	//LabelTransaction adds a label to a transaction. If the transaction already
	//has a label the call will fail unless the overwrite bool is set. This will
	//overwrite the exiting transaction label. Labels must not be empty, and
	//cannot exceed 500 characters.
	LabelTransaction(ctx context.Context, in *LabelTransactionRequest, opts ...grpc.CallOption) (*LabelTransactionResponse, error)
	//
	//FundPsbt creates a fully populated PSBT that contains enough inputs to fund
	//the outputs specified in the template. There are two ways of specifying a
	//template: Either by passing in a PSBT with at least one output declared or
	//by passing in a raw TxTemplate message.
	//
	//If there are no inputs specified in the template, coin selection is
	//performed automatically. If the template does contain any inputs, it is
	//assumed that full coin selection happened externally and no additional
	//inputs are added. If the specified inputs aren't enough to fund the outputs
	//with the given fee rate, an error is returned.
	//
	//After either selecting or verifying the inputs, all input UTXOs are locked
	//with an internal app ID.
	//
	//NOTE: If this method returns without an error, it is the caller's
	//responsibility to either spend the locked UTXOs (by finalizing and then
	//publishing the transaction) or to unlock/release the locked UTXOs in case of
	//an error on the caller's side.
	FundPsbt(ctx context.Context, in *FundPsbtRequest, opts ...grpc.CallOption) (*FundPsbtResponse, error)
	//
	//FinalizePsbt expects a partial transaction with all inputs and outputs fully
	//declared and tries to sign all inputs that belong to the wallet. Lnd must be
	//the last signer of the transaction. That means, if there are any unsigned
	//non-witness inputs or inputs without UTXO information attached or inputs
	//without witness data that do not belong to lnd's wallet, this method will
	//fail. If no error is returned, the PSBT is ready to be extracted and the
	//final TX within to be broadcast.
	//
	//NOTE: This method does NOT publish the transaction once finalized. It is the
	//caller's responsibility to either publish the transaction on success or
	//unlock/release any locked UTXOs in case of an error in this method.
	FinalizePsbt(ctx context.Context, in *FinalizePsbtRequest, opts ...grpc.CallOption) (*FinalizePsbtResponse, error)
}

type walletKitClient struct {
	cc grpc.ClientConnInterface
}

func NewWalletKitClient(cc grpc.ClientConnInterface) WalletKitClient {
	return &walletKitClient{cc}
}

func (c *walletKitClient) ListUnspent(ctx context.Context, in *ListUnspentRequest, opts ...grpc.CallOption) (*ListUnspentResponse, error) {
	out := new(ListUnspentResponse)
	err := c.cc.Invoke(ctx, "/walletrpc.WalletKit/ListUnspent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletKitClient) LeaseOutput(ctx context.Context, in *LeaseOutputRequest, opts ...grpc.CallOption) (*LeaseOutputResponse, error) {
	out := new(LeaseOutputResponse)
	err := c.cc.Invoke(ctx, "/walletrpc.WalletKit/LeaseOutput", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletKitClient) ReleaseOutput(ctx context.Context, in *ReleaseOutputRequest, opts ...grpc.CallOption) (*ReleaseOutputResponse, error) {
	out := new(ReleaseOutputResponse)
	err := c.cc.Invoke(ctx, "/walletrpc.WalletKit/ReleaseOutput", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletKitClient) DeriveNextKey(ctx context.Context, in *KeyReq, opts ...grpc.CallOption) (*signrpc.KeyDescriptor, error) {
	out := new(signrpc.KeyDescriptor)
	err := c.cc.Invoke(ctx, "/walletrpc.WalletKit/DeriveNextKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletKitClient) DeriveKey(ctx context.Context, in *signrpc.KeyLocator, opts ...grpc.CallOption) (*signrpc.KeyDescriptor, error) {
	out := new(signrpc.KeyDescriptor)
	err := c.cc.Invoke(ctx, "/walletrpc.WalletKit/DeriveKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletKitClient) NextAddr(ctx context.Context, in *AddrRequest, opts ...grpc.CallOption) (*AddrResponse, error) {
	out := new(AddrResponse)
	err := c.cc.Invoke(ctx, "/walletrpc.WalletKit/NextAddr", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletKitClient) PublishTransaction(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*PublishResponse, error) {
	out := new(PublishResponse)
	err := c.cc.Invoke(ctx, "/walletrpc.WalletKit/PublishTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletKitClient) SendOutputs(ctx context.Context, in *SendOutputsRequest, opts ...grpc.CallOption) (*SendOutputsResponse, error) {
	out := new(SendOutputsResponse)
	err := c.cc.Invoke(ctx, "/walletrpc.WalletKit/SendOutputs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletKitClient) EstimateFee(ctx context.Context, in *EstimateFeeRequest, opts ...grpc.CallOption) (*EstimateFeeResponse, error) {
	out := new(EstimateFeeResponse)
	err := c.cc.Invoke(ctx, "/walletrpc.WalletKit/EstimateFee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletKitClient) PendingSweeps(ctx context.Context, in *PendingSweepsRequest, opts ...grpc.CallOption) (*PendingSweepsResponse, error) {
	out := new(PendingSweepsResponse)
	err := c.cc.Invoke(ctx, "/walletrpc.WalletKit/PendingSweeps", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletKitClient) BumpFee(ctx context.Context, in *BumpFeeRequest, opts ...grpc.CallOption) (*BumpFeeResponse, error) {
	out := new(BumpFeeResponse)
	err := c.cc.Invoke(ctx, "/walletrpc.WalletKit/BumpFee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletKitClient) ListSweeps(ctx context.Context, in *ListSweepsRequest, opts ...grpc.CallOption) (*ListSweepsResponse, error) {
	out := new(ListSweepsResponse)
	err := c.cc.Invoke(ctx, "/walletrpc.WalletKit/ListSweeps", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletKitClient) LabelTransaction(ctx context.Context, in *LabelTransactionRequest, opts ...grpc.CallOption) (*LabelTransactionResponse, error) {
	out := new(LabelTransactionResponse)
	err := c.cc.Invoke(ctx, "/walletrpc.WalletKit/LabelTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletKitClient) FundPsbt(ctx context.Context, in *FundPsbtRequest, opts ...grpc.CallOption) (*FundPsbtResponse, error) {
	out := new(FundPsbtResponse)
	err := c.cc.Invoke(ctx, "/walletrpc.WalletKit/FundPsbt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletKitClient) FinalizePsbt(ctx context.Context, in *FinalizePsbtRequest, opts ...grpc.CallOption) (*FinalizePsbtResponse, error) {
	out := new(FinalizePsbtResponse)
	err := c.cc.Invoke(ctx, "/walletrpc.WalletKit/FinalizePsbt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WalletKitServer is the server API for WalletKit service.
// All implementations should embed UnimplementedWalletKitServer
// for forward compatibility
type WalletKitServer interface {
	//
	//$pld.category: `Unspent`
	//$pld.short_description: `List utxos available for spending`
	//
	//ListUnspent returns a list of all utxos spendable by the wallet with a
	//number of confirmations between the specified minimum and maximum.
	ListUnspent(context.Context, *ListUnspentRequest) (*ListUnspentResponse, error)
	//
	//LeaseOutput locks an output to the given ID, preventing it from being
	//available for any future coin selection attempts. The absolute time of the
	//lock's expiration is returned. The expiration of the lock can be extended by
	//successive invocations of this RPC. Outputs can be unlocked before their
	//expiration through `ReleaseOutput`.
	LeaseOutput(context.Context, *LeaseOutputRequest) (*LeaseOutputResponse, error)
	//
	//ReleaseOutput unlocks an output, allowing it to be available for coin
	//selection if it remains unspent. The ID should match the one used to
	//originally lock the output.
	ReleaseOutput(context.Context, *ReleaseOutputRequest) (*ReleaseOutputResponse, error)
	//
	//DeriveNextKey attempts to derive the *next* key within the key family
	//(account in BIP43) specified. This method should return the next external
	//child within this branch.
	DeriveNextKey(context.Context, *KeyReq) (*signrpc.KeyDescriptor, error)
	//
	//DeriveKey attempts to derive an arbitrary key specified by the passed
	//KeyLocator.
	DeriveKey(context.Context, *signrpc.KeyLocator) (*signrpc.KeyDescriptor, error)
	//
	//NextAddr returns the next unused address within the wallet.
	NextAddr(context.Context, *AddrRequest) (*AddrResponse, error)
	//
	//PublishTransaction attempts to publish the passed transaction to the
	//network. Once this returns without an error, the wallet will continually
	//attempt to re-broadcast the transaction on start up, until it enters the
	//chain.
	PublishTransaction(context.Context, *Transaction) (*PublishResponse, error)
	//
	//SendOutputs is similar to the existing sendmany call in Bitcoind, and
	//allows the caller to create a transaction that sends to several outputs at
	//once. This is ideal when wanting to batch create a set of transactions.
	SendOutputs(context.Context, *SendOutputsRequest) (*SendOutputsResponse, error)
	//
	//EstimateFee attempts to query the internal fee estimator of the wallet to
	//determine the fee (in sat/kw) to attach to a transaction in order to
	//achieve the confirmation target.
	EstimateFee(context.Context, *EstimateFeeRequest) (*EstimateFeeResponse, error)
	//
	//PendingSweeps returns lists of on-chain outputs that lnd is currently
	//attempting to sweep within its central batching engine. Outputs with similar
	//fee rates are batched together in order to sweep them within a single
	//transaction.
	//
	//NOTE: Some of the fields within PendingSweepsRequest are not guaranteed to
	//remain supported. This is an advanced API that depends on the internals of
	//the UtxoSweeper, so things may change.
	PendingSweeps(context.Context, *PendingSweepsRequest) (*PendingSweepsResponse, error)
	//
	//BumpFee bumps the fee of an arbitrary input within a transaction. This RPC
	//takes a different approach than bitcoind's bumpfee command. lnd has a
	//central batching engine in which inputs with similar fee rates are batched
	//together to save on transaction fees. Due to this, we cannot rely on
	//bumping the fee on a specific transaction, since transactions can change at
	//any point with the addition of new inputs. The list of inputs that
	//currently exist within lnd's central batching engine can be retrieved
	//through the PendingSweeps RPC.
	//
	//When bumping the fee of an input that currently exists within lnd's central
	//batching engine, a higher fee transaction will be created that replaces the
	//lower fee transaction through the Replace-By-Fee (RBF) policy. If it
	//
	//This RPC also serves useful when wanting to perform a Child-Pays-For-Parent
	//(CPFP), where the child transaction pays for its parent's fee. This can be
	//done by specifying an outpoint within the low fee transaction that is under
	//the control of the wallet.
	//
	//The fee preference can be expressed either as a specific fee rate or a delta
	//of blocks in which the output should be swept on-chain within. If a fee
	//preference is not explicitly specified, then an error is returned.
	//
	//Note that this RPC currently doesn't perform any validation checks on the
	//fee preference being provided. For now, the responsibility of ensuring that
	//the new fee preference is sufficient is delegated to the user.
	BumpFee(context.Context, *BumpFeeRequest) (*BumpFeeResponse, error)
	//
	//ListSweeps returns a list of the sweep transactions our node has produced.
	//Note that these sweeps may not be confirmed yet, as we record sweeps on
	//broadcast, not confirmation.
	ListSweeps(context.Context, *ListSweepsRequest) (*ListSweepsResponse, error)
	//
	//LabelTransaction adds a label to a transaction. If the transaction already
	//has a label the call will fail unless the overwrite bool is set. This will
	//overwrite the exiting transaction label. Labels must not be empty, and
	//cannot exceed 500 characters.
	LabelTransaction(context.Context, *LabelTransactionRequest) (*LabelTransactionResponse, error)
	//
	//FundPsbt creates a fully populated PSBT that contains enough inputs to fund
	//the outputs specified in the template. There are two ways of specifying a
	//template: Either by passing in a PSBT with at least one output declared or
	//by passing in a raw TxTemplate message.
	//
	//If there are no inputs specified in the template, coin selection is
	//performed automatically. If the template does contain any inputs, it is
	//assumed that full coin selection happened externally and no additional
	//inputs are added. If the specified inputs aren't enough to fund the outputs
	//with the given fee rate, an error is returned.
	//
	//After either selecting or verifying the inputs, all input UTXOs are locked
	//with an internal app ID.
	//
	//NOTE: If this method returns without an error, it is the caller's
	//responsibility to either spend the locked UTXOs (by finalizing and then
	//publishing the transaction) or to unlock/release the locked UTXOs in case of
	//an error on the caller's side.
	FundPsbt(context.Context, *FundPsbtRequest) (*FundPsbtResponse, error)
	//
	//FinalizePsbt expects a partial transaction with all inputs and outputs fully
	//declared and tries to sign all inputs that belong to the wallet. Lnd must be
	//the last signer of the transaction. That means, if there are any unsigned
	//non-witness inputs or inputs without UTXO information attached or inputs
	//without witness data that do not belong to lnd's wallet, this method will
	//fail. If no error is returned, the PSBT is ready to be extracted and the
	//final TX within to be broadcast.
	//
	//NOTE: This method does NOT publish the transaction once finalized. It is the
	//caller's responsibility to either publish the transaction on success or
	//unlock/release any locked UTXOs in case of an error in this method.
	FinalizePsbt(context.Context, *FinalizePsbtRequest) (*FinalizePsbtResponse, error)
}

// UnimplementedWalletKitServer should be embedded to have forward compatible implementations.
type UnimplementedWalletKitServer struct {
}

func (UnimplementedWalletKitServer) ListUnspent(context.Context, *ListUnspentRequest) (*ListUnspentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUnspent not implemented")
}
func (UnimplementedWalletKitServer) LeaseOutput(context.Context, *LeaseOutputRequest) (*LeaseOutputResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeaseOutput not implemented")
}
func (UnimplementedWalletKitServer) ReleaseOutput(context.Context, *ReleaseOutputRequest) (*ReleaseOutputResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReleaseOutput not implemented")
}
func (UnimplementedWalletKitServer) DeriveNextKey(context.Context, *KeyReq) (*signrpc.KeyDescriptor, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeriveNextKey not implemented")
}
func (UnimplementedWalletKitServer) DeriveKey(context.Context, *signrpc.KeyLocator) (*signrpc.KeyDescriptor, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeriveKey not implemented")
}
func (UnimplementedWalletKitServer) NextAddr(context.Context, *AddrRequest) (*AddrResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NextAddr not implemented")
}
func (UnimplementedWalletKitServer) PublishTransaction(context.Context, *Transaction) (*PublishResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishTransaction not implemented")
}
func (UnimplementedWalletKitServer) SendOutputs(context.Context, *SendOutputsRequest) (*SendOutputsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendOutputs not implemented")
}
func (UnimplementedWalletKitServer) EstimateFee(context.Context, *EstimateFeeRequest) (*EstimateFeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EstimateFee not implemented")
}
func (UnimplementedWalletKitServer) PendingSweeps(context.Context, *PendingSweepsRequest) (*PendingSweepsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PendingSweeps not implemented")
}
func (UnimplementedWalletKitServer) BumpFee(context.Context, *BumpFeeRequest) (*BumpFeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BumpFee not implemented")
}
func (UnimplementedWalletKitServer) ListSweeps(context.Context, *ListSweepsRequest) (*ListSweepsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSweeps not implemented")
}
func (UnimplementedWalletKitServer) LabelTransaction(context.Context, *LabelTransactionRequest) (*LabelTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LabelTransaction not implemented")
}
func (UnimplementedWalletKitServer) FundPsbt(context.Context, *FundPsbtRequest) (*FundPsbtResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FundPsbt not implemented")
}
func (UnimplementedWalletKitServer) FinalizePsbt(context.Context, *FinalizePsbtRequest) (*FinalizePsbtResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FinalizePsbt not implemented")
}

// UnsafeWalletKitServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WalletKitServer will
// result in compilation errors.
type UnsafeWalletKitServer interface {
	mustEmbedUnimplementedWalletKitServer()
}

func RegisterWalletKitServer(s grpc.ServiceRegistrar, srv WalletKitServer) {
	s.RegisterService(&WalletKit_ServiceDesc, srv)
}

func _WalletKit_ListUnspent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUnspentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletKitServer).ListUnspent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/walletrpc.WalletKit/ListUnspent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletKitServer).ListUnspent(ctx, req.(*ListUnspentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletKit_LeaseOutput_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeaseOutputRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletKitServer).LeaseOutput(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/walletrpc.WalletKit/LeaseOutput",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletKitServer).LeaseOutput(ctx, req.(*LeaseOutputRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletKit_ReleaseOutput_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReleaseOutputRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletKitServer).ReleaseOutput(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/walletrpc.WalletKit/ReleaseOutput",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletKitServer).ReleaseOutput(ctx, req.(*ReleaseOutputRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletKit_DeriveNextKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletKitServer).DeriveNextKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/walletrpc.WalletKit/DeriveNextKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletKitServer).DeriveNextKey(ctx, req.(*KeyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletKit_DeriveKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(signrpc.KeyLocator)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletKitServer).DeriveKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/walletrpc.WalletKit/DeriveKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletKitServer).DeriveKey(ctx, req.(*signrpc.KeyLocator))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletKit_NextAddr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddrRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletKitServer).NextAddr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/walletrpc.WalletKit/NextAddr",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletKitServer).NextAddr(ctx, req.(*AddrRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletKit_PublishTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletKitServer).PublishTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/walletrpc.WalletKit/PublishTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletKitServer).PublishTransaction(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletKit_SendOutputs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendOutputsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletKitServer).SendOutputs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/walletrpc.WalletKit/SendOutputs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletKitServer).SendOutputs(ctx, req.(*SendOutputsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletKit_EstimateFee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EstimateFeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletKitServer).EstimateFee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/walletrpc.WalletKit/EstimateFee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletKitServer).EstimateFee(ctx, req.(*EstimateFeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletKit_PendingSweeps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PendingSweepsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletKitServer).PendingSweeps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/walletrpc.WalletKit/PendingSweeps",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletKitServer).PendingSweeps(ctx, req.(*PendingSweepsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletKit_BumpFee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BumpFeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletKitServer).BumpFee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/walletrpc.WalletKit/BumpFee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletKitServer).BumpFee(ctx, req.(*BumpFeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletKit_ListSweeps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSweepsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletKitServer).ListSweeps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/walletrpc.WalletKit/ListSweeps",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletKitServer).ListSweeps(ctx, req.(*ListSweepsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletKit_LabelTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LabelTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletKitServer).LabelTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/walletrpc.WalletKit/LabelTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletKitServer).LabelTransaction(ctx, req.(*LabelTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletKit_FundPsbt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FundPsbtRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletKitServer).FundPsbt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/walletrpc.WalletKit/FundPsbt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletKitServer).FundPsbt(ctx, req.(*FundPsbtRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletKit_FinalizePsbt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FinalizePsbtRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletKitServer).FinalizePsbt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/walletrpc.WalletKit/FinalizePsbt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletKitServer).FinalizePsbt(ctx, req.(*FinalizePsbtRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WalletKit_ServiceDesc is the grpc.ServiceDesc for WalletKit service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WalletKit_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "walletrpc.WalletKit",
	HandlerType: (*WalletKitServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListUnspent",
			Handler:    _WalletKit_ListUnspent_Handler,
		},
		{
			MethodName: "LeaseOutput",
			Handler:    _WalletKit_LeaseOutput_Handler,
		},
		{
			MethodName: "ReleaseOutput",
			Handler:    _WalletKit_ReleaseOutput_Handler,
		},
		{
			MethodName: "DeriveNextKey",
			Handler:    _WalletKit_DeriveNextKey_Handler,
		},
		{
			MethodName: "DeriveKey",
			Handler:    _WalletKit_DeriveKey_Handler,
		},
		{
			MethodName: "NextAddr",
			Handler:    _WalletKit_NextAddr_Handler,
		},
		{
			MethodName: "PublishTransaction",
			Handler:    _WalletKit_PublishTransaction_Handler,
		},
		{
			MethodName: "SendOutputs",
			Handler:    _WalletKit_SendOutputs_Handler,
		},
		{
			MethodName: "EstimateFee",
			Handler:    _WalletKit_EstimateFee_Handler,
		},
		{
			MethodName: "PendingSweeps",
			Handler:    _WalletKit_PendingSweeps_Handler,
		},
		{
			MethodName: "BumpFee",
			Handler:    _WalletKit_BumpFee_Handler,
		},
		{
			MethodName: "ListSweeps",
			Handler:    _WalletKit_ListSweeps_Handler,
		},
		{
			MethodName: "LabelTransaction",
			Handler:    _WalletKit_LabelTransaction_Handler,
		},
		{
			MethodName: "FundPsbt",
			Handler:    _WalletKit_FundPsbt_Handler,
		},
		{
			MethodName: "FinalizePsbt",
			Handler:    _WalletKit_FinalizePsbt_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "walletrpc/walletkit.proto",
}