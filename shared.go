package gouse

type ITest struct {
	Name    string
	Input   string
	Want    bool
	WantErr bool
}

// These maybe not supported by all terminals
const (
	DEFAULT_FG string = "\033[0m"
	RED_FG     string = "\033[31m"
	GREEN_FG   string = "\033[32m"
	YELLOW_FG  string = "\033[33m"
	PURPLE_FG  string = "\033[34m"
	PINK_FG    string = "\033[35m"
	CYAN_FG    string = "\033[36m"
	WHITE_FG   string = "\033[97m"
	ORANGE_FG  string = "\033[38;5;208m"
	BLUE_FG    string = "\033[38;5;27m"
	MAGENTA_FG string = "\033[38;5;13m"
	Gray       string = "\033[37m"
)

// const (
// 	EmptyReg   = `^$`
// 	SpaceReg   = `\s+`
// 	WordNumReg = `^[a-zA-Z0-9]+$`
// )

const (
	UsernameReg = `^[a-zA-Z0-9_]{3,20}$`
)

const (
	PasswordLenReg     = `^.{8,32}$`
	PasswordLowerReg   = `[a-z]`
	PasswordUpperReg   = `[A-Z]`
	PasswordDigitReg   = `\d`
	PasswordSpecialReg = `[!@#$%^&*]`
)

const (
	EmailLenReg = `^.{8,32}$`
)

const (
	PhoneReg = `^\+\d{1,2}\s?\(\d{1,4}\)\s?\d{1,6}-\d{1,6}$`
)

type IFontBannerType map[string][3]string

var DOUBLE_ALPHA = IFontBannerType{
	`A`: {
		`╔═╗`,
		`╠═╣`,
		`╩ ╩`,
	},
	`B`: {
		`╔╗ `,
		`╠╩╗`,
		`╚═╝`,
	},
	`C`: {
		`╔═╗`,
		`║  `,
		`╚═╝`,
	},
	`D`: {
		`╦═╗`,
		`║ ║`,
		`╩═╝`,
	},
	`E`: {
		`╔═╗`,
		`║╣ `,
		`╚═╝`,
	},
	`F`: {
		`╔═╗`,
		`║╣ `,
		`╩  `,
	},
	`G`: {
		`╔═╗`,
		`║ ╦`,
		`╚═╝`,
	},
	`H`: {
		`╦ ╦`,
		`╠═╣`,
		`╩ ╩`,
	},
	`I`: {
		` ╦ `,
		` ║ `,
		` ╩ `,
	},
	`J`: {
		`╦═╦`,
		`  ║`,
		` ═╝`,
	},
	`K`: {
		`╦  ╦`,
		`╠═╔╝`,
		`╝ ╚`,
	},
	`L`: {
		`╦  `,
		`║  `,
		`╩═╝`,
	},
	`M`: {
		`╔╗╔╗`,
		`║║║║`,
		`╝╚╝╚`,
	},
	`N`: {
		`╔╗╔`,
		`║║║`,
		`╝╚╝`,
	},
	`O`: {
		`╔═╗`,
		`║ ║`,
		`╚═╝`,
	},
	`P`: {
		`╔═╗`,
		`╠═╝`,
		`╩  `,
	},
	`Q`: {
		`╔═╗`,
		`║ ║`,
		`╚═╚═`,
	},
	`R`: {
		`╔═╗`,
		`╠╔╝`,
		`╝╚╝`,
	},
	`S`: {
		`╔═╗`,
		`╚═╗`,
		`╚═╝`,
	},
	`T`: {
		`╔╦╗`,
		` ║ `,
		` ╩ `,
	},
	`U`: {
		`╦ ╦`,
		`║ ║`,
		`╚═╝`,
	},
	`V`: {
		`╦   ╦`,
		`╚╗ ╔╝`,
		` ╚╔╝`,
	},
	`W`: {
		`╦ ╦`,
		`║║║`,
		`╚╩╝`,
	},
	`X`: {
		`╚╔╝`,
		`╔╝╗`,
	},
	`Y`: {
		`╦ ╦`,
		`╚╔╝`,
		` ╩ `,
	},
	`Z`: {
		`╔═╗`,
		`╔═╝`,
		`╚═╝`,
	},
	`0`: {
		`╔═╗`,
		`║ ║`,
		`╚═╝`,
	},
	`1`: {
		` ╦ `,
		` ║ `,
		` ╩ `,
	},
	`2`: {
		`╔═╗`,
		`╔═╝`,
		`╚══`,
	},
	`3`: {
		`╔═╗`,
		`╚═╗`,
		`╚═╝`,
	},
	`4`: {
		`╦ ╦`,
		`╚═╣`,
		`  ╩`,
	},
	`5`: {
		`╔═╗`,
		`╚═╗`,
		`╚═╝`,
	},
	`6`: {
		`╔═╗`,
		`╠═╗`,
		`╚═╝`,
	},
	`7`: {
		`╔═╗`,
		` ═║`,
		`  ╩`,
	},
	`8`: {
		`╔═╗`,
		`╠═╣`,
		`╚═╝`,
	},
	`9`: {
		`╔═╗`,
		`╚═╣`,
		`╚═╝`,
	},
}

// var zero =
//   `███
//    █ █
//    █ █
//    █ █
//    ███`

// var one =
//   `██
//     █
//     █
//     █
//    ███`

// var two =
//   `███
//      █
//    ███
//    █
//    ███`

// var three =
//   `███
//      █
//    ███
//      █
//    ███`

// var four =
//   `█ █
//    █ █
//    ███
//      █
//      █`

// var five =
//   `███
//    █
//    ███
//      █
//    ███`

// var six =
//   `███
//    █
//    ███
//    █ █
//    ███`

// var seven =
// `███
//     █
//     █
//     █
//     █`

// var eight =
//   `███
//    █ █
//    ███
//    █ █
//    ███`

// var nine =
//   `███
//    █ █
//    ███
//      █
//    ███`

// var colon =
//   `
//     ░

//     ░
//       `

const (
	// basic symbols
	SCOPE_SYM     string = " ⦿ "
	INPUT_SYM     string = " » "
	AGAIN_SYM     string = " ↺ "
	EXIT_SYM      string = " ↵ "
	INFO_SYM      string = " ⚠ "
	QUESTION_SYM  string = " ? "
	SPAN_HOR_SYM  string = " ↔ "
	SPAN_VER_SYM  string = " ↕ "
	MORE_SYM      string = " ... "
	OK_SYM        string = " ✔ "
	NO_SYM        string = " ✘ "
	CROSS_SYM     string = " ✖ "
	ITEM_SYM      string = " ➤ "
	COPYRIGHT_SYM string = " © "
	TRADEMARK_SYM string = " ™ "
	REGISTER_SYM  string = " ® "
	AT_SYM        string = " @ "
	NUMERO_SYM    string = " № "
)

const (
	// arrow symbols
	TO_RIGHT_CON string = " → "
	TO_LEFT_SYM  string = " ← "
	TO_UP_SYM    string = " ↑ "
	TO_DOWN_SYM  string = " ↓ "
	ARROW3_SYM   string = " ⇨ "
	ARROW4_SYM   string = " ⇾ "
	ARROW5_SYM   string = " ⇛ "
	ARROW6_SYM   string = " ⇝ "
	ARROW7_SYM   string = " ⇢ "
	ARROW8_SYM   string = " ⇥ "
	LEFT_SYM     string = " ◀ "
	RIGHT_SYM    string = " ▶ "
	UP_SYM       string = " ▲ "
	DOWN_SYM     string = " ▼ "
)

const (
	// math symbols
	PLUS_SYM           string = " + "
	MINUS_SYM          string = " - "
	MULT_SYM           string = " * "
	DIV_SYM            string = " / "
	MOD_SYM            string = " % "
	OR_SYM             string = " | "
	AND_SYM            string = " & "
	NOT_EQUAL_SYM2     string = " != "
	LESS_SYM           string = " < "
	GREATER_SYM        string = " > "
	LESS_EQUAL_SYM2    string = " <= "
	GREATER_EQUAL_SYM2 string = " >= "
	APPROX_EQUAL_SYM2  string = " ~= "
	ASSIGN_SYM         string = " = "
	PLUS_ASSIGN_SYM    string = " += "
	MINUS_ASSIGN_SYM   string = " -= "
	MULT_ASSIGN_SYM    string = " *= "
	DIV_ASSIGN_SYM     string = " /= "
	SECTION_SYM        string = " § "
	PI_SYM             string = " π "
	INFINITY_SYM       string = " ∞ "
	NOT_EQUAL_SYM      string = " ≠ "
	LESS_EQUAL_SYM     string = " ≤ "
	GREATER_EQUAL_SYM  string = " ≥ "
	APPROX_EQUAL_SYM   string = " ≈ "
	PLUSMINUS_SYM      string = " ± "
	NOT_SYM            string = " ¬ "
	NANO_SYM           string = " n "
	MILLI_SYM          string = " m "
	KILO_SYM           string = " k "
	MEGA_SYM           string = " M "
	GIGA_SYM           string = " G "
	TERA_SYM           string = " T "
	PICO_SYM           string = " p "
	FEMTO_SYM          string = " f "
	ATTO_SYM           string = " a "
	ZEPTO_SYM          string = " z "
	YOCTO_SYM          string = " y "
	SQUARED_SYM        string = " ² "
	CUBED_SYM          string = " ³ "
	POWER_6_SYM        string = " ⁶ "
	POWER_9_SYM        string = " ⁹ "
	POWER_12_SYM       string = " ¹² "
	SQUARE_ROOT_SYM    string = " √ "
	CUBE_ROOT_SYM      string = " ∛ "
	// set symbols
	EMPTY_SET_SYM      string = " ∅ "
	SUBSET_SYM         string = " ⊂ "
	SUBSET_EQUAL_SYM   string = " ⊆ "
	SUPERSET_SYM       string = " ⊃ "
	SUPERSET_EQUAL_SYM string = " ⊇ "
	INTERSECTION_SYM   string = " ∩ "
	UNION_SYM          string = " ∪ "
	// logic symbols
	EXISTS_SYM     string = " ∃ "
	FORALL_SYM     string = " ∀ "
	NOT_EXISTS_SYM string = " ∄ "
	NOT_FORALL_SYM string = " ∁ "
	IMPLIES_SYM    string = " ⇒ "
	IMPLIES2_SYM   string = " ⟹ "
	IMPLIES3_SYM   string = " ⟼ "
	IMPLIES4_SYM   string = " ⟾ "
	IMPLIES5_SYM   string = " ⟺ "
	IMPLIES6_SYM   string = " ⟸ "
	IMPLIES7_SYM   string = " ⇔ "
	IMPLIES8_SYM   string = " ⇎ "
	IMPLIES9_SYM   string = " ⇏ "
	IMPLIES10_SYM  string = " ⇍ "
	IMPLIES13_SYM  string = " ⇐ "
	IMPLIES15_SYM  string = " ⇌ "
)

const (
	// physics symbols
	OHM_SYM            string = " Ω "
	ANGSTROM_SYM       string = " Å "
	DEGREE_C_SYM       string = " ℃ "
	DEGREE_F_SYM       string = " ℉ "
	DEGREE_SYM         string = " ° "
	DEGREE_K_SYM       string = " K "
	DEGREE_R_SYM       string = " °R "
	DEGREE_REAUMUR_SYM string = " °Ré "
	DEGREE_RANKINE_SYM string = " °Ra "
	DEGREE_NEWTON_SYM  string = " °N "
	DEGREE_DELISLE_SYM string = " °De "
	DEGREE_ROMER_SYM   string = " °Rø "
	DEGREE_GALILEO_SYM string = " °G "
)

const (
	// chemistry symbols
	DEGREE_BAUME_SYM    string = " °Bé "
	DEGREE_TWADDELL_SYM string = " °Tw "
)

const (
	// greek symbols
	ALPHA_SYM   string = " α "
	BETA_SYM    string = " β "
	GAMMA_SYM   string = " γ "
	DELTA_SYM   string = " δ "
	EPSILON_SYM string = " ε "
	ZETA_SYM    string = " ζ "
	ETA_SYM     string = " η "
	THETA_SYM   string = " θ "
	IOTA_SYM    string = " ι "
	KAPPA_SYM   string = " κ "
	LAMBDA_SYM  string = " λ "
	MU_SYM      string = " μ "
	NU_SYM      string = " ν "
	XI_SYM      string = " ξ "
	OMICRON_SYM string = " ο "
	PI_SYM2     string = " π "
	RHO_SYM     string = " ρ "
	SIGMA_SYM   string = " σ "
	TAU_SYM     string = " τ "
	UPSILON_SYM string = " υ "
	PHI_SYM     string = " φ "
	CHI_SYM     string = " χ "
	PSI_SYM     string = " ψ "
	OMEGA_SYM   string = " ω "
)

const (
	// geometric symbols
	DOT_SYM        string = " • "
	CIRCLE_SYM     string = " ○ "
	SQUARE_SYM     string = " ■ "
	TRIANGLE_SYM   string = " △ "
	QUADRANGLE_SYM string = " □ "
	PENTAGON_SYM   string = " ⬟ "
	HEXAGON_SYM    string = " ⬡ "
	HEPTAGON_SYM   string = " ⬢ "
	OCTAGON_SYM    string = " ⬣ "
	ENNEAGON_SYM   string = " ⬤ "
	DECAGON_SYM    string = " ⬥ "
	HEXAGON2_SYM   string = " ⬦ "
	HEPTAGON2_SYM  string = " ⬧ "
	OCTAGON2_SYM   string = " ⬨ "
	ENNEAGON2_SYM  string = " ⬩ "
	DECAGON2_SYM   string = " ⬪ "
	HEXAGON3_SYM   string = " ⬫ "
	HEPTAGON3_SYM  string = " ⬬ "
	OCTAGON3_SYM   string = " ⬭ "
	ENNEAGON3_SYM  string = " ⬮ "
	DECAGON3_SYM   string = " ⬯ "
	HEXAGON4_SYM   string = " ⬰ "
	HEPTAGON4_SYM  string = " ⬱ "
	OCTAGON4_SYM   string = " ⬲ "
	ENNEAGON4_SYM  string = " ⬳ "
	DECAGON4_SYM   string = " ⬴ "
)

const (
	// currency symbols
	POUND_SYM     string = " £ "
	YEN_SYM       string = " ¥ "
	EURO_SYM      string = " € "
	CENT_SYM      string = " ¢ "
	DOLLAR_SYM    string = " $ "
	CURRENCY_SYM  string = " ¤ "
	PERMILLE_SYM  string = " ‰ "
	PERMYRIAD_SYM string = " ‱ "
	VIETNAM_SYM   string = " ₫ "
	ROUBLE_SYM    string = " ₽ "
	BAHT_SYM      string = " ฿ "
	INDIAN_SYM    string = " ₹ "
	RIAL_SYM      string = " ﷼ "
	LIRA_SYM      string = " ₤ "
	DRACHMA_SYM   string = " ₯ "
	NAIRA_SYM     string = " ₦ "
	PESETA_SYM    string = " ₧ "
	RUPEE_SYM     string = " ₨ "
	FRANC_SYM     string = " ₣ "
	LEU_SYM       string = " ₗ "
	AFGHANI_SYM   string = " ؋ "
)

const (
	// card symbols
	CLUB_SYM    string = " ♣ "
	DIAMOND_SYM string = " ♦ "
	HEART_SYM   string = " ♥ "
	SPADE_SYM   string = " ♠ "
)

const (
	// chess symbols
	WHITE_KING_SYM   string = " ♔ "
	WHITE_QUEEN_SYM  string = " ♕ "
	WHITE_ROOK_SYM   string = " ♖ "
	WHITE_BISHOP_SYM string = " ♗ "
	WHITE_KNIGHT_SYM string = " ♘ "
	WHITE_PAWN_SYM   string = " ♙ "
	BLACK_KING_SYM   string = " ♚ "
	BLACK_QUEEN_SYM  string = " ♛ "
	BLACK_ROOK_SYM   string = " ♜ "
	BLACK_BISHOP_SYM string = " ♝ "
	BLACK_KNIGHT_SYM string = " ♞ "
	BLACK_PAWN_SYM   string = " ♟ "
)

const (
	// weather symbols
	SUNNY_SYM              string = " ☀ "
	CLOUDY_SYM             string = " ☁ "
	RAIN_SYM               string = " ☂ "
	SNOW_SYM               string = " ☃ "
	THUNDERSTORM_SYM       string = " ☇ "
	SUNNY_CLOUDY_SYM       string = " ☼ "
	RAIN_CLOUDY_SYM        string = " ☂☁ "
	SNOW_CLOUDY_SYM        string = " ☃☁ "
	THUNDERSTORM_CLOUDY    string = " ☇☁ "
	SUNNY_RAIN_SYM         string = " ☀☂ "
	SUNNY_SNOW_SYM         string = " ☀☃ "
	SUNNY_THUNDERSTORM_SYM string = " ☀☇ "
	CLOUDY_RAIN_SYM        string = " ☁☂ "
	CLOUDY_SNOW_SYM        string = " ☁☃ "
	CLOUDY_THUNDERSTORM    string = " ☁☇ "
	RAIN_SNOW_SYM          string = " ☂☃ "
	RAIN_THUNDERSTORM_SYM  string = " ☂☇ "
	SNOW_THUNDERSTORM_SYM  string = " ☃☇ "
)

const (
	// astrological symbols
	SUN_SYM         string = " ☉ "
	MOON_SYM        string = " ☽ "
	MERCURY_SYM     string = " ☿ "
	VENUS_SYM       string = " ♀ "
	EARTH_SYM       string = " ⊕ "
	MARS_SYM        string = " ♂ "
	JUPITER_SYM     string = " ♃ "
	SATURN_SYM      string = " ♄ "
	URANUS_SYM      string = " ♅ "
	NEPTUNE_SYM     string = " ♆ "
	PLUTO_SYM       string = " ♇ "
	CONJUNCTION_SYM string = " ☌ "
	OPPOSITION_SYM  string = " ☍ "
	SEXTILE_SYM     string = " ⚹ "
	TRINE_SYM       string = " ⚸ "
	ASC_NODE_SYM    string = " ☊ "
	DESC_NODE_SYM   string = " ☋ "
	BLACK_MOON_SYM  string = " ☾ "
)

const (
	// religious symbols
	CHRISTIAN_SYM string = " ✝ "
	JEWISH_SYM    string = " ✡ "
	ISLAMIC_SYM   string = " ☪ "
	YIN_YANG_SYM  string = " ☯ "
)