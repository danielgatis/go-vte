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

	NoneAction        Action = 0
	ClearAction       Action = 1
	CollectAction     Action = 2
	CsiDispatchAction Action = 3
	EscDispatchAction Action = 4
	ExecuteAction     Action = 5
	HookAction        Action = 6
	IgnoreAction      Action = 7
	OscEndAction      Action = 8
	OscPutAction      Action = 9
	OscStartAction    Action = 10
	ParamAction       Action = 11
	PrintAction       Action = 12
	PutAction         Action = 13
	UnhookAction      Action = 14
	BeginUtf8Action   Action = 15
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

	stateTable = [][]byte{
		{ //State AnywhereState = 0
			0,                                  // 00
			0,                                  // 01
			0,                                  // 02
			0,                                  // 03
			0,                                  // 04
			0,                                  // 05
			0,                                  // 06
			0,                                  // 07
			0,                                  // 08
			0,                                  // 09
			0,                                  // 0A
			0,                                  // 0B
			0,                                  // 0C
			0,                                  // 0D
			0,                                  // 0E
			0,                                  // 0F
			0,                                  // 10
			0,                                  // 11
			0,                                  // 12
			0,                                  // 13
			0,                                  // 14
			0,                                  // 15
			0,                                  // 16
			0,                                  // 17
			GroundState | (ExecuteAction << 4), // 18
			0,                                  // 19
			GroundState | (ExecuteAction << 4), // 1A
			EscapeState | (NoneAction << 4),    // 1B
			0,                                  // 1C
			0,                                  // 1D
			0,                                  // 1E
			0,                                  // 1F
			0,                                  // 20
			0,                                  // 21
			0,                                  // 22
			0,                                  // 23
			0,                                  // 24
			0,                                  // 25
			0,                                  // 26
			0,                                  // 27
			0,                                  // 28
			0,                                  // 29
			0,                                  // 2A
			0,                                  // 2B
			0,                                  // 2C
			0,                                  // 2D
			0,                                  // 2E
			0,                                  // 2F
			0,                                  // 30
			0,                                  // 31
			0,                                  // 32
			0,                                  // 33
			0,                                  // 34
			0,                                  // 35
			0,                                  // 36
			0,                                  // 37
			0,                                  // 38
			0,                                  // 39
			0,                                  // 3A
			0,                                  // 3B
			0,                                  // 3C
			0,                                  // 3D
			0,                                  // 3E
			0,                                  // 3F
			0,                                  // 40
			0,                                  // 41
			0,                                  // 42
			0,                                  // 43
			0,                                  // 44
			0,                                  // 45
			0,                                  // 46
			0,                                  // 47
			0,                                  // 48
			0,                                  // 49
			0,                                  // 4A
			0,                                  // 4B
			0,                                  // 4C
			0,                                  // 4D
			0,                                  // 4E
			0,                                  // 4F
			0,                                  // 50
			0,                                  // 51
			0,                                  // 52
			0,                                  // 53
			0,                                  // 54
			0,                                  // 55
			0,                                  // 56
			0,                                  // 57
			0,                                  // 58
			0,                                  // 59
			0,                                  // 5A
			0,                                  // 5B
			0,                                  // 5C
			0,                                  // 5D
			0,                                  // 5E
			0,                                  // 5F
			0,                                  // 60
			0,                                  // 61
			0,                                  // 62
			0,                                  // 63
			0,                                  // 64
			0,                                  // 65
			0,                                  // 66
			0,                                  // 67
			0,                                  // 68
			0,                                  // 69
			0,                                  // 6A
			0,                                  // 6B
			0,                                  // 6C
			0,                                  // 6D
			0,                                  // 6E
			0,                                  // 6F
			0,                                  // 70
			0,                                  // 71
			0,                                  // 72
			0,                                  // 73
			0,                                  // 74
			0,                                  // 75
			0,                                  // 76
			0,                                  // 77
			0,                                  // 78
			0,                                  // 79
			0,                                  // 7A
			0,                                  // 7B
			0,                                  // 7C
			0,                                  // 7D
			0,                                  // 7E
			0,                                  // 7F
			0,                                  // 80
			0,                                  // 81
			0,                                  // 82
			0,                                  // 83
			0,                                  // 84
			0,                                  // 85
			0,                                  // 86
			0,                                  // 87
			0,                                  // 88
			0,                                  // 89
			0,                                  // 8A
			0,                                  // 8B
			0,                                  // 8C
			0,                                  // 8D
			0,                                  // 8E
			0,                                  // 8F
			0,                                  // 90
			0,                                  // 91
			0,                                  // 92
			0,                                  // 93
			0,                                  // 94
			0,                                  // 95
			0,                                  // 96
			0,                                  // 97
			0,                                  // 98
			0,                                  // 99
			0,                                  // 9A
			0,                                  // 9B
			0,                                  // 9C
			0,                                  // 9D
			0,                                  // 9E
			0,                                  // 9F
			0,                                  // A0
			0,                                  // A1
			0,                                  // A2
			0,                                  // A3
			0,                                  // A4
			0,                                  // A5
			0,                                  // A6
			0,                                  // A7
			0,                                  // A8
			0,                                  // A9
			0,                                  // AA
			0,                                  // AB
			0,                                  // AC
			0,                                  // AD
			0,                                  // AE
			0,                                  // AF
			0,                                  // B0
			0,                                  // B1
			0,                                  // B2
			0,                                  // B3
			0,                                  // B4
			0,                                  // B5
			0,                                  // B6
			0,                                  // B7
			0,                                  // B8
			0,                                  // B9
			0,                                  // BA
			0,                                  // BB
			0,                                  // BC
			0,                                  // BD
			0,                                  // BE
			0,                                  // BF
			0,                                  // C0
			0,                                  // C1
			0,                                  // C2
			0,                                  // C3
			0,                                  // C4
			0,                                  // C5
			0,                                  // C6
			0,                                  // C7
			0,                                  // C8
			0,                                  // C9
			0,                                  // CA
			0,                                  // CB
			0,                                  // CC
			0,                                  // CD
			0,                                  // CE
			0,                                  // CF
			0,                                  // D0
			0,                                  // D1
			0,                                  // D2
			0,                                  // D3
			0,                                  // D4
			0,                                  // D5
			0,                                  // D6
			0,                                  // D7
			0,                                  // D8
			0,                                  // D9
			0,                                  // DA
			0,                                  // DB
			0,                                  // DC
			0,                                  // DD
			0,                                  // DE
			0,                                  // DF
			0,                                  // E0
			0,                                  // E1
			0,                                  // E2
			0,                                  // E3
			0,                                  // E4
			0,                                  // E5
			0,                                  // E6
			0,                                  // E7
			0,                                  // E8
			0,                                  // E9
			0,                                  // EA
			0,                                  // EB
			0,                                  // EC
			0,                                  // ED
			0,                                  // EE
			0,                                  // EF
			0,                                  // F0
			0,                                  // F1
			0,                                  // F2
			0,                                  // F3
			0,                                  // F4
			0,                                  // F5
			0,                                  // F6
			0,                                  // F7
			0,                                  // F8
			0,                                  // F9
			0,                                  // FA
			0,                                  // FB
			0,                                  // FC
			0,                                  // FD
			0,                                  // FE
			0,                                  // FF
		},
		{ //State CsiEntryState = 1
			AnywhereState | (ExecuteAction << 4),        // 00
			AnywhereState | (ExecuteAction << 4),        // 01
			AnywhereState | (ExecuteAction << 4),        // 02
			AnywhereState | (ExecuteAction << 4),        // 03
			AnywhereState | (ExecuteAction << 4),        // 04
			AnywhereState | (ExecuteAction << 4),        // 05
			AnywhereState | (ExecuteAction << 4),        // 06
			AnywhereState | (ExecuteAction << 4),        // 07
			AnywhereState | (ExecuteAction << 4),        // 08
			AnywhereState | (ExecuteAction << 4),        // 09
			AnywhereState | (ExecuteAction << 4),        // 0A
			AnywhereState | (ExecuteAction << 4),        // 0B
			AnywhereState | (ExecuteAction << 4),        // 0C
			AnywhereState | (ExecuteAction << 4),        // 0D
			AnywhereState | (ExecuteAction << 4),        // 0E
			AnywhereState | (ExecuteAction << 4),        // 0F
			AnywhereState | (ExecuteAction << 4),        // 10
			AnywhereState | (ExecuteAction << 4),        // 11
			AnywhereState | (ExecuteAction << 4),        // 12
			AnywhereState | (ExecuteAction << 4),        // 13
			AnywhereState | (ExecuteAction << 4),        // 14
			AnywhereState | (ExecuteAction << 4),        // 15
			AnywhereState | (ExecuteAction << 4),        // 16
			AnywhereState | (ExecuteAction << 4),        // 17
			0,                                           // 18
			AnywhereState | (ExecuteAction << 4),        // 19
			0,                                           // 1A
			0,                                           // 1B
			AnywhereState | (ExecuteAction << 4),        // 1C
			AnywhereState | (ExecuteAction << 4),        // 1D
			AnywhereState | (ExecuteAction << 4),        // 1E
			AnywhereState | (ExecuteAction << 4),        // 1F
			CsiIntermediateState | (CollectAction << 4), // 20
			CsiIntermediateState | (CollectAction << 4), // 21
			CsiIntermediateState | (CollectAction << 4), // 22
			CsiIntermediateState | (CollectAction << 4), // 23
			CsiIntermediateState | (CollectAction << 4), // 24
			CsiIntermediateState | (CollectAction << 4), // 25
			CsiIntermediateState | (CollectAction << 4), // 26
			CsiIntermediateState | (CollectAction << 4), // 27
			CsiIntermediateState | (CollectAction << 4), // 28
			CsiIntermediateState | (CollectAction << 4), // 29
			CsiIntermediateState | (CollectAction << 4), // 2A
			CsiIntermediateState | (CollectAction << 4), // 2B
			CsiIntermediateState | (CollectAction << 4), // 2C
			CsiIntermediateState | (CollectAction << 4), // 2D
			CsiIntermediateState | (CollectAction << 4), // 2E
			CsiIntermediateState | (CollectAction << 4), // 2F
			CsiParamState | (ParamAction << 4),          // 30
			CsiParamState | (ParamAction << 4),          // 31
			CsiParamState | (ParamAction << 4),          // 32
			CsiParamState | (ParamAction << 4),          // 33
			CsiParamState | (ParamAction << 4),          // 34
			CsiParamState | (ParamAction << 4),          // 35
			CsiParamState | (ParamAction << 4),          // 36
			CsiParamState | (ParamAction << 4),          // 37
			CsiParamState | (ParamAction << 4),          // 38
			CsiParamState | (ParamAction << 4),          // 39
			CsiParamState | (ParamAction << 4),          // 3A
			CsiParamState | (ParamAction << 4),          // 3B
			CsiParamState | (CollectAction << 4),        // 3C
			CsiParamState | (CollectAction << 4),        // 3D
			CsiParamState | (CollectAction << 4),        // 3E
			CsiParamState | (CollectAction << 4),        // 3F
			GroundState | (CsiDispatchAction << 4),      // 40
			GroundState | (CsiDispatchAction << 4),      // 41
			GroundState | (CsiDispatchAction << 4),      // 42
			GroundState | (CsiDispatchAction << 4),      // 43
			GroundState | (CsiDispatchAction << 4),      // 44
			GroundState | (CsiDispatchAction << 4),      // 45
			GroundState | (CsiDispatchAction << 4),      // 46
			GroundState | (CsiDispatchAction << 4),      // 47
			GroundState | (CsiDispatchAction << 4),      // 48
			GroundState | (CsiDispatchAction << 4),      // 49
			GroundState | (CsiDispatchAction << 4),      // 4A
			GroundState | (CsiDispatchAction << 4),      // 4B
			GroundState | (CsiDispatchAction << 4),      // 4C
			GroundState | (CsiDispatchAction << 4),      // 4D
			GroundState | (CsiDispatchAction << 4),      // 4E
			GroundState | (CsiDispatchAction << 4),      // 4F
			GroundState | (CsiDispatchAction << 4),      // 50
			GroundState | (CsiDispatchAction << 4),      // 51
			GroundState | (CsiDispatchAction << 4),      // 52
			GroundState | (CsiDispatchAction << 4),      // 53
			GroundState | (CsiDispatchAction << 4),      // 54
			GroundState | (CsiDispatchAction << 4),      // 55
			GroundState | (CsiDispatchAction << 4),      // 56
			GroundState | (CsiDispatchAction << 4),      // 57
			GroundState | (CsiDispatchAction << 4),      // 58
			GroundState | (CsiDispatchAction << 4),      // 59
			GroundState | (CsiDispatchAction << 4),      // 5A
			GroundState | (CsiDispatchAction << 4),      // 5B
			GroundState | (CsiDispatchAction << 4),      // 5C
			GroundState | (CsiDispatchAction << 4),      // 5D
			GroundState | (CsiDispatchAction << 4),      // 5E
			GroundState | (CsiDispatchAction << 4),      // 5F
			GroundState | (CsiDispatchAction << 4),      // 60
			GroundState | (CsiDispatchAction << 4),      // 61
			GroundState | (CsiDispatchAction << 4),      // 62
			GroundState | (CsiDispatchAction << 4),      // 63
			GroundState | (CsiDispatchAction << 4),      // 64
			GroundState | (CsiDispatchAction << 4),      // 65
			GroundState | (CsiDispatchAction << 4),      // 66
			GroundState | (CsiDispatchAction << 4),      // 67
			GroundState | (CsiDispatchAction << 4),      // 68
			GroundState | (CsiDispatchAction << 4),      // 69
			GroundState | (CsiDispatchAction << 4),      // 6A
			GroundState | (CsiDispatchAction << 4),      // 6B
			GroundState | (CsiDispatchAction << 4),      // 6C
			GroundState | (CsiDispatchAction << 4),      // 6D
			GroundState | (CsiDispatchAction << 4),      // 6E
			GroundState | (CsiDispatchAction << 4),      // 6F
			GroundState | (CsiDispatchAction << 4),      // 70
			GroundState | (CsiDispatchAction << 4),      // 71
			GroundState | (CsiDispatchAction << 4),      // 72
			GroundState | (CsiDispatchAction << 4),      // 73
			GroundState | (CsiDispatchAction << 4),      // 74
			GroundState | (CsiDispatchAction << 4),      // 75
			GroundState | (CsiDispatchAction << 4),      // 76
			GroundState | (CsiDispatchAction << 4),      // 77
			GroundState | (CsiDispatchAction << 4),      // 78
			GroundState | (CsiDispatchAction << 4),      // 79
			GroundState | (CsiDispatchAction << 4),      // 7A
			GroundState | (CsiDispatchAction << 4),      // 7B
			GroundState | (CsiDispatchAction << 4),      // 7C
			GroundState | (CsiDispatchAction << 4),      // 7D
			GroundState | (CsiDispatchAction << 4),      // 7E
			AnywhereState | (IgnoreAction << 4),         // 7F
			0,                                           // 80
			0,                                           // 81
			0,                                           // 82
			0,                                           // 83
			0,                                           // 84
			0,                                           // 85
			0,                                           // 86
			0,                                           // 87
			0,                                           // 88
			0,                                           // 89
			0,                                           // 8A
			0,                                           // 8B
			0,                                           // 8C
			0,                                           // 8D
			0,                                           // 8E
			0,                                           // 8F
			0,                                           // 90
			0,                                           // 91
			0,                                           // 92
			0,                                           // 93
			0,                                           // 94
			0,                                           // 95
			0,                                           // 96
			0,                                           // 97
			0,                                           // 98
			0,                                           // 99
			0,                                           // 9A
			0,                                           // 9B
			0,                                           // 9C
			0,                                           // 9D
			0,                                           // 9E
			0,                                           // 9F
			0,                                           // A0
			0,                                           // A1
			0,                                           // A2
			0,                                           // A3
			0,                                           // A4
			0,                                           // A5
			0,                                           // A6
			0,                                           // A7
			0,                                           // A8
			0,                                           // A9
			0,                                           // AA
			0,                                           // AB
			0,                                           // AC
			0,                                           // AD
			0,                                           // AE
			0,                                           // AF
			0,                                           // B0
			0,                                           // B1
			0,                                           // B2
			0,                                           // B3
			0,                                           // B4
			0,                                           // B5
			0,                                           // B6
			0,                                           // B7
			0,                                           // B8
			0,                                           // B9
			0,                                           // BA
			0,                                           // BB
			0,                                           // BC
			0,                                           // BD
			0,                                           // BE
			0,                                           // BF
			0,                                           // C0
			0,                                           // C1
			0,                                           // C2
			0,                                           // C3
			0,                                           // C4
			0,                                           // C5
			0,                                           // C6
			0,                                           // C7
			0,                                           // C8
			0,                                           // C9
			0,                                           // CA
			0,                                           // CB
			0,                                           // CC
			0,                                           // CD
			0,                                           // CE
			0,                                           // CF
			0,                                           // D0
			0,                                           // D1
			0,                                           // D2
			0,                                           // D3
			0,                                           // D4
			0,                                           // D5
			0,                                           // D6
			0,                                           // D7
			0,                                           // D8
			0,                                           // D9
			0,                                           // DA
			0,                                           // DB
			0,                                           // DC
			0,                                           // DD
			0,                                           // DE
			0,                                           // DF
			0,                                           // E0
			0,                                           // E1
			0,                                           // E2
			0,                                           // E3
			0,                                           // E4
			0,                                           // E5
			0,                                           // E6
			0,                                           // E7
			0,                                           // E8
			0,                                           // E9
			0,                                           // EA
			0,                                           // EB
			0,                                           // EC
			0,                                           // ED
			0,                                           // EE
			0,                                           // EF
			0,                                           // F0
			0,                                           // F1
			0,                                           // F2
			0,                                           // F3
			0,                                           // F4
			0,                                           // F5
			0,                                           // F6
			0,                                           // F7
			0,                                           // F8
			0,                                           // F9
			0,                                           // FA
			0,                                           // FB
			0,                                           // FC
			0,                                           // FD
			0,                                           // FE
			0,                                           // FF
		},
		{ //State CsiIgnoreState = 2
			AnywhereState | (ExecuteAction << 4), // 00
			AnywhereState | (ExecuteAction << 4), // 01
			AnywhereState | (ExecuteAction << 4), // 02
			AnywhereState | (ExecuteAction << 4), // 03
			AnywhereState | (ExecuteAction << 4), // 04
			AnywhereState | (ExecuteAction << 4), // 05
			AnywhereState | (ExecuteAction << 4), // 06
			AnywhereState | (ExecuteAction << 4), // 07
			AnywhereState | (ExecuteAction << 4), // 08
			AnywhereState | (ExecuteAction << 4), // 09
			AnywhereState | (ExecuteAction << 4), // 0A
			AnywhereState | (ExecuteAction << 4), // 0B
			AnywhereState | (ExecuteAction << 4), // 0C
			AnywhereState | (ExecuteAction << 4), // 0D
			AnywhereState | (ExecuteAction << 4), // 0E
			AnywhereState | (ExecuteAction << 4), // 0F
			AnywhereState | (ExecuteAction << 4), // 10
			AnywhereState | (ExecuteAction << 4), // 11
			AnywhereState | (ExecuteAction << 4), // 12
			AnywhereState | (ExecuteAction << 4), // 13
			AnywhereState | (ExecuteAction << 4), // 14
			AnywhereState | (ExecuteAction << 4), // 15
			AnywhereState | (ExecuteAction << 4), // 16
			AnywhereState | (ExecuteAction << 4), // 17
			0,                                    // 18
			AnywhereState | (ExecuteAction << 4), // 19
			0,                                    // 1A
			0,                                    // 1B
			AnywhereState | (ExecuteAction << 4), // 1C
			AnywhereState | (ExecuteAction << 4), // 1D
			AnywhereState | (ExecuteAction << 4), // 1E
			AnywhereState | (ExecuteAction << 4), // 1F
			AnywhereState | (IgnoreAction << 4),  // 20
			AnywhereState | (IgnoreAction << 4),  // 21
			AnywhereState | (IgnoreAction << 4),  // 22
			AnywhereState | (IgnoreAction << 4),  // 23
			AnywhereState | (IgnoreAction << 4),  // 24
			AnywhereState | (IgnoreAction << 4),  // 25
			AnywhereState | (IgnoreAction << 4),  // 26
			AnywhereState | (IgnoreAction << 4),  // 27
			AnywhereState | (IgnoreAction << 4),  // 28
			AnywhereState | (IgnoreAction << 4),  // 29
			AnywhereState | (IgnoreAction << 4),  // 2A
			AnywhereState | (IgnoreAction << 4),  // 2B
			AnywhereState | (IgnoreAction << 4),  // 2C
			AnywhereState | (IgnoreAction << 4),  // 2D
			AnywhereState | (IgnoreAction << 4),  // 2E
			AnywhereState | (IgnoreAction << 4),  // 2F
			AnywhereState | (IgnoreAction << 4),  // 30
			AnywhereState | (IgnoreAction << 4),  // 31
			AnywhereState | (IgnoreAction << 4),  // 32
			AnywhereState | (IgnoreAction << 4),  // 33
			AnywhereState | (IgnoreAction << 4),  // 34
			AnywhereState | (IgnoreAction << 4),  // 35
			AnywhereState | (IgnoreAction << 4),  // 36
			AnywhereState | (IgnoreAction << 4),  // 37
			AnywhereState | (IgnoreAction << 4),  // 38
			AnywhereState | (IgnoreAction << 4),  // 39
			AnywhereState | (IgnoreAction << 4),  // 3A
			AnywhereState | (IgnoreAction << 4),  // 3B
			AnywhereState | (IgnoreAction << 4),  // 3C
			AnywhereState | (IgnoreAction << 4),  // 3D
			AnywhereState | (IgnoreAction << 4),  // 3E
			AnywhereState | (IgnoreAction << 4),  // 3F
			GroundState | (NoneAction << 4),      // 40
			GroundState | (NoneAction << 4),      // 41
			GroundState | (NoneAction << 4),      // 42
			GroundState | (NoneAction << 4),      // 43
			GroundState | (NoneAction << 4),      // 44
			GroundState | (NoneAction << 4),      // 45
			GroundState | (NoneAction << 4),      // 46
			GroundState | (NoneAction << 4),      // 47
			GroundState | (NoneAction << 4),      // 48
			GroundState | (NoneAction << 4),      // 49
			GroundState | (NoneAction << 4),      // 4A
			GroundState | (NoneAction << 4),      // 4B
			GroundState | (NoneAction << 4),      // 4C
			GroundState | (NoneAction << 4),      // 4D
			GroundState | (NoneAction << 4),      // 4E
			GroundState | (NoneAction << 4),      // 4F
			GroundState | (NoneAction << 4),      // 50
			GroundState | (NoneAction << 4),      // 51
			GroundState | (NoneAction << 4),      // 52
			GroundState | (NoneAction << 4),      // 53
			GroundState | (NoneAction << 4),      // 54
			GroundState | (NoneAction << 4),      // 55
			GroundState | (NoneAction << 4),      // 56
			GroundState | (NoneAction << 4),      // 57
			GroundState | (NoneAction << 4),      // 58
			GroundState | (NoneAction << 4),      // 59
			GroundState | (NoneAction << 4),      // 5A
			GroundState | (NoneAction << 4),      // 5B
			GroundState | (NoneAction << 4),      // 5C
			GroundState | (NoneAction << 4),      // 5D
			GroundState | (NoneAction << 4),      // 5E
			GroundState | (NoneAction << 4),      // 5F
			GroundState | (NoneAction << 4),      // 60
			GroundState | (NoneAction << 4),      // 61
			GroundState | (NoneAction << 4),      // 62
			GroundState | (NoneAction << 4),      // 63
			GroundState | (NoneAction << 4),      // 64
			GroundState | (NoneAction << 4),      // 65
			GroundState | (NoneAction << 4),      // 66
			GroundState | (NoneAction << 4),      // 67
			GroundState | (NoneAction << 4),      // 68
			GroundState | (NoneAction << 4),      // 69
			GroundState | (NoneAction << 4),      // 6A
			GroundState | (NoneAction << 4),      // 6B
			GroundState | (NoneAction << 4),      // 6C
			GroundState | (NoneAction << 4),      // 6D
			GroundState | (NoneAction << 4),      // 6E
			GroundState | (NoneAction << 4),      // 6F
			GroundState | (NoneAction << 4),      // 70
			GroundState | (NoneAction << 4),      // 71
			GroundState | (NoneAction << 4),      // 72
			GroundState | (NoneAction << 4),      // 73
			GroundState | (NoneAction << 4),      // 74
			GroundState | (NoneAction << 4),      // 75
			GroundState | (NoneAction << 4),      // 76
			GroundState | (NoneAction << 4),      // 77
			GroundState | (NoneAction << 4),      // 78
			GroundState | (NoneAction << 4),      // 79
			GroundState | (NoneAction << 4),      // 7A
			GroundState | (NoneAction << 4),      // 7B
			GroundState | (NoneAction << 4),      // 7C
			GroundState | (NoneAction << 4),      // 7D
			GroundState | (NoneAction << 4),      // 7E
			AnywhereState | (IgnoreAction << 4),  // 7F
			0,                                    // 80
			0,                                    // 81
			0,                                    // 82
			0,                                    // 83
			0,                                    // 84
			0,                                    // 85
			0,                                    // 86
			0,                                    // 87
			0,                                    // 88
			0,                                    // 89
			0,                                    // 8A
			0,                                    // 8B
			0,                                    // 8C
			0,                                    // 8D
			0,                                    // 8E
			0,                                    // 8F
			0,                                    // 90
			0,                                    // 91
			0,                                    // 92
			0,                                    // 93
			0,                                    // 94
			0,                                    // 95
			0,                                    // 96
			0,                                    // 97
			0,                                    // 98
			0,                                    // 99
			0,                                    // 9A
			0,                                    // 9B
			0,                                    // 9C
			0,                                    // 9D
			0,                                    // 9E
			0,                                    // 9F
			0,                                    // A0
			0,                                    // A1
			0,                                    // A2
			0,                                    // A3
			0,                                    // A4
			0,                                    // A5
			0,                                    // A6
			0,                                    // A7
			0,                                    // A8
			0,                                    // A9
			0,                                    // AA
			0,                                    // AB
			0,                                    // AC
			0,                                    // AD
			0,                                    // AE
			0,                                    // AF
			0,                                    // B0
			0,                                    // B1
			0,                                    // B2
			0,                                    // B3
			0,                                    // B4
			0,                                    // B5
			0,                                    // B6
			0,                                    // B7
			0,                                    // B8
			0,                                    // B9
			0,                                    // BA
			0,                                    // BB
			0,                                    // BC
			0,                                    // BD
			0,                                    // BE
			0,                                    // BF
			0,                                    // C0
			0,                                    // C1
			0,                                    // C2
			0,                                    // C3
			0,                                    // C4
			0,                                    // C5
			0,                                    // C6
			0,                                    // C7
			0,                                    // C8
			0,                                    // C9
			0,                                    // CA
			0,                                    // CB
			0,                                    // CC
			0,                                    // CD
			0,                                    // CE
			0,                                    // CF
			0,                                    // D0
			0,                                    // D1
			0,                                    // D2
			0,                                    // D3
			0,                                    // D4
			0,                                    // D5
			0,                                    // D6
			0,                                    // D7
			0,                                    // D8
			0,                                    // D9
			0,                                    // DA
			0,                                    // DB
			0,                                    // DC
			0,                                    // DD
			0,                                    // DE
			0,                                    // DF
			0,                                    // E0
			0,                                    // E1
			0,                                    // E2
			0,                                    // E3
			0,                                    // E4
			0,                                    // E5
			0,                                    // E6
			0,                                    // E7
			0,                                    // E8
			0,                                    // E9
			0,                                    // EA
			0,                                    // EB
			0,                                    // EC
			0,                                    // ED
			0,                                    // EE
			0,                                    // EF
			0,                                    // F0
			0,                                    // F1
			0,                                    // F2
			0,                                    // F3
			0,                                    // F4
			0,                                    // F5
			0,                                    // F6
			0,                                    // F7
			0,                                    // F8
			0,                                    // F9
			0,                                    // FA
			0,                                    // FB
			0,                                    // FC
			0,                                    // FD
			0,                                    // FE
			0,                                    // FF
		},
		{ //State CsiIntermediateState = 3
			AnywhereState | (ExecuteAction << 4),   // 00
			AnywhereState | (ExecuteAction << 4),   // 01
			AnywhereState | (ExecuteAction << 4),   // 02
			AnywhereState | (ExecuteAction << 4),   // 03
			AnywhereState | (ExecuteAction << 4),   // 04
			AnywhereState | (ExecuteAction << 4),   // 05
			AnywhereState | (ExecuteAction << 4),   // 06
			AnywhereState | (ExecuteAction << 4),   // 07
			AnywhereState | (ExecuteAction << 4),   // 08
			AnywhereState | (ExecuteAction << 4),   // 09
			AnywhereState | (ExecuteAction << 4),   // 0A
			AnywhereState | (ExecuteAction << 4),   // 0B
			AnywhereState | (ExecuteAction << 4),   // 0C
			AnywhereState | (ExecuteAction << 4),   // 0D
			AnywhereState | (ExecuteAction << 4),   // 0E
			AnywhereState | (ExecuteAction << 4),   // 0F
			AnywhereState | (ExecuteAction << 4),   // 10
			AnywhereState | (ExecuteAction << 4),   // 11
			AnywhereState | (ExecuteAction << 4),   // 12
			AnywhereState | (ExecuteAction << 4),   // 13
			AnywhereState | (ExecuteAction << 4),   // 14
			AnywhereState | (ExecuteAction << 4),   // 15
			AnywhereState | (ExecuteAction << 4),   // 16
			AnywhereState | (ExecuteAction << 4),   // 17
			0,                                      // 18
			AnywhereState | (ExecuteAction << 4),   // 19
			0,                                      // 1A
			0,                                      // 1B
			AnywhereState | (ExecuteAction << 4),   // 1C
			AnywhereState | (ExecuteAction << 4),   // 1D
			AnywhereState | (ExecuteAction << 4),   // 1E
			AnywhereState | (ExecuteAction << 4),   // 1F
			AnywhereState | (CollectAction << 4),   // 20
			AnywhereState | (CollectAction << 4),   // 21
			AnywhereState | (CollectAction << 4),   // 22
			AnywhereState | (CollectAction << 4),   // 23
			AnywhereState | (CollectAction << 4),   // 24
			AnywhereState | (CollectAction << 4),   // 25
			AnywhereState | (CollectAction << 4),   // 26
			AnywhereState | (CollectAction << 4),   // 27
			AnywhereState | (CollectAction << 4),   // 28
			AnywhereState | (CollectAction << 4),   // 29
			AnywhereState | (CollectAction << 4),   // 2A
			AnywhereState | (CollectAction << 4),   // 2B
			AnywhereState | (CollectAction << 4),   // 2C
			AnywhereState | (CollectAction << 4),   // 2D
			AnywhereState | (CollectAction << 4),   // 2E
			AnywhereState | (CollectAction << 4),   // 2F
			CsiIgnoreState | (NoneAction << 4),     // 30
			CsiIgnoreState | (NoneAction << 4),     // 31
			CsiIgnoreState | (NoneAction << 4),     // 32
			CsiIgnoreState | (NoneAction << 4),     // 33
			CsiIgnoreState | (NoneAction << 4),     // 34
			CsiIgnoreState | (NoneAction << 4),     // 35
			CsiIgnoreState | (NoneAction << 4),     // 36
			CsiIgnoreState | (NoneAction << 4),     // 37
			CsiIgnoreState | (NoneAction << 4),     // 38
			CsiIgnoreState | (NoneAction << 4),     // 39
			CsiIgnoreState | (NoneAction << 4),     // 3A
			CsiIgnoreState | (NoneAction << 4),     // 3B
			CsiIgnoreState | (NoneAction << 4),     // 3C
			CsiIgnoreState | (NoneAction << 4),     // 3D
			CsiIgnoreState | (NoneAction << 4),     // 3E
			CsiIgnoreState | (NoneAction << 4),     // 3F
			GroundState | (CsiDispatchAction << 4), // 40
			GroundState | (CsiDispatchAction << 4), // 41
			GroundState | (CsiDispatchAction << 4), // 42
			GroundState | (CsiDispatchAction << 4), // 43
			GroundState | (CsiDispatchAction << 4), // 44
			GroundState | (CsiDispatchAction << 4), // 45
			GroundState | (CsiDispatchAction << 4), // 46
			GroundState | (CsiDispatchAction << 4), // 47
			GroundState | (CsiDispatchAction << 4), // 48
			GroundState | (CsiDispatchAction << 4), // 49
			GroundState | (CsiDispatchAction << 4), // 4A
			GroundState | (CsiDispatchAction << 4), // 4B
			GroundState | (CsiDispatchAction << 4), // 4C
			GroundState | (CsiDispatchAction << 4), // 4D
			GroundState | (CsiDispatchAction << 4), // 4E
			GroundState | (CsiDispatchAction << 4), // 4F
			GroundState | (CsiDispatchAction << 4), // 50
			GroundState | (CsiDispatchAction << 4), // 51
			GroundState | (CsiDispatchAction << 4), // 52
			GroundState | (CsiDispatchAction << 4), // 53
			GroundState | (CsiDispatchAction << 4), // 54
			GroundState | (CsiDispatchAction << 4), // 55
			GroundState | (CsiDispatchAction << 4), // 56
			GroundState | (CsiDispatchAction << 4), // 57
			GroundState | (CsiDispatchAction << 4), // 58
			GroundState | (CsiDispatchAction << 4), // 59
			GroundState | (CsiDispatchAction << 4), // 5A
			GroundState | (CsiDispatchAction << 4), // 5B
			GroundState | (CsiDispatchAction << 4), // 5C
			GroundState | (CsiDispatchAction << 4), // 5D
			GroundState | (CsiDispatchAction << 4), // 5E
			GroundState | (CsiDispatchAction << 4), // 5F
			GroundState | (CsiDispatchAction << 4), // 60
			GroundState | (CsiDispatchAction << 4), // 61
			GroundState | (CsiDispatchAction << 4), // 62
			GroundState | (CsiDispatchAction << 4), // 63
			GroundState | (CsiDispatchAction << 4), // 64
			GroundState | (CsiDispatchAction << 4), // 65
			GroundState | (CsiDispatchAction << 4), // 66
			GroundState | (CsiDispatchAction << 4), // 67
			GroundState | (CsiDispatchAction << 4), // 68
			GroundState | (CsiDispatchAction << 4), // 69
			GroundState | (CsiDispatchAction << 4), // 6A
			GroundState | (CsiDispatchAction << 4), // 6B
			GroundState | (CsiDispatchAction << 4), // 6C
			GroundState | (CsiDispatchAction << 4), // 6D
			GroundState | (CsiDispatchAction << 4), // 6E
			GroundState | (CsiDispatchAction << 4), // 6F
			GroundState | (CsiDispatchAction << 4), // 70
			GroundState | (CsiDispatchAction << 4), // 71
			GroundState | (CsiDispatchAction << 4), // 72
			GroundState | (CsiDispatchAction << 4), // 73
			GroundState | (CsiDispatchAction << 4), // 74
			GroundState | (CsiDispatchAction << 4), // 75
			GroundState | (CsiDispatchAction << 4), // 76
			GroundState | (CsiDispatchAction << 4), // 77
			GroundState | (CsiDispatchAction << 4), // 78
			GroundState | (CsiDispatchAction << 4), // 79
			GroundState | (CsiDispatchAction << 4), // 7A
			GroundState | (CsiDispatchAction << 4), // 7B
			GroundState | (CsiDispatchAction << 4), // 7C
			GroundState | (CsiDispatchAction << 4), // 7D
			GroundState | (CsiDispatchAction << 4), // 7E
			AnywhereState | (IgnoreAction << 4),    // 7F
			0,                                      // 80
			0,                                      // 81
			0,                                      // 82
			0,                                      // 83
			0,                                      // 84
			0,                                      // 85
			0,                                      // 86
			0,                                      // 87
			0,                                      // 88
			0,                                      // 89
			0,                                      // 8A
			0,                                      // 8B
			0,                                      // 8C
			0,                                      // 8D
			0,                                      // 8E
			0,                                      // 8F
			0,                                      // 90
			0,                                      // 91
			0,                                      // 92
			0,                                      // 93
			0,                                      // 94
			0,                                      // 95
			0,                                      // 96
			0,                                      // 97
			0,                                      // 98
			0,                                      // 99
			0,                                      // 9A
			0,                                      // 9B
			0,                                      // 9C
			0,                                      // 9D
			0,                                      // 9E
			0,                                      // 9F
			0,                                      // A0
			0,                                      // A1
			0,                                      // A2
			0,                                      // A3
			0,                                      // A4
			0,                                      // A5
			0,                                      // A6
			0,                                      // A7
			0,                                      // A8
			0,                                      // A9
			0,                                      // AA
			0,                                      // AB
			0,                                      // AC
			0,                                      // AD
			0,                                      // AE
			0,                                      // AF
			0,                                      // B0
			0,                                      // B1
			0,                                      // B2
			0,                                      // B3
			0,                                      // B4
			0,                                      // B5
			0,                                      // B6
			0,                                      // B7
			0,                                      // B8
			0,                                      // B9
			0,                                      // BA
			0,                                      // BB
			0,                                      // BC
			0,                                      // BD
			0,                                      // BE
			0,                                      // BF
			0,                                      // C0
			0,                                      // C1
			0,                                      // C2
			0,                                      // C3
			0,                                      // C4
			0,                                      // C5
			0,                                      // C6
			0,                                      // C7
			0,                                      // C8
			0,                                      // C9
			0,                                      // CA
			0,                                      // CB
			0,                                      // CC
			0,                                      // CD
			0,                                      // CE
			0,                                      // CF
			0,                                      // D0
			0,                                      // D1
			0,                                      // D2
			0,                                      // D3
			0,                                      // D4
			0,                                      // D5
			0,                                      // D6
			0,                                      // D7
			0,                                      // D8
			0,                                      // D9
			0,                                      // DA
			0,                                      // DB
			0,                                      // DC
			0,                                      // DD
			0,                                      // DE
			0,                                      // DF
			0,                                      // E0
			0,                                      // E1
			0,                                      // E2
			0,                                      // E3
			0,                                      // E4
			0,                                      // E5
			0,                                      // E6
			0,                                      // E7
			0,                                      // E8
			0,                                      // E9
			0,                                      // EA
			0,                                      // EB
			0,                                      // EC
			0,                                      // ED
			0,                                      // EE
			0,                                      // EF
			0,                                      // F0
			0,                                      // F1
			0,                                      // F2
			0,                                      // F3
			0,                                      // F4
			0,                                      // F5
			0,                                      // F6
			0,                                      // F7
			0,                                      // F8
			0,                                      // F9
			0,                                      // FA
			0,                                      // FB
			0,                                      // FC
			0,                                      // FD
			0,                                      // FE
			0,                                      // FF
		},
		{ //State CsiParamState = 4
			AnywhereState | (ExecuteAction << 4),        // 00
			AnywhereState | (ExecuteAction << 4),        // 01
			AnywhereState | (ExecuteAction << 4),        // 02
			AnywhereState | (ExecuteAction << 4),        // 03
			AnywhereState | (ExecuteAction << 4),        // 04
			AnywhereState | (ExecuteAction << 4),        // 05
			AnywhereState | (ExecuteAction << 4),        // 06
			AnywhereState | (ExecuteAction << 4),        // 07
			AnywhereState | (ExecuteAction << 4),        // 08
			AnywhereState | (ExecuteAction << 4),        // 09
			AnywhereState | (ExecuteAction << 4),        // 0A
			AnywhereState | (ExecuteAction << 4),        // 0B
			AnywhereState | (ExecuteAction << 4),        // 0C
			AnywhereState | (ExecuteAction << 4),        // 0D
			AnywhereState | (ExecuteAction << 4),        // 0E
			AnywhereState | (ExecuteAction << 4),        // 0F
			AnywhereState | (ExecuteAction << 4),        // 10
			AnywhereState | (ExecuteAction << 4),        // 11
			AnywhereState | (ExecuteAction << 4),        // 12
			AnywhereState | (ExecuteAction << 4),        // 13
			AnywhereState | (ExecuteAction << 4),        // 14
			AnywhereState | (ExecuteAction << 4),        // 15
			AnywhereState | (ExecuteAction << 4),        // 16
			AnywhereState | (ExecuteAction << 4),        // 17
			0,                                           // 18
			AnywhereState | (ExecuteAction << 4),        // 19
			0,                                           // 1A
			0,                                           // 1B
			AnywhereState | (ExecuteAction << 4),        // 1C
			AnywhereState | (ExecuteAction << 4),        // 1D
			AnywhereState | (ExecuteAction << 4),        // 1E
			AnywhereState | (ExecuteAction << 4),        // 1F
			CsiIntermediateState | (CollectAction << 4), // 20
			CsiIntermediateState | (CollectAction << 4), // 21
			CsiIntermediateState | (CollectAction << 4), // 22
			CsiIntermediateState | (CollectAction << 4), // 23
			CsiIntermediateState | (CollectAction << 4), // 24
			CsiIntermediateState | (CollectAction << 4), // 25
			CsiIntermediateState | (CollectAction << 4), // 26
			CsiIntermediateState | (CollectAction << 4), // 27
			CsiIntermediateState | (CollectAction << 4), // 28
			CsiIntermediateState | (CollectAction << 4), // 29
			CsiIntermediateState | (CollectAction << 4), // 2A
			CsiIntermediateState | (CollectAction << 4), // 2B
			CsiIntermediateState | (CollectAction << 4), // 2C
			CsiIntermediateState | (CollectAction << 4), // 2D
			CsiIntermediateState | (CollectAction << 4), // 2E
			CsiIntermediateState | (CollectAction << 4), // 2F
			AnywhereState | (ParamAction << 4),          // 30
			AnywhereState | (ParamAction << 4),          // 31
			AnywhereState | (ParamAction << 4),          // 32
			AnywhereState | (ParamAction << 4),          // 33
			AnywhereState | (ParamAction << 4),          // 34
			AnywhereState | (ParamAction << 4),          // 35
			AnywhereState | (ParamAction << 4),          // 36
			AnywhereState | (ParamAction << 4),          // 37
			AnywhereState | (ParamAction << 4),          // 38
			AnywhereState | (ParamAction << 4),          // 39
			AnywhereState | (ParamAction << 4),          // 3A
			AnywhereState | (ParamAction << 4),          // 3B
			CsiIgnoreState | (NoneAction << 4),          // 3C
			CsiIgnoreState | (NoneAction << 4),          // 3D
			CsiIgnoreState | (NoneAction << 4),          // 3E
			CsiIgnoreState | (NoneAction << 4),          // 3F
			GroundState | (CsiDispatchAction << 4),      // 40
			GroundState | (CsiDispatchAction << 4),      // 41
			GroundState | (CsiDispatchAction << 4),      // 42
			GroundState | (CsiDispatchAction << 4),      // 43
			GroundState | (CsiDispatchAction << 4),      // 44
			GroundState | (CsiDispatchAction << 4),      // 45
			GroundState | (CsiDispatchAction << 4),      // 46
			GroundState | (CsiDispatchAction << 4),      // 47
			GroundState | (CsiDispatchAction << 4),      // 48
			GroundState | (CsiDispatchAction << 4),      // 49
			GroundState | (CsiDispatchAction << 4),      // 4A
			GroundState | (CsiDispatchAction << 4),      // 4B
			GroundState | (CsiDispatchAction << 4),      // 4C
			GroundState | (CsiDispatchAction << 4),      // 4D
			GroundState | (CsiDispatchAction << 4),      // 4E
			GroundState | (CsiDispatchAction << 4),      // 4F
			GroundState | (CsiDispatchAction << 4),      // 50
			GroundState | (CsiDispatchAction << 4),      // 51
			GroundState | (CsiDispatchAction << 4),      // 52
			GroundState | (CsiDispatchAction << 4),      // 53
			GroundState | (CsiDispatchAction << 4),      // 54
			GroundState | (CsiDispatchAction << 4),      // 55
			GroundState | (CsiDispatchAction << 4),      // 56
			GroundState | (CsiDispatchAction << 4),      // 57
			GroundState | (CsiDispatchAction << 4),      // 58
			GroundState | (CsiDispatchAction << 4),      // 59
			GroundState | (CsiDispatchAction << 4),      // 5A
			GroundState | (CsiDispatchAction << 4),      // 5B
			GroundState | (CsiDispatchAction << 4),      // 5C
			GroundState | (CsiDispatchAction << 4),      // 5D
			GroundState | (CsiDispatchAction << 4),      // 5E
			GroundState | (CsiDispatchAction << 4),      // 5F
			GroundState | (CsiDispatchAction << 4),      // 60
			GroundState | (CsiDispatchAction << 4),      // 61
			GroundState | (CsiDispatchAction << 4),      // 62
			GroundState | (CsiDispatchAction << 4),      // 63
			GroundState | (CsiDispatchAction << 4),      // 64
			GroundState | (CsiDispatchAction << 4),      // 65
			GroundState | (CsiDispatchAction << 4),      // 66
			GroundState | (CsiDispatchAction << 4),      // 67
			GroundState | (CsiDispatchAction << 4),      // 68
			GroundState | (CsiDispatchAction << 4),      // 69
			GroundState | (CsiDispatchAction << 4),      // 6A
			GroundState | (CsiDispatchAction << 4),      // 6B
			GroundState | (CsiDispatchAction << 4),      // 6C
			GroundState | (CsiDispatchAction << 4),      // 6D
			GroundState | (CsiDispatchAction << 4),      // 6E
			GroundState | (CsiDispatchAction << 4),      // 6F
			GroundState | (CsiDispatchAction << 4),      // 70
			GroundState | (CsiDispatchAction << 4),      // 71
			GroundState | (CsiDispatchAction << 4),      // 72
			GroundState | (CsiDispatchAction << 4),      // 73
			GroundState | (CsiDispatchAction << 4),      // 74
			GroundState | (CsiDispatchAction << 4),      // 75
			GroundState | (CsiDispatchAction << 4),      // 76
			GroundState | (CsiDispatchAction << 4),      // 77
			GroundState | (CsiDispatchAction << 4),      // 78
			GroundState | (CsiDispatchAction << 4),      // 79
			GroundState | (CsiDispatchAction << 4),      // 7A
			GroundState | (CsiDispatchAction << 4),      // 7B
			GroundState | (CsiDispatchAction << 4),      // 7C
			GroundState | (CsiDispatchAction << 4),      // 7D
			GroundState | (CsiDispatchAction << 4),      // 7E
			AnywhereState | (IgnoreAction << 4),         // 7F
			0,                                           // 80
			0,                                           // 81
			0,                                           // 82
			0,                                           // 83
			0,                                           // 84
			0,                                           // 85
			0,                                           // 86
			0,                                           // 87
			0,                                           // 88
			0,                                           // 89
			0,                                           // 8A
			0,                                           // 8B
			0,                                           // 8C
			0,                                           // 8D
			0,                                           // 8E
			0,                                           // 8F
			0,                                           // 90
			0,                                           // 91
			0,                                           // 92
			0,                                           // 93
			0,                                           // 94
			0,                                           // 95
			0,                                           // 96
			0,                                           // 97
			0,                                           // 98
			0,                                           // 99
			0,                                           // 9A
			0,                                           // 9B
			0,                                           // 9C
			0,                                           // 9D
			0,                                           // 9E
			0,                                           // 9F
			0,                                           // A0
			0,                                           // A1
			0,                                           // A2
			0,                                           // A3
			0,                                           // A4
			0,                                           // A5
			0,                                           // A6
			0,                                           // A7
			0,                                           // A8
			0,                                           // A9
			0,                                           // AA
			0,                                           // AB
			0,                                           // AC
			0,                                           // AD
			0,                                           // AE
			0,                                           // AF
			0,                                           // B0
			0,                                           // B1
			0,                                           // B2
			0,                                           // B3
			0,                                           // B4
			0,                                           // B5
			0,                                           // B6
			0,                                           // B7
			0,                                           // B8
			0,                                           // B9
			0,                                           // BA
			0,                                           // BB
			0,                                           // BC
			0,                                           // BD
			0,                                           // BE
			0,                                           // BF
			0,                                           // C0
			0,                                           // C1
			0,                                           // C2
			0,                                           // C3
			0,                                           // C4
			0,                                           // C5
			0,                                           // C6
			0,                                           // C7
			0,                                           // C8
			0,                                           // C9
			0,                                           // CA
			0,                                           // CB
			0,                                           // CC
			0,                                           // CD
			0,                                           // CE
			0,                                           // CF
			0,                                           // D0
			0,                                           // D1
			0,                                           // D2
			0,                                           // D3
			0,                                           // D4
			0,                                           // D5
			0,                                           // D6
			0,                                           // D7
			0,                                           // D8
			0,                                           // D9
			0,                                           // DA
			0,                                           // DB
			0,                                           // DC
			0,                                           // DD
			0,                                           // DE
			0,                                           // DF
			0,                                           // E0
			0,                                           // E1
			0,                                           // E2
			0,                                           // E3
			0,                                           // E4
			0,                                           // E5
			0,                                           // E6
			0,                                           // E7
			0,                                           // E8
			0,                                           // E9
			0,                                           // EA
			0,                                           // EB
			0,                                           // EC
			0,                                           // ED
			0,                                           // EE
			0,                                           // EF
			0,                                           // F0
			0,                                           // F1
			0,                                           // F2
			0,                                           // F3
			0,                                           // F4
			0,                                           // F5
			0,                                           // F6
			0,                                           // F7
			0,                                           // F8
			0,                                           // F9
			0,                                           // FA
			0,                                           // FB
			0,                                           // FC
			0,                                           // FD
			0,                                           // FE
			0,                                           // FF
		},
		{ //State DcsEntryState = 5
			AnywhereState | (IgnoreAction << 4),         // 00
			AnywhereState | (IgnoreAction << 4),         // 01
			AnywhereState | (IgnoreAction << 4),         // 02
			AnywhereState | (IgnoreAction << 4),         // 03
			AnywhereState | (IgnoreAction << 4),         // 04
			AnywhereState | (IgnoreAction << 4),         // 05
			AnywhereState | (IgnoreAction << 4),         // 06
			AnywhereState | (IgnoreAction << 4),         // 07
			AnywhereState | (IgnoreAction << 4),         // 08
			AnywhereState | (IgnoreAction << 4),         // 09
			AnywhereState | (IgnoreAction << 4),         // 0A
			AnywhereState | (IgnoreAction << 4),         // 0B
			AnywhereState | (IgnoreAction << 4),         // 0C
			AnywhereState | (IgnoreAction << 4),         // 0D
			AnywhereState | (IgnoreAction << 4),         // 0E
			AnywhereState | (IgnoreAction << 4),         // 0F
			AnywhereState | (IgnoreAction << 4),         // 10
			AnywhereState | (IgnoreAction << 4),         // 11
			AnywhereState | (IgnoreAction << 4),         // 12
			AnywhereState | (IgnoreAction << 4),         // 13
			AnywhereState | (IgnoreAction << 4),         // 14
			AnywhereState | (IgnoreAction << 4),         // 15
			AnywhereState | (IgnoreAction << 4),         // 16
			AnywhereState | (IgnoreAction << 4),         // 17
			0,                                           // 18
			AnywhereState | (IgnoreAction << 4),         // 19
			0,                                           // 1A
			0,                                           // 1B
			AnywhereState | (IgnoreAction << 4),         // 1C
			AnywhereState | (IgnoreAction << 4),         // 1D
			AnywhereState | (IgnoreAction << 4),         // 1E
			AnywhereState | (IgnoreAction << 4),         // 1F
			DcsIntermediateState | (CollectAction << 4), // 20
			DcsIntermediateState | (CollectAction << 4), // 21
			DcsIntermediateState | (CollectAction << 4), // 22
			DcsIntermediateState | (CollectAction << 4), // 23
			DcsIntermediateState | (CollectAction << 4), // 24
			DcsIntermediateState | (CollectAction << 4), // 25
			DcsIntermediateState | (CollectAction << 4), // 26
			DcsIntermediateState | (CollectAction << 4), // 27
			DcsIntermediateState | (CollectAction << 4), // 28
			DcsIntermediateState | (CollectAction << 4), // 29
			DcsIntermediateState | (CollectAction << 4), // 2A
			DcsIntermediateState | (CollectAction << 4), // 2B
			DcsIntermediateState | (CollectAction << 4), // 2C
			DcsIntermediateState | (CollectAction << 4), // 2D
			DcsIntermediateState | (CollectAction << 4), // 2E
			DcsIntermediateState | (CollectAction << 4), // 2F
			DcsParamState | (ParamAction << 4),          // 30
			DcsParamState | (ParamAction << 4),          // 31
			DcsParamState | (ParamAction << 4),          // 32
			DcsParamState | (ParamAction << 4),          // 33
			DcsParamState | (ParamAction << 4),          // 34
			DcsParamState | (ParamAction << 4),          // 35
			DcsParamState | (ParamAction << 4),          // 36
			DcsParamState | (ParamAction << 4),          // 37
			DcsParamState | (ParamAction << 4),          // 38
			DcsParamState | (ParamAction << 4),          // 39
			DcsParamState | (ParamAction << 4),          // 3A
			DcsParamState | (ParamAction << 4),          // 3B
			DcsParamState | (CollectAction << 4),        // 3C
			DcsParamState | (CollectAction << 4),        // 3D
			DcsParamState | (CollectAction << 4),        // 3E
			DcsParamState | (CollectAction << 4),        // 3F
			DcsPassthroughState | (NoneAction << 4),     // 40
			DcsPassthroughState | (NoneAction << 4),     // 41
			DcsPassthroughState | (NoneAction << 4),     // 42
			DcsPassthroughState | (NoneAction << 4),     // 43
			DcsPassthroughState | (NoneAction << 4),     // 44
			DcsPassthroughState | (NoneAction << 4),     // 45
			DcsPassthroughState | (NoneAction << 4),     // 46
			DcsPassthroughState | (NoneAction << 4),     // 47
			DcsPassthroughState | (NoneAction << 4),     // 48
			DcsPassthroughState | (NoneAction << 4),     // 49
			DcsPassthroughState | (NoneAction << 4),     // 4A
			DcsPassthroughState | (NoneAction << 4),     // 4B
			DcsPassthroughState | (NoneAction << 4),     // 4C
			DcsPassthroughState | (NoneAction << 4),     // 4D
			DcsPassthroughState | (NoneAction << 4),     // 4E
			DcsPassthroughState | (NoneAction << 4),     // 4F
			DcsPassthroughState | (NoneAction << 4),     // 50
			DcsPassthroughState | (NoneAction << 4),     // 51
			DcsPassthroughState | (NoneAction << 4),     // 52
			DcsPassthroughState | (NoneAction << 4),     // 53
			DcsPassthroughState | (NoneAction << 4),     // 54
			DcsPassthroughState | (NoneAction << 4),     // 55
			DcsPassthroughState | (NoneAction << 4),     // 56
			DcsPassthroughState | (NoneAction << 4),     // 57
			DcsPassthroughState | (NoneAction << 4),     // 58
			DcsPassthroughState | (NoneAction << 4),     // 59
			DcsPassthroughState | (NoneAction << 4),     // 5A
			DcsPassthroughState | (NoneAction << 4),     // 5B
			DcsPassthroughState | (NoneAction << 4),     // 5C
			DcsPassthroughState | (NoneAction << 4),     // 5D
			DcsPassthroughState | (NoneAction << 4),     // 5E
			DcsPassthroughState | (NoneAction << 4),     // 5F
			DcsPassthroughState | (NoneAction << 4),     // 60
			DcsPassthroughState | (NoneAction << 4),     // 61
			DcsPassthroughState | (NoneAction << 4),     // 62
			DcsPassthroughState | (NoneAction << 4),     // 63
			DcsPassthroughState | (NoneAction << 4),     // 64
			DcsPassthroughState | (NoneAction << 4),     // 65
			DcsPassthroughState | (NoneAction << 4),     // 66
			DcsPassthroughState | (NoneAction << 4),     // 67
			DcsPassthroughState | (NoneAction << 4),     // 68
			DcsPassthroughState | (NoneAction << 4),     // 69
			DcsPassthroughState | (NoneAction << 4),     // 6A
			DcsPassthroughState | (NoneAction << 4),     // 6B
			DcsPassthroughState | (NoneAction << 4),     // 6C
			DcsPassthroughState | (NoneAction << 4),     // 6D
			DcsPassthroughState | (NoneAction << 4),     // 6E
			DcsPassthroughState | (NoneAction << 4),     // 6F
			DcsPassthroughState | (NoneAction << 4),     // 70
			DcsPassthroughState | (NoneAction << 4),     // 71
			DcsPassthroughState | (NoneAction << 4),     // 72
			DcsPassthroughState | (NoneAction << 4),     // 73
			DcsPassthroughState | (NoneAction << 4),     // 74
			DcsPassthroughState | (NoneAction << 4),     // 75
			DcsPassthroughState | (NoneAction << 4),     // 76
			DcsPassthroughState | (NoneAction << 4),     // 77
			DcsPassthroughState | (NoneAction << 4),     // 78
			DcsPassthroughState | (NoneAction << 4),     // 79
			DcsPassthroughState | (NoneAction << 4),     // 7A
			DcsPassthroughState | (NoneAction << 4),     // 7B
			DcsPassthroughState | (NoneAction << 4),     // 7C
			DcsPassthroughState | (NoneAction << 4),     // 7D
			DcsPassthroughState | (NoneAction << 4),     // 7E
			AnywhereState | (IgnoreAction << 4),         // 7F
			0,                                           // 80
			0,                                           // 81
			0,                                           // 82
			0,                                           // 83
			0,                                           // 84
			0,                                           // 85
			0,                                           // 86
			0,                                           // 87
			0,                                           // 88
			0,                                           // 89
			0,                                           // 8A
			0,                                           // 8B
			0,                                           // 8C
			0,                                           // 8D
			0,                                           // 8E
			0,                                           // 8F
			0,                                           // 90
			0,                                           // 91
			0,                                           // 92
			0,                                           // 93
			0,                                           // 94
			0,                                           // 95
			0,                                           // 96
			0,                                           // 97
			0,                                           // 98
			0,                                           // 99
			0,                                           // 9A
			0,                                           // 9B
			0,                                           // 9C
			0,                                           // 9D
			0,                                           // 9E
			0,                                           // 9F
			0,                                           // A0
			0,                                           // A1
			0,                                           // A2
			0,                                           // A3
			0,                                           // A4
			0,                                           // A5
			0,                                           // A6
			0,                                           // A7
			0,                                           // A8
			0,                                           // A9
			0,                                           // AA
			0,                                           // AB
			0,                                           // AC
			0,                                           // AD
			0,                                           // AE
			0,                                           // AF
			0,                                           // B0
			0,                                           // B1
			0,                                           // B2
			0,                                           // B3
			0,                                           // B4
			0,                                           // B5
			0,                                           // B6
			0,                                           // B7
			0,                                           // B8
			0,                                           // B9
			0,                                           // BA
			0,                                           // BB
			0,                                           // BC
			0,                                           // BD
			0,                                           // BE
			0,                                           // BF
			0,                                           // C0
			0,                                           // C1
			0,                                           // C2
			0,                                           // C3
			0,                                           // C4
			0,                                           // C5
			0,                                           // C6
			0,                                           // C7
			0,                                           // C8
			0,                                           // C9
			0,                                           // CA
			0,                                           // CB
			0,                                           // CC
			0,                                           // CD
			0,                                           // CE
			0,                                           // CF
			0,                                           // D0
			0,                                           // D1
			0,                                           // D2
			0,                                           // D3
			0,                                           // D4
			0,                                           // D5
			0,                                           // D6
			0,                                           // D7
			0,                                           // D8
			0,                                           // D9
			0,                                           // DA
			0,                                           // DB
			0,                                           // DC
			0,                                           // DD
			0,                                           // DE
			0,                                           // DF
			0,                                           // E0
			0,                                           // E1
			0,                                           // E2
			0,                                           // E3
			0,                                           // E4
			0,                                           // E5
			0,                                           // E6
			0,                                           // E7
			0,                                           // E8
			0,                                           // E9
			0,                                           // EA
			0,                                           // EB
			0,                                           // EC
			0,                                           // ED
			0,                                           // EE
			0,                                           // EF
			0,                                           // F0
			0,                                           // F1
			0,                                           // F2
			0,                                           // F3
			0,                                           // F4
			0,                                           // F5
			0,                                           // F6
			0,                                           // F7
			0,                                           // F8
			0,                                           // F9
			0,                                           // FA
			0,                                           // FB
			0,                                           // FC
			0,                                           // FD
			0,                                           // FE
			0,                                           // FF
		},
		{ //State DcsIgnoreState = 6
			AnywhereState | (IgnoreAction << 4), // 00
			AnywhereState | (IgnoreAction << 4), // 01
			AnywhereState | (IgnoreAction << 4), // 02
			AnywhereState | (IgnoreAction << 4), // 03
			AnywhereState | (IgnoreAction << 4), // 04
			AnywhereState | (IgnoreAction << 4), // 05
			AnywhereState | (IgnoreAction << 4), // 06
			AnywhereState | (IgnoreAction << 4), // 07
			AnywhereState | (IgnoreAction << 4), // 08
			AnywhereState | (IgnoreAction << 4), // 09
			AnywhereState | (IgnoreAction << 4), // 0A
			AnywhereState | (IgnoreAction << 4), // 0B
			AnywhereState | (IgnoreAction << 4), // 0C
			AnywhereState | (IgnoreAction << 4), // 0D
			AnywhereState | (IgnoreAction << 4), // 0E
			AnywhereState | (IgnoreAction << 4), // 0F
			AnywhereState | (IgnoreAction << 4), // 10
			AnywhereState | (IgnoreAction << 4), // 11
			AnywhereState | (IgnoreAction << 4), // 12
			AnywhereState | (IgnoreAction << 4), // 13
			AnywhereState | (IgnoreAction << 4), // 14
			AnywhereState | (IgnoreAction << 4), // 15
			AnywhereState | (IgnoreAction << 4), // 16
			AnywhereState | (IgnoreAction << 4), // 17
			0,                                   // 18
			AnywhereState | (IgnoreAction << 4), // 19
			0,                                   // 1A
			0,                                   // 1B
			AnywhereState | (IgnoreAction << 4), // 1C
			AnywhereState | (IgnoreAction << 4), // 1D
			AnywhereState | (IgnoreAction << 4), // 1E
			AnywhereState | (IgnoreAction << 4), // 1F
			AnywhereState | (IgnoreAction << 4), // 20
			AnywhereState | (IgnoreAction << 4), // 21
			AnywhereState | (IgnoreAction << 4), // 22
			AnywhereState | (IgnoreAction << 4), // 23
			AnywhereState | (IgnoreAction << 4), // 24
			AnywhereState | (IgnoreAction << 4), // 25
			AnywhereState | (IgnoreAction << 4), // 26
			AnywhereState | (IgnoreAction << 4), // 27
			AnywhereState | (IgnoreAction << 4), // 28
			AnywhereState | (IgnoreAction << 4), // 29
			AnywhereState | (IgnoreAction << 4), // 2A
			AnywhereState | (IgnoreAction << 4), // 2B
			AnywhereState | (IgnoreAction << 4), // 2C
			AnywhereState | (IgnoreAction << 4), // 2D
			AnywhereState | (IgnoreAction << 4), // 2E
			AnywhereState | (IgnoreAction << 4), // 2F
			AnywhereState | (IgnoreAction << 4), // 30
			AnywhereState | (IgnoreAction << 4), // 31
			AnywhereState | (IgnoreAction << 4), // 32
			AnywhereState | (IgnoreAction << 4), // 33
			AnywhereState | (IgnoreAction << 4), // 34
			AnywhereState | (IgnoreAction << 4), // 35
			AnywhereState | (IgnoreAction << 4), // 36
			AnywhereState | (IgnoreAction << 4), // 37
			AnywhereState | (IgnoreAction << 4), // 38
			AnywhereState | (IgnoreAction << 4), // 39
			AnywhereState | (IgnoreAction << 4), // 3A
			AnywhereState | (IgnoreAction << 4), // 3B
			AnywhereState | (IgnoreAction << 4), // 3C
			AnywhereState | (IgnoreAction << 4), // 3D
			AnywhereState | (IgnoreAction << 4), // 3E
			AnywhereState | (IgnoreAction << 4), // 3F
			AnywhereState | (IgnoreAction << 4), // 40
			AnywhereState | (IgnoreAction << 4), // 41
			AnywhereState | (IgnoreAction << 4), // 42
			AnywhereState | (IgnoreAction << 4), // 43
			AnywhereState | (IgnoreAction << 4), // 44
			AnywhereState | (IgnoreAction << 4), // 45
			AnywhereState | (IgnoreAction << 4), // 46
			AnywhereState | (IgnoreAction << 4), // 47
			AnywhereState | (IgnoreAction << 4), // 48
			AnywhereState | (IgnoreAction << 4), // 49
			AnywhereState | (IgnoreAction << 4), // 4A
			AnywhereState | (IgnoreAction << 4), // 4B
			AnywhereState | (IgnoreAction << 4), // 4C
			AnywhereState | (IgnoreAction << 4), // 4D
			AnywhereState | (IgnoreAction << 4), // 4E
			AnywhereState | (IgnoreAction << 4), // 4F
			AnywhereState | (IgnoreAction << 4), // 50
			AnywhereState | (IgnoreAction << 4), // 51
			AnywhereState | (IgnoreAction << 4), // 52
			AnywhereState | (IgnoreAction << 4), // 53
			AnywhereState | (IgnoreAction << 4), // 54
			AnywhereState | (IgnoreAction << 4), // 55
			AnywhereState | (IgnoreAction << 4), // 56
			AnywhereState | (IgnoreAction << 4), // 57
			AnywhereState | (IgnoreAction << 4), // 58
			AnywhereState | (IgnoreAction << 4), // 59
			AnywhereState | (IgnoreAction << 4), // 5A
			AnywhereState | (IgnoreAction << 4), // 5B
			AnywhereState | (IgnoreAction << 4), // 5C
			AnywhereState | (IgnoreAction << 4), // 5D
			AnywhereState | (IgnoreAction << 4), // 5E
			AnywhereState | (IgnoreAction << 4), // 5F
			AnywhereState | (IgnoreAction << 4), // 60
			AnywhereState | (IgnoreAction << 4), // 61
			AnywhereState | (IgnoreAction << 4), // 62
			AnywhereState | (IgnoreAction << 4), // 63
			AnywhereState | (IgnoreAction << 4), // 64
			AnywhereState | (IgnoreAction << 4), // 65
			AnywhereState | (IgnoreAction << 4), // 66
			AnywhereState | (IgnoreAction << 4), // 67
			AnywhereState | (IgnoreAction << 4), // 68
			AnywhereState | (IgnoreAction << 4), // 69
			AnywhereState | (IgnoreAction << 4), // 6A
			AnywhereState | (IgnoreAction << 4), // 6B
			AnywhereState | (IgnoreAction << 4), // 6C
			AnywhereState | (IgnoreAction << 4), // 6D
			AnywhereState | (IgnoreAction << 4), // 6E
			AnywhereState | (IgnoreAction << 4), // 6F
			AnywhereState | (IgnoreAction << 4), // 70
			AnywhereState | (IgnoreAction << 4), // 71
			AnywhereState | (IgnoreAction << 4), // 72
			AnywhereState | (IgnoreAction << 4), // 73
			AnywhereState | (IgnoreAction << 4), // 74
			AnywhereState | (IgnoreAction << 4), // 75
			AnywhereState | (IgnoreAction << 4), // 76
			AnywhereState | (IgnoreAction << 4), // 77
			AnywhereState | (IgnoreAction << 4), // 78
			AnywhereState | (IgnoreAction << 4), // 79
			AnywhereState | (IgnoreAction << 4), // 7A
			AnywhereState | (IgnoreAction << 4), // 7B
			AnywhereState | (IgnoreAction << 4), // 7C
			AnywhereState | (IgnoreAction << 4), // 7D
			AnywhereState | (IgnoreAction << 4), // 7E
			AnywhereState | (IgnoreAction << 4), // 7F
			0,                                   // 80
			0,                                   // 81
			0,                                   // 82
			0,                                   // 83
			0,                                   // 84
			0,                                   // 85
			0,                                   // 86
			0,                                   // 87
			0,                                   // 88
			0,                                   // 89
			0,                                   // 8A
			0,                                   // 8B
			0,                                   // 8C
			0,                                   // 8D
			0,                                   // 8E
			0,                                   // 8F
			0,                                   // 90
			0,                                   // 91
			0,                                   // 92
			0,                                   // 93
			0,                                   // 94
			0,                                   // 95
			0,                                   // 96
			0,                                   // 97
			0,                                   // 98
			0,                                   // 99
			0,                                   // 9A
			0,                                   // 9B
			GroundState | (NoneAction << 4),     // 9C
			0,                                   // 9D
			0,                                   // 9E
			0,                                   // 9F
			0,                                   // A0
			0,                                   // A1
			0,                                   // A2
			0,                                   // A3
			0,                                   // A4
			0,                                   // A5
			0,                                   // A6
			0,                                   // A7
			0,                                   // A8
			0,                                   // A9
			0,                                   // AA
			0,                                   // AB
			0,                                   // AC
			0,                                   // AD
			0,                                   // AE
			0,                                   // AF
			0,                                   // B0
			0,                                   // B1
			0,                                   // B2
			0,                                   // B3
			0,                                   // B4
			0,                                   // B5
			0,                                   // B6
			0,                                   // B7
			0,                                   // B8
			0,                                   // B9
			0,                                   // BA
			0,                                   // BB
			0,                                   // BC
			0,                                   // BD
			0,                                   // BE
			0,                                   // BF
			0,                                   // C0
			0,                                   // C1
			0,                                   // C2
			0,                                   // C3
			0,                                   // C4
			0,                                   // C5
			0,                                   // C6
			0,                                   // C7
			0,                                   // C8
			0,                                   // C9
			0,                                   // CA
			0,                                   // CB
			0,                                   // CC
			0,                                   // CD
			0,                                   // CE
			0,                                   // CF
			0,                                   // D0
			0,                                   // D1
			0,                                   // D2
			0,                                   // D3
			0,                                   // D4
			0,                                   // D5
			0,                                   // D6
			0,                                   // D7
			0,                                   // D8
			0,                                   // D9
			0,                                   // DA
			0,                                   // DB
			0,                                   // DC
			0,                                   // DD
			0,                                   // DE
			0,                                   // DF
			0,                                   // E0
			0,                                   // E1
			0,                                   // E2
			0,                                   // E3
			0,                                   // E4
			0,                                   // E5
			0,                                   // E6
			0,                                   // E7
			0,                                   // E8
			0,                                   // E9
			0,                                   // EA
			0,                                   // EB
			0,                                   // EC
			0,                                   // ED
			0,                                   // EE
			0,                                   // EF
			0,                                   // F0
			0,                                   // F1
			0,                                   // F2
			0,                                   // F3
			0,                                   // F4
			0,                                   // F5
			0,                                   // F6
			0,                                   // F7
			0,                                   // F8
			0,                                   // F9
			0,                                   // FA
			0,                                   // FB
			0,                                   // FC
			0,                                   // FD
			0,                                   // FE
			0,                                   // FF
		},
		{ //State DcsIntermediateState = 7
			AnywhereState | (IgnoreAction << 4),     // 00
			AnywhereState | (IgnoreAction << 4),     // 01
			AnywhereState | (IgnoreAction << 4),     // 02
			AnywhereState | (IgnoreAction << 4),     // 03
			AnywhereState | (IgnoreAction << 4),     // 04
			AnywhereState | (IgnoreAction << 4),     // 05
			AnywhereState | (IgnoreAction << 4),     // 06
			AnywhereState | (IgnoreAction << 4),     // 07
			AnywhereState | (IgnoreAction << 4),     // 08
			AnywhereState | (IgnoreAction << 4),     // 09
			AnywhereState | (IgnoreAction << 4),     // 0A
			AnywhereState | (IgnoreAction << 4),     // 0B
			AnywhereState | (IgnoreAction << 4),     // 0C
			AnywhereState | (IgnoreAction << 4),     // 0D
			AnywhereState | (IgnoreAction << 4),     // 0E
			AnywhereState | (IgnoreAction << 4),     // 0F
			AnywhereState | (IgnoreAction << 4),     // 10
			AnywhereState | (IgnoreAction << 4),     // 11
			AnywhereState | (IgnoreAction << 4),     // 12
			AnywhereState | (IgnoreAction << 4),     // 13
			AnywhereState | (IgnoreAction << 4),     // 14
			AnywhereState | (IgnoreAction << 4),     // 15
			AnywhereState | (IgnoreAction << 4),     // 16
			AnywhereState | (IgnoreAction << 4),     // 17
			0,                                       // 18
			AnywhereState | (IgnoreAction << 4),     // 19
			0,                                       // 1A
			0,                                       // 1B
			AnywhereState | (IgnoreAction << 4),     // 1C
			AnywhereState | (IgnoreAction << 4),     // 1D
			AnywhereState | (IgnoreAction << 4),     // 1E
			AnywhereState | (IgnoreAction << 4),     // 1F
			AnywhereState | (CollectAction << 4),    // 20
			AnywhereState | (CollectAction << 4),    // 21
			AnywhereState | (CollectAction << 4),    // 22
			AnywhereState | (CollectAction << 4),    // 23
			AnywhereState | (CollectAction << 4),    // 24
			AnywhereState | (CollectAction << 4),    // 25
			AnywhereState | (CollectAction << 4),    // 26
			AnywhereState | (CollectAction << 4),    // 27
			AnywhereState | (CollectAction << 4),    // 28
			AnywhereState | (CollectAction << 4),    // 29
			AnywhereState | (CollectAction << 4),    // 2A
			AnywhereState | (CollectAction << 4),    // 2B
			AnywhereState | (CollectAction << 4),    // 2C
			AnywhereState | (CollectAction << 4),    // 2D
			AnywhereState | (CollectAction << 4),    // 2E
			AnywhereState | (CollectAction << 4),    // 2F
			DcsIgnoreState | (NoneAction << 4),      // 30
			DcsIgnoreState | (NoneAction << 4),      // 31
			DcsIgnoreState | (NoneAction << 4),      // 32
			DcsIgnoreState | (NoneAction << 4),      // 33
			DcsIgnoreState | (NoneAction << 4),      // 34
			DcsIgnoreState | (NoneAction << 4),      // 35
			DcsIgnoreState | (NoneAction << 4),      // 36
			DcsIgnoreState | (NoneAction << 4),      // 37
			DcsIgnoreState | (NoneAction << 4),      // 38
			DcsIgnoreState | (NoneAction << 4),      // 39
			DcsIgnoreState | (NoneAction << 4),      // 3A
			DcsIgnoreState | (NoneAction << 4),      // 3B
			DcsIgnoreState | (NoneAction << 4),      // 3C
			DcsIgnoreState | (NoneAction << 4),      // 3D
			DcsIgnoreState | (NoneAction << 4),      // 3E
			DcsIgnoreState | (NoneAction << 4),      // 3F
			DcsPassthroughState | (NoneAction << 4), // 40
			DcsPassthroughState | (NoneAction << 4), // 41
			DcsPassthroughState | (NoneAction << 4), // 42
			DcsPassthroughState | (NoneAction << 4), // 43
			DcsPassthroughState | (NoneAction << 4), // 44
			DcsPassthroughState | (NoneAction << 4), // 45
			DcsPassthroughState | (NoneAction << 4), // 46
			DcsPassthroughState | (NoneAction << 4), // 47
			DcsPassthroughState | (NoneAction << 4), // 48
			DcsPassthroughState | (NoneAction << 4), // 49
			DcsPassthroughState | (NoneAction << 4), // 4A
			DcsPassthroughState | (NoneAction << 4), // 4B
			DcsPassthroughState | (NoneAction << 4), // 4C
			DcsPassthroughState | (NoneAction << 4), // 4D
			DcsPassthroughState | (NoneAction << 4), // 4E
			DcsPassthroughState | (NoneAction << 4), // 4F
			DcsPassthroughState | (NoneAction << 4), // 50
			DcsPassthroughState | (NoneAction << 4), // 51
			DcsPassthroughState | (NoneAction << 4), // 52
			DcsPassthroughState | (NoneAction << 4), // 53
			DcsPassthroughState | (NoneAction << 4), // 54
			DcsPassthroughState | (NoneAction << 4), // 55
			DcsPassthroughState | (NoneAction << 4), // 56
			DcsPassthroughState | (NoneAction << 4), // 57
			DcsPassthroughState | (NoneAction << 4), // 58
			DcsPassthroughState | (NoneAction << 4), // 59
			DcsPassthroughState | (NoneAction << 4), // 5A
			DcsPassthroughState | (NoneAction << 4), // 5B
			DcsPassthroughState | (NoneAction << 4), // 5C
			DcsPassthroughState | (NoneAction << 4), // 5D
			DcsPassthroughState | (NoneAction << 4), // 5E
			DcsPassthroughState | (NoneAction << 4), // 5F
			DcsPassthroughState | (NoneAction << 4), // 60
			DcsPassthroughState | (NoneAction << 4), // 61
			DcsPassthroughState | (NoneAction << 4), // 62
			DcsPassthroughState | (NoneAction << 4), // 63
			DcsPassthroughState | (NoneAction << 4), // 64
			DcsPassthroughState | (NoneAction << 4), // 65
			DcsPassthroughState | (NoneAction << 4), // 66
			DcsPassthroughState | (NoneAction << 4), // 67
			DcsPassthroughState | (NoneAction << 4), // 68
			DcsPassthroughState | (NoneAction << 4), // 69
			DcsPassthroughState | (NoneAction << 4), // 6A
			DcsPassthroughState | (NoneAction << 4), // 6B
			DcsPassthroughState | (NoneAction << 4), // 6C
			DcsPassthroughState | (NoneAction << 4), // 6D
			DcsPassthroughState | (NoneAction << 4), // 6E
			DcsPassthroughState | (NoneAction << 4), // 6F
			DcsPassthroughState | (NoneAction << 4), // 70
			DcsPassthroughState | (NoneAction << 4), // 71
			DcsPassthroughState | (NoneAction << 4), // 72
			DcsPassthroughState | (NoneAction << 4), // 73
			DcsPassthroughState | (NoneAction << 4), // 74
			DcsPassthroughState | (NoneAction << 4), // 75
			DcsPassthroughState | (NoneAction << 4), // 76
			DcsPassthroughState | (NoneAction << 4), // 77
			DcsPassthroughState | (NoneAction << 4), // 78
			DcsPassthroughState | (NoneAction << 4), // 79
			DcsPassthroughState | (NoneAction << 4), // 7A
			DcsPassthroughState | (NoneAction << 4), // 7B
			DcsPassthroughState | (NoneAction << 4), // 7C
			DcsPassthroughState | (NoneAction << 4), // 7D
			DcsPassthroughState | (NoneAction << 4), // 7E
			AnywhereState | (IgnoreAction << 4),     // 7F
			0,                                       // 80
			0,                                       // 81
			0,                                       // 82
			0,                                       // 83
			0,                                       // 84
			0,                                       // 85
			0,                                       // 86
			0,                                       // 87
			0,                                       // 88
			0,                                       // 89
			0,                                       // 8A
			0,                                       // 8B
			0,                                       // 8C
			0,                                       // 8D
			0,                                       // 8E
			0,                                       // 8F
			0,                                       // 90
			0,                                       // 91
			0,                                       // 92
			0,                                       // 93
			0,                                       // 94
			0,                                       // 95
			0,                                       // 96
			0,                                       // 97
			0,                                       // 98
			0,                                       // 99
			0,                                       // 9A
			0,                                       // 9B
			0,                                       // 9C
			0,                                       // 9D
			0,                                       // 9E
			0,                                       // 9F
			0,                                       // A0
			0,                                       // A1
			0,                                       // A2
			0,                                       // A3
			0,                                       // A4
			0,                                       // A5
			0,                                       // A6
			0,                                       // A7
			0,                                       // A8
			0,                                       // A9
			0,                                       // AA
			0,                                       // AB
			0,                                       // AC
			0,                                       // AD
			0,                                       // AE
			0,                                       // AF
			0,                                       // B0
			0,                                       // B1
			0,                                       // B2
			0,                                       // B3
			0,                                       // B4
			0,                                       // B5
			0,                                       // B6
			0,                                       // B7
			0,                                       // B8
			0,                                       // B9
			0,                                       // BA
			0,                                       // BB
			0,                                       // BC
			0,                                       // BD
			0,                                       // BE
			0,                                       // BF
			0,                                       // C0
			0,                                       // C1
			0,                                       // C2
			0,                                       // C3
			0,                                       // C4
			0,                                       // C5
			0,                                       // C6
			0,                                       // C7
			0,                                       // C8
			0,                                       // C9
			0,                                       // CA
			0,                                       // CB
			0,                                       // CC
			0,                                       // CD
			0,                                       // CE
			0,                                       // CF
			0,                                       // D0
			0,                                       // D1
			0,                                       // D2
			0,                                       // D3
			0,                                       // D4
			0,                                       // D5
			0,                                       // D6
			0,                                       // D7
			0,                                       // D8
			0,                                       // D9
			0,                                       // DA
			0,                                       // DB
			0,                                       // DC
			0,                                       // DD
			0,                                       // DE
			0,                                       // DF
			0,                                       // E0
			0,                                       // E1
			0,                                       // E2
			0,                                       // E3
			0,                                       // E4
			0,                                       // E5
			0,                                       // E6
			0,                                       // E7
			0,                                       // E8
			0,                                       // E9
			0,                                       // EA
			0,                                       // EB
			0,                                       // EC
			0,                                       // ED
			0,                                       // EE
			0,                                       // EF
			0,                                       // F0
			0,                                       // F1
			0,                                       // F2
			0,                                       // F3
			0,                                       // F4
			0,                                       // F5
			0,                                       // F6
			0,                                       // F7
			0,                                       // F8
			0,                                       // F9
			0,                                       // FA
			0,                                       // FB
			0,                                       // FC
			0,                                       // FD
			0,                                       // FE
			0,                                       // FF
		},
		{ //State DcsParamState = 8
			AnywhereState | (IgnoreAction << 4),         // 00
			AnywhereState | (IgnoreAction << 4),         // 01
			AnywhereState | (IgnoreAction << 4),         // 02
			AnywhereState | (IgnoreAction << 4),         // 03
			AnywhereState | (IgnoreAction << 4),         // 04
			AnywhereState | (IgnoreAction << 4),         // 05
			AnywhereState | (IgnoreAction << 4),         // 06
			AnywhereState | (IgnoreAction << 4),         // 07
			AnywhereState | (IgnoreAction << 4),         // 08
			AnywhereState | (IgnoreAction << 4),         // 09
			AnywhereState | (IgnoreAction << 4),         // 0A
			AnywhereState | (IgnoreAction << 4),         // 0B
			AnywhereState | (IgnoreAction << 4),         // 0C
			AnywhereState | (IgnoreAction << 4),         // 0D
			AnywhereState | (IgnoreAction << 4),         // 0E
			AnywhereState | (IgnoreAction << 4),         // 0F
			AnywhereState | (IgnoreAction << 4),         // 10
			AnywhereState | (IgnoreAction << 4),         // 11
			AnywhereState | (IgnoreAction << 4),         // 12
			AnywhereState | (IgnoreAction << 4),         // 13
			AnywhereState | (IgnoreAction << 4),         // 14
			AnywhereState | (IgnoreAction << 4),         // 15
			AnywhereState | (IgnoreAction << 4),         // 16
			AnywhereState | (IgnoreAction << 4),         // 17
			0,                                           // 18
			AnywhereState | (IgnoreAction << 4),         // 19
			0,                                           // 1A
			0,                                           // 1B
			AnywhereState | (IgnoreAction << 4),         // 1C
			AnywhereState | (IgnoreAction << 4),         // 1D
			AnywhereState | (IgnoreAction << 4),         // 1E
			AnywhereState | (IgnoreAction << 4),         // 1F
			DcsIntermediateState | (CollectAction << 4), // 20
			DcsIntermediateState | (CollectAction << 4), // 21
			DcsIntermediateState | (CollectAction << 4), // 22
			DcsIntermediateState | (CollectAction << 4), // 23
			DcsIntermediateState | (CollectAction << 4), // 24
			DcsIntermediateState | (CollectAction << 4), // 25
			DcsIntermediateState | (CollectAction << 4), // 26
			DcsIntermediateState | (CollectAction << 4), // 27
			DcsIntermediateState | (CollectAction << 4), // 28
			DcsIntermediateState | (CollectAction << 4), // 29
			DcsIntermediateState | (CollectAction << 4), // 2A
			DcsIntermediateState | (CollectAction << 4), // 2B
			DcsIntermediateState | (CollectAction << 4), // 2C
			DcsIntermediateState | (CollectAction << 4), // 2D
			DcsIntermediateState | (CollectAction << 4), // 2E
			DcsIntermediateState | (CollectAction << 4), // 2F
			AnywhereState | (ParamAction << 4),          // 30
			AnywhereState | (ParamAction << 4),          // 31
			AnywhereState | (ParamAction << 4),          // 32
			AnywhereState | (ParamAction << 4),          // 33
			AnywhereState | (ParamAction << 4),          // 34
			AnywhereState | (ParamAction << 4),          // 35
			AnywhereState | (ParamAction << 4),          // 36
			AnywhereState | (ParamAction << 4),          // 37
			AnywhereState | (ParamAction << 4),          // 38
			AnywhereState | (ParamAction << 4),          // 39
			AnywhereState | (ParamAction << 4),          // 3A
			AnywhereState | (ParamAction << 4),          // 3B
			DcsIgnoreState | (NoneAction << 4),          // 3C
			DcsIgnoreState | (NoneAction << 4),          // 3D
			DcsIgnoreState | (NoneAction << 4),          // 3E
			DcsIgnoreState | (NoneAction << 4),          // 3F
			DcsPassthroughState | (NoneAction << 4),     // 40
			DcsPassthroughState | (NoneAction << 4),     // 41
			DcsPassthroughState | (NoneAction << 4),     // 42
			DcsPassthroughState | (NoneAction << 4),     // 43
			DcsPassthroughState | (NoneAction << 4),     // 44
			DcsPassthroughState | (NoneAction << 4),     // 45
			DcsPassthroughState | (NoneAction << 4),     // 46
			DcsPassthroughState | (NoneAction << 4),     // 47
			DcsPassthroughState | (NoneAction << 4),     // 48
			DcsPassthroughState | (NoneAction << 4),     // 49
			DcsPassthroughState | (NoneAction << 4),     // 4A
			DcsPassthroughState | (NoneAction << 4),     // 4B
			DcsPassthroughState | (NoneAction << 4),     // 4C
			DcsPassthroughState | (NoneAction << 4),     // 4D
			DcsPassthroughState | (NoneAction << 4),     // 4E
			DcsPassthroughState | (NoneAction << 4),     // 4F
			DcsPassthroughState | (NoneAction << 4),     // 50
			DcsPassthroughState | (NoneAction << 4),     // 51
			DcsPassthroughState | (NoneAction << 4),     // 52
			DcsPassthroughState | (NoneAction << 4),     // 53
			DcsPassthroughState | (NoneAction << 4),     // 54
			DcsPassthroughState | (NoneAction << 4),     // 55
			DcsPassthroughState | (NoneAction << 4),     // 56
			DcsPassthroughState | (NoneAction << 4),     // 57
			DcsPassthroughState | (NoneAction << 4),     // 58
			DcsPassthroughState | (NoneAction << 4),     // 59
			DcsPassthroughState | (NoneAction << 4),     // 5A
			DcsPassthroughState | (NoneAction << 4),     // 5B
			DcsPassthroughState | (NoneAction << 4),     // 5C
			DcsPassthroughState | (NoneAction << 4),     // 5D
			DcsPassthroughState | (NoneAction << 4),     // 5E
			DcsPassthroughState | (NoneAction << 4),     // 5F
			DcsPassthroughState | (NoneAction << 4),     // 60
			DcsPassthroughState | (NoneAction << 4),     // 61
			DcsPassthroughState | (NoneAction << 4),     // 62
			DcsPassthroughState | (NoneAction << 4),     // 63
			DcsPassthroughState | (NoneAction << 4),     // 64
			DcsPassthroughState | (NoneAction << 4),     // 65
			DcsPassthroughState | (NoneAction << 4),     // 66
			DcsPassthroughState | (NoneAction << 4),     // 67
			DcsPassthroughState | (NoneAction << 4),     // 68
			DcsPassthroughState | (NoneAction << 4),     // 69
			DcsPassthroughState | (NoneAction << 4),     // 6A
			DcsPassthroughState | (NoneAction << 4),     // 6B
			DcsPassthroughState | (NoneAction << 4),     // 6C
			DcsPassthroughState | (NoneAction << 4),     // 6D
			DcsPassthroughState | (NoneAction << 4),     // 6E
			DcsPassthroughState | (NoneAction << 4),     // 6F
			DcsPassthroughState | (NoneAction << 4),     // 70
			DcsPassthroughState | (NoneAction << 4),     // 71
			DcsPassthroughState | (NoneAction << 4),     // 72
			DcsPassthroughState | (NoneAction << 4),     // 73
			DcsPassthroughState | (NoneAction << 4),     // 74
			DcsPassthroughState | (NoneAction << 4),     // 75
			DcsPassthroughState | (NoneAction << 4),     // 76
			DcsPassthroughState | (NoneAction << 4),     // 77
			DcsPassthroughState | (NoneAction << 4),     // 78
			DcsPassthroughState | (NoneAction << 4),     // 79
			DcsPassthroughState | (NoneAction << 4),     // 7A
			DcsPassthroughState | (NoneAction << 4),     // 7B
			DcsPassthroughState | (NoneAction << 4),     // 7C
			DcsPassthroughState | (NoneAction << 4),     // 7D
			DcsPassthroughState | (NoneAction << 4),     // 7E
			AnywhereState | (IgnoreAction << 4),         // 7F
			0,                                           // 80
			0,                                           // 81
			0,                                           // 82
			0,                                           // 83
			0,                                           // 84
			0,                                           // 85
			0,                                           // 86
			0,                                           // 87
			0,                                           // 88
			0,                                           // 89
			0,                                           // 8A
			0,                                           // 8B
			0,                                           // 8C
			0,                                           // 8D
			0,                                           // 8E
			0,                                           // 8F
			0,                                           // 90
			0,                                           // 91
			0,                                           // 92
			0,                                           // 93
			0,                                           // 94
			0,                                           // 95
			0,                                           // 96
			0,                                           // 97
			0,                                           // 98
			0,                                           // 99
			0,                                           // 9A
			0,                                           // 9B
			0,                                           // 9C
			0,                                           // 9D
			0,                                           // 9E
			0,                                           // 9F
			0,                                           // A0
			0,                                           // A1
			0,                                           // A2
			0,                                           // A3
			0,                                           // A4
			0,                                           // A5
			0,                                           // A6
			0,                                           // A7
			0,                                           // A8
			0,                                           // A9
			0,                                           // AA
			0,                                           // AB
			0,                                           // AC
			0,                                           // AD
			0,                                           // AE
			0,                                           // AF
			0,                                           // B0
			0,                                           // B1
			0,                                           // B2
			0,                                           // B3
			0,                                           // B4
			0,                                           // B5
			0,                                           // B6
			0,                                           // B7
			0,                                           // B8
			0,                                           // B9
			0,                                           // BA
			0,                                           // BB
			0,                                           // BC
			0,                                           // BD
			0,                                           // BE
			0,                                           // BF
			0,                                           // C0
			0,                                           // C1
			0,                                           // C2
			0,                                           // C3
			0,                                           // C4
			0,                                           // C5
			0,                                           // C6
			0,                                           // C7
			0,                                           // C8
			0,                                           // C9
			0,                                           // CA
			0,                                           // CB
			0,                                           // CC
			0,                                           // CD
			0,                                           // CE
			0,                                           // CF
			0,                                           // D0
			0,                                           // D1
			0,                                           // D2
			0,                                           // D3
			0,                                           // D4
			0,                                           // D5
			0,                                           // D6
			0,                                           // D7
			0,                                           // D8
			0,                                           // D9
			0,                                           // DA
			0,                                           // DB
			0,                                           // DC
			0,                                           // DD
			0,                                           // DE
			0,                                           // DF
			0,                                           // E0
			0,                                           // E1
			0,                                           // E2
			0,                                           // E3
			0,                                           // E4
			0,                                           // E5
			0,                                           // E6
			0,                                           // E7
			0,                                           // E8
			0,                                           // E9
			0,                                           // EA
			0,                                           // EB
			0,                                           // EC
			0,                                           // ED
			0,                                           // EE
			0,                                           // EF
			0,                                           // F0
			0,                                           // F1
			0,                                           // F2
			0,                                           // F3
			0,                                           // F4
			0,                                           // F5
			0,                                           // F6
			0,                                           // F7
			0,                                           // F8
			0,                                           // F9
			0,                                           // FA
			0,                                           // FB
			0,                                           // FC
			0,                                           // FD
			0,                                           // FE
			0,                                           // FF
		},
		{ //State DcsPassthroughState = 9
			AnywhereState | (PutAction << 4),    // 00
			AnywhereState | (PutAction << 4),    // 01
			AnywhereState | (PutAction << 4),    // 02
			AnywhereState | (PutAction << 4),    // 03
			AnywhereState | (PutAction << 4),    // 04
			AnywhereState | (PutAction << 4),    // 05
			AnywhereState | (PutAction << 4),    // 06
			AnywhereState | (PutAction << 4),    // 07
			AnywhereState | (PutAction << 4),    // 08
			AnywhereState | (PutAction << 4),    // 09
			AnywhereState | (PutAction << 4),    // 0A
			AnywhereState | (PutAction << 4),    // 0B
			AnywhereState | (PutAction << 4),    // 0C
			AnywhereState | (PutAction << 4),    // 0D
			AnywhereState | (PutAction << 4),    // 0E
			AnywhereState | (PutAction << 4),    // 0F
			AnywhereState | (PutAction << 4),    // 10
			AnywhereState | (PutAction << 4),    // 11
			AnywhereState | (PutAction << 4),    // 12
			AnywhereState | (PutAction << 4),    // 13
			AnywhereState | (PutAction << 4),    // 14
			AnywhereState | (PutAction << 4),    // 15
			AnywhereState | (PutAction << 4),    // 16
			AnywhereState | (PutAction << 4),    // 17
			0,                                   // 18
			AnywhereState | (PutAction << 4),    // 19
			0,                                   // 1A
			0,                                   // 1B
			AnywhereState | (PutAction << 4),    // 1C
			AnywhereState | (PutAction << 4),    // 1D
			AnywhereState | (PutAction << 4),    // 1E
			AnywhereState | (PutAction << 4),    // 1F
			AnywhereState | (PutAction << 4),    // 20
			AnywhereState | (PutAction << 4),    // 21
			AnywhereState | (PutAction << 4),    // 22
			AnywhereState | (PutAction << 4),    // 23
			AnywhereState | (PutAction << 4),    // 24
			AnywhereState | (PutAction << 4),    // 25
			AnywhereState | (PutAction << 4),    // 26
			AnywhereState | (PutAction << 4),    // 27
			AnywhereState | (PutAction << 4),    // 28
			AnywhereState | (PutAction << 4),    // 29
			AnywhereState | (PutAction << 4),    // 2A
			AnywhereState | (PutAction << 4),    // 2B
			AnywhereState | (PutAction << 4),    // 2C
			AnywhereState | (PutAction << 4),    // 2D
			AnywhereState | (PutAction << 4),    // 2E
			AnywhereState | (PutAction << 4),    // 2F
			AnywhereState | (PutAction << 4),    // 30
			AnywhereState | (PutAction << 4),    // 31
			AnywhereState | (PutAction << 4),    // 32
			AnywhereState | (PutAction << 4),    // 33
			AnywhereState | (PutAction << 4),    // 34
			AnywhereState | (PutAction << 4),    // 35
			AnywhereState | (PutAction << 4),    // 36
			AnywhereState | (PutAction << 4),    // 37
			AnywhereState | (PutAction << 4),    // 38
			AnywhereState | (PutAction << 4),    // 39
			AnywhereState | (PutAction << 4),    // 3A
			AnywhereState | (PutAction << 4),    // 3B
			AnywhereState | (PutAction << 4),    // 3C
			AnywhereState | (PutAction << 4),    // 3D
			AnywhereState | (PutAction << 4),    // 3E
			AnywhereState | (PutAction << 4),    // 3F
			AnywhereState | (PutAction << 4),    // 40
			AnywhereState | (PutAction << 4),    // 41
			AnywhereState | (PutAction << 4),    // 42
			AnywhereState | (PutAction << 4),    // 43
			AnywhereState | (PutAction << 4),    // 44
			AnywhereState | (PutAction << 4),    // 45
			AnywhereState | (PutAction << 4),    // 46
			AnywhereState | (PutAction << 4),    // 47
			AnywhereState | (PutAction << 4),    // 48
			AnywhereState | (PutAction << 4),    // 49
			AnywhereState | (PutAction << 4),    // 4A
			AnywhereState | (PutAction << 4),    // 4B
			AnywhereState | (PutAction << 4),    // 4C
			AnywhereState | (PutAction << 4),    // 4D
			AnywhereState | (PutAction << 4),    // 4E
			AnywhereState | (PutAction << 4),    // 4F
			AnywhereState | (PutAction << 4),    // 50
			AnywhereState | (PutAction << 4),    // 51
			AnywhereState | (PutAction << 4),    // 52
			AnywhereState | (PutAction << 4),    // 53
			AnywhereState | (PutAction << 4),    // 54
			AnywhereState | (PutAction << 4),    // 55
			AnywhereState | (PutAction << 4),    // 56
			AnywhereState | (PutAction << 4),    // 57
			AnywhereState | (PutAction << 4),    // 58
			AnywhereState | (PutAction << 4),    // 59
			AnywhereState | (PutAction << 4),    // 5A
			AnywhereState | (PutAction << 4),    // 5B
			AnywhereState | (PutAction << 4),    // 5C
			AnywhereState | (PutAction << 4),    // 5D
			AnywhereState | (PutAction << 4),    // 5E
			AnywhereState | (PutAction << 4),    // 5F
			AnywhereState | (PutAction << 4),    // 60
			AnywhereState | (PutAction << 4),    // 61
			AnywhereState | (PutAction << 4),    // 62
			AnywhereState | (PutAction << 4),    // 63
			AnywhereState | (PutAction << 4),    // 64
			AnywhereState | (PutAction << 4),    // 65
			AnywhereState | (PutAction << 4),    // 66
			AnywhereState | (PutAction << 4),    // 67
			AnywhereState | (PutAction << 4),    // 68
			AnywhereState | (PutAction << 4),    // 69
			AnywhereState | (PutAction << 4),    // 6A
			AnywhereState | (PutAction << 4),    // 6B
			AnywhereState | (PutAction << 4),    // 6C
			AnywhereState | (PutAction << 4),    // 6D
			AnywhereState | (PutAction << 4),    // 6E
			AnywhereState | (PutAction << 4),    // 6F
			AnywhereState | (PutAction << 4),    // 70
			AnywhereState | (PutAction << 4),    // 71
			AnywhereState | (PutAction << 4),    // 72
			AnywhereState | (PutAction << 4),    // 73
			AnywhereState | (PutAction << 4),    // 74
			AnywhereState | (PutAction << 4),    // 75
			AnywhereState | (PutAction << 4),    // 76
			AnywhereState | (PutAction << 4),    // 77
			AnywhereState | (PutAction << 4),    // 78
			AnywhereState | (PutAction << 4),    // 79
			AnywhereState | (PutAction << 4),    // 7A
			AnywhereState | (PutAction << 4),    // 7B
			AnywhereState | (PutAction << 4),    // 7C
			AnywhereState | (PutAction << 4),    // 7D
			AnywhereState | (PutAction << 4),    // 7E
			AnywhereState | (IgnoreAction << 4), // 7F
			0,                                   // 80
			0,                                   // 81
			0,                                   // 82
			0,                                   // 83
			0,                                   // 84
			0,                                   // 85
			0,                                   // 86
			0,                                   // 87
			0,                                   // 88
			0,                                   // 89
			0,                                   // 8A
			0,                                   // 8B
			0,                                   // 8C
			0,                                   // 8D
			0,                                   // 8E
			0,                                   // 8F
			0,                                   // 90
			0,                                   // 91
			0,                                   // 92
			0,                                   // 93
			0,                                   // 94
			0,                                   // 95
			0,                                   // 96
			0,                                   // 97
			0,                                   // 98
			0,                                   // 99
			0,                                   // 9A
			0,                                   // 9B
			GroundState | (NoneAction << 4),     // 9C
			0,                                   // 9D
			0,                                   // 9E
			0,                                   // 9F
			0,                                   // A0
			0,                                   // A1
			0,                                   // A2
			0,                                   // A3
			0,                                   // A4
			0,                                   // A5
			0,                                   // A6
			0,                                   // A7
			0,                                   // A8
			0,                                   // A9
			0,                                   // AA
			0,                                   // AB
			0,                                   // AC
			0,                                   // AD
			0,                                   // AE
			0,                                   // AF
			0,                                   // B0
			0,                                   // B1
			0,                                   // B2
			0,                                   // B3
			0,                                   // B4
			0,                                   // B5
			0,                                   // B6
			0,                                   // B7
			0,                                   // B8
			0,                                   // B9
			0,                                   // BA
			0,                                   // BB
			0,                                   // BC
			0,                                   // BD
			0,                                   // BE
			0,                                   // BF
			0,                                   // C0
			0,                                   // C1
			0,                                   // C2
			0,                                   // C3
			0,                                   // C4
			0,                                   // C5
			0,                                   // C6
			0,                                   // C7
			0,                                   // C8
			0,                                   // C9
			0,                                   // CA
			0,                                   // CB
			0,                                   // CC
			0,                                   // CD
			0,                                   // CE
			0,                                   // CF
			0,                                   // D0
			0,                                   // D1
			0,                                   // D2
			0,                                   // D3
			0,                                   // D4
			0,                                   // D5
			0,                                   // D6
			0,                                   // D7
			0,                                   // D8
			0,                                   // D9
			0,                                   // DA
			0,                                   // DB
			0,                                   // DC
			0,                                   // DD
			0,                                   // DE
			0,                                   // DF
			0,                                   // E0
			0,                                   // E1
			0,                                   // E2
			0,                                   // E3
			0,                                   // E4
			0,                                   // E5
			0,                                   // E6
			0,                                   // E7
			0,                                   // E8
			0,                                   // E9
			0,                                   // EA
			0,                                   // EB
			0,                                   // EC
			0,                                   // ED
			0,                                   // EE
			0,                                   // EF
			0,                                   // F0
			0,                                   // F1
			0,                                   // F2
			0,                                   // F3
			0,                                   // F4
			0,                                   // F5
			0,                                   // F6
			0,                                   // F7
			0,                                   // F8
			0,                                   // F9
			0,                                   // FA
			0,                                   // FB
			0,                                   // FC
			0,                                   // FD
			0,                                   // FE
			0,                                   // FF
		},
		{ //State EscapeState = 10
			AnywhereState | (ExecuteAction << 4),           // 00
			AnywhereState | (ExecuteAction << 4),           // 01
			AnywhereState | (ExecuteAction << 4),           // 02
			AnywhereState | (ExecuteAction << 4),           // 03
			AnywhereState | (ExecuteAction << 4),           // 04
			AnywhereState | (ExecuteAction << 4),           // 05
			AnywhereState | (ExecuteAction << 4),           // 06
			AnywhereState | (ExecuteAction << 4),           // 07
			AnywhereState | (ExecuteAction << 4),           // 08
			AnywhereState | (ExecuteAction << 4),           // 09
			AnywhereState | (ExecuteAction << 4),           // 0A
			AnywhereState | (ExecuteAction << 4),           // 0B
			AnywhereState | (ExecuteAction << 4),           // 0C
			AnywhereState | (ExecuteAction << 4),           // 0D
			AnywhereState | (ExecuteAction << 4),           // 0E
			AnywhereState | (ExecuteAction << 4),           // 0F
			AnywhereState | (ExecuteAction << 4),           // 10
			AnywhereState | (ExecuteAction << 4),           // 11
			AnywhereState | (ExecuteAction << 4),           // 12
			AnywhereState | (ExecuteAction << 4),           // 13
			AnywhereState | (ExecuteAction << 4),           // 14
			AnywhereState | (ExecuteAction << 4),           // 15
			AnywhereState | (ExecuteAction << 4),           // 16
			AnywhereState | (ExecuteAction << 4),           // 17
			0,                                              // 18
			AnywhereState | (ExecuteAction << 4),           // 19
			0,                                              // 1A
			0,                                              // 1B
			AnywhereState | (ExecuteAction << 4),           // 1C
			AnywhereState | (ExecuteAction << 4),           // 1D
			AnywhereState | (ExecuteAction << 4),           // 1E
			AnywhereState | (ExecuteAction << 4),           // 1F
			EscapeIntermediateState | (CollectAction << 4), // 20
			EscapeIntermediateState | (CollectAction << 4), // 21
			EscapeIntermediateState | (CollectAction << 4), // 22
			EscapeIntermediateState | (CollectAction << 4), // 23
			EscapeIntermediateState | (CollectAction << 4), // 24
			EscapeIntermediateState | (CollectAction << 4), // 25
			EscapeIntermediateState | (CollectAction << 4), // 26
			EscapeIntermediateState | (CollectAction << 4), // 27
			EscapeIntermediateState | (CollectAction << 4), // 28
			EscapeIntermediateState | (CollectAction << 4), // 29
			EscapeIntermediateState | (CollectAction << 4), // 2A
			EscapeIntermediateState | (CollectAction << 4), // 2B
			EscapeIntermediateState | (CollectAction << 4), // 2C
			EscapeIntermediateState | (CollectAction << 4), // 2D
			EscapeIntermediateState | (CollectAction << 4), // 2E
			EscapeIntermediateState | (CollectAction << 4), // 2F
			GroundState | (EscDispatchAction << 4),         // 30
			GroundState | (EscDispatchAction << 4),         // 31
			GroundState | (EscDispatchAction << 4),         // 32
			GroundState | (EscDispatchAction << 4),         // 33
			GroundState | (EscDispatchAction << 4),         // 34
			GroundState | (EscDispatchAction << 4),         // 35
			GroundState | (EscDispatchAction << 4),         // 36
			GroundState | (EscDispatchAction << 4),         // 37
			GroundState | (EscDispatchAction << 4),         // 38
			GroundState | (EscDispatchAction << 4),         // 39
			GroundState | (EscDispatchAction << 4),         // 3A
			GroundState | (EscDispatchAction << 4),         // 3B
			GroundState | (EscDispatchAction << 4),         // 3C
			GroundState | (EscDispatchAction << 4),         // 3D
			GroundState | (EscDispatchAction << 4),         // 3E
			GroundState | (EscDispatchAction << 4),         // 3F
			GroundState | (EscDispatchAction << 4),         // 40
			GroundState | (EscDispatchAction << 4),         // 41
			GroundState | (EscDispatchAction << 4),         // 42
			GroundState | (EscDispatchAction << 4),         // 43
			GroundState | (EscDispatchAction << 4),         // 44
			GroundState | (EscDispatchAction << 4),         // 45
			GroundState | (EscDispatchAction << 4),         // 46
			GroundState | (EscDispatchAction << 4),         // 47
			GroundState | (EscDispatchAction << 4),         // 48
			GroundState | (EscDispatchAction << 4),         // 49
			GroundState | (EscDispatchAction << 4),         // 4A
			GroundState | (EscDispatchAction << 4),         // 4B
			GroundState | (EscDispatchAction << 4),         // 4C
			GroundState | (EscDispatchAction << 4),         // 4D
			GroundState | (EscDispatchAction << 4),         // 4E
			GroundState | (EscDispatchAction << 4),         // 4F
			DcsEntryState | (NoneAction << 4),              // 50
			GroundState | (EscDispatchAction << 4),         // 51
			GroundState | (EscDispatchAction << 4),         // 52
			GroundState | (EscDispatchAction << 4),         // 53
			GroundState | (EscDispatchAction << 4),         // 54
			GroundState | (EscDispatchAction << 4),         // 55
			GroundState | (EscDispatchAction << 4),         // 56
			GroundState | (EscDispatchAction << 4),         // 57
			SosPmApcStringState | (NoneAction << 4),        // 58
			GroundState | (EscDispatchAction << 4),         // 59
			GroundState | (EscDispatchAction << 4),         // 5A
			CsiEntryState | (NoneAction << 4),              // 5B
			GroundState | (EscDispatchAction << 4),         // 5C
			OscStringState | (NoneAction << 4),             // 5D
			SosPmApcStringState | (NoneAction << 4),        // 5E
			SosPmApcStringState | (NoneAction << 4),        // 5F
			GroundState | (EscDispatchAction << 4),         // 60
			GroundState | (EscDispatchAction << 4),         // 61
			GroundState | (EscDispatchAction << 4),         // 62
			GroundState | (EscDispatchAction << 4),         // 63
			GroundState | (EscDispatchAction << 4),         // 64
			GroundState | (EscDispatchAction << 4),         // 65
			GroundState | (EscDispatchAction << 4),         // 66
			GroundState | (EscDispatchAction << 4),         // 67
			GroundState | (EscDispatchAction << 4),         // 68
			GroundState | (EscDispatchAction << 4),         // 69
			GroundState | (EscDispatchAction << 4),         // 6A
			GroundState | (EscDispatchAction << 4),         // 6B
			GroundState | (EscDispatchAction << 4),         // 6C
			GroundState | (EscDispatchAction << 4),         // 6D
			GroundState | (EscDispatchAction << 4),         // 6E
			GroundState | (EscDispatchAction << 4),         // 6F
			GroundState | (EscDispatchAction << 4),         // 70
			GroundState | (EscDispatchAction << 4),         // 71
			GroundState | (EscDispatchAction << 4),         // 72
			GroundState | (EscDispatchAction << 4),         // 73
			GroundState | (EscDispatchAction << 4),         // 74
			GroundState | (EscDispatchAction << 4),         // 75
			GroundState | (EscDispatchAction << 4),         // 76
			GroundState | (EscDispatchAction << 4),         // 77
			GroundState | (EscDispatchAction << 4),         // 78
			GroundState | (EscDispatchAction << 4),         // 79
			GroundState | (EscDispatchAction << 4),         // 7A
			GroundState | (EscDispatchAction << 4),         // 7B
			GroundState | (EscDispatchAction << 4),         // 7C
			GroundState | (EscDispatchAction << 4),         // 7D
			GroundState | (EscDispatchAction << 4),         // 7E
			AnywhereState | (IgnoreAction << 4),            // 7F
			0,                                              // 80
			0,                                              // 81
			0,                                              // 82
			0,                                              // 83
			0,                                              // 84
			0,                                              // 85
			0,                                              // 86
			0,                                              // 87
			0,                                              // 88
			0,                                              // 89
			0,                                              // 8A
			0,                                              // 8B
			0,                                              // 8C
			0,                                              // 8D
			0,                                              // 8E
			0,                                              // 8F
			0,                                              // 90
			0,                                              // 91
			0,                                              // 92
			0,                                              // 93
			0,                                              // 94
			0,                                              // 95
			0,                                              // 96
			0,                                              // 97
			0,                                              // 98
			0,                                              // 99
			0,                                              // 9A
			0,                                              // 9B
			0,                                              // 9C
			0,                                              // 9D
			0,                                              // 9E
			0,                                              // 9F
			0,                                              // A0
			0,                                              // A1
			0,                                              // A2
			0,                                              // A3
			0,                                              // A4
			0,                                              // A5
			0,                                              // A6
			0,                                              // A7
			0,                                              // A8
			0,                                              // A9
			0,                                              // AA
			0,                                              // AB
			0,                                              // AC
			0,                                              // AD
			0,                                              // AE
			0,                                              // AF
			0,                                              // B0
			0,                                              // B1
			0,                                              // B2
			0,                                              // B3
			0,                                              // B4
			0,                                              // B5
			0,                                              // B6
			0,                                              // B7
			0,                                              // B8
			0,                                              // B9
			0,                                              // BA
			0,                                              // BB
			0,                                              // BC
			0,                                              // BD
			0,                                              // BE
			0,                                              // BF
			0,                                              // C0
			0,                                              // C1
			0,                                              // C2
			0,                                              // C3
			0,                                              // C4
			0,                                              // C5
			0,                                              // C6
			0,                                              // C7
			0,                                              // C8
			0,                                              // C9
			0,                                              // CA
			0,                                              // CB
			0,                                              // CC
			0,                                              // CD
			0,                                              // CE
			0,                                              // CF
			0,                                              // D0
			0,                                              // D1
			0,                                              // D2
			0,                                              // D3
			0,                                              // D4
			0,                                              // D5
			0,                                              // D6
			0,                                              // D7
			0,                                              // D8
			0,                                              // D9
			0,                                              // DA
			0,                                              // DB
			0,                                              // DC
			0,                                              // DD
			0,                                              // DE
			0,                                              // DF
			0,                                              // E0
			0,                                              // E1
			0,                                              // E2
			0,                                              // E3
			0,                                              // E4
			0,                                              // E5
			0,                                              // E6
			0,                                              // E7
			0,                                              // E8
			0,                                              // E9
			0,                                              // EA
			0,                                              // EB
			0,                                              // EC
			0,                                              // ED
			0,                                              // EE
			0,                                              // EF
			0,                                              // F0
			0,                                              // F1
			0,                                              // F2
			0,                                              // F3
			0,                                              // F4
			0,                                              // F5
			0,                                              // F6
			0,                                              // F7
			0,                                              // F8
			0,                                              // F9
			0,                                              // FA
			0,                                              // FB
			0,                                              // FC
			0,                                              // FD
			0,                                              // FE
			0,                                              // FF
		},
		{ //State EscapeIntermediateState = 11
			AnywhereState | (ExecuteAction << 4),   // 00
			AnywhereState | (ExecuteAction << 4),   // 01
			AnywhereState | (ExecuteAction << 4),   // 02
			AnywhereState | (ExecuteAction << 4),   // 03
			AnywhereState | (ExecuteAction << 4),   // 04
			AnywhereState | (ExecuteAction << 4),   // 05
			AnywhereState | (ExecuteAction << 4),   // 06
			AnywhereState | (ExecuteAction << 4),   // 07
			AnywhereState | (ExecuteAction << 4),   // 08
			AnywhereState | (ExecuteAction << 4),   // 09
			AnywhereState | (ExecuteAction << 4),   // 0A
			AnywhereState | (ExecuteAction << 4),   // 0B
			AnywhereState | (ExecuteAction << 4),   // 0C
			AnywhereState | (ExecuteAction << 4),   // 0D
			AnywhereState | (ExecuteAction << 4),   // 0E
			AnywhereState | (ExecuteAction << 4),   // 0F
			AnywhereState | (ExecuteAction << 4),   // 10
			AnywhereState | (ExecuteAction << 4),   // 11
			AnywhereState | (ExecuteAction << 4),   // 12
			AnywhereState | (ExecuteAction << 4),   // 13
			AnywhereState | (ExecuteAction << 4),   // 14
			AnywhereState | (ExecuteAction << 4),   // 15
			AnywhereState | (ExecuteAction << 4),   // 16
			AnywhereState | (ExecuteAction << 4),   // 17
			0,                                      // 18
			AnywhereState | (ExecuteAction << 4),   // 19
			0,                                      // 1A
			0,                                      // 1B
			AnywhereState | (ExecuteAction << 4),   // 1C
			AnywhereState | (ExecuteAction << 4),   // 1D
			AnywhereState | (ExecuteAction << 4),   // 1E
			AnywhereState | (ExecuteAction << 4),   // 1F
			AnywhereState | (CollectAction << 4),   // 20
			AnywhereState | (CollectAction << 4),   // 21
			AnywhereState | (CollectAction << 4),   // 22
			AnywhereState | (CollectAction << 4),   // 23
			AnywhereState | (CollectAction << 4),   // 24
			AnywhereState | (CollectAction << 4),   // 25
			AnywhereState | (CollectAction << 4),   // 26
			AnywhereState | (CollectAction << 4),   // 27
			AnywhereState | (CollectAction << 4),   // 28
			AnywhereState | (CollectAction << 4),   // 29
			AnywhereState | (CollectAction << 4),   // 2A
			AnywhereState | (CollectAction << 4),   // 2B
			AnywhereState | (CollectAction << 4),   // 2C
			AnywhereState | (CollectAction << 4),   // 2D
			AnywhereState | (CollectAction << 4),   // 2E
			AnywhereState | (CollectAction << 4),   // 2F
			GroundState | (EscDispatchAction << 4), // 30
			GroundState | (EscDispatchAction << 4), // 31
			GroundState | (EscDispatchAction << 4), // 32
			GroundState | (EscDispatchAction << 4), // 33
			GroundState | (EscDispatchAction << 4), // 34
			GroundState | (EscDispatchAction << 4), // 35
			GroundState | (EscDispatchAction << 4), // 36
			GroundState | (EscDispatchAction << 4), // 37
			GroundState | (EscDispatchAction << 4), // 38
			GroundState | (EscDispatchAction << 4), // 39
			GroundState | (EscDispatchAction << 4), // 3A
			GroundState | (EscDispatchAction << 4), // 3B
			GroundState | (EscDispatchAction << 4), // 3C
			GroundState | (EscDispatchAction << 4), // 3D
			GroundState | (EscDispatchAction << 4), // 3E
			GroundState | (EscDispatchAction << 4), // 3F
			GroundState | (EscDispatchAction << 4), // 40
			GroundState | (EscDispatchAction << 4), // 41
			GroundState | (EscDispatchAction << 4), // 42
			GroundState | (EscDispatchAction << 4), // 43
			GroundState | (EscDispatchAction << 4), // 44
			GroundState | (EscDispatchAction << 4), // 45
			GroundState | (EscDispatchAction << 4), // 46
			GroundState | (EscDispatchAction << 4), // 47
			GroundState | (EscDispatchAction << 4), // 48
			GroundState | (EscDispatchAction << 4), // 49
			GroundState | (EscDispatchAction << 4), // 4A
			GroundState | (EscDispatchAction << 4), // 4B
			GroundState | (EscDispatchAction << 4), // 4C
			GroundState | (EscDispatchAction << 4), // 4D
			GroundState | (EscDispatchAction << 4), // 4E
			GroundState | (EscDispatchAction << 4), // 4F
			GroundState | (EscDispatchAction << 4), // 50
			GroundState | (EscDispatchAction << 4), // 51
			GroundState | (EscDispatchAction << 4), // 52
			GroundState | (EscDispatchAction << 4), // 53
			GroundState | (EscDispatchAction << 4), // 54
			GroundState | (EscDispatchAction << 4), // 55
			GroundState | (EscDispatchAction << 4), // 56
			GroundState | (EscDispatchAction << 4), // 57
			GroundState | (EscDispatchAction << 4), // 58
			GroundState | (EscDispatchAction << 4), // 59
			GroundState | (EscDispatchAction << 4), // 5A
			GroundState | (EscDispatchAction << 4), // 5B
			GroundState | (EscDispatchAction << 4), // 5C
			GroundState | (EscDispatchAction << 4), // 5D
			GroundState | (EscDispatchAction << 4), // 5E
			GroundState | (EscDispatchAction << 4), // 5F
			GroundState | (EscDispatchAction << 4), // 60
			GroundState | (EscDispatchAction << 4), // 61
			GroundState | (EscDispatchAction << 4), // 62
			GroundState | (EscDispatchAction << 4), // 63
			GroundState | (EscDispatchAction << 4), // 64
			GroundState | (EscDispatchAction << 4), // 65
			GroundState | (EscDispatchAction << 4), // 66
			GroundState | (EscDispatchAction << 4), // 67
			GroundState | (EscDispatchAction << 4), // 68
			GroundState | (EscDispatchAction << 4), // 69
			GroundState | (EscDispatchAction << 4), // 6A
			GroundState | (EscDispatchAction << 4), // 6B
			GroundState | (EscDispatchAction << 4), // 6C
			GroundState | (EscDispatchAction << 4), // 6D
			GroundState | (EscDispatchAction << 4), // 6E
			GroundState | (EscDispatchAction << 4), // 6F
			GroundState | (EscDispatchAction << 4), // 70
			GroundState | (EscDispatchAction << 4), // 71
			GroundState | (EscDispatchAction << 4), // 72
			GroundState | (EscDispatchAction << 4), // 73
			GroundState | (EscDispatchAction << 4), // 74
			GroundState | (EscDispatchAction << 4), // 75
			GroundState | (EscDispatchAction << 4), // 76
			GroundState | (EscDispatchAction << 4), // 77
			GroundState | (EscDispatchAction << 4), // 78
			GroundState | (EscDispatchAction << 4), // 79
			GroundState | (EscDispatchAction << 4), // 7A
			GroundState | (EscDispatchAction << 4), // 7B
			GroundState | (EscDispatchAction << 4), // 7C
			GroundState | (EscDispatchAction << 4), // 7D
			GroundState | (EscDispatchAction << 4), // 7E
			AnywhereState | (IgnoreAction << 4),    // 7F
			0,                                      // 80
			0,                                      // 81
			0,                                      // 82
			0,                                      // 83
			0,                                      // 84
			0,                                      // 85
			0,                                      // 86
			0,                                      // 87
			0,                                      // 88
			0,                                      // 89
			0,                                      // 8A
			0,                                      // 8B
			0,                                      // 8C
			0,                                      // 8D
			0,                                      // 8E
			0,                                      // 8F
			0,                                      // 90
			0,                                      // 91
			0,                                      // 92
			0,                                      // 93
			0,                                      // 94
			0,                                      // 95
			0,                                      // 96
			0,                                      // 97
			0,                                      // 98
			0,                                      // 99
			0,                                      // 9A
			0,                                      // 9B
			0,                                      // 9C
			0,                                      // 9D
			0,                                      // 9E
			0,                                      // 9F
			0,                                      // A0
			0,                                      // A1
			0,                                      // A2
			0,                                      // A3
			0,                                      // A4
			0,                                      // A5
			0,                                      // A6
			0,                                      // A7
			0,                                      // A8
			0,                                      // A9
			0,                                      // AA
			0,                                      // AB
			0,                                      // AC
			0,                                      // AD
			0,                                      // AE
			0,                                      // AF
			0,                                      // B0
			0,                                      // B1
			0,                                      // B2
			0,                                      // B3
			0,                                      // B4
			0,                                      // B5
			0,                                      // B6
			0,                                      // B7
			0,                                      // B8
			0,                                      // B9
			0,                                      // BA
			0,                                      // BB
			0,                                      // BC
			0,                                      // BD
			0,                                      // BE
			0,                                      // BF
			0,                                      // C0
			0,                                      // C1
			0,                                      // C2
			0,                                      // C3
			0,                                      // C4
			0,                                      // C5
			0,                                      // C6
			0,                                      // C7
			0,                                      // C8
			0,                                      // C9
			0,                                      // CA
			0,                                      // CB
			0,                                      // CC
			0,                                      // CD
			0,                                      // CE
			0,                                      // CF
			0,                                      // D0
			0,                                      // D1
			0,                                      // D2
			0,                                      // D3
			0,                                      // D4
			0,                                      // D5
			0,                                      // D6
			0,                                      // D7
			0,                                      // D8
			0,                                      // D9
			0,                                      // DA
			0,                                      // DB
			0,                                      // DC
			0,                                      // DD
			0,                                      // DE
			0,                                      // DF
			0,                                      // E0
			0,                                      // E1
			0,                                      // E2
			0,                                      // E3
			0,                                      // E4
			0,                                      // E5
			0,                                      // E6
			0,                                      // E7
			0,                                      // E8
			0,                                      // E9
			0,                                      // EA
			0,                                      // EB
			0,                                      // EC
			0,                                      // ED
			0,                                      // EE
			0,                                      // EF
			0,                                      // F0
			0,                                      // F1
			0,                                      // F2
			0,                                      // F3
			0,                                      // F4
			0,                                      // F5
			0,                                      // F6
			0,                                      // F7
			0,                                      // F8
			0,                                      // F9
			0,                                      // FA
			0,                                      // FB
			0,                                      // FC
			0,                                      // FD
			0,                                      // FE
			0,                                      // FF
		},
		{ //State GroundState = 12
			AnywhereState | (ExecuteAction << 4), // 00
			AnywhereState | (ExecuteAction << 4), // 01
			AnywhereState | (ExecuteAction << 4), // 02
			AnywhereState | (ExecuteAction << 4), // 03
			AnywhereState | (ExecuteAction << 4), // 04
			AnywhereState | (ExecuteAction << 4), // 05
			AnywhereState | (ExecuteAction << 4), // 06
			AnywhereState | (ExecuteAction << 4), // 07
			AnywhereState | (ExecuteAction << 4), // 08
			AnywhereState | (ExecuteAction << 4), // 09
			AnywhereState | (ExecuteAction << 4), // 0A
			AnywhereState | (ExecuteAction << 4), // 0B
			AnywhereState | (ExecuteAction << 4), // 0C
			AnywhereState | (ExecuteAction << 4), // 0D
			AnywhereState | (ExecuteAction << 4), // 0E
			AnywhereState | (ExecuteAction << 4), // 0F
			AnywhereState | (ExecuteAction << 4), // 10
			AnywhereState | (ExecuteAction << 4), // 11
			AnywhereState | (ExecuteAction << 4), // 12
			AnywhereState | (ExecuteAction << 4), // 13
			AnywhereState | (ExecuteAction << 4), // 14
			AnywhereState | (ExecuteAction << 4), // 15
			AnywhereState | (ExecuteAction << 4), // 16
			AnywhereState | (ExecuteAction << 4), // 17
			0,                                    // 18
			AnywhereState | (ExecuteAction << 4), // 19
			0,                                    // 1A
			0,                                    // 1B
			AnywhereState | (ExecuteAction << 4), // 1C
			AnywhereState | (ExecuteAction << 4), // 1D
			AnywhereState | (ExecuteAction << 4), // 1E
			AnywhereState | (ExecuteAction << 4), // 1F
			AnywhereState | (PrintAction << 4),   // 20
			AnywhereState | (PrintAction << 4),   // 21
			AnywhereState | (PrintAction << 4),   // 22
			AnywhereState | (PrintAction << 4),   // 23
			AnywhereState | (PrintAction << 4),   // 24
			AnywhereState | (PrintAction << 4),   // 25
			AnywhereState | (PrintAction << 4),   // 26
			AnywhereState | (PrintAction << 4),   // 27
			AnywhereState | (PrintAction << 4),   // 28
			AnywhereState | (PrintAction << 4),   // 29
			AnywhereState | (PrintAction << 4),   // 2A
			AnywhereState | (PrintAction << 4),   // 2B
			AnywhereState | (PrintAction << 4),   // 2C
			AnywhereState | (PrintAction << 4),   // 2D
			AnywhereState | (PrintAction << 4),   // 2E
			AnywhereState | (PrintAction << 4),   // 2F
			AnywhereState | (PrintAction << 4),   // 30
			AnywhereState | (PrintAction << 4),   // 31
			AnywhereState | (PrintAction << 4),   // 32
			AnywhereState | (PrintAction << 4),   // 33
			AnywhereState | (PrintAction << 4),   // 34
			AnywhereState | (PrintAction << 4),   // 35
			AnywhereState | (PrintAction << 4),   // 36
			AnywhereState | (PrintAction << 4),   // 37
			AnywhereState | (PrintAction << 4),   // 38
			AnywhereState | (PrintAction << 4),   // 39
			AnywhereState | (PrintAction << 4),   // 3A
			AnywhereState | (PrintAction << 4),   // 3B
			AnywhereState | (PrintAction << 4),   // 3C
			AnywhereState | (PrintAction << 4),   // 3D
			AnywhereState | (PrintAction << 4),   // 3E
			AnywhereState | (PrintAction << 4),   // 3F
			AnywhereState | (PrintAction << 4),   // 40
			AnywhereState | (PrintAction << 4),   // 41
			AnywhereState | (PrintAction << 4),   // 42
			AnywhereState | (PrintAction << 4),   // 43
			AnywhereState | (PrintAction << 4),   // 44
			AnywhereState | (PrintAction << 4),   // 45
			AnywhereState | (PrintAction << 4),   // 46
			AnywhereState | (PrintAction << 4),   // 47
			AnywhereState | (PrintAction << 4),   // 48
			AnywhereState | (PrintAction << 4),   // 49
			AnywhereState | (PrintAction << 4),   // 4A
			AnywhereState | (PrintAction << 4),   // 4B
			AnywhereState | (PrintAction << 4),   // 4C
			AnywhereState | (PrintAction << 4),   // 4D
			AnywhereState | (PrintAction << 4),   // 4E
			AnywhereState | (PrintAction << 4),   // 4F
			AnywhereState | (PrintAction << 4),   // 50
			AnywhereState | (PrintAction << 4),   // 51
			AnywhereState | (PrintAction << 4),   // 52
			AnywhereState | (PrintAction << 4),   // 53
			AnywhereState | (PrintAction << 4),   // 54
			AnywhereState | (PrintAction << 4),   // 55
			AnywhereState | (PrintAction << 4),   // 56
			AnywhereState | (PrintAction << 4),   // 57
			AnywhereState | (PrintAction << 4),   // 58
			AnywhereState | (PrintAction << 4),   // 59
			AnywhereState | (PrintAction << 4),   // 5A
			AnywhereState | (PrintAction << 4),   // 5B
			AnywhereState | (PrintAction << 4),   // 5C
			AnywhereState | (PrintAction << 4),   // 5D
			AnywhereState | (PrintAction << 4),   // 5E
			AnywhereState | (PrintAction << 4),   // 5F
			AnywhereState | (PrintAction << 4),   // 60
			AnywhereState | (PrintAction << 4),   // 61
			AnywhereState | (PrintAction << 4),   // 62
			AnywhereState | (PrintAction << 4),   // 63
			AnywhereState | (PrintAction << 4),   // 64
			AnywhereState | (PrintAction << 4),   // 65
			AnywhereState | (PrintAction << 4),   // 66
			AnywhereState | (PrintAction << 4),   // 67
			AnywhereState | (PrintAction << 4),   // 68
			AnywhereState | (PrintAction << 4),   // 69
			AnywhereState | (PrintAction << 4),   // 6A
			AnywhereState | (PrintAction << 4),   // 6B
			AnywhereState | (PrintAction << 4),   // 6C
			AnywhereState | (PrintAction << 4),   // 6D
			AnywhereState | (PrintAction << 4),   // 6E
			AnywhereState | (PrintAction << 4),   // 6F
			AnywhereState | (PrintAction << 4),   // 70
			AnywhereState | (PrintAction << 4),   // 71
			AnywhereState | (PrintAction << 4),   // 72
			AnywhereState | (PrintAction << 4),   // 73
			AnywhereState | (PrintAction << 4),   // 74
			AnywhereState | (PrintAction << 4),   // 75
			AnywhereState | (PrintAction << 4),   // 76
			AnywhereState | (PrintAction << 4),   // 77
			AnywhereState | (PrintAction << 4),   // 78
			AnywhereState | (PrintAction << 4),   // 79
			AnywhereState | (PrintAction << 4),   // 7A
			AnywhereState | (PrintAction << 4),   // 7B
			AnywhereState | (PrintAction << 4),   // 7C
			AnywhereState | (PrintAction << 4),   // 7D
			AnywhereState | (PrintAction << 4),   // 7E
			AnywhereState | (PrintAction << 4),   // 7F
			AnywhereState | (ExecuteAction << 4), // 80
			AnywhereState | (ExecuteAction << 4), // 81
			AnywhereState | (ExecuteAction << 4), // 82
			AnywhereState | (ExecuteAction << 4), // 83
			AnywhereState | (ExecuteAction << 4), // 84
			AnywhereState | (ExecuteAction << 4), // 85
			AnywhereState | (ExecuteAction << 4), // 86
			AnywhereState | (ExecuteAction << 4), // 87
			AnywhereState | (ExecuteAction << 4), // 88
			AnywhereState | (ExecuteAction << 4), // 89
			AnywhereState | (ExecuteAction << 4), // 8A
			AnywhereState | (ExecuteAction << 4), // 8B
			AnywhereState | (ExecuteAction << 4), // 8C
			AnywhereState | (ExecuteAction << 4), // 8D
			AnywhereState | (ExecuteAction << 4), // 8E
			AnywhereState | (ExecuteAction << 4), // 8F
			0,                                    // 90
			AnywhereState | (ExecuteAction << 4), // 91
			AnywhereState | (ExecuteAction << 4), // 92
			AnywhereState | (ExecuteAction << 4), // 93
			AnywhereState | (ExecuteAction << 4), // 94
			AnywhereState | (ExecuteAction << 4), // 95
			AnywhereState | (ExecuteAction << 4), // 96
			AnywhereState | (ExecuteAction << 4), // 97
			AnywhereState | (ExecuteAction << 4), // 98
			AnywhereState | (ExecuteAction << 4), // 99
			AnywhereState | (ExecuteAction << 4), // 9A
			0,                                    // 9B
			AnywhereState | (ExecuteAction << 4), // 9C
			0,                                    // 9D
			0,                                    // 9E
			0,                                    // 9F
			0,                                    // A0
			0,                                    // A1
			0,                                    // A2
			0,                                    // A3
			0,                                    // A4
			0,                                    // A5
			0,                                    // A6
			0,                                    // A7
			0,                                    // A8
			0,                                    // A9
			0,                                    // AA
			0,                                    // AB
			0,                                    // AC
			0,                                    // AD
			0,                                    // AE
			0,                                    // AF
			0,                                    // B0
			0,                                    // B1
			0,                                    // B2
			0,                                    // B3
			0,                                    // B4
			0,                                    // B5
			0,                                    // B6
			0,                                    // B7
			0,                                    // B8
			0,                                    // B9
			0,                                    // BA
			0,                                    // BB
			0,                                    // BC
			0,                                    // BD
			0,                                    // BE
			0,                                    // BF
			0,                                    // C0
			0,                                    // C1
			Utf8State | (BeginUtf8Action << 4),   // C2
			Utf8State | (BeginUtf8Action << 4),   // C3
			Utf8State | (BeginUtf8Action << 4),   // C4
			Utf8State | (BeginUtf8Action << 4),   // C5
			Utf8State | (BeginUtf8Action << 4),   // C6
			Utf8State | (BeginUtf8Action << 4),   // C7
			Utf8State | (BeginUtf8Action << 4),   // C8
			Utf8State | (BeginUtf8Action << 4),   // C9
			Utf8State | (BeginUtf8Action << 4),   // CA
			Utf8State | (BeginUtf8Action << 4),   // CB
			Utf8State | (BeginUtf8Action << 4),   // CC
			Utf8State | (BeginUtf8Action << 4),   // CD
			Utf8State | (BeginUtf8Action << 4),   // CE
			Utf8State | (BeginUtf8Action << 4),   // CF
			Utf8State | (BeginUtf8Action << 4),   // D0
			Utf8State | (BeginUtf8Action << 4),   // D1
			Utf8State | (BeginUtf8Action << 4),   // D2
			Utf8State | (BeginUtf8Action << 4),   // D3
			Utf8State | (BeginUtf8Action << 4),   // D4
			Utf8State | (BeginUtf8Action << 4),   // D5
			Utf8State | (BeginUtf8Action << 4),   // D6
			Utf8State | (BeginUtf8Action << 4),   // D7
			Utf8State | (BeginUtf8Action << 4),   // D8
			Utf8State | (BeginUtf8Action << 4),   // D9
			Utf8State | (BeginUtf8Action << 4),   // DA
			Utf8State | (BeginUtf8Action << 4),   // DB
			Utf8State | (BeginUtf8Action << 4),   // DC
			Utf8State | (BeginUtf8Action << 4),   // DD
			Utf8State | (BeginUtf8Action << 4),   // DE
			Utf8State | (BeginUtf8Action << 4),   // DF
			Utf8State | (BeginUtf8Action << 4),   // E0
			Utf8State | (BeginUtf8Action << 4),   // E1
			Utf8State | (BeginUtf8Action << 4),   // E2
			Utf8State | (BeginUtf8Action << 4),   // E3
			Utf8State | (BeginUtf8Action << 4),   // E4
			Utf8State | (BeginUtf8Action << 4),   // E5
			Utf8State | (BeginUtf8Action << 4),   // E6
			Utf8State | (BeginUtf8Action << 4),   // E7
			Utf8State | (BeginUtf8Action << 4),   // E8
			Utf8State | (BeginUtf8Action << 4),   // E9
			Utf8State | (BeginUtf8Action << 4),   // EA
			Utf8State | (BeginUtf8Action << 4),   // EB
			Utf8State | (BeginUtf8Action << 4),   // EC
			Utf8State | (BeginUtf8Action << 4),   // ED
			Utf8State | (BeginUtf8Action << 4),   // EE
			Utf8State | (BeginUtf8Action << 4),   // EF
			Utf8State | (BeginUtf8Action << 4),   // F0
			Utf8State | (BeginUtf8Action << 4),   // F1
			Utf8State | (BeginUtf8Action << 4),   // F2
			Utf8State | (BeginUtf8Action << 4),   // F3
			Utf8State | (BeginUtf8Action << 4),   // F4
			0,                                    // F5
			0,                                    // F6
			0,                                    // F7
			0,                                    // F8
			0,                                    // F9
			0,                                    // FA
			0,                                    // FB
			0,                                    // FC
			0,                                    // FD
			0,                                    // FE
			0,                                    // FF
		},
		{ //State OscStringState = 13
			AnywhereState | (IgnoreAction << 4), // 00
			AnywhereState | (IgnoreAction << 4), // 01
			AnywhereState | (IgnoreAction << 4), // 02
			AnywhereState | (IgnoreAction << 4), // 03
			AnywhereState | (IgnoreAction << 4), // 04
			AnywhereState | (IgnoreAction << 4), // 05
			AnywhereState | (IgnoreAction << 4), // 06
			GroundState | (NoneAction << 4),     // 07
			AnywhereState | (IgnoreAction << 4), // 08
			AnywhereState | (IgnoreAction << 4), // 09
			AnywhereState | (IgnoreAction << 4), // 0A
			AnywhereState | (IgnoreAction << 4), // 0B
			AnywhereState | (IgnoreAction << 4), // 0C
			AnywhereState | (IgnoreAction << 4), // 0D
			AnywhereState | (IgnoreAction << 4), // 0E
			AnywhereState | (IgnoreAction << 4), // 0F
			AnywhereState | (IgnoreAction << 4), // 10
			AnywhereState | (IgnoreAction << 4), // 11
			AnywhereState | (IgnoreAction << 4), // 12
			AnywhereState | (IgnoreAction << 4), // 13
			AnywhereState | (IgnoreAction << 4), // 14
			AnywhereState | (IgnoreAction << 4), // 15
			AnywhereState | (IgnoreAction << 4), // 16
			AnywhereState | (IgnoreAction << 4), // 17
			0,                                   // 18
			AnywhereState | (IgnoreAction << 4), // 19
			0,                                   // 1A
			0,                                   // 1B
			AnywhereState | (IgnoreAction << 4), // 1C
			AnywhereState | (IgnoreAction << 4), // 1D
			AnywhereState | (IgnoreAction << 4), // 1E
			AnywhereState | (IgnoreAction << 4), // 1F
			AnywhereState | (OscPutAction << 4), // 20
			AnywhereState | (OscPutAction << 4), // 21
			AnywhereState | (OscPutAction << 4), // 22
			AnywhereState | (OscPutAction << 4), // 23
			AnywhereState | (OscPutAction << 4), // 24
			AnywhereState | (OscPutAction << 4), // 25
			AnywhereState | (OscPutAction << 4), // 26
			AnywhereState | (OscPutAction << 4), // 27
			AnywhereState | (OscPutAction << 4), // 28
			AnywhereState | (OscPutAction << 4), // 29
			AnywhereState | (OscPutAction << 4), // 2A
			AnywhereState | (OscPutAction << 4), // 2B
			AnywhereState | (OscPutAction << 4), // 2C
			AnywhereState | (OscPutAction << 4), // 2D
			AnywhereState | (OscPutAction << 4), // 2E
			AnywhereState | (OscPutAction << 4), // 2F
			AnywhereState | (OscPutAction << 4), // 30
			AnywhereState | (OscPutAction << 4), // 31
			AnywhereState | (OscPutAction << 4), // 32
			AnywhereState | (OscPutAction << 4), // 33
			AnywhereState | (OscPutAction << 4), // 34
			AnywhereState | (OscPutAction << 4), // 35
			AnywhereState | (OscPutAction << 4), // 36
			AnywhereState | (OscPutAction << 4), // 37
			AnywhereState | (OscPutAction << 4), // 38
			AnywhereState | (OscPutAction << 4), // 39
			AnywhereState | (OscPutAction << 4), // 3A
			AnywhereState | (OscPutAction << 4), // 3B
			AnywhereState | (OscPutAction << 4), // 3C
			AnywhereState | (OscPutAction << 4), // 3D
			AnywhereState | (OscPutAction << 4), // 3E
			AnywhereState | (OscPutAction << 4), // 3F
			AnywhereState | (OscPutAction << 4), // 40
			AnywhereState | (OscPutAction << 4), // 41
			AnywhereState | (OscPutAction << 4), // 42
			AnywhereState | (OscPutAction << 4), // 43
			AnywhereState | (OscPutAction << 4), // 44
			AnywhereState | (OscPutAction << 4), // 45
			AnywhereState | (OscPutAction << 4), // 46
			AnywhereState | (OscPutAction << 4), // 47
			AnywhereState | (OscPutAction << 4), // 48
			AnywhereState | (OscPutAction << 4), // 49
			AnywhereState | (OscPutAction << 4), // 4A
			AnywhereState | (OscPutAction << 4), // 4B
			AnywhereState | (OscPutAction << 4), // 4C
			AnywhereState | (OscPutAction << 4), // 4D
			AnywhereState | (OscPutAction << 4), // 4E
			AnywhereState | (OscPutAction << 4), // 4F
			AnywhereState | (OscPutAction << 4), // 50
			AnywhereState | (OscPutAction << 4), // 51
			AnywhereState | (OscPutAction << 4), // 52
			AnywhereState | (OscPutAction << 4), // 53
			AnywhereState | (OscPutAction << 4), // 54
			AnywhereState | (OscPutAction << 4), // 55
			AnywhereState | (OscPutAction << 4), // 56
			AnywhereState | (OscPutAction << 4), // 57
			AnywhereState | (OscPutAction << 4), // 58
			AnywhereState | (OscPutAction << 4), // 59
			AnywhereState | (OscPutAction << 4), // 5A
			AnywhereState | (OscPutAction << 4), // 5B
			AnywhereState | (OscPutAction << 4), // 5C
			AnywhereState | (OscPutAction << 4), // 5D
			AnywhereState | (OscPutAction << 4), // 5E
			AnywhereState | (OscPutAction << 4), // 5F
			AnywhereState | (OscPutAction << 4), // 60
			AnywhereState | (OscPutAction << 4), // 61
			AnywhereState | (OscPutAction << 4), // 62
			AnywhereState | (OscPutAction << 4), // 63
			AnywhereState | (OscPutAction << 4), // 64
			AnywhereState | (OscPutAction << 4), // 65
			AnywhereState | (OscPutAction << 4), // 66
			AnywhereState | (OscPutAction << 4), // 67
			AnywhereState | (OscPutAction << 4), // 68
			AnywhereState | (OscPutAction << 4), // 69
			AnywhereState | (OscPutAction << 4), // 6A
			AnywhereState | (OscPutAction << 4), // 6B
			AnywhereState | (OscPutAction << 4), // 6C
			AnywhereState | (OscPutAction << 4), // 6D
			AnywhereState | (OscPutAction << 4), // 6E
			AnywhereState | (OscPutAction << 4), // 6F
			AnywhereState | (OscPutAction << 4), // 70
			AnywhereState | (OscPutAction << 4), // 71
			AnywhereState | (OscPutAction << 4), // 72
			AnywhereState | (OscPutAction << 4), // 73
			AnywhereState | (OscPutAction << 4), // 74
			AnywhereState | (OscPutAction << 4), // 75
			AnywhereState | (OscPutAction << 4), // 76
			AnywhereState | (OscPutAction << 4), // 77
			AnywhereState | (OscPutAction << 4), // 78
			AnywhereState | (OscPutAction << 4), // 79
			AnywhereState | (OscPutAction << 4), // 7A
			AnywhereState | (OscPutAction << 4), // 7B
			AnywhereState | (OscPutAction << 4), // 7C
			AnywhereState | (OscPutAction << 4), // 7D
			AnywhereState | (OscPutAction << 4), // 7E
			AnywhereState | (OscPutAction << 4), // 7F
			AnywhereState | (OscPutAction << 4), // 80
			AnywhereState | (OscPutAction << 4), // 81
			AnywhereState | (OscPutAction << 4), // 82
			AnywhereState | (OscPutAction << 4), // 83
			AnywhereState | (OscPutAction << 4), // 84
			AnywhereState | (OscPutAction << 4), // 85
			AnywhereState | (OscPutAction << 4), // 86
			AnywhereState | (OscPutAction << 4), // 87
			AnywhereState | (OscPutAction << 4), // 88
			AnywhereState | (OscPutAction << 4), // 89
			AnywhereState | (OscPutAction << 4), // 8A
			AnywhereState | (OscPutAction << 4), // 8B
			AnywhereState | (OscPutAction << 4), // 8C
			AnywhereState | (OscPutAction << 4), // 8D
			AnywhereState | (OscPutAction << 4), // 8E
			AnywhereState | (OscPutAction << 4), // 8F
			AnywhereState | (OscPutAction << 4), // 90
			AnywhereState | (OscPutAction << 4), // 91
			AnywhereState | (OscPutAction << 4), // 92
			AnywhereState | (OscPutAction << 4), // 93
			AnywhereState | (OscPutAction << 4), // 94
			AnywhereState | (OscPutAction << 4), // 95
			AnywhereState | (OscPutAction << 4), // 96
			AnywhereState | (OscPutAction << 4), // 97
			AnywhereState | (OscPutAction << 4), // 98
			AnywhereState | (OscPutAction << 4), // 99
			AnywhereState | (OscPutAction << 4), // 9A
			AnywhereState | (OscPutAction << 4), // 9B
			AnywhereState | (OscPutAction << 4), // 9C
			AnywhereState | (OscPutAction << 4), // 9D
			AnywhereState | (OscPutAction << 4), // 9E
			AnywhereState | (OscPutAction << 4), // 9F
			AnywhereState | (OscPutAction << 4), // A0
			AnywhereState | (OscPutAction << 4), // A1
			AnywhereState | (OscPutAction << 4), // A2
			AnywhereState | (OscPutAction << 4), // A3
			AnywhereState | (OscPutAction << 4), // A4
			AnywhereState | (OscPutAction << 4), // A5
			AnywhereState | (OscPutAction << 4), // A6
			AnywhereState | (OscPutAction << 4), // A7
			AnywhereState | (OscPutAction << 4), // A8
			AnywhereState | (OscPutAction << 4), // A9
			AnywhereState | (OscPutAction << 4), // AA
			AnywhereState | (OscPutAction << 4), // AB
			AnywhereState | (OscPutAction << 4), // AC
			AnywhereState | (OscPutAction << 4), // AD
			AnywhereState | (OscPutAction << 4), // AE
			AnywhereState | (OscPutAction << 4), // AF
			AnywhereState | (OscPutAction << 4), // B0
			AnywhereState | (OscPutAction << 4), // B1
			AnywhereState | (OscPutAction << 4), // B2
			AnywhereState | (OscPutAction << 4), // B3
			AnywhereState | (OscPutAction << 4), // B4
			AnywhereState | (OscPutAction << 4), // B5
			AnywhereState | (OscPutAction << 4), // B6
			AnywhereState | (OscPutAction << 4), // B7
			AnywhereState | (OscPutAction << 4), // B8
			AnywhereState | (OscPutAction << 4), // B9
			AnywhereState | (OscPutAction << 4), // BA
			AnywhereState | (OscPutAction << 4), // BB
			AnywhereState | (OscPutAction << 4), // BC
			AnywhereState | (OscPutAction << 4), // BD
			AnywhereState | (OscPutAction << 4), // BE
			AnywhereState | (OscPutAction << 4), // BF
			AnywhereState | (OscPutAction << 4), // C0
			AnywhereState | (OscPutAction << 4), // C1
			AnywhereState | (OscPutAction << 4), // C2
			AnywhereState | (OscPutAction << 4), // C3
			AnywhereState | (OscPutAction << 4), // C4
			AnywhereState | (OscPutAction << 4), // C5
			AnywhereState | (OscPutAction << 4), // C6
			AnywhereState | (OscPutAction << 4), // C7
			AnywhereState | (OscPutAction << 4), // C8
			AnywhereState | (OscPutAction << 4), // C9
			AnywhereState | (OscPutAction << 4), // CA
			AnywhereState | (OscPutAction << 4), // CB
			AnywhereState | (OscPutAction << 4), // CC
			AnywhereState | (OscPutAction << 4), // CD
			AnywhereState | (OscPutAction << 4), // CE
			AnywhereState | (OscPutAction << 4), // CF
			AnywhereState | (OscPutAction << 4), // D0
			AnywhereState | (OscPutAction << 4), // D1
			AnywhereState | (OscPutAction << 4), // D2
			AnywhereState | (OscPutAction << 4), // D3
			AnywhereState | (OscPutAction << 4), // D4
			AnywhereState | (OscPutAction << 4), // D5
			AnywhereState | (OscPutAction << 4), // D6
			AnywhereState | (OscPutAction << 4), // D7
			AnywhereState | (OscPutAction << 4), // D8
			AnywhereState | (OscPutAction << 4), // D9
			AnywhereState | (OscPutAction << 4), // DA
			AnywhereState | (OscPutAction << 4), // DB
			AnywhereState | (OscPutAction << 4), // DC
			AnywhereState | (OscPutAction << 4), // DD
			AnywhereState | (OscPutAction << 4), // DE
			AnywhereState | (OscPutAction << 4), // DF
			AnywhereState | (OscPutAction << 4), // E0
			AnywhereState | (OscPutAction << 4), // E1
			AnywhereState | (OscPutAction << 4), // E2
			AnywhereState | (OscPutAction << 4), // E3
			AnywhereState | (OscPutAction << 4), // E4
			AnywhereState | (OscPutAction << 4), // E5
			AnywhereState | (OscPutAction << 4), // E6
			AnywhereState | (OscPutAction << 4), // E7
			AnywhereState | (OscPutAction << 4), // E8
			AnywhereState | (OscPutAction << 4), // E9
			AnywhereState | (OscPutAction << 4), // EA
			AnywhereState | (OscPutAction << 4), // EB
			AnywhereState | (OscPutAction << 4), // EC
			AnywhereState | (OscPutAction << 4), // ED
			AnywhereState | (OscPutAction << 4), // EE
			AnywhereState | (OscPutAction << 4), // EF
			AnywhereState | (OscPutAction << 4), // F0
			AnywhereState | (OscPutAction << 4), // F1
			AnywhereState | (OscPutAction << 4), // F2
			AnywhereState | (OscPutAction << 4), // F3
			AnywhereState | (OscPutAction << 4), // F4
			AnywhereState | (OscPutAction << 4), // F5
			AnywhereState | (OscPutAction << 4), // F6
			AnywhereState | (OscPutAction << 4), // F7
			AnywhereState | (OscPutAction << 4), // F8
			AnywhereState | (OscPutAction << 4), // F9
			AnywhereState | (OscPutAction << 4), // FA
			AnywhereState | (OscPutAction << 4), // FB
			AnywhereState | (OscPutAction << 4), // FC
			AnywhereState | (OscPutAction << 4), // FD
			AnywhereState | (OscPutAction << 4), // FE
			AnywhereState | (OscPutAction << 4), // FF
		},
		{ //State SosPmApcStringState = 14
			AnywhereState | (IgnoreAction << 4), // 00
			AnywhereState | (IgnoreAction << 4), // 01
			AnywhereState | (IgnoreAction << 4), // 02
			AnywhereState | (IgnoreAction << 4), // 03
			AnywhereState | (IgnoreAction << 4), // 04
			AnywhereState | (IgnoreAction << 4), // 05
			AnywhereState | (IgnoreAction << 4), // 06
			AnywhereState | (IgnoreAction << 4), // 07
			AnywhereState | (IgnoreAction << 4), // 08
			AnywhereState | (IgnoreAction << 4), // 09
			AnywhereState | (IgnoreAction << 4), // 0A
			AnywhereState | (IgnoreAction << 4), // 0B
			AnywhereState | (IgnoreAction << 4), // 0C
			AnywhereState | (IgnoreAction << 4), // 0D
			AnywhereState | (IgnoreAction << 4), // 0E
			AnywhereState | (IgnoreAction << 4), // 0F
			AnywhereState | (IgnoreAction << 4), // 10
			AnywhereState | (IgnoreAction << 4), // 11
			AnywhereState | (IgnoreAction << 4), // 12
			AnywhereState | (IgnoreAction << 4), // 13
			AnywhereState | (IgnoreAction << 4), // 14
			AnywhereState | (IgnoreAction << 4), // 15
			AnywhereState | (IgnoreAction << 4), // 16
			AnywhereState | (IgnoreAction << 4), // 17
			0,                                   // 18
			AnywhereState | (IgnoreAction << 4), // 19
			0,                                   // 1A
			0,                                   // 1B
			AnywhereState | (IgnoreAction << 4), // 1C
			AnywhereState | (IgnoreAction << 4), // 1D
			AnywhereState | (IgnoreAction << 4), // 1E
			AnywhereState | (IgnoreAction << 4), // 1F
			AnywhereState | (IgnoreAction << 4), // 20
			AnywhereState | (IgnoreAction << 4), // 21
			AnywhereState | (IgnoreAction << 4), // 22
			AnywhereState | (IgnoreAction << 4), // 23
			AnywhereState | (IgnoreAction << 4), // 24
			AnywhereState | (IgnoreAction << 4), // 25
			AnywhereState | (IgnoreAction << 4), // 26
			AnywhereState | (IgnoreAction << 4), // 27
			AnywhereState | (IgnoreAction << 4), // 28
			AnywhereState | (IgnoreAction << 4), // 29
			AnywhereState | (IgnoreAction << 4), // 2A
			AnywhereState | (IgnoreAction << 4), // 2B
			AnywhereState | (IgnoreAction << 4), // 2C
			AnywhereState | (IgnoreAction << 4), // 2D
			AnywhereState | (IgnoreAction << 4), // 2E
			AnywhereState | (IgnoreAction << 4), // 2F
			AnywhereState | (IgnoreAction << 4), // 30
			AnywhereState | (IgnoreAction << 4), // 31
			AnywhereState | (IgnoreAction << 4), // 32
			AnywhereState | (IgnoreAction << 4), // 33
			AnywhereState | (IgnoreAction << 4), // 34
			AnywhereState | (IgnoreAction << 4), // 35
			AnywhereState | (IgnoreAction << 4), // 36
			AnywhereState | (IgnoreAction << 4), // 37
			AnywhereState | (IgnoreAction << 4), // 38
			AnywhereState | (IgnoreAction << 4), // 39
			AnywhereState | (IgnoreAction << 4), // 3A
			AnywhereState | (IgnoreAction << 4), // 3B
			AnywhereState | (IgnoreAction << 4), // 3C
			AnywhereState | (IgnoreAction << 4), // 3D
			AnywhereState | (IgnoreAction << 4), // 3E
			AnywhereState | (IgnoreAction << 4), // 3F
			AnywhereState | (IgnoreAction << 4), // 40
			AnywhereState | (IgnoreAction << 4), // 41
			AnywhereState | (IgnoreAction << 4), // 42
			AnywhereState | (IgnoreAction << 4), // 43
			AnywhereState | (IgnoreAction << 4), // 44
			AnywhereState | (IgnoreAction << 4), // 45
			AnywhereState | (IgnoreAction << 4), // 46
			AnywhereState | (IgnoreAction << 4), // 47
			AnywhereState | (IgnoreAction << 4), // 48
			AnywhereState | (IgnoreAction << 4), // 49
			AnywhereState | (IgnoreAction << 4), // 4A
			AnywhereState | (IgnoreAction << 4), // 4B
			AnywhereState | (IgnoreAction << 4), // 4C
			AnywhereState | (IgnoreAction << 4), // 4D
			AnywhereState | (IgnoreAction << 4), // 4E
			AnywhereState | (IgnoreAction << 4), // 4F
			AnywhereState | (IgnoreAction << 4), // 50
			AnywhereState | (IgnoreAction << 4), // 51
			AnywhereState | (IgnoreAction << 4), // 52
			AnywhereState | (IgnoreAction << 4), // 53
			AnywhereState | (IgnoreAction << 4), // 54
			AnywhereState | (IgnoreAction << 4), // 55
			AnywhereState | (IgnoreAction << 4), // 56
			AnywhereState | (IgnoreAction << 4), // 57
			AnywhereState | (IgnoreAction << 4), // 58
			AnywhereState | (IgnoreAction << 4), // 59
			AnywhereState | (IgnoreAction << 4), // 5A
			AnywhereState | (IgnoreAction << 4), // 5B
			AnywhereState | (IgnoreAction << 4), // 5C
			AnywhereState | (IgnoreAction << 4), // 5D
			AnywhereState | (IgnoreAction << 4), // 5E
			AnywhereState | (IgnoreAction << 4), // 5F
			AnywhereState | (IgnoreAction << 4), // 60
			AnywhereState | (IgnoreAction << 4), // 61
			AnywhereState | (IgnoreAction << 4), // 62
			AnywhereState | (IgnoreAction << 4), // 63
			AnywhereState | (IgnoreAction << 4), // 64
			AnywhereState | (IgnoreAction << 4), // 65
			AnywhereState | (IgnoreAction << 4), // 66
			AnywhereState | (IgnoreAction << 4), // 67
			AnywhereState | (IgnoreAction << 4), // 68
			AnywhereState | (IgnoreAction << 4), // 69
			AnywhereState | (IgnoreAction << 4), // 6A
			AnywhereState | (IgnoreAction << 4), // 6B
			AnywhereState | (IgnoreAction << 4), // 6C
			AnywhereState | (IgnoreAction << 4), // 6D
			AnywhereState | (IgnoreAction << 4), // 6E
			AnywhereState | (IgnoreAction << 4), // 6F
			AnywhereState | (IgnoreAction << 4), // 70
			AnywhereState | (IgnoreAction << 4), // 71
			AnywhereState | (IgnoreAction << 4), // 72
			AnywhereState | (IgnoreAction << 4), // 73
			AnywhereState | (IgnoreAction << 4), // 74
			AnywhereState | (IgnoreAction << 4), // 75
			AnywhereState | (IgnoreAction << 4), // 76
			AnywhereState | (IgnoreAction << 4), // 77
			AnywhereState | (IgnoreAction << 4), // 78
			AnywhereState | (IgnoreAction << 4), // 79
			AnywhereState | (IgnoreAction << 4), // 7A
			AnywhereState | (IgnoreAction << 4), // 7B
			AnywhereState | (IgnoreAction << 4), // 7C
			AnywhereState | (IgnoreAction << 4), // 7D
			AnywhereState | (IgnoreAction << 4), // 7E
			AnywhereState | (IgnoreAction << 4), // 7F
			0,                                   // 80
			0,                                   // 81
			0,                                   // 82
			0,                                   // 83
			0,                                   // 84
			0,                                   // 85
			0,                                   // 86
			0,                                   // 87
			0,                                   // 88
			0,                                   // 89
			0,                                   // 8A
			0,                                   // 8B
			0,                                   // 8C
			0,                                   // 8D
			0,                                   // 8E
			0,                                   // 8F
			0,                                   // 90
			0,                                   // 91
			0,                                   // 92
			0,                                   // 93
			0,                                   // 94
			0,                                   // 95
			0,                                   // 96
			0,                                   // 97
			0,                                   // 98
			0,                                   // 99
			0,                                   // 9A
			0,                                   // 9B
			GroundState | (NoneAction << 4),     // 9C
			0,                                   // 9D
			0,                                   // 9E
			0,                                   // 9F
			0,                                   // A0
			0,                                   // A1
			0,                                   // A2
			0,                                   // A3
			0,                                   // A4
			0,                                   // A5
			0,                                   // A6
			0,                                   // A7
			0,                                   // A8
			0,                                   // A9
			0,                                   // AA
			0,                                   // AB
			0,                                   // AC
			0,                                   // AD
			0,                                   // AE
			0,                                   // AF
			0,                                   // B0
			0,                                   // B1
			0,                                   // B2
			0,                                   // B3
			0,                                   // B4
			0,                                   // B5
			0,                                   // B6
			0,                                   // B7
			0,                                   // B8
			0,                                   // B9
			0,                                   // BA
			0,                                   // BB
			0,                                   // BC
			0,                                   // BD
			0,                                   // BE
			0,                                   // BF
			0,                                   // C0
			0,                                   // C1
			0,                                   // C2
			0,                                   // C3
			0,                                   // C4
			0,                                   // C5
			0,                                   // C6
			0,                                   // C7
			0,                                   // C8
			0,                                   // C9
			0,                                   // CA
			0,                                   // CB
			0,                                   // CC
			0,                                   // CD
			0,                                   // CE
			0,                                   // CF
			0,                                   // D0
			0,                                   // D1
			0,                                   // D2
			0,                                   // D3
			0,                                   // D4
			0,                                   // D5
			0,                                   // D6
			0,                                   // D7
			0,                                   // D8
			0,                                   // D9
			0,                                   // DA
			0,                                   // DB
			0,                                   // DC
			0,                                   // DD
			0,                                   // DE
			0,                                   // DF
			0,                                   // E0
			0,                                   // E1
			0,                                   // E2
			0,                                   // E3
			0,                                   // E4
			0,                                   // E5
			0,                                   // E6
			0,                                   // E7
			0,                                   // E8
			0,                                   // E9
			0,                                   // EA
			0,                                   // EB
			0,                                   // EC
			0,                                   // ED
			0,                                   // EE
			0,                                   // EF
			0,                                   // F0
			0,                                   // F1
			0,                                   // F2
			0,                                   // F3
			0,                                   // F4
			0,                                   // F5
			0,                                   // F6
			0,                                   // F7
			0,                                   // F8
			0,                                   // F9
			0,                                   // FA
			0,                                   // FB
			0,                                   // FC
			0,                                   // FD
			0,                                   // FE
			0,                                   // FF
		},
	}

	entryActions = []byte{
		0,              //AnywhereState
		ClearAction,    //CsiEntryState
		0,              //CsiIgnoreState
		0,              //CsiIntermediateState
		0,              //CsiParamState
		ClearAction,    //DcsEntryState
		0,              //DcsIgnoreState
		0,              //DcsIntermediateState
		0,              //DcsParamState
		HookAction,     //DcsPassthroughState
		ClearAction,    //EscapeState
		0,              //EscapeIntermediateState
		0,              //GroundState
		OscStartAction, //OscStringState
		0,              //SosPmApcStringState
		0,              //Utf8State
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
