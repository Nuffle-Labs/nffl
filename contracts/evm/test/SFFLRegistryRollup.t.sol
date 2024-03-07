// SPDX-License-Identifier: MIT
pragma solidity ^0.8.12;

import {BN254} from "eigenlayer-middleware/src/libraries/BN254.sol";

import {SFFLRegistryRollup} from "../src/rollup/SFFLRegistryRollup.sol";
import {StateRootUpdate} from "../src/base/message/StateRootUpdate.sol";
import {Operators} from "../src/rollup/utils/Operators.sol";
import {OperatorSetUpdate} from "../src/rollup/message/OperatorSetUpdate.sol";

import {TestUtils} from "./utils/TestUtils.sol";

contract SFFLRegistryRollupTest is TestUtils {
    using BN254 for BN254.G1Point;
    using OperatorSetUpdate for OperatorSetUpdate.Message;
    using StateRootUpdate for StateRootUpdate.Message;

    SFFLRegistryRollup public registry;

    Operators.Operator[] public initialOperators;
    Operators.Operator[] public extraOperators;

    uint128 public constant DEFAULT_WEIGHT = 100;
    uint128 public QUORUM_THRESHOLD = 2 * uint128(100) / 3;

    event StateRootUpdated(uint32 indexed rollupId, uint64 indexed blockHeight, bytes32 stateRoot);
    event OperatorUpdated(bytes32 indexed pubkeyHash, uint128 weight);
    event QuorumThresholdUpdated(uint128 indexed newQuorumThreshold);

    function setUp() public {
        // BLSUtilsFFI.keygen(4, 100)
        initialOperators.push(
            Operators.Operator(
                BN254.G1Point(
                    9616480996400718794846151252530035789548914510416131398073228825370885543387,
                    1193202739907461244540326755381374372389880624644682929124544554638046348630
                ),
                DEFAULT_WEIGHT
            )
        );
        initialOperators.push(
            Operators.Operator(
                BN254.G1Point(
                    18127872302725521126948039783576335344235515541656874329008995106922633337074,
                    8399728219671791397177034566462901981672518419761436513477052772303615710768
                ),
                DEFAULT_WEIGHT
            )
        );
        initialOperators.push(
            Operators.Operator(
                BN254.G1Point(
                    21052107500402163350976152667562860793896335512669052649238089783827995103691,
                    19540845318985121430596366476154886876615197304682640344094567483512876396969
                ),
                DEFAULT_WEIGHT
            )
        );
        initialOperators.push(
            Operators.Operator(
                BN254.G1Point(
                    2219290000546820918614914859472334932465240992233723879760496915453797571330,
                    12792826480108937261317636923053265154992326678284813916740018134087766211155
                ),
                DEFAULT_WEIGHT
            )
        );

        extraOperators.push(
            Operators.Operator(
                BN254.G1Point(
                    1768235322131906328721284719328963673934008473718528701566488601237085071962,
                    4002933128315738291771882099514828917162455208473397665597366416040939248693
                ),
                DEFAULT_WEIGHT
            )
        );

        vm.prank(addr("owner"));
        registry = new SFFLRegistryRollup(initialOperators, QUORUM_THRESHOLD, 0);
    }

    function test_setUp() public {
        assertEq(
            registry.getApk().hashG1Point(),
            BN254.G1Point(
                7422372290785096726190210501277844820262413627971980991387239130751122169828,
                12427069946071247197599765748799305802712487068319377150664147036368161174602
            ).hashG1Point()
        );
        assertEq(registry.getTotalWeight(), 400);
        assertEq(registry.getQuorumThreshold(), QUORUM_THRESHOLD);
    }

    function test_updateOperatorSet() public {
        Operators.Operator[] memory operators = new Operators.Operator[](3);

        operators[0] = Operators.Operator(initialOperators[3].pubkey, 0);
        operators[1] = Operators.Operator(initialOperators[2].pubkey, 3 * DEFAULT_WEIGHT);
        operators[2] = extraOperators[0];

        OperatorSetUpdate.Message memory message =
            OperatorSetUpdate.Message(registry.nextOperatorUpdateId(), 0, operators);

        BN254.G1Point[] memory nonSignerPubkeys = new BN254.G1Point[](1);
        nonSignerPubkeys[0] = initialOperators[3].pubkey;

        Operators.SignatureInfo memory signatureInfo = Operators.SignatureInfo({
            nonSignerPubkeys: nonSignerPubkeys,
            apkG2: BN254.G2Point(
                [
                    21774854595736935906777183372431491423672246101465086449723107940773462536091,
                    11859388993407979358677113204795514610964422675159446451278647734574620707784
                ],
                [
                    3453374196609277266042659107600871924832557088868662992636101033001416801985,
                    2630500117064331827715800222355515273572786883080373379723474133051328147838
                ]
                ),
            sigma: BN254.hashToG1(message.hash()).scalar_mul(
                6305737925830641523797682626723526790077499630761662964405387941160208990354
                )
        });

        vm.expectEmit(true, false, false, true);
        emit OperatorUpdated(operators[0].pubkey.hashG1Point(), 0);
        vm.expectEmit(true, false, false, true);
        emit OperatorUpdated(operators[1].pubkey.hashG1Point(), 3 * DEFAULT_WEIGHT);
        vm.expectEmit(true, false, false, true);
        emit OperatorUpdated(operators[2].pubkey.hashG1Point(), DEFAULT_WEIGHT);

        registry.updateOperatorSet(message, signatureInfo);

        assertEq(
            registry.getApk().hashG1Point(),
            BN254.G1Point(
                20722407922923263883576268605784225769666234232258906224209479286774407267165,
                2757196487678270995513080287135364337546940931709504737604569117674180167662
            ).hashG1Point()
        );
        assertEq(registry.getTotalWeight(), 400 - DEFAULT_WEIGHT - DEFAULT_WEIGHT + 3 * DEFAULT_WEIGHT + DEFAULT_WEIGHT);
        assertEq(registry.getOperatorWeight(operators[0].pubkey.hashG1Point()), 0);
        assertEq(registry.getOperatorWeight(operators[1].pubkey.hashG1Point()), 3 * DEFAULT_WEIGHT);
        assertEq(registry.getOperatorWeight(operators[2].pubkey.hashG1Point()), DEFAULT_WEIGHT);
        assertEq(registry.nextOperatorUpdateId(), message.id + 1);
    }

    function test_updateOperatorSet_RevertWhen_QuorumNotMet() public {
        Operators.Operator[] memory operators = new Operators.Operator[](3);

        operators[0] = Operators.Operator(initialOperators[3].pubkey, 0);
        operators[1] = Operators.Operator(initialOperators[2].pubkey, 3 * DEFAULT_WEIGHT);
        operators[2] = extraOperators[0];

        OperatorSetUpdate.Message memory message =
            OperatorSetUpdate.Message(registry.nextOperatorUpdateId(), 0, operators);

        BN254.G1Point[] memory nonSignerPubkeys = new BN254.G1Point[](2);
        nonSignerPubkeys[0] = initialOperators[3].pubkey;
        nonSignerPubkeys[1] = initialOperators[2].pubkey;

        Operators.SignatureInfo memory signatureInfo = Operators.SignatureInfo({
            nonSignerPubkeys: nonSignerPubkeys,
            apkG2: BN254.G2Point(
                [
                    2907061045990700054725359562461748672178882330951313836195327790258647029271,
                    6281563211311446916608055318441209721476802551492427187203759791135728100948
                ],
                [
                    16914633983767821662837413448515342510443742248193301243910656017619171484704,
                    19066719044691333956823624407701006018002836358629451345855468619321548553433
                ]
                ),
            sigma: BN254.hashToG1(message.hash()).scalar_mul(
                10871270083209376487778842013958292562863808577713565975978123572762179443915
                )
        });

        vm.expectRevert("Quorum not met");

        registry.updateOperatorSet(message, signatureInfo);
    }

    function test_updateOperatorSet_RevertWhen_WrongMessageId() public {
        Operators.Operator[] memory operators = new Operators.Operator[](3);

        operators[0] = Operators.Operator(initialOperators[3].pubkey, 0);
        operators[1] = Operators.Operator(initialOperators[2].pubkey, 3 * DEFAULT_WEIGHT);
        operators[2] = extraOperators[0];

        OperatorSetUpdate.Message memory message =
            OperatorSetUpdate.Message(registry.nextOperatorUpdateId() + 1, 0, operators);

        BN254.G1Point[] memory nonSignerPubkeys = new BN254.G1Point[](2);
        nonSignerPubkeys[0] = initialOperators[3].pubkey;
        nonSignerPubkeys[1] = initialOperators[2].pubkey;

        Operators.SignatureInfo memory signatureInfo = Operators.SignatureInfo({
            nonSignerPubkeys: nonSignerPubkeys,
            apkG2: BN254.G2Point(
                [
                    2907061045990700054725359562461748672178882330951313836195327790258647029271,
                    6281563211311446916608055318441209721476802551492427187203759791135728100948
                ],
                [
                    16914633983767821662837413448515342510443742248193301243910656017619171484704,
                    19066719044691333956823624407701006018002836358629451345855468619321548553433
                ]
                ),
            sigma: BN254.hashToG1(message.hash()).scalar_mul(
                10871270083209376487778842013958292562863808577713565975978123572762179443915
                )
        });

        vm.expectRevert("Wrong message ID");

        registry.updateOperatorSet(message, signatureInfo);
    }

    function test_updateStateRoot() public {
        StateRootUpdate.Message memory message = StateRootUpdate.Message(0, 1, 0, keccak256(hex"f00d"));

        BN254.G1Point[] memory nonSignerPubkeys = new BN254.G1Point[](1);
        nonSignerPubkeys[0] = initialOperators[3].pubkey;

        Operators.SignatureInfo memory signatureInfo = Operators.SignatureInfo({
            nonSignerPubkeys: nonSignerPubkeys,
            apkG2: BN254.G2Point(
                [
                    21774854595736935906777183372431491423672246101465086449723107940773462536091,
                    11859388993407979358677113204795514610964422675159446451278647734574620707784
                ],
                [
                    3453374196609277266042659107600871924832557088868662992636101033001416801985,
                    2630500117064331827715800222355515273572786883080373379723474133051328147838
                ]
                ),
            sigma: BN254.hashToG1(message.hash()).scalar_mul(
                6305737925830641523797682626723526790077499630761662964405387941160208990354
                )
        });

        assertEq(registry.getStateRoot(message.rollupId, message.blockHeight), bytes32(0));

        vm.expectEmit(true, false, false, true);
        emit StateRootUpdated(message.rollupId, message.blockHeight, message.stateRoot);

        registry.updateStateRoot(message, signatureInfo);

        assertEq(registry.getStateRoot(message.rollupId, message.blockHeight), message.stateRoot);
    }

    function test_updateStateRoot_RevertWhen_QuorumNotMet() public {
        StateRootUpdate.Message memory message = StateRootUpdate.Message(0, 1, 0, keccak256(hex"f00d"));

        BN254.G1Point[] memory nonSignerPubkeys = new BN254.G1Point[](2);
        nonSignerPubkeys[0] = initialOperators[3].pubkey;
        nonSignerPubkeys[1] = initialOperators[2].pubkey;

        Operators.SignatureInfo memory signatureInfo = Operators.SignatureInfo({
            nonSignerPubkeys: nonSignerPubkeys,
            apkG2: BN254.G2Point(
                [
                    2907061045990700054725359562461748672178882330951313836195327790258647029271,
                    6281563211311446916608055318441209721476802551492427187203759791135728100948
                ],
                [
                    16914633983767821662837413448515342510443742248193301243910656017619171484704,
                    19066719044691333956823624407701006018002836358629451345855468619321548553433
                ]
                ),
            sigma: BN254.hashToG1(message.hash()).scalar_mul(
                10871270083209376487778842013958292562863808577713565975978123572762179443915
                )
        });

        vm.expectRevert("Quorum not met");

        registry.updateStateRoot(message, signatureInfo);
    }

    function test_setQuorumThreshold() public {
        assertEq(registry.getQuorumThreshold(), QUORUM_THRESHOLD);

        uint128 denominator = registry.THRESHOLD_DENOMINATOR();

        vm.expectEmit(true, false, false, false);
        emit QuorumThresholdUpdated(denominator - 1);

        vm.prank(addr("owner"));
        registry.setQuorumThreshold(denominator - 1);

        assertEq(registry.getQuorumThreshold(), denominator - 1);
    }

    function test_setQuorumThreshold_RevertWhen_CallerNotOwner() public {
        vm.expectRevert("Ownable: caller is not the owner");

        registry.setQuorumThreshold(1000);
    }

    function test_setQuorumThreshold_RevertWhen_ThresholdGreaterThanDenominator() public {
        uint128 denominator = registry.THRESHOLD_DENOMINATOR();

        vm.expectRevert("Quorum threshold greater than denominator");

        vm.prank(addr("owner"));
        registry.setQuorumThreshold(denominator + 1);
    }
}
