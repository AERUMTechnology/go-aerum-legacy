pragma solidity ^0.4.25;

contract AtmosGovernance {

    address reptileOverlord;
    ComposerOperation[] operations;
    BootstrapDelegates[] bootstrappers; 

    struct ComposerOperation {
        uint256 block;
        address composer;
        OperationType opType;
    }

    struct BootstrapDelegates {
        uint16 index;
        address delegate;
        OperationType opType;
    }

    enum OperationType { Add, Remove }

    constructor() public {
        reptileOverlord = msg.sender;
    }

    function addComposer(uint256 aerumBlock, address composer) external {
        ComposerOperation memory operation = ComposerOperation({
            block: aerumBlock,
            composer: composer,
            opType: OperationType.Add
        });
        operations.push(operation);
    }

    function removeComposer(uint256 aerumBlock, address composer) external {
        ComposerOperation memory operation = ComposerOperation({
            block: aerumBlock,
            composer: composer,
            opType: OperationType.Remove
        });
        operations.push(operation);
    }

    function getComposers(uint256 aerumBlock, uint256 timestamp) external view returns (address[]) {
        address[] memory candidates = getValidComposers(aerumBlock);
        if (candidates.length <= 10) {
            return candidates;
        }
        address[] memory composers = new address[](10);
        uint256 first = (aerumBlock / 1000) % candidates.length;
        uint256 last = first + 10;
        uint256 fromStart = 0;
        if (last > candidates.length) {
            fromStart = last - candidates.length;
            last = candidates.length;
        }
        for (uint256 i = first; i < last; i++) {
            composers[i - first] = candidates[i];
        }
        for (uint256 j = 0; j < fromStart; j++) {
            composers[10 - j - 1] = candidates[j];
        }
        return composers;
    }

    function getValidComposers(uint256 aerumBlock) public view returns (address[]) {
        uint256 operationIndex = 0;
        uint256 composerIndex = 0;
        uint256 composerCount = 0;

        // There should not be more composers than operations
        address[] memory composers = new address[](operations.length);
        bool[] memory isValid = new bool[](operations.length);

        while (operationIndex < operations.length && operations[operationIndex].block <= aerumBlock) {
            ComposerOperation memory operation = operations[operationIndex];

            for(composerIndex = 0; composerIndex < composerCount; composerIndex++) {
                // Find if we already have composer in the list
                if(composers[composerIndex] == operation.composer) {
                    break;
                }
            }

            if(composerIndex == composerCount) {
                // We don't have composer yet, add one
                composers[composerIndex] = operation.composer;
                composerCount++;
            }

            // Update valid or not
            isValid[composerIndex] = bool(operation.opType == OperationType.Add);
            operationIndex++;
        }

        uint256 validComposersCount = 0;
        address[] memory validComposers = new address[](composerCount);
        for (composerIndex = 0; composerIndex < composerCount; composerIndex++) {
            if(isValid[composerIndex]) {
                validComposers[validComposersCount] = composers[composerIndex];
                validComposersCount++;
            }
        }

        address[] memory result = new address[](validComposersCount);
        for (uint256 index = 0; index < validComposersCount; index++) {
            result[index] = validComposers[index];
        }
        return result;
    }

    function addBoostrapDelegate(uint16 delegateIndex, address delegate) external {
        BootstrapDelegates memory strap = BootstrapDelegates({
            index: delegateIndex,
            delegate: delegate,
            opType: OperationType.Add
        });
        bootstrappers.push(strap);
    }

    function removeComposer(uint16 delegateIndex, address delegate) external {
        BootstrapDelegates memory strap = BootstrapDelegates({
            index: delegateIndex,
            delegate: delegate,
            opType: OperationType.Remove
        });
        bootstrappers.push(strap);
    }
    
    function showBootstrapDelegates() public view returns (address[]) {       
        uint256 i = 0;
        address[] memory collectedDelegates = new address[](bootstrappers.length);

        for(i = 0; i < bootstrappers.length; i++) {
            collectedDelegates[i] = bootstrappers[i].delegate;
        }
        return collectedDelegates;
    }

    function cleanBootstrapDelegates() external {
        delete bootstrappers;
    }

    function clean() external {
        delete operations;
    }
}