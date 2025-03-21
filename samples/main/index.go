package main

import (
	"flag"

	"github.com/thuongtruong109/gouse/samples"
)

func main() {
	isDevFlag := flag.Bool("isDev", false, "toggle environment mode")
	flag.Parse()
	if *isDevFlag {
		apiSample()
		arraySample()
		cacheSample()
		chartSample()
		colorSample()
		configSample()
		connectionSample()
		consoleSample()
		cronSample()
		cryptoSample()
		dateSample()
		functionSample()
		ioSample()
		loggerSample()
		mathSample()
		mediaSample()
		netSample()
		numberSample()
		osSample()
		randomSample()
		regexSample()
		stringSample()
		structSample()
		timeSample()
		typeSample()
		utilsSample()
	} else {
		print("Please enable isDev flag to run samples")
	}
}

func apiSample() {
	samples.Loadbalancer()
	samples.GracefulShutdown()
	// samples.Validate()
}

func arraySample() {
	samples.ArrayMin()
	samples.ArrayMax()
	samples.ArrayIntersect()
	samples.ArrayIncludes()
	samples.ArrayEqual()
	samples.ArrayUnique()
	samples.ArrayMost()
	samples.ArrayLeast()
	samples.ArrayChunk()
	samples.ArrayDifference()
	samples.ArrayDrop()
	samples.ArrayFill()
	samples.ArrayShift()
	samples.ArrayUnshift()
	samples.ArrayPop()
	samples.ArrayPush()
	samples.ArraySlice()
	samples.ArraySplice()
	samples.ArrayTake()
	samples.ArrayReverse()
	samples.ArrayIndexOf()
	samples.ArrayMerge()
	samples.ArrayCompact()
	samples.ArrayFlatten()
	samples.ArraySort()
	samples.ArrayShuffle()
	samples.ArrayFilterBy()
	samples.ArrayFindBy()
	samples.ArrayForBy()
	samples.ArrayIndexBy()
	samples.ArrayKeyBy()
	samples.ArrayMapBy()
	samples.ArrayRejectBy()
}

func cacheSample() {
	samples.CacheLocal()
	samples.CacheTmp()
}

func chartSample() {
	samples.ChartBar()
	samples.ChartLine()
	samples.ChartPie()
}

func colorSample() {
	// samples.ColorHueToRGB()
	// samples.ColorHSLAToRGBA()
	// samples.ColorRGBAToHSLA()
	// samples.ColorHexToHSLA()
	// samples.ColorHexToRGBA()
}

func configSample() {
	samples.ConfigJson()
	samples.ConfigYaml()
	samples.ConfigToml()
}

func connectionSample() {
	samples.ConnectRedis()
	samples.ConnectRedisUri()
	samples.ConnectPostgres()
	samples.ConnectMongo()
	samples.ConnectMinio()
}

func consoleSample() {
	samples.ConsoleCmd()
	samples.ConsoleClear()
	samples.ConsoleWithColor()
	samples.ConsoleBanner()
	samples.ConsoleSelect()
	samples.ConsoleHelp()

	samples.ConsoleList()
	// samples.ConsolePaper()
	samples.ConsoleProgress()
	samples.ConsoleRealtime()
	samples.ConsoleChoice()
	samples.ConsoleSpinner()
	samples.ConsoleSplit()
	samples.ConsoleStopwatch()
	samples.ConsoleTable()
	// samples.ConsoleTab()
	samples.ConsoleCountdown()
	samples.ConsoleSequence()
	samples.ConsoleInline()
	samples.ConsoleParallel()
	samples.ConsoleDir()
	samples.ConsoleGlamour()
}

func cronSample() {
	samples.CronRun()
}

func cryptoSample() {
	samples.CryptoEncodeData()
	samples.CryptoDecodeData()
	samples.CryptoEncryptPassword()
	samples.CryptoComparePassword()
	samples.CryptoEncryptFile()
	samples.CryptoDecryptFile()
}

func dateSample() {
	samples.DateISO()
	samples.DateShort()
	samples.DateLong()
	samples.DateUTC()
}

func functionSample() {
	samples.FunctionDelay()
	samples.FunctionRetry()
	samples.FunctionRemain()
	samples.FunctionInterval()
	samples.FunctionLock()
	samples.FunctionRunTime()
	samples.FunctionDeferWrapper()
	samples.FunctionParallel()
}

func ioSample() {
	samples.IoCheckDir()
	samples.IoCreateDir()
	samples.IoRemoveDir()
	samples.IoLsDir()
	samples.IoHierarchyDir()
	samples.IoCurrentDir()

	samples.IoCheckFile()
	samples.IoCreateFile()
	samples.IoRemoveFile()
	samples.IoReadFileByLine()
	samples.IoFileInfo()
	samples.IoRenameFile()
	samples.IoCopyFile()
	samples.IoTruncateFile()
	samples.IoCleanFile()
	samples.IoWriteToFile()
	samples.IoAppendToFile()
	samples.IoFileObj()

	samples.IoCreatePath()
	samples.IoReadPath()
	samples.IoWritePath()

	samples.IoZip()
	samples.IoUnzip()
}

func loggerSample() {
	samples.LoggerWriteDefault()
	samples.LoggerWriteAsGouse()
	samples.LoggerWriteWithType()
	samples.LoggerAutoRotateTruncate()
}

func mathSample() {
	samples.CheckNumber()
	samples.Abs()
	samples.Floor()
	samples.Ceil()
	samples.Round()
	samples.MinMaxMean()
	samples.Sum()
	samples.Operators()
	samples.Power()
	samples.Factorial()
	samples.Root()
	samples.Log()
	samples.Pytago()
	samples.Trigonometry()
	samples.Transition()
	samples.Rect()
	samples.Circle()
	samples.Triangle()
	samples.Square()
	samples.Cube()
	samples.Sphere()
	samples.Cylinder()
	samples.Cone()
	samples.Trapezoid()
	samples.Parallelogram()
	samples.Rhombus()
	samples.Ellipse()
	samples.Polygon()
}

func mediaSample() {
	samples.MediaCanvas()
	// samples.MediaPNGToJPG()
}

func netSample() {
	samples.NetOpenURL()
	samples.NetEncodeURL()
	samples.NetDecodeURL()
	samples.NetCheckURL()
	samples.NetHeaderURL()
	samples.NetConnectTime()
	samples.NetProxy()
	samples.NetPortScanner()
	samples.NetPortChecker()
}

func numberSample() {
	samples.NumberClamp()
	samples.NumberCheckInRange()
	// samples.NumberSortNum()
}

func osSample() {
	// samples.OsIo()
	samples.OsDisk()
	samples.OsCpu()
	// samples.OsMemory()
	// samples.OsUser()
	samples.OsSystem()
}

func randomSample() {
	samples.RandomNumber()
	samples.RandomID()
	samples.RandomString()
	samples.RandomDigit()
	samples.RandomUUID()
}

func regexSample() {
	samples.RegexIsMatch()
	samples.RegexMatch()
	samples.RegexMatchIndex()
}

func stringSample() {
	samples.StringCapitalize()
	samples.StringCamelCase()
	samples.StringSnakeCase()
	samples.StringKebabCase()
	samples.StringSpaceCase()
	samples.StringCustomCase()

	samples.StringIsLetter()
	samples.StringIsDigit()
	samples.StringIncludes()
	samples.StringIsLower()
	samples.StringIsUpper()
	samples.StringStartsWith()
	samples.StringEndsWith()

	samples.StringSplit()
	samples.StringWords()
	samples.StringReverse()
	samples.StringLower()
	samples.StringUpper()
	samples.StringRepeat()
	samples.StringTruncate()
	samples.StringReplace()
	samples.StringTrim()
	samples.StringTrimBlank()
	samples.StringJoin()
	samples.StringConcat()
	samples.StringSubStr()
	samples.StringSlice()
	samples.StringSplice()
	samples.StringEscape()
	samples.StringUnescape()
	samples.StringPad()
	samples.StringCount()
	samples.StringLines()
	samples.StringIndex()
	samples.StringAt()
	samples.StringCodePointAt()
	samples.StringCodePoint()
	samples.StringFromCodePointAt()
	samples.StringFromCodePoint()
}

func structSample() {
	// samples.StructGetTag()
	samples.StructMerge()
	samples.StructRemove()
	samples.StructAdd()
	samples.StructSet()
	samples.StructGet()
	samples.StructClone()
	samples.StructHas()
}

func timeSample() {
	samples.TimeElement()
	// samples.TimeToSecond()
	// samples.TimeToMinute()
	// samples.TimeToHour()
	// samples.TimeSleepSecond()
	// samples.TimeSleepMinute()
	// samples.TimeSleepHour()
	// samples.TimeClock()
	// samples.TimeDiff()
	// samples.TimeDiffNow()
}

func typeSample() {
	samples.TypeCheck()
	samples.TypeStringConvert()
	samples.TypeStructConvert()
	samples.TypeCastToString()
	samples.TypeCheckUUID()
	samples.TypeCheckGmail()
	samples.TypeCheckYahoo()
	samples.TypeCheckOutlook()
	samples.TypeCheckEdu()
	samples.TypeCheckEmail()
	samples.TypeCheckUsername()
	samples.TypeCheckPassword()
	samples.TypeCheckPhone()
	// samples.TypeTimeConvert()
}

func utilsSample() {
	// samples.SetCtx()
	// samples.GetCtx()
	// samples.DetectError()
}
