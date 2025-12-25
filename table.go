package vte

// THIS FILE WAS AUTO-GENERATED. DO NOT MODIFY.

const (
	AnywhereState           State = 0
	CsiEntryState           State = 1
	CsiIgnoreState          State = 2
	CsiIntermediateState    State = 3
	CsiParamState           State = 4
	DcsEntryState           State = 5
	DcsIgnoreState          State = 6
	DcsIntermediateState    State = 7
	DcsParamState           State = 8
	DcsPassthroughState     State = 9
	EscapeState             State = 10
	EscapeIntermediateState State = 11
	GroundState             State = 12
	OscStringState          State = 13
	SosPmApcStringState     State = 14
	Utf8State               State = 15

	NoneAction          Action = 0
	ClearAction         Action = 1
	CollectAction       Action = 2
	CsiDispatchAction   Action = 3
	EscDispatchAction   Action = 4
	ExecuteAction       Action = 5
	HookAction          Action = 6
	IgnoreAction        Action = 7
	OscEndAction        Action = 8
	OscPutAction        Action = 9
	OscStartAction      Action = 10
	ParamAction         Action = 11
	PrintAction         Action = 12
	PutAction           Action = 13
	UnhookAction        Action = 14
	BeginUtf8Action     Action = 15
	SosPmApcStartAction Action = 16
	SosPmApcPutAction   Action = 17
	SosPmApcEndAction   Action = 18
)

var (
	stateNames = []string{
		"AnywhereState",
		"CsiEntryState",
		"CsiIgnoreState",
		"CsiIntermediateState",
		"CsiParamState",
		"DcsEntryState",
		"DcsIgnoreState",
		"DcsIntermediateState",
		"DcsParamState",
		"DcsPassthroughState",
		"EscapeState",
		"EscapeIntermediateState",
		"GroundState",
		"OscStringState",
		"SosPmApcStringState",
		"Utf8State",
	}

	stateTable = [][]uint16{
		{ //State AnywhereState = 0
			0, // 00
			0, // 01
			0, // 02
			0, // 03
			0, // 04
			0, // 05
			0, // 06
			0, // 07
			0, // 08
			0, // 09
			0, // 0A
			0, // 0B
			0, // 0C
			0, // 0D
			0, // 0E
			0, // 0F
			0, // 10
			0, // 11
			0, // 12
			0, // 13
			0, // 14
			0, // 15
			0, // 16
			0, // 17
			uint16(GroundState) | (uint16(ExecuteAction) << 8), // 18
			0, // 19
			uint16(GroundState) | (uint16(ExecuteAction) << 8), // 1A
			uint16(EscapeState) | (uint16(NoneAction) << 8),    // 1B
			0, // 1C
			0, // 1D
			0, // 1E
			0, // 1F
			0, // 20
			0, // 21
			0, // 22
			0, // 23
			0, // 24
			0, // 25
			0, // 26
			0, // 27
			0, // 28
			0, // 29
			0, // 2A
			0, // 2B
			0, // 2C
			0, // 2D
			0, // 2E
			0, // 2F
			0, // 30
			0, // 31
			0, // 32
			0, // 33
			0, // 34
			0, // 35
			0, // 36
			0, // 37
			0, // 38
			0, // 39
			0, // 3A
			0, // 3B
			0, // 3C
			0, // 3D
			0, // 3E
			0, // 3F
			0, // 40
			0, // 41
			0, // 42
			0, // 43
			0, // 44
			0, // 45
			0, // 46
			0, // 47
			0, // 48
			0, // 49
			0, // 4A
			0, // 4B
			0, // 4C
			0, // 4D
			0, // 4E
			0, // 4F
			0, // 50
			0, // 51
			0, // 52
			0, // 53
			0, // 54
			0, // 55
			0, // 56
			0, // 57
			0, // 58
			0, // 59
			0, // 5A
			0, // 5B
			0, // 5C
			0, // 5D
			0, // 5E
			0, // 5F
			0, // 60
			0, // 61
			0, // 62
			0, // 63
			0, // 64
			0, // 65
			0, // 66
			0, // 67
			0, // 68
			0, // 69
			0, // 6A
			0, // 6B
			0, // 6C
			0, // 6D
			0, // 6E
			0, // 6F
			0, // 70
			0, // 71
			0, // 72
			0, // 73
			0, // 74
			0, // 75
			0, // 76
			0, // 77
			0, // 78
			0, // 79
			0, // 7A
			0, // 7B
			0, // 7C
			0, // 7D
			0, // 7E
			0, // 7F
			0, // 80
			0, // 81
			0, // 82
			0, // 83
			0, // 84
			0, // 85
			0, // 86
			0, // 87
			0, // 88
			0, // 89
			0, // 8A
			0, // 8B
			0, // 8C
			0, // 8D
			0, // 8E
			0, // 8F
			0, // 90
			0, // 91
			0, // 92
			0, // 93
			0, // 94
			0, // 95
			0, // 96
			0, // 97
			0, // 98
			0, // 99
			0, // 9A
			0, // 9B
			0, // 9C
			0, // 9D
			0, // 9E
			0, // 9F
			0, // A0
			0, // A1
			0, // A2
			0, // A3
			0, // A4
			0, // A5
			0, // A6
			0, // A7
			0, // A8
			0, // A9
			0, // AA
			0, // AB
			0, // AC
			0, // AD
			0, // AE
			0, // AF
			0, // B0
			0, // B1
			0, // B2
			0, // B3
			0, // B4
			0, // B5
			0, // B6
			0, // B7
			0, // B8
			0, // B9
			0, // BA
			0, // BB
			0, // BC
			0, // BD
			0, // BE
			0, // BF
			0, // C0
			0, // C1
			0, // C2
			0, // C3
			0, // C4
			0, // C5
			0, // C6
			0, // C7
			0, // C8
			0, // C9
			0, // CA
			0, // CB
			0, // CC
			0, // CD
			0, // CE
			0, // CF
			0, // D0
			0, // D1
			0, // D2
			0, // D3
			0, // D4
			0, // D5
			0, // D6
			0, // D7
			0, // D8
			0, // D9
			0, // DA
			0, // DB
			0, // DC
			0, // DD
			0, // DE
			0, // DF
			0, // E0
			0, // E1
			0, // E2
			0, // E3
			0, // E4
			0, // E5
			0, // E6
			0, // E7
			0, // E8
			0, // E9
			0, // EA
			0, // EB
			0, // EC
			0, // ED
			0, // EE
			0, // EF
			0, // F0
			0, // F1
			0, // F2
			0, // F3
			0, // F4
			0, // F5
			0, // F6
			0, // F7
			0, // F8
			0, // F9
			0, // FA
			0, // FB
			0, // FC
			0, // FD
			0, // FE
			0, // FF
		},
		{ //State CsiEntryState = 1
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 00
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 01
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 02
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 03
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 04
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 05
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 06
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 07
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 08
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 09
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 0A
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 0B
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 0C
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 0D
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 0E
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 0F
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 10
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 11
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 12
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 13
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 14
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 15
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 16
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 17
			0, // 18
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 19
			0, // 1A
			0, // 1B
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8),        // 1C
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8),        // 1D
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8),        // 1E
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8),        // 1F
			uint16(CsiIntermediateState) | (uint16(CollectAction) << 8), // 20
			uint16(CsiIntermediateState) | (uint16(CollectAction) << 8), // 21
			uint16(CsiIntermediateState) | (uint16(CollectAction) << 8), // 22
			uint16(CsiIntermediateState) | (uint16(CollectAction) << 8), // 23
			uint16(CsiIntermediateState) | (uint16(CollectAction) << 8), // 24
			uint16(CsiIntermediateState) | (uint16(CollectAction) << 8), // 25
			uint16(CsiIntermediateState) | (uint16(CollectAction) << 8), // 26
			uint16(CsiIntermediateState) | (uint16(CollectAction) << 8), // 27
			uint16(CsiIntermediateState) | (uint16(CollectAction) << 8), // 28
			uint16(CsiIntermediateState) | (uint16(CollectAction) << 8), // 29
			uint16(CsiIntermediateState) | (uint16(CollectAction) << 8), // 2A
			uint16(CsiIntermediateState) | (uint16(CollectAction) << 8), // 2B
			uint16(CsiIntermediateState) | (uint16(CollectAction) << 8), // 2C
			uint16(CsiIntermediateState) | (uint16(CollectAction) << 8), // 2D
			uint16(CsiIntermediateState) | (uint16(CollectAction) << 8), // 2E
			uint16(CsiIntermediateState) | (uint16(CollectAction) << 8), // 2F
			uint16(CsiParamState) | (uint16(ParamAction) << 8),          // 30
			uint16(CsiParamState) | (uint16(ParamAction) << 8),          // 31
			uint16(CsiParamState) | (uint16(ParamAction) << 8),          // 32
			uint16(CsiParamState) | (uint16(ParamAction) << 8),          // 33
			uint16(CsiParamState) | (uint16(ParamAction) << 8),          // 34
			uint16(CsiParamState) | (uint16(ParamAction) << 8),          // 35
			uint16(CsiParamState) | (uint16(ParamAction) << 8),          // 36
			uint16(CsiParamState) | (uint16(ParamAction) << 8),          // 37
			uint16(CsiParamState) | (uint16(ParamAction) << 8),          // 38
			uint16(CsiParamState) | (uint16(ParamAction) << 8),          // 39
			uint16(CsiParamState) | (uint16(ParamAction) << 8),          // 3A
			uint16(CsiParamState) | (uint16(ParamAction) << 8),          // 3B
			uint16(CsiParamState) | (uint16(CollectAction) << 8),        // 3C
			uint16(CsiParamState) | (uint16(CollectAction) << 8),        // 3D
			uint16(CsiParamState) | (uint16(CollectAction) << 8),        // 3E
			uint16(CsiParamState) | (uint16(CollectAction) << 8),        // 3F
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 40
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 41
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 42
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 43
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 44
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 45
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 46
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 47
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 48
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 49
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 4A
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 4B
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 4C
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 4D
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 4E
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 4F
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 50
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 51
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 52
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 53
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 54
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 55
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 56
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 57
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 58
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 59
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 5A
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 5B
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 5C
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 5D
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 5E
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 5F
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 60
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 61
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 62
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 63
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 64
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 65
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 66
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 67
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 68
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 69
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 6A
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 6B
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 6C
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 6D
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 6E
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 6F
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 70
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 71
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 72
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 73
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 74
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 75
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 76
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 77
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 78
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 79
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 7A
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 7B
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 7C
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 7D
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 7E
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8),         // 7F
			0, // 80
			0, // 81
			0, // 82
			0, // 83
			0, // 84
			0, // 85
			0, // 86
			0, // 87
			0, // 88
			0, // 89
			0, // 8A
			0, // 8B
			0, // 8C
			0, // 8D
			0, // 8E
			0, // 8F
			0, // 90
			0, // 91
			0, // 92
			0, // 93
			0, // 94
			0, // 95
			0, // 96
			0, // 97
			0, // 98
			0, // 99
			0, // 9A
			0, // 9B
			0, // 9C
			0, // 9D
			0, // 9E
			0, // 9F
			0, // A0
			0, // A1
			0, // A2
			0, // A3
			0, // A4
			0, // A5
			0, // A6
			0, // A7
			0, // A8
			0, // A9
			0, // AA
			0, // AB
			0, // AC
			0, // AD
			0, // AE
			0, // AF
			0, // B0
			0, // B1
			0, // B2
			0, // B3
			0, // B4
			0, // B5
			0, // B6
			0, // B7
			0, // B8
			0, // B9
			0, // BA
			0, // BB
			0, // BC
			0, // BD
			0, // BE
			0, // BF
			0, // C0
			0, // C1
			0, // C2
			0, // C3
			0, // C4
			0, // C5
			0, // C6
			0, // C7
			0, // C8
			0, // C9
			0, // CA
			0, // CB
			0, // CC
			0, // CD
			0, // CE
			0, // CF
			0, // D0
			0, // D1
			0, // D2
			0, // D3
			0, // D4
			0, // D5
			0, // D6
			0, // D7
			0, // D8
			0, // D9
			0, // DA
			0, // DB
			0, // DC
			0, // DD
			0, // DE
			0, // DF
			0, // E0
			0, // E1
			0, // E2
			0, // E3
			0, // E4
			0, // E5
			0, // E6
			0, // E7
			0, // E8
			0, // E9
			0, // EA
			0, // EB
			0, // EC
			0, // ED
			0, // EE
			0, // EF
			0, // F0
			0, // F1
			0, // F2
			0, // F3
			0, // F4
			0, // F5
			0, // F6
			0, // F7
			0, // F8
			0, // F9
			0, // FA
			0, // FB
			0, // FC
			0, // FD
			0, // FE
			0, // FF
		},
		{ //State CsiIgnoreState = 2
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 00
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 01
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 02
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 03
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 04
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 05
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 06
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 07
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 08
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 09
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 0A
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 0B
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 0C
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 0D
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 0E
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 0F
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 10
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 11
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 12
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 13
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 14
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 15
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 16
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 17
			0, // 18
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 19
			0, // 1A
			0, // 1B
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 1C
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 1D
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 1E
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 1F
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8),  // 20
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8),  // 21
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8),  // 22
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8),  // 23
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8),  // 24
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8),  // 25
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8),  // 26
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8),  // 27
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8),  // 28
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8),  // 29
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8),  // 2A
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8),  // 2B
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8),  // 2C
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8),  // 2D
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8),  // 2E
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8),  // 2F
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8),  // 30
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8),  // 31
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8),  // 32
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8),  // 33
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8),  // 34
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8),  // 35
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8),  // 36
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8),  // 37
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8),  // 38
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8),  // 39
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8),  // 3A
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8),  // 3B
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8),  // 3C
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8),  // 3D
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8),  // 3E
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8),  // 3F
			uint16(GroundState) | (uint16(NoneAction) << 8),      // 40
			uint16(GroundState) | (uint16(NoneAction) << 8),      // 41
			uint16(GroundState) | (uint16(NoneAction) << 8),      // 42
			uint16(GroundState) | (uint16(NoneAction) << 8),      // 43
			uint16(GroundState) | (uint16(NoneAction) << 8),      // 44
			uint16(GroundState) | (uint16(NoneAction) << 8),      // 45
			uint16(GroundState) | (uint16(NoneAction) << 8),      // 46
			uint16(GroundState) | (uint16(NoneAction) << 8),      // 47
			uint16(GroundState) | (uint16(NoneAction) << 8),      // 48
			uint16(GroundState) | (uint16(NoneAction) << 8),      // 49
			uint16(GroundState) | (uint16(NoneAction) << 8),      // 4A
			uint16(GroundState) | (uint16(NoneAction) << 8),      // 4B
			uint16(GroundState) | (uint16(NoneAction) << 8),      // 4C
			uint16(GroundState) | (uint16(NoneAction) << 8),      // 4D
			uint16(GroundState) | (uint16(NoneAction) << 8),      // 4E
			uint16(GroundState) | (uint16(NoneAction) << 8),      // 4F
			uint16(GroundState) | (uint16(NoneAction) << 8),      // 50
			uint16(GroundState) | (uint16(NoneAction) << 8),      // 51
			uint16(GroundState) | (uint16(NoneAction) << 8),      // 52
			uint16(GroundState) | (uint16(NoneAction) << 8),      // 53
			uint16(GroundState) | (uint16(NoneAction) << 8),      // 54
			uint16(GroundState) | (uint16(NoneAction) << 8),      // 55
			uint16(GroundState) | (uint16(NoneAction) << 8),      // 56
			uint16(GroundState) | (uint16(NoneAction) << 8),      // 57
			uint16(GroundState) | (uint16(NoneAction) << 8),      // 58
			uint16(GroundState) | (uint16(NoneAction) << 8),      // 59
			uint16(GroundState) | (uint16(NoneAction) << 8),      // 5A
			uint16(GroundState) | (uint16(NoneAction) << 8),      // 5B
			uint16(GroundState) | (uint16(NoneAction) << 8),      // 5C
			uint16(GroundState) | (uint16(NoneAction) << 8),      // 5D
			uint16(GroundState) | (uint16(NoneAction) << 8),      // 5E
			uint16(GroundState) | (uint16(NoneAction) << 8),      // 5F
			uint16(GroundState) | (uint16(NoneAction) << 8),      // 60
			uint16(GroundState) | (uint16(NoneAction) << 8),      // 61
			uint16(GroundState) | (uint16(NoneAction) << 8),      // 62
			uint16(GroundState) | (uint16(NoneAction) << 8),      // 63
			uint16(GroundState) | (uint16(NoneAction) << 8),      // 64
			uint16(GroundState) | (uint16(NoneAction) << 8),      // 65
			uint16(GroundState) | (uint16(NoneAction) << 8),      // 66
			uint16(GroundState) | (uint16(NoneAction) << 8),      // 67
			uint16(GroundState) | (uint16(NoneAction) << 8),      // 68
			uint16(GroundState) | (uint16(NoneAction) << 8),      // 69
			uint16(GroundState) | (uint16(NoneAction) << 8),      // 6A
			uint16(GroundState) | (uint16(NoneAction) << 8),      // 6B
			uint16(GroundState) | (uint16(NoneAction) << 8),      // 6C
			uint16(GroundState) | (uint16(NoneAction) << 8),      // 6D
			uint16(GroundState) | (uint16(NoneAction) << 8),      // 6E
			uint16(GroundState) | (uint16(NoneAction) << 8),      // 6F
			uint16(GroundState) | (uint16(NoneAction) << 8),      // 70
			uint16(GroundState) | (uint16(NoneAction) << 8),      // 71
			uint16(GroundState) | (uint16(NoneAction) << 8),      // 72
			uint16(GroundState) | (uint16(NoneAction) << 8),      // 73
			uint16(GroundState) | (uint16(NoneAction) << 8),      // 74
			uint16(GroundState) | (uint16(NoneAction) << 8),      // 75
			uint16(GroundState) | (uint16(NoneAction) << 8),      // 76
			uint16(GroundState) | (uint16(NoneAction) << 8),      // 77
			uint16(GroundState) | (uint16(NoneAction) << 8),      // 78
			uint16(GroundState) | (uint16(NoneAction) << 8),      // 79
			uint16(GroundState) | (uint16(NoneAction) << 8),      // 7A
			uint16(GroundState) | (uint16(NoneAction) << 8),      // 7B
			uint16(GroundState) | (uint16(NoneAction) << 8),      // 7C
			uint16(GroundState) | (uint16(NoneAction) << 8),      // 7D
			uint16(GroundState) | (uint16(NoneAction) << 8),      // 7E
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8),  // 7F
			0, // 80
			0, // 81
			0, // 82
			0, // 83
			0, // 84
			0, // 85
			0, // 86
			0, // 87
			0, // 88
			0, // 89
			0, // 8A
			0, // 8B
			0, // 8C
			0, // 8D
			0, // 8E
			0, // 8F
			0, // 90
			0, // 91
			0, // 92
			0, // 93
			0, // 94
			0, // 95
			0, // 96
			0, // 97
			0, // 98
			0, // 99
			0, // 9A
			0, // 9B
			0, // 9C
			0, // 9D
			0, // 9E
			0, // 9F
			0, // A0
			0, // A1
			0, // A2
			0, // A3
			0, // A4
			0, // A5
			0, // A6
			0, // A7
			0, // A8
			0, // A9
			0, // AA
			0, // AB
			0, // AC
			0, // AD
			0, // AE
			0, // AF
			0, // B0
			0, // B1
			0, // B2
			0, // B3
			0, // B4
			0, // B5
			0, // B6
			0, // B7
			0, // B8
			0, // B9
			0, // BA
			0, // BB
			0, // BC
			0, // BD
			0, // BE
			0, // BF
			0, // C0
			0, // C1
			0, // C2
			0, // C3
			0, // C4
			0, // C5
			0, // C6
			0, // C7
			0, // C8
			0, // C9
			0, // CA
			0, // CB
			0, // CC
			0, // CD
			0, // CE
			0, // CF
			0, // D0
			0, // D1
			0, // D2
			0, // D3
			0, // D4
			0, // D5
			0, // D6
			0, // D7
			0, // D8
			0, // D9
			0, // DA
			0, // DB
			0, // DC
			0, // DD
			0, // DE
			0, // DF
			0, // E0
			0, // E1
			0, // E2
			0, // E3
			0, // E4
			0, // E5
			0, // E6
			0, // E7
			0, // E8
			0, // E9
			0, // EA
			0, // EB
			0, // EC
			0, // ED
			0, // EE
			0, // EF
			0, // F0
			0, // F1
			0, // F2
			0, // F3
			0, // F4
			0, // F5
			0, // F6
			0, // F7
			0, // F8
			0, // F9
			0, // FA
			0, // FB
			0, // FC
			0, // FD
			0, // FE
			0, // FF
		},
		{ //State CsiIntermediateState = 3
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 00
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 01
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 02
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 03
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 04
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 05
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 06
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 07
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 08
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 09
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 0A
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 0B
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 0C
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 0D
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 0E
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 0F
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 10
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 11
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 12
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 13
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 14
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 15
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 16
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 17
			0, // 18
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 19
			0, // 1A
			0, // 1B
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8),   // 1C
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8),   // 1D
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8),   // 1E
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8),   // 1F
			uint16(AnywhereState) | (uint16(CollectAction) << 8),   // 20
			uint16(AnywhereState) | (uint16(CollectAction) << 8),   // 21
			uint16(AnywhereState) | (uint16(CollectAction) << 8),   // 22
			uint16(AnywhereState) | (uint16(CollectAction) << 8),   // 23
			uint16(AnywhereState) | (uint16(CollectAction) << 8),   // 24
			uint16(AnywhereState) | (uint16(CollectAction) << 8),   // 25
			uint16(AnywhereState) | (uint16(CollectAction) << 8),   // 26
			uint16(AnywhereState) | (uint16(CollectAction) << 8),   // 27
			uint16(AnywhereState) | (uint16(CollectAction) << 8),   // 28
			uint16(AnywhereState) | (uint16(CollectAction) << 8),   // 29
			uint16(AnywhereState) | (uint16(CollectAction) << 8),   // 2A
			uint16(AnywhereState) | (uint16(CollectAction) << 8),   // 2B
			uint16(AnywhereState) | (uint16(CollectAction) << 8),   // 2C
			uint16(AnywhereState) | (uint16(CollectAction) << 8),   // 2D
			uint16(AnywhereState) | (uint16(CollectAction) << 8),   // 2E
			uint16(AnywhereState) | (uint16(CollectAction) << 8),   // 2F
			uint16(CsiIgnoreState) | (uint16(NoneAction) << 8),     // 30
			uint16(CsiIgnoreState) | (uint16(NoneAction) << 8),     // 31
			uint16(CsiIgnoreState) | (uint16(NoneAction) << 8),     // 32
			uint16(CsiIgnoreState) | (uint16(NoneAction) << 8),     // 33
			uint16(CsiIgnoreState) | (uint16(NoneAction) << 8),     // 34
			uint16(CsiIgnoreState) | (uint16(NoneAction) << 8),     // 35
			uint16(CsiIgnoreState) | (uint16(NoneAction) << 8),     // 36
			uint16(CsiIgnoreState) | (uint16(NoneAction) << 8),     // 37
			uint16(CsiIgnoreState) | (uint16(NoneAction) << 8),     // 38
			uint16(CsiIgnoreState) | (uint16(NoneAction) << 8),     // 39
			uint16(CsiIgnoreState) | (uint16(NoneAction) << 8),     // 3A
			uint16(CsiIgnoreState) | (uint16(NoneAction) << 8),     // 3B
			uint16(CsiIgnoreState) | (uint16(NoneAction) << 8),     // 3C
			uint16(CsiIgnoreState) | (uint16(NoneAction) << 8),     // 3D
			uint16(CsiIgnoreState) | (uint16(NoneAction) << 8),     // 3E
			uint16(CsiIgnoreState) | (uint16(NoneAction) << 8),     // 3F
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8), // 40
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8), // 41
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8), // 42
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8), // 43
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8), // 44
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8), // 45
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8), // 46
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8), // 47
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8), // 48
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8), // 49
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8), // 4A
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8), // 4B
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8), // 4C
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8), // 4D
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8), // 4E
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8), // 4F
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8), // 50
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8), // 51
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8), // 52
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8), // 53
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8), // 54
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8), // 55
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8), // 56
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8), // 57
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8), // 58
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8), // 59
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8), // 5A
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8), // 5B
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8), // 5C
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8), // 5D
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8), // 5E
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8), // 5F
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8), // 60
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8), // 61
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8), // 62
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8), // 63
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8), // 64
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8), // 65
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8), // 66
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8), // 67
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8), // 68
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8), // 69
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8), // 6A
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8), // 6B
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8), // 6C
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8), // 6D
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8), // 6E
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8), // 6F
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8), // 70
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8), // 71
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8), // 72
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8), // 73
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8), // 74
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8), // 75
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8), // 76
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8), // 77
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8), // 78
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8), // 79
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8), // 7A
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8), // 7B
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8), // 7C
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8), // 7D
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8), // 7E
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8),    // 7F
			0, // 80
			0, // 81
			0, // 82
			0, // 83
			0, // 84
			0, // 85
			0, // 86
			0, // 87
			0, // 88
			0, // 89
			0, // 8A
			0, // 8B
			0, // 8C
			0, // 8D
			0, // 8E
			0, // 8F
			0, // 90
			0, // 91
			0, // 92
			0, // 93
			0, // 94
			0, // 95
			0, // 96
			0, // 97
			0, // 98
			0, // 99
			0, // 9A
			0, // 9B
			0, // 9C
			0, // 9D
			0, // 9E
			0, // 9F
			0, // A0
			0, // A1
			0, // A2
			0, // A3
			0, // A4
			0, // A5
			0, // A6
			0, // A7
			0, // A8
			0, // A9
			0, // AA
			0, // AB
			0, // AC
			0, // AD
			0, // AE
			0, // AF
			0, // B0
			0, // B1
			0, // B2
			0, // B3
			0, // B4
			0, // B5
			0, // B6
			0, // B7
			0, // B8
			0, // B9
			0, // BA
			0, // BB
			0, // BC
			0, // BD
			0, // BE
			0, // BF
			0, // C0
			0, // C1
			0, // C2
			0, // C3
			0, // C4
			0, // C5
			0, // C6
			0, // C7
			0, // C8
			0, // C9
			0, // CA
			0, // CB
			0, // CC
			0, // CD
			0, // CE
			0, // CF
			0, // D0
			0, // D1
			0, // D2
			0, // D3
			0, // D4
			0, // D5
			0, // D6
			0, // D7
			0, // D8
			0, // D9
			0, // DA
			0, // DB
			0, // DC
			0, // DD
			0, // DE
			0, // DF
			0, // E0
			0, // E1
			0, // E2
			0, // E3
			0, // E4
			0, // E5
			0, // E6
			0, // E7
			0, // E8
			0, // E9
			0, // EA
			0, // EB
			0, // EC
			0, // ED
			0, // EE
			0, // EF
			0, // F0
			0, // F1
			0, // F2
			0, // F3
			0, // F4
			0, // F5
			0, // F6
			0, // F7
			0, // F8
			0, // F9
			0, // FA
			0, // FB
			0, // FC
			0, // FD
			0, // FE
			0, // FF
		},
		{ //State CsiParamState = 4
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 00
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 01
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 02
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 03
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 04
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 05
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 06
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 07
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 08
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 09
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 0A
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 0B
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 0C
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 0D
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 0E
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 0F
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 10
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 11
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 12
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 13
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 14
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 15
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 16
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 17
			0, // 18
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 19
			0, // 1A
			0, // 1B
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8),        // 1C
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8),        // 1D
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8),        // 1E
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8),        // 1F
			uint16(CsiIntermediateState) | (uint16(CollectAction) << 8), // 20
			uint16(CsiIntermediateState) | (uint16(CollectAction) << 8), // 21
			uint16(CsiIntermediateState) | (uint16(CollectAction) << 8), // 22
			uint16(CsiIntermediateState) | (uint16(CollectAction) << 8), // 23
			uint16(CsiIntermediateState) | (uint16(CollectAction) << 8), // 24
			uint16(CsiIntermediateState) | (uint16(CollectAction) << 8), // 25
			uint16(CsiIntermediateState) | (uint16(CollectAction) << 8), // 26
			uint16(CsiIntermediateState) | (uint16(CollectAction) << 8), // 27
			uint16(CsiIntermediateState) | (uint16(CollectAction) << 8), // 28
			uint16(CsiIntermediateState) | (uint16(CollectAction) << 8), // 29
			uint16(CsiIntermediateState) | (uint16(CollectAction) << 8), // 2A
			uint16(CsiIntermediateState) | (uint16(CollectAction) << 8), // 2B
			uint16(CsiIntermediateState) | (uint16(CollectAction) << 8), // 2C
			uint16(CsiIntermediateState) | (uint16(CollectAction) << 8), // 2D
			uint16(CsiIntermediateState) | (uint16(CollectAction) << 8), // 2E
			uint16(CsiIntermediateState) | (uint16(CollectAction) << 8), // 2F
			uint16(AnywhereState) | (uint16(ParamAction) << 8),          // 30
			uint16(AnywhereState) | (uint16(ParamAction) << 8),          // 31
			uint16(AnywhereState) | (uint16(ParamAction) << 8),          // 32
			uint16(AnywhereState) | (uint16(ParamAction) << 8),          // 33
			uint16(AnywhereState) | (uint16(ParamAction) << 8),          // 34
			uint16(AnywhereState) | (uint16(ParamAction) << 8),          // 35
			uint16(AnywhereState) | (uint16(ParamAction) << 8),          // 36
			uint16(AnywhereState) | (uint16(ParamAction) << 8),          // 37
			uint16(AnywhereState) | (uint16(ParamAction) << 8),          // 38
			uint16(AnywhereState) | (uint16(ParamAction) << 8),          // 39
			uint16(AnywhereState) | (uint16(ParamAction) << 8),          // 3A
			uint16(AnywhereState) | (uint16(ParamAction) << 8),          // 3B
			uint16(CsiIgnoreState) | (uint16(NoneAction) << 8),          // 3C
			uint16(CsiIgnoreState) | (uint16(NoneAction) << 8),          // 3D
			uint16(CsiIgnoreState) | (uint16(NoneAction) << 8),          // 3E
			uint16(CsiIgnoreState) | (uint16(NoneAction) << 8),          // 3F
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 40
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 41
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 42
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 43
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 44
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 45
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 46
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 47
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 48
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 49
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 4A
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 4B
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 4C
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 4D
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 4E
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 4F
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 50
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 51
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 52
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 53
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 54
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 55
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 56
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 57
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 58
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 59
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 5A
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 5B
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 5C
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 5D
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 5E
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 5F
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 60
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 61
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 62
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 63
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 64
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 65
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 66
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 67
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 68
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 69
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 6A
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 6B
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 6C
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 6D
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 6E
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 6F
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 70
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 71
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 72
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 73
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 74
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 75
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 76
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 77
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 78
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 79
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 7A
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 7B
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 7C
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 7D
			uint16(GroundState) | (uint16(CsiDispatchAction) << 8),      // 7E
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8),         // 7F
			0, // 80
			0, // 81
			0, // 82
			0, // 83
			0, // 84
			0, // 85
			0, // 86
			0, // 87
			0, // 88
			0, // 89
			0, // 8A
			0, // 8B
			0, // 8C
			0, // 8D
			0, // 8E
			0, // 8F
			0, // 90
			0, // 91
			0, // 92
			0, // 93
			0, // 94
			0, // 95
			0, // 96
			0, // 97
			0, // 98
			0, // 99
			0, // 9A
			0, // 9B
			0, // 9C
			0, // 9D
			0, // 9E
			0, // 9F
			0, // A0
			0, // A1
			0, // A2
			0, // A3
			0, // A4
			0, // A5
			0, // A6
			0, // A7
			0, // A8
			0, // A9
			0, // AA
			0, // AB
			0, // AC
			0, // AD
			0, // AE
			0, // AF
			0, // B0
			0, // B1
			0, // B2
			0, // B3
			0, // B4
			0, // B5
			0, // B6
			0, // B7
			0, // B8
			0, // B9
			0, // BA
			0, // BB
			0, // BC
			0, // BD
			0, // BE
			0, // BF
			0, // C0
			0, // C1
			0, // C2
			0, // C3
			0, // C4
			0, // C5
			0, // C6
			0, // C7
			0, // C8
			0, // C9
			0, // CA
			0, // CB
			0, // CC
			0, // CD
			0, // CE
			0, // CF
			0, // D0
			0, // D1
			0, // D2
			0, // D3
			0, // D4
			0, // D5
			0, // D6
			0, // D7
			0, // D8
			0, // D9
			0, // DA
			0, // DB
			0, // DC
			0, // DD
			0, // DE
			0, // DF
			0, // E0
			0, // E1
			0, // E2
			0, // E3
			0, // E4
			0, // E5
			0, // E6
			0, // E7
			0, // E8
			0, // E9
			0, // EA
			0, // EB
			0, // EC
			0, // ED
			0, // EE
			0, // EF
			0, // F0
			0, // F1
			0, // F2
			0, // F3
			0, // F4
			0, // F5
			0, // F6
			0, // F7
			0, // F8
			0, // F9
			0, // FA
			0, // FB
			0, // FC
			0, // FD
			0, // FE
			0, // FF
		},
		{ //State DcsEntryState = 5
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 00
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 01
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 02
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 03
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 04
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 05
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 06
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 07
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 08
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 09
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 0A
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 0B
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 0C
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 0D
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 0E
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 0F
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 10
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 11
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 12
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 13
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 14
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 15
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 16
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 17
			0, // 18
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 19
			0, // 1A
			0, // 1B
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8),         // 1C
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8),         // 1D
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8),         // 1E
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8),         // 1F
			uint16(DcsIntermediateState) | (uint16(CollectAction) << 8), // 20
			uint16(DcsIntermediateState) | (uint16(CollectAction) << 8), // 21
			uint16(DcsIntermediateState) | (uint16(CollectAction) << 8), // 22
			uint16(DcsIntermediateState) | (uint16(CollectAction) << 8), // 23
			uint16(DcsIntermediateState) | (uint16(CollectAction) << 8), // 24
			uint16(DcsIntermediateState) | (uint16(CollectAction) << 8), // 25
			uint16(DcsIntermediateState) | (uint16(CollectAction) << 8), // 26
			uint16(DcsIntermediateState) | (uint16(CollectAction) << 8), // 27
			uint16(DcsIntermediateState) | (uint16(CollectAction) << 8), // 28
			uint16(DcsIntermediateState) | (uint16(CollectAction) << 8), // 29
			uint16(DcsIntermediateState) | (uint16(CollectAction) << 8), // 2A
			uint16(DcsIntermediateState) | (uint16(CollectAction) << 8), // 2B
			uint16(DcsIntermediateState) | (uint16(CollectAction) << 8), // 2C
			uint16(DcsIntermediateState) | (uint16(CollectAction) << 8), // 2D
			uint16(DcsIntermediateState) | (uint16(CollectAction) << 8), // 2E
			uint16(DcsIntermediateState) | (uint16(CollectAction) << 8), // 2F
			uint16(DcsParamState) | (uint16(ParamAction) << 8),          // 30
			uint16(DcsParamState) | (uint16(ParamAction) << 8),          // 31
			uint16(DcsParamState) | (uint16(ParamAction) << 8),          // 32
			uint16(DcsParamState) | (uint16(ParamAction) << 8),          // 33
			uint16(DcsParamState) | (uint16(ParamAction) << 8),          // 34
			uint16(DcsParamState) | (uint16(ParamAction) << 8),          // 35
			uint16(DcsParamState) | (uint16(ParamAction) << 8),          // 36
			uint16(DcsParamState) | (uint16(ParamAction) << 8),          // 37
			uint16(DcsParamState) | (uint16(ParamAction) << 8),          // 38
			uint16(DcsParamState) | (uint16(ParamAction) << 8),          // 39
			uint16(DcsParamState) | (uint16(ParamAction) << 8),          // 3A
			uint16(DcsParamState) | (uint16(ParamAction) << 8),          // 3B
			uint16(DcsParamState) | (uint16(CollectAction) << 8),        // 3C
			uint16(DcsParamState) | (uint16(CollectAction) << 8),        // 3D
			uint16(DcsParamState) | (uint16(CollectAction) << 8),        // 3E
			uint16(DcsParamState) | (uint16(CollectAction) << 8),        // 3F
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 40
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 41
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 42
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 43
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 44
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 45
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 46
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 47
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 48
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 49
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 4A
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 4B
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 4C
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 4D
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 4E
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 4F
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 50
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 51
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 52
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 53
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 54
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 55
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 56
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 57
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 58
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 59
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 5A
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 5B
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 5C
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 5D
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 5E
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 5F
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 60
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 61
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 62
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 63
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 64
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 65
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 66
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 67
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 68
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 69
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 6A
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 6B
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 6C
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 6D
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 6E
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 6F
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 70
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 71
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 72
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 73
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 74
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 75
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 76
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 77
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 78
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 79
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 7A
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 7B
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 7C
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 7D
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 7E
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8),         // 7F
			0, // 80
			0, // 81
			0, // 82
			0, // 83
			0, // 84
			0, // 85
			0, // 86
			0, // 87
			0, // 88
			0, // 89
			0, // 8A
			0, // 8B
			0, // 8C
			0, // 8D
			0, // 8E
			0, // 8F
			0, // 90
			0, // 91
			0, // 92
			0, // 93
			0, // 94
			0, // 95
			0, // 96
			0, // 97
			0, // 98
			0, // 99
			0, // 9A
			0, // 9B
			0, // 9C
			0, // 9D
			0, // 9E
			0, // 9F
			0, // A0
			0, // A1
			0, // A2
			0, // A3
			0, // A4
			0, // A5
			0, // A6
			0, // A7
			0, // A8
			0, // A9
			0, // AA
			0, // AB
			0, // AC
			0, // AD
			0, // AE
			0, // AF
			0, // B0
			0, // B1
			0, // B2
			0, // B3
			0, // B4
			0, // B5
			0, // B6
			0, // B7
			0, // B8
			0, // B9
			0, // BA
			0, // BB
			0, // BC
			0, // BD
			0, // BE
			0, // BF
			0, // C0
			0, // C1
			0, // C2
			0, // C3
			0, // C4
			0, // C5
			0, // C6
			0, // C7
			0, // C8
			0, // C9
			0, // CA
			0, // CB
			0, // CC
			0, // CD
			0, // CE
			0, // CF
			0, // D0
			0, // D1
			0, // D2
			0, // D3
			0, // D4
			0, // D5
			0, // D6
			0, // D7
			0, // D8
			0, // D9
			0, // DA
			0, // DB
			0, // DC
			0, // DD
			0, // DE
			0, // DF
			0, // E0
			0, // E1
			0, // E2
			0, // E3
			0, // E4
			0, // E5
			0, // E6
			0, // E7
			0, // E8
			0, // E9
			0, // EA
			0, // EB
			0, // EC
			0, // ED
			0, // EE
			0, // EF
			0, // F0
			0, // F1
			0, // F2
			0, // F3
			0, // F4
			0, // F5
			0, // F6
			0, // F7
			0, // F8
			0, // F9
			0, // FA
			0, // FB
			0, // FC
			0, // FD
			0, // FE
			0, // FF
		},
		{ //State DcsIgnoreState = 6
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 00
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 01
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 02
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 03
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 04
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 05
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 06
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 07
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 08
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 09
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 0A
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 0B
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 0C
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 0D
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 0E
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 0F
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 10
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 11
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 12
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 13
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 14
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 15
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 16
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 17
			0, // 18
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 19
			0, // 1A
			0, // 1B
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 1C
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 1D
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 1E
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 1F
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 20
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 21
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 22
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 23
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 24
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 25
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 26
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 27
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 28
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 29
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 2A
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 2B
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 2C
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 2D
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 2E
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 2F
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 30
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 31
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 32
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 33
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 34
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 35
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 36
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 37
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 38
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 39
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 3A
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 3B
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 3C
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 3D
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 3E
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 3F
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 40
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 41
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 42
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 43
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 44
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 45
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 46
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 47
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 48
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 49
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 4A
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 4B
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 4C
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 4D
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 4E
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 4F
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 50
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 51
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 52
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 53
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 54
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 55
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 56
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 57
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 58
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 59
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 5A
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 5B
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 5C
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 5D
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 5E
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 5F
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 60
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 61
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 62
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 63
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 64
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 65
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 66
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 67
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 68
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 69
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 6A
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 6B
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 6C
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 6D
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 6E
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 6F
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 70
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 71
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 72
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 73
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 74
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 75
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 76
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 77
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 78
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 79
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 7A
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 7B
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 7C
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 7D
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 7E
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 7F
			0, // 80
			0, // 81
			0, // 82
			0, // 83
			0, // 84
			0, // 85
			0, // 86
			0, // 87
			0, // 88
			0, // 89
			0, // 8A
			0, // 8B
			0, // 8C
			0, // 8D
			0, // 8E
			0, // 8F
			0, // 90
			0, // 91
			0, // 92
			0, // 93
			0, // 94
			0, // 95
			0, // 96
			0, // 97
			0, // 98
			0, // 99
			0, // 9A
			0, // 9B
			uint16(GroundState) | (uint16(NoneAction) << 8), // 9C
			0, // 9D
			0, // 9E
			0, // 9F
			0, // A0
			0, // A1
			0, // A2
			0, // A3
			0, // A4
			0, // A5
			0, // A6
			0, // A7
			0, // A8
			0, // A9
			0, // AA
			0, // AB
			0, // AC
			0, // AD
			0, // AE
			0, // AF
			0, // B0
			0, // B1
			0, // B2
			0, // B3
			0, // B4
			0, // B5
			0, // B6
			0, // B7
			0, // B8
			0, // B9
			0, // BA
			0, // BB
			0, // BC
			0, // BD
			0, // BE
			0, // BF
			0, // C0
			0, // C1
			0, // C2
			0, // C3
			0, // C4
			0, // C5
			0, // C6
			0, // C7
			0, // C8
			0, // C9
			0, // CA
			0, // CB
			0, // CC
			0, // CD
			0, // CE
			0, // CF
			0, // D0
			0, // D1
			0, // D2
			0, // D3
			0, // D4
			0, // D5
			0, // D6
			0, // D7
			0, // D8
			0, // D9
			0, // DA
			0, // DB
			0, // DC
			0, // DD
			0, // DE
			0, // DF
			0, // E0
			0, // E1
			0, // E2
			0, // E3
			0, // E4
			0, // E5
			0, // E6
			0, // E7
			0, // E8
			0, // E9
			0, // EA
			0, // EB
			0, // EC
			0, // ED
			0, // EE
			0, // EF
			0, // F0
			0, // F1
			0, // F2
			0, // F3
			0, // F4
			0, // F5
			0, // F6
			0, // F7
			0, // F8
			0, // F9
			0, // FA
			0, // FB
			0, // FC
			0, // FD
			0, // FE
			0, // FF
		},
		{ //State DcsIntermediateState = 7
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 00
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 01
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 02
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 03
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 04
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 05
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 06
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 07
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 08
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 09
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 0A
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 0B
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 0C
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 0D
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 0E
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 0F
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 10
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 11
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 12
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 13
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 14
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 15
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 16
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 17
			0, // 18
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 19
			0, // 1A
			0, // 1B
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8),     // 1C
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8),     // 1D
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8),     // 1E
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8),     // 1F
			uint16(AnywhereState) | (uint16(CollectAction) << 8),    // 20
			uint16(AnywhereState) | (uint16(CollectAction) << 8),    // 21
			uint16(AnywhereState) | (uint16(CollectAction) << 8),    // 22
			uint16(AnywhereState) | (uint16(CollectAction) << 8),    // 23
			uint16(AnywhereState) | (uint16(CollectAction) << 8),    // 24
			uint16(AnywhereState) | (uint16(CollectAction) << 8),    // 25
			uint16(AnywhereState) | (uint16(CollectAction) << 8),    // 26
			uint16(AnywhereState) | (uint16(CollectAction) << 8),    // 27
			uint16(AnywhereState) | (uint16(CollectAction) << 8),    // 28
			uint16(AnywhereState) | (uint16(CollectAction) << 8),    // 29
			uint16(AnywhereState) | (uint16(CollectAction) << 8),    // 2A
			uint16(AnywhereState) | (uint16(CollectAction) << 8),    // 2B
			uint16(AnywhereState) | (uint16(CollectAction) << 8),    // 2C
			uint16(AnywhereState) | (uint16(CollectAction) << 8),    // 2D
			uint16(AnywhereState) | (uint16(CollectAction) << 8),    // 2E
			uint16(AnywhereState) | (uint16(CollectAction) << 8),    // 2F
			uint16(DcsIgnoreState) | (uint16(NoneAction) << 8),      // 30
			uint16(DcsIgnoreState) | (uint16(NoneAction) << 8),      // 31
			uint16(DcsIgnoreState) | (uint16(NoneAction) << 8),      // 32
			uint16(DcsIgnoreState) | (uint16(NoneAction) << 8),      // 33
			uint16(DcsIgnoreState) | (uint16(NoneAction) << 8),      // 34
			uint16(DcsIgnoreState) | (uint16(NoneAction) << 8),      // 35
			uint16(DcsIgnoreState) | (uint16(NoneAction) << 8),      // 36
			uint16(DcsIgnoreState) | (uint16(NoneAction) << 8),      // 37
			uint16(DcsIgnoreState) | (uint16(NoneAction) << 8),      // 38
			uint16(DcsIgnoreState) | (uint16(NoneAction) << 8),      // 39
			uint16(DcsIgnoreState) | (uint16(NoneAction) << 8),      // 3A
			uint16(DcsIgnoreState) | (uint16(NoneAction) << 8),      // 3B
			uint16(DcsIgnoreState) | (uint16(NoneAction) << 8),      // 3C
			uint16(DcsIgnoreState) | (uint16(NoneAction) << 8),      // 3D
			uint16(DcsIgnoreState) | (uint16(NoneAction) << 8),      // 3E
			uint16(DcsIgnoreState) | (uint16(NoneAction) << 8),      // 3F
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8), // 40
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8), // 41
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8), // 42
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8), // 43
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8), // 44
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8), // 45
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8), // 46
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8), // 47
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8), // 48
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8), // 49
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8), // 4A
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8), // 4B
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8), // 4C
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8), // 4D
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8), // 4E
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8), // 4F
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8), // 50
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8), // 51
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8), // 52
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8), // 53
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8), // 54
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8), // 55
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8), // 56
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8), // 57
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8), // 58
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8), // 59
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8), // 5A
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8), // 5B
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8), // 5C
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8), // 5D
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8), // 5E
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8), // 5F
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8), // 60
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8), // 61
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8), // 62
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8), // 63
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8), // 64
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8), // 65
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8), // 66
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8), // 67
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8), // 68
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8), // 69
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8), // 6A
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8), // 6B
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8), // 6C
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8), // 6D
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8), // 6E
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8), // 6F
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8), // 70
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8), // 71
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8), // 72
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8), // 73
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8), // 74
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8), // 75
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8), // 76
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8), // 77
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8), // 78
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8), // 79
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8), // 7A
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8), // 7B
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8), // 7C
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8), // 7D
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8), // 7E
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8),     // 7F
			0, // 80
			0, // 81
			0, // 82
			0, // 83
			0, // 84
			0, // 85
			0, // 86
			0, // 87
			0, // 88
			0, // 89
			0, // 8A
			0, // 8B
			0, // 8C
			0, // 8D
			0, // 8E
			0, // 8F
			0, // 90
			0, // 91
			0, // 92
			0, // 93
			0, // 94
			0, // 95
			0, // 96
			0, // 97
			0, // 98
			0, // 99
			0, // 9A
			0, // 9B
			0, // 9C
			0, // 9D
			0, // 9E
			0, // 9F
			0, // A0
			0, // A1
			0, // A2
			0, // A3
			0, // A4
			0, // A5
			0, // A6
			0, // A7
			0, // A8
			0, // A9
			0, // AA
			0, // AB
			0, // AC
			0, // AD
			0, // AE
			0, // AF
			0, // B0
			0, // B1
			0, // B2
			0, // B3
			0, // B4
			0, // B5
			0, // B6
			0, // B7
			0, // B8
			0, // B9
			0, // BA
			0, // BB
			0, // BC
			0, // BD
			0, // BE
			0, // BF
			0, // C0
			0, // C1
			0, // C2
			0, // C3
			0, // C4
			0, // C5
			0, // C6
			0, // C7
			0, // C8
			0, // C9
			0, // CA
			0, // CB
			0, // CC
			0, // CD
			0, // CE
			0, // CF
			0, // D0
			0, // D1
			0, // D2
			0, // D3
			0, // D4
			0, // D5
			0, // D6
			0, // D7
			0, // D8
			0, // D9
			0, // DA
			0, // DB
			0, // DC
			0, // DD
			0, // DE
			0, // DF
			0, // E0
			0, // E1
			0, // E2
			0, // E3
			0, // E4
			0, // E5
			0, // E6
			0, // E7
			0, // E8
			0, // E9
			0, // EA
			0, // EB
			0, // EC
			0, // ED
			0, // EE
			0, // EF
			0, // F0
			0, // F1
			0, // F2
			0, // F3
			0, // F4
			0, // F5
			0, // F6
			0, // F7
			0, // F8
			0, // F9
			0, // FA
			0, // FB
			0, // FC
			0, // FD
			0, // FE
			0, // FF
		},
		{ //State DcsParamState = 8
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 00
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 01
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 02
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 03
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 04
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 05
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 06
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 07
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 08
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 09
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 0A
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 0B
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 0C
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 0D
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 0E
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 0F
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 10
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 11
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 12
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 13
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 14
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 15
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 16
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 17
			0, // 18
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 19
			0, // 1A
			0, // 1B
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8),         // 1C
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8),         // 1D
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8),         // 1E
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8),         // 1F
			uint16(DcsIntermediateState) | (uint16(CollectAction) << 8), // 20
			uint16(DcsIntermediateState) | (uint16(CollectAction) << 8), // 21
			uint16(DcsIntermediateState) | (uint16(CollectAction) << 8), // 22
			uint16(DcsIntermediateState) | (uint16(CollectAction) << 8), // 23
			uint16(DcsIntermediateState) | (uint16(CollectAction) << 8), // 24
			uint16(DcsIntermediateState) | (uint16(CollectAction) << 8), // 25
			uint16(DcsIntermediateState) | (uint16(CollectAction) << 8), // 26
			uint16(DcsIntermediateState) | (uint16(CollectAction) << 8), // 27
			uint16(DcsIntermediateState) | (uint16(CollectAction) << 8), // 28
			uint16(DcsIntermediateState) | (uint16(CollectAction) << 8), // 29
			uint16(DcsIntermediateState) | (uint16(CollectAction) << 8), // 2A
			uint16(DcsIntermediateState) | (uint16(CollectAction) << 8), // 2B
			uint16(DcsIntermediateState) | (uint16(CollectAction) << 8), // 2C
			uint16(DcsIntermediateState) | (uint16(CollectAction) << 8), // 2D
			uint16(DcsIntermediateState) | (uint16(CollectAction) << 8), // 2E
			uint16(DcsIntermediateState) | (uint16(CollectAction) << 8), // 2F
			uint16(AnywhereState) | (uint16(ParamAction) << 8),          // 30
			uint16(AnywhereState) | (uint16(ParamAction) << 8),          // 31
			uint16(AnywhereState) | (uint16(ParamAction) << 8),          // 32
			uint16(AnywhereState) | (uint16(ParamAction) << 8),          // 33
			uint16(AnywhereState) | (uint16(ParamAction) << 8),          // 34
			uint16(AnywhereState) | (uint16(ParamAction) << 8),          // 35
			uint16(AnywhereState) | (uint16(ParamAction) << 8),          // 36
			uint16(AnywhereState) | (uint16(ParamAction) << 8),          // 37
			uint16(AnywhereState) | (uint16(ParamAction) << 8),          // 38
			uint16(AnywhereState) | (uint16(ParamAction) << 8),          // 39
			uint16(AnywhereState) | (uint16(ParamAction) << 8),          // 3A
			uint16(AnywhereState) | (uint16(ParamAction) << 8),          // 3B
			uint16(DcsIgnoreState) | (uint16(NoneAction) << 8),          // 3C
			uint16(DcsIgnoreState) | (uint16(NoneAction) << 8),          // 3D
			uint16(DcsIgnoreState) | (uint16(NoneAction) << 8),          // 3E
			uint16(DcsIgnoreState) | (uint16(NoneAction) << 8),          // 3F
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 40
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 41
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 42
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 43
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 44
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 45
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 46
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 47
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 48
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 49
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 4A
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 4B
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 4C
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 4D
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 4E
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 4F
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 50
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 51
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 52
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 53
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 54
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 55
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 56
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 57
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 58
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 59
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 5A
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 5B
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 5C
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 5D
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 5E
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 5F
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 60
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 61
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 62
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 63
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 64
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 65
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 66
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 67
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 68
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 69
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 6A
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 6B
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 6C
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 6D
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 6E
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 6F
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 70
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 71
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 72
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 73
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 74
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 75
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 76
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 77
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 78
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 79
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 7A
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 7B
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 7C
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 7D
			uint16(DcsPassthroughState) | (uint16(NoneAction) << 8),     // 7E
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8),         // 7F
			0, // 80
			0, // 81
			0, // 82
			0, // 83
			0, // 84
			0, // 85
			0, // 86
			0, // 87
			0, // 88
			0, // 89
			0, // 8A
			0, // 8B
			0, // 8C
			0, // 8D
			0, // 8E
			0, // 8F
			0, // 90
			0, // 91
			0, // 92
			0, // 93
			0, // 94
			0, // 95
			0, // 96
			0, // 97
			0, // 98
			0, // 99
			0, // 9A
			0, // 9B
			0, // 9C
			0, // 9D
			0, // 9E
			0, // 9F
			0, // A0
			0, // A1
			0, // A2
			0, // A3
			0, // A4
			0, // A5
			0, // A6
			0, // A7
			0, // A8
			0, // A9
			0, // AA
			0, // AB
			0, // AC
			0, // AD
			0, // AE
			0, // AF
			0, // B0
			0, // B1
			0, // B2
			0, // B3
			0, // B4
			0, // B5
			0, // B6
			0, // B7
			0, // B8
			0, // B9
			0, // BA
			0, // BB
			0, // BC
			0, // BD
			0, // BE
			0, // BF
			0, // C0
			0, // C1
			0, // C2
			0, // C3
			0, // C4
			0, // C5
			0, // C6
			0, // C7
			0, // C8
			0, // C9
			0, // CA
			0, // CB
			0, // CC
			0, // CD
			0, // CE
			0, // CF
			0, // D0
			0, // D1
			0, // D2
			0, // D3
			0, // D4
			0, // D5
			0, // D6
			0, // D7
			0, // D8
			0, // D9
			0, // DA
			0, // DB
			0, // DC
			0, // DD
			0, // DE
			0, // DF
			0, // E0
			0, // E1
			0, // E2
			0, // E3
			0, // E4
			0, // E5
			0, // E6
			0, // E7
			0, // E8
			0, // E9
			0, // EA
			0, // EB
			0, // EC
			0, // ED
			0, // EE
			0, // EF
			0, // F0
			0, // F1
			0, // F2
			0, // F3
			0, // F4
			0, // F5
			0, // F6
			0, // F7
			0, // F8
			0, // F9
			0, // FA
			0, // FB
			0, // FC
			0, // FD
			0, // FE
			0, // FF
		},
		{ //State DcsPassthroughState = 9
			uint16(AnywhereState) | (uint16(PutAction) << 8), // 00
			uint16(AnywhereState) | (uint16(PutAction) << 8), // 01
			uint16(AnywhereState) | (uint16(PutAction) << 8), // 02
			uint16(AnywhereState) | (uint16(PutAction) << 8), // 03
			uint16(AnywhereState) | (uint16(PutAction) << 8), // 04
			uint16(AnywhereState) | (uint16(PutAction) << 8), // 05
			uint16(AnywhereState) | (uint16(PutAction) << 8), // 06
			uint16(AnywhereState) | (uint16(PutAction) << 8), // 07
			uint16(AnywhereState) | (uint16(PutAction) << 8), // 08
			uint16(AnywhereState) | (uint16(PutAction) << 8), // 09
			uint16(AnywhereState) | (uint16(PutAction) << 8), // 0A
			uint16(AnywhereState) | (uint16(PutAction) << 8), // 0B
			uint16(AnywhereState) | (uint16(PutAction) << 8), // 0C
			uint16(AnywhereState) | (uint16(PutAction) << 8), // 0D
			uint16(AnywhereState) | (uint16(PutAction) << 8), // 0E
			uint16(AnywhereState) | (uint16(PutAction) << 8), // 0F
			uint16(AnywhereState) | (uint16(PutAction) << 8), // 10
			uint16(AnywhereState) | (uint16(PutAction) << 8), // 11
			uint16(AnywhereState) | (uint16(PutAction) << 8), // 12
			uint16(AnywhereState) | (uint16(PutAction) << 8), // 13
			uint16(AnywhereState) | (uint16(PutAction) << 8), // 14
			uint16(AnywhereState) | (uint16(PutAction) << 8), // 15
			uint16(AnywhereState) | (uint16(PutAction) << 8), // 16
			uint16(AnywhereState) | (uint16(PutAction) << 8), // 17
			0, // 18
			uint16(AnywhereState) | (uint16(PutAction) << 8), // 19
			0, // 1A
			0, // 1B
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 1C
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 1D
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 1E
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 1F
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 20
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 21
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 22
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 23
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 24
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 25
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 26
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 27
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 28
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 29
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 2A
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 2B
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 2C
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 2D
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 2E
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 2F
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 30
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 31
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 32
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 33
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 34
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 35
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 36
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 37
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 38
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 39
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 3A
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 3B
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 3C
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 3D
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 3E
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 3F
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 40
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 41
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 42
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 43
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 44
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 45
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 46
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 47
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 48
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 49
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 4A
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 4B
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 4C
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 4D
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 4E
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 4F
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 50
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 51
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 52
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 53
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 54
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 55
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 56
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 57
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 58
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 59
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 5A
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 5B
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 5C
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 5D
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 5E
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 5F
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 60
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 61
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 62
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 63
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 64
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 65
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 66
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 67
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 68
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 69
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 6A
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 6B
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 6C
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 6D
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 6E
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 6F
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 70
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 71
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 72
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 73
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 74
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 75
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 76
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 77
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 78
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 79
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 7A
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 7B
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 7C
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 7D
			uint16(AnywhereState) | (uint16(PutAction) << 8),    // 7E
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 7F
			0, // 80
			0, // 81
			0, // 82
			0, // 83
			0, // 84
			0, // 85
			0, // 86
			0, // 87
			0, // 88
			0, // 89
			0, // 8A
			0, // 8B
			0, // 8C
			0, // 8D
			0, // 8E
			0, // 8F
			0, // 90
			0, // 91
			0, // 92
			0, // 93
			0, // 94
			0, // 95
			0, // 96
			0, // 97
			0, // 98
			0, // 99
			0, // 9A
			0, // 9B
			uint16(GroundState) | (uint16(NoneAction) << 8), // 9C
			0, // 9D
			0, // 9E
			0, // 9F
			0, // A0
			0, // A1
			0, // A2
			0, // A3
			0, // A4
			0, // A5
			0, // A6
			0, // A7
			0, // A8
			0, // A9
			0, // AA
			0, // AB
			0, // AC
			0, // AD
			0, // AE
			0, // AF
			0, // B0
			0, // B1
			0, // B2
			0, // B3
			0, // B4
			0, // B5
			0, // B6
			0, // B7
			0, // B8
			0, // B9
			0, // BA
			0, // BB
			0, // BC
			0, // BD
			0, // BE
			0, // BF
			0, // C0
			0, // C1
			0, // C2
			0, // C3
			0, // C4
			0, // C5
			0, // C6
			0, // C7
			0, // C8
			0, // C9
			0, // CA
			0, // CB
			0, // CC
			0, // CD
			0, // CE
			0, // CF
			0, // D0
			0, // D1
			0, // D2
			0, // D3
			0, // D4
			0, // D5
			0, // D6
			0, // D7
			0, // D8
			0, // D9
			0, // DA
			0, // DB
			0, // DC
			0, // DD
			0, // DE
			0, // DF
			0, // E0
			0, // E1
			0, // E2
			0, // E3
			0, // E4
			0, // E5
			0, // E6
			0, // E7
			0, // E8
			0, // E9
			0, // EA
			0, // EB
			0, // EC
			0, // ED
			0, // EE
			0, // EF
			0, // F0
			0, // F1
			0, // F2
			0, // F3
			0, // F4
			0, // F5
			0, // F6
			0, // F7
			0, // F8
			0, // F9
			0, // FA
			0, // FB
			0, // FC
			0, // FD
			0, // FE
			0, // FF
		},
		{ //State EscapeState = 10
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 00
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 01
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 02
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 03
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 04
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 05
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 06
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 07
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 08
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 09
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 0A
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 0B
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 0C
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 0D
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 0E
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 0F
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 10
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 11
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 12
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 13
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 14
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 15
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 16
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 17
			0, // 18
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 19
			0, // 1A
			0, // 1B
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8),             // 1C
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8),             // 1D
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8),             // 1E
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8),             // 1F
			uint16(EscapeIntermediateState) | (uint16(CollectAction) << 8),   // 20
			uint16(EscapeIntermediateState) | (uint16(CollectAction) << 8),   // 21
			uint16(EscapeIntermediateState) | (uint16(CollectAction) << 8),   // 22
			uint16(EscapeIntermediateState) | (uint16(CollectAction) << 8),   // 23
			uint16(EscapeIntermediateState) | (uint16(CollectAction) << 8),   // 24
			uint16(EscapeIntermediateState) | (uint16(CollectAction) << 8),   // 25
			uint16(EscapeIntermediateState) | (uint16(CollectAction) << 8),   // 26
			uint16(EscapeIntermediateState) | (uint16(CollectAction) << 8),   // 27
			uint16(EscapeIntermediateState) | (uint16(CollectAction) << 8),   // 28
			uint16(EscapeIntermediateState) | (uint16(CollectAction) << 8),   // 29
			uint16(EscapeIntermediateState) | (uint16(CollectAction) << 8),   // 2A
			uint16(EscapeIntermediateState) | (uint16(CollectAction) << 8),   // 2B
			uint16(EscapeIntermediateState) | (uint16(CollectAction) << 8),   // 2C
			uint16(EscapeIntermediateState) | (uint16(CollectAction) << 8),   // 2D
			uint16(EscapeIntermediateState) | (uint16(CollectAction) << 8),   // 2E
			uint16(EscapeIntermediateState) | (uint16(CollectAction) << 8),   // 2F
			uint16(GroundState) | (uint16(EscDispatchAction) << 8),           // 30
			uint16(GroundState) | (uint16(EscDispatchAction) << 8),           // 31
			uint16(GroundState) | (uint16(EscDispatchAction) << 8),           // 32
			uint16(GroundState) | (uint16(EscDispatchAction) << 8),           // 33
			uint16(GroundState) | (uint16(EscDispatchAction) << 8),           // 34
			uint16(GroundState) | (uint16(EscDispatchAction) << 8),           // 35
			uint16(GroundState) | (uint16(EscDispatchAction) << 8),           // 36
			uint16(GroundState) | (uint16(EscDispatchAction) << 8),           // 37
			uint16(GroundState) | (uint16(EscDispatchAction) << 8),           // 38
			uint16(GroundState) | (uint16(EscDispatchAction) << 8),           // 39
			uint16(GroundState) | (uint16(EscDispatchAction) << 8),           // 3A
			uint16(GroundState) | (uint16(EscDispatchAction) << 8),           // 3B
			uint16(GroundState) | (uint16(EscDispatchAction) << 8),           // 3C
			uint16(GroundState) | (uint16(EscDispatchAction) << 8),           // 3D
			uint16(GroundState) | (uint16(EscDispatchAction) << 8),           // 3E
			uint16(GroundState) | (uint16(EscDispatchAction) << 8),           // 3F
			uint16(GroundState) | (uint16(EscDispatchAction) << 8),           // 40
			uint16(GroundState) | (uint16(EscDispatchAction) << 8),           // 41
			uint16(GroundState) | (uint16(EscDispatchAction) << 8),           // 42
			uint16(GroundState) | (uint16(EscDispatchAction) << 8),           // 43
			uint16(GroundState) | (uint16(EscDispatchAction) << 8),           // 44
			uint16(GroundState) | (uint16(EscDispatchAction) << 8),           // 45
			uint16(GroundState) | (uint16(EscDispatchAction) << 8),           // 46
			uint16(GroundState) | (uint16(EscDispatchAction) << 8),           // 47
			uint16(GroundState) | (uint16(EscDispatchAction) << 8),           // 48
			uint16(GroundState) | (uint16(EscDispatchAction) << 8),           // 49
			uint16(GroundState) | (uint16(EscDispatchAction) << 8),           // 4A
			uint16(GroundState) | (uint16(EscDispatchAction) << 8),           // 4B
			uint16(GroundState) | (uint16(EscDispatchAction) << 8),           // 4C
			uint16(GroundState) | (uint16(EscDispatchAction) << 8),           // 4D
			uint16(GroundState) | (uint16(EscDispatchAction) << 8),           // 4E
			uint16(GroundState) | (uint16(EscDispatchAction) << 8),           // 4F
			uint16(DcsEntryState) | (uint16(NoneAction) << 8),                // 50
			uint16(GroundState) | (uint16(EscDispatchAction) << 8),           // 51
			uint16(GroundState) | (uint16(EscDispatchAction) << 8),           // 52
			uint16(GroundState) | (uint16(EscDispatchAction) << 8),           // 53
			uint16(GroundState) | (uint16(EscDispatchAction) << 8),           // 54
			uint16(GroundState) | (uint16(EscDispatchAction) << 8),           // 55
			uint16(GroundState) | (uint16(EscDispatchAction) << 8),           // 56
			uint16(GroundState) | (uint16(EscDispatchAction) << 8),           // 57
			uint16(SosPmApcStringState) | (uint16(SosPmApcStartAction) << 8), // 58
			uint16(GroundState) | (uint16(EscDispatchAction) << 8),           // 59
			uint16(GroundState) | (uint16(EscDispatchAction) << 8),           // 5A
			uint16(CsiEntryState) | (uint16(NoneAction) << 8),                // 5B
			uint16(GroundState) | (uint16(EscDispatchAction) << 8),           // 5C
			uint16(OscStringState) | (uint16(NoneAction) << 8),               // 5D
			uint16(SosPmApcStringState) | (uint16(SosPmApcStartAction) << 8), // 5E
			uint16(SosPmApcStringState) | (uint16(SosPmApcStartAction) << 8), // 5F
			uint16(GroundState) | (uint16(EscDispatchAction) << 8),           // 60
			uint16(GroundState) | (uint16(EscDispatchAction) << 8),           // 61
			uint16(GroundState) | (uint16(EscDispatchAction) << 8),           // 62
			uint16(GroundState) | (uint16(EscDispatchAction) << 8),           // 63
			uint16(GroundState) | (uint16(EscDispatchAction) << 8),           // 64
			uint16(GroundState) | (uint16(EscDispatchAction) << 8),           // 65
			uint16(GroundState) | (uint16(EscDispatchAction) << 8),           // 66
			uint16(GroundState) | (uint16(EscDispatchAction) << 8),           // 67
			uint16(GroundState) | (uint16(EscDispatchAction) << 8),           // 68
			uint16(GroundState) | (uint16(EscDispatchAction) << 8),           // 69
			uint16(GroundState) | (uint16(EscDispatchAction) << 8),           // 6A
			uint16(GroundState) | (uint16(EscDispatchAction) << 8),           // 6B
			uint16(GroundState) | (uint16(EscDispatchAction) << 8),           // 6C
			uint16(GroundState) | (uint16(EscDispatchAction) << 8),           // 6D
			uint16(GroundState) | (uint16(EscDispatchAction) << 8),           // 6E
			uint16(GroundState) | (uint16(EscDispatchAction) << 8),           // 6F
			uint16(GroundState) | (uint16(EscDispatchAction) << 8),           // 70
			uint16(GroundState) | (uint16(EscDispatchAction) << 8),           // 71
			uint16(GroundState) | (uint16(EscDispatchAction) << 8),           // 72
			uint16(GroundState) | (uint16(EscDispatchAction) << 8),           // 73
			uint16(GroundState) | (uint16(EscDispatchAction) << 8),           // 74
			uint16(GroundState) | (uint16(EscDispatchAction) << 8),           // 75
			uint16(GroundState) | (uint16(EscDispatchAction) << 8),           // 76
			uint16(GroundState) | (uint16(EscDispatchAction) << 8),           // 77
			uint16(GroundState) | (uint16(EscDispatchAction) << 8),           // 78
			uint16(GroundState) | (uint16(EscDispatchAction) << 8),           // 79
			uint16(GroundState) | (uint16(EscDispatchAction) << 8),           // 7A
			uint16(GroundState) | (uint16(EscDispatchAction) << 8),           // 7B
			uint16(GroundState) | (uint16(EscDispatchAction) << 8),           // 7C
			uint16(GroundState) | (uint16(EscDispatchAction) << 8),           // 7D
			uint16(GroundState) | (uint16(EscDispatchAction) << 8),           // 7E
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8),              // 7F
			0, // 80
			0, // 81
			0, // 82
			0, // 83
			0, // 84
			0, // 85
			0, // 86
			0, // 87
			0, // 88
			0, // 89
			0, // 8A
			0, // 8B
			0, // 8C
			0, // 8D
			0, // 8E
			0, // 8F
			0, // 90
			0, // 91
			0, // 92
			0, // 93
			0, // 94
			0, // 95
			0, // 96
			0, // 97
			0, // 98
			0, // 99
			0, // 9A
			0, // 9B
			0, // 9C
			0, // 9D
			0, // 9E
			0, // 9F
			0, // A0
			0, // A1
			0, // A2
			0, // A3
			0, // A4
			0, // A5
			0, // A6
			0, // A7
			0, // A8
			0, // A9
			0, // AA
			0, // AB
			0, // AC
			0, // AD
			0, // AE
			0, // AF
			0, // B0
			0, // B1
			0, // B2
			0, // B3
			0, // B4
			0, // B5
			0, // B6
			0, // B7
			0, // B8
			0, // B9
			0, // BA
			0, // BB
			0, // BC
			0, // BD
			0, // BE
			0, // BF
			0, // C0
			0, // C1
			0, // C2
			0, // C3
			0, // C4
			0, // C5
			0, // C6
			0, // C7
			0, // C8
			0, // C9
			0, // CA
			0, // CB
			0, // CC
			0, // CD
			0, // CE
			0, // CF
			0, // D0
			0, // D1
			0, // D2
			0, // D3
			0, // D4
			0, // D5
			0, // D6
			0, // D7
			0, // D8
			0, // D9
			0, // DA
			0, // DB
			0, // DC
			0, // DD
			0, // DE
			0, // DF
			0, // E0
			0, // E1
			0, // E2
			0, // E3
			0, // E4
			0, // E5
			0, // E6
			0, // E7
			0, // E8
			0, // E9
			0, // EA
			0, // EB
			0, // EC
			0, // ED
			0, // EE
			0, // EF
			0, // F0
			0, // F1
			0, // F2
			0, // F3
			0, // F4
			0, // F5
			0, // F6
			0, // F7
			0, // F8
			0, // F9
			0, // FA
			0, // FB
			0, // FC
			0, // FD
			0, // FE
			0, // FF
		},
		{ //State EscapeIntermediateState = 11
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 00
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 01
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 02
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 03
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 04
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 05
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 06
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 07
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 08
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 09
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 0A
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 0B
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 0C
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 0D
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 0E
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 0F
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 10
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 11
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 12
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 13
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 14
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 15
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 16
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 17
			0, // 18
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 19
			0, // 1A
			0, // 1B
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8),   // 1C
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8),   // 1D
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8),   // 1E
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8),   // 1F
			uint16(AnywhereState) | (uint16(CollectAction) << 8),   // 20
			uint16(AnywhereState) | (uint16(CollectAction) << 8),   // 21
			uint16(AnywhereState) | (uint16(CollectAction) << 8),   // 22
			uint16(AnywhereState) | (uint16(CollectAction) << 8),   // 23
			uint16(AnywhereState) | (uint16(CollectAction) << 8),   // 24
			uint16(AnywhereState) | (uint16(CollectAction) << 8),   // 25
			uint16(AnywhereState) | (uint16(CollectAction) << 8),   // 26
			uint16(AnywhereState) | (uint16(CollectAction) << 8),   // 27
			uint16(AnywhereState) | (uint16(CollectAction) << 8),   // 28
			uint16(AnywhereState) | (uint16(CollectAction) << 8),   // 29
			uint16(AnywhereState) | (uint16(CollectAction) << 8),   // 2A
			uint16(AnywhereState) | (uint16(CollectAction) << 8),   // 2B
			uint16(AnywhereState) | (uint16(CollectAction) << 8),   // 2C
			uint16(AnywhereState) | (uint16(CollectAction) << 8),   // 2D
			uint16(AnywhereState) | (uint16(CollectAction) << 8),   // 2E
			uint16(AnywhereState) | (uint16(CollectAction) << 8),   // 2F
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 30
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 31
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 32
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 33
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 34
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 35
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 36
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 37
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 38
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 39
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 3A
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 3B
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 3C
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 3D
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 3E
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 3F
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 40
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 41
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 42
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 43
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 44
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 45
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 46
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 47
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 48
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 49
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 4A
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 4B
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 4C
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 4D
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 4E
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 4F
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 50
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 51
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 52
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 53
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 54
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 55
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 56
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 57
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 58
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 59
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 5A
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 5B
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 5C
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 5D
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 5E
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 5F
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 60
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 61
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 62
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 63
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 64
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 65
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 66
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 67
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 68
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 69
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 6A
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 6B
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 6C
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 6D
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 6E
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 6F
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 70
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 71
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 72
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 73
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 74
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 75
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 76
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 77
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 78
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 79
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 7A
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 7B
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 7C
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 7D
			uint16(GroundState) | (uint16(EscDispatchAction) << 8), // 7E
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8),    // 7F
			0, // 80
			0, // 81
			0, // 82
			0, // 83
			0, // 84
			0, // 85
			0, // 86
			0, // 87
			0, // 88
			0, // 89
			0, // 8A
			0, // 8B
			0, // 8C
			0, // 8D
			0, // 8E
			0, // 8F
			0, // 90
			0, // 91
			0, // 92
			0, // 93
			0, // 94
			0, // 95
			0, // 96
			0, // 97
			0, // 98
			0, // 99
			0, // 9A
			0, // 9B
			0, // 9C
			0, // 9D
			0, // 9E
			0, // 9F
			0, // A0
			0, // A1
			0, // A2
			0, // A3
			0, // A4
			0, // A5
			0, // A6
			0, // A7
			0, // A8
			0, // A9
			0, // AA
			0, // AB
			0, // AC
			0, // AD
			0, // AE
			0, // AF
			0, // B0
			0, // B1
			0, // B2
			0, // B3
			0, // B4
			0, // B5
			0, // B6
			0, // B7
			0, // B8
			0, // B9
			0, // BA
			0, // BB
			0, // BC
			0, // BD
			0, // BE
			0, // BF
			0, // C0
			0, // C1
			0, // C2
			0, // C3
			0, // C4
			0, // C5
			0, // C6
			0, // C7
			0, // C8
			0, // C9
			0, // CA
			0, // CB
			0, // CC
			0, // CD
			0, // CE
			0, // CF
			0, // D0
			0, // D1
			0, // D2
			0, // D3
			0, // D4
			0, // D5
			0, // D6
			0, // D7
			0, // D8
			0, // D9
			0, // DA
			0, // DB
			0, // DC
			0, // DD
			0, // DE
			0, // DF
			0, // E0
			0, // E1
			0, // E2
			0, // E3
			0, // E4
			0, // E5
			0, // E6
			0, // E7
			0, // E8
			0, // E9
			0, // EA
			0, // EB
			0, // EC
			0, // ED
			0, // EE
			0, // EF
			0, // F0
			0, // F1
			0, // F2
			0, // F3
			0, // F4
			0, // F5
			0, // F6
			0, // F7
			0, // F8
			0, // F9
			0, // FA
			0, // FB
			0, // FC
			0, // FD
			0, // FE
			0, // FF
		},
		{ //State GroundState = 12
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 00
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 01
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 02
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 03
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 04
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 05
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 06
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 07
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 08
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 09
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 0A
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 0B
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 0C
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 0D
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 0E
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 0F
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 10
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 11
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 12
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 13
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 14
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 15
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 16
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 17
			0, // 18
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 19
			0, // 1A
			0, // 1B
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 1C
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 1D
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 1E
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 1F
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 20
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 21
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 22
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 23
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 24
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 25
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 26
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 27
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 28
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 29
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 2A
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 2B
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 2C
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 2D
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 2E
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 2F
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 30
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 31
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 32
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 33
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 34
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 35
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 36
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 37
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 38
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 39
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 3A
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 3B
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 3C
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 3D
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 3E
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 3F
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 40
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 41
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 42
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 43
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 44
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 45
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 46
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 47
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 48
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 49
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 4A
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 4B
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 4C
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 4D
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 4E
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 4F
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 50
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 51
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 52
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 53
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 54
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 55
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 56
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 57
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 58
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 59
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 5A
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 5B
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 5C
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 5D
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 5E
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 5F
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 60
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 61
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 62
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 63
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 64
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 65
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 66
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 67
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 68
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 69
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 6A
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 6B
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 6C
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 6D
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 6E
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 6F
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 70
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 71
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 72
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 73
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 74
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 75
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 76
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 77
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 78
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 79
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 7A
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 7B
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 7C
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 7D
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 7E
			uint16(AnywhereState) | (uint16(PrintAction) << 8),   // 7F
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 80
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 81
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 82
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 83
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 84
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 85
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 86
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 87
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 88
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 89
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 8A
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 8B
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 8C
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 8D
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 8E
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 8F
			0, // 90
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 91
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 92
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 93
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 94
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 95
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 96
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 97
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 98
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 99
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 9A
			0, // 9B
			uint16(AnywhereState) | (uint16(ExecuteAction) << 8), // 9C
			0, // 9D
			0, // 9E
			0, // 9F
			0, // A0
			0, // A1
			0, // A2
			0, // A3
			0, // A4
			0, // A5
			0, // A6
			0, // A7
			0, // A8
			0, // A9
			0, // AA
			0, // AB
			0, // AC
			0, // AD
			0, // AE
			0, // AF
			0, // B0
			0, // B1
			0, // B2
			0, // B3
			0, // B4
			0, // B5
			0, // B6
			0, // B7
			0, // B8
			0, // B9
			0, // BA
			0, // BB
			0, // BC
			0, // BD
			0, // BE
			0, // BF
			0, // C0
			0, // C1
			uint16(Utf8State) | (uint16(BeginUtf8Action) << 8), // C2
			uint16(Utf8State) | (uint16(BeginUtf8Action) << 8), // C3
			uint16(Utf8State) | (uint16(BeginUtf8Action) << 8), // C4
			uint16(Utf8State) | (uint16(BeginUtf8Action) << 8), // C5
			uint16(Utf8State) | (uint16(BeginUtf8Action) << 8), // C6
			uint16(Utf8State) | (uint16(BeginUtf8Action) << 8), // C7
			uint16(Utf8State) | (uint16(BeginUtf8Action) << 8), // C8
			uint16(Utf8State) | (uint16(BeginUtf8Action) << 8), // C9
			uint16(Utf8State) | (uint16(BeginUtf8Action) << 8), // CA
			uint16(Utf8State) | (uint16(BeginUtf8Action) << 8), // CB
			uint16(Utf8State) | (uint16(BeginUtf8Action) << 8), // CC
			uint16(Utf8State) | (uint16(BeginUtf8Action) << 8), // CD
			uint16(Utf8State) | (uint16(BeginUtf8Action) << 8), // CE
			uint16(Utf8State) | (uint16(BeginUtf8Action) << 8), // CF
			uint16(Utf8State) | (uint16(BeginUtf8Action) << 8), // D0
			uint16(Utf8State) | (uint16(BeginUtf8Action) << 8), // D1
			uint16(Utf8State) | (uint16(BeginUtf8Action) << 8), // D2
			uint16(Utf8State) | (uint16(BeginUtf8Action) << 8), // D3
			uint16(Utf8State) | (uint16(BeginUtf8Action) << 8), // D4
			uint16(Utf8State) | (uint16(BeginUtf8Action) << 8), // D5
			uint16(Utf8State) | (uint16(BeginUtf8Action) << 8), // D6
			uint16(Utf8State) | (uint16(BeginUtf8Action) << 8), // D7
			uint16(Utf8State) | (uint16(BeginUtf8Action) << 8), // D8
			uint16(Utf8State) | (uint16(BeginUtf8Action) << 8), // D9
			uint16(Utf8State) | (uint16(BeginUtf8Action) << 8), // DA
			uint16(Utf8State) | (uint16(BeginUtf8Action) << 8), // DB
			uint16(Utf8State) | (uint16(BeginUtf8Action) << 8), // DC
			uint16(Utf8State) | (uint16(BeginUtf8Action) << 8), // DD
			uint16(Utf8State) | (uint16(BeginUtf8Action) << 8), // DE
			uint16(Utf8State) | (uint16(BeginUtf8Action) << 8), // DF
			uint16(Utf8State) | (uint16(BeginUtf8Action) << 8), // E0
			uint16(Utf8State) | (uint16(BeginUtf8Action) << 8), // E1
			uint16(Utf8State) | (uint16(BeginUtf8Action) << 8), // E2
			uint16(Utf8State) | (uint16(BeginUtf8Action) << 8), // E3
			uint16(Utf8State) | (uint16(BeginUtf8Action) << 8), // E4
			uint16(Utf8State) | (uint16(BeginUtf8Action) << 8), // E5
			uint16(Utf8State) | (uint16(BeginUtf8Action) << 8), // E6
			uint16(Utf8State) | (uint16(BeginUtf8Action) << 8), // E7
			uint16(Utf8State) | (uint16(BeginUtf8Action) << 8), // E8
			uint16(Utf8State) | (uint16(BeginUtf8Action) << 8), // E9
			uint16(Utf8State) | (uint16(BeginUtf8Action) << 8), // EA
			uint16(Utf8State) | (uint16(BeginUtf8Action) << 8), // EB
			uint16(Utf8State) | (uint16(BeginUtf8Action) << 8), // EC
			uint16(Utf8State) | (uint16(BeginUtf8Action) << 8), // ED
			uint16(Utf8State) | (uint16(BeginUtf8Action) << 8), // EE
			uint16(Utf8State) | (uint16(BeginUtf8Action) << 8), // EF
			uint16(Utf8State) | (uint16(BeginUtf8Action) << 8), // F0
			uint16(Utf8State) | (uint16(BeginUtf8Action) << 8), // F1
			uint16(Utf8State) | (uint16(BeginUtf8Action) << 8), // F2
			uint16(Utf8State) | (uint16(BeginUtf8Action) << 8), // F3
			uint16(Utf8State) | (uint16(BeginUtf8Action) << 8), // F4
			0, // F5
			0, // F6
			0, // F7
			0, // F8
			0, // F9
			0, // FA
			0, // FB
			0, // FC
			0, // FD
			0, // FE
			0, // FF
		},
		{ //State OscStringState = 13
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 00
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 01
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 02
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 03
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 04
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 05
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 06
			uint16(GroundState) | (uint16(NoneAction) << 8),     // 07
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 08
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 09
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 0A
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 0B
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 0C
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 0D
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 0E
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 0F
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 10
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 11
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 12
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 13
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 14
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 15
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 16
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 17
			0, // 18
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 19
			0, // 1A
			0, // 1B
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 1C
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 1D
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 1E
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 1F
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 20
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 21
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 22
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 23
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 24
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 25
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 26
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 27
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 28
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 29
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 2A
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 2B
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 2C
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 2D
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 2E
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 2F
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 30
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 31
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 32
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 33
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 34
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 35
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 36
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 37
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 38
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 39
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 3A
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 3B
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 3C
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 3D
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 3E
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 3F
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 40
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 41
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 42
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 43
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 44
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 45
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 46
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 47
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 48
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 49
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 4A
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 4B
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 4C
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 4D
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 4E
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 4F
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 50
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 51
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 52
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 53
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 54
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 55
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 56
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 57
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 58
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 59
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 5A
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 5B
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 5C
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 5D
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 5E
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 5F
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 60
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 61
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 62
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 63
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 64
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 65
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 66
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 67
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 68
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 69
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 6A
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 6B
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 6C
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 6D
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 6E
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 6F
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 70
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 71
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 72
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 73
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 74
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 75
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 76
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 77
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 78
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 79
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 7A
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 7B
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 7C
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 7D
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 7E
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 7F
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 80
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 81
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 82
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 83
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 84
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 85
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 86
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 87
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 88
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 89
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 8A
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 8B
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 8C
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 8D
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 8E
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 8F
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 90
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 91
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 92
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 93
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 94
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 95
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 96
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 97
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 98
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 99
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 9A
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 9B
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 9C
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 9D
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 9E
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // 9F
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // A0
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // A1
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // A2
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // A3
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // A4
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // A5
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // A6
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // A7
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // A8
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // A9
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // AA
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // AB
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // AC
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // AD
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // AE
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // AF
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // B0
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // B1
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // B2
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // B3
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // B4
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // B5
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // B6
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // B7
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // B8
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // B9
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // BA
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // BB
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // BC
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // BD
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // BE
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // BF
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // C0
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // C1
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // C2
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // C3
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // C4
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // C5
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // C6
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // C7
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // C8
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // C9
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // CA
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // CB
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // CC
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // CD
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // CE
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // CF
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // D0
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // D1
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // D2
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // D3
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // D4
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // D5
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // D6
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // D7
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // D8
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // D9
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // DA
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // DB
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // DC
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // DD
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // DE
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // DF
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // E0
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // E1
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // E2
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // E3
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // E4
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // E5
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // E6
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // E7
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // E8
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // E9
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // EA
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // EB
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // EC
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // ED
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // EE
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // EF
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // F0
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // F1
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // F2
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // F3
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // F4
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // F5
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // F6
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // F7
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // F8
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // F9
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // FA
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // FB
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // FC
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // FD
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // FE
			uint16(AnywhereState) | (uint16(OscPutAction) << 8), // FF
		},
		{ //State SosPmApcStringState = 14
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 00
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 01
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 02
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 03
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 04
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 05
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 06
			uint16(GroundState) | (uint16(NoneAction) << 8),     // 07
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 08
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 09
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 0A
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 0B
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 0C
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 0D
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 0E
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 0F
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 10
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 11
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 12
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 13
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 14
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 15
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 16
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 17
			0, // 18
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8), // 19
			0, // 1A
			0, // 1B
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8),      // 1C
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8),      // 1D
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8),      // 1E
			uint16(AnywhereState) | (uint16(IgnoreAction) << 8),      // 1F
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 20
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 21
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 22
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 23
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 24
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 25
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 26
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 27
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 28
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 29
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 2A
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 2B
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 2C
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 2D
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 2E
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 2F
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 30
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 31
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 32
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 33
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 34
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 35
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 36
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 37
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 38
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 39
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 3A
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 3B
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 3C
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 3D
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 3E
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 3F
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 40
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 41
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 42
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 43
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 44
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 45
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 46
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 47
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 48
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 49
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 4A
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 4B
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 4C
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 4D
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 4E
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 4F
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 50
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 51
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 52
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 53
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 54
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 55
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 56
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 57
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 58
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 59
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 5A
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 5B
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 5C
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 5D
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 5E
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 5F
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 60
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 61
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 62
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 63
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 64
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 65
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 66
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 67
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 68
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 69
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 6A
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 6B
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 6C
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 6D
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 6E
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 6F
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 70
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 71
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 72
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 73
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 74
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 75
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 76
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 77
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 78
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 79
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 7A
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 7B
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 7C
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 7D
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 7E
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 7F
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 80
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 81
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 82
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 83
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 84
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 85
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 86
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 87
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 88
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 89
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 8A
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 8B
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 8C
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 8D
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 8E
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 8F
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 90
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 91
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 92
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 93
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 94
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 95
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 96
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 97
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 98
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 99
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 9A
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 9B
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 9C
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 9D
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 9E
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // 9F
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // A0
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // A1
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // A2
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // A3
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // A4
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // A5
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // A6
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // A7
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // A8
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // A9
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // AA
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // AB
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // AC
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // AD
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // AE
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // AF
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // B0
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // B1
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // B2
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // B3
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // B4
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // B5
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // B6
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // B7
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // B8
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // B9
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // BA
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // BB
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // BC
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // BD
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // BE
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // BF
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // C0
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // C1
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // C2
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // C3
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // C4
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // C5
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // C6
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // C7
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // C8
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // C9
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // CA
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // CB
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // CC
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // CD
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // CE
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // CF
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // D0
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // D1
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // D2
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // D3
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // D4
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // D5
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // D6
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // D7
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // D8
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // D9
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // DA
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // DB
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // DC
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // DD
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // DE
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // DF
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // E0
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // E1
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // E2
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // E3
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // E4
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // E5
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // E6
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // E7
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // E8
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // E9
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // EA
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // EB
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // EC
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // ED
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // EE
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // EF
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // F0
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // F1
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // F2
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // F3
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // F4
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // F5
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // F6
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // F7
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // F8
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // F9
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // FA
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // FB
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // FC
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // FD
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // FE
			uint16(AnywhereState) | (uint16(SosPmApcPutAction) << 8), // FF
		},
	}

	exitActions = []byte{
		0,            //AnywhereState
		0,            //CsiEntryState
		0,            //CsiIgnoreState
		0,            //CsiIntermediateState
		0,            //CsiParamState
		0,            //DcsEntryState
		0,            //DcsIgnoreState
		0,            //DcsIntermediateState
		0,            //DcsParamState
		UnhookAction, //DcsPassthroughState
		0,            //EscapeState
		0,            //EscapeIntermediateState
		0,            //GroundState
		OscEndAction, //OscStringState
		0,            //SosPmApcStringState
		0,            //Utf8State
	}
)
