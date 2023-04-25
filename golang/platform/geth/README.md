# GETH

## [ Install ]

```sh
sudo add-apt-repository -y ppa:ethereum/ethereum
sudo apt-get update
sudo apt-get install ethereum
```

## [ Concept ]

### - Transaction Types

 - `LegacyTxType`: EIP-1559 이전의 일반적인 트랜잭션 타입입니다. 가스 가격과 가스 한도를 명시적으로 지정합니다.
 - `AccessListTxType`: EIP-2930에서 추가된 타입으로, 트랜잭션 발신자가 액세스하려는 계정 목록을 지정할 수 있습니다. 이를 통해 실행되는 스마트 계약의 가스 비용을 줄일 수 있습니다.
 - `DynamicFeeTxType`: EIP-1559에서 추가된 타입으로, 블록 내부 가스 가격 경매를 통해 동적으로 가스 가격을 결정합니다. 이를 통해 블록의 가스 한도를 유지하면서 트랜잭션 처리 속도와 수수료 측면에서 개선이 이루어집니다.

즉, LegacyTxType는 기존의 방식대로 가스 가격과 가스 한도를 지정하는 방식이며, 
AccessListTxType은 일부 계정에 대해서만 가스를 소비하도록 할 수 있는 방식이고, 
DynamicFeeTxType은 동적으로 가스 가격이 결정되는 방식입니다.

현재로서는 EIP-1559에서 추가된 DynamicFeeTxType이 가장 최신의 트랜잭션 타입입니다.
DynamicFeeTxType은 블록 내부 가스 가격 경매를 통해 동적으로 가스 가격을 결정하므로, 트랜잭션 처리 속도와 수수료 측면에서 이전 방식에 비해 향상된 성능을 제공합니다. 
하지만 이전 방식의 LegacyTxType도 아직까지 많이 사용되고 있으며, 일부 경우에는 AccessListTxType도 사용될 수 있습니다.

### - Signer Types

 - LatestSigner: 현재 사용 중인 네트워크 버전에 해당하는 최신 서명 방식을 사용합니다.
 - EIP2930Signer:  EIP-155에서 도입된 CHAIN_ID 필드를 사용하여 서명합니다.
 - LondonSigner: Frontier 버전의 네트워크에서 사용되는 서명 방식을 사용합니다.
 - HomesteadSigner:  Homestead 버전의 네트워크에서 사용되는 서명 방식을 사용합니다.
 - FrontierSigner: London 업그레이드 이후 네트워크에서 사용되는 서명 방식을 사용합니다.
 - EIP155Signer: EIP-2930에서 도입된 access list를 사용하여 서명합니다.

### - RLP(Recursive Length Prefix)

Ethereum에서 데이터를 직렬화하는 방식 중 하나입니다. RLP는 바이트 배열로 표현된 계층 구조 데이터를 직렬화할 수 있습니다.