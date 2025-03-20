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
		connectionSample()
		consoleSample()
		cronSample()
		cryptoSample()
		dateSample()
		functionSample()
		helperSample()
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
	} else {
		print("Please enable isDev flag to run samples")
	}
}

func apiSample() {
	samples.Loadbalancer()
	samples.GracefulShutdown()
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
	samples.EncodeData()
	samples.DecodeData()
	samples.EncryptPassword()
	samples.ComparePassword()
	samples.EncryptFile()
	samples.DecryptFile()
}

func dateSample() {
	samples.TimeElement()

	samples.DateISO()
	samples.DateShort()
	samples.DateLong()
	samples.DateUTC()

	samples.ToSecond()
	samples.ToMinute()
	samples.ToHour()
	samples.SleepSecond()
	samples.SleepMinute()
	samples.SleepHour()

	samples.Clock()
}

func functionSample() {
	samples.FunctionDelay()
	samples.FunctionRetry()
	samples.FunctionRemain()
	samples.FunctionInterval()
	samples.FunctionLock()
	samples.FunctionRunTime()
}

func helperSample() {
	samples.UUID()
	samples.AutoMarkdownDocument()
}

func loggerSample() {
	samples.WriteLogDefault()
	samples.WriteLogAsGouse()
	samples.WriteLogWithType()
	samples.AutoRotateTruncateLog()
}

func ioSample() {
	samples.CheckDir()
	samples.CreateDir()
	samples.RemoveDir()
	samples.LsDir()
	samples.HierarchyDir()
	samples.CurrentDir()

	samples.CheckFile()
	samples.CreateFile()
	samples.RemoveFile()
	samples.ReadFileByLine()
	samples.FileInfo()
	samples.RenameFile()
	samples.CopyFile()
	samples.TruncateFile()
	samples.CleanFile()
	samples.WriteToFile()
	samples.AppendToFile()
	samples.FileObj()

	samples.CreatePath()
	samples.ReadPath()
	samples.WritePath()

	samples.Zip()
	samples.Unzip()
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
	samples.Canvas()
	samples.PngToJpg()
}

func netSample() {
	samples.OpenURL()
	samples.EncodeURL()
	samples.DecodeURL()
	samples.CheckURL()
	samples.HeaderURL()
	samples.ConnectTime()
	samples.Proxy()
	samples.PortScanner()
	samples.PortChecker()
}

func numberSample() {
	samples.Clamp()
	samples.CheckInRange()
}

func osSample() {
	samples.System()
	samples.Disk()
	samples.Cpu()
	samples.Profile()
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
	samples.Capitalize()
	samples.CamelCase()
	samples.SnakeCase()
	samples.KebabCase()

	samples.IsLetter()
	samples.IsDigit()
	samples.Includes()
	samples.IsLower()
	samples.IsUpper()

	samples.Split()
	samples.Words()
	samples.Reverse()
	samples.Lower()
	samples.Upper()
	samples.Repeat()
	samples.Truncate()
	samples.Replace()
	samples.Trim()
	samples.TrimBlank()
	samples.TrimPrefix()
	samples.TrimSuffix()
	samples.Join()
	samples.Concat()
	samples.SubStr()
	samples.Slice()
	samples.Splice()
	samples.StartsWith()
	samples.EndsWith()
	samples.Escape()
	samples.Unescape()
	samples.Pad()

	samples.Count()
	samples.Lines()
	samples.Index()
	samples.RandomChain()
	samples.At()
	samples.CodePointAt()
	samples.CodePoint()
	samples.FromCodePointAt()
	samples.FromCodePoint()
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

func timeSample() {
	samples.TimeElement()

	samples.DateISO()
	samples.DateShort()
	samples.DateLong()
	samples.DateUTC()

	samples.ToSecond()
	samples.ToMinute()
	samples.ToHour()
	samples.SleepSecond()
	samples.SleepMinute()
	samples.SleepHour()

	samples.Clock()
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
