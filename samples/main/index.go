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
		configSample()
		consoleSample()
		cronSample()
		cryptoSample()
		dateSample()
		functionSample()
		helperSample()
		ioSample()
		mathSample()
		mediaSample()
		netSample()
		numberSample()
		osSample()
		randomSample()
		regexSample()
		stringSample()
		structSample()
		typeSample()
	} else {
		print("Please enable isDev flag to run samples")
	}
}

func apiSample() {
	samples.ApiLoadbalancer()
	samples.ApiGracefulShutdown()
}

func arraySample() {
	samples.ArrayEqual()
	samples.ArrayIncludes()
	samples.ArrayMost()
	samples.ArraySum()
	samples.ArrayChunk()
	samples.ArrayDifference()
	samples.ArrayDrop()
	samples.ArrayIndex()
	samples.ArrayMerge()
	samples.ArrayCompact()

	samples.ArrayIntersect()
	samples.ArrayMin()
	samples.ArrayMax()

	samples.ArrayKeyBy()
	samples.ArrayIndexBy()
	samples.ArrayFilterBy()
	samples.ArrayRejectBy()
	samples.ArrayFindBy()
	samples.ArrayForBy()
	samples.ArrayMapBy()
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

func configSample() {
	samples.ConfigJson()
	samples.ConfigYaml()
	samples.ConfigToml()
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
	samples.CryptoEncode()
	samples.CryptoDecode()
	samples.CryptoEncryptPassword()
	samples.CryptoDecryptPassword()
	samples.CryptoEncryptFile()
	samples.CryptoDecryptFile()
}

func dateSample() {
	samples.DateTime()

	samples.DateISO()
	samples.DateShort()
	samples.DateLong()
	samples.DateUTC()

	samples.DateToSecond()
	samples.DateToMinute()
	samples.DateToHour()
	samples.DateSleepSecond()
	samples.DateSleepMinute()
	samples.DateSleepHour()

	samples.DateClock()
}

func functionSample() {
	samples.FuncDelay()
	samples.FuncRetry()
	samples.FuncRemain()
	samples.FuncInterval()
	samples.FuncLock()
	samples.FuncRunTime()
}

func helperSample() {
	samples.HelperRandomID()
	samples.HelperUUID()

	samples.HelperAutoMdDoc()
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

func mathSample() {
	samples.MathIsPrime()
	samples.MathIsEven()
	samples.MathIsOdd()
	samples.MathIsPerfectSquare()

	samples.MathAbs()
	samples.MathFloor()
	samples.MathCeil()
	samples.MathRound()
	samples.MathMin()
	samples.MathMax()
	samples.MathSum()
	samples.MathMean()
	samples.MathOperators()
	samples.MathPower()
	samples.MathFactorial()
	samples.MathRoot()

	samples.MathLog()
	samples.MathPytago()
	samples.MathTrigonometry()
	samples.MathTransition()

	samples.MathRect()
	samples.MathCircle()
	samples.MathTriangle()
	samples.MathSquare()
	samples.MathCube()
	samples.MathSphere()
	samples.MathCylinder()
	samples.MathCone()
	samples.MathTrapezoid()
	samples.MathParallelogram()
	samples.MathRhombus()
	samples.MathEllipse()
	samples.MathPolygon()
}

func mediaSample() {
	samples.MediaCanvas()
	samples.MediaPngToJpg()
}

func netSample() {
	samples.NetOpen()
	samples.NetEncode()
	samples.NetDecode()
	samples.NetCheck()
	samples.NetCheckWithStatusCode()
	samples.NetHeader()
	samples.NetConnectTime()
	samples.NetProxy()
	samples.ApiPortScanner()
	samples.ApiPortChecker()
}

func numberSample() {
	samples.NumRandom()
	samples.NumClamp()
	samples.NumInRange()
}

func osSample() {
	samples.OsSystem()
	samples.OsDisk()
	samples.OsCpu()
	samples.OsProfile()
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

	samples.StringIsLetter()
	samples.StringIsDigit()
	samples.StringIncludes()
	samples.StringIsLower()
	samples.StringIsUpper()

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
	samples.StringTrimPrefix()
	samples.StringTrimSuffix()
	samples.StringJoin()
	samples.StringConcat()
	samples.StringSubStr()
	samples.StringSlice()
	samples.StringSplice()
	samples.StringStartsWith()
	samples.StringEndsWith()
	samples.StringEscape()
	samples.StringUnescape()
	samples.StringPad()

	samples.StringCount()
	samples.StringLines()
	samples.StringIndex()
	samples.StringRandom()
	samples.StringAt()
	samples.StringCodePointAt()
	samples.StringCodePoint()
	samples.StringFromCodePointAt()
	samples.StringFromCodePoint()
}

func structSample() {
	samples.StructMerge()
	samples.StructRemove()
	samples.StructAdd()
	samples.StructSet()
	samples.StructGet()

	samples.StructClone()
	samples.StructHas()
}

func typeSample() {
	samples.TypeStringConvert()
	samples.TypeStructConvert()
	samples.TypeCastToString()

	samples.TypeCheck()
	samples.TypeCheckUUID()
	samples.TypeCheckGmail()
	samples.TypeCheckYahoo()
	samples.TypeCheckOutlook()
	samples.TypeCheckEdu()
	samples.TypeCheckEmail()
	samples.TypeCheckUsername()
	samples.TypeCheckPassword()
	samples.TypeCheckPhone()
}
